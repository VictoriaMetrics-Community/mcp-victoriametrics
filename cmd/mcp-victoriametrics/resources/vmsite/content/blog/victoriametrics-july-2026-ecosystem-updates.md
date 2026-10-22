---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-08-10
enableComments: true
title: "VictoriaMetrics July 2026 Ecosystem Updates"
summary: "VictoriaMetrics July updates bring a new LTS release, vmestimator, and a set of improvements across the stack that make operations simpler and observability more practical."
description: "VictoriaMetrics July updates: v1.148.0 LTS, v1.147.0, vmestimator, and product improvements for easier operations and better observability."
toc: true
categories:
  - Product News
  - Company News
tags:
  - victoriametrics
  - vmestimator
  - observability
  - new features
  - vmauth
  - vmagent
  - release notes
  - vmsingle
  - vmbackup
images:
  - /blog/victoriametrics-july-2026-ecosystem-updates/preview.webp
---

July brought a new long-term support line for VictoriaMetrics, providing Enterprise users with a consistent release stream and bug and security patches for the next 12 months.

We also [launched vmestimator](https://victoriametrics.com/blog/vmestimator-announcement/), a new open-source component that estimates metric cardinality in real time. It continuously tracks cardinality across label combinations, so you can spot growth and churn early, before they lead to slow queries, increased storage usage, or unexpected cost spikes.

## VictoriaMetrics v1.148.0 is out: a new LTS release

> VictoriaMetrics v1.148.0 is our latest LTS release. This means **support for **[v1.122.27](https://docs.victoriametrics.com/victoriametrics/changelog/#v112227) has ended** and will no longer receive fixes or security backports. Meanwhile [v1.136.14](https://docs.victoriametrics.com/victoriametrics/changelog/#v113614) remains supported following the usual LTS release schedule.

- New LTS release: [v1.148.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11480) is the new long-term support line, giving Enterprise users a stable release stream with bug and security fixes over the next 12 months.
- More resilient persistent queue in [vmagent](https://docs.victoriametrics.com/victoriametrics/vmagent/): ungraceful shutdowns are now much less likely to drop the entire persistent queue file, so in most cases only the latest written block is lost instead of all queued metrics.
- Fixed Prometheus Remote Write protocol downgrade in [vmauth](https://docs.victoriametrics.com/victoriametrics/vmauth/): vmauth now returns 408 Request Timeout rather than 400 Bad Request when requests time out, preventing vmagent from permanently downgrading the Prometheus Remote-Write protocol and dropping subsequent slow requests.
- More reliable S3 backups: [vmbackup](https://docs.victoriametrics.com/victoriametrics/vmbackup/) and [vmbackupmanager](https://docs.victoriametrics.com/victoriametrics/vmbackupmanager/) now retry HTTP 429 Too Many Requests responses automatically, making backups more resilient to temporary S3 rate limiting.
- Client-side least-loaded load balancing for vmagent: DNS-based service discovery can now route traffic to the least-loaded backend, which simplifies Kubernetes deployments and can remove the need for vmauth between vmagent and vminsert in many cases.
- Fill function modifiers: MetricsQL now supports [fill modifiers](https://prometheus.io/docs/prometheus/latest/querying/operators/#filling-in-missing-matches) in binary operations, which improves Prometheus compatibility and makes gaps in series easier to handle.
- Scrape metrics over Unix domain sockets: [scraping over Unix domain sockets](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/11156) is now supported in vmagent and vmsingle through the __unix_socket__ target label.
- Better relabeling debug UI: The relabeling debug UI in vmagent now auto-loads rules from -promscrape.config, -remoteWrite.relabelConfig, and -remoteWrite.urlRelabelConfig, with additional [usability improvements](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/10607#issuecomment-4956698085) for debugging relabeling flows.
- Screenshot of the new relabel UI in VMUI
- Also in this version: vmagent and vmsingle can now override max_scrape_size per target via the __max_scrape_size__ label, background discovery performance for http_sd is improved, and a new -maxBackfillAge flag limits ingestion of samples with old timestamps (useful when data moves between storage tiers). Full list in the changelog.

[Read the changelog](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11480).

## VictoriaMetrics 1.147.0 is out: easier migration to cluster

> **Important**: vmagent contains a bug that causes increased CPU and memory usage when -remoteWrite.urlRelabelConfig or -remoteWrite.streamAggr.config flags are used. The bug was introduced in [#11250](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/11250). We recommend upgrading to v1.148.0 or rolling back to v1.146.0.

- Easier migration from single-node to cluster: vmsingle now supports the vmselect clusternative RPC protocol and can act as a vmstorage node, so you can [migrate from single-node to cluster](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#from-single-node-to-cluster) with a smoother migration path (vmctl is no longer needed for the move). This also enables lightweight cluster setups with vmagent, vmsingle, and vmselect with limited [multitenancy](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#multi-tenancy) support.
- vmagent MDX routing: vmagent now supports [Monitoring Data eXchange (MDX)](https://docs.victoriametrics.com/victoriametrics/vmagent/#monitoring-data-exchange), which lets you route only VictoriaMetrics telemetry to a specific -remoteWrite.url and simplifies monitoring-of-monitoring deployments in both self-hosted mode and in our hosted MoM/MaaS offering.
- More flexible JWT auth in vmauth: the new default_vm_access_claim option makes the vm_access JWT claim optional, so existing tokens can keep working while access control still uses [claim matching](https://docs.victoriametrics.com/victoriametrics/vmauth/#jwt-claim-matching) and configured backends.
- Improved stream aggregation correctness: stream aggregation now returns more accurate rate_avg and rate_sum results.
- Also in this version: a ~10% CPU usage reduction for [sharding among remote storages](https://docs.victoriametrics.com/victoriametrics/vmagent/#sharding-among-remote-storages), a 64KiB size limit on metric metadata fields, new vm_data_size_bytes metrics for metaindex memory tracking, an optimize_repeated_binary_op_subexprs query arg for reusing rollup cache across binary-op subexpressions, an artificial 2 to 3 second delay on failed vmauth logins to slow brute-force attempts, and new alerting rules for invalid authentication tokens and rules approaching their result limits

This release includes 7 bug fixes and 14 feature improvements. [Read the changelog](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11470) for more details.

## Blogs and tutorials

- [VictoriaMetrics 2026 Mid Year Roundup](https://victoriametrics.com/blog/victoriametrics-2026-mid-year-roundup/): a quick look at the biggest things we shipped in the first half of 2026.
- [Observability for LLM Apps and Agents: OpenLIT SDK + VictoriaMetrics observability stack](https://victoriametrics.com/blog/victoriametrics-openlit-agents-observability/): a practical guide to instrumenting LLM apps and agents with OpenLIT, then sending metrics, logs, and traces into the VictoriaMetrics stack for self-hosted observability.
- [GPU Observability with the OpenLIT Collector and the VictoriaMetrics observability stack](https://victoriametrics.com/blog/victoriametrics-openlit-gpu-observability/): a walkthrough of how to monitor GPUs with OpenLIT and VictoriaMetrics, so you can connect application behavior with what is happening on the hardware underneath.

## Community news

The VictoriaMetrics team attended the [WeAreDevelopers World Congress](https://www.wearedevelopers.com/world-congress) as a bronze Sponsor. A packed room attended Aliaksandr Valialkin's (Founder and CTO) talk about [Strategies for Efficient Log Management in Large-Scale Kubernetes Clusters](https://www.wearedevelopers.com/world-congress/agenda/sessions/strategies-for-efficient-log-management-in-large-scale-kubernetes-clusters-1121797). The crew had fantastic talks during the event, and the booth was bustling with guests. Looking forward to the next WeAreDevelopers San Jose on September 23 to 25!

DevEx Talks Episode 6 focused on [neurodiversity in tech](https://www.youtube.com/watch?v=x6YHeuv4fHc), with Diana Todea speaking with Carl Alexander and Zach Stepek about their experiences as neurodivergent professionals. The conversation highlights what has helped them succeed, what has not, and why community plays a crucial role in long-term growth and support.

VictoriaMetrics latest Virtual Meetup took place on July 16 and [is now available on YouTube](https://www.youtube.com/watch?v=E9DauzMUSO8&list=PLXT8DSiuv5yl47-tv4Tl1Uty6nLNX6ar1). In this episode, the engineering team shares product updates, roadmap highlights, new features, cloud updates, and community news.

## Wrap up

This month's updates focus on stability, visibility, and easier operations across the VictoriaMetrics stack. From the new LTS release to vmestimator and the other improvements, the goal is to make day-to-day monitoring simpler and more predictable.

Stay tuned for more ecosystem updates from the VictoriaMetrics team. Thank you for reading!
