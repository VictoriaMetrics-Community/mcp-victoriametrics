---
draft: false    
page: blog blog_post
authors:
  - Roman Khavronenko
  - Aliaksandr Valialkin 
date: 2023-11-17
enableComments: true
title: "Performance optimization techniques in time series databases: function caching"
summary: "This blog post is a second in the series of the blog posts based on the
talk about 'Performance optimizations in Go', GopherCon 2023. It is dedicated to various optimization techniques used in VictoriaMetrics for
improving performance and resource usage."
categories:
  - Performance
  - Time Series Database
tags:
  - performance
  - go
images:
  - /blog/tsdb-performance-techniques-function-caching/preview.webp
---

This blog post is also available as a [recorded talk](https://www.youtube.com/watch?v=NdjuW98ep_w&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj)
with [slides](https://docs.google.com/presentation/d/1hquMVEwuefqCefPI-A1YulitQvCaib1MVXUao9DzHxQ/edit).

**Table of Contents**

Performance optimization techniques in time series databases:
- [Strings interning](https://victoriametrics.com/blog/tsdb-performance-techniques-strings-interning/);
- Function caching (you are here);
- [Limiting concurrency for CPU-bound load](https://victoriametrics.com/blog/tsdb-performance-techniques-limiting-concurrency/);
- [sync.Pool for CPU-bound operations](https://victoriametrics.com/blog/tsdb-performance-techniques-sync-pool/).

---

[Relabeling](https://docs.victoriametrics.com/relabeling.html) is an important feature that allows users to modify 
metadata (labels) of scraped metrics before they ever make it to the database.

As an example, some of your scrape targets may generate metric labels with underscores (`_`),
and some of your targets may generate labels with hyphens (`-`). Relabeling allows you to make this consistent, 
making database queries easier to write:'

{{< image href="/blog/tsdb-performance-techniques-function-caching/relabeling-example.webp" alt="An example of relabeling rule to replace hyphens with underscores. You can play with VictoriaMetrics' relabeling functionality <a href='https://play.victoriametrics.com/select/accounting/1/6a716b0f-38bc-4856-90ce-448fd713e3fe/prometheus/graph/#/relabeling?config=-+action%3A+labelmap_all%0A++regex%3A+%22-%22%0A++replacement%3A+%22_%22&labels=%7B__name__%3D%22metric%22%2C+foo-bar-baz%3D%22qux%22%7D' target='_blank'>in our playground</a>." >}}

Relabeling, if defined, happens every time `vmagent` scrapes metrics from your targets, but as we've seen before,
`vmagent` is likely to see the same metric label many times. That means if we once saw `foo-bar-baz` and changed
it to `foo_bar_baz`, then it is very likely we'll have to do the same transformation on the next scrape as well.
In this case, caching the results of the relabeling function is likely to reduce CPU usage.

Internally, we implement caching for relabeling functions via struct called `Transformer`:
```go
type Transformer struct {
    m sync.Map
    transformFunc func(s string) string
}
```

`Transformer` contains a [sync.Map](https://pkg.go.dev/sync#Map) for thread-safe access to cached results,
and a function `transformFunc` that will do the actual relabeling.

`Transformer` implements function `Transform` which we use during relabeling:
```go
func (t *Transformer) Transform(s string) string {
    v, ok := t.m.Load(s)
    if ok {
         // Fast path - the transformed `s` is found in the cache.
         return v.(string)
    }
    // Slow path - transform `s` and store it in the cache.
    sTransformed := t.transformFunc(s)
    t.m.Store(s, sTransformed)
    return sTransformed
}
```

The `Transform` function first checks the cache using the `Load` function. If a cached result is found, 
then it returns the result from the cache. Otherwise, it will call `transformFunc` to do the transformation,
store the result in the cache, and return it.

As an example, here's a `Transformer` that replaces any character [not allowed in Prometheus data model](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels)
with an underscore:
```go
// SanitizeName replaces unsupported by Prometheus chars
// in metric names and label names with _.
func SanitizeName(name string) string {
    return promSanitizer.Transform(name)
}

var promSanitizer = NewTransformer(func(s string) string {
    return unsupportedPromChars.ReplaceAllString(s, "_")
})

var unsupportedPromChars = regexp.MustCompile(`[^a-zA-Z0-9_:]`)
```

In the above example, `promSanitizer` is created using our `Transformer` constructor. This constructor creates 
a new `sync.Map`, and stores the reference to the passed function. Now we can use `SanitizeName` function 
in the code "hot path" to sanitize scraped label names.

Function result caching allows you to trade off reduced CPU time for increased memory usage in certain cases. 
It works best when caching CPU-heavy functions that take a limited amount of possible values. 
Examples of CPU-heavy functions include those that do string transforms or regex matching.

### Summary

VictoriaMetrics uses function result caching for its relabeling feature, but doesn't use it for caching database queries.
In the case of database queries, the range of possible values is too large, and it's likely our cache hit rate would be low.
As with [strings interning](https://victoriametrics.com/blog/tsdb-performance-techniques-strings-interning/), 
functions results caching works the best if **number of cached variants is limited**, so you can achieve **high
cache hit rate**.

Stay tuned for the new blog post in this series!
