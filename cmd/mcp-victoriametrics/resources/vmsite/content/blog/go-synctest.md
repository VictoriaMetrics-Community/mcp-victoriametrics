---
draft: false
page: blog blog_post
authors:
 - Phuong Le
date: 2025-05-23
title: "Go synctest: Solving Flaky Tests"
summary: "Traditional concurrent Go tests can be flaky due to non-deterministic scheduler behavior and timing. Go 1.24's experimental synctest feature provides deterministic testing by running goroutines in isolated 'bubbles' where a synthetic clock only advances when all internally managed goroutines are durably blocked."
enableComments: true
toc: true
categories:
  - Go @ VictoriaMetrics
  - Open Source Tech
tags:
  - go
  - golang
  - synctest
  - goroutine
  - test
images:
  - /blog/go-synctest/preview.webp
---

To understand what `synctest` solves, we must first look at the core issue: non-determinism in concurrent tests.

```go
func TestSharedValue(t *testing.T) {
	var shared atomic.Int64
	go func() {
		shared.Store(1)
		time.Sleep(1 * time.Microsecond)
		shared.Store(2)
	}()

	// Check the shared value after 5 microseconds
	time.Sleep(5 * time.Microsecond)
	if shared.Load() != 2 {
		t.Errorf("shared = %d, want 2", shared.Load())
	}
}
```

This test starts a goroutine that modifies a shared variable. It sets `shared` to 1, sleeps for 1 microsecond, and then sets it to 2.

Meanwhile, the main test function waits for 5 microseconds before checking if `shared` has reached 2. At first glance, it seems like this test should always pass. After all, 5 microseconds should be enough time for the goroutine to complete its execution.

However, running the test repeatedly using:

```bash
go test -run TestSharedValue -count=1000
```

will show that the test sometimes fails. You might see outputs like:

```bash
shared = 0, want 2
```

or

```bash
shared = 1, want 2
```

This happens because the test is flaky. Sometimes the goroutine hasn't completed by the time the check runs or even started. The result depends on the system scheduler and how quickly the goroutine is picked up by the runtime.

The accuracy of `time.Sleep` and the behavior of the scheduler can vary widely. Factors such as operating system differences and system load can affect timing. This makes any synchronization strategy based solely on sleep unreliable.

While this example uses microsecond delays for demonstration, real-world issues often involve delays at the millisecond or second level, especially under high load. 

Real systems affected by this type of flakiness include background cleanup, retry logic, time-based cache eviction, heartbeat monitoring, leader election in distributed environments, etc.

Tests like this depend on timing and can also be time-consuming. Imagine if it had to wait 5 seconds instead of just 5 microseconds.

## What is synctest?

`synctest` is a new feature introduced in Go 1.24. It enables deterministic testing of concurrent code by running goroutines in controlled, isolated environments.

Consider this example that does not use `synctest`:

```go
func TestTimingWithoutSynctest(t *testing.T) {
	start := time.Now().UTC()
	time.Sleep(5 * time.Second)
	t.Log(time.Since(start))
}
```

When you run this test with:

```bash
go test . -v
```

You will see that the output is never exactly `5s`. Instead, it might look like `5.329s`, `5.394s`, or `5.456s`. These variations come from delays in system scheduling and timing resolution.

With `synctest`, time is completely controlled. The duration becomes consistent, and the output will always show `5s`.

To use `synctest`, wrap your test logic inside a function and pass it to `synctest.Run()`:

```go
import "testing/synctest"

func TestTimingWithSynctest(t *testing.T) {
	synctest.Run(func() {
		start := time.Now().UTC()
		time.Sleep(5 * time.Second)
		t.Log(time.Since(start))
	})
}
```

Then run the test with the `GOEXPERIMENT=synctest` flag:

```bash
GOEXPERIMENT=synctest go test -run TestTimingWithSynctest -v
```

Sample output:

```bash {hl_line=9}
=== RUN   TestTimingWithSynctest
    main_test.go:8: 5s
--- PASS: TestTimingWithSynctest (0.00s)
PASS
```

Note that `time.Sleep` inside `synctest` returns immediately. The test does not actually wait 5 seconds. This makes tests run much faster while still being accurate.

Now that we know `synctest` manipulates time to produce deterministic behavior, we can use it to fix the earlier flaky test. Simply wrap the test body with `synctest.Run`:

```go
func TestSharedValue(t *testing.T) {
	synctest.Run(func() {
		var shared atomic.Int64
		go func() {
			shared.Store(1)
			time.Sleep(1 * time.Microsecond)
			shared.Store(2)
		}()

		// Check the shared value after 5 microseconds
		time.Sleep(5 * time.Microsecond)
		if shared.Load() != 2 {
			t.Errorf("shared = %d, want 2", shared.Load())
		}
	})
}
```

With this change, the test will pass every time. But how does it fix the problem that Go runtime scheduler does not pick up the goroutine to run?

The reason is that time is controlled. The 5 microseconds is simulated rather than real. When the code runs, time is effectively frozen, and `synctest` manages its progression. In other words, the logic doesn't rely on real time, but instead depends on a deterministic execution order.

### Wait Mechanism

In addition to synthetic time, `synctest` also provides a powerful synchronization primitive: the `synctest.Wait` function.

When you call `synctest.Wait()`, it blocks until all other goroutines (in the same `synctest` group) have either finished or are durably blocked. The most common use of `Wait()` is to start background goroutines, then pause until they reach a stable point before making assertions.

