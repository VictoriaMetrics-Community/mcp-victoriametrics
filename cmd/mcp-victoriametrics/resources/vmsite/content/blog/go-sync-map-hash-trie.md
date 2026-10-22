---
draft: false
page: blog blog_post
authors:
  - Phuong Le
date: 2026-08-26
title: "Understanding Go's sync.Map from API to Hash Trie"
summary: "Go's sync.Map now uses a hash trie. This article builds the implementation from a normal map, then explains lock-free reads, fine-grained writer locking, collisions, retries, deletion, Clear, Range, and the trade-offs against map plus RWMutex."
enableComments: true
toc: true
categories:
  - Go @ VictoriaMetrics
  - Open Source Tech
tags:
  - go
  - golang
  - sync
  - map
  - concurrency
images:
  - /blog/go-sync-map-hash-trie/preview.webp
---

This post is part of a series about handling concurrency in Go:

- [Go sync.Mutex: Normal and Starvation Mode](/blog/go-sync-mutex/)
- [Go sync.WaitGroup and The Alignment Problem](/blog/go-sync-waitgroup/)
- [Go sync.Pool and the Mechanics Behind It](/blog/go-sync-pool/)
- [Go sync.Cond, the Most Overlooked Sync Mechanism](/blog/go-sync-cond/)
- [Go sync.Map: The Right Tool for the Right Job](/blog/go-sync-map/)
- Understanding Go's sync.Map from API to Hash Trie (We're here)
- [Go Sync.Once is Simple... Does It Really?](/blog/go-sync-once/)
- [Go Singleflight Melts in Your Code, Not in Your DB](/blog/go-singleflight/)

