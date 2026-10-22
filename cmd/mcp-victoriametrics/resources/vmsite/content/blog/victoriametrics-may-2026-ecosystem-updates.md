---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-06-01
enableComments: true
title: "VictoriaMetrics May 2026 Ecosystem Updates"
summary: "May 2026 VictoriaMetrics release roundup: v1.144.0 brings 15 bug fixes and 9 UX improvements for reliability and observability, while v1.143.0 adds native Prometheus histogram ingestion support across vmagent, vmsingle, and vminsert. Also includes the first LTS release for VictoriaMetrics Operator."
toc: true
categories:
  - Product News
  - Company News
tags:
  - victoriametrics
  - observability
  - new features
images:
  - /blog/victoriametrics-may-2026-ecosystem-updates/preview.webp
---

The VictoriaMetrics Observability Stack included two releases for VictoriaMetrics and a [new LTS release for the VictoriaMetrics Operator](https://victoriametrics.com/blog/operator-lts/) in May.

We've also published two detailed articles in the last few weeks:

- [How Airbnb Built a High-Volume Metrics Pipeline with OpenTelemetry and vmagent](https://victoriametrics.com/blog/how-airbnb-built-a-high-volume-metrics-pipeline-with-opentelemetry-and-vmagent/): an in-depth guide on how Airbnb created a massive aggregation pipeline using vmagent.
- [Multi-tiered Observability: A Practical Way to Handle Diverse Workloads](https://victoriametrics.com/blog/multi-tier-observability-a-practical-way-to-handle-diverse-workloads/): a practical architecture for handling diverse workloads without forcing every signal into the same storage and query path.

This release roundup covers updates for:

- [VictoriaMetrics](#vm)

## VictoriaMetrics v1.144.0 brings many bugfixes and improvements {#vm}

This release includes 15 bug fixes and 9 feature improvements focused on reliability, observability, and UX.

- Improved VMUI Top Queries UX: Enhanced user experience for the Top Queries feature in [VMUI](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#vmui).
- vmagent graceful shutdown: [vmagent](https://docs.victoriametrics.com/victoriametrics/vmalert/) now drains in-memory queues on graceful shutdown before falling back to the persisted queue, reducing delivery delays and helping avoid disruptions to alerts and recording rules.
- Improved [vminsert](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#) slowness-rerouting: Enhanced slowness-rerouting logic to avoid rerouting storms.
- VMUI Raw Query debugging: VMUI Raw Query now displays NaN, stale markers, and null values for easier debugging.
- Fixed vmagent rerouting bug: Fixed unintended rerouting in vmagent with `-remoteWrite.shardByURL`.
- VMUI graph tooltips: Graph tooltips now show exact values without rounding for precise metric inspection.
- Bug fixes: 15 bug fixes are included in this release.
- Quality-of-life improvements: Various small improvements to enhance user experience

Read the [changelog](https://docs.victoriametrics.com/victoriametrics/changelog/#v11440) to see all that's changed.

## VictoriaMetrics v1.143.0 is out with native Prometheus histogram support

- [Prometheus native histogram](https://prometheus.io/docs/specs/native_histograms/) support: VictoriaMetrics now supports ingesting Prometheus native histograms in vmagent, single-node, and cluster deployments. During ingestion, native histograms are automatically converted into the VictoriaMetrics [histogram format](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#histogram) and exposed as classic histogram series such as `_count`, `_sum`, and `_bucket`, which can be queried with standard functions like [`histogram_quantile()`](https://docs.victoriametrics.com/victoriametrics/metricsql/#histogram_quantile).
- Bug fixes and improvements: The release includes various bug fixes and small quality-of-life improvements.

Read the [changelog](https://docs.victoriametrics.com/victoriametrics/changelog/#v11430) for more details.

## Community News

As always, VictoriaMetrics was active in the community during May, enjoying the energy of two major conferences:

- **Observability Summit North America**: We joined engineers and practitioners discussing the latest trends in observability, sharing insights on handling diverse workloads and building scalable monitoring solutions.
- **Linux NorthWest Fest**: We engaged with the Linux community at this regional festival, connecting with developers and sysadmins who rely on VictoriaMetrics for their monitoring needs.

Looking ahead, we're excited to participate in:

- **SouthEast Linux Fest (SELF)**: June 12-14, 2026, in Charlotte, North Carolina. Join us to connect with the Southeast Linux community and learn how VictoriaMetrics powers reliable, cost-effective monitoring at scale.
- **DevConf.cz**: June 18-19, in Brno, Czech Republic. The VictoriaMetrics team will present 4 talks over two days of the conference. These sessions showcase VictoriaMetrics expertise across large-scale data analysis, OpenTelemetry integration, intelligent anomaly detection, and community engagement.

See you at next June's roundup. Thanks for reading!


