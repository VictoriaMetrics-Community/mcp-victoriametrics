---
draft: false
page: blog blog_post
authors:
 - Aliaksandr Valialkin
date: 2022-01-07
enableComments: true
title: "What’s new in VictoriaMetrics 2021?"
summary: "With more than 20 new releases of VictoriaMetrics published during 2021, a features roundup seemed appropriate. This blog walks you through the key VM features released in 2021."
categories: 
 - Company News
 - Product News
tags:
 - features roundup
 - product releases
 - open source
 - monitoring
 - time series
 - database
 - victoriametrics
 - 2021
---
The 2021 year is finished, so it's time to look at changes VictoriaMetrics has gained during the past year. The first release in 2021 was [v1.52.0](https://docs.victoriametrics.com/CHANGELOG.html#v1520). The last release in 2021 was [v1.71.0](https://docs.victoriametrics.com/CHANGELOG.html#v1710). More than 20 new releases of VictoriaMetrics were published during the 2021. The full changelog is available at [this page](https://docs.victoriametrics.com/CHANGELOG.html). Let's look at the most interesting changes.


## Querying Graphite data

VictoriaMetrics was able to accept Graphite data from 2019 - see [these docs](https://docs.victoriametrics.com/#how-to-send-data-from-graphite-compatible-agents-such-as-statsd). It provides much lower disk space usage for the ingested Graphite data and uses much lower disk IO compared to the default Graphite database - [Whisper](https://graphite.readthedocs.io/en/latest/whisper.html). The ingested Graphite data could be queried via [MetricsQL](https://docs.victoriametrics.com/metricsql/) - PromQL-compatible query language. While MetricsQL and PromQL are [powerful query languages](https://valyala.medium.com/promql-tutorial-for-beginners-9ab455142085), Graphite users wanted to use Graphite query language directly from VictoriaMetrics, since this allows seamless migration from Graphite to VictoriaMetrics without the need to modify dashboards and alerting rules in Grafana. So we added support for Graphite query language into VictoriaMetrics in the beginning of 2021 - see [these docs](https://docs.victoriametrics.com/#graphite-api-usage) for details. This allows using VictoriaMetrics as a drop-in replacement for other Graphite backends.

Additionally, we introduced the following Graphite-related features to VictoriaMetrics in 2021:

* `__graphite__` pseudo-label in MetricsQL, which allows selecting metrics with [Graphite query paths and wildcards](https://graphite.readthedocs.io/en/latest/render_api.html#paths-and-wildcards). For example, `{__graphite__=~"dc*.{appA,appB}.host*.memory"}` matches Graphite metrics for the following Graphite wildcard - `dc*.{appA,appB}.host*.memory`. See [these docs](https://docs.victoriametrics.com/#selecting-graphite-metrics) for more details.
* `label_graphite_group(q, groupNum1, …, groupNumN)` function, which allows selecting particular groups from Graphite metric names. See [these docs](https://docs.victoriametrics.com/metricsql/#label_graphite_group).
* Ability to use Graphite queries for alerting and recording rules in vmalert. See [these docs](https://docs.victoriametrics.com/vmalert.html#graphite).


## Downsampling for the stored data

Downsampling automatically reduces the number of stored raw samples per time series. This may help reducing disk space usage, since lower number of samples are stored to disk. This also may improve performance for queries over long time ranges, since lower number of samples are scanned during the query. VictoriaMetrics gained support for downampling in 2021 - see [these docs](https://docs.victoriametrics.com/#downsampling) for details.


## Web UI for data exploring and troubleshooting

VictoriaMetrics now supports web UI for data exploring and troubleshooting - see [vmui docs](https://docs.victoriametrics.com/#vmui). It supports the same set of query args as `/graph` web UI in Prometheus, so it can be seamlessly used for query troubleshooting when editing graphs in Grafana.


## Prometheus-compatible staleness markers

Prometheus-compatible staleness markers fix many edge cases when monitoring highly dynamic environments with frequently changed Prometheus-compatible scrape targets (for example, Kubernetes monitoring). See [these docs](https://docs.victoriametrics.com/vmagent.html#prometheus-staleness-markers) for details.


## Horizontally scalable scraping for Prometheus-compatible targets

[vmagent](https://docs.victoriametrics.com/vmagent.html) gained the ability to spread scrape targets among multiple vmagent instances. This may be useful if you need to scrape tens of thousands of targets. In this case a single vmagent instance can reach scalability limits, so the solution is to spread the targets among multiple vmagent instances. This is easy to do now according to [these instructions](https://docs.victoriametrics.com/vmagent.html#scraping-big-number-of-targets).


## Backfilling for recording rules in vmalert

[vmalert](https://docs.victoriametrics.com/vmalert.html) gained support for backfilling historical data for recording rules in 2021. See [these docs](https://docs.victoriametrics.com/vmalert.html#rules-backfilling).


## vmctl tool

[vmctl tool](https://docs.victoriametrics.com/vmctl.html) has been introduced in 2020 in a separate Github project. This tool can be used for migrating data from other monitoring systems to VictoriaMetrics. vmctl has been moved into the main VictoriaMetrics repository in 2021, so it is released together with other VictoriaMetrics components. See [vmctl docs](https://docs.victoriametrics.com/vmctl.html) for more details.

vmctl tool gained support for [data migration from OpenTSDB to VictoriaMetrics](https://docs.victoriametrics.com/vmctl.html#migrating-data-from-opentsdb) in 2021 thanks to [John Seekins](https://github.com/johnseekins).


## Other interesting features

VictoriaMetrics gained many other features during 2021. Below is a list of the most interesting features:

* [All the enterprise components of VictoriaMetrics](https://victoriametrics.com/products/enterprise/) became available for download and evaluation at [releases page](https://github.com/VictoriaMetrics/VictoriaMetrics/releases) and at [DockerHub](https://hub.docker.com/u/victoriametrics).

* [Ability to accept metrics from DataDog agent and DogStatsD](https://docs.victoriametrics.com/Single-server-VictoriaMetrics.html#how-to-send-data-from-datadog-agent).

* [Support for data reading and data writing from/to Kafka](https://docs.victoriametrics.com/vmagent.html#kafka-integration).

* [The official alerting rules for VictoriaMetrics components](https://docs.victoriametrics.com/#monitoring).

* Many new MetricsQL functions, which were requested by our users - see [these docs](https://docs.victoriametrics.com/metricsql/#metricsql-functions).

* Ability to set additional labels for all the ingested samples when ingesting data into VictoriaMetrics via supported HTTP-based data ingestion protocols. Additional labels can be set via `extra_label` query argument. See [these docs](https://docs.victoriametrics.com/#how-to-import-time-series-data).

* Ability to set additional label filters to all the queries (MetricsQL and Graphite) via `extra_label` and `extra_filters[]` query args. This can be used for label-based multitenancy setup. See [these docs](https://docs.victoriametrics.com/#prometheus-querying-api-enhancements).

* Numerous improvements to the official Grafana dashboards for VictoriaMetrics components. See [VictoriaMetrics dashboards for Grafana](https://grafana.com/orgs/victoriametrics).

* Improved routing rules and load balancing for incoming queries in vmauth. See [these docs](https://docs.victoriametrics.com/vmauth.html#auth-config).

* Official Windows builds for vmagent, vmctl, vmauth and vmalert are now published at [releases page](https://github.com/VictoriaMetrics/VictoriaMetrics/releases).

* Improved service discovery in vmagent for Kubernetes, EC2, GCE and Consul.

* Added Prometheus-compatible service discovery in vmagent for Docker, DockerSwarm, DigitalOcean and [http_sd_config](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#http_sd_config).

* Reduced memory usage and improved query performance for typical workloads at VictoriaMetrics. Reduced memory usage for vmagent.

* Official ARM and ARM64 builds for VictoriaMetrics components are now published at [releases page](https://github.com/VictoriaMetrics/VictoriaMetrics/releases) and at [DockerHub](https://hub.docker.com/u/victoriametrics).

* Ability to limit the number of unique time series during data ingestion at vmagent and VictoriaMetrics. See [these](https://docs.victoriametrics.com/#cardinality-limiter) and [these](https://docs.victoriametrics.com/#cardinality-limiter) docs.

* [Multitenant support in vmalert](https://docs.victoriametrics.com/vmalert.html#multitenancy).

* [Multitenant support in vmagent](https://docs.victoriametrics.com/vmagent.html#multitenancy).

* Ability to filter `/api/v1/status/tsdb` output with arbitrary label filters - see [these docs](https://docs.victoriametrics.com/#tsdb-stats).

* Improved logging for VictoriaMetrics components aimed towards simplified troubleshooting.

* [Ability to read Prometheus-compatible scrape configs from multiple files](https://docs.victoriametrics.com/vmagent.html#loading-scrape-configs-from-multiple-files).

* Ability to read file-based configs such as `-promscrape.config` from http and https urls.

* [Web UI in vmalert for alerting and recording rules](https://docs.victoriametrics.com/vmalert.html#web).

* Improved relabeling - see [these docs](https://docs.victoriametrics.com/vmagent.html#relabeling).

* [Ability to scrape targets via http, https or socks5 proxies](https://docs.victoriametrics.com/vmagent.html#scraping-targets-via-a-proxy).

* [Automatic switching to read-only mode under low amounts of free disk space](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/269).

* [Improved compatibility with PromQL](https://medium.com/@romanhavronenko/victoriametrics-promql-compliance-d4318203f51e).


## Conclusion

There were many useful changes in VictoriaMetrics during 2021 thanks to our users and customers. The full changelog is available [here](https://docs.victoriametrics.com/CHANGELOG.html). Probably it is time to upgrade VictoriaMetrics to newer versions according to [these instructions](https://docs.victoriametrics.com/#how-to-upgrade-victoriametrics). The 2022 year should bring more new features and improvements. If you miss some features in VictoriaMetrics, then please file a feature request of vote for already existing features at [GitHub project for VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/issues).
