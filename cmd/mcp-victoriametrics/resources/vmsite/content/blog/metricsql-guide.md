---
draft: true
page: blog blog_post
authors:
  - Phuong Le
date: 2025-03-04
title: "Inside vmselect: The Query Processing Engine of VictoriaMetrics"
summary: "The article explains vmselect's core functionalities: concurrent request processing, query parsing and validation, data fetching and merging from vmstorage nodes, and memory-efficient result caching."
enableComments: true
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

Implicit conversion is a feature in VictoriaMetrics' MetricsQL that automatically transforms queries to make them valid without requiring users to explicitly specify certain parameters. While this feature aims to make queries more user-friendly, it can sometimes lead to confusion or unexpected behavior.

## Types of Implicit Conversions

MetricsQL performs several types of implicit conversions:

### Automatic Lookbehind Window Addition

When a window (time range in square brackets) is missing inside a rollup function, MetricsQL automatically adds one: For most rollup functions (except `default_rollup` and `rate`), it uses the step value from the query (known as `$__interval` in Grafana or `1i` in MetricsQL).

Example: `avg_over_time(temperature)` is automatically converted to `avg_over_time(temperature[1i])`

For `default_rollup` and `rate` functions, it uses max(step, scrape_interval), where scrape_interval is the interval between raw samples. This prevents gaps in graphs when the step is smaller than the scrape interval.

### Automatic Wrapping of Series Selectors

Series selectors that aren't wrapped in rollup functions are automatically wrapped in the `default_rollup` function:
- `foo` becomes `default_rollup(foo)`
- `foo + bar` becomes `default_rollup(foo) + default_rollup(bar)`
- `count(up)` becomes `count(default_rollup(up))` because `count` is an aggregate function, not a rollup function
- `abs(temperature)` becomes `abs(default_rollup(temperature))` because `abs` is a transform function, not a rollup function

### Automatic Step Addition in Subqueries

If the step parameter is missing in a subquery, a `1i` step is automatically added:
Example: `avg_over_time(rate(http_requests_total[5m])[1h])` becomes `avg_over_time(rate(http_requests_total[5m])[1h:1i])`

### Automatic Subquery Formation

If a non-series selector is passed to a rollup function, a subquery with 1i lookbehind window and 1i step is automatically formed:
Example: `rate(sum(up))` becomes `rate((sum(default_rollup(up)))[1i:1i])`

## Detection of Implicit Conversions

VictoriaMetrics uses the `IsLikelyInvalid` function in the MetricsQL package to detect queries that rely on implicit conversions. This function checks for:

1. Rollup functions that receive arguments other than simple metric expressions
2. Rollup expressions without an explicit window parameter

## Control Mechanisms

Starting from version v1.101.0 (released on April 26, 2024), VictoriaMetrics introduced two command-line flags to help manage implicit conversions:
- `-search.disableImplicitConversion`: When enabled, VictoriaMetrics returns an error for queries that rely on implicit conversions, forcing users to write explicit queries.
- `-search.logImplicitConversion`: When enabled, VictoriaMetrics logs queries that use implicit conversions without blocking them, helping users identify and potentially refactor problematic queries.
In version v1.102.0-rc2, additional validation was added to check for:
- Ranged vector arguments in non-rollup expressions (e.g., `sum(up[5m])`)
- Missing ranged vector arguments in rollup expressions (e.g., `rate(metric)`)

## Why Control Implicit Conversions?

The ability to disable or log implicit conversions was added in response to community feedback. According to GitHub issue #4338, implicit conversions often caused more confusion than convenience, leading to various support issues and discussions in community channels.

The main problems with implicit conversions include:
- Unexpected Behavior: Users might not realize their queries are being transformed, leading to confusion when results don't match expectations.
- Performance Impact: Some implicit conversions, especially those creating subqueries, can significantly impact query performance.
- Debugging Difficulty: When queries are implicitly transformed, it can be harder to debug issues since the executed query differs from what was written.
- Inconsistency with PromQL: While MetricsQL extends PromQL, these implicit conversions create differences in behavior that can confuse users familiar with Prometheus.

## Best Practices
To avoid issues with implicit conversions:
- Write Explicit Queries: Always specify lookbehind windows and steps explicitly in your queries.
- Use the Logging Flag: Enable `-search.logImplicitConversion` to identify queries that rely on implicit conversions.
- Gradually Refactor: If you have many queries that rely on implicit conversions, use the logging flag to identify them and refactor them gradually.
- Consider Strict Mode: For production environments where query correctness is critical, consider enabling `-search.disableImplicitConversion` to enforce explicit query writing.