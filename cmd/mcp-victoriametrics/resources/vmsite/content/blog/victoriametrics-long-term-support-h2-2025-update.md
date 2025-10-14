---
draft: false
page: blog blog_post
authors:
  - Adam Yates
  - Artem Navoiev
date: 2025-09-26
enableComments: true
title: "VictoriaMetrics Long-Term Support (LTS): H2 2025 Update"
summary: "As we’re half-way through the year, we’d like to take this opportunity to provide an update on the most recent changes in our Long-Term Support (LTS) releases."
categories: 
 - Product News
 - Company News
tags:
 - victoriametrics
 - LTS release
 - long-term support
 - victoriametrics enterprise
images:
 - /blog/lts-status-h2-2025/preview.webp
---

![VictoriaMetrics LTS release timeline](/blog/lts-status-h2-2025/victoriametrics-lts-release-timeline.webp)
<figcaption style="text-align: center; font-style: italic;">VictoriaMetrics LTS release timeline</figcaption>

As we’re into the second half of the year, we’d like to take this opportunity to provide an update on the most recent changes in our Long-Term Support (LTS) releases.

LTS releases are published for the Enterprise versions of VictoriaMetrics and are designed for production workloads under SLA, providing long-term support lines of releases for VictoriaMetrics customers.

Every LTS line receives bug fixes and security fixes for 12 months after the initial release. New LTS lines are published every 6 months, meaning the latest two LTS lines are supported at any given moment. This gives up to 6 months for the migration to new LTS lines for [VictoriaMetrics Enterprise](https://docs.victoriametrics.com/victoriametrics/enterprise/) users.

## Important for our users

While the latest LTS versions are an enterprise-license feature, VictoriaMetrics Enterprise is based on open source, meaning version 1.122.x is open source and everyone can update to it.

All the bugfixes and security fixes, which are included in LTS releases, are available in [the latest open source release](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/latest). Please upgrade VictoriaMetrics products regularly to [the latest available open source release](https://docs.victoriametrics.com/victoriametrics/changelog/).

## Introducing our latest LTS release

We’ve recently released **VictoriaMetrics LTS** [**v1.122.1**](https://docs.victoriametrics.com/victoriametrics/changelog/#v11221), which begins a new line of **Long-Term Support** releases. This version will receive **bugfixes and security patches for 12 months**, ensuring maximum stability for production environments.

The previous LTS release, [v1.111](https://docs.victoriametrics.com/victoriametrics/changelog/#v11110), will continue to receive bug fixes until February 2026.

### Please note:

The **v1.102 LTS** line has reached its end-of-life and is **no longer supported**. We strongly recommend upgrading to **v1.122.1** to continue receiving critical updates.

### More information about LTS releases:

[https://docs.victoriametrics.com/victoriametrics/lts-releases/](https://docs.victoriametrics.com/victoriametrics/lts-releases/)

## Update notes for versions 1.111-1.122:

### VictoriaMetrics Single-node and vmstorage in VictoriaMetrics cluster
The *-snapshotsMaxAge* flag default has been changed to 3d. This enables automatic deletion of snapshots older than 3 days. If you want to keep the previous behavior (never automatically deleting snapshots), please set *-snapshotsMaxAge=0*.<br />
[See GitHub issue #9344 for details](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9344)

[*Full changelog for v1.122.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11220)

### vmagent
The *-retryMaxTime* flag has been deprecated. Please use *-retryMaxInterval* flag instead.<br />
[See GitHub issue #9169 for more details](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9169)

[*Full changelog for v1.121.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11210)

### Stable tag for Docker images will no longer be updated

You need to use specific version tags for docker images to continue receiving updates.<br />
[See GitHub issue #7336 for more details](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7336)

[*Full changelog for v1.117.1*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11171)

### Updated the RPC cluster protocol version for the TSDB status API

Calls to */api/v1/status/tsdb* may temporarily fail until **vmstorage** and **vmselect** are updated to the same version.

[*Full changelog for v1.116.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11160)

### vmagent data distribution changed

[vmagent](https://docs.victoriametrics.com/victoriametrics/vmagent/)’s data distribution algorithm of remote write is changed from round-robin to consistent hashing when *-remoteWrite.shardByURL* is enabled. This means **vmagents** with *-remoteWrite.shardByURL* will re-shard series after the upgrade, which may result in temporary higher churn rate and memory usage on remote destinations.<br />
[See GitHub issue #8546 for more details](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8546)

[*Full changelog for v1.116.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11160)

### Metric *vm_mmaped_files* renamed

Metric **vm_mmaped_files** was renamed to **vm_mmapped_files** to fix the typo in word mmapped.

[*Full changelog for v1.114.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11140)

### IPv6 addresses fix for VictoriaMetrics Single-node & vmagent

[**VictoriaMetrics Single-node**](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/) and [**vmagent**](https://docs.victoriametrics.com/victoriametrics/vmagent/) include a fix which enforces IPv6 addresses escaping for containers discovered with [Kubernetes service discovery](https://docs.victoriametrics.com/victoriametrics/sd_configs/#kubernetes_sd_configs) and role: pods which do not have exposed ports defined. This means that addresses for these containers will always be wrapped in square brackets. This might affect some relabeling rules which were relying on previous behaviour.

[*Full changelog for v1.113.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11130)

### vmalert disallows using time-buckets stats pipe

[**vmalert**](https://docs.victoriametrics.com/victoriametrics/vmalert/) disallows using [time-buckets stats pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-time-buckets) in alerting or recording rules with [VictoriaLogs](https://docs.victoriametrics.com/victorialogs/) as datasource. Time buckets used with [stats query API](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-stats) may produce unexpected results for users and result in cardinality issues.

[*Full changelog for v1.113.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11130)

### vmalert disallows specifying eval_offset and eval_delay options

[vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/) disallows specifying *eval_offset* and *eval_delay* options in the same [group](https://docs.victoriametrics.com/victoriametrics/vmalert/#groups). The *eval_offset* option ensures the group is evaluated at the exact offset in the range of [0…interval]. However, with *eval_delay*, this behavior cannot be guaranteed without further adjusting the evaluation time, which could lead to more confusion.

### Single-node VictoriaMetrics and vmstorage stop exposing vm_index_search_duration_seconds histogram metric

This metric records time spent on search operations in the index. It was introduced in [v1.56.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.56.0). However, this metric was used neither in [dashboards](https://grafana.com/orgs/victoriametrics/dashboards) nor in [alerting rules](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/docker/rules). It also has high cardinality because index search operations latency can differ by 3 orders of magnitude. Hence, dropping it as unused.

[*Full changelog for v1.111.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11110)

## Update notes for VictoriaMetrics Enterprise

### All the VictoriaMetrics Enterprise components

*-eula* command-line flag is skipped when validating the VictoriaMetrics Enterprise license. Instead, the *-license* or *-licenseFile* command-line flags must be used to provide a valid license key.<br />
[See these docs for configuration examples](https://docs.victoriametrics.com/victoriametrics/enterprise/)

[*Full changelog for v1.122.0*](https://docs.victoriametrics.com/victoriametrics/changelog/#v11220)

## Full changelog:

[https://docs.victoriametrics.com/victoriametrics/changelog/](https://docs.victoriametrics.com/victoriametrics/changelog/)

## Please note

Security patches are only guaranteed in LTS lines (Enterprise) or in [the latest open source release](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/latest).

If you're using the open source version, be advised to regularly [upgrade](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-upgrade-victoriametrics) VictoriaMetrics products to [their latest available releases](https://docs.victoriametrics.com/victoriametrics/changelog/).

To learn more about LTS releases for the Enterprise versions of VictoriaMetrics, [get in touch](https://victoriametrics.com/contact-us/).

If you have any questions or need assistance, feel free to reach out to our team in our [public Slack channel](https://inviter.co/victoriametrics).
