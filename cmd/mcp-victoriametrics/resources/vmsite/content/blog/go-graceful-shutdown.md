---
draft: false
page: blog blog_post
authors:
  - Phuong Le
date: 2025-05-02
title: "Graceful Shutdown in Go: Practical Patterns"
summary: "Go applications can implement graceful shutdown by handling termination signals (SIGTERM, SIGINT) via os/signal or signal.NotifyContext. Shutdown must complete within a specified timeout (e.g., Kubernetes' terminationGracePeriodSeconds)..."
enableComments: true
toc: true
categories:
  - Go @ VictoriaMetrics
  - Open Source Tech
tags:
  - go
  - golang
  - http
  - graceful shutdown
  - kubernetes
images:
  - /blog/go-graceful-shutdown/preview.webp
---

Graceful shutdown in any application generally satisfies three minimum conditions:

1. Close the entry point by stopping new requests or messages from sources like HTTP, pub/sub systems, etc. However, keep outgoing connections to third-party services like databases or caches active.
2. Wait for all ongoing requests to finish. If a request takes too long, respond with a graceful error.
3. Release critical resources such as database connections, file locks, or network listeners. Do any final cleanup.

This article focuses on HTTP servers and containerized applications, but the core ideas apply to all types of applications.

For further discussion on graceful shutdown practices in Go, see the conversations on [Reddit](https://www.reddit.com/r/golang/comments/1kd6um8/graceful_shutdown_in_go_practical_patterns/) and [Hacker News](https://news.ycombinator.com/item?id=43889610).

## 1. Catching the Signal

Before we handle graceful shutdown, we first need to catch termination signals. These signals tell our application it's time to exit and begin the shutdown process.

So, what are signals?

In Unix-like systems, signals are software interrupts. They notify a process that something has happened and it should take action. When a signal is sent, the operating system interrupts the normal flow of the process to deliver the notification.

Here are a few possible behaviors:

- **Signal handler**: A process can register a handler (a function) for a specific signal. This function runs when that signal is received.
- **Default action**: If no handler is registered, the process follows the default behavior for that signal. This might mean terminating, stopping, continuing, or ignoring the process.
- **Unblockable signals**: Some signals, like `SIGKILL` (signal number 9), cannot be caught or ignored. They may terminate the process.

When your Go application starts, even before your `main` function runs, the Go runtime automatically registers signal handlers for many signals (`SIGTERM`, `SIGQUIT`, `SIGILL`, `SIGTRAP`, and others). However, for graceful shutdown, only three termination signals are typically important:

- `SIGTERM` (Termination): A standard and polite way to ask a process to terminate. It does not force the process to stop. Kubernetes sends this signal when it wants your application to exit before it forcibly kills it.
- `SIGINT` (Interrupt): Sent when the user wants to stop a process from the terminal, usually by pressing `Ctrl+C`.
- `SIGHUP` (Hang up): Originally used when a terminal disconnected. Now, it is often repurposed to signal an application to reload its configuration.

People mostly care about `SIGTERM` and `SIGINT`. `SIGHUP` is less used today for shutdown and more for reloading configs. You can find more about this in [SIGHUP Signal for Configuration Reloads](https://blog.devtrovert.com/p/sighup-signal-for-configuration-reloads).

By default, when your application receives a `SIGTERM`, `SIGINT`, or `SIGHUP`, the Go runtime will terminate the application.

> [!NOTE] Insight: How Go Terminates Your Application  
> When your Go app gets a `SIGTERM`, the runtime first catches it using a built-in handler. It checks if a custom handler is registered. If not, the runtime disables its own handler temporarily, and sends the same signal (`SIGTERM`) to the application again. This time, the OS handles it using the default behavior, which is to terminate the process.

You can override this by registering your own signal handler using the `os/signal` package.

```go
func main() {
  signalChan := make(chan os.Signal, 1)
  signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

  // Setup work here

  <-signalChan

  fmt.Println("Received termination signal, shutting down...")
}
```

![Graceful shutdown begins with signal setup](/blog/go-graceful-shutdown/signal-setup-before-init.webp)	
<figcaption style="text-align: center; font-style: italic;">Graceful shutdown begins with signal setup</figcaption>

`signal.Notify` tells the Go runtime to deliver specified signals to a channel instead of using the default behavior. This allows you to handle them manually and prevents the application from terminating automatically.

A buffered channel with a capacity of 1 is a good choice for reliable signal handling. Internally, Go sends signals to this channel using a `select` statement with a default case:

```go
select {
case c <- sig:
default:
}
```

This is different from the usual `select` used with receiving channels. When used for sending:

- If the buffer has space, the signal is sent and the code continues.
- If the buffer is full, the signal is discarded, and the `default` case runs. If you're using an unbuffered channel and no goroutine is actively receiving, the signal will be missed.

Even though it can only hold one signal, this buffered channel helps avoid missing that first signal while your app is still initializing and not yet listening.

> [!NOTE] Note
> You can call `Notify` multiple times for the same signal. Go will send that signal to all registered channels.

When you press `Ctrl+C` more than once, it doesn't automatically kill the app. The first `Ctrl+C` sends a `SIGINT` to the foreground process. Pressing it again usually sends another `SIGINT`, not `SIGKILL`. Most terminals, like bash or other Linux shells, do not escalate the signal automatically. If you want to force a stop, you must send `SIGKILL` manually using `kill -9`.

This is not ideal for local development, where you may want the second `Ctrl+C` to terminate the app forcefully. You can stop the app from listening to further signals by using `signal.Stop` right after the first signal is received:

```go
func main() {
  signalChan := make(chan os.Signal, 1)
  signal.Notify(signalChan, syscall.SIGINT)

  <-signalChan

  signal.Stop(signalChan)
  select {}
}
```

Starting with Go 1.16, you can simplify signal handling by using `signal.NotifyContext`, which ties signal handling to context cancellation:

```go
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
defer stop()

// Setup tasks here

<-ctx.Done()
stop()
```

You should still call `stop()` after `ctx.Done()` to allow a second `Ctrl+C` to forcefully terminate the application.

## 2. Timeout Awareness

It is important to know how long your application has to shut down after receiving a termination signal. For example, in Kubernetes, the default grace period is 30 seconds, unless otherwise specified using the `terminationGracePeriodSeconds` field. After this period, Kubernetes sends a `SIGKILL` to forcefully stop the application. This signal cannot be caught or handled.

Your shutdown logic must complete within this time, including processing any remaining requests and releasing resources.

Assume the default is 30 seconds. It is a good practice to reserve about 20 percent of the time as a safety margin to avoid being killed before cleanup finishes. This means aiming to finish everything within 25 seconds to avoid data loss or inconsistency.

## 3. Stop Accepting New Requests

When using `net/http`, you can handle graceful shutdown by calling the `http.Server.Shutdown` method. This method stops the server from accepting new connections and waits for all active requests to complete before shutting down idle connections.

Here is how it behaves:

- If a request is already in progress on an existing connection, the server will allow it to complete. After that, the connection is marked as idle and is closed.
- If a client tries to make a new connection during shutdown, it will fail because the server's listeners are already closed. This typically results in a "connection refused" error.

In a containerized environment (and many other orchestrated environments with external load balancers), do not stop accepting new requests immediately. Even after a pod is marked for termination, it might still receive traffic for a few moments.

Kubernetes internal components like `kube-proxy` are quickly aware of the change in pod status to "Terminating". They then prioritize routing **internal traffic** to `Ready,Serving` endpoints over `Terminating,Serving` ones.

The external load balancer, however, operates independently from Kubernetes. It typically uses its own health check mechanisms to determine which backend nodes should receive traffic. This health check indicates whether there are healthy (`Ready`) and non-terminating pods on the node. However, this check needs a little time to propagate.

There are 2 ways to handle this:

1. Use a `preStop` hook to sleep for a while, so the external load balancer has time to recognize that the pod is terminating.
    ```yaml
	lifecycle:
	preStop:
		exec:
		command: ["/bin/sh", "-c", "sleep 10"]
    ```
	And really importantly, the time taken by the `preStop` hook is included within the `terminationGracePeriodSeconds`.

2. Fail the readiness probe and sleep at the code level. This approach is not only applicable to Kubernetes environments, but also to other environments with load balancers that need to know the pod is not ready.

> [!NOTE] What is readiness probe?  
> A readiness probe determines when a container is prepared to accept traffic by periodically checking its health through configured methods like HTTP requests, TCP connections, or command executions. If the probe fails, Kubernetes removes the pod from the service's endpoints, preventing it from receiving traffic until it becomes ready again.

To avoid connection errors during this short window, the correct strategy is to fail the readiness probe first. This tells the orchestrator that your pod should no longer receive traffic:

```go
var isShuttingDown atomic.Bool

func readinessHandler(w http.ResponseWriter, r *http.Request) {
    if isShuttingDown.Load() {
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte("shutting down"))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
}
```

![Delay shutdown by failing readiness first](/blog/go-graceful-shutdown/fail-readiness-before-shutdown.webp)
<figcaption style="text-align: center; font-style: italic;">Delay shutdown by failing readiness first</figcaption>

This pattern is also used as a code example in the test images. In [their implementation](https://github.com/kubernetes/kubernetes/blob/95860cff1c418ea6f5494e4a6168e7acd1c390ec/test/images/agnhost/netexec/netexec.go#L357), a closed channel is used to signal the readiness probe to return HTTP 503 when the application is preparing to shut down.

After updating the readiness probe to indicate that the pod is no longer ready, wait a few seconds to allow the system time to propagate the change.

The exact wait time depends on your readiness probe configuration; we will use 5 seconds for this article with the following simple configuration:

```yaml
readinessProbe:
  httpGet:
    path: /healthz
    port: 8080
  periodSeconds: 5
```

_This guide only gives you the idea behind graceful shutdown. Planning your graceful shutdown strategy depends on your application's characteristics._

<!-- > [!IMPORTANT] Question!
> _"Isn't it better to still use terminating pod as a fallback if there are no other pods?"_
> >
> There are 2 situations to consider:
> - During normal operation, when a pod is marked for termination, there's typically another pod that's already running and handling traffic.
> - During rolling updates, Kubernetes creates a new pod first and waits until it's ready before sending SIGTERM to the pod being replaced.
> >
> However, if the other pod suddenly breaks while a pod is terminating, the terminating pod will still receive traffic as a fallback mechanism. This raises a question: should we avoid failing the readiness probe during termination to ensure this fallback works?
> >
> The answer is most likely no. If we don't fail the readiness probe, we might face worse consequences if the terminating pod is abruptly killed with `SIGKILL`. This could lead to corrupted processes or data and cause more serious issues. -->

## 4. Handle Pending Requests

Now that we are shutting down the server gracefully, we need to choose a timeout based on your shutdown budget:

```go
ctx, cancelFn := context.WithTimeout(context.Background(), timeout)
err := server.Shutdown(ctx)
```

The `server.Shutdown` function returns in only two situations:

1. All active connections are closed and all handlers have finished processing.
2. The context passed to `Shutdown(ctx)` expires before the handlers finish. In this case, the server gives up waiting.

In either case, `Shutdown` only returns after the server has completely stopped handling requests. This is why your handlers must be fast and context-aware. Otherwise, they may be cut off mid-process in case 2, which can cause issues like partial writes, data loss, inconsistent state, open transactions, or corrupted data.

A common issue is that handlers are not automatically aware when the server is shutting down.

So, how can we notify our handlers that the server is shutting down? The answer is by using context. There are two main ways to do this:

### a. Use context middleware to inject cancellation logic

This middleware wraps each request with a context that listens to a shutdown signal:

```go
func WithGracefulShutdown(next http.Handler, cancelCh <-chan struct{}) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := WithCancellation(r.Context(), cancelCh)
        defer cancel()

        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
    })
}
```

### b. Use `BaseContext` to provide a global context to all connections

Here, we create a server with a custom `BaseContext` that can be canceled during shutdown. This context is shared across all incoming requests:

```go
ongoingCtx, cancelFn := context.WithCancel(context.Background())
server := &http.Server{
    Addr: ":8080",
    Handler: yourHandler,
    BaseContext: func(l net.Listener) context.Context {
        return ongoingCtx
    },
}

// After attempting graceful shutdown:
cancelFn()
time.Sleep(5 * time.Second) // optional delay to allow context propagation
```

In an HTTP server, you can customize two types of contexts: `BaseContext` and `ConnContext`. For graceful shutdown, `BaseContext` is more suitable. It allows you to create a global context with cancellation that applies to the entire server, and you can cancel it to signal all active requests that the server is shutting down.

![Full graceful shutdown with propagation delay](/blog/go-graceful-shutdown/shutdown-context-propagation-timeline.webp)
<figcaption style="text-align: center; font-style: italic;">Full graceful shutdown with propagation delay</figcaption>

All of this work around graceful shutdown won't help if your functions do not respect context cancellation. Try to avoid using `context.Background()`, `time.Sleep()`, or any other function that ignores context.

For example, `time.Sleep(duration)` can be replaced with a context-aware version like this:

```go
func Sleep(ctx context.Context, duration time.Duration) error {
    select {
    case <-time.After(duration):
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

> [!WARNING] Leaking Resources  
> In older versions of Go, `time.After` can leak memory until the timer fires. This was fixed in Go 1.23 and newer. If you're unsure which version you're using, consider using `time.NewTimer` along with `Stop` and an optional `<-t.C` check if `Stop` returns false.  
> > 
> See: [time: stop requiring Timer/Ticker.Stop for prompt GC](https://github.com/golang/go/issues/61542)

Although this article focuses on HTTP servers, the same concept applies to third-party services as well. For example, the `database/sql` package has a `DB.Close` method. It closes the database connection and prevents new queries from starting. It also waits for any ongoing queries to finish before fully shutting down.

The core principle of graceful shutdown is the same across all systems: **Stop accepting new requests or messages, and give existing operations time to finish within a defined grace period.**

Some may wonder about the `server.Close()` method, which shuts down the ongoing connections immediately without waiting for requests to finish. Can it be used after `server.Shutdown()` returns an error?

The short answer is yes, but it depends on your shutdown strategy. The `Close` method forcefully closes all active listeners and connections:

- Handlers that are actively using the network will receive errors when they try to read or write.
- The client will immediately receive a connection error, such as `ECONNRESET` ('socket hang up')
- However, long-running handlers that are not interacting with the network may **continue running** in the background.

This is why using context to propagate a shutdown signal is still the more reliable and graceful approach.

## 5. Release Critical Resources

A common mistake is releasing critical resources as soon as the termination signal is received. At that point, your handlers and in-flight requests may still be using those resources. You should delay the resource cleanup until the shutdown timeout has passed or all requests are done.

In many cases, simply letting the process exit is enough. The operating system will automatically reclaim resources. For instance:

- Memory allocated by Go is automatically freed when the process terminates.
- File descriptors are closed by the OS.
- OS-level resources like process handles are reclaimed.

However, there are important cases where explicit cleanup is still necessary during shutdown:

- **Database connections** should be closed properly. If any transactions are still open, they need to be committed or rolled back. Without a proper shutdown, the database has to rely on connection timeouts.
- **Message queues and brokers** often require a clean shutdown. This may involve flushing messages, committing offsets, or signaling to the broker that the client is exiting. Without this, there can be rebalancing issues or message loss.
- **External services** may not detect the disconnect immediately. Closing connections manually allows those systems to clean up faster than waiting for TCP timeouts.

A good rule is to shut down components in the reverse order of how they were initialized. This respects dependencies between components. 

Go's `defer` statement makes this easier since the last deferred function is executed first:

```go
db := connectDB()
defer db.Close()

cache := connectCache()
defer cache.Close()
```

Some components require special handling. For example, if you cache data in memory, you might need to write that data to disk before exiting. In those cases, design a shutdown routine specific to that component to handle the cleanup properly.

## Summary

This is a complete example of a graceful shutdown mechanism. It is written in a flat and straightforward structure to make it easier to understand. You can customize it to fit your application as needed, using `errgroup`, `WaitGroup`, or any other patterns:

```go
const (
	_shutdownPeriod      = 15 * time.Second
	_shutdownHardPeriod  = 3 * time.Second
	_readinessDrainDelay = 5 * time.Second
)

var isShuttingDown atomic.Bool

func main() {
	// Setup signal context
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Readiness endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if isShuttingDown.Load() {
			http.Error(w, "Shutting down", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintln(w, "OK")
	})

	// Sample business logic
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Fprintln(w, "Hello, world!")
		case <-r.Context().Done():
			http.Error(w, "Request cancelled.", http.StatusRequestTimeout)
		}
	})

	// Ensure in-flight requests aren't cancelled immediately on SIGTERM
	ongoingCtx, stopOngoingGracefully := context.WithCancel(context.Background())
	server := &http.Server{
		Addr: ":8080",
		BaseContext: func(_ net.Listener) context.Context {
			return ongoingCtx
		},
	}

	go func() {
		log.Println("Server starting on :8080.")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Wait for signal
	<-rootCtx.Done()
	stop()
	isShuttingDown.Store(true)
	log.Println("Received shutdown signal, shutting down.")

	// Give time for readiness check to propagate
	time.Sleep(_readinessDrainDelay)
	log.Println("Readiness check propagated, now waiting for ongoing requests to finish.")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), _shutdownPeriod)
	defer cancel()
	err := server.Shutdown(shutdownCtx)
	stopOngoingGracefully()
	if err != nil {
		log.Println("Failed to wait for ongoing requests to finish, waiting for forced cancellation.")
		time.Sleep(_shutdownHardPeriod)
	}

	log.Println("Server shut down gracefully.")
}
```

## Who We Are

If you want to monitor your services, track metrics, and see how everything performs, you might want to check out [VictoriaMetrics](https://docs.victoriametrics.com/). It's a fast, **open-source**, and cost-saving way to keep an eye on your infrastructure.

And we're Gophers, enthusiasts who love researching, experimenting, and sharing knowledge about Go and its ecosystem. If you spot anything that's outdated or if you have questions, don't hesitate to reach out. You can drop me a DM on [X(@func25)](https://twitter.com/func25).

Related articles:

- [Golang Series at VictoriaMetrics](/categories/go-@-victoriametrics)
- [How Go Arrays Work and Get Tricky with For-Range](/blog/go-array/)
- [Slices in Go: Grow Big or Go Home](/blog/go-slice/)
- [Go Maps Explained: How Key-Value Pairs Are Actually Stored](/blog/go-map/)
- [Golang Defer: From Basic To Traps](/blog/defer-in-go/)
- [Vendoring, or go mod vendor: What is it?](/blog/vendoring-go-mod-vendor/)