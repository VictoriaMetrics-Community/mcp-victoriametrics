---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2022-04-08
enableComments: true
title: "Q1/2022 Release Roundup: Announcing VictoriaMetrics v1.76 & More"
summary: "The VictoriaMetrics v1.76 release headlines our first VictoriaMetrics releases roundup blog post, which summarises all the releases we published in the first quarter of 2022; and includes feature highlights such as multi-level downsampling (the most wanted Vicky feature in 2021)."
description: ""
categories: 
 - Company News
 - Product News
tags:
 - new release
 - victoriametrics
 - new features
 - community
 - open source
 - kubernetes
keywords: 
 - new release
 - victoriametrics
 - downsampling
 - open source
 - performance
 - kubernetes
 - zero-trust
 - ARM64
images:
 - /blog/q1-2022-release-roundup/q1-2022-release-roundup.webp
---
<p><img src="/blog/q1-2022-release-roundup/q1-2022-release-roundup.webp" style="width:100%" alt="Q1/2022 Release Roundup"></p>
Since the beginning of the year, our team has been busy working with the open source community of VictoriaMetrics users and our customers as we continuously enhance and improve Vicky!

Thanks to everyone who has contributed with their feedback, questions, feature requests, bug reports, etc.

We push out a new release every 3-4 weeks, sometimes more, to make sure that the user community and our customers can benefit as quickly as possible from the new features and improvements that are being requested by users and customers alike.

We publish release notes for each new release on GitHub and will now also be publishing a release roundup blog every quarter that presents a summary of the releases published in the previous three months.

This is our first release roundup blog and the past quarter saw some great additions to VictoriaMetrics, including:

* 6 new releases
* 54 new features
* 36 bug fixes
* 23 contributors across 6 releases
* 332 commits

## Top Q1/2022 Features

* Multi-level downsampling (the most wanted Vicky feature in 2021)
* Performance for arm64 builds of VictoriaMetrics components improved by up to 15%
* Improve ingesting performance for series with high level of churn (e.g. Kubernetes metrics)
* VictoriaMetrics Cluster (Enterprise version): now includes support for mTLS communications (and Zero-Trust) between cluster components
* Vicky now has its own built-in UI (vmui) for plotting ad hoc graphs. Try it, tell us what you think!


## [Announcing VictoriaMetrics v1.76](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.76.0)


With this release roundup blog post we’re happy to announce the availability of our latest release, [VictoriaMetrics v1.76.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.76.0).


#### Feature Highlights include:

