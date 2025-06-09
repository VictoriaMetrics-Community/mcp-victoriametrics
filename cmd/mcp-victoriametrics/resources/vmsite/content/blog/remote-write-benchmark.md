---
draft: false
page: blog blog_post
authors:
 - Roman Khavronenko
date: 2022-01-17
title: "Benchmarking Prometheus-compatible time series databases"
summary: "A Helm chart for pushing node_exporter metrics to Prometheus-compatible systems via remote_write protocol"
enableComments: true
categories: 
 - Company News
 - Product News
tags:
 - benchmark
 - open source
 - monitoring
 - time series database
 - Prometheus
 - victoriametrics
 - node_exporter
images:
 - /blog/remote-write-benchmark/benchmark-architecture.webp
---
  
## Prometheus remote_write benchmark

Some time ago, Aliaksandr Valialkin published a medium post about 
[comparing VictoriaMetrics and Prometheus resource usage](https://valyala.medium.com/prometheus-vs-victoriametrics-benchmark-on-node-exporter-metrics-4ca29c75590f)
when scraping metrics from thousands of targets. He used [node_exporter](https://github.com/prometheus/node_exporter) 
as a source for metrics to scrape, which is very close to most real-world scenarios. 
However, the benchmark itself was just a bunch of scripts and a lot of manual work for every test.

For running internal comparisons between different VictoriaMetrics versions or between VictoriaMetrics and other solutions 
with [Prometheus remote_write protocol](https://docs.google.com/document/d/1LPhVRSFkGNSuU1fBd81ulhsCPR4hkSZyyBj1SZ8fWOM/edit#heading=h.n0d0vphea3fe) 
support we created [Prometheus-benchmark](https://github.com/VictoriaMetrics/prometheus-benchmark).
The idea behind this is very simple:
- `node_exporter` is used as a source of production-like metrics;
- `nginx` is used as caching proxy in front of `node_exporter`. It reduces the load on `node_exporter` when too many concurrent scrapes are happening;
- `vmagent` is used for scraping `node_exporter` metrics and forwarding them via Prometheus `remote_write` protocol to the configured
destinations. If multiple destinations are set multiple vmagent instances independently push the scraped data to these destinations.


<p><img src="/blog/remote-write-benchmark/benchmark-architecture.webp" style="width:100%" alt="benchmark architecture"></p>

Please note, the benchmark does not collect metrics from the configured remote_write destinations.
It collects metrics for its internal components - `vmagent` and `vmalert`, so they can be inspected later.
It is assumed that the monitoring of the tested Prometheus storage systems is done separately - see [these docs](https://github.com/VictoriaMetrics/prometheus-benchmark#monitoring).

Let's go through the most important configuration settings.

### Targets count

[targetsCount](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/f6a69052413618c607758d5469e43e508792aff7/chart/values.yaml#L10)
defines how many node_exporter scrape targets are added to vmagent's scrape config (each with unique `instance` label).
This param affects the volume of scraped metrics and cardinality. Typically, one node_exporter produces around 800 unique metrics.

### Scrape interval

[scrapeInterval](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/f6a69052413618c607758d5469e43e508792aff7/chart/values.yaml#L16)
defines how frequently to scrape each target. This param affects data ingestion rate. The lower the interval, the higher
 the data ingestion rate is.

### Remote storages

[remoteStorages](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/f6a69052413618c607758d5469e43e508792aff7/chart/values.yaml#L50)
contains a list of tested systems where to push the scraped metrics. If multiple destinations
are set multiple vmagent instances individually push the same data to multiple destinations.

### Churn rate

[scrapeConfigUpdatePercent](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/f6a69052413618c607758d5469e43e508792aff7/chart/values.yaml#L30)
and [scrapeConfigUpdateInterval](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/f6a69052413618c607758d5469e43e508792aff7/chart/values.yaml#L38)
can be used for generating non-zero [time series churn rate](https://docs.victoriametrics.com/FAQ.html#what-is-high-churn-rate),
which is typical in Kubernetes monitoring.

## How do we use it?

A typical scenario is to run multiple VictoriaMetrics installations and list their addresses
in [remoteStorages](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/f6a69052413618c607758d5469e43e508792aff7/chart/values.yaml#L50) section.
The default config for such tests is `targetsCount=1000` and `scrapeInterval=10s` which results in about 80k samples/s:

`800 metrics-per-target * 1k targets / 10s = 80k samples/s`

We have separate monitoring for every remote-write destination, so later we can compare the resource usage, data compression
and overall performance via [the official Grafana dashboards for VictoriaMetrics](https://grafana.com/orgs/victoriametrics/dashboards).

## Bonus: read load

As a bonus, the helm chart also contains a [vmalert](https://docs.victoriametrics.com/vmalert.html)
configuration for running read queries.
These are standard [alerting rules](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/main/chart/files/alerts.yaml)
for node_exporter. Running vmalert is optional and allows generating more production-like workload, where metrics storage
receives production-like read requests additionally to data ingestion. 
The [alerting rules file](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/main/chart/files/alerts.yaml) can be easily
replaced with custom set of rules. The frequency of rules evaluation is controlled 
by the [queryInterval](https://github.com/VictoriaMetrics/prometheus-benchmark/blob/f6a69052413618c607758d5469e43e508792aff7/chart/values.yaml#L21) parameter.


## Conclusion

The benchmark proved to be useful for our internal tests. We believe that the community may also benefit from it
when comparing different solutions or versions of the same solution, which accept data via Prometheus remote_write protocol. For example, [Prometheus itself](https://prometheus.io/docs/prometheus/latest/storage/#overview), [Cortex](https://cortexmetrics.io/docs/api/#remote-write), [Thanos](https://thanos.io/tip/components/receive.md/), [M3DB](https://github.com/m3db/m3/blob/master/site/content/reference/m3coordinator/api/remote.md) and [TimescaleDB](https://docs.timescale.com/promscale/latest/send-data/prometheus/).
However, we always recommend to not simply believe synthetic
benchmarks, but validate the numbers and resource usage on production-like data.