Here is an example where `Wait()` ensures that the `afterFunc` callback has been called:

```go
synctest.Run(func() {
    ctx, cancel := context.WithCancel(context.Background())
    
    afterFuncCalled := false
    context.AfterFunc(ctx, func() {
        afterFuncCalled = true
    })
    
    // Cancel the context and wait for the AfterFunc to complete
    cancel()
    synctest.Wait()

    // Now we can safely check that the callback has been called
    fmt.Printf("after context is canceled: afterFuncCalled=%v\n", afterFuncCalled)
})
```

When we call `cancel()`, the function passed to `context.AfterFunc` runs in a separate goroutine. Without coordination, we cannot be sure when that goroutine will be scheduled or when it will finish.

Because `synctest` tracks all goroutines in the test bubble, it knows their exact state. When `Wait()` returns, it guarantees that all other goroutines are either finished or blocked. This allows you to make reliable and deterministic assertions about the program's state.

## How synctest works

`synctest` works by creating isolated environments called "bubbles." A bubble is a set of goroutines that run in a controlled and independent environment, separated from the normal execution of the program.

When you call `synctest.Run(f)`, the Go runtime creates a new execution bubble. This bubble has several unique characteristics that make it different from regular Go behavior:

### 1. Synthetic Time

Each bubble has its own synthetic clock. This synthetic time starts at midnight UTC on January 1, 2000 (epoch 946684800000000000):

```go
func TestTimingWithSynctest(t *testing.T) {
	synctest.Run(func() {
		t.Log(time.Now().UTC())
	})
}

// Output:
// 2000-01-01 00:00:00 +0000 UTC
```

Inside the bubble, time does not move forward in real time. Instead, Go pauses time and observes what the goroutines are doing. If any goroutine is still active (not blocked), synthetic time stays frozen:

```go
func TestTimingWithSynctest(t *testing.T) {
	synctest.Run(func() {
		t.Log(time.Now().UnixNano())

		var now int64
		for range 10000000 {
			now = time.Now().UnixNano()
		}

		t.Log(now)
	})
}

// Output:
// 946684800000000000
// 946684800000000000
```

Time only advances when all goroutines in the bubble are blocked. This means they are waiting on operations such as `time.Sleep`, channel receives, mutexes, or other blocking calls.

In a `synctest` bubble, time only progresses to trigger scheduled events. This gives the test complete control over execution timing and ordering.

For example, if a goroutine is sleeping for 5 seconds, and all others are also blocked, Go will instantly move the synthetic time forward by 5 seconds. This allows the goroutine to resume immediately, without waiting for real time to pass.

### 2. Goroutine Coordination

When `synctest.Run(f)` is called, the current goroutine becomes the root of the bubble. This root goroutine manages synthetic time and coordinates the execution of all other goroutines inside the bubble.

The function `f`, passed to `synctest.Run`, is launched in a new goroutine and becomes part of the bubble. The root goroutine then enters a loop to manage time and control the scheduling of other bubble goroutines.

![Illustrates synctest goroutine states](/blog/go-synctest/synctest-goroutine-states.webp)
<figcaption style="text-align: center; font-style: italic;">Illustrates synctest goroutine states</figcaption>

There are two categories of blocked goroutines: **external blocked** and **durably blocked**.

**Durably blocked** means the goroutine cannot proceed until something else triggers an unblock, and that "something" is controlled inside the test environment. Examples include:

- `time.Sleep()`
- `sync.Cond.Wait()`
- `sync.WaitGroup.Wait()`
- Operations on `nil` channels
- `select` statements where all cases involve channels within the bubble
- Sends and receives on channels created within the bubble

Goroutines are **not durably blocked** if they are waiting on events outside the bubble. These include:

- System calls like file or network I/O
- External event handling (such as reading from a socket)
- Channel operations on channels that were created outside the bubble

From `synctest`'s perspective, goroutines blocked on external events are considered to be running, because their progress depends on real-world state.

So if you have a goroutine that is forever externally blocked, and another goroutine that is durably blocked on something like `time.Sleep(5 * time.Microsecond)`, the sleep will never complete. Since the external block prevents the system from reaching a fully blocked state, synthetic time will not advance, and the durably blocked goroutine will remain paused.

When there are no running goroutines and all active ones are durably blocked, `synctest` proceeds to either wake a goroutine waiting on `synctest.Wait()` or continue executing the root goroutine. The decision logic looks like this:

```go
func (sg *synctestGroup) maybeWakeLocked() *g {
	if sg.running > 0 || sg.active > 0 {
		return nil
	}

	sg.active++
	if gp := sg.waiter; gp != nil {
		return gp
	}

	return sg.root
}
```

The role of the root goroutine at this point is to find the next scheduled timer event. This could be triggered by functions like `time.Sleep`, `time.Timer`, `time.Ticker`, or `time.AfterFunc`. All of these create timers internally.

Once the root finds the next event, it sets the synthetic time to that moment (`sg.now = next`), then parks itself and waits for the test scheduler to resume the goroutine that should now run.

Remember that `synctest` is primarily designed to test the timing and correctness of synchronization logic, not to simulate real-world timing behavior exactly. If used incorrectly, it may hide bugs that would appear in real-world conditions.

And as a final note, this article was written while `synctest` is still experimental. Some details may change over time, but the core concepts are expected to stay the same.

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
