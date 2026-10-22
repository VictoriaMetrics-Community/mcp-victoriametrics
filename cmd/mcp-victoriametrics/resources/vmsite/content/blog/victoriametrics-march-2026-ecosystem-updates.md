---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-04-05
enableComments: true
title: "VictoriaMetrics March 2026 Ecosystem Updates"
summary: "March 2026 VictoriaMetrics Stack releases deliver Grafana Cardinality Explorer, VMUI pagination, JWT auth enhancements, log UI sidebar, and Kubernetes log enrichment for scalable observability. Upgrade now for better UI usability, query performance, and production-ready monitoring workflows."
toc: true
categories:
  - Product News
  - Company News
tags:
  - victoriametrics
  - victorialogs
  - observability
  - open source
  - kubernetes
  - new features
images:
  - /blog/victoriametrics-march-2026-ecosystem-updates/preview.webp
---

Welcome to the March release roundup of VictoriaMetrics Stack, covering key enhancements in VictoriaMetrics and VictoriaLogs. These updates deliver improved UI scalability, enhanced authentication flexibility, improved query performance, and logging tools that streamline observability workflows in production environments.

This roundup covers releases for:

- [VictoriaMetrics](#vm)
- [VictoriaLogs](#vl)

## VictoriaMetrics v1.139.0, v1.138.0, and v1.137.0 {#vm}

**VictoriaMetrics [v1.139.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11390) release highlights**

- Cardinality Explorer in Grafana: Added an experimental dashboard for [exploring metric cardinality](https://docs.victoriametrics.com/victoriametrics/#cardinality-explorer) directly in Grafana.
- Improved VMUI scalability: Added pagination for alerts and rules in [VMUI](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#vmui), making the UI easier to use with thousands of entries.
- [`vmauth`](https://docs.victoriametrics.com/victoriametrics/vmauth/) improvements: [JWT handling](https://docs.victoriametrics.com/victoriametrics/vmauth/#jwt-token-auth-proxy) is now more flexible and standards-compliant, and basic auth no longer requires a username.
- Better OpenTelemetry compatibility: Added support for negative buckets in [exponential histograms](https://opentelemetry.io/docs/specs/otel/metrics/data-model/#exponentialhistogram).
- Faster restores: [`vmrestore`](https://docs.victoriametrics.com/victoriametrics/vmrestore/) can now be significantly faster on Linux, with optimized write paths and file preallocation when the network is not the bottleneck.
- More robust cluster behavior: Improved handling of stale connections in [`vmselect`](https://docs.victoriametrics.com/victoriametrics/quick-start/#installing-vmselect), reducing query failures after `vmstorage` restarts.
- Improved NFS stability: Fixed a panic that could occur during directory deletion on NFS mounts.

As always, this release also includes bug fixes and smaller quality-of-life improvements.

**New features in VictoriaMetrics [v1.138.0](https://docs.victoriametrics.com/victoriametrics/changelog/#v11380)**

- [vmauth](https://docs.victoriametrics.com/victoriametrics/vmauth/) access logs: Added an [`access_log`](https://docs.victoriametrics.com/victoriametrics/vmauth/#access-log) option so per-user requests can be logged to stdout, with filtering by HTTP status codes for easier debugging of unusual traffic and request sources.
- JWT auth in `vmauth`: Expanded [JWT tokens features](https://docs.victoriametrics.com/victoriametrics/vmauth/#jwt-token-auth-proxy). You can now use [`match_claims`](https://docs.victoriametrics.com/victoriametrics/vmauth/#jwt-claim-matching) for JWT-based request routing, alongside support [claim-based request templating](https://docs.victoriametrics.com/victoriametrics/vmauth/#jwt-claim-based-request-templating), and [OIDC Discovery](https://docs.victoriametrics.com/victoriametrics/vmauth/#oidc-discovery) for automatically fetching and rotating public keys.
- New and updated JWT guides: Expanded the JWT flow with [guides for Grafana + VictoriaMetrics](https://docs.victoriametrics.com/guides/grafana-vmauth-openid-configuration/) datasource setups, `vmagent` [JWT token retrieval for `remoteWrite`](https://docs.victoriametrics.com/guides/vmagent-openid-configuration/), and [`vmgateway` migration to `vmauth`](https://docs.victoriametrics.com/victoriametrics/vmgateway/#access-control-migration-to-vmauth).
- Bug fixes: Also shipped several fixes, including correct routing to `default_url`, improved histogram metadata handling, and vmauth dashboard memory panel fixes.

**VictoriaMetrics [v1.137.0](https://docs.victoriametrics.com/victoriametrics/changelog/#v11370) release highlights**

> [!NOTE] Important
> `vmstorage` in VictoriaMetrics cluster has a bug affecting the `/select/multitenant/*` APIs, which may return incorrect results due to a tenant search issue.
> The affected and fixed versions are below:
>
> - v1.135.0 -> v1.137.0
> - v1.136.0 -> v1.136.1
> - v1.122.14 -> v1.122.16
> - v1.110.29 -> v1.110.31
>
> We advise skipping the affected releases if you use [multitenancy and `/select/multitenant/*`](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#multitenancy).

- Cross-backend backups with [`vmbackup`](https://docs.victoriametrics.com/victoriametrics/vmbackup/): Backups can now be copied between different storage backends, such as S3 to local disk or GCS to S3.
- JWT authentication in `vmauth`: Added [JWT authentication](https://docs.victoriametrics.com/victoriametrics/vmauth/#jwt-token-auth-proxy) with claim-based routing.
- macOS process metrics: All components now expose process-level metrics on macOS, including `process_cpu_seconds_total` and `process_resident_memory_bytes`.
- [`vmagent`](https://docs.victoriametrics.com/victoriametrics/vmagent/) default debug UI restored: `-promscrape.dropOriginalLabels` now defaults to `false`, bringing back the **Discovered targets** UI without extra CPU or memory overhead.
- Metadata ingestion enabled by default: Metrics metadata ingestion and in-memory storage are now enabled by default, and [can be disabled with `-enableMetadata=false`](https://docs.victoriametrics.com/victoriametrics/vmagent/#metric-metadata).
- New [`histogram_fraction` function](https://docs.victoriametrics.com/victoriametrics/metricsql/#histogram_fraction): Added a MetricsQL function for calculating bucket fractions between bounds, improving compatibility with SLO and SLI tooling.
- SBOM attestations for container images: Software Bill of Materials (SBOM) attestations are now published for Docker images in [SPDX format](https://spdx.dev/), improving supply chain transparency and compliance.

## VictoriaLogs v1.48.0 and v1.47.0 {#vl}

**VictoriaLogs [v1.48.0](https://docs.victoriametrics.com/victorialogs/changelog/#v1480) release highlights**

- Web UI sidebar: The built-in web UI now includes a sidebar for quick filtering by log stream fields and their values. This gives us a much faster way to narrow down noisy result sets without manually composing every filter from scratch. (Thanks to [@hagen1778](https://github.com/hagen1778) and [@Loori-R](https://github.com/loori-r) for pushing hard to implement this feature!)
    ![Screenshot of VictoriaLogs VMUI](/blog/victoriametrics-march-2026-ecosystem-updates/vl-sidebar.webp)
- `ignore_pipes` support: Supporting that new sidebar, HTTP query endpoints can now [ignore pipes when `ignore_pipes=1`](https://docs.victoriametrics.com/victorialogs/querying/#querying-hits-stats) is passed as a query argument. In practice, this means helper requests for stream fields and values can still work even if the current query contains transformation pipes, so the UI can keep providing filtering options without being confused by the later stages of the query.
- CSV export: Query results from `/select/logsql/query` can now be returned as CSV by [adding `format=csv`](https://docs.victoriametrics.com/victorialogs/querying/#querying-logs-in-csv-format) query argument. That makes it easier to export data into spreadsheets, reporting workflows, and external tools, and it also enables cleaner CSV downloads from the web UI.
- Kubernetes Namespace metadata: [`vlagent`](https://docs.victoriametrics.com/victorialogs/vlagent/) can now include Namespace labels and Namespace annotations in Kubernetes log entries. The same metadata is also available to `-kubernetesCollector.excludeFilter` even when you choose not to store it in every log entry, so unwanted logs can still be filtered out earlier in the pipeline.
- offset in `total_stats`: LogsQL now supports the offset modifier for [`first(...)` and `last(...)` inside `total_stats`](https://docs.victoriametrics.com/victorialogs/logsql/#total_stats-pipe-functions). This makes sequence analysis more practical because you are no longer limited to the earliest or latest event and can also inspect the second, third, or earlier items in the ordered set.
- `last()` in `running_stats`: [`running_stats` now supports `last(...)`](https://docs.victoriametrics.com/victorialogs/logsql/#running_stats-pipe-functions), which can be used for comparing each record with the previously seen value. This is useful for spotting changes between consecutive events, such as state transitions, previous timestamps, or earlier field values.
- `maxQueryLen` limit: Users can now limit accepted query size with `-search.maxQueryLen`. This provides a straightforward guardrail against extremely long queries that would otherwise consume extra CPU and memory during parsing and execution.
- Ingestion stream skip: During ingestion, logs containing `_stream` or `_stream_id` are now skipped, and a warning is emitted for each affected entry. This prevents user-supplied fields from colliding with internally generated stream metadata, making ingestion behavior safer and easier to debug when input data is malformed.
- SBOM attestations for container images: Container images now publish Software Bill of Materials (SBOM) attestations for Docker images in [SPDX format](https://spdx.dev/), improving supply chain transparency and compliance.

**VictoriaLogs [v1.47.0](https://docs.victoriametrics.com/victorialogs/changelog/#v1470) release highlights**

- `json_array_contains_any` filter: [New `json_array_contains_any(...)` filter](https://docs.victoriametrics.com/victorialogs/logsql/#json_array_contains_any-filter) lets you filter logs by values inside JSON arrays such as `tags=["prod","canary"]`, which makes it much easier to work with structured JSON logs without first reshaping the data.
- first/last in `total_stats`: [`total_stats` now has `first(...)` and `last(...)`](https://docs.victoriametrics.com/victorialogs/logsql/#total_stats-pipe-functions), so you can pull the earliest and latest field values in a time range and analyze event sequences more directly.
- New stats pipe functions: The [stats pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe-functions) gained `any(...)`, `field_max(...)`, and `field_min(...)`, which makes it simpler to return a sample field value or the field attached to the highest or lowest point, without the extra unpacking work older patterns needed.
- Nested `total_stats` labels: `stats_query` and `stats_query_range` [now allow `total_stats by (...)`](https://docs.victoriametrics.com/victorialogs/logsql/#total_stats-by-fields) to use a smaller set of labels than the preceding stats `by (...)`, which helps when you want to roll up grouped results into cleaner totals for dashboards, alerting, or further analysis.
- Interval selector UI: The web UI replaces the old Bars control with an Interval selector on the hits chart, so choosing chart resolution is more intuitive and better aligned with the time range you are investigating.
- Improved hits chart: The hits chart is now easier to drill into. Clicking a bar sets the exact time range for that bucket, and zooming requires a minimum drag distance, reducing accidental selections during exploration.

## Full Changelogs

Upgrade today to leverage these enhancements and explore the full changelogs for implementation details.

- VictoriaMetrics: [v1.139.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11390), [v1.138.0](https://docs.victoriametrics.com/victoriametrics/changelog/#v11380), and [v1.137.0](https://docs.victoriametrics.com/victoriametrics/changelog/#v11370)
- VictoriaLogs: [v1.48.0](https://docs.victoriametrics.com/victorialogs/changelog/#v1480), and [v1.47.0](https://docs.victoriametrics.com/victorialogs/changelog/#v1470)

Stay tuned for more ecosystem updates from the VictoriaMetrics team. Thank you for reading!
