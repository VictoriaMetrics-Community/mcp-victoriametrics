---
draft: false
page: blog blog_post
authors:
  - Phuong Le
date: 2025-02-02
title: "How vmstorage's IndexDB Works"
summary: "IndexDB acts as vmstorage's memory - it remembers which numbers (TSIDs) belong to which metrics, making sure your queries get answered fast. This article walks through how this system works, from the way it organizes data to how it keeps track of millions of timeseries."
enableComments: true
toc: true
categories:
  - Open Source Tech
  - Monitoring
  - Time Series Database
tags:
  - vmstorage
  - indexdb
  - open source
  - database
  - monitoring
  - high-availability
  - time series
images:
  - /blog/vmstorage-how-it-handles-query-requests/vmstorage-how-indexdb-works-preview.webp
---

This piece is part of our ongoing VictoriaMetrics series, where we break down how different components of the system function:

1. [How VictoriaMetrics Agent (**vmagent**) Works](/blog/vmagent-how-it-works)
2. [How **vmstorage** Handles Data Ingestion](/blog/vmstorage-how-it-handles-data-ingestion)
3. [How **vmstorage** Processes Data: Retention, Merging, Deduplication,...](/blog/vmstorage-retention-merging-deduplication)
4. [When Metrics Meet **vminsert**: A Data-Delivery Story](/blog/vminsert-how-it-works)
5. How vmstorage's IndexDB Works? (We're here)
6. [How **vmstorage** Handles Query Requests From vmselect](/blog/vmstorage-how-it-handles-query-requests)
7. [Inside **vmselect**: The Query Processing Engine of VictoriaMetrics](/blog/vmselect-how-it-works)

> [!IMPORTANT]
> This discussion assumes you've checked out the earlier articles, which cover how vmstorage handles data ingestion and what goes on when it processes data.

