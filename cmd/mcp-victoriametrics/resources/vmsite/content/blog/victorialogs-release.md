---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2023-06-22
enableComments: true
title: "VictoriaMetrics bolsters move from monitoring to observability with VictoriaLogs release"
summary: "Read the announcement blog about the release of VictoriaLogs, our new open source logs management solution."
categories: 
 - Company News
tags:
 - new release 
 - victorialogs
 - logs management
 - logging solution
 - open source
 - monitoring
 - observability
images:
 - /blog/victorialogs-release/preview.webp
---

## **New scalable open source solution simplifies logs complexity for users & enterprises**

Today we’re happy to announce our new open source, scalable logging solution, [VictoriaLogs](/products/victorialogs), which helps users and enterprises expand their current monitoring of applications into a more strategic ‘state of all systems’ enterprise-wide observability.

Many existing logging solutions on the market today offer IT professionals a limited window into live operations of databases and clusters. 

As operations scale up globally and complexity rises, these earlier generations of monitoring solutions cannot offer the granular, yet simplified view needed to observe accurately how time series data affects business performance in near real-time. 

[VictoriaLogs](/products/victorialogs), built by engineers with previous experience at Google and Lyft, solves this.

It is built upon the same principles that drive VictoriaMetrics since its inception: simplicity, reliability and cost-efficiency. 

## **Key Highlights of VictoriaLogs**

* Requires up to 10x less disk space and RAM than ElasticSearch on production workloads
* Easier to configure & operate than ElasticSearch and Grafana Loki
* [LogsQL](/products/logsql): A simple,  yet powerful query language

{{< image href="/blog/victorialogs-release/image1.webp">}}

For details on this benchmark, please view: 

[https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/logs-benchmark](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/logs-benchmark)


## **Open Source Benchmark**

Don’t take our word for it and run the VictoriaLogs preview alongside existing solutions in production, compare their resource usage and share your experience.

We’ve published an open source benchmark  that lets you do just that.

Try it out here and let us know your results: [https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/logs-benchmark](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/logs-benchmark)

## **Efficient, easy-to-use monitoring for better observability**

Highly efficient in operation, [VictoriaLogs](/products/victorialogs) works with both structured and unstructured logs for maximum backwards compatibility with the complex large-scale infrastructure needed by users whether they are academic or commercial, working in e-commerce or video gaming teams.

Designed for ease of install and simplicity of use, VictoriaLogs speeds up the analysis of infrastructure performance and the mean time to resolution for the complex issues which appear quickly in live time-series environments and where every second matters. 

VictoriaLogs also dramatically improves the observability of systems to help businesses identify and analyze database performance issues, debug them and predict future behavior. This empowers internal and external teams to reduce both the immediate and longer term business risks of outages, to smooth product updates and to enhance customer experiences. 

## **Faster log querying**

To further improve usability, VictoriaMetrics’s new Query Language, [LogsQL](/products/logsql), is an easy to use yet powerful Query Language for VictoriaLogs. It delivers a full-text search capability, which allows querying arbitrary analytics for billions of logs. VictoriaLogs accepts logs from existing logging agents, pipelines and streams, and efficiently stores them in a highly-optimized log database which can then be queried using LogsQL at lightning-fast speeds.

Roman Khavronenko, Co-Founder of VictoriaMetrics says, “As engineers, we are all too familiar with frustrations caused by modern logging systems that create further complexity, rather than removing it. Logs have been around far longer than monitoring and so it is easy to forget just how useful they can be for modern observability. We’ve successfully created a simple and easy-to-use monitoring solution that scales easily and turning to logs was a natural next step. [VictoriaLogs](/products/victorialogs) gives under-pressure teams enhanced observability of complex systems and their interactions. It is perfect for teams who need to rapidly identify data outliers and anomalies or identify site reliability and availability.”

VictoriaLogs fits ideally for SRE, DevOps and system engineers who need to provide logging solutions as a platform for the entire company or team. 

It’s available to trial today with general availability from summer 2023. 

Over time it will accept data from FluentBit, Promtail and Logstash and is suitable for running in Kubernetes, Docker, cloud, edge or bare metal.

## **Related resources**

* [VictoriaLogs product page](/products/victorialogs)
* [LogsQL page](/products/logsql)
* [VictoriaLogs & LogsQL on GitHub](https://github.com/VictoriaMetrics/VictoriaMetrics)
* [VictoriaLogs & LogsQL documentation](https://docs.victoriametrics.com/victorialogs/)
* [Open source benchmark](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/logs-benchmark)

