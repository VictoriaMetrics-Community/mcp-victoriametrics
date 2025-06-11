---
draft: false
page: blog blog_post
authors:
 - Aliaksandr Valialkin
date: 2024-09-11
title: "The Rise of Open Source Time Series Databases"
enableComments: true
summary: "Time series databases are essential tools in any software engineer’s toolbelt. Their development has been shaped by user needs and countless open source contributors, leading to the healthy ecosystem of options we see today. In this article, you’ll see how time series databases came about, and why so many are open source."
categories:
 - Company News
 - Observability
 - Monitoring
 - Time Series Database
 - Open Source Tech
tags:
 - victoriametrics
 - monitoring
 - observability
 - time series
 - timeseries
 - database
 - open source
images:
 - /blog/the-rise-of-open-source-time-series-databases/preview.webp
---

Time series databases allow you to store and query metrics efficiently. For example, if you want to forecast load on your servers, or identify intermittent faults with your production services, time series databases can help. Besides infrastructure monitoring, time series databases have been invaluable in finance, IoT applications, manufacturing, and more.

Many time series databases, including VictoriaMetrics, are open source. In this article, you’ll see how time series databases came about, and why so many are open source. We’ll also share our insider take on the future of this space.

## What is a time series database?

When people think of databases, they think of relational databases like PostgreSQL and Oracle. Relational databases store data in a table format, where each row contains an instance of the data being described, with each column describing a different aspect of the data. For example, you may have an `employees` table with `name`, `role`, and `salary` columns. Each row would represent a different employee:

| Name | Role | Salary |
| :-------- | :------- | :-------: |
| Alice | Charlie | $100,000 |
| Bob  | Moral support | $50,000 |
| Charlie | Dog | $0 |

This model works great for many types of data. As long as the amount of data doesn’t get too large, it supports fast querying for specific information and easy updates if the value of a cell changes. However, when data size reaches billions or trillions of rows (typical sizes for time series data), performance degrades.

A time series database is a database designed specifically for time series data. Time series data is a sequence of measurements, or observations that have an associated timestamp. In a database like VictoriaMetrics, these measurements can also have associated metadata known as labels.

Time series databases can ingest and query vast amounts of data, and use clever tricks to use fewer resources than a relational database would when storing the same data. Nothing comes for free, of course, and time series databases aren’t suitable for all data types. For example, time series databases:

- Take a long time to update data that has already been stored
- Can only store numbers, not other types of data

These trade-offs work well for typical time series data, e.g. stock prices, or the number of requests made to a specific endpoint. However, they wouldn’t work for our employee example earlier.

## A short history of time series databases

To explain time series databases, we’ll go back to the early 2010s before today’s popular time series databases existed. Back then, infrastructure monitoring would often be real-time-only. If you were lucky, you might have a dashboard showing you the current state of the system you were working on, but historical data was hard to come by. Many problems, such as intermittent faults, could go unnoticed for a long time.

Tracking the state of your system across time requires storing billions, or even trillions of rows. In the early 2010’s the majority of databases were traditional relational databases, which weren’t designed to handle this volume of data.

### The first time-series database

[InfluxDB](https://en.wikipedia.org/wiki/InfluxDB) launched in 2013 and was the first mainstream time series database. There were a couple of efforts beforehand, but they were either unknown or intended for niche use cases. Importantly for our story, InfluxDB was, and still is an open source project. Being open source did wonders for the popularity of both InfluxDB and time series databases as a whole.

As with any piece of software, however, InfluxDB wasn’t perfect. One of the main sources of user complaints was the database’s stability. A search for “[influxdb crash](https://www.google.com/search?q=influxdb+crash)” still brings up results full of confused and frustrated users. Worse, every major version release has broken backward compatibility and automatic update tools have been late, missing, or buggy. Version upgrade processes have been time-consuming and frustrating for users.

### Prometheus

Around the same time InfluxDB was started, engineers at SoundCloud were [developing their own](https://thenewstack.io/prometheus-at-10-whats-been-its-impact-on-observability/) in-house infrastructure monitoring system called Prometheus. While InfluxDB aims to cater to all time series use cases, Prometheus is focused explicitly on metrics and event monitoring. For example, Prometheus only supports recording metrics as floating point numbers, while InfluxDB supports many data types.

With a more narrow focus, Prometheus developed greater stability and efficiency than InfluxDB. It is known for being easier to run, configure, upgrade, and troubleshoot than InfluxDB. Prometheus’ excellent developer experience has made it the de facto observability solution today.

## Enter VictoriaMetrics

While Prometheus is a great product, it’s still possible to reach its performance limits when working with large quantities of data. Before founding VictoriaMetrics, I was working with another of our co-founders, Roman, at an ad-tech company. The systems we worked on served millions of requests per second, providing an observability and analytics challenge.

Our team were early adopters of Prometheus for observability. It revolutionized our workflow and helped us see more than just the current state of their system, allowing us to see and analyze historical data. Historical observability data is critical for building reliable, performant systems. Prometheus was a big help for us, and we stored so much data in it that we reached its limits.

At the same time as adopting Prometheus, we migrated from PostgreSQL to ClickHouse for our analytics workloads. ClickHouse’s architecture was very efficient, allowing us to downscale from 12 servers to just one to run their analytics. This got us thinking, “_What if we had Prometheus, but with ClickHouse’s architecture?_”.

This question was the birth of VictoriaMetrics, a new time series database that carried on Prometheus’ legacy while adopting some of the designs that made ClickHouse so efficient. For example, VictoriaMetrics:

- Uses advanced compression techniques that use less disk space and less memory
- Stores and processes data in blocks for speed and efficiency
- Makes use of all available CPU cores to maximize performance

VictoriaMetrics’ heritage has given it outstanding efficiency, simplicity, and performance. Our [benchmarks](https://victoriametrics.com/blog/reducing-costs-p1/) show VictoriaMetrics using 2.5x less disk space and servicing queries 16x faster than Prometheus.

## Landscape

InfluxDB, Prometheus, and VictoriaMetrics are just three of the time series databases available today. VictoriaMetrics arrived during a time series database boom. The late 2010s saw many open source time series databases launch, including VictoriaMetrics, TimescaleDB, QuestDB, and more.

Many other categories of software are dominated by proprietary options. So you may be wondering where are all the proprietary time series databases? The answer is that while there are successful proprietary time series databases, they compete in niche use cases. For example, [kdb+](https://en.wikipedia.org/wiki/Kdb%2B) is very popular with high-frequency trading firms.

## The future of time series databases

It’s not easy to predict the future, but looking at the pain points of today can provide some hints on what problems will be solved next. Time series databases today have two well-known pain points—the high cardinality problem, and high time series churn rates.

### High cardinality

The performance of a time series database is usually directly correlated with the number of active time series. For a database like VictoriaMetrics or Prometheus, the number of active time series is a product of the number of unique label combinations. For example, consider the following metric:

```PromQL
requests_total{path="/", code="200", machine="vm01"}
```

The above metric counts the total number of requests served for the root path, that returned HTTP 200, on the machine called `vm01`. Each of `path`, `code`, and `machine` are a **label** — a key-value pair that contains metadata about a metric. This metric would be associated with a single time series.

Now, imagine that your application serves 1000 paths, might return 20 different HTTP codes from each path, and is served from 100 machines. This means that your monitoring infrastructure would need to deal with 1000 x 20 x 100 = 2,000,000 time series just to count the total number of requests served.

The number of active time series is referred to as cardinality. [High cardinality](https://victoriametrics.com/categories/high-cardinality/), or many active time series, leads to database slowdowns and failures.

VictoriaMetrics helps you deal with high cardinality through raw performance — it supports [over 10 million](https://valyala.medium.com/insert-benchmarks-with-inch-influxdb-vs-victoriametrics-e31a41ae2893) active time series on a single machine. That said, it is still easy to accidentally significantly increase cardinality with a simple label change, so this is not a solved problem.

## Time series churn

Tracking infrastructure metrics generally requires recording which machine generated the metric. With a database like VictoriaMetrics, this is done by recording the machine name with a label. Machine names were relatively static when physical machines ran applications, but containers and Kubernetes have changed the field significantly.

The parallel for a physical machine in Kubernetes is the pod. A pod is a running instance of a container that is uniquely identifiable. Due to their lightweight nature, pods can be created and destroyed much more easily than physical machines can.

In Kubernetes, a pod restart generates a new name. A new pod name translates into a new label value, invalidating the old time series and creating a new one. Over a year, a typical Kubernetes installation can expect to generate and invalidate **tens of billions** of time series this way. Huge numbers of time series can make it hard to query historical data and slow down your database.

This problem isn’t as visible as high cardinality. The typical use case for a time series database involves querying hours or days of historical data, reducing the number of time series involved. As users ask more of their time series databases and perform deeper analysis over longer periods, we expect this problem will become more prevalent.

## Wide events

There has been a lot of buzz lately about the concept of wide events. A wide event is a bit like a structured log entry. It records something happening at a particular time, along with associated attributes in the form of key-value pairs. They are called “wide events”, instead of just “events” because they are supposed to capture all the context around the event, leading to many attributes per-event.

Wide events are an exciting development for observability, as they can help to replace unstructured logs and traces that are hard to query. We don’t believe wide events will replace metrics as the storage space required for a wide event is orders of magnitude greater than for a metric. A single metric can be stored in VictoriaMetrics in **less than a single byte** thanks to clever compression.

If you’re interested in using wide events to help enhance your logging, it’s worth checking out VictoriaLogs. VictoriaLogs allows you to store arbitrary key-value pairs for each log entry, which you can then filter on and group by at query time. [VictoriaLogs](https://victoriametrics.com/products/victorialogs/) also supports unstructured logs, allowing you to transition at your own pace.

## Conclusion

Time series databases are essential tools in any software engineer’s toolbelt. Their development has been shaped by user needs and countless open source contributors, leading to the healthy ecosystem of options we see today. 

If you’re curious about time series databases, why not give VictoriaMetrics a go? VictoriaMetrics is incredibly efficient, simple to deploy, and completely open source. Check out our [quick start guide](https://docs.victoriametrics.com/quick-start/) or [request a demo](https://victoriametrics.com/contact-us/) with our team.