![How vmstorage's IndexDB Works](/blog/vmstorage-how-it-handles-query-requests/vmstorage-how-indexdb-works-preview.webp)
<figcaption style="text-align: center; font-style: italic;">How vmstorage's IndexDB Works</figcaption>

vmstorage doesn't store the timeseries (e.g. `node_cpu_seconds_total{mode="idle"}`) directly into its **main storage**. Instead, it stores `TSIDs` (unique identifiers for each timeseries), along with the actual values and timestamps.  

![vmstorage's main storage](/blog/vmstorage-how-it-handles-query-requests/vmstorage-main-storage.webp)
<figcaption style="text-align: center; font-style: italic;">vmstorage's main storage</figcaption>

So, how does vmstorage know which `TSID` belongs to which metric? That's where **IndexDB** steps in — it's the database that keeps track of the relationship between timeseries and their `TSIDs`.

For example, when you make a query — whether you're using PromQL or the enhanced [MetricsQL](https://docs.victoriametrics.com/metricsql/), such as  `sum_over_time(node_cpu_seconds_total{mode="idle"}[5m])`, vmstorage focuses on two things: the metric you're asking for, `node_cpu_seconds_total{mode="idle"}`, and the time range. 

The aggregation function `sum_over_time` is actually handled by **vmselect** once it gets the samples from vmstorage.

![Metric query flow: vmselect and vmstorage interaction](/blog/vmstorage-how-it-handles-query-requests/vmstorage-vmselect-query-flow.webp)  
<figcaption style="text-align: center; font-style: italic;">Metric query flow: vmselect and vmstorage interaction</figcaption>  

So, what's the main job of IndexDB? 

Simply put, it translates your human-readable metric names, like `node_cpu_seconds_total{mode="idle"}`, into unique numeric IDs — these are the `TSIDs`.

> [!NOTE] Note
> `TSID` (Timeseries ID) is technically a wrapper for a metric ID, with a few extras like account ID or project ID. The metric ID itself is a big, unique number that identifies each timeseries. From the user's perspective, though, there's not much difference between `TSIDs` and metric IDs.

Once the `TSIDs` are sorted out, vmstorage uses them to dig into its main storage and pull out the actual metric data. Basically, this storage is just a collection of blocks organized by `TSID`. After retrieving those blocks, vmstorage sends them over to vmselect for further processing.

![Human-readable metrics converted into numeric IDs](/blog/vmstorage-how-it-handles-query-requests/indexdb-tsid-mapping.webp)  
<figcaption style="text-align: center; font-style: italic;">Human-readable metrics converted into numeric IDs</figcaption>  

Now, here's something we haven't touched on yet: how exactly does IndexDB deal with new metrics when they're added to vmstorage? And how does vmstorage handle requests from vmselect? That's what we'll cover in this discussion.

## How IndexDB is Structured  

vmstorage uses a three-stage IndexDB setup, which includes the next IndexDB, current IndexDB, and previous IndexDB. These indexes rotate periodically based on the retention period.  

- The **current IndexDB** holds the active index data. Any new timeseries data gets added here.  
- The **previous IndexDB** stores older data that's still within the retention period, so it's still available for queries.  
- The **next IndexDB** is essentially a placeholder, set up ahead of time to take over as the current IndexDB during the next rotation.  

![Three-stage IndexDB rotation for data retention](/blog/vmstorage-how-it-handles-query-requests/indexdb-retention-system.webp)  
<figcaption style="text-align: center; font-style: italic;">Three-stage IndexDB rotation for data retention</figcaption>  

The rotation happens once the retention period hits its deadline (with a small offset for buffer).  

Let's take an example. If you've got a 365-day retention period starting on 2023-01-01, the rotation would trigger on 2023-12-31T04:00:00Z. During this process, the next IndexDB takes over as the current IndexDB, the old current IndexDB becomes the previous IndexDB, and the old previous IndexDB is marked for deletion.  

If the retention period is updated, the system recalculates the rotation schedule based on the new setting.  

> [!WARNING] Warning: Be cautious when updating the retention policy.  
> There is a known issue [#7609: Changing -retentionPeriod may cause earlier deletion of previous indexDB](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7609). 

Each row in the IndexDB comes with a numeric prefix (e.g. `1`, `2`, `3`, etc.). This prefix determines what the mapping of that row is supposed to represent. Right now, there are 7 prefixes (or indexes) in play:  

1. **Tag to metric IDs (Global index)** – This maps a specific tag to metric IDs. For example, `1 method=GET 49,53` or `1 status=200 67,99,100,120,130`. It's one of the main mappings used to locate metric IDs.  
2. **Metric ID to TSID (Global index)** – After finding the metric IDs from the first mapping, this one links each metric ID to its TSID. Example: `2 49 TSID{metricID=49,...}`.  
3. **Metric ID to metric name (Global index)** – This connects a unique metric ID to the actual timeseries we've stored. Example: `3 http_request_total{method="GET",status="200"}`.  
4. **Deleted metric ID** – Not really a mapping, but more of a tracker for deleted metric IDs. Example: `4 152`. It's worth noting that deleting timeseries is resource-intensive, so it's best to minimize how often this happens.  
5. **Date to metric ID (Per-day index)** – This maps a specific date to a metric ID, helping quickly determine if the metric exists on that date. Example: `5 2024-01-01 152`, `5 2024-01-01 153`.  
6. **Date with tag to metric IDs (Per-day index)** – Similar to the first mapping, but scoped to a specific date for faster lookups. Example: `6 2024-01-01 method=GET 152,156,201`.  
7. **Date with metric name to TSID (Per-day index)** – This one looks up the TSID of a specific metric on a specific date. Example: `7 2024-01-01 http_request_total{method="GET",status="200"} TSID{metricID=49,...}`.  

You don't need to commit all of these to memory. As we keep going, they'll pop up when needed, and you'll see how they work in practice.  

### IndexDB on Disk

IndexDB follows a structure that's similar to the main storage, but the internal data it stores is a bit different. Instead of things like `values` and `timestamps`, it works with `items`, and `lengths`.  

Here's a quick look at how it's organized on disk: 

```bash
/path/to/indexdb/
├── 181B3F97816592AA/       # Previous indexdb (16 hex digits)
│   ├── parts.json          # List of all parts
│   ├── 181B3F9782297961/   # Part directory (random numbers)
│   │   ├── metadata.json   # Part's metadata
│   │   ├── items.bin 
│   │   ├── lens.bin
│   │   ├── index.bin 
│   │   └── metaindex.bin
│   ├── 181D447DFCF1147A/   # Another part
│   │   └── (same structure as above)
│   └── ...                 # More parts
│
├── 181B3F97816592AB/       # Current indexdb
│   └── (same structure as prev)
│
└── 181B3F97816592AC/       # Next indexdb
    └── (same structure as prev)
```

The `items.bin` file stores the entries we talked about earlier (e.g., `1 method=GET 49,53` or `2 status=200 67,99,100,120,130`). The `lens.bin` file helps locate the start and end of each item in `items.bin`.

![IndexDB structure: metaindex, index, items, lens](/blog/vmstorage-how-it-handles-query-requests/indexdb-file-structure.webp)
<figcaption style="text-align: center; font-style: italic;">IndexDB structure: metaindex, index, items, lens</figcaption>

The `index.bin` file, on the other hand, is where block headers live. These headers store metadata about each block, like the block's starting and ending positions in `items.bin` and `lens.bin`. They also include the block's shared prefix, which is the common prefix for all the items in that block:

![index.bin stores block prefix metadata](/blog/vmstorage-how-it-handles-query-requests/block-header-index.webp)
<figcaption style="text-align: center; font-style: italic;">index.bin stores block prefix metadata</figcaption>

During startup, vmstorage goes through a few key steps to get everything ready:  

- It starts by reading the `parts.json` file. This tells vmstorage which parts belong to which IndexDB. At this stage, it opens (but doesn't actually read) the heavier files like `index.bin`, `items.bin`, and `lens.bin`.  
- It then loads all the `metadata.json` files from every part into memory. These files include details like the number of items and blocks, as well as the first and last items in the part. This makes quick comparisons possible later.
- Finally, it loads all the `metaindex.bin` files from every part, whether big or small. These files are relatively lightweight but critical — they help vmstorage quickly find the right data inside the larger `index.bin` files.  

## How IndexDB Handles Data Ingestion  

When data arrives at vmstorage, the first thing it does is check whether the metric already has a unique ID (`TSID`). It does this by looking in the **`TSID` cache**. If it finds a match, great—things move quickly. But if there's a cache miss, it's considered a slow insert, and vmstorage falls back to consulting IndexDB instead.  

![Cache miss triggers IndexDB lookup](/blog/vmstorage-how-it-handles-query-requests/tsid-cache-indexdb.webp)  
<figcaption style="text-align: center; font-style: italic;">Cache miss triggers IndexDB lookup</figcaption>  

Here's how it works: IndexDB uses its seventh prefix mapping (`date + metric name -> TSID`) to look for the `TSID` in both the current and previous IndexDB (as long as the previous one is still relevant). If the `TSID` isn't found, vmstorage creates a brand-new `TSID` for the timeseries and adds it to all seven indexes.  

> [!TIP] Tip: Useful metrics
> - **Slow inserts** (cache misses): `vm_slow_row_inserts_total`  
> - **New timeseries** (cache misses that also weren't in IndexDB): `vm_new_timeseries_created_total`  

Most indexes get one new entry — except for the `tag to metric ID` index.

To explain this better, imagine we get a timeseries like `http_request_total{method="GET",status="200"}` on January 1st, 2024. In this case, IndexDB creates 3 entries: one for the metric name `http_request_total`, one for `method`, and one for `status` in both global and per-day `tag to metric IDs` indexes, right?

So, it's 6 entries in total. Let's see if our assumption is correct:

![Global and per-day indexes for metrics](/blog/vmstorage-how-it-handles-query-requests/global-perday-index.webp)  
<figcaption style="text-align: center; font-style: italic;">Global and per-day indexes for metrics</figcaption>  

What we just discussed isn't correct. The actual number of new rows added for the `tag to metric ID` index isn't 6 — it's 10. Why? Because of something called a **composite index**.  

A composite index combines the metric name with its tags. This narrows down the search space even further. It means you can search using just the tag (e.g., `{method="GET"}`) or combine it with the metric name (e.g., `http_request_total{method="GET",status="200"}`). It's designed to speed things up for most scenarios.

At the end of the ingestion process, two warmup steps kick in:  

- **One hour before the next rotation** (when the current and next IndexDB switch roles), the next IndexDB gets pre-filled with metrics being ingested during that time.  
- **One hour before the new day starts**, vmstorage pre-populates all the per-day indexes for the next day, including the last three indexes.  

> [!TIP] Tip: Useful metrics
> - How many rows were added to IndexDB: `vm_indexdb_items_added_total`  
> - How many bytes those rows took up: `vm_indexdb_items_added_size_bytes_total`  
> - How many metrics were pre-populated into the next IndexDB: `vm_timeseries_precreated_total`  
<!-- > How many items are sitting in the next IndexDB: `vm_cache_entries{type="storage/next_day_metric_ids"}`. (There're many other types of cache you can explore.) -->
<!-- > How many bytes are sitting in the next IndexDB: `vm_cache_size_bytes{type="storage/next_day_metric_ids"}`. -->

### The Merge Process  

IndexDB's structure has a lot in common with the main storage. Both include in-memory buffers, in-memory parts, small parts, and big parts.  

However, the entries we just talked about don't go straight to disk right away. Just like the main storage uses **raw-row shards** as a buffer, IndexDB has its own version of a buffer. Let's call it the `in-memory block shard`, since each entry here is a block (not a TSID row, like in the main storage's shard).  

So, how many shards are there? The number depends on your CPU cores, based on this formula:  

```go
shard_count = cpu_cores * min(16, cpu_cores)
```

In other words, the more CPU cores, the more shards. For example, if vmstorage has access to 4 CPU cores, it will create 16 shards. Each shard can hold up to 256 in-memory blocks. A block, in turn, is a set of index entries and can hold up to 64 KB of data.

When a shard is full, its blocks are flushed into an intermediate stage called "pending blocks":

![Shards hold in-memory blocks before flushing](/blog/vmstorage-how-it-handles-query-requests/shard-buffering-process.webp)
<figcaption style="text-align: center; font-style: italic;">Shards hold in-memory blocks before flushing</figcaption>

From there, the pending blocks are flushed into in-memory parts in two situations: either **periodically** every 1 second, or when **too many blocks** accumulate.

The flush process doesn't just group all the blocks into a part — it merges them. It works like this: the system sorts all the entries in the blocks, reads multiple blocks at once, picks the smallest item from these blocks, and adds it to a new sorted block. This way, items in the part are sorted by prefix.

In addition to everything else, the `tag to metric IDs` mapping gets even more compact thanks to a combining metric ID technique. 

Initially, each row in the `tag to metric IDs` mapping contains a single metric ID. But after the merge process, these rows are transformed into sets of metric IDs:

![Combining metric IDs for compact storage](/blog/vmstorage-how-it-handles-query-requests/prefix-id-merging.webp)  
<figcaption style="text-align: center; font-style: italic;">Combining metric IDs for compact storage</figcaption>  

Now, here's the big question:  

> _"Why is the per-day index faster than the global index?"_  

With per-day indexing, searches are naturally split by date. When looking for a metric within a specific time range, the system only searches through entries that match the date prefix. This dramatically reduces the search space since it only focuses on metrics that were active on the relevant days.  

For example, if you're searching for `http_request_total{status="200"}` between 13:00 and 14:00 on 2024-01-01:

- The global index would need to scan every `1 status=200` entry across the entire retention period. Most of those entries wouldn't even apply to the one-hour window you're looking at. 
- With the per-day index, IndexDB starts with the date prefix `6 2024-01-01 status=200` and quickly narrows down the search to entries from January 1st, 2024. It doesn't waste time on data from other days, like December or later dates.  

The downside is that per-day indexes use significantly more disk and memory compared to the global index.  

If you have a timeseries that appears daily, `http_request_total{status="200"}`, the per-day `tag to metric IDs` index will generate 21 entries over 7 days (3 entries per day × 7 days). Meanwhile, the global index will only need 3 entries for the same timeseries, no matter how many days it spans.

_Finally, when it comes to merging small parts, big parts, and flushing data to disk, the process works similarly to how it's done [in the main storage](/blog/vmstorage-retention-merging-deduplication). So, we'll skip that here to avoid repeating ourselves._

> [!NOTE] Read next: [How **vmstorage** Handles Query Requests From vmselect](/blog/vmstorage-how-it-handles-query-requests)

## Who We Are

Need to monitor your services to see how everything performs and to troubleshoot issues? [VictoriaMetrics](https://docs.victoriametrics.com/) is a fast, **open-source**, and cost-efficient way to stay on top of your infrastructure's performance.

If you come across anything outdated or have questions, feel free to send me a DM on [X (@func25)](https://twitter.com/func25).