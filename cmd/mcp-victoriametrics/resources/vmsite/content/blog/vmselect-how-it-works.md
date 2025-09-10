---
draft: false
page: blog blog_post
authors:
  - Phuong Le
date: 2025-03-07
title: "Inside vmselect: The Query Processing Engine of VictoriaMetrics"
summary: "The article explains vmselect's core functionalities: concurrent request processing, query parsing and validation, data fetching and merging from vmstorage nodes, and memory-efficient result caching."
enableComments: true
categories:
  - Open Source Tech
  - Monitoring
  - Time Series Database
toc: true
tags:
  - vmstorage
  - indexdb
  - open source
  - database
  - monitoring
  - high-availability
  - time series
images:
  - /blog/vmselect-how-it-works/vmselect-how-it-works-preview.webp
---

This piece is part of our ongoing VictoriaMetrics series, where we break down how different components of the system function:

1. [How VictoriaMetrics Agent (**vmagent**) Works](/blog/vmagent-how-it-works)
2. [How **vmstorage** Handles Data Ingestion](/blog/vmstorage-how-it-handles-data-ingestion)
3. [How **vmstorage** Processes Data: Retention, Merging, Deduplication,...](/blog/vmstorage-retention-merging-deduplication)
4. [When Metrics Meet **vminsert**: A Data-Delivery Story](/blog/vminsert-how-it-works)
5. [How vmstorage's IndexDB Works?](/blog/vmstorage-how-indexdb-works)
6. [How **vmstorage** Handles Query Requests From vmselect](/blog/vmstorage-how-it-handles-query-requests)
7. Inside vmselect: The Query Processing Engine of VictoriaMetrics

![Inside vmselect: The Query Processing Engine of VictoriaMetrics](/blog/vmselect-how-it-works/vmselect-how-it-works-preview.webp)
<figcaption style="text-align: center; font-style: italic;">Inside vmselect: The Query Processing Engine of VictoriaMetrics</figcaption>

## Receive Requests

vmselect handles multiple requests from different APIs at the same time. But there's a limit—twice the number of CPU cores, capped at 8 (unless you configure it differently using `-search.maxConcurrentRequests`). 

The reason for this 8-core limit is that running too many concurrent requests is not beneficial, as each request typically uses all CPU cores during query execution. So, if you're running on more than 4 cores, vmselect allows up to 8 concurrent requests.

When requests exceed this limit, they wait in a queue for up to 10 seconds (`-search.maxQueueDuration`). If they're still waiting after that, they get dropped.  

> [!TIP] Metrics: Is vmselect overwhelmed? 
> - How many requests are actively being processed: `vm_concurrent_select_current`  
> - The fixed number of concurrent requests allowed: `vm_concurrent_select_capacity`   
> - How often vmselect hits its request limit: `vm_concurrent_select_limit_reached_total`  
> - How many requests time out in the queue: `vm_concurrent_select_limit_timeout_total`  

vmselect processes range query requests through these HTTP paths:  
- In a cluster setup: `select/{tenantID}/prometheus/api/v1/query_range`  
- In a single-node setup: `/api/v1/query_range`  

A few request settings are worth noting:  

- **Deadlines**: By default, a request has 10 seconds to complete (`-search.maxQueryDuration`). You can override this with the `timeout` query parameter.  
- **Caching**: Query results may be cached in memory. To bypass caching, use the `nocache` parameter to disable caching for individual requests or the `-disableCache` flag to disable caching entirely.  
- **Lookback delta**: is controlled by `-search.maxLookback` or can be set per request using the `max_lookback` parameter. If neither is set, vmselect automatically determines a value based on the data.  
- **Step**: If the `step` parameter isn't provided, vmselect defaults to 5 minutes. Typically, query editors like Grafana or [vmui](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/app/vmui) handle this for you.  

Not sure what _lookback delta_ or _step_ means? Check out [Instant Queries and Range Queries Explained](/blog/prometheus-monitoring-instant-range-query).  

## Parse Query  

vmselect takes your query as a raw string, like `rate(http_requests_total{method="GET"}[10m])`. Before it can do anything, it has to break that down into structured data for analysis:

![Understanding query structure: function, filters, and window](/blog/vmselect-how-it-works/query-parsing-breakdown.webp)
<figcaption style="text-align: center; font-style: italic;">Understanding query structure: function, filters, and window</figcaption>

To avoid the overhead of parsing the same query over and over, vmselect caches structured query results in memory.  

<!-- This cache holds around 10,000 entries, split across 128 buckets. Each bucket stores 78 entries. When a bucket fills up, it evicts 10% of its entries. -->  

> [!TIP] Is the parsing cache working well?  
> - How many requests hit the cache: `vm_cache_requests_total{type="promql/parse"}`  
> - How many cache misses: `vm_cache_misses_total{type="promql/parse"}`  
> - How many entries are currently in the cache: `vm_cache_entries{type="promql/parse"}`  

If a query is likely invalid, vmselect automatically tweaks it to make it valid—no need to manually adjust parameters. For example, `avg_over_time(temperature)` wouldn't work because `avg_over_time` needs a range vector. Instead of rejecting it outright, vmselect rewrites it as `avg_over_time(temperature[1i])`. In this case, `1i` equals the step parameter provided in the API request. 

This implicit conversion makes queries easier to write, but it can also hide mistakes. If you prefer strict validation, set the `-search.disableImplicitConversion` flag to force vmselect to reject invalid queries. If you like the convenience but still want visibility, enable `-search.logImplicitConversion` to get warnings when adjustments happen.  

At this point, vmselect has a basic understanding of the query and moves on to fetch data from vmstorage.  

## Fetch Data

vmselect first figures out what kind of expressions are in your query. It could be a rollup function like `sum_over_time` or `rate`, a transformation function like `abs` or `sqrt`, or a binary operation like `+`, `-`, `*`, `/`. Each type needs to be handled differently. No need to go into all the details here—there's a separate article for that later.  

<!-- TODO(@phuong,article): link to the article -->

Some functions need historical data to work, like `increase` and `rate`. vmselect automatically pulls in older data, even from before your query's start time. For example, if you query data starting at 10:00, it might fetch data from 09:00 to 10:00. But how much historical data is needed? 

The exact amount depends on your query and the settings:  

```go
period = max(step, window) + silence_interval
```

So the actual `start` of the query is shifted by the period:

```go
start = requested_start - period
```

Let's take an example: 

- The silence interval is 5 minutes (`-search.minStalenessInterval`)
- You query `rate(http_requests_total[10m])` at 10:00, so the window is 10 minutes drived from the query (`[10m]`).
- Assume the `step` is very small compare to the window.

The actual start is:

```go
period = max(step, window) + silence_interval
       = max(step, 10m) + 5m
       = 15m

start = requested_start - period
      = 10:00 - 15m
      = 09:45
```

So, vmselect will actually fetch data from 09:45 instead of your original query 10:00. If your metrics are collected at an unusual interval—say, every 15 minutes instead of the usual scrape frequency—you might need to increase silence interval value to make sure vmselect gathers enough data.

_In the last section of this discussion, you will know that the most recent data points (up to 30 seconds) are not included in the results._

Once the time range is set, vmselect builds a search query with tag filters and sends it to all vmstorage nodes:

![vmselect extracts filters and determines time range](/blog/vmselect-how-it-works/vmselect-query-distribution.webp)
<figcaption style="text-align: center; font-style: italic;">vmselect extracts filters and determines time range</figcaption>

Each vmstorage node streams back matching data in blocks. These blocks are first stored in an in-memory buffer for each vmstorage node:

![vmselect stores incoming data in per-node buffers](/blog/vmselect-how-it-works/vmselect-inmemory-buffers.webp)
<figcaption style="text-align: center; font-style: italic;">vmselect stores incoming data in per-node buffers</figcaption>

The buffer size depends on your system, ranging from 64 KB to 4 MB. It's calculated as:

```go
bufferSize = clamp(system memory / 1024, 64 KB, 4 MB)
```

So, if your system has 16 GB of memory, the buffer size is intended to be 16 MB but will be capped at 4 MB.

_Note that, these buffers are not shared across queries. Instead, each query will have its own set of in-memory buffers._

To check how much memory is allocated for these buffers, use `vm_tmp_blocks_max_inmemory_file_size_bytes`. This shows the maximum size a temporary file can reach before spilling to disk. The actual size of a buffer after use is recorded in `vm_tmp_blocks_inmemory_file_size_bytes`. So, what these numbers tell you:

- If the actual sizes are much smaller than the maximum, the buffer might be too large.  
- If they often hit the maximum, queries might be spilling to disk too frequently.

You can actually tune the buffer size using `-search.inmemoryBufSizeBytes`.

That said, any extra data that doesn't fit in memory is written to temporary files on disk (`-cacheDataPath/tmp/searchResults`). When query load is high, you might notice increased disk I/O on vmselect nodes.  

> [!TIP] Metrics: Is your disk handling the load?
> - How many temporary files have been created: `vm_tmp_blocks_files_created_total`.
> - Free space available in the temporary directory: `vm_tmp_blocks_files_directory_free_bytes`.

If some vmstorage nodes fail and can't respond with their data, vmselect will still process the query with the available data by default. There are two scenarios here:

- For recent data, the available data is considered complete because when one vmstorage node fails, the other nodes receive the data for that failed node.
- For older data, the available data is partial because the failed node may have retained part of the data.

In the second case, this can be misleading since the results may be incomplete. To make errors more obvious, set the `-search.denyPartialResponse` flag. For better fault tolerance and data availability, you can organize vmstorage nodes into groups with a replication factor. More details on that here: [vmstorage groups at vmselect](https://docs.victoriametrics.com/cluster-victoriametrics/#vmstorage-groups-at-vmselect).

> [!TIP] Metrics: Are vmstorage nodes working properly?
> - How many search requests were sent to each node: `vm_requests_total{action="search", type="rpcClient", name="vmselect"}`
> - How many search requests failed: `vm_request_errors_total{action="search", type="rpcClient", name="vmselect"}`

## Evaluating & Memory Usage  

As blocks arrive from multiple vmstorage nodes, vmselect merges them while keeping the data in order. If replication is enabled on vminsert, some data might be duplicated:  

![vmselect merges and deduplicates replicated data blocks](/blog/vmselect-how-it-works/vmselect-deduplication.webp)
<figcaption style="text-align: center; font-style: italic;">vmselect merges and deduplicates replicated data blocks</figcaption>

To handle this, vmselect should be set up with deduplication using a small interval, like 1 millisecond (`-dedup.minScrapeInterval=1ms`). When multiple data points have the same timestamp (or are within 1 millisecond), vmselect keeps only the latest one. Without deduplication, functions like `sum_over_time`, `count_over_time`, and `quantile_over_time` can produce incorrect results.  

vmselect uses up to 60% of system memory (`-memory.allowedPercent`) for its activities. It will try to stay under this limit, but it doesn't guarantee that it will use the full amount.

Before executing a query, vmselect estimates its memory usage and applies several safeguards. All running queries combined can use up to 50% of the allowed memory. If a query would exceed this limit, vmselect rejects it and returns an error.  

> [!IMPORTANT] Quiz: How much memory is allocated for running queries in your system?
> Suppose your vmselect memory limit is 60 GB:
> - The allowed memory is 60 * 60% = 36 GB.
> - All running queries combined can use up to 50% of the allowed memory, which is 36 * 50% = 18 GB.

To monitor memory usage more closely, set `-search.logQueryMemoryUsage` with a threshold. If a query crosses this limit, vmselect logs a warning and increments the `vm_memory_intensive_queries_total` metric. For stricter control, limit memory per query using `-search.maxMemoryPerQuery`.  

Once the data is ready, vmselect processes it based on your query. If the query includes aggregation functions like `sum`, `avg`, `min`, or `max`, vmselect groups the data and runs calculations across multiple time series. This can be memory-intensive, especially for high-cardinality metrics.  

To reduce memory usage, vmselect applies _incremental aggregation_ where possible. Instead of loading everything at once, it processes time series in batches, updating results step by step.  

> [!TIP] Metrics: Is aggregation running efficiently?  
> - Total incremental aggregation calls: `vm_incremental_aggr_function_calls_total{}`  
> - Number of series processed per incremental aggregation: `vm_incremental_aggr_function_series_processed_total{}`  

For more complex functions like `quantile` or `histogram_quantile`, incremental processing isn't an option. These functions need all data points at once to calculate an accurate result, which means higher memory usage.  

Once the query is evaluated, vmselect caches the results if caching isn't disabled. This **rollup result cache** uses up to 12.5% of allowed memory.

The cache avoids storing very recent data; anything from `now-5m` to `now` is left out (`-search.cacheTimestampOffset`) because this 5-minute window of data is considered incomplete. Some of the samples are delayed and haven't been ingested into VictoriaMetrics yet.

This is an in-memory cache, but it can be made persistent by setting `-cacheDataPath`. With this option, cached results are saved to disk at `-cacheDataPath/rollupResult` when vmselect shuts down and reloaded on startup if available.

In rare cases, cached results become outdated due to major changes in underlying data. If needed, the cache can be reset with the `-search.resetRollupResultCacheOnStartup` flag or through the `/internal/resetRollupResultCache` API.  

> [!TIP] Metrics: Is rollup caching improving performance?  
> - Requests fully served from cache: `vm_rollup_result_cache_full_hits_total`  
> - Requests partially served from cache (with the remaining data fetched and processed in real-time): `vm_rollup_result_cache_partial_hits_total`  
> - Requests that had to be processed from scratch (cache miss): `vm_rollup_result_cache_miss_total`

Last but not least, the data returned by vmselect has a delay of 30 seconds (`-search.latencyOffset`). 

This means that the most recent data points will not appear in your chart immediately. The reason for this delay is to account for 'incomplete data.' The 30-second window represents the period during which `vmselect` considers data points to be [potentially incomplete](https://docs.victoriametrics.com/keyconcepts/#query-latency) and therefore may be misleading.

## Who We Are

Need to monitor your services to see how everything performs and to troubleshoot issues? [VictoriaMetrics](https://docs.victoriametrics.com/) is a fast, **open-source**, and cost-efficient way to stay on top of your infrastructure's performance.

If you come across anything outdated or have questions, feel free to send me a DM on [X (@func25)](https://twitter.com/func25).