---
draft: false
page: blog blog_post
authors:
 - Roman Khavronenko
date: 2023-01-16
title: "Monitoring benchmark: how to generate 100 million samples/s of production-like data"
summary: "One of the latest benchmarks we made was 'VictoriaMetrics: scaling to 100 million metrics per second'. 
While the fact of such scale for VictoriaMetrics is noteworthy on its own, the benchmark tool used to generate that 
load is usually overlooked. In this blog post I'll explain in more details the challenge of running such benchmarks."
enableComments: true
categories:
 - Performance
tags:
 - benchmark
 - open source
 - monitoring
 - time series database
---

One of the latest benchmarks we did was for [OSMC 2022](https://osmc.de/) talk
`VictoriaMetrics: scaling to 100 million metrics per second` -
see [the video](https://www.youtube.com/watch?v=xfed9_Q0_qU)
and [slides](https://www.slideshare.net/NETWAYS/osmc-2022-victoriametrics-scaling-to-100-million-metrics-per-second-by-aliaksandr-valialkin).

While the fact that VictoriaMetrics can handle data ingestion rate at 100 million samples per second
for one billion of active time series is newsworthy on its own, the benchmark tool used to generate that kind
of load is usually overlooked. This blog post explains the challenges of scaling the [prometheus-benchmark tool](https://github.com/VictoriaMetrics/prometheus-benchmark)
for generating such a load.

The [prometheus-benchmark tool](https://github.com/VictoriaMetrics/prometheus-benchmark) has been already mentioned
in [Benchmarking Prometheus-compatible time series databases](https://victoriametrics.com/blog/remote-write-benchmark/)
and [Grafana Mimir and VictoriaMetrics: performance tests](https://victoriametrics.com/blog/mimir-benchmark/) blog posts.

-----------------------

## Why data matters

The most crucial part of any database benchmark is **the data**. The efficiency of the read queries, indexes,
compression, and even CPU usage would depend on the "shape of data" used during the tests. For public benchmarks,
people usually pick a neutral data set and perform all the tests on it. For example,
the [New York City Taxi](https://github.com/toddwschneider/nyc-taxi-data) data set is heavily used
for [testing and comparing analytical databases](https://tech.marksblogg.com/benchmarks.html).

But what if public data sets aren't fitting your needs: are too big or too small; too specific or too broad?
The answer is usually "to generate the data yourself"! While it may sound like a simple solution, data generation is a very complex
task since data never exists apart from the application area.

For example, we want our database to store meter readings. We can generate a data set of N meters,
each having M readings. But how do we choose the values for each reading? Would it be random? No, random data isn't
a good idea in most cases, since it is too far from real-world data seen in production.

Meters aren't showing random values: each type of meter has its own distribution nature of
values (aka "data shape"). For example, temperature readings are usually related to each other and have a sinusoidal shape over days,
while voltage readings are more "chaotic":

{{< image-col-2 href-1="/blog/benchmark-100m/voltage.webp" href-2="/blog/benchmark-100m/temperature.webp" alt="Images for <a href='https://webhome.phy.duke.edu/~qelectron/proj/BooleanChaos/results.php' target = '_blank'>voltage</a> and <a href='https://stats.stackexchange.com/questions/490753/time-series-analysis-of-daily-temperature-data-in-r' target = '_blank'>temperature</a> time series are captured from the public resources as example.">}}

But the "chaotic" voltage data isn't random too - it is, at least, bounded by min and max values, speed of change, etc.
So think twice before writing your own data generator from scratch.


## TSBS - Time Series Benchmark Suite

One of the most popular benchmark suites for time series data is [TSBS](https://github.com/timescale/tsbs).
VictoriaMetrics [is represented](https://github.com/timescale/tsbs/blob/master/docs/victoriametrics.md) in the suite
as well, but we stopped using it long ago. TSBS was no match for our purposes due to the following reasons:
* TSBS ingests "random" data into the tested system. Such a data is far from most real-world cases.
* TSBS requires pre-generated data to run the benchmark on it. This means TSBS cannot be used as a continuous benchmark,
  which runs over long periods of time.
* TSBS tries ingesting the data to the tested system at the maximum possible speed, while the data in production is usually ingested at some fixed speed.
* The "read" query tests in TSBS are far from typical PromQL or [MetricsQL](https://docs.victoriametrics.com/metricsql/) queries used in production.

The way how the TSBS works assumes that the user completes the following actions one by one:
1. Pre-generate the data.
2. Ingest generated data into the tested database at the maximum possible speed.
3. Run "read" queries at the maximum possible speed.

But this is not how "real world" works. In reality, steps 2 and 3 are simultaneous and rate-limited: the database constantly processes
the ingestion load and serves read queries **at the same time**. It is very important to understand how these
two processes affect each other and how the database manages the given resources.

Another important point is tests duration: checking how the tested database behaves under constant pressure over
long periods of time. For example, in [Scaling Grafana Mimir to 500 million active series](https://grafana.com/blog/2022/05/24/scaling-grafana-mimir-to-500-million-active-series-on-customer-infrastructure-with-grafana-enterprise-metrics/)
the benchmark was running for 72h, and the authors detected an anomaly in the queries error rate with 2 hours interval.
It appeared, that the errors were caused by the scheduled compaction process that occurs every 2 hours.
Such things could reveal themselves only during continuous tests.

## Benchmark for Prometheus-compatible systems

So how could the benchmark be done better? From our experience, both Prometheus and VictoriaMetrics are mostly used
for hardware monitoring. The most popular Prometheus metrics exporter is [node-exporter](https://github.com/prometheus/node_exporter),
so we built a benchmark tool that uses it as a **source of data**. The benefits of such an approach are the following:
* The produced data is very similar to "production" data.
* The data is endless - there is no need in pre-generating it in advance.
* It is easy to scale by just deploying more node-exporter instances.

We needed something very efficient and transparent from a monitoring perspective for delivering the data
from node-exporter to the tested database. So we used [vmagent](https://docs.victoriametrics.com/vmagent.html) -
a swiss army knife for metrics collection. It is optimized for speed and it comes along with
a [Grafana dashboard](https://grafana.com/grafana/dashboards/12683) for monitoring.

For read load, we use [vmalert](https://docs.victoriametrics.com/vmalert.html) configured with typical alerting rules
used in production for the node-exporter metrics. `vmalert` captures latency and success rate of executed queries,
and also provides a [Grafana dashboard](https://grafana.com/grafana/dashboards/14950) for monitoring.

Altogether, these components form a tool that can ingest production-like data as long as it is needed. It can be easily
scaled from a thousand samples/s to 100 million samples/s with the configured number of unique time series (aka cardinality).
The source of metrics - node-exporter - can be easily replaced with any other exporter or application to provide a different set of metrics.
The same goes for read load, which can be easily adjusted by changing the list of alerting rules and the interval of their execution.

{{< image href="/blog/remote-write-benchmark/benchmark-architecture.webp" alt="Prometheus remote write benchmark architecture diagram." >}}

More details can be found [here](https://victoriametrics.com/blog/remote-write-benchmark/).

## The 100M benchmark

So how did it work for the 100M benchmark? To generate 100M samples/s for a billion of active time series, our benchmark suite
[used 16 replicas](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/bm-100/bench-overrides.yaml),
8 vCPU and 25GB RAM each.

{{< image href="/blog/benchmark-100m/dash-overview.webp" class="wide-img" alt="Screenshot of the official Grafana dashboard for VictoriaMetrics vmagent during the benchmark." >}}

The most CPU-extensive part of work for vmagent is the parsing of the scraped data. Each scrape also generates additional
pressure on Nginx and node-exporter processes. So in order to save some resources, we decided to benefit
from [VictoriaMetrics API enhancements](https://docs.victoriametrics.com/#prometheus-querying-api-enhancements).
Adding `extra_label` param to the remote write URL instructs VictoriaMetrics to add an extra label to all
the received data via this URL. This made it possible to introduce a `writeURLReplicas: N` benchmark config,
which makes vmagent fan-out the scraped data `N` times, each with unique `url_replica` label in URL.
This is why vmagent scrapes 3.13M samples/s from node-exporter instances, but writes `3.13 * 32 replicas = 101M samples/s` to the remote storage.

This is cheating a bit, but it has no significant effect on VictoriaMetrics itself. All the written data remains
unique thanks to `url_replica` label, so the database can't apply any optimizations here. On the other hand, it helps
save a lot of resources on the benchmark suite and provides a sustainable level of write load.

In total, the benchmark was able to sustain a stable **2 GB/s** data stream to the configured remote storage:

{{< image href="/blog/benchmark-100m/dash-network.webp" alt="Screenshot of the official Grafana dashboard for VictoriaMetrics vmagent during the benchmark. Network." >}}

You can find a snapshot of the vmagent dashboard during one of the test rounds [here](https://snapshots.raintank.io/dashboard/snapshot/oi1DG6s358j62LB6iVRx3RJSqVhGhoJc).

## Conclusion

We're pretty happy with the [prometheus-benchmark tool](https://github.com/VictoriaMetrics/prometheus-benchmark).
It helps us not only perform [benchmarks on a tremendous scale](https://www.slideshare.net/NETWAYS/osmc-2022-victoriametrics-scaling-to-100-million-metrics-per-second-by-aliaksandr-valialkin),
but also helps us constantly test VictoriaMetrics releases internally.

Thanks to the infinite source of data, we can run benchmarking constantly and get results very close to what
we observe in various production environments. All VictoriaMetrics components are shipped with [alerting rules](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/docker)
and [Grafana dashboards](https://grafana.com/orgs/victoriametrics/dashboards), so once something is wrong we're
getting notifications immediately. Good observability for benchmarking is critical and we're glad it helps us to keep everything transparent.

If you are evaluating different Prometheus-compatible monitoring solutions, then try testing them with the [prometheus-benchmark tool](https://github.com/VictoriaMetrics/prometheus-benchmark) on the expected workload. This will help determining the best solution for your needs.

Check other relevant blog posts:
* [Benchmarking Prometheus-compatible time series databases](https://victoriametrics.com/blog/remote-write-benchmark/)
* [Grafana Mimir and VictoriaMetrics: performance tests](https://victoriametrics.com/blog/mimir-benchmark/)
* [VictoriaMetrics Monitoring](https://victoriametrics.com/blog/victoriametrics-monitoring/)
