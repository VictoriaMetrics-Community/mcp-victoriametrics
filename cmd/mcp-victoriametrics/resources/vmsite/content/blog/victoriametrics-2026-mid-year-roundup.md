---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-07-13
featuredPost: true
enableComments: true
title: "VictoriaMetrics 2026 Mid Year Roundup"
summary: "Take a look at what we shipped in the first half of 2026 across VictoriaMetrics: metrics, logs, traces, and cloud,with improvements in performance, scalability, and usability."
toc: true
categories:
  - Product News
  - Company News
tags:
  - victoriametrics
  - victorialogs
  - victoriatraces
  - observability
  - open source
  - new features
  - community
  - talks
images:
  - /blog/victoriametrics-2026-mid-year-roundup/preview.webp
---

In the first half of 2026, we shipped a wide range of improvements across metrics, logs, traces, cloud, and the Kubernetes operator. Our main focus across open-source components and enterprise solutions was on performance, stability, and making observability easier to adopt and operate day‑to‑day.

This roundup brings together the most important changes to date, including a quick look back at key anomaly detection improvements from 2025 that are now paying off today.

## The VictoriaMetrics Observability Stack In Numbers

Across all VictoriaMetrics OSS products we have recorded so far...

- 3300+ additional GitHub stars (raising the total above 21,000 ⭐️ across all our repositories).
- 4,476,294 downloads from GitHub.
- 32,670,000 downloads from Docker Hub and Quay.
- 70+ releases.
- 30+ active third-party contributors.
- 622 issues closed.
- 1,715 PRs merged (excluding bots).

## First Half of 2026 Retrospective

### Open‑source projects

VictoriaMetrics gained multiple improvements, including better proxy support, VMUI enhancements, safer backend routing, improved CORS behavior, Prometheus-native histogram ingestion support, and fixes for ingestion stalls and startup panics.

VictoriaLogs continued to evolve with [Splunk HEC support](https://victoriametrics.com/blog/victoriametrics-april-2026-ecosystem-updates/), CSV export, Kubernetes‑aware collection improvements, new LogsQL operators, and a more powerful Web UI for exploration.

VictoriaTraces saw active development in 2026, with 10 new releases (and counting). The main themes were API compatibility, [Tempo query support](https://docs.victoriametrics.com/victoriatraces/querying/grafana/#grafana-tempo-datasource), ingestion reliability, and better observability of the tracing pipeline. These changes landed just as work began to bring VictoriaTraces into VictoriaMetrics Cloud.

On the operations side, the VictoriaMetrics Kubernetes Operator [introduced a Long‑Term Support (LTS) line (v0.68.x)](https://victoriametrics.com/blog/operator-lts/) with a clearly defined support window and a new dry‑run mode.

### Enterprise and Cloud

In 2026, VictoriaMetrics introduced a new LTS line ([v1.136.x](https://docs.victoriametrics.com/victoriametrics/changelog/#v113612)) with multiple performance and stability improvements across ingestion and querying, as well as important bug fixes for OpenTelemetry ingestion and query evaluation precedence.

[VictoriaLogs became generally available in VictoriaMetrics Cloud](https://victoriametrics.com/blog/q1-2026-whats-new-victoriametrics-cloud/), bringing a fully managed log solution for Cloud users. VictoriaMetrics Cloud also added a [hosted MCP Server](https://victoriametrics.com/blog/q1-2026-whats-new-victoriametrics-cloud/), a new rules editor with a richer preset catalog, more transparent token limits, PrivateLink in all regions, and incremental backups.

VictoriaMetrics Anomaly Detection (vmanomaly) was built on 2025's foundations with new UI ergonomics, an [AI copilot built into the UI](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2026-q1/), additional playgrounds, and model‑agnostic controls to reduce false positives.

## Announcing VictoriaMetrics v1.146.0

June brought a new release of [VictoriaMetrics](https://docs.victoriametrics.com/victoriametrics/). Version 1.146.0 introduces the following changes:

- Stream Aggregation: Use the aggregation rule interval as the [default staleness_interval](https://docs.victoriametrics.com/victoriametrics/stream-aggregation/configuration/#stream-aggregation-config), rather than `2 * interval`, reducing spikes when there are gaps between received samples.
- vmagent: add the [-remoteWrite.inmemoryQueues flag](https://docs.victoriametrics.com/victoriametrics/vmagent/#on-disk-persistence-and-data-processing-order) to prioritize recently ingested data over historical data stored in the file-based persistent queue, so real-time metrics stay visible when the on-disk queue grows too large and takes hours (or days) to drain; this can help users recover observability during backlog buildup, but it may introduce gaps or out-of-order samples for older data still in the queue.
- vmagent: add [-promscrape.cluster.shardByLabels flag](shardByLabels) for choosing target labels used to shard scrape targets among vmagent instances in cluster mode. This supports setups that shard by service name and then use stream aggregation to drop the instance label in the same vmagent.
- vmstorage: prevent more cases of panic during directory deletion on NFS-based mounts.
- Metrics metadata: fix several issues in metadata parsing and exposure.

Read the [full changelog](https://docs.victoriametrics.com/victoriametrics/changelog/#v11460).

## Community, Airbnb, OpenAI, and Conferences

Community usage continued to increase, with two notable high-level use cases involving VictoriaMetrics OSS.

In April, [Airbnb described migrating its high-volume metrics pipeline](https://victoriametrics.com/blog/how-airbnb-built-a-high-volume-metrics-pipeline-with-opentelemetry-and-vmagent/) from StatsD and Veneur to OpenTelemetry and vmagent. Their setup processes 100 million samples per second in a single production cluster.

In February, the [OpenAI Harness experiment used the VictoriaMetrics stack](https://victoriametrics.com/blog/observability-lessons-from-openai/) (VictoriaMetrics, VictoriaLogs, and VictoriaTraces) as the observability layer for a coding loop, providing metrics, logs, and traces to Codex, bringing unified observability to AI-driven workflows. For similar setups, Alexander’s guide on [observing AI agents](https://victoriametrics.com/blog/ai-agents-observability/index.html) with OpenTelemetry and VictoriaMetrics provides a practical reference.

On the events side, the team participated in [KubeCon Amsterdam and CloudNativeCon Europe](https://victoriametrics.com/blog/victoriametrics-at-kubecon-amsterdam/), as well as regional conferences including LinuxFest Northwest, StackConf, Observability Summit North America, and Linux NorthWest Fest. Presentations focused on reducing observability costs, designing multi-tier architectures, and operating OpenTelemetry with VictoriaMetrics at scale. Meetups and AMAs were also used to gather user feedback to plan the roadmap.

## What's Next

Looking ahead to the rest of 2026, the focus remains on deepening the integration across metrics, logs, traces, anomaly detection, and the managed cloud platform while continuing to prioritize performance and operability. On the cloud side, that means bringing VictoriaTraces to VictoriaMetrics Cloud, completing SOC 2 certification, simplifying private connectivity, and evolving cluster pricing towards a compute‑based model with clearer visibility into consumption.

