---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2023-04-19
title: "Q1 Roadmap Review & Q2 2023 Look Ahead"
summary: "Read about our early achievements in 2023, the roadmap for VictoriaMetrics, initial details on the upcoming VictoriaLogs, as well as where to find our team in the coming weeks."
enableComments: true
categories: 
 - Company News
tags:
 - open source
 - victoriametrics
images:
 - /blog/q1-roadmap-review-2023/preview.webp
---
In our recent [virtual meetup](https://www.youtube.com/watch?v=Gu96Fj2l7ls&t=1780s) the VictoriaMetrics  Founders team discussed some of our Q1 2023 highlights, including features highlights, the 2023 roadmap for VictoriaMetrics as well as first introduction to the upcoming VictoriaLogs.

In this blog post, we’d like to share a summary of these highlights and a heads up on where to find our team in the coming weeks and months - starting with our participation at [KubeCon Europe 2023](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/).

Let’s look at some of our highlights for the beginning of this year!

## **VictoriaMetrics 2023 Q1 Stats**

The Vicky community and our team have been busy with contributions as you can see from the stats below - thank you so much to everyone involved!

* 180+ issues
260+ PRs
* 40 contributors
* 12 releases, from 1.86 to 1.89:
  * 114 FEATURES
  * 103 BUG FIXES

## **What's new in VictoriaMetrics at Q1 2023**

### **[vmalert](https://docs.victoriametrics.com/vmalert.html) - in [VictoriaMetrics Enterprise](https://victoriametrics.com/products/enterprise/)**

* GCS and S3 support for config rules

### **Streaming Aggregation**

* StatsD alternative
  * Counting input samples
  * Summing input metrics
  * Quantiles over input metrics
  * Histograms over input metrics
* More efficient recording rules alternative
* Reducing the number of stored samples

{{< image href="/blog/q1-roadmap-review-2023/image1.webp" >}}

### **[VictoriaMetrics Remote Write Protocol](https://docs.victoriametrics.com/vmagent.html?highlight=%20remote%20write#victoriametrics-remote-write-protocol])**

* The VictoriaMetrics remote write protocol allows reducing network traffic costs by 2x-4x!

Read this blog also for more details: [Save network costs with VictoriaMetrics remote write protocol](https://victoriametrics.com/blog/victoriametrics-remote-write/).


### **[vmui Features](https://play.victoriametrics.com/select/accounting/1/6a716b0f-38bc-4856-90ce-448fd713e3fe/prometheus/graph/#/?g0.range_input=30m&g0.end_input=2023-04-18T13%3A36%3A52&g0.relative_time=last_30_minutes)**

* Dark theme
* Explore mode
* Sticky tooltips
* Cardinality explorer

## **VictoriaMetrics Roadmap Review: What’s Coming up**

* OpenTelemetry ingestion protocol support
* vmalert: UI for rules management
* vmalert: hysteresis support
* Features that were in the roadmap for 2023 and are already released:
  * vmui explore tab
  * [Grafana Data Source Plugin](https://github.com/VictoriaMetrics/grafana-datasource)
  * Grafana datasource plugin: query trace

[Watch the replay of our March virtual meet up](https://www.youtube.com/watch?v=Gu96Fj2l7ls), where Roman walks us through the current roadmap for VictoriaMetrics.

## **[VictoriaLogs Preview](https://www.youtube.com/watch?v=Gu96Fj2l7ls)**

The big news of course is the preview and pre-announcement of the upcoming VictoriaLogs!

Here’s a summary of the highlights:

**What is VictoriaLogs**

* Open source log management system from VictoriaMetrics
* Easy to setup and operate
* Scales vertically and horizontally
* Optimized for low resource usage (CPU, RAM, disk space)
* Accepts data from Logstash and Fluentbit in Elasticsearch format
* Accepts data from Promtail in Loki format
* Supports stream concept from Loki
* Provides easy to use yet powerful query language - LogsQL

For the full run-down, [watch Aliaksandr](https://www.youtube.com/watch?v=Gu96Fj2l7ls) provide some of the initial details on VictoriaLogs.

Finally, let’s look at where our team can be found over the next few months as well some as of our recent news. 

## **Where to Find VictoriaMetrics**

This is a summary of where you can find us over the coming weeks / months: We'd love to meet you and talk to you in person; or chat with you during online talks.

* 17th of April : [Cloud Native Rejekts](http://eventbrite.com/e/465064138357) in Amsterdam
* 17th - 21st of April: [KubeCon Europe](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/) in Amsterdam
* 15th of May: [SloConf](https://www.sloconf.com/) (Online)
* 22nd to 24th of May: [Percona Live](https://www.percona.com/live/conferences) in Denver
* 26th to 28th of June: [Monitorama](https://monitorama.com/2023/pdx.html) in Portland
* 26th to 28th of June: [GopherCon Europe](https://gophercon.eu/) in Berlin

We're looking forward to meeting you at one or more of these events!

If there are events that you know and recommend we take part in, please let us know!


## **Recent VictoriaMetrics Blogs**

* [Save network costs with VictoriaMetrics remote write protocol](https://victoriametrics.com/blog/victoriametrics-remote-write/)
* [VictoriaMetrics Long-Term Support (LTS): Commitment, Current and Next LTS Versions](https://victoriametrics.com/blog/lts-status-h1-2023/)
* [Rules backfilling via vmalert](https://victoriametrics.com/blog/rules-replay/)
* [Monitoring benchmark: how to generate 100 million samples/s of production-like data](https://victoriametrics.com/blog/benchmark-100m/)
* [Latest updates about backup components of VictoriaMetrics](https://victoriametrics.com/blog/latest-updates-for-backup-compnents-2023-q1/)

## **VictoriaMetrics in the News**

* ComputerWeekly: [Is it time for time series databases?](https://www.computerweekly.com/blog/Open-Source-Insider/Is-it-time-for-time-series-databases)
* datanami: [Open source time series database VictoriaMetrics sees significant growth](https://www.datanami.com/2023/01/26/open-source-times-series-database-victoriametrics-sees-significant-growth/)
* Data Centre & Network News: [VictoriaMetrics announces 252% growth in 2022](https://dcnnmagazine.com/news/victoriametrics-growth-2022/)
* Information Age: [Kubernetes Best Monitoring Tools](https://www.information-age.com/kubernetes-monitoring-best-tools-123501818/)
* Intelligent CIO: [VictoriaMetrics leads sustainable monitoring - slashing corporate energy usage by up to 90%](https://www.intelligentcio.com/north-america/2023/03/02/victoriametrics-leads-sustainable-data-monitoring-slashing-corporate-energy-usage-by-up-to-90/)
* Connected Technology Solutions: [Sustainable data monitoring for high performance](https://connectedtechnologysolutions.co.uk/sustainable-data-monitoring-for-high-performance/?utm_content=240538211&utm_medium=social&utm_source=twitter&hss_channel=tw-1009364910281805824)
* The Stack: [One to watch - The story of a startup from Ukraine](https://thestack.technology/one-to-watch-victoriametrics-the-story-of-a-startup-from-ukraine/amp/)
* ComputerWeekly: [KubeCon + CloudNativeCon 2023: Container no-brainers](https://www.computerweekly.com/blog/Open-Source-Insider/KubeCon-CloudNativeCon-2023-Container-no-brainers)

As always, we welcome your feedback and questions, so feel free to use the comments box below!

