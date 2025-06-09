---
draft: false
page: blog blog_post
authors:
 - Phuong Le
date: 2024-11-01
title: "Go sync.Once is Simple... Does It Really?"
summary: "The sync.Once is probably the easiest sync primitive to use, but there's more under the hood than you might think. It's also a good opportunity to understand how it works by juggling both atomic operations and mutexes."
enableComments: true
categories:
 - Go @ VictoriaMetrics
 - Open Source Tech
tags:
 - go
 - golang
 - sync
 - once
images:
 - /blog/go-sync-once/go-sync-once-preview.webp
---

This article is part of our ongoing series about handling concurrency in Go, a quick rundown of what we've covered so far:

- [Go sync.Mutex: Normal and Starvation Mode](/blog/go-sync-mutex)
- [Go sync.WaitGroup and The Alignment Problem](/blog/go-sync-waitgroup)
- [Go sync.Pool and the Mechanics Behind It](/blog/go-sync-pool)
- [Go sync.Cond, the Most Overlooked Sync Mechanism](/blog/go-sync-cond)
- [Go sync.Map: The Right Tool for the Right Job](/blog/go-sync-map)
- Go sync.Once is Simple... Does It Really? (We're here) 
- [Go Singleflight Melts in Your Code, Not in Your DB](/blog/go-singleflight)

![Go sync.Once is Simple... Does It Really?](/blog/go-sync-once/go-sync-once-preview.webp) 
<figcaption style="text-align: center; font-style: italic;">Go sync.Once is Simple... Does It Really?</figcaption>

The `sync.Once` is probably the easiest sync primitive to use, but there's more under the hood than you might think.

It's a good opportunity to understand how it works by juggling both atomic operations and mutexes.

In this discussion, we're going to break down what `sync.Once` is, how you can use it properly, and - maybe most importantly, how it actually works under the hood. We'll also take a look at its cousins: `OnceFunc`, `OnceValue[T]`, and `OnceValues[T, K]`.

## What is sync.Once?

The `sync.Once` is exactly what you reach for when you need a function to run just one time, no matter how many times it gets called or how many goroutines hit it simultaneously.

It's perfect for initializing a singleton resource, something that should only happen once in your application's lifecycle, such as setting up a database connection pool, initializing a logger, or maybe configuring your metrics system, etc.

```go
var once sync.Once
var conf Config

func GetConfig() Config {
    once.Do(func() {
        conf = fetchConfig()
    })
    return conf
}
```

If `GetConfig()` is called multiple times, `fetchConfig()` is executed only once.

The real benefit of `sync.Once` is that it delays certain operations until they are first needed (lazy-loading), which can improve runtime performance and reduce initial memory usage. For instance, if a large lookup table is created only when accessed for the first time, then memory and processing time for creating the table are saved until that moment.

Most of the time, it is better for initializing (external) resources than using `init()`.

Now, something important to keep in mind: once you've used a `sync.Once` object to run a function, that's it — you can't reuse it. Once it's done, it's done. 

After `sync.Once` successfully runs the function with `Do(f)`, it marks itself as "complete" internally with a flag (`done`), and from that point on, any further calls to `Do(f)` won't run the function again, even if it's a different function.

```go
var once sync.Once

func main() {
    once.Do(func() {
        fmt.Println("This will be printed once")
    })

    once.Do(func() {
        fmt.Println("This will not be printed")
    })
}

// Output:
// This will be printed once
```

There's no built-in way to reset a sync.Once either. Once it's done its job, it's retired for good.

Now, here's an interesting twist. If the function you pass to `Once.Do` panics while running, `sync.Once` still treats that as "mission accomplished." That means future calls to `Do(f)` won't run the function again. This can be tricky, especially if you're trying to catch the panic and handle the error afterward, there is no retry.

Also, if you need to handle errors that might come out of `f`, it can get a little awkward to write:

```go
var once sync.Once
var config Config

func GetConfig() (Config, error) {
    var err error
    once.Do(func() {
        config, err = fetchConfig() 
    })
    return config, err
}
```

This works well, but only the first time.

The thing is, if `fetchConfig()` fails, only the first call will receive the error. Any subsequent calls—second, third, and so on—won’t return that error (as [venicedreamway](https://www.reddit.com/user/venicedreamway/) pointed out). To make the behavior consistent across calls, we need to make `err` a package-scoped variable, like this:

```go
var once sync.Once
var config Config
var err error

func GetConfig() (Config, error) {
    var err error
    once.Do(func() {
        config, err = fetchConfig() 
    })
    return config, err
}
```

Good news, though! From Go 1.21 onward, we get `OnceFunc`, `OnceValue`, and `OnceValues`. 

These are basically handy wrappers around `sync.Once` that make things smoother, without sacrificing any performance. `OnceFunc` is pretty straightforward, it takes your function `f` and wraps it in another function that you can call as many times as you want, but `f` itself will only run once:

```go
var wrapper = sync.OnceFunc(printOnce)

func printOnce() {
  fmt.Println("This will be printed once")
}

func main() {
  wrapper()
  wrapper()
}

// Output:
// This will be printed once
```

Even if you call `wrapper()` a bunch of times, `printOnce()` runs only on the first call. 

Now, if `f()` panics during that first execution, every future call to `wrapper()` will also panic with the same error. It's like it locks in the failure state, so your app doesn't continue like nothing went wrong when something critical didn't initialize correctly. 

But this doesn't really solve the problem of catching errors in a nice way.

Let's move on to something even more useful: `OnceValue` and `OnceValues`. These are cool because they remember the result of `f` after the first execution and just return the cached result on future calls.

```go
var getConfigOnce = sync.OnceValue(func() Config {
	fmt.Println("Loading config...")
	return fetchConfig() // Pretend this is expensive
})

func main() {
	config1 := getConfigOnce() // Loading config...
	config2 := getConfigOnce() // No print, just returns the cached config
  ...
}

// Output:
// Loading config...
```

After that first call to `getConfigOnce()`, it just hands you the same result without re-running `fetchConfig()`. Beside the lazy loading, you don't need to deal with closures here.

But what about errors? Fetching something usually involves error handling.

This is where `sync.OnceValues` comes in. It works like `OnceValue`, but lets you return multiple values, including errors. So, you can cache both the result and any error that comes up during the first run.

```go
var config Config

var getConfigOnce = sync.OnceValues(fetchConfig)

func main() {
  var err error

  config, err = getConfigOnce()
  if err != nil {
    log.Fatalf("Failed to fetch config: %v", err)
  }
  ...
}
```

Now, `getConfigOnce()` behaves just like a normal function — it's still concurrent-safe, caches the result, and only incurs the cost of running the function once. After that, every call is cheap.

> _"So, if an error happens, is that cached too?"_

Unfortunately, yes. Whether it's a panic or an error, both the result and the failure state are cached. So, the calling code needs to be aware that it might be dealing with a cached error or failure. If you need to retry, you'd have to create a new instance from `sync.OnceValues` to re-run the initialization.

Also, in the example, we return an error as the second value to match the typical function signature we're used to, but honestly, it could be anything depending on what you need.

## How it works?

If you're not familiar with atomic operations or synchronization techniques in Go, then `sync.Once` is a great starting point, as it's one of the simplest synchronization primitives.

It's really simple compared to something like `sync.Mutex`, which is [one of the trickiest synchronization](/blog/go-sync-mutex) tools in Go. So let's take a step back and think about how `sync.Once` is actually implemented.

Here's the basic structure of `sync.Once`:

```go
type Once struct {
	done atomic.Uint32
	m    Mutex
}
```

The first thing you'll notice is the `done` field, which uses an atomic operation. 

And here's an interesting detail, `done` is placed at the very top of the struct for a reason. On many CPU architectures (like x86-64), accessing the first field in a struct is faster because it sits at the base address of the memory block. This little optimization allows the CPU to load the first field more directly, without calculating memory offsets.

Also, putting it at the top helps with inlining optimization, which we'll get to in a bit.

So, what’s the first thing that comes to mind when implementing `sync.Once` to make sure a function runs only once? Using a mutex, right?

```go
func (o *Once) Do(f func()) {
	o.m.Lock()
	defer o.m.Unlock()

	if o.done.Load() == 0 {
		o.done.Store(1)
		f()
	}
}
```

It's simple and gets the job done. The idea is that the mutex (`o.m.Lock()`) allows only one goroutine to enter the critical section at a time. Then, if `done` is still 0 (meaning the function hasn't run yet), it sets `done` to 1 and runs the function `f()`.

This is actually the original version of `sync.Once`, written by Rob Pike back in 2010. 

Now, the version we just looked at works fine, but it's not the most performant one. Even after the first call, every time `Do(f)` is called, it still grabs a lock, which means goroutines are waiting on each other. We can definitely do better by adding a quick exit if the task is already done.

```go
func (o *Once) Do(f func()) {
  if atomic.LoadUint32(&o.done) == 1 {
    return
  }

  // slow path
  o.m.Lock()
  defer o.m.Unlock()

  if o.done.Load() == 0 {
    o.done.Store(1)
    f()
  }
}
```

This gives us a **fast path**, when the `done` flag is set, we skip the lock entirely and just return immediately. Nice and quick. But, if the flag isn't set, we fall back to the **slower path**, which locks the mutex, rechecks `done`, and then runs the function.

Now, we have to re-check `done` after acquiring the lock because there's a small window between checking the flag and actually locking the mutex, where another goroutine might have already run `f()` and set the flag. We also set the `done` flag before calling `f()`. The idea is that even if `f()` panics, we still mark it as "success" to prevent it from running again. 

But, this action is also a mistake.

Imagine this scenario, we set `done` to 1, but `f()` hasn't finished yet, maybe it's stuck on a long network call. 

![sync.Once with race condition](/blog/go-sync-once/go-sync-once-done-mistake.webp) 
<figcaption style="text-align: center; font-style: italic;">sync.Once with race condition</figcaption>

Now, a second goroutine comes along, checks the flag, sees that it's set, and mistakenly thinks, "Great, the resource is ready to go!" But in reality, it's still being fetched. So what happens? Nil dereference and panic! The resource isn't ready, and the system tries to use it too early.

We can fix the problem by using `defer` like this:

```go
func (o *Once) Do(f func()) {
  if o.done.Load() == 1 {
    return
  }

  // slow path
  o.m.Lock()
  defer o.m.Unlock()

  if o.done.Load() == 0 {
    defer o.done.Store(1)
    f()
  }
}
```

You might think, "Okay, this looks pretty solid now." But it's still not perfect.

The idea is that Go supports something called inlining optimization.

If a function is simple enough, the Go compiler will "inline" it, meaning it'll take the function's code and paste it directly where the function is called, making it faster. Our `Do()` function is still too complex for inlining, though, because it has multiple branches, a defer, and function calls.

To help the compiler make better decisions about inlining the code, we can move the slow-path logic to another function:

```go
func (o *Once) Do(f func()) {
	if o.done.Load() == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
  
	if o.done.Load() == 0 {
		defer o.done.Store(1)
		f()
	}
}
```

This makes the `once.Do()` function much simpler and can be inlined by the compiler.

Even though from our point of view, it now has 2 function calls, it's not quite that way in practice. The `o.done.Load()` is an atomic operation that Go's compiler handles in a special way (compiler intrinsic), so it doesn't count toward the function call complexity.

> _"Why not just inline `doSlow()`?"_

The reason is that after the first call to `Do(f)`, the common scenario is the fast path — just checking if the function has already run. 

In real-world applications, after `f()` runs once (which is the slow path), there are usually many more calls to `once.Do(f)` that just need to quickly check the `done` flag without locking or re-running the function.

That's why we optimize for the fast path, where we just check if it's already done and immediately return. And remember when we talked about why the `done` field is placed first in the `Once` struct? That's because it makes the fast path quicker by being easier to access.

Now, we have the perfect version of `sync.Once`, but here's the final quiz. The Go team also mentioned an implementation version using a compare-and-swap (CAS) operation, which makes our `Do()` function much simpler:

```go
func (o *Once) Do(f func()) {
	if o.done.CompareAndSwap(0, 1) {
		f()
	}
}
```

The idea is that whichever goroutine can successfully swap the value of done from 0 to 1 would "win" the race and run `f()`, while all the other goroutines would just return.

But why doesn’t the Go team use this version? Can you guess why before reading the next section, as we already discussed this mistake?

Yes, this brings us back to the same mistake we talked about earlier:

![sync.Once with CAS](/blog/go-sync-once/go-sync-once-cas.webp) 
<figcaption style="text-align: center; font-style: italic;">sync.Once with Compare-And-Swap (CAS) operation</figcaption>

While the "winning" goroutine is still running `f()`, other goroutines might come and check the `done` flag, think `f()` is already finished, and proceed to use resources that aren't fully ready yet.

And... that's it! `sync.Once` is simple in both implementation and usage, but it turns out to be quite tricky to get it right.

## Stay Connected

Hi, I'm Phuong Le, a software engineer at VictoriaMetrics. The writing style above focuses on clarity and simplicity, explaining concepts in a way that's easy to understand, even if it's not always perfectly aligned with academic precision.

If you spot anything that's outdated or if you have questions, don't hesitate to reach out. You can drop me a DM on [X(@func25)](https://twitter.com/func25).

Related articles:

- [Golang Series at VictoriaMetrics](/categories/go-@-victoriametrics)
- [How Go Arrays Work and Get Tricky with For-Range](/blog/go-array)
- [Slices in Go: Grow Big or Go Home](/blog/go-slice)
- [Go Maps Explained: How Key-Value Pairs Are Actually Stored](/blog/go-map)
- [Golang Defer: From Basic To Traps](/blog/defer-in-go)
- [Inside Go's Unique Package: String Interning Simplified](/blog/go-unique-package-intern-string)
- [Vendoring, or go mod vendor: What is it?](/blog/vendoring-go-mod-vendor)

## Who We Are

If you want to monitor your services, track metrics, and see how everything performs, you might want to check out [VictoriaMetrics](https://docs.victoriametrics.com/). It's a fast, **open-source**, and cost-saving way to keep an eye on your infrastructure.

And we're Gophers, enthusiasts who love researching, experimenting, and sharing knowledge about Go and its ecosystem.
