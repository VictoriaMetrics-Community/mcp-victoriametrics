---
draft: false
page: blog blog_post
authors:
  - Jean-Jerome Schmidt-Soisson
date: 2022-01-06
enableComments: true
title: "With 36M+ Downloads, VictoriaMetrics Skyrockets to New Heights: 2021 in Review"
summary: "This ‘VictoriaMetrics 2021 Momentum Milestones’ blog post provides a summary of this year’s main achievements with our top features, blogs and talks"
categories: 
  - Company News
tags:
  - victoriametrics
  - momentum
  - 2021
  - achievements
  - open source
  - database
  - monitoring
  - timeseries
images:
  - /blog/momentum-2021.webp
---

We took advantage of the quiet days between holidays to look back on the year past and thank our users and customers for their  support in 2021 - and wish you a very happy 2022!

As our co-founder, Roman recently pointed out: “You can’t improve what you don’t measure!”

Our aim is to make monitoring simple, fast  and reliable for everyone by providing an open source time series database for monitoring that has what it takes to become a standard component of modern observability stacks. We love it when we hear from our users that “VictoriaMetrics just works.”

Whether you are a longtime or new member of the VictoriaMetrics Community, please share in this year’s success with our main highlights and stats.
<p>&nbsp;</p>

# 2021 VictoriaMetrics Momentum Milestones

  - 36M+ Docker Pulls
  - 800K+ GitHub Downloads 
  - 5K+ GitHub Stars for the VictoriaMetrics time series database
  - 7K+ GitHub Stars for [VictoriaMetrics overall](https://coderstats.net/github/#victoriametrics)
  - New [customer & users wins](https://victoriametrics.com/case-studies/) with Grammarly, DFKI, GrooveX, WiX, CERN, and more
  - Launched the new [victoriametrics.com](https://victoriametrics.com)

2021 has been a great year for VictoriaMetrics with 36M+ downloads, 20+ releases, thousands of new users and customers, and we’re planning to increase this momentum in 2022 with the community’s support.

VictoriaMetrics, the high performance open source time series database and monitoring solution, is fast, easy-to-use, and optimized for high cardinality. It’s also highly scalable on cloud, kubernetes or on-premise setups.

We’ve been delighted to see the great uptake of our product and the vibrant user community that’s organising itself around it.

We are proudly a self-funded startup that generates profitability from [services that we offer](https://victoriametrics.com/support/) in support to VictoriaMetrics as well as it’s [Enterprise version](https://victoriametrics.com/products/enterprise/). Our team is laser-focused on solving our customer and community user needs, while constantly perfecting and enhancing our software.

We have a very short and prolific release cycle, which can be followed on our GitHub page: [https://github.com/VictoriaMetrics](https://github.com/VictoriaMetrics)

This blog post provides a summary of our main achievements this year with our top features, blogs and talks.

<a href="https://victoriametrics.com" rel="some text"><img src="/blog/momentum-2021.webp" style="width:100%" title="momentum"/></a>

## Top 5 New VictoriaMetrics Features

- [vmoperator](https://docs.victoriametrics.com/guides/getting-started-with-vm-operator.html?highlight=vmoperator)
  - This is a Kubernetes operator for automated provisioning, scaling & management. This feature is highly popular and currently our most downloaded component.

- [UI for VictoriaMetrics](https://github.com/VictoriaMetrics/vmui)
  - vmui is - as the name suggests, a user interface for VictoriaMetrics, which makes it even easier to get started and to manage your monitoring solution.

- [Clustering in vmagent](https://docs.victoriametrics.com/vmagent.html#scraping-big-number-of-targets)
  - A single vmagent instance can scrape tens of thousands of scrape targets, but this isn't always enough due to limitations of CPU, network, RAM, etc. With this new feature, scrape targets can be split among multiple vmagent instances (aka vmagent horizontal scaling, sharding and clustering).

- [extra_labels](https://docs.victoriametrics.com/Single-server-VictoriaMetrics.html#prometheus-querying-api-enhancements)
  - Help to enforce filters on user's queries. This feature is particularly  powerful as it  allows multi-tenancy on a single server (without the need for clustering).

- [vmalert replay](https://docs.victoriametrics.com/vmalert.html#rules-backfilling)
  - Allows users to evaluate recording and alerting rules from the past and backfill results of these evaluations back to the database.

## Top 3 New VictoriaMetrics Enterprise Features

- [Downsampling](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#downsampling)
  - Also referred to as Rollups: this is a tool that rewrites old data with a configurable, lower sample rate - and helps achieve significant savings on storage costs.

- [Graphite Query API](https://docs.victoriametrics.com/#graphite-api-usage)
  - Provides the ability to query data using Graphite Query Language in addition to the currently supported MetricsQL and PromQL.

- [Kafka Integration](https://docs.victoriametrics.com/vmagent.html?highlight=Kafka%20integration#kafka-integration)
  - Supports reading (consumer) & writing (producer) from Kafka & lets users build highly reliable data pipelines across multiple regions.

Read Aliaksandr Valialkin’s [2021 VictoriaMetrics Features Roundup](/blog/features-roundup-2021/) for all the details!

## Top 3 Blogs

- [VictoriaMetrics: PromQL Compliance](https://medium.com/@romanhavronenko/victoriametrics-promql-compliance-d4318203f51e) by Roman Khavronenko
  - This blog post was written and posted in response to discussions that were taking place online as to whether or not VictoriaMetrics is fully compatible with PromQL. While our MetricsQL is backward-compatible with PromQL in the sense that Grafana dashboards backed by a Prometheus datasource should work the same after switching from Prometheus to VictoriaMetrics, VictoriaMetrics overall is not 100% compatible with PromQL and we believe is better for it. [Please read](https://medium.com/@romanhavronenko/victoriametrics-promql-compliance-d4318203f51e) on as Roman discusses why that is.
  
- [How to optimize PromQL and MetricsQL queries](https://valyala.medium.com/how-to-optimize-promql-and-metricsql-queries-85a1b75bf986) by Aliaksandr Valialkin
  - PromQL and MetricsQL are powerful query languages. They allow writing simple queries to build nice looking graphs of time series data. They also allow users to write sophisticated queries for SLI / SLO calculations and alerts, but it may be hard to optimize PromQL queries. This article shows how to determine slow PromQL queries, how to understand query costs and how to optimize these queries so they execute faster and consume lower amounts of CPU and RAM. [Read the blog](https://valyala.medium.com/how-to-optimize-promql-and-metricsql-queries-85a1b75bf986) for all the details.
  
- [How to monitor Go applications with VictoriaMetrics](https://victoriametrics.medium.com/how-to-monitor-go-applications-with-victoriametrics-c04703110870) by Roman Khavronenko
  - Monitoring is fun! It is so fun that once you get started you’ll never leave any of your apps without some fancy metrics. But sometimes beginners are afraid to touch this area, mostly because the rest of the tech appears overwhelming with complexity, standards, and conventions. In this article, Roman shows how simple it can be to start using metrics, storing them in VictoriaMetrics TSDB and visualizing via Grafana. [Read the blog](https://victoriametrics.medium.com/how-to-monitor-go-applications-with-victoriametrics-c04703110870) for all the details.


## Top 3 Talks

- [Open Source Strategy at VictoriaMetrics](https://www.youtube.com/watch?v=-DbbIZzFHIY) by Roman Khavronenko
  - Building a company around a free software product isn’t something new. What's less common is creating a company in order to build a free software product. This talk by Roman covers our story of the creation of a time series database, the lessons we learned, the mistakes we made. While the free software world has changed over the last few years, one thing remains essential: the importance of a community, i.e. people who use the product. [View the talk here.](https://www.youtube.com/watch?v=-DbbIZzFHIY)
  
- [Migration From Prometheus to VictoriaMetrics for Percona’s PMM](https://www.youtube.com/watch?v=kB8SXNpET14)  by Aliaksandr Valialkin and Roma Novikov
  - Recently, PMM replaced Prometheus with VictoriaMetrics. In the talk we want to cover the motivation behind this transition, the architecture and internals of PMM and technical details of the replacement. The talk was given by members of both organizations who took part in the migration: Percona and VictoriaMetrics. [View the talk here.](https://www.youtube.com/watch?v=kB8SXNpET14)
  
- [How ClickHouse Inspired Us to Build a High Performance Time Series Database](https://www.youtube.com/watch?v=p9qjb_yoBro) by Aliaksandr Valialkin
  - Join Aliaksandr Valialkin as he walks you through the internals of the processing pipeline inside the VictoriaMetrics time series database, the architectural decisions made and the optimizations used for getting the highest performance possible at OSACon 2021. [View the talk here.](https://www.youtube.com/watch?v=p9qjb_yoBro)

Thanks again for your support this past year - have a successful 2022!

Happy New Year from everyone at VictoriaMetrics! Here’s to continuously improving and innovating!

PS.: If you would be interested in learning more about our Enterprise features or getting more personalized support, [please click here](https://victoriametrics.com/products/).
