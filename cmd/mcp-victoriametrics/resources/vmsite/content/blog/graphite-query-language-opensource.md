---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2023-04-19
title: "Releasing Graphite Query Language in Open Source VictoriaMetrics"
summary: "We are releasing Graphite Query Language  in open source VictoriaMetrics starting with VictoriaMetrics v1.90 - i.e. we’re open sourcing Graphite Query Language in VictoriaMetrics."
enableComments: true
categories: 
 - Company News
tags:
 - graphite
 - graphite query language
 - open source
 - victoriametrics
 - monitoring tools
keywords: 
 - graphite
 - graphite query language
 - open source
 - victoriametrics
 - monitoring tools
images:
 - /blog/graphite-query-language-opensource/preview.webp
---

As many of our users and the wider monitoring community will know, Graphite Query Language is a query language for Graphite monitoring tools, which helps analyze data stored in it.

Graphite is a well-known and respected pioneer in the monitoring space, which has seen a number of next generation monitoring solutions enter the scene … such as ourselves. It’s been used by a wide range of companies, which started using monitoring tools more than a decade ago. These companies [include](http://graphiteapp.org/case-studies/) Etsy, Salesforce, Reddit, GitHub and Booking.com to name but a few; Booking.com probably being [one of the largest](https://www.infoq.com/news/2019/03/graphite-scaling-booking/) implementations and users of Graphite.

Historically, VictoriaMetrics included support for [data ingestion in Graphite protocol](https://docs.victoriametrics.com/?highlight=graphite#how-to-send-data-from-graphite-compatible-agents-such-as-statsd) as well as support for the following Graphite querying APIs, which are needed for [Graphite datasource in Grafana](https://grafana.com/docs/grafana/latest/datasources/graphite/):

* [Render API](https://docs.victoriametrics.com/?highlight=graphite#graphite-render-api-usage)
* [Metrics API](https://docs.victoriametrics.com/?highlight=graphite#graphite-metrics-api-usage)
* [Tags API](https://docs.victoriametrics.com/?highlight=graphite#graphite-tags-api-usage)

Our powerful query language for VictoriaMetrics components, MetricsQL, also includes features such as Graphite-compatible filters that can be passed via `{__graphite__="foo.*.bar"}` syntax.

## **What are we announcing?**

Whereas Graphite Query Language has been a feature in VictoriaMetrics Enterprise (compatible with Graphite render API) thus far, we have decided to release it in open source VictoriaMetrics starting with [VictoriaMetrics v1.90](https://docs.victoriametrics.com/CHANGELOG.html#v1900) - in other words, we’re open sourcing Graphite Query Language in VictoriaMetrics.

## **Why are we making this announcement?**

We’ve been seeing high user demand in the VictoriaMetrics community and want to help our users migrate to modern monitoring solutions seamlessly. With the Graphite Query Language in open source VictoriaMetrics, engineers can switch the storage and get all benefits of using open source VictoriaMetrics while retaining the graphite protocol for data ingestion and querying time series.

This significantly helps with migration as VictoriaMetrics becomes a drop-in replacement for Graphite.

How to get started!

To get started with VictoriaMetrics, please visit our [Quick Start page](https://docs.victoriametrics.com/quick-start/) and [Graphite API usage](https://docs.victoriametrics.com/#graphite-api-usage). 