`sync.Map` has changed since my [previous sync.Map article](https://victoriametrics.com/blog/go-sync-map/). Its public API stayed the same, but Go 1.24 changed its default implementation to an experimental hash trie.

Go 1.26 removed the experiment flag and made the hash trie the official implementation of `sync.Map` in the standard library.

We will discuss the current implementation from the beginning, so you do not need to read the old article first. We will build everything up from a normal map to the design behind `sync.Map`. If you already know the setup, feel free to skip those sections.

## 1. Start with a normal Go map

### How a map finds a key

Before we look at the `sync.Map`, we need one piece shared by every map implementation: how a map finds 1 key among everything it stores.

A normal Go map gives us the simplest place to start:

```go {hl_lines="6"}
m := map[string]int{}

m["cat"] = 31
m["dog"] = 28

value, ok := m["cat"]
fmt.Println(value, ok) // 31 true
```

The lookup is simple: find the `"cat"` key in the map storage and pull its value. Internally, it doesn't check and scan every key linearly. That approach would make each lookup do more work as the map grew, since the map may contain thousands of other keys.

Instead, the lookup starts by applying a hash function to the key:

![Hashing a key into a fixed-width integer.](/blog/go-sync-map-hash-trie/key-to-hash.webp)

<figcaption style="text-align: center; font-style: italic;">Hashing a key into a fixed-width integer.</figcaption>

The hash result has 32 bits on a 32-bit architecture and 64 bits on a 64-bit architecture. We use 64-bit examples in this article.

For example, the first 4 hash bits can narrow the search to the `"dog"` and `"cat"` keys:

![Hash lookup and equality check.](/blog/go-sync-map-hash-trie/hash-then-equality.webp)

<figcaption style="text-align: center; font-style: italic;">Hash lookup and equality check.</figcaption>

Inside that smaller part of the storage, an equality check (`==`) compares the actual key and confirms the exact match.

Both a built-in `map[K]T` and `sync.Map` share this strategy: use the hash to narrow the search, then compare the actual key. They organize their storage in separate ways, so their lookup code is not the same.

### A normal map is not safe for concurrent read and write

A normal map does not synchronize access between goroutines. This snippet lets one goroutine write while another reads:

```go {hl_lines=["6","12"]}
func main() {
	m := map[string]int{}

	go func() {
		for {
			m["requests"]++
		}
	}()

	go func() {
		for {
			fmt.Println(m["requests"])
		}
	}()

	select {}
}
```

This snippet contains a data race and can terminate with:

```text
fatal error: concurrent map read and map write
```

The reason is that a write can change several parts of the map that a lookup also reads. It may update a key or value, change lookup metadata and counters, grow the map by copying entries into new storage, etc.

A reader uses the same storage and metadata to find its key. Without synchronization, the reader may access those parts while the writer is changing them because the write is not atomic.

> [!NOTE]
>
> This is also why you cannot take the address of a map entry:
>
> ```go
> p := &m["cat"] // invalid operation: cannot take address of m["cat"]
> ```
>
> There is no stable address to hand out.

The standard typed solution is a map protected by a mutex, either `sync.Mutex` or `sync.RWMutex`:

```go
type Counters struct {
	mu sync.RWMutex
	m  map[string]int
}

func (c *Counters) Load(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[key]
	return v, ok
}

func (c *Counters) Store(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}
```

This is often the best design. You keep the concrete key and value types, so no type assertion is needed here. And because one lock protects the whole struct, you can change several entries together without another goroutine seeing a half-finished state.

The same wrapper can be written once as a generic type:

```go
type Map[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}
```

The limitation is that all operations coordinate through the same `sync.RWMutex`. A read lock can run concurrently with other read locks, but a writer waits for existing readers and excludes every other operation.

![External locking and internal locking.](/blog/go-sync-map-hash-trie/plain-map-vs-sync-map.webp)

<figcaption style="text-align: center; font-style: italic;">External locking and internal locking.</figcaption>

This design uses one lock for the entire map, so concurrent access can create contention around that lock. We need a way to shard the lock, and that is where `sync.Map` comes in.

## 2. What sync.Map provides

`sync.Map` provides concurrent versions of common map operations. Multiple goroutines can call `Load`, `Store`, and `Delete` without protecting the map with an external lock:

```go
var m sync.Map

m.Store("cat", 31)

value, ok := m.Load("cat")
if ok {
	number := value.(int)
	fmt.Println(number)
}

m.Delete("cat")
```

`sync.Map` has a valid zero value, so we can declare and use it immediately without initialization. Once we have used it, however, we must not copy it. See also [How Go detects struct copies with sync.noCopy](https://func25.dev/posts/go-sync-nocopy/).

This convenience comes with a type-information tradeoff. `sync.Map` accepts keys and values as `any`, so `Load` also returns an `any` value.

We usually need a type assertion before using that value as its concrete type:

```go {hl_lines=["4"]}
m.Store("cat", 31)

value, _ := m.Load("cat")
number := value.(int) // value is any, so the type has to be restored
```

The key parameter is typed as `any`, but the value passed as the key must still be comparable. A slice therefore compiles at the call site and panics when `sync.Map` tries to hash it.

```go {hl_lines=[7]}
package main

import "sync"

func main() {
	var m sync.Map
	m.Store([]int{1, 2, 3}, "value")
}
```

```text
panic: runtime error: hash of unhashable type []int
```

There is another limitation: `sync.Map` does not let us update several selected keys together as one atomic operation.

If 2 keys must change together, another goroutine can read the map between those changes and see only half of the update. A plain `map[K]T` with your own mutex avoids this problem because the same lock can protect both changes.

The standard library documentation says most code should use a plain map instead. Do not reach for `sync.Map` every time you need a map with a mutex. It recommends `sync.Map` for two situations:

- an entry is written once and read many times, such as a cache that only grows.
- different goroutines mostly use different keys.

## 3. Before you reach for sync.Map

`sync.Map` is safe to call from many goroutines. This is the easy part. A few things are different from a plain map, and those need a little more care.

Here are the main methods of `sync.Map`:

```go
Load(key) // returns the value and whether the key was found
Store(key, value) // stores the value under the key
Delete(key) // removes the key and its value
LoadOrStore(key, value) // returns the current value or stores the new value
LoadAndDelete(key) // removes the key and returns its previous value
Swap(key, value) // replaces the value and returns its previous value
CompareAndSwap(key, old, new) // replaces old with new only when the current value equals old
CompareAndDelete(key, old) // deletes the key only when the current value equals old
Range(func(key, value any) bool) // calls the function for entries until it returns false
Clear() // removes every entry
```

You may notice that `sync.Map` has no `Len` method. The only way to count entries is to walk the whole map with `Range`:

```go
n := 0
m.Range(func(key, value any) bool {
	n++
	return true
})
```

But `Range` does not give us a snapshot. It walks the live structure, as we will see later. If another goroutine stores or deletes a key while the walk is running, you may or may not see that key, and the value you get for a key may be from any moment during the walk.

There are a few interesting details about key order when we use `for range` with a plain `map[K]T` or the `Range` method with `sync.Map`.

Each `sync.Map` chooses a random hash seed when it is first used. That seed affects the hash of every key, so 2 maps containing the same keys can place them differently in their storage. A normal `map[K]T` uses the same general idea.

If we check their order through the `Range` method, it may look like this:

```text
sync.Map #1: [6 3 5 1 2 4 0 7]
sync.Map #2: [2 6 3 1 5 4 0 7]
sync.Map #3: [6 5 1 0 2 4 7 3]
```

That is how the order can change between different `sync.Map` values. What if we iterate over the same `sync.Map` multiple times? The `Range` order is still not specified, just like the order of a plain map, but the current implementation can make it look stable.

Here is a comparison between `sync.Map.Range` and a plain map range:

```text
sync.Map    [7 2 6 0 3 5 1 4]   [7 2 6 0 3 5 1 4]   [7 2 6 0 3 5 1 4]
plain map   [5 6 7 0 1 2 3 4]   [3 4 5 6 7 0 1 2]   [4 5 6 7 0 1 2 3]
```

A built-in `map` chooses random starting offsets for every `range` loop, so its iteration order is more likely to change. In contrast, `sync.Map.Range` always walks the child slots from 0 through 15. If the stored structure does not change, repeated calls can return the same order.

> [!NOTE]
> This behavior is not part of the `sync.Map` contract. It is an internal detail. Insertions, deletions, and concurrent updates can change which entries `Range` reads and when it reads them. We should never write code that depends on the order of either map type.

Now, each call is safe on its own, but `sync.Map` does not turn separate calls into one transaction. If `Clear` and `Store` run at the same time, the `Store` can be lost. `Clear` throws the whole map away at once by swapping in a new empty map.

![Concurrent Store and Clear.](/blog/go-sync-map-hash-trie/clear-loses-store.webp)

<figcaption style="text-align: center; font-style: italic;">Concurrent Store and Clear.</figcaption>

If `Store` starts first, it may read the current storage. `Clear` can run immediately afterward and replace that storage with a new empty one. The call returns normally but `Store` finishes its write using the discarded storage it already read.

`sync.Map` usually needs more memory than a plain map. To measure the difference, we record heap usage before inserting any entries and again after inserting the same number of `int` keys and values into each map.

We run GC before each measurement to clear unrelated garbage. After subtracting the starting heap size, we get roughly how much memory each filled map adds:

![Heap memory added by each map.](/blog/go-sync-map-hash-trie/memory-usage.webp)

<figcaption style="text-align: center; font-style: italic;">Heap memory added by each map.</figcaption>

A plain map stores its `int` keys and values together in arrays of slots. At 1 million entries, those arrays use 36.1 MiB. `sync.Map` has a separate object for each key-value pair, stores both through `any`, and builds trie nodes that contain child pointers and mutexes. Those allocations bring its total to 115.9 MiB.

Of course, the exact cost of both map types also depends on the key and value types, the entry count, and the Go version. Across these measurements, `sync.Map` uses about 3 to 5 times as much memory as the plain map.

To be fair to `sync.Map`, 2 things can make it worth considering:

1. It gives us one-call atomic operations such as `LoadOrStore` and `CompareAndSwap`. A plain map needs our own lock around the combined read and write.
2. `sync.Map` can help when a shared mutex is genuinely the bottleneck, but we should measure that first. Concurrency safety alone is not a reason to choose it, because a plain map under a mutex is also safe and usually leads to simpler code.

## 4. The public type is now very small

We now have the whole picture from the outside. The rest of this article opens `sync.Map` up and follows what happens inside.

The public type is:

```go {hl_lines="5"}
package sync

type Map struct {
	_ noCopy
	m isync.HashTrieMap[any, any]
}
```

The first field `_ noCopy` is only there so that `go vet` warns you if you copy a `sync.Map` by accident:

```go
var a sync.Map
b := a // assignment copies lock value to b: sync.Map contains sync.noCopy
```

_See [How Go detects struct copies with sync.noCopy](https://func25.dev/posts/go-sync-nocopy/)._

The second field `m isync.HashTrieMap[any, any]` is the one we care about. It holds every key and value in the map, and its type is where the entire hash trie lives. `isync` is an alias for `internal/sync`.

> [!NOTE]
> `HashTrieMap` was originally created for the `unique` package. It later moved to `internal/sync` when Go planned to reuse it for `sync.Map`.

Now, look at the concrete type behind that field. It is a generic struct:

```go
type HashTrieMap[K comparable, V any] struct { ... }
```

So the internal type can hold concrete key and value types. Then why does `sync.Map` instantiate it as `isync.HashTrieMap[any, any]`?

If `sync.Map` could use concrete key and value types instead of `any`,

- `Load` could return the value with its actual type, so callers would not need a type assertion.
- The trie could also store keys and values directly in those types, without the interface boxing that contributed to the memory cost above.

The API could look like this:

```go
var prices sync.Map[string, int]
prices.Store("pizza", 9)
price, ok := prices.Load("pizza") // price is an int
```

The problem is that `sync.Map` was added before Go had generics. Go cannot replace it with `sync.Map[K, V]` because existing code such as `var m sync.Map` would stop compiling. That would violate the Go 1 compatibility promise.

> [!NOTE]
> There are two interesting ideas for solving this problem.
>
> - The proposed [`sync/v2`](https://github.com/golang/go/issues/71076) package would provide a new generic API.
> - A separate [discussion about default type arguments](https://github.com/golang/go/discussions/48287) considered making `sync.Map` mean `sync.Map[any, any]` when no types are specified.

Back to `sync.Map`. Each public method delegates directly to the internal map:

```go
func (m *Map) Load(key any) (value any, ok bool) {
	return m.m.Load(key)
}

func (m *Map) Store(key, value any) {
	m.m.Store(key, value)
}

func (m *Map) Delete(key any) {
	m.m.Delete(key)
}
```

Therefore, understanding current `sync.Map` means understanding `internal/sync.HashTrieMap`.

Before we can read that type, we need to know what a trie is.

## 5. From a trie to a hash trie

A trie breaks a key into smaller pieces and uses one tree level for each piece. A lookup reads the pieces in order and chooses a child at every level. Here, each piece is one character:

![A string trie with shared prefixes.](/blog/go-sync-map-hash-trie/string-trie.webp)

<figcaption style="text-align: center; font-style: italic;">A string trie with shared prefixes.</figcaption>

To find `"cat"`, lookup reads `c`, then `a`, then `t`. To find `"can"`, it reads `c`, then `a`, and fails because there is no `n` child.

But strings are not the main point here. What matters is that each level takes the next piece of whatever we are looking up. A piece can be

- a character,
- a group of characters,
- a byte,
- a single bit,
- **a group of bits**,
- ...

That last one is what a hash trie uses. Instead of walking the key, it walks the hash of the key.

### What is a hash trie?

`sync.Map` stores its keys in a hash trie. In Go, this trie has one root and adds more nodes below it as the tree grows. The internal `HashTrieMap` stores the root pointer directly:

```go {hl_lines=["4"]}
type HashTrieMap[K comparable, V any] struct {
	inited   atomic.Uint32
	initMu   Mutex
	root     atomic.Pointer[indirect[K, V]]
	...
	seed     uintptr
}
```

If you notice, `root` is an atomic pointer to an `indirect[K, V]` node. We will explain this indirect node type in a moment. Every lookup loads this pointer first, then uses the hash to select a child at each level:

![The hash trie root.](/blog/go-sync-map-hash-trie/hash-trie-tree.webp)

<figcaption style="text-align: center; font-style: italic;">The hash trie root.</figcaption>

The lookup calculates a 64-bit hash from the key, then divides the hash into 4-bit groups. The first group selects a child of the root, and each later group selects a child at the next level.

Let's say we want to look up the key `"cat"`, and its hash looks like this:

![A 64-bit hash in 4-bit groups.](/blog/go-sync-map-hash-trie/hash-groups.webp)

<figcaption style="text-align: center; font-style: italic;">A 64-bit hash in 4-bit groups.</figcaption>

Four bits give 16 possible values, from `0000` to `1111`, so each group is a number from 0 to 15. So the tree can never be deeper than 16 levels.

> [!NOTE] Why does each node have 16 children instead of fewer or more?
> The source comment calls 16 children the sweet spot for `Load` performance. Using fewer children reduced that performance by 50% or more, but increasing the node to 32 children improved it by only about 1%. Moving from 16 to 32 would therefore double each child array for very little additional performance.

The first group for `"cat"` is `1011`, which equals 11 in decimal. The lookup therefore reads child 11 under the root node, finds the `"cat"` entry, and does not need another level:

![The first hash group selects child 11.](/blog/go-sync-map-hash-trie/trie-walk.webp)

<figcaption style="text-align: center; font-style: italic;">The first hash group selects child 11.</figcaption>

The node at child 11 is called an _entry node_, and it stores both the key `"cat"` and its value.

Let's say we want to store another key-value pair with `"dog"` as the key. Its hash starts with `1011`, the same as `"cat"`, so it selects the same child slot:

![A colliding child expands into an indirect node.](/blog/go-sync-map-hash-trie/trie-expansion.webp)

<figcaption style="text-align: center; font-style: italic;">A colliding child expands into an indirect node.</figcaption>

The map cannot leave both keys in one slot, so it puts a new node there and stores both entries one level deeper. This new node is called an _indirect node_: it holds child pointers instead of a key and value.

Their second hash groups no longer select the same child. `"cat"` uses `1001` to select child 9, while `"dog"` uses `0110` to select child 6.

If the second group selected the same child again, the map would need another indirect node to check the third group. So it keeps adding indirect nodes until the keys select separate children.

Now let's see how Go represents these nodes in the source code. Both node types begin with the same header:

```go {hl_lines=["2","7"]}
type indirect[K comparable, V any] struct {
	node[K, V]
	...
}

type entry[K comparable, V any] struct {
	node[K, V]
	...
}

type node[K comparable, V any] struct {
	isEntry bool
}
```

The `isEntry` field tells the map whether a node is an indirect node or an entry node. Both concrete structs place this shared header first. An indirect node then provides 16 child slots, one for each value represented by a four-bit hash group:

```go {hl_lines=["6"]}
type indirect[K comparable, V any] struct {
	node[K, V]
	dead     atomic.Bool
	mu       Mutex
	parent   *indirect[K, V]
	children [16]atomic.Pointer[node[K, V]]
}
```

An entry node contains one key and its value, which the lookup uses to verify the key and return the stored value:

```go {hl_lines=["4","5"]}
type entry[K comparable, V any] struct {
	node[K, V]
	overflow atomic.Pointer[entry[K, V]]
	key      K
	value    V
}
```

But what happens in the rare case where two different keys have the same complete hash? All 16 four-bit groups are the same, so every level selects the same child for both keys. After the 16th group, the map has no hash bits left to create another indirect node.

That's the responsibility of `overflow` in an entry node, a field we haven't explained yet:

```go {hl_lines=["3"]}
type entry[K comparable, V any] struct {
	node[K, V]
	overflow atomic.Pointer[entry[K, V]]
	key      K
	value    V
}
```

The child slot points to one entry, and its `overflow` field points to the next entry with the same hash:

![An overflow list for a full hash collision.](/blog/go-sync-map-hash-trie/hash-collision-overflow.webp)

<figcaption style="text-align: center; font-style: italic;">An overflow list for a full hash collision.</figcaption>

Congratulations on making it this far. We are done with the main idea of the hash trie and can now look at how each `sync.Map` operation reads and updates it under concurrency.

## 6. How each operation works

Most `sync.Map` methods begin by finding the child selected for a key, then do the work specific to that operation. `Clear` and `Range` are exceptions, and they use separate paths because neither starts with a key.

### Load

We can reuse the previous `sync.Map` example, where `"cat"` and `"dog"` are stored under the same indirect node. The diagram below shows the complete flow for `Load("cat")`. It should look familiar now that we have "mastered" the hash trie:

![Load("cat") through two trie levels.](/blog/go-sync-map-hash-trie/load-cat-example.webp)

<figcaption style="text-align: center; font-style: italic;">Load("cat") through two trie levels.</figcaption>

If the selected child is `nil`, the key is not in the map, so `Load` returns the zero value and `false`. If the child contains a different key, `Load` checks the overflow entries one by one.

Here is the complete shape of `Load`, with some syntax shortened:

```go
func (ht *HashTrieMap[K, V]) Load(key K) (value V, ok bool) {
	// lazy init
	ht.init()
	// Get the hash of the key using this sync.Map's seed.
	hash := ht.keyHash(unsafe.Pointer(&key), ht.seed)

	i := ht.root.Load()
	hashShift := 8 * goarch.PtrSize
	for hashShift != 0 {
		// The next four bits are read from the hash and used as a child index.
		hashShift -= 4
		index := (hash >> hashShift) & 15
		n := i.children[index].Load()

		// A nil child means that the key is missing from the map.
		if n == nil {
			return *new(V), false
		}
		if n.isEntry {
			// Compare the actual key and check overflow entries.
			return n.entry().lookup(key)
		}

		// An indirect node continues at the next trie level.
		i = n.indirect()
	}
	panic("ran out of hash bits")
}
```

_If you do not want to follow every line of the function, its comments are enough to understand the flow. We will explain each part below._

`Load` begins with `init()`, as most hash-trie operations do. This is lazy initialization. It creates the trie state on first use, which is why we can use the zero value of `sync.Map` without initializing it ourselves.

```go
var m sync.Map

m.Store("cat", 31) // Store directly without panicking
value, ok := m.Load("cat")
```

Now we come to the main part. The same process can be written as a reusable flow that applies at every trie level:

![The repeated steps in Load.](/blog/go-sync-map-hash-trie/load-flow.webp)

<figcaption style="text-align: center; font-style: italic;">The repeated steps in Load.</figcaption>

After the four-bit hash groups lead us to an entry node, we still need to confirm that it contains the requested key.

The call to `n.entry().lookup(key)` runs this equality check (`==`). It starts with the entry stored in the child slot, then follows the overflow list until it finds a matching key or reaches the end. If it finds a match, it returns the stored value and `true`. If it reaches the end, it returns the zero value and `false`.

`Load` does not acquire a node mutex while reading. The only map-wide lock it may take is `initMu`, during the first lazy initialization. Concurrent calls to `Load` can therefore run at the same time.

So while `Load` is running, a writer may update one of its child pointers while holding the node lock, but `Load` does not care about or wait for that lock:

![Atomic child pointer publication.](/blog/go-sync-map-hash-trie/load-atomic-pointer.webp)

<figcaption style="text-align: center; font-style: italic;">Atomic child pointer publication.</figcaption>

`Load` reads each child pointer atomically. It sees either the old pointer or the new pointer, never a partially written pointer. The pointer is still safe to read even if another writer has replaced it in the map.

We can now look at how `Store` publishes that new pointer.

### Store

The public `sync.Map.Store` eventually reaches this internal method:

```go {hl_lines="2"}
func (ht *HashTrieMap[K, V]) Store(key K, new V) {
	_, _ = ht.Swap(key, new)
}
```

`Swap` here has to handle both possibilities:

- the key exists, so replace its value and return the previous value.
- the key does not exist, so insert it.

The writer first walks through the trie without a lock and stops at either a nil child or an entry:

```go {hl_lines=["6","11"]}
for {
	i = ht.root.Load()
	hashShift = 8 * goarch.PtrSize

	// find the node for the key
	for hashShift != 0 {
		...
	}

	// Lock the parent and check the child again.
	i.mu.Lock()
	n = slot.Load()
	if (n == nil || n.isEntry) && !i.dead.Load() {
		break
	}
	i.mu.Unlock()
}
```

We discussed the `for hashShift != 0 {}` loop in `Load`, where it walks the trie to find an entry node, so there is no need to repeat the traversal here. The writer then locks only the indirect node containing the selected child pointer, not every node in the map.

The diagram below shows how this local locking works.

- An update to A locks `root.mu`.
- Updates to B or C lock `X.mu`.
- An update to D locks `Y.mu`.

![Immediate-parent locking for writers.](/blog/go-sync-map-hash-trie/writer-immediate-parent-lock.webp)

<figcaption style="text-align: center; font-style: italic;">Immediate-parent locking for writers.</figcaption>

Therefore, updates to B and C must run one at a time because they share `X.mu`. Updates to B and D use separate mutexes and may run at the same time.

However, there is a window between the search loop (`for hashShift != 0`) and locking the indirect node (`i.mu.Lock()`). Another writer can update the selected child during that window. If the child changes, the result from the search loop is stale:

![Revalidation after acquiring the parent lock.](/blog/go-sync-map-hash-trie/writer-revalidation.webp)

<figcaption style="text-align: center; font-style: italic;">Revalidation after acquiring the parent lock.</figcaption>

So writer A in the diagram above must not continue using the pointer it read earlier, right? It's already "outdated".

Writer A reloads the slot with `slot.Load()` after acquiring the lock. If the selected child is now an indirect node or its parent is marked `dead`, the writer starts again from the root.

That is why the previous snippet has an outer loop:

```go {hl_lines=["1","9"]}
for {
	// find the node for the key
	for hashShift != 0 {
		...
	}

	// Lock the parent and check the child again.
	i.mu.Lock()
	n = slot.Load()
	if (n == nil || n.isEntry) && !i.dead.Load() {
		break
	}
	i.mu.Unlock()
	...
}
...
```

It can continue when the child is `nil` or an entry node and the parent indirect node has not been marked `dead` (deleted).

We can see the same pattern throughout insert, update, and delete operations. The writer reads without a lock, locks the immediate parent indirect node, checks the node state again, and then modifies and publishes the update.

When the key exists, the map checks the entry stored in the selected slot and then its overflow entries. After finding a matching key, it creates a _replacement_ entry with the new value and publishes the updated list:

```go {hl_lines="3"}
newEntry, old, swapped := oldEntry.swap(key, new)
if swapped {
	slot.Store(&newEntry.node)
	return old, true
}
```

What is a replacement?

![Replacing an entry without changing the surrounding overflow list.](/blog/go-sync-map-hash-trie/swap-entry-replacement.webp)

<figcaption style="text-align: center; font-style: italic;">Replacing an entry without changing the surrounding overflow list.</figcaption>

Replacement means the map does not modify the value in the existing entry. Instead, it creates another entry with the same key and the new value. A reader may still hold the old entry, and it continues to work as expected.

If the slot is nil, insertion is more direct:

```go
newEntry := newEntryNode(key, new)
if oldEntry == nil {
	slot.Store(&newEntry.node)
}
```

The entry is fully initialized before the atomic store publishes it.

But what if the writer checks the entire overflow list and still cannot find the key? Now `Swap` is inserting a new key, so the slot must hold both the existing entries and the new entry, right? At this point, `Swap` needs either

- an indirect node that becomes the root of a subtree containing both the old entry and the new entry, or
  ![Expanding an entry into an indirect node.](/blog/go-sync-map-hash-trie/swap-expand-entry.webp)
  <figcaption style="text-align: center; font-style: italic;">Expanding an entry into an indirect node.</figcaption>
- an overflow list when the old entry and the new entry have the same complete 64-bit hash. The new entry becomes the head of the overflow list.
  ![A full hash collision becomes an overflow list.](/blog/go-sync-map-hash-trie/swap-full-hash-overflow.webp)
  <figcaption style="text-align: center; font-style: italic;">A full hash collision becomes an overflow list.</figcaption>

### LoadOrStore and the compare operations

The way `LoadOrStore`, `CompareAndSwap`, and `CompareAndDelete` work is similar to `Store`. They use the same four-step write process:

1. Search the trie without a lock.
2. Lock the immediate parent of the selected slot.
3. Read the slot again and check that it is still valid.
4. Publish the replacement with an atomic pointer store.

The difference is the condition checked in the 3rd step, before the method changes the slot:

- `LoadOrStore` stores the new value only if the key is still missing. If another writer stores the key first, it simply returns that value.
- `CompareAndSwap` receives the expected current value as `old`. It replaces the value only if the key exists and its current value still matches that expectation.
- `CompareAndDelete` uses the same condition as `CompareAndSwap`, but removes the entry instead of replacing its value.

Although their names describe 2 actions, these methods combine the check and the update into one atomic operation. The mutex on the parent indirect node prevents another writer from changing the slot between those two steps.

### Delete

`Delete` is just a wrapper around `LoadAndDelete`. It ignores both results: the deleted value and the `loaded` boolean, which reports whether the key was present:

```go
func (ht *HashTrieMap[K, V]) Delete(key K) {
	_, _ = ht.LoadAndDelete(key)
}
```

Like `Store`, `LoadAndDelete` uses the same 4 steps described above. If the key is missing after the parent is locked and the slot is checked again, it returns `false` without changing the trie. When the key is still present, the deletion has three cases:

1. **The indirect node still contains another child, or it is the root.**

   The map clears the selected child pointer and then stops. A non-root indirect node may be left with one child, and the root stays even when all 16 of its child pointers are `nil`.

2. **A non-root indirect node becomes empty.**

   At this point, `LoadAndDelete` still holds the empty node's lock. It locks the parent, marks the empty node as `dead`, clears the pointer to it from the parent, and then unlocks the empty node. The parent stays locked because it may also need to be pruned.

   ![Removing empty indirect nodes.](/blog/go-sync-map-hash-trie/delete-pruning.webp)
   <figcaption style="text-align: center; font-style: italic;">Removing empty indirect nodes.</figcaption>

   If the next parent also becomes empty, the same process repeats with the next ancestor. This pruning stops at an indirect node that still has a child or at the root.

3. **The overflow list still contains another entry.**

   The map atomically replaces the child pointer with the rest of the list, then stops because the child is not empty.

   ![Deleting the first overflow entry.](/blog/go-sync-map-hash-trie/delete-overflow-entry.webp)
   <figcaption style="text-align: center; font-style: italic;">Deleting the first overflow entry.</figcaption>

What surprised me is that when an indirect node has only one entry node left, Go does not promote that entry to replace the indirect node.

What happens if `Store` and `Delete` race? `Store` may read a pointer to an indirect node before `Delete` removes that node from the trie.

![Store retry after concurrent deletion.](/blog/go-sync-map-hash-trie/store-delete-race.webp)

<figcaption style="text-align: center; font-style: italic;">Store retry after concurrent deletion.</figcaption>

This is just a quick question to exercise our memory. We already discussed the answer in the second bullet point above.

- While `Store` waits for the node's mutex, `Delete` marks the node as `dead`, clears the pointer from its parent, and unlocks the node.
- `Store` then acquires the mutex and checks the node again. The `dead` flag tells `Store` that its old pointer no longer belongs to the current trie, so it unlocks the node and retries from the current root.

### Clear

`Clear` does not walk through the trie and remove its entry nodes one by one. It creates a new empty root and replaces the current root pointer:

```go
func (ht *HashTrieMap[K, V]) Clear() {
	ht.init()
	ht.root.Store(newIndirectNode[K, V](nil))
}
```

`newIndirectNode(nil)` creates a root with all 16 `nil` child pointers. The atomic `root.Store` then makes that node the new root.

A concurrent operation may have loaded the previous root before `Clear` replaced it. That operation can still finish using the old nodes since those pointers are still valid. The old trie is no longer reachable from the map, and Go's garbage collector can reclaim it after no goroutine holds one of those pointers.

The number of stored entries does not change the work required by `Clear`, so the method itself is O(1).

### Range

`Range` loads the current root once and passes it to `iter`:

```go
func (ht *HashTrieMap[K, V]) Range(yield func(K, V) bool) {
	ht.init()
	ht.iter(ht.root.Load(), yield)
}
```

This is literally depth-first search (DFS). At each indirect node, `iter` reads the children from index 0 through index 15. A child that points to another indirect node starts a recursive call, and that call reads the entire branch before the loop moves to the next child.

When the child points to an entry node, `iter` passes its key and value to the callback and then reads any entries in its overflow list:

```go
func (ht *HashTrieMap[K, V]) iter(
	i *indirect[K, V],
	yield func(key K, value V) bool,
) bool {
	for j := range i.children {
		n := i.children[j].Load()
		if n == nil {
			continue
		}
		if !n.isEntry {
			if !ht.iter(n.indirect(), yield) {
				return false
			}
			continue
		}
		for e := n.entry(); e != nil; e = e.overflow.Load() {
			if !yield(e.key, e.value) {
				return false
			}
		}
	}
	return true
}
```

`Range` does not lock the nodes or copy all entries before it calls the callback, so another goroutine can insert, update, delete, or clear while `Range` is reading the trie.

One child may be read before a concurrent change, and another child may be read after it. So the callback can receive values that existed at different times during the same call.

Whatever order we get, the public API guarantees that `Range` calls the callback no more than once for each key.

## What the internals tell us

The hash trie source describes its main priority directly:

> The implementation is designed around frequent loads.

The code now explains that statement:

- `Load` hashes the key and follows atomic pointers without taking a node mutex or allocating a new trie node.
- `Store` has more work to do: it searches for an entry, locks the parent indirect node, reads the child again, and may allocate an entry or another indirect node. A concurrent structural change may also force it to retry from the root.

A workload with frequent reads is therefore the easiest case for `sync.Map`, because `Load` avoids contention on a shared lock.

But a write-heavy workload is not automatically a bad fit. Writes that reach separate indirect nodes can use separate mutexes, so several goroutines may update the trie at the same time. Of course, separate keys do not guarantee separate locks, as we discussed.

A plain map paired with an `RWMutex` has the opposite cost. Its lookup is direct and simple, but every operation uses the same lock. That design can be cheaper when concurrency is low because it avoids

- the trie traversal,
- atomic child pointers,
- extra node allocations.

Go's own benchmark suite compares `sync.Map` with an `RWMutexMap`. The following results came from an Apple M4 Pro, using the benchmark's parallel runner:

```sh
go test sync -run='^$' \
  -bench='BenchmarkMap(SwapCollision|SwapMostlyHits)$' \
  -benchmem -benchtime=200ms -count=1 -cpu=1,8
```

The `Memory` column lists `RWMutexMap` first and `sync.Map` second.

| Swap pattern                  | CPUs | `map + RWMutex` |  `sync.Map` |       Memory |
| ----------------------------- | ---: | --------------: | ----------: | -----------: |
| Repeatedly update one key     |    1 |     19.43 ns/op | 28.99 ns/op |  0 / 48 B/op |
| Repeatedly update one key     |    8 |     101.5 ns/op | 124.4 ns/op |  0 / 48 B/op |
| Mostly hits, 1 miss per 1,024 |    1 |     35.35 ns/op | 58.90 ns/op | 12 / 60 B/op |
| Mostly hits, 1 miss per 1,024 |    8 |     195.1 ns/op | 32.87 ns/op | 12 / 60 B/op |

In the same-key test, both maps must handle the writes one at a time because every goroutine updates the same key.

`sync.Map` creates a new entry and walks the hash trie, so `RWMutexMap` is faster with both 1 CPU and 8 CPUs.

The second benchmark is called `SwapMostlyHits`. The map starts with keys `0` through `1022`, and the benchmark follows a repeating pattern of 1,024 operations:

- The first 1,023 operations call `Swap` with one of those existing keys, so each operation finds its key in the map.
- The last operation calls `Swap` with a new key, so the map has to insert a new entry. The benchmark deletes that key immediately. The next round therefore starts with the same 1,023 existing keys.

This produces 1 miss for every 1,023 hits.

This mostly-hit test changes the result as expected.

- `RWMutexMap` wins with 1 CPU because its single lock has no competing goroutine.
- With 8 CPUs, that lock serializes both the updates and the occasional insertion and deletion.

The hash trie can update separate indirect nodes concurrently, so `sync.Map` has much higher throughput in this specific workload.

These benchmarks do not prove that either map is universally faster, but they show what your benchmark must get right about the real application: goroutine count, key distribution, map size, how often the application calls `Load`, `Store`, and `Delete`, etc.

## Bonus: What was wrong with the old sync.Map?

Before Go 1.24, `sync.Map` used two native maps.

- The `read` map handled the fast path without a lock,
- the `dirty` map held newer entries behind one mutex.

The process is simplified below:

1. A `Load` checked `read` first.
2. If the key was missing there but the `dirty` map contained newer keys, `Load` acquired the mutex, checked `dirty`, and incremented a miss counter.
3. Once the number of misses reached the size of `dirty`, the entire `dirty` map became the new `read` map.

![The promotion cycle in the old sync.Map implementation.](/blog/go-sync-map-hash-trie/old-sync-map-promotion.webp)

<figcaption style="text-align: center; font-style: italic;">New keys used the locked path until enough misses promoted the dirty map.</figcaption>

This design worked especially well when keys were written once and then read many times. Keys already stored in `read` could be loaded without locking. The problem was not that every operation took the mutex. The extra work happened when keys were frequently added or deleted:

- A new key stayed in `dirty` until the next promotion. Loading that key required the same mutex during this period.
- Creating a new `dirty` map required copying every non-deleted entry from `read` into another native map.
- Deleted entries could remain in `read` until a later promotion. A native map could also keep memory for slots that were no longer used.
- Writers changing unrelated keys could still wait for the same mutex whenever they needed the `dirty` map.

The hash trie does not use a `read` and `dirty` map together. A load follows atomic child pointers immediately. A write that changes the trie locks the closest indirect node. If a deletion leaves nodes empty, Go can remove those nodes from that key's path. There is no `dirty` map, miss counter, or promotion threshold. Writes to separate branches can also use separate node locks.

That does not mean the hash trie is always faster. A hit in the old `read` map was very cheap. The hash trie must walk through its nodes, load atomic pointers, and sometimes allocate new nodes. The main difference is that new keys no longer wait for a `dirty` map promotion, and writes to different branches do not always share one mutex.

## Source references

- [`src/sync/map.go`](https://go.dev/src/sync/map.go)
- [`src/internal/sync/hashtriemap.go`](https://go.dev/src/internal/sync/hashtriemap.go)
- [`src/sync/map_bench_test.go`](https://go.dev/src/sync/map_bench_test.go)
- [`sync.Map` package documentation](https://pkg.go.dev/sync#Map)
- [Go 1.24 release notes](https://go.dev/doc/go1.24#syncpkgsync)
- [The commit that made the hash trie the default](https://github.com/golang/go/commit/aecb8faa91e2b40e3651ee93c8e70635ed367677)
- [The commit that removed the old implementation](https://github.com/golang/go/commit/e15a14c4ddcb7854ecb5fb2f6dc01e8933a11652)
- [The proposal and benchmark discussion for replacing the old implementation](https://github.com/golang/go/issues/70683)
- [The Go memory model](https://go.dev/ref/mem)

## Who We Are

We're VictoriaMetrics, a team providing open-source, highly scalable, and cost-efficient solutions for monitoring, logging, and tracing, trusted by users worldwide to reduce their observability costs. Check out our [VictoriaMetrics](https://docs.victoriametrics.com/), [VictoriaLogs](https://docs.victoriametrics.com/victorialogs/), and [VictoriaTraces](https://docs.victoriametrics.com/victoriatraces/) for more details.
