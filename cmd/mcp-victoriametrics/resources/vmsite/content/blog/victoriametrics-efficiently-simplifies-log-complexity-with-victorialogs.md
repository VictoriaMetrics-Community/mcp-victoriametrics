---
draft: false
page: blog blog_post
authors:
  - Jean-Jerome Schmidt-Soisson
date: 2024-11-13
title: "VictoriaMetrics Efficiently Simplifies Log Complexity with VictoriaLogs"
summary: "We’re delighted to announce the GA release of our innovative logging solution: VictoriaLogs. It’s perfect for managing and analyzing large volumes of log data, especially in containerized environments such as Kubernetes. "
enableComments: true
categories:
  - Company News
  - Product News
tags:
  - victorialogs
  - log management
  - kubernetes
  - elastic
  - grafana
  - loki
  - wide events
  - syslog
  - log collector
  - monitoring
  - observability
images:
  - /blog/victoriametrics-efficiently-simplifies-log-complexity-with-victorialogs/preview.webp
---

# VictoriaLogs General Availability Delivers Unparalleled Performance and Scalability

**Salt Lake City, Utah, 13th November 2024** – Today we’re delighted to announce the GA release of our innovative logging solution - [VictoriaLogs](https://victoriametrics.com/products/victorialogs/).

Our easy-to-use, open source log management solution combines a powerful query language for easy log searching with minimal resource requirements. It’s perfect for managing and analyzing large volumes of log data, especially in containerized environments such as Kubernetes.

We’re also announcing  the upcoming preview-release of VictoriaLogs Cluster, with the full release scheduled for 2025.

Read the full announcement below to get all the details!

**VictoriaLogs - Key Highlights:**
* Improves query performance for haystack searches by up to 1000 times
* Uses up to 30 times less RAM and 15 times less disk space than comparable solutions
* Accepts logs from popular log collectors, including OpenTelemetry Collector, Vector, Fluentd, Logstash, Syslog, Rsyslog and Syslog-ng, Filebeat, Fluentbit, Fluentd, Logstash, Promtail, Telegraf, Journald, and DataDog
* Supports ingestion directly via Syslog protocol, removing the need for proxies and converters

*"VictoriaLogs addresses the major challenges of traditional log management tools. It significantly reduces memory usage and infrastructure costs, making it a game-changer for those frustrated by slow searches in large log volumes. By using bloom filters, VictoriaLogs accelerates haystack search times and quickly pinpoints relevant data, making queries up to 1,000 times faster than comparable Grafana Loki. Even the single-node version of VictoriaLogs is capable of replacing an Elasticsearch cluster of up to 30 nodes. It’s a high-performance tool that simplifies log management."*

<p style="text-align:center;font-weight: bold;">- Aliaksandr Valialkin, Co-founder and CTO at VictoriaMetrics.</p>

**Seamless Integration for Comprehensive Observability**
VictoriaLogs integrates seamlessly with the broader observability ecosystem. By allowing users to correlate logs with metrics, it provides a holistic view of system performance and behavior. Support for popular log shippers like Vector, Fluentd, and Logstash, enable easy adoption without disrupting existing logging workflows. This integrated approach not only simplifies the observability stack, but also enhances troubleshooting capabilities by allowing users to swiftly navigate between related logs and metrics.

**A New Era of Efficiency**
VictoriaLogs distinguishes itself from competitors by using automatically adjusted bloom filters instead of inverted indexes, allowing for accurate word searches providing definitive yes or no answers. By minimizing CPU time and disk read IO spent on unpacking, parsing, and reading logs, VictoriaLogs significantly saves computing resources at large scale.

Offering robust support for Syslog ingestion, VictoriaLogs simplifies this transition by allowing direct ingestion over Syslog without intermediaries. While Syslog remains a widely used standard, many existing solutions require additional proxies or converters, complicating the ingestion process. Users can easily migrate to VictoriaLogs by updating just the address or URL for their existing log shippers, ensuring a smooth transition.

*“VictoriaLogs sets a new standard in log management performance, addressing the demands of today's data-intensive environments. For example, where Grafana Loki creates new log streams each time a detail (such as an IP address or user ID) changes, VictoriaLogs treats these details as regular information within each log entry - resulting in fewer log streams and improving system performance and resource utilization. VictoriaLogs also achieves this while using up to 30 times less RAM and 15 times less disk space compared to Elasticsearch, making it an ideal choice for organizations dealing with massive log volumes.”*

<p style="text-align:center;font-weight: bold;">- Aliaksandr Valialkin, Co-founder and CTO at VictoriaMetrics.</p>

**Designed for Wide Events**
VictoriaLogs is optimized for efficient storing and querying of wide events containing hundreds of fields, accepting wide events with different sets of fields without the need to configure. The solution’s LogsQL query language simplifies querying wide events' stats at high speed.

**Smart architecture for cost-effective operations**
VictoriaLogs' efficient data management also extends to its disk I/O performance, addressing a critical bottleneck in log analytics. By requiring substantially less disk space, the solution dramatically reduces the volume of data read during resource-intensive queries. Industry tests have shown that this approach can accelerate query performance by up to two orders of magnitude, with heavy queries executing up to 100 times faster than on comparable solutions.
Try it out and let us know your feedback: [https://victoriametrics.com/products/victorialogs/](https://victoriametrics.com/products/victorialogs/)