* We’ve added the ability to verify files obtained via [native export](https://docs.victoriametrics.com/#how-to-export-data-in-native-format) to [vmctl](https://docs.victoriametrics.com/vmctl.html). See [these docs](https://docs.victoriametrics.com/vmctl.html#verifying-exported-blocks-from-victoriametrics) and [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/2362).
* New feature in [VictoriaMetrics Cluster](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html): reduce memory usage by up to 50% for `vminsert` and `vmstorage` under high ingestion rate.
* Added five command-line flags, which can be used for fine-grained limiting of CPU and memory usage during various API calls.

Read the release notes for all the details on the new v1.76 release: [https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.76.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.76.0)

And see below the summaries of all of our Q1/2022 releases:

## [VictoriaMetrics v1.75.0 release - released 30 March 2022](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.75.0)

#### Feature Highlights include:

* [VictoriaMetrics Cluster](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html) (Enterprise version): now includes support for mTLS communications (and Zero-Trust) between cluster components - this is typically only needed in enterprise environments. See [these docs](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#mtls-protection) and [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/550) for more details on this feature.
* Properly free up memory occupied by deleted cache entries for the following caches: `indexdb/dataBlocks`, `indexdb/indexBlocks`, `storage/indexBlocks`. This should reduce the increased memory usage starting from v1.73.0. See [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/2242) and [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/2007) issue for more information.
* [vmalert](https://docs.victoriametrics.com/vmalert.html): add ability to use OAuth2 for `-datasource.url`, `-notifier.url` and `-remoteRead.url`. See the corresponding command-line flags containing `oauth2` in their names [here](https://docs.victoriametrics.com/vmalert.html#flags).
    * And ability to use Bearer Token for `-notifier.url` via `-notifier.bearerToken` and `-notifier.bearerTokenFile` command-line flags. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1824).

## [VictoriaMetrics v1.74.0 - release 03 March 2022](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.74.0)

#### Feature Highlights include:

* Support for conditional relabeling via `if` filter
* Improve performance when registering new time series. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/2247). Thanks to @ahfuzhang
* How to push data from vmagent to Kafka: reuse Kafka clients when pushing data from many tenants to #Kafka. See the details here: [https://docs.victoriametrics.com/vmagent.html#writing-metrics-to-kafka](https://docs.victoriametrics.com/vmagent.html#writing-metrics-to-kafka)


## [VictoriaMetrics v1.73.1 - release 22 February 2022](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.73.1)

#### Feature Highlights include:

* Allow overriding default limits for the following in-memory caches, which usually occupy the most memory:
    * `storage/tsid` - the cache speeds up lookups of internal metric ids by `metric_name{labels...}` during data ingestion. The size for this cache can be tuned with `-storage.cacheSizeStorageTSID` command-line flag.
    * `indexdb/dataBlocks` - the cache speeds up data lookups in `<-storageDataPath>/indexdb` files. The size for this cache can be tuned with `-storage.cacheSizeIndexDBDataBlocks` command-line flag.
    * `indexdb/indexBlocks` - the cache speeds up index lookups in `<-storageDataPath>/indexdb` files. The size for this cache can be tuned with `-storage.cacheSizeIndexDBIndexBlocks` command-line flag. See also [cache tuning docs](https://docs.victoriametrics.com/#cache-tuning). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1940).
* Add `-influxDBLabel` command-line flag for overriding db label name for the data [imported into VictoriaMetrics via InfluxDB line protocol](https://docs.victoriametrics.com/#how-to-send-data-from-influxdb-compatible-agents-such-as-telegraf). Thanks to [@johnatannvmd](https://github.com/johnatannvmd) for the pull request.
* Return `X-Influxdb-Version` HTTP header in responses to [InfluxDB write requests](https://docs.victoriametrics.com/#how-to-send-data-from-influxdb-compatible-agents-such-as-telegraf). This is needed for some InfluxDB clients. See [this comment](https://github.com/ntop/ntopng/issues/5449#issuecomment-1005347597) and [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/2209).


## [VictoriaMetrics v1.73.0 - released 14 February 2022](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.73.0)

Amongst other things, we’re now publishing VictoriaMetrics binaries for MacOS amd64 & MacOS arm64 (aka MacBook M1) in our release notes.

From this release onwards, performance for arm64 builds of VictoriaMetrics components improved by up to 15%!

#### Feature Highlights include:

Further Highlights include:

* Reduce CPU and disk IO usage during `indexdb` rotation once per `-retentionPeriod`. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1401).
* [VictoriaMetrics Cluster](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html): add `-dropSamplesOnOverload` command-line flag for `vminsert`. If this flag is set, then `vminsert` drops incoming data if the destination `vmstorage` is temporarily unavailable or cannot keep up with the ingestion rate. The number of dropped rows can be [monitored](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#monitoring) via `vm_rpc_rows_dropped_on_overload_total` metric at `vminsert`.
* [VictoriaMetrics Cluster](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html): improve re-routing logic, so it re-routes incoming data more evenly if some of `vmstorage` nodes are temporarily unavailable and/or accept data at slower rate than other `vmstorage` nodes. Also significantly reduce possible re-routing storm when `vminsert` runs with `-disableRerouting=false` command-line flag. This should help the following issues: [one](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1337), [two](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1165), [three](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1054), [four](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/791), [five](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1544).
* [MetricsQL](https://docs.victoriametrics.com/metricsql/): cover more cases with the [label filters' propagation optimization](https://utcc.utoronto.ca/~cks/space/blog/sysadmin/PrometheusLabelNonOptimization). This should improve the average performance for practical queries.
* [MetricsQL](https://docs.victoriametrics.com/metricsql/): optimize joining with `*_info labels`. For example: `kube_pod_created{namespace="prod"} * on (uid) group_left(node) kube_pod_info` now automatically adds the needed filters on `uid` label to `kube_pod_info` before selecting series for the right side of `*` operation. This may save CPU, RAM and disk IO resources. See [this article](https://www.robustperception.io/exposing-the-software-version-to-prometheus) for details on `*_info` labels. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/1827).

## [VictoriaMetrics v1.72.0 - released 18 January 2022](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.72.0)

#### Feature Highlights:

* Extended support for @ modifier from PromQL, which is enabled by default in Prometheus starting from Prometheus v2.33.0
    * It can contain arbitrary expression
    * It can be put anywhere in the query
* Ability to keep the original metric names when calculating arbitrary [MetricsQL](https://docs.victoriametrics.com/metricsql/) and PromQL functions. For example, rate(foo) keep_metric_names .
* Improved UX in vmui based on user requests (thank you!)

And as reminder, some feature highlights from release v1.71.0:
* Multi-level downsampling (the most wanted Vicky feature in 2021)
* Ability to compare two queries on the same graph in vmui
* Ability to read configs from http and https urls in vmagent

This sums up our release roundup for the beginning of 2022!


You can find all our release notes here: [https://github.com/VictoriaMetrics/VictoriaMetrics/releases](https://github.com/VictoriaMetrics/VictoriaMetrics/releases)

If you need support, please visit our support page: [https://victoriametrics.com/support/](https://victoriametrics.com/support/)

And if you’re interested in finding out more about our Enterprise offering, please visit: [https://victoriametrics.com/products/enterprise/](https://victoriametrics.com/products/enterprise/)


