---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2023-07-12
title: "Q2 Round Up: Roadmap Review & Q3 2023 Look Ahead "
summary: "Read about our Q2 achievements in 2023, the roadmap for VictoriaMetrics, the launch of VictoriaLogs, and more!"
enableComments: true
categories:
 - Company News
tags:
 - victoriametrics
 - roadmap
 - achievements
 - open source
 - database
 - monitoring
 - timeseries
 - victorialogs
 - logs management
images:
 - /blog/q2-round-up-roadmap-review-q3-2023-look-ahead/preview.webp
---
Many thanks to everyone who joined us for our recent [virtual meetup](https://www.youtube.com/watch?v=yt0ukL5X2pQ&t=239s), during which we discussed some of our Q2 2023 highlights, including features highlights, the 2023 roadmap for VictoriaMetrics and of course:

[The launch of  VictoriaLogs!](/blog/victorialogs-release/)

In this blog post, we’d like to share a summary of these highlights.

## **What's new in VictoriaMetrics at Q2 2023**

### 1. [Releasing Graphite Query Language in Open Source VictoriaMetrics](/blog/graphite-query-language-opensource/)

Support for Graphite querying APIs needed for [Graphite datasource in Grafana](https://grafana.com/docs/grafana/latest/datasources/graphite/):

* [Render API](https://docs.victoriametrics.com/?highlight=graphite#graphite-render-api-usage)
* [Metrics API](https://docs.victoriametrics.com/?highlight=graphite#graphite-metrics-api-usage)
* [Tags API](https://docs.victoriametrics.com/?highlight=graphite#graphite-tags-api-usage)

Starting from VictoriaMetrics [v1.90](https://docs.victoriametrics.com/CHANGELOG.html#v1900)

Read more at [Graphite API usage docs](https://docs.victoriametrics.com/#graphite-api-usage) and/or our [announcement blog](/blog/graphite-query-language-opensource/).

{{< image href="/blog/q2-round-up-roadmap-review-q3-2023-look-ahead/image1.webp" src="/blog/graphite-query-language-opensource/" >}}

### 2. VictoriaMetrics is Now on Windows!

Release Windows binaries for:

* [single-node VictoriaMetrics](https://docs.victoriametrics.com/)
* [cluster VictoriaMetrics](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html)
* [vmbackup](https://docs.victoriametrics.com/vmbackup.html) and [vmrestore](https://docs.victoriametrics.com/vmrestore.html)

Starting from VictoriaMetrics [v1.90](https://docs.victoriametrics.com/CHANGELOG.html#v1900)

{{< image href="/blog/q2-round-up-roadmap-review-q3-2023-look-ahead/image2.webp" src="https://docs.victoriametrics.com/quick-start/" >}}



### 3. [New Features in vmui](https://play.victoriametrics.com/)

* Heatmap support
* Relabeling playground
* WITH template playground
* Cardinality explorer change tracking
* vmui detects expressions matching no series

See it all in action here: [https://play.victoriametrics.com/](https://play.victoriametrics.com/)

{{< image href="/blog/q2-round-up-roadmap-review-q3-2023-look-ahead/image3.webp" src="https://play.victoriametrics.com/" >}}

### 4. vmauth Improvements

* Filter incoming requests by IP
* Proxy requests to the specified backends for unauthorized users
* Default route for unmatched requests
* Automatically retry POST requests on the remaining backends if the currently selected backend isn't reachable

See here for details: [https://docs.victoriametrics.com/vmauth.html](https://docs.victoriametrics.com/vmauth.html)

### 5. vmagent Kafka Integration Improvements

* Support for Kafka producer and consumer on arm64 machines
* Allow tuning of consumer concurrency via -kafka.consumer.topic.concurrency
* Support [VictoriaMetrics remote write protocol](https://docs.victoriametrics.com/vmagent.html#victoriametrics-remote-write-protocol) while consuming or pushing data to Kafka

See here for details: [https://docs.victoriametrics.com/vmagent.html#kafka-integration](https://docs.victoriametrics.com/vmagent.html#kafka-integration)
 

### 6. New Features in vmalert

* vmalert detects alerting rules with no matching series
* Alerts for alerting and recording rules

{{< image href="/blog/q2-round-up-roadmap-review-q3-2023-look-ahead/image4.webp" src="https://docs.victoriametrics.com/vmalert.html" >}}

See here for details: [https://docs.victoriametrics.com/vmalert.html](https://docs.victoriametrics.com/vmalert.html)


## **2023 Roadmap Review**

The following features are currently in the works:

* Grafana datasource plugin
* Grafana datasource plugin: WITH templates support
* OpenTelemetry ingestion protocol support

For a more detailed roadmap review update, please watch this extract of our recent virtual meetup: [https://www.youtube.com/watch?v=yt0ukL5X2pQ&t=3323s](https://www.youtube.com/watch?v=yt0ukL5X2pQ&t=3323s)

## **[Announcing VictoriaLogs](/blog/victorialogs-release/)**

The big news last quarter for us was of course the launch of the long-awaited VictoriaLogs!

Built by engineers for engineers, VictoriaLogs is our new open source, scalable logging solution.
It is built upon the same principles that drive VictoriaMetrics since its inception:

* Simplicity
* Reliability
* Cost-efficiency

Key highlights include:

* Requires up to 10x less disk space and RAM than ElasticSearch on production workloads
* Easier to configure & operate than ElasticSearch and Grafana Loki
* [LogsQL](/products/logsql): A simple, yet powerful query language

Get all details here:

* [Announcement blog](/blog/victorialogs-release/)
* [Product page](/products/victorialogs/)
* [Documentation](https://docs.victoriametrics.com/victorialogs/)
* [LogsQL page](/products/logsql/)

{{< image href="/blog/q2-round-up-roadmap-review-q3-2023-look-ahead/image5.webp" src="/products/victorialogs/">}}

## **[Q2 VictoriaMetrics Virtual Meetup](https://www.youtube.com/watch?v=yt0ukL5X2pQ&t=3323s)**

For a more detailed description and discussion of the topics covered in this blog, please watch the recording of our second virtual meetup this year.
Thanks to all of you who participated in the meetup - you can watch the recording here:

[https://www.youtube.com/watch?v=yt0ukL5X2pQ&t=3323s](https://www.youtube.com/watch?v=yt0ukL5X2pQ&t=3323s)

## **VictoriaMetrics in the News!**

We’ve had some nice press coverage in the past few months as well as some community driven articles that have been published - thank you for these, and please let us know if you’d like to talk to us about what we do, or if you have any suggestions on articles that we could help with also.

* Forbes: [The Agility in Cloud Observability](https://www.forbes.com/sites/adrianbridgwater/2023/07/05/the-agility-in-cloud-observability/)
* Datanami: [Why Roblox picked VictoriaMetrics for Observability Data Overhaul](https://www.datanami.com/2023/05/30/why-roblox-picked-victoriametrics-for-observability-data-overhaul/)
* ComputerWeekly: [KubeCon 2023 - Container No-brainers](https://www.computerweekly.com/blog/Open-Source-Insider/KubeCon-CloudNativeCon-2023-Container-no-brainers)
* Digitalization World: [Turbocharge your data monitoring while slashing costs](https://sdc-channel.news/blogs/57501/turbocharge-your-data-monitoring-whilst-slashing-costs)
* vm-blog: [Pushing your platform full stream ahead](https://vmblog.com/archive/2023/04/03/pushing-your-platform-full-stream-ahead.aspx#.ZCrgKi1Q1p9)

<p></p>

* RTFM (English): [VictoriaMetrics: An overview and its use instead of Prometheus](https://rtfm.co.ua/en/victoriametrics-an-overview-and-its-use-instead-of-prometheus/)
* RFTM (Ukrainian): [VictoriaMetrics: знайомство та використання замість Prometheus](https://rtfm.co.ua/victoriametrics-znajomstvo-ta-vikoristannya-zamist-prometheus/)
* dbi-services: [Prometheus vs VictoriaMetrics](https://www.dbi-services.com/blog/prometheus-vs-victoriametrics/)
* dbi-services: [Migrating Monitoring Data from Prometheus to VictoriaMetrics](https://www.dbi-services.com/blog/migrating-monitoring-data-from-prometheus-to-victoriametrics/)

## **Our Recent Blog Posts**

Catch up on our latest blog posts and please let us know if there are any topics you’d like to see us cover in upcoming blogs.

* [VictoriaMetrics bolsters move from monitoring to observability with VictoriaLogs release](/blog/victorialogs-release/)
* [Never-firing alerts: What they are and how to deal with them](/blog/never-firing-alerts/)
* [How to use VictoriaMetrics for monitoring with Netdata Agent](/blog/using-victoriametrics-and-netdata/)
* [Releasing Graphite Query Language in Open Source VictoriaMetrics](/blog/graphite-query-language-opensource/)
* [Q1 Roadmap Review & Q2 2023 Look Ahead](/blog/q1-roadmap-review-2023/)

This sums up our Q2 2023!

As always, we welcome your feedback and questions, so feel free to use the comments box below!


