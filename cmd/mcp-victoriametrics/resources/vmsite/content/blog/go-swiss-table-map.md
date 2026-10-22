---
draft: false
page: blog blog_post
authors:
  - Phuong Le
date: 2026-09-03
title: "How Go’s Built-in Map Works with Swiss Tables"
summary: "Go 1.24 replaced the built-in map's bucket-based runtime with Swiss Tables. This article explains groups, control bytes, H1 and H2, probing, table growth, directories, deletion, load factor, and the experimental split-group layout."
enableComments: true
toc: true
categories:
  - Go @ VictoriaMetrics
  - Open Source Tech
tags:
  - go
  - golang
  - map
  - runtime
images:
  - /blog/go-swiss-table-map/preview.webp
---

We have already written about Go maps and their old runtime implementation in [Go Maps Explained: How Key-Value Pairs Are Actually Stored](/blog/go-map/). Go 1.24 replaced that implementation with a design based on Swiss Tables, so it is time for an update.

You do not need to go back and read the old article. We will explain the new runtime internals step by step and introduce each concept as we need it.

_The Go blog also has an excellent article, [Faster Go maps with Swiss Tables](https://go.dev/blog/swisstable). It goes deeper and assumes a little more background knowledge. We take a different approach. We will discuss the same implementation more gradually and in a visual way, so you can relax your brain a little and still understand what Go is doing._

## What is a map at runtime?

Let's start with what a map actually is.

```go
m := make(map[string]int)
```

`make` initializes the map. `map[string]int` is the language-level type, which tells us that the map uses strings as keys and integers as values. Underneath that type, the runtime representation of `m` is a pointer to `internal/runtime/maps.Map`.

```go
type Map struct {
	used uint64
	seed uintptr

	dirPtr unsafe.Pointer
	dirLen int
	...
}
```

We can easily inspect this with `println`, which prints that pointer:

```go
m := make(map[string]int)
m2 := m

println(m)  // 0x14000122000
println(m2) // 0x14000122000
```

Copying `m` to another map variable copies this pointer, so both variables refer to the same runtime `Map` and the same entries.

![Copying a map variable makes m and m2 point to the same runtime Map.](/blog/go-swiss-table-map/map-copy-pointer.webp)

<figcaption style="text-align: center; font-style: italic;">Copying a map variable makes m and m2 point to the same runtime Map.</figcaption>

The 2 fields at the top describe the map itself, not the storage for its entries.

```go {hl_lines=["2","3"]}
type Map struct {
	used uint64
	seed uintptr
	...
}
```

`used` counts how many entries are currently stored. Since Go knows exactly where to find the number of entries, when you write `len(m)`, Go replaces this call with an access to the first field of `Map` and converts it to an `int`. That is why `len(m)` is O(1) instead of scanning the entire map.

`seed` is an interesting field because it causes different maps to distribute the same keys differently. Go initializes this field with a random number for every map.

![The same entries are arranged differently when maps use different seeds.](/blog/go-swiss-table-map/map-seed-distribution.webp)

<figcaption style="text-align: center; font-style: italic;">The same entries are arranged differently when maps use different seeds.</figcaption>

_The array above is only a simplified representation used for this explanation. The actual data structure is more complicated._

When Go hashes a key to find where it is stored, it uses both the key and the map's seed to calculate the hash. Each map receives its own seed, so the same key can produce different hash values in 2 maps and end up in different storage locations.

### Storage layout

Go does not use the same storage layout for every map. The layout changes based on how many key-value pairs the map stores.

Before we go into the details, let’s spoil what the map will look like in the 3 main cases:

![Three example layouts for maps with 4, 100, and 1,000 keys.](/blog/go-swiss-table-map/map-storage-overview.webp)

<figcaption style="text-align: center; font-style: italic;">The map adds tables and a directory as it stores more entries.</figcaption>

Here we can already see the 3 main parts of the data structure: a **directory**, one or more **tables**, and the **groups** inside each table. A group is the smallest storage unit. It is where Go stores the key-value pairs.

### Group

We will start with a small map that only needs 1 group to store its key-value pairs. A group can hold up to 8 pairs, so the map does not need a directory or a table yet. Instead, `dirPtr` points directly to the group.

A group is the smallest piece of storage that Go’s Swiss Table implementation checks at one time. Each group contains:

- 8 slots for key-value entries.
- 8 control bytes, one for each slot. Go stores these 8 bytes together in one `uint64`.

![A group pairs each control byte with the key-value slot below it.](/blog/go-swiss-table-map/group-anatomy.webp)

<figcaption style="text-align: center; font-style: italic;">A group pairs each control byte with the key-value slot below it.</figcaption>

The group's concrete type depends on the map’s key and value types, so the compiler generates an internal anonymous struct for each map type. Conceptually, `map[string]int` has this layout:

```go
type group struct {
	ctrl uint64

	slots [8]struct {
		key  Key
		elem Elem
	}
}
```

Go is also testing a new group layout with separate key and value arrays to improve key lookup locality and remove repeated alignment padding, as explained in the [split group layout](#the-split-group-layout) section.

#### Control bytes and the control word

Let's first look at the top row of the group. These are the 8 control bytes. Together, they form the 8-byte **control word**.

Each control byte describes the slot directly below it, so control byte 0 belongs to slot 0, control byte 1 belongs to slot 1, and the same relationship continues through slot 7.

But where do those bytes come from?

Go hashes the key using the `seed` from `Map`, then divides that hash into 2 parts. On most 64-bit targets, the upper 57 bits are called **H1**, and the lower 7 bits are called **H2**. Suppose we have another key, `"cow"`, which produces H2 `42` in our illustration:

![A 64-bit hash contains H1 and a 7-bit H2.](/blog/go-swiss-table-map/hash-split.webp)

<figcaption style="text-align: center; font-style: italic;">A 64-bit hash contains H1 and a 7-bit H2.</figcaption>

> [!NOTE]
> Go uses a 32-bit hash layout on 32-bit targets (and Wasm). We will follow the 64-bit layout in the rest of this article.

H1 is the first part of the hash that Go uses to choose where a search starts in the map’s storage. A small map has only 1 group, so there is nothing to choose. Let's leave it aside until the map grows.

H2 is the part stored in the control byte above a live slot.

But a control byte has 8 bits, while H2 uses only 7, so we still have 1 bit left. Go uses this highest bit to tell whether the slot contains a live entry or a special state. If this bit is `0`, the lower 7 bits contain H2. If this bit is `1`, the complete control byte represents `empty` or `deleted`:

![A control byte represents a live, empty, or deleted slot.](/blog/go-swiss-table-map/control-byte-states.webp)

<figcaption style="text-align: center; font-style: italic;">A control byte represents a live, empty, or deleted slot.</figcaption>

When its slot contains a key-value entry, the highest bit is `0`, while the lower 7 bits contain H2. H2 `42` is `0101010` in binary, so the complete control byte for `"cow"` is `00101010`.

When the highest bit is `1`, the control byte stores a special value instead of H2. An empty slot uses `10000000`. A deleted slot uses `11111110` and is also called a **tombstone**. Both states contain no live key-value entry, but a lookup can stop at `empty` while it must continue past `deleted`. We will return to this distinction in the [deletion](#deletion) section.

With this layout, a control byte lets Go answer 2 questions before it reads the complete key from a slot:

- Does this slot contain a live entry, or is it `empty` or `deleted`?
- If the slot contains a live entry, could its key be the one we are looking for?

Now return to the original group above:

![A group pairs each control byte with the key-value slot below it.](/blog/go-swiss-table-map/group-anatomy.webp)

<figcaption style="text-align: center; font-style: italic;">A group pairs each control byte with the key-value slot below it.</figcaption>

Next, assign a value to the `"cow"` key that produced H2 `42`:

```go
m["cow"] = 4
```

Before storing `"cow"`, Go must know whether this assignment updates an existing key or adds a new one. It uses H2 to find slots that may already store the key, then confirms each candidate with a complete key equality check.

- H2 for `"cow"` is `42`, which is also the value stored in the control byte of `"dog"`.
- Go then compares the complete keys with an equality check (`==`), but `"dog"` is not equal to `"cow"`.
- No other control byte contains H2 `42`, so Go knows that this assignment is adding a new key.

The map selects the first empty slot in the group, which is slot 2, writes `"cow"` and `4` into that slot, then writes H2 `42` into control byte 2 directly above it:

![Cow uses the first empty slot and stores H2 42 above it.](/blog/go-swiss-table-map/group-insert-cow.webp)

<figcaption style="text-align: center; font-style: italic;">Cow uses the first empty slot and stores H2 42 above it.</figcaption>

Now the group contains 2 control bytes with the same H2 value, `42`: one above `"dog"` and one above `"cow"`. Suppose we later assign another value to `"cow"`:

```go
m["cow"] = 5
```

Before Go can update the value, it must find the existing key. It takes H2 `42` from the hash and compares it with all 8 control bytes in the group at once:

![H2 42 selects dog and cow as candidate keys.](/blog/go-swiss-table-map/control-word-small-map.webp)

<figcaption style="text-align: center; font-style: italic;">H2 42 selects dog and cow as candidate keys.</figcaption>

Go does not visit the 8 slots one by one and compare H2 with each control byte separately. On AMD64, Go uses SIMD instructions to compare H2 `42` with those 8 control bytes at the same time.

SIMD lets the CPU apply the same comparison to several byte values in parallel. On AMD64, the result is a packed bitmap with one bit for each slot:

![One control-word comparison produces the candidate bitmap.](/blog/go-swiss-table-map/simd-h2-bitmap.webp)

<figcaption style="text-align: center; font-style: italic;">One control-word comparison produces the candidate bitmap.</figcaption>

In our group, the bits for slots 0 and 2 are set because both control bytes contain `42`. The other bits are clear, which masks out the other 6 slots without reading their complete keys.

Other architectures produce the same candidate mask with arithmetic and bitwise operations on the 64-bit control word, but they use one byte per slot instead of packing the result into 8 bits.

Go then reads the complete keys from slots 0 and 2 and compares them with `"cow"`. `"dog"` fails the equality check (`==`), while `"cow"` matches, so the assignment updates the value stored for `"cow"`.

### Table

A group has only 8 slots, so if we keep adding key-value pairs beyond its capacity, Go needs another storage structure to store them.

Go doubles the number of groups from 1 to 2 and introduces a new structure called a **table** to manage them. It moves the 8 existing entries from the small group into that table, redistributes them between the 2 groups, and then stores the new entry.

![The ninth entry turns one group into a table with two groups.](/blog/go-swiss-table-map/group-to-table.webp)

<figcaption style="text-align: center; font-style: italic;">The ninth entry turns one group into a table with two groups.</figcaption>

A table is a complete Swiss Table that owns one or more groups together:

```go
type table struct {
	used       uint16
	capacity   uint16
	growthLeft uint16
	...

	groups groupsReference
}
```

Our first table has `capacity = 16`, `used = 9`, and 2 groups:

- `groups` points to the contiguous allocation that contains the groups.
- `capacity` counts all slots across those groups.
- `used` counts the live entries in this table.

Now, why do the existing key-value pairs in the first group need to be redistributed, and how does Go know which of the 2 groups each pair should go to?

Let me introduce H1, which is used for this exact purpose. Go uses H1 to calculate the starting group for each key:

```text
starting group = H1 % number of groups
```

Since this table has 2 groups, `% 2` only needs the lowest bit of H1. When we write H1 with its highest bit on the left, the lowest bit is the rightmost bit, directly beside H2 in the original hash.

For example, suppose H1 for `"cow"` ends in `0`, while H1 for `"dog"` ends in `1`:

![The lowest H1 bit redistributes entries between two groups.](/blog/go-swiss-table-map/table-groups.webp)

<figcaption style="text-align: center; font-style: italic;">The lowest H1 bit redistributes entries between two groups.</figcaption>

Because the number of groups changed, Go runs this calculation again for each existing key. Some key-value pairs are inserted into the new group 0, while others are inserted into the new group 1.

This calculation gives us only the starting group. Let's use a different map with 16 slots and 10 live entries to see why this matters. Its keys are distributed unevenly, leaving group 0 full while group 1 still has empty slots:

![One group can be full while another still has empty slots.](/blog/go-swiss-table-map/uneven-groups.webp)

<figcaption style="text-align: center; font-style: italic;">One group can be full while another still has empty slots.</figcaption>

In this case, another key may also select the full group 0 above as its starting group. Go cannot store the key there, so it checks group 1 next. The list of groups Go checks, together with the order used to check them, is called the **triangular probe sequence**.

> [!NOTE] What is the triangular probe sequence?
> Let's say the table has 8 groups, and H1 selects group 3 as the starting group for a key. If group 3 has no empty slot, Go needs to check other groups.
>
> Go uses quadratic probing between groups. Starting from group 3, it moves by `+1`, then `+2`, then `+3`, wrapping around the table. The resulting sequence is `3, 4, 6, 1, 5, 2, 0, 7`.
>
> ![An eight-group triangular probe sequence.](/blog/go-swiss-table-map/triangular-probe-sequence.webp)
>
> <figcaption style="text-align: center; font-style: italic;">An eight-group triangular probe sequence.</figcaption>
>
> The cumulative offsets are `0, 1, 3, 6, 10, ...`, the triangular numbers, which is where the name comes from. Because the group count is a power of two, Go visits every group exactly once before the sequence repeats.
>
> Why not simply check the next group each time: `3, 4, 5, 6, 7, 0, 1, 2`?
>
> Imagine that groups 3, 4, and 5 are already full:
>
> 1. A new key that starts at group 3 has to pass all 3 groups before it can use group 6.
> 2. New keys that start at group 4 or 5 also reach group 6 because the groups before it are full.
> 3. When group 6 becomes full, these new insertions continue to group 7.
>
> Over time, the full area keeps growing in one direction, and Go has to check more and more groups to insert a new key that starts inside it. This is called **clustering**.
>
> The triangular sequence avoids always pushing these keys into the next neighboring group. Its jumps get larger, so after checking groups 3 and 4, it jumps to group 6, then group 1, then group 5. This spreads the collisions around the table instead of extending the same full area. Inside each group, Go still checks all 8 slots together.

The storage pointer in `Map` also changes when the small map becomes a table-backed map. `dirPtr` no longer points directly to a group. It points to a one-entry array, and that entry points to the table.

The runtime calls this pointer array the **directory**. Here is how it connects our 10-entry map to its table:

![dirPtr reaches the table through directory entry 0.](/blog/go-swiss-table-map/table-backed-map.webp)

<figcaption style="text-align: center; font-style: italic;">dirPtr reaches the table through directory entry 0.</figcaption>

The table above starts with 2 groups. Since each group contains 8 slots, this table currently has 16 slots.

If we keep adding keys, Go can replace it with a larger table containing 4 groups, then 8 groups, then 16 groups, and continue doubling the number of groups when more storage is needed. One table can grow up to 128 groups, giving it a maximum capacity of 1024 slots:

```text
128 groups * 8 slots = 1024 slots
```

If another insertion requires more storage, Go splits the table into 2 tables instead. We will follow that split shortly.

But after the small map with 1 group becomes a table with 2 groups, Go uses a different growth threshold:

- The map does not wait for all groups in the table to become full before doubling the group count from 2 to 4, from 4 to 8, and so on.
- Once a table has grown to 1024 slots, it also does not wait for every slot to contain an entry before splitting it into 2 tables.

A regular table reaches its insertion limit before every slot is used, and that limit is controlled by the **load factor**. It tells Go how full the table is. Go includes both live entries and deleted slots in this calculation:

```text
load = (live entries + deleted slots) / table slots
```

Our table has no deleted slots, so with 16 slots and 10 live entries, its load factor is `10 / 16`, or `62.5%`.

Go does not let a regular table reach `100%`. Its maximum load factor is `7 / 8`, or `87.5%`, which means that live entries and deleted slots together may account for 7 out of every 8 slots on average. The limit applies across the whole table.

> [!NOTE]
> We explain why Go uses a `7 / 8` load factor in [Why load factor?](#why-load-factor).

The table field `growthLeft` tracks how many additional empty slots new keys may consume before the table reaches that limit.

```go {hl_lines="4"}
type table struct {
	used       uint16
	capacity   uint16
	growthLeft uint16
	...

	groups groupsReference
}
```

A 16-slot table containing only live entries and empty slots has an insertion limit of 14 entries because `16 * 7 / 8 = 14`. After the 10 entries above are stored, `growthLeft` is `4`.

`growthLeft` reaches `0` after the table stores 14 live entries. If another new key needs an empty slot, Go doubles the table's number of groups if the table has fewer than 1024 slots.

![A table doubles until it reaches 1024 slots.](/blog/go-swiss-table-map/table-growth.webp)

<figcaption style="text-align: center; font-style: italic;">A table doubles until it reaches 1024 slots.</figcaption>

Each time the table doubles its number of groups, Go hashes every live key in that table again and redistributes all its entries across the new groups.

A table with 1024 slots already contains the maximum of 128 groups. In this case, Go **splits** the selected table into 2 new tables instead, each with 128 groups and 1024 slots.

Why not keep adding groups to the same table?

Each time a table grows, Go needs to hash its existing keys again and move its entries into the new groups. With one large table, an insertion that triggers growth would have to wait for all those entries to move.

Go uses multiple smaller tables to avoid that. When one table needs more space, Go only rebuilds that table. Entries in the other tables stay where they are. Each table has at most 1024 slots, which limits how many entries each rebuild needs to move.

> [!NOTE]
> Here, when we say **split**, we do not mean that the number of tables doubles. The number of tables increases gradually because only the table that needs more space splits into 2 new tables. The other tables are unchanged.

This split uses a different end of H1 than the calculation that chooses a group. When the number of groups doubles, Go uses one more bit from the right side of H1, where the low bits are, to calculate the starting group inside the same table.

So when a table splits, Go uses the next unused bit from the left side of H1, where the high bits are, to choose between the 2 new tables.

Suppose H1 for `"dog"` begins with `0`, while H1 for `"cat"` begins with `1`. The `0` sends `"dog"` to table 0, and the `1` sends `"cat"` to table 1:

![The leftmost H1 bit separates the first two tables.](/blog/go-swiss-table-map/grow-split.webp)

<figcaption style="text-align: center; font-style: italic;">The leftmost H1 bit separates the first two tables.</figcaption>

So the map hashes the key again with its seed, reads the same leftmost H1 bit, and inserts the key-value pair into the selected new table. After all live entries have moved, Go retries the insertion that triggered the split, just as it does when doubling the number of groups.

### Directory

Before getting into the directory, let's understand what problem it solves.

#### Redistribution

If Go redistributed every key-value pair in the map whenever one table needed more storage, each growth operation would become more and more expensive as the map became larger. Instead, Go only rebuilds the table that needs more storage.

**When the number of groups in a table doubles, Go redistributes only the live key-value pairs from that table across the new groups.**

The leftmost H1 bits select a table, while the rightmost H1 bits select a starting group inside that table.

When the number of groups changes, Go only changes how many rightmost H1 bits that table uses for its group selection.

![The left and right ends of H1 select a table and a group.](/blog/go-swiss-table-map/h1-table-group-selection.webp)

<figcaption style="text-align: center; font-style: italic;">The left and right ends of H1 select a table and a group.</figcaption>

It does not change the leftmost H1 prefix that selected the table. A key stored in another table has a different leftmost H1 prefix, so it continues to select that table and does not take part in this redistribution.

**When a table splits into 2 tables, Go redistributes only the live key-value pairs from the original table between the 2 new tables.**

When the map has 2 tables, the directory uses the leftmost H1 bit to select between them: `0` points to table 0 and `1` points to table 1. If table 1 splits, the directory increases the number of leftmost H1 bits it reads from 1 to 2:

![Splitting table 1 leaves table 0 unchanged.](/blog/go-swiss-table-map/redistribute-table-split.webp)

<figcaption style="text-align: center; font-style: italic;">Splitting table 1 leaves table 0 unchanged.</figcaption>

H1 prefixes `00` and `01` select directory entries `0` and `1`, which both point to table 0 because table 0 has not split. Prefix `10` selects entry `2`, which points to new table 1, while prefix `11` selects entry `3`, which points to new table 2.

You can see that the map now contains 3 tables, but 2 leftmost H1 bits produce 4 possible combinations, right?

This is where the directory becomes useful. It has 4 entries, one for each bit combination, but those entries do not need to point to 4 different tables.

![Four directory entries can point to three tables.](/blog/go-swiss-table-map/map-directory-after-split.webp)

<figcaption style="text-align: center; font-style: italic;">Four directory entries can point to three tables.</figcaption>

During a lookup, Go uses the 2 leftmost H1 bits to select a directory entry, and that entry tells Go which table to search. This allows one table to split without requiring every other table in the map to split with it.

In other words, the leftmost H1 bits select a directory entry, not a table directly, and that directory entry tells Go which table may contain the key.

### Global and local depth

Let's say table 0 has reached its insertion limit and needs to split. However, we cannot simply split it as we did with table 1, because 2 directory entries point to table 0.

Table 0 therefore needs a way to know:

- How many directory entries point to it?
- Does the directory already have enough entries for the 2 new child tables?

This is what **global depth** and **local depth** tell us.

```go {hl_lines=["5","10"]}
type Map struct {
	dirPtr unsafe.Pointer
	dirLen int

	globalDepth uint8
	...
}

type table struct {
	localDepth uint8
	...
}
```

Global depth belongs to the whole directory. It is the number of high hash bits used to select one directory entry. In this case, we have 4 directory entries selected by the 2 leftmost bits of H1 (`00`, `01`, `10`, and `11`), so `globalDepth = 2`.

Local depth belongs to each table. It is the number of high hash bits needed to identify that table. In this case:

- Every key with an H1 prefix of `0` chooses table 0. Table 0 is identified by the leftmost bit (`0`), so its `localDepth = 1`.
- Every key with an H1 prefix of `10` chooses table 1, while every key with an H1 prefix of `11` chooses table 2. Tables 1 and 2 need the 2 leftmost bits, so each has `localDepth = 2`.

![Four directory entries point to three tables.](/blog/go-swiss-table-map/directory.webp)

<figcaption style="text-align: center; font-style: italic;">Four directory entries point to three tables.</figcaption>

Table 0 knows that its local depth (`1`) is less than the global depth (`2`), which means the directory already has enough entries for the split. Go can make entry `0` point to one child and entry `1` point to the other without growing the directory.

![Table 0 splits into tables 0 and 1 without growing the directory.](/blog/go-swiss-table-map/directory-after-table-0-split.webp)

<figcaption style="text-align: center; font-style: italic;">Table 0 splits into tables 0 and 1 without growing the directory.</figcaption>

What if table 1 reaches its insertion limit and needs to split instead of table 0? In this case, its `localDepth` (`2`) is equal to the map's `globalDepth` (`2`). The map first doubles the directory from 4 entries to 8 and increases `globalDepth` to `3`. It then splits table 1 into tables 1 and 2:

![Splitting table 1 doubles the directory before creating tables 1 and 2.](/blog/go-swiss-table-map/directory-after-table-1-split.webp)

<figcaption style="text-align: center; font-style: italic;">Splitting table 1 doubles the directory before creating tables 1 and 2.</figcaption>

So far, so good. Let's recap what we have discussed so far using only a key's hash.

The highest bits of H1 select a directory entry and therefore a table. The lower bits of H1 select a starting group inside that table. H2 filters the 8 slots in the group. Any H1 bits between the directory selection and the group selection may be unused for the current map shape:

![The hash supplies bits for a directory entry, a group, and H2.](/blog/go-swiss-table-map/hash-anatomy.webp)

<figcaption style="text-align: center; font-style: italic;">The hash supplies bits for a directory entry, a group, and H2.</figcaption>

## What changed in Go 1.24?

After following the new implementation from a group to a directory, let's compare it with the implementation that Go used before version 1.24.

The old implementation stored up to 8 key-value pairs in one bucket. If that bucket could not hold another entry, the runtime could connect an **overflow bucket** that provided 8 more slots.

If the overflow bucket also became full, another overflow bucket could follow it:

![A bucket points to a chain of overflow buckets.](/blog/go-swiss-table-map/old-map-overflow.webp)

<figcaption style="text-align: center; font-style: italic;">A bucket points to a chain of overflow buckets.</figcaption>

An overflow bucket allowed 1 full bucket to receive more space without redistributing the other buckets. But an overflow chain also introduced several costs:

- A lookup had to load the overflow pointer before it could find and check the next 8 slots.
- A longer chain repeated that pointer load and bucket scan for every additional overflow bucket.
- Extending the chain could require another allocation.

Go cannot read the overflow bucket until it has loaded the pointer from the current bucket. Every additional overflow bucket adds another dependent pointer load:

```
bucket
  -> load pointer
overflow 1
  -> load pointer
overflow 2
```

The new implementation allocates a table's groups together in one array. It already knows exactly where each group is and uses the triangular probe sequence to find another group inside that table when the starting group has no empty slot:

![Old maps use overflow buckets while new maps probe table groups.](/blog/go-swiss-table-map/old-vs-new-memory.webp)

<figcaption style="text-align: center; font-style: italic;">Old maps use overflow buckets while new maps probe table groups.</figcaption>

Growth is the 2nd major change.

The old implementation had one primary bucket array for the map. When the bucket array doubled, Go created a new array with twice as many buckets. The key-value pairs from each old bucket were then distributed between 2 buckets in the new array.

![Entries from old bucket 0 move into new buckets 0 and 4.](/blog/go-swiss-table-map/old-map-growth.webp)

<figcaption style="text-align: center; font-style: italic;">Entries from old bucket 0 move into new buckets 0 and 4.</figcaption>

Go repeated this process gradually for every old bucket as we assigned or deleted entries, so both arrays stayed alive until every key-value pair had moved.

In contrast, Go's Swiss Table growth does not move one group at a time, as we discussed. The new implementation divides a large map into tables with at most 1024 slots.

One assignment may rebuild one complete table, but the advantage is that it does not rebuild the entry storage of the other tables.

![Swiss Table growth rebuilds only the selected table.](/blog/go-swiss-table-map/swiss-table-growth.webp)

<figcaption style="text-align: center; font-style: italic;">Swiss Table growth rebuilds only the selected table.</figcaption>

In the Go team's microbenchmarks, map operations ran up to 60% faster than in Go 1.23, although some edge cases became slower. Full application benchmarks showed a geometric mean CPU time improvement of around 1.5%.

That’s all we want to discuss about the data structure. Let’s summarize the design and what each part helps with:

| Part | What it helps with |
| --- | --- |
| Control bytes + groups of 8 | Compare H2 with 8 control bytes together, then compare complete keys only in candidate slots. |
| Open addressing | Look for space in the same table array, without adding overflow buckets or following their pointers. |
| Triangular probing | Reduce the long runs of full groups that linear probing can build up. |
| Load factor of `7 / 8` | Keep empty slots so a missing lookup can stop early, before it has to check every group. We will explain this in [Why load factor?](#why-load-factor) below. |
| Directory + tables | Let one table grow or split without moving entries from every other table. |
| Maximum of 1024 slots per table | Limit the entries moved by each table rebuild, helping avoid long delays in the insertion that triggers it. |

## Bonus

### Deletion

Let's return to the 3 control states we introduced earlier, this time focusing on `deleted`:

![The deleted control state is the focus of this section.](/blog/go-swiss-table-map/control-byte-states-deletion.webp)

<figcaption style="text-align: center; font-style: italic;">The deleted control state is the focus of this section.</figcaption>

An `empty` slot and a `deleted` slot both contain no live key-value entry, but they have different meanings during lookup.

During a lookup:

1. Go first checks whether any candidate slot in the current group contains the requested key.
2. If Go does not find the key in those candidate slots and the group contains an `empty` control byte, Go can stop early.
3. A `deleted` control byte cannot provide the same guarantee because another key may have been inserted into a later group before this slot was deleted, so Go must continue checking other groups in the probe sequence.

The `empty` state does not mean no keys exist after the empty slot or in later groups. Other keys may still be stored there. It only means the requested key cannot sit in a later group along the same probe sequence.

But deleting a key does not always produce a `deleted` slot. In a small map with only 1 group, the removed slot always becomes `empty` because lookup has no later group to continue to.

In a larger table-backed map, the removed slot also becomes `empty` if its group already contains at least 1 other `empty` slot. If the group has no `empty` slot, Go marks the removed slot as `deleted` so lookup can continue to later groups.

### Why load factor?

Earlier, we said Go lets live entries and deleted slots account for up to `7 / 8` of a table before the next new key needs table maintenance. This ratio is the table's **load factor**:

```text
load factor = (live entries + deleted slots) / table slots
```

Our current table has 16 slots and no deleted slots. A load factor of `7 / 8` gives it space for 14 live entries:

```text
14 / 16 = 7 / 8 = 87.5%
```

Why does Go stop here instead of filling all 16 slots?

The triangular probe sequence chooses the group indexes, and Go scans the selected groups one at a time. This work is a linear scan across the probed groups, so checking 4 groups requires 4 separate group checks.

Now extend the same load factor to a 32-slot table. It can store 28 live entries and still keep 4 empty slots. All 4 empty slots may sit in one group, while the other 3 groups are full:

![An empty slot stops lookup before 2 later groups.](/blog/go-swiss-table-map/load-factor.webp)

<figcaption style="text-align: center; font-style: italic;">An empty slot stops lookup before 2 later groups.</figcaption>

Suppose a missing lookup starts at group 0. With 4 groups, the triangular probe sequence is `0 -> 1 -> 3 -> 2`. Go checks group 0, then group 1. The empty control bytes in group 1 stop the lookup, so Go never checks groups 3 and 2 even though both contain entries.

As the table fills, empty slots become harder to find. A missing lookup may scan more groups before it reaches an empty control byte. If every slot were full, no empty control byte could stop the scan, so Go would need to check every group before reporting a missing key.

The `7 / 8` limit keeps some empty slots in the table before this scanning cost becomes too high.

Could Go choose another limit? Yes.

- A lower limit would use more memory but usually shorten the probe sequence.
- A higher limit would leave less unused memory but usually lengthen the probe sequence.

Go currently uses the same `7 / 8` limit as [Abseil's Swiss Table](https://abseil.io/about/design/swisstables).

### The split group layout

There is one more map change worth discussing. Go 1.27 includes an experimental group layout named `mapsplitgroup`, which we can enable with `GOEXPERIMENT=mapsplitgroup`.

Despite the name, this experiment does not change the map algorithm, the hash split, or the probe sequence. It only changes how one group arranges its 8 keys and values in memory.

By default, Go 1.27 stores each key beside its value:

```text
control:  42     17     -      -
slots:    dog:1  cat:2  empty  empty
```

The internal layout is approximately:

```go
slots [8]struct {
	key   string
	value int
}
```

The experimental layout separates the keys from the values:

```text
control:  42     17     -      -
keys:     dog    cat    empty  empty
values:   1      2
```

Its internal layout is approximately:

```go
keys   [8]string
values [8]int
```

The 2 arrays still describe the same entries. At index 1, `keys[1]` stores `"cat"`, and `values[1]` stores `2`.

Let's look up `"cat"`. Its control byte identifies index 1 as a candidate, so Go reads `keys[1]` and compares it with `"cat"`. But do we need `values[1]` at this point? No:

![Lookup reads keys[1] before values[1].](/blog/go-swiss-table-map/split-group-lookup.webp)

<figcaption style="text-align: center; font-style: italic;">Lookup reads keys[1] before values[1].</figcaption>

Go reads `values[1]` only after the key matches. This is why the new layout keeps the keys together. During the search, Go reads candidate keys, while the values become useful only after one key matches. The values no longer sit between those key reads.

The split layout can also use less memory. Consider `map[int64]struct{}`, a map often used as a set. Each key needs 8 bytes, and the empty value stores no data. But the old layout still uses 16 bytes for every `{key, value}` slot because the empty value at the end adds 8 bytes of padding.

The control word is unchanged, so we only need to compare the storage below it:

![The old layout uses 128 bytes below the control word.](/blog/go-swiss-table-map/split-group-padding-old.webp)

<figcaption style="text-align: center; font-style: italic;">The old layout uses 128 bytes below the control word.</figcaption>

The old slot layout uses 128 bytes across 8 slots. The 8 keys use 64 bytes, but each slot also contains 8 bytes of padding, adding another 64 bytes.

The split layout keeps all 8 keys together instead:

![The split layout uses 72 bytes below the control word.](/blog/go-swiss-table-map/split-group-padding-new.webp)

<figcaption style="text-align: center; font-style: italic;">The split layout uses 72 bytes below the control word.</figcaption>

The 8 keys still use 64 bytes. The empty value array adds no element data, but Go adds one 8-byte trailing padding area after it. The storage below the control word therefore uses 72 bytes in total, saving 56 bytes in one group.

Go introduced this layout behind `GOEXPERIMENT=mapsplitgroup` in [CL 711560](https://go-review.googlesource.com/c/go/+/711560), then enabled it by default on the development branch in [CL 820500](https://go-review.googlesource.com/c/go/+/820500).

## Resources

- [The Go specification: Map types](https://go.dev/ref/spec#Map_types)
- [Go 1.24 release notes: Runtime](https://go.dev/doc/go1.24#runtime)
- [Faster Go maps with Swiss Tables](https://go.dev/blog/swisstable)
- Go 1.27 runtime source: [`map.go`](https://go.googlesource.com/go/+/refs/tags/go1.27.0/src/internal/runtime/maps/map.go), [`table.go`](https://go.googlesource.com/go/+/refs/tags/go1.27.0/src/internal/runtime/maps/table.go), [`group.go`](https://go.googlesource.com/go/+/refs/tags/go1.27.0/src/internal/runtime/maps/group.go), and [`runtime.go`](https://go.googlesource.com/go/+/refs/tags/go1.27.0/src/internal/runtime/maps/runtime.go)
- [Go 1.27 map ABI and group layout](https://go.googlesource.com/go/+/refs/tags/go1.27.0/src/internal/abi/map.go)
- [Go 1.23.12 map implementation](https://go.googlesource.com/go/+/refs/tags/go1.23.12/src/runtime/map.go)
- [Abseil Swiss Tables Design Notes](https://abseil.io/about/design/swisstables)
- Split group layout: [CL 711560](https://go-review.googlesource.com/c/go/+/711560) and [CL 820500](https://go-review.googlesource.com/c/go/+/820500)
- [Go Maps Explained: How Key-Value Pairs Are Actually Stored](/blog/go-map/)

## Who We Are

We're VictoriaMetrics, a team providing open-source, highly scalable, and cost-efficient solutions for monitoring, logging, and tracing, trusted by users worldwide to reduce their observability costs. Check out our [VictoriaMetrics](https://docs.victoriametrics.com/), [VictoriaLogs](https://docs.victoriametrics.com/victorialogs/), and [VictoriaTraces](https://docs.victoriametrics.com/victoriatraces/) for more details.
