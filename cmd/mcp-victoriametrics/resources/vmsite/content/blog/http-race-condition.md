---
page: blog blog_post
authors:
  - Jesús Espino
date: 2026-10-06
title: "The Data Race That Wasn't a Bug (and the One That Was)"
summary: "You turn on Go's race detector and get a data race inside net/http that makes no sense at first. This post digs into how Go's HTTP client really sends a request body, why it can keep reading your buffer after Client.Do has returned, why we decided that race is harmless in vmagent, and why the same behavior could corrupt data in vmauth, so we fixed it there."
enableComments: true
toc: true
categories:
  - Open Source Tech
  - VictoriaMetrics
tags:
  - go
  - golang
  - net/http
  - data race
  - vmagent
  - vmauth
  - open source
images:
  - /blog/http-race-condition/preview.webp
---

Imagine this: you are testing the performance of some part of your application. Everything is going smoothly, the numbers look good, and as a last check you turn on Go's race detector. Then, out of nowhere, it prints a warning you didn't expect:

```
WARNING: DATA RACE
```

So you look at it. You look at it again, and again, and you think: *"What the...?"* The race is between your code and a goroutine you never started, somewhere deep inside `net/http`. You have no idea how that is possible, or why. Time to dig in.

That is more or less what happened to [Vadim Alekseev](https://x.com/vadimaleksv) while running [vmagent](https://docs.victoriametrics.com/victoriametrics/vmagent/) under the race detector. He tracked the behavior down, reduced it to a small program, and reported it upstream as [golang/go#81445](https://github.com/golang/go/issues/81445). The same behavior also showed up in [vmauth](https://docs.victoriametrics.com/victoriametrics/vmauth/). In the end we left one of those races in place on purpose, and fixed the other. To see why, we first need to understand what the race actually is.

> [!NOTE]
> The Go internals in this post (function names, buffer sizes, timeouts) come from the Go 1.27 `net/http` source. The same mechanics have been there for many releases.

## The code and the race

Here is the client side of the program from [the upstream issue](https://github.com/golang/go/issues/81445) (the full runnable version is there). It sends a batch of logs to a server, and retries if the server answers with an error. To avoid allocating a new buffer for every attempt, it reuses a single global `bytes.Buffer`:

```go
type LogEntry struct {
 Timestamp int64
 Content   string
}

var buf = bytes.NewBuffer(nil)

func insertLogs(client *http.Client, logs []LogEntry) error {
 buf.Reset()
 if err := json.MarshalWrite(buf, logs); err != nil {
  panic(err)
 }

 body := bytes.NewReader(buf.Bytes())
 resp, err := client.Post("http://example.com/insert/logs", "application/json", body)
 if err != nil {
  return err
 }
 defer resp.Body.Close()

 _, _ = io.Copy(io.Discard, resp.Body)
 if resp.StatusCode/100 != 2 {
  return fmt.Errorf("unexpected status code %d", resp.StatusCode)
 }

 return nil
}

// ...and in main, retry until it works:
for range 100 {
 err := insertLogs(client, logs)
 if err != nil {
  fmt.Println("error inserting logs:", err)
  continue
 }
 break
}
```

`logs` is 1024 entries with 256 bytes of content each, so the body is about 300 KB. On the other side there's a small test server that always answers `502 Bad Gateway`, without even looking at the request body. The program sends, gets a `502`, and tries again. Nothing unusual. Run it with `-race` and, after a few attempts:

```
error inserting logs: unexpected status code 502
error inserting logs: unexpected status code 502
...
==================
WARNING: DATA RACE
Write at 0x00c000390f60 by main goroutine:
  runtime.slicecopy()
  encoding/json/v2.makeStructArshaler.func2()
  ...
  encoding/json/v2.MarshalWrite()
  main.insertLogs()
  main.main()

Previous read at 0x00c000390f60 by goroutine 49:
  runtime.slicecopy()
  bytes.(*Reader).Read()
      bytes/reader.go:44
  io.(*LimitedReader).Read()
  ...
  io.Copy()
  net.genericReadFrom()
  net.(*TCPConn).ReadFrom()
  ...
  net/http.persistConnWriter.ReadFrom()
  bufio.(*Writer).ReadFrom()
  ...
  net/http.(*transferWriter).doBodyCopy()
  net/http.(*transferWriter).writeBody()
  net/http.(*Request).write()
  net/http.(*persistConn).writeLoop()
==================
exit status 66
```

The report describes two accesses to the same memory address, `0x00c000390f60`:

- **The write** happens in the `main` goroutine, inside `insertLogs`. It's `json.MarshalWrite` writing the JSON for a new attempt into `buf`.
- **The read** happens in another goroutine, one we never started. It's a `bytes.Reader` reading `buf`, called from a function named `writeLoop` inside `net/http`.

So something inside `net/http` was reading our buffer while our code was writing the next attempt into it. That's surprising: by the time we write, `client.Post` for the previous attempt has already returned, and we have read and closed its response.

To understand how that can happen, we need to take a step back and look at how Go's HTTP client really sends a request.

## How the HTTP client sends a request: three goroutines

When you call `client.Post`, it feels like a single operation: send the request, get the response. Under the hood, at least **three goroutines** are involved.

For every HTTP/1.1 connection it opens, the `http.Transport` creates a `persistConn` with two goroutines of its own:

- **`writeLoop`** sends requests over the connection: the request line, the headers, then the body.
- **`readLoop`** reads responses from the connection and delivers them.

The third goroutine is **yours**, the one that called `client.Post`.

So `writeLoop` and `readLoop` are already running, one pair per connection, waiting for work. When your goroutine calls `client.Post`, it goes down through the client and the transport until it reaches `persistConn.roundTrip`, the function that connects your request with those two goroutines. It does two things.

First, it gives the request to **both** goroutines, one message on a channel to each:

```go
// net/http/transport.go (simplified)

// Write the request concurrently with waiting for a response,
// in case the server decides to reply before reading our full
// request body.
pc.writech <- writeRequest{req, writeErrCh, continueCh} // writeLoop: "send this"
pc.reqch <- requestAndChan{treq: req, ch: resc, ...}    // readLoop: "wait for its answer"
```

From this point on, `writeLoop` is sending the request and `readLoop` is waiting for the response, both at the same time.

Second, your goroutine waits to see what happens: either `writeLoop` reports that it finished writing the request, or `readLoop` delivers a response:

```go
for {
 select {
 case err := <-writeErrCh: // writeLoop finished writing the request
  ...
 case re := <-resc: // readLoop got a response
  return handleResponse(re)
 ...
 }
}
```

Let's see what this looks like visually:

![Behind a single client.Post call: your goroutine hands the request to writeLoop, which sends it, and to readLoop, which waits for the response, and returns as soon as readLoop delivers one.](/blog/http-race-condition/three-goroutines.webp)

On the left, your goroutine does the two channel sends we just saw: `writech` tells `writeLoop` to send the request, and `reqch` tells `readLoop` to wait for its answer. Then it parks in the `select`, waiting for whichever arrow comes back first.

On the right side there are two separate paths. The **request + body** arrow, from `writeLoop` to the server, is the request going out: first the request line and the headers, then the body, chunk after chunk. The arrow from the server back to `readLoop` is the response coming in. A TCP connection carries data in both directions at the same time, and each direction is independent. The response doesn't have to wait for the request to finish before it can travel back.

Typically, `writeLoop` sends the request line and the headers, and then starts sending the body. The server receives the request line and the headers, and based on them it calls the right handler. The handler reads and processes the body, and then sends a reply back, which `readLoop` picks up and passes to your goroutine along the **response** arrow.

But it isn't always like that. Sometimes the headers alone are enough for the server to make a decision. The request may be unauthorized, or its `Content-Length` may say the body is bigger than the server accepts, or the server may be a proxy that can't reach its backend. In those cases, the server can reply right away, without waiting for the body. The reply travels back on the other direction of the connection, `readLoop` passes it to your goroutine, your `select` takes the `resc` case, and `roundTrip` returns. So the client is still sending the body while the server has already sent its reply. That's the interesting case for us here.

And here's the important consequence. Your `select` returns as soon as the response arrives, and it doesn't wait for `writeLoop` to finish. So `client.Post` can return, and your code can move on with the response in hand, **while `writeLoop` is still running in the background, sending the rest of your body**.

Go does document this. The [`http.RoundTripper`](https://pkg.go.dev/net/http#RoundTripper) interface says:

> RoundTrip must always close the body, including on errors, but depending on the implementation may do so in a separate goroutine even after RoundTrip returns. This means that callers wanting to reuse the body for subsequent requests must arrange to wait for the Close call before doing so.

and [`Client.Do`](https://pkg.go.dev/net/http#Client.Do) warns that "the Body may be closed asynchronously after Do returns".

If we go back to our `DATA RACE` report, the goroutine doing the read was running `writeLoop`. So now we have the *who*: it's the goroutine that sends our request, and it can keep running after `client.Post` has returned. What we still don't know is *what* exactly it reads from our buffer, and when. For that, we have to follow the bytes.

## Following the bytes: your buffer, the reader, and the scratch buffer

Let's follow our JSON from the moment we write it until it leaves the machine:

![The path of the body: your buffer's array, copied in 32 KiB chunks into writeLoop's own buffer, written into the kernel's socket buffer, and sent over the network. Only the first copy touches your memory.](/blog/http-race-condition/body-bytes-path.webp)

It all starts inside the **your code** box, with **buf's array**. Remember that `buf` is a global variable: it lives for the whole program, so every call to `insertLogs` uses the same buffer, and the same array inside it, which stays in memory between calls. `buf.Reset()` doesn't throw that array away, it just marks it as empty, so every time we retry `insertLogs`, the new JSON goes into the same memory. That's the point of reusing `buf`: no new allocation on every retry. The `bytes.Reader` we pass to `client.Post` doesn't copy anything either: it reads straight from that same array.

To send the body, `writeLoop` follows the first arrow in our diagram: it copies 32 KiB chunks from our buffer into a buffer of its own, and sends them to the server through the network.

Then the arrow at the bottom takes us back to buf's array for the next chunk, and the whole thing **repeats until EOF or until the socket is closed**.

Now we have all the pieces. Let's put them together.

## Putting it together: how the race happens

Now let's put all the pieces on one timeline:

![Timeline of the race: the caller writes the buffer and calls Do, writeLoop copies it out in 32 KiB chunks, readLoop hands back the 502 without waiting for writeLoop, and the caller's retry overwrites the array while writeLoop may still be copying from it.](/blog/http-race-condition/race-timeline.webp)

Each row is one goroutine, and the bottom row is the memory they share: **buf's underlying array**. Let's follow it from left to right.

First, in **your code**, we call `buf.Reset()` and encode our logs, which writes the JSON into the array. Then we call `client.Do`, and our goroutine starts waiting.

That call sets the other two goroutines in motion, in parallel. `writeLoop` starts sending our request to the server: it copies the body out of the array, 32 KiB at a time, and writes each chunk to the socket. At the same time, `readLoop` starts waiting for the server to answer.

The answer comes back before the body is fully sent. The server only needed our headers to decide, so it replies `502` without reading the body. And because so much of the body is left unread, it adds **Connection: close**. `readLoop` hands that response back to us, and **Do returns 502**. For us, the request is done.

But `writeLoop` doesn't know the server already answered. It's **still reading!** from our array, sending the rest of the body.

Right then, our code retries: `buf.Reset()` and encode again, and **WRITE again** goes into the same array `writeLoop` is still reading. That's the **overlap window**: one goroutine writing new data into the array while another is still reading the old data from it. That's our data race.

Finally, because the server asked to close the connection, `readLoop` **closes the socket**. `writeLoop`'s next write fails, it closes our body, and it exits.

This race doesn't show up every time: locally the body is sent almost instantly, and if the connection is going to be reused, Go waits up to 50 ms for `writeLoop` before letting us continue. But what is the actual problem when the race does happen?

### What can actually go wrong

First, the good news: our code writes the array and `writeLoop` only reads it, so our buffer itself is never corrupted. And in our example, every retry encodes the same logs, so the bytes `writeLoop` reads are the same either way.

But let's do a small mental exercise. Imagine we reuse this buffer for *different* requests, like sending a new batch of logs each time. What would the leftover `writeLoop` send for the previous request?

Say the previous request was:

```
[{"Timestamp":1790000000111111111,"Content":"disk full"}]
```

and the next one, written into the same array while `writeLoop` is still copying, is:

```
[{"Timestamp":1790000000999999999,"Content":"disk ok!!"}]
```

If `writeLoop` had copied the first half when the new data landed, it sends the old first half and the new second half:

```
[{"Timestamp":1790000000111119999,"Content":"disk ok!!"}]
```

That's valid JSON, but with a timestamp that neither request ever had. Both payloads have the same shape, so the pieces fit together perfectly.

Now imagine the next request is shorter than the previous one. `writeLoop` still thinks the body has the old length, so it copies the new data and then keeps going into what's left of the old data, because `buf.Reset()` doesn't clear anything:

```
previous : [{"Timestamp":1,"Content":"first"},{"Timestamp":2,"Content":"second"}]
next     : [{"Timestamp":9,"Content":"new"}]
sent     : [{"Timestamp":9,"Content":"new"}]},{"Timestamp":2,"Content":"second"}]
```

This time it's not even valid JSON.

That sounds scary. But whether it matters depends on one question: **who, if anyone, reads that mixed copy?**

## vmagent: a real race we decided to keep

vmagent's remote write client ([#11507](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/11507)) is the same pattern at a bigger scale. A worker pulls a block of compressed samples from its queue into a byte slice that it **reuses** on every iteration, and sends it with a fresh reader:

```go
// app/vmagent/remotewrite/client.go (simplified)
func (c *client) runWorker(readBlock func(dst []byte) ([]byte, bool)) {
 var block []byte
 for {
  block, ok = readBlock(block[:0]) // <- overwrites the previous block's bytes
  ...
  c.sendBlock(block)
 }
}

func (c *client) newRequest(url string, body []byte) (*http.Request, error) {
 reqBody := bytes.NewBuffer(body) // <- a fresh reader for every request
 req, err := http.NewRequest(http.MethodPost, url, reqBody)
 ...
}
```

Point vmagent at a remote storage that answers `200` without reading the body (Vadim used `httpbin.org/status/200`), push a big import through it, and the race detector reports the queue writing the next block into `block` while `writeLoop` is still reading the previous one. It's exactly the race we just dissected. We even merged a fix for it. Three days later we [reverted it](https://github.com/VictoriaMetrics/VictoriaMetrics/commit/8e9af3911f045ed69ed150393806cbebd37b4cd3), and later closed the issue without a fix. Here's why.

### Why we left it alone

The key to understanding this is how we build the request body: with `bytes.NewReader` in our example, or `bytes.NewBuffer` in vmagent. Both wrap our array with their own length and read position, so that part isn't shared between requests. The only shared part is the array underneath. And the [Go memory model](https://go.dev/ref/mem#restrictions) guarantees that reading a byte while it's being overwritten gives us either the old value or the new one, never some corrupted in-between state. So our readers are safe: the data they send can be mixed, but the program itself won't break. Nothing crashes, and no other memory gets corrupted.

So the worst case is some mixed data, and in vmagent nobody uses it. It goes into a request the server has **already answered**, without reading the body, so the server didn't want it anyway. Usually the connection is being closed too, so nothing on the other side ever reads them.

With nothing to protect, fixing it anyway would only cost us. A new buffer per request would undo the savings of reusing it. And waiting for the transport to finish with the body, which we tried and then [reverted](https://github.com/VictoriaMetrics/VictoriaMetrics/commit/8e9af3911f045ed69ed150393806cbebd37b4cd3), could stall workers, deadlock in rare cases, and still didn't cover everything.

Paying in performance and complexity to protect bytes nobody reads isn't a good trade, so we kept the race.

> [!NOTE]
> "Harmless" is not "free". This is a known, accepted race-detector report, not a silent one. We'll revisit it if Go gains an official way to wait for the transport to be done with a request body. In [the upstream issue](https://github.com/golang/go/issues/81445), Damien Neil suggested a new `Request.Close` method for exactly that.

## vmauth: when the race is a real bug

vmauth hit the same behavior ([#11508](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/11508)), but here it was a real problem. vmauth is a proxy, and it can **retry**: if a backend fails or answers with an error like `503`, vmauth sends the same request to the next backend.

To send the same body twice, vmauth keeps it in memory in a type called `bufferedBody`. Before the fix, `bufferedBody` was also the reader: it held the bytes plus its own read position. On a retry, vmauth rewound that position to zero and handed **the same `bufferedBody`** to the next request.

That's the key difference from vmagent. In vmagent, each request has its own reader, and only the bytes are shared. In vmauth, both requests shared **the reader itself**, including its read position. So the leftover `writeLoop` from the failed request could keep moving that position, and when it finished, it even reset it to zero, right in the middle of the retry's upload.

The result: the next backend could receive the body with pieces missing or repeated. Sometimes the length came out wrong and the request failed. But sometimes it came out exactly right, and the backend accepted a corrupted request without anyone noticing. This time the mixed data doesn't go nowhere: it goes to a healthy backend that processes it.

### The fix: correctness first

The fix ([#11647](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/11647)) does what vmagent already does: **never give two requests the same reader.** vmauth still keeps the bytes in `bufferedBody`, but every attempt now gets its own new `bytes.Buffer` over them. The leftover `writeLoop` can keep reading its own reader as long as it wants, and it can't touch the retry's. In the case of vmauth, the shared bytes are only ever read, never written, so there's no race at all.

Here, correctness wins easily: a proxy that can silently send corrupted data to a backend isn't acceptable at any speed. And the cost is tiny: a couple of small allocations per attempt, with no copy of the data, on a path where the network round trip to the backend costs far more.

So the same race detector warning led to two opposite decisions: keep it in vmagent, fix it in vmauth. Let's wrap up with what made the difference.

## What to take away

Both warnings came from the same `net/http` behavior: the transport writes the request body in its own goroutine, and it can return the response to you before it's done reading your body. The race detector was right both times. What it can't tell you is whether the race matters, and that came down to two questions:

1. **What exactly is shared?** Plain bytes that get copied somewhere, or *state* that decides behavior, like an offset, a length or a pointer? A race on bytes gives you stale bytes. A race on a cursor gives you wrong behavior.
2. **Who consumes the result?** In vmagent, the mixed copy goes into a request the server has already answered, on a connection that's being closed. In vmauth, the racing cursor decided what a *healthy backend* received and ingested.

So in vmagent we accepted the race and documented why, because fixing it would cost performance and complexity to protect bytes nobody reads. In vmauth we fixed it, by removing the sharing rather than adding synchronization, because correctness comes first.

If you use Go's HTTP client, the rules are short:

- Assume the transport may still be reading your request body after `Do` returns.
- Avoid giving the same stateful `io.Reader` to two requests.
- Be careful with reusing, pooling or mutating the memory behind a body you've already sent, until the transport has called `Close()` on it.

And keep an eye on [golang/go#81445](https://github.com/golang/go/issues/81445): if Go gets an official way to wait until the transport is done with a request, most of this goes away.
