---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2024-05-30
title: "VictoriaMetrics slashes data storage bills by 90% with world’s most cost-efficient monitoring"
summary: "We’re happy to share customer research today demonstrating that VictoriaMetrics is the world's most cost-efficient monitoring solution! Read the post for details!"
enableComments: true
categories:
 - Company News
tags:
 - victoriametrics
 - monitoring
 - open source
 - prometheus
 - metrics
 - cost-efficiency
 - cost savings
keywords:
 - monitoring
images:
 - /blog/victoriametrics-slashes-data-storage-bills-with-worlds-most-cost-efficient-monitoring/preview.webp
---

## VictoriaMetrics outpaces open-source standard Prometheus’ query latency by 16x

We’re happy to share customer research today demonstrating that VictoriaMetrics is the world's most cost-efficient monitoring solution! 

The real-world results show customers can save energy costs and achieve Net Zero carbon compliance faster with VictoriaMetrics in their tech stack.

A combination of VictoriaMetrics’ optimized data structures and efficiently coded algorithms reduces the energy costs for data processing and [storage by up to 90%](https://www.grammarly.com/blog/engineering/monitoring-with-victoriametrics/). Compared to similar solutions such as InfluxDB and Prometheus, VictoriaMetrics customers and users need significantly less Central Processing Unit (CPU) compute power and Random Access Memory (RAM) storage space. In practice, this can result in 10x cloud cost savings and also save [up to x4](https://victoriametrics.com/blog/victoriametrics-remote-write/) network costs.

At a time of increased pressure on IT budgets, VictoriaMetrics minimizes the hardware needed for data storage reducing the reliance on physical servers which can be expensive, unreliable and hard to manage. VictoriaMetrics requires less hardware and compute resources to run, without sacrificing performance and subsequently saving further money on equipment upgrades.

## VictoriaMetrics laps Prometheus

As a drop-in replacement for the most widely-used open-source monitoring system in the world, Prometheus, VictoriaMetrics provides a more efficient software solution. In a benchmark test, [VictoriaMetrics outperformed Prometheus](https://victoriametrics.com/blog/reducing-costs-p1/), using 1.7x less memory, 2.5x less disk space, and delivering 16x faster query latency on average. This significant improvement in performance translates into substantial cost savings.

Existing monitoring solutions come with hefty price tags and limitations in scalability due to their complex, resource-intensive architectures. These limitations force businesses to choose between monitoring a subset of their data, or incurring significant expenses. VictoriaMetrics' open-source approach shatters this cost barrier and businesses can optimize performance, reduce costs associated with data movement, and gain a comprehensive view of their entire data landscape – avoiding the high licensing fees and hardware investments typically associated with proprietary software.

<i>"Businesses need a monitoring solution that can keep up with the ever-increasing volume of data without breaking the bank", said Roman Khavronenko, Co-Founder of VictoriaMetrics. "With VictoriaMetrics, businesses can benefit from the power of open-source software, smart data management techniques, and scalability, while enjoying unparalleled cost efficiency."</i>

## VictoriaMetrics in action

Organizations are increasingly drawn to VictoriaMetrics due to its simplicity, reliability and efficiency. The Compact Muon Solenoid (CMS) experiment at [CERN turned to VictoriaMetrics](https://docs.victoriametrics.com/casestudies/#cern) after encountering storage and scalability issues. The architecture of VictoriaMetrics minimizes the energy needed for data processing and storage by as much as 90% in comparison to other comparable technologies.

Similarly, [Roblox](https://www.datanami.com/2023/05/30/why-roblox-picked-victoriametrics-for-observability-data-overhaul/) and [Grammarly](https://www.grammarly.com/blog/engineering/monitoring-with-victoriametrics/) turned to VictoriaMetrics to help with their monitoring needs. With over 200 million active monthly users, Roblox can now ingest 120 million data points per second into its VictoriaMetrics cluster, whilst Grammarly can retain 18 months’ worth of data without any concern about exorbitant storage costs. Grammarly’s costs were reduced by 10x with VictoriaMetrics.

## Stripping 'dependency bloat'

Many organizations struggle with dependency bloat, which occurs when a software project has an excessive number of unnecessary dependencies, resulting in slower performance, larger size, and increased complexity. By removing these unnecessary dependencies, VictoriaMetrics not only improves its performance, but also significantly enhances cost efficiency. By taking only the necessary functionalities from libraries and implementing tailored solutions, VictoriaMetrics reduces  resource usage, leading to lower infrastructure costs. 

## VictoriaLogs, soon to be announced…

VictoriaMetrics will soon be announcing the general availability of its open source, user friendly database for logs - VictoriaLogs. Similarly to VictoriaMetrics, VictoriaLogs is more cost-efficient and reliable than other solutions on the market, such as Elasticsearch and Grafana Loki. VictoriaLogs’ increased efficiency means it can handle up to 30x bigger data volumes than competitors, whilst running on the same hardware.

## Getting Started

Get started with [VictoriaMetrics](https://docs.victoriametrics.com/quick-start/?_gl=1*rrrckw*_ga*MTI4MTAxNDMxMC4xNzEwNzExNTM2*_ga_N9SVT8S3HK*MTcxNzAwMTk0OC43OC4xLjE3MTcwMDE5NTAuNTguMC4w)

Get started with [VictoriaLogs](https://docs.victoriametrics.com/victorialogs/quickstart/?_gl=1*1oxff7v*_ga*MTI4MTAxNDMxMC4xNzEwNzExNTM2*_ga_N9SVT8S3HK*MTcxNzAwMTk0OC43OC4xLjE3MTcwMDIwMzMuNTYuMC4w)
