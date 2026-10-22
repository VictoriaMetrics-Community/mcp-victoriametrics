---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-10-12
enableComments: true
title: "VictoriaMetrics September 2026 Ecosystem Updates"
summary: "September release roundup brings new releases across the VictoriaMetrics ecosystem, plus the launch of VictoriaTraces in Cloud, completing our managed observability offering of metrics, logs, and traces stack."
toc: true
categories:
  - Product News
  - Company News
tags:
  - victoriametrics
  - vmestimator
  - victorialogs
  - victoriatraces
  - vmanomaly
  - observability
  - new features
images:
  - /blog/victoriametrics-september-2026-ecosystem-updates/preview-1.webp
---

Another month whooshed by, bringing a slew of updates: 3 vmanomaly and 2 VictoriaMetrics releases, and new versions of vmestimator, VictoriaTraces, and our first ever LTS release of VictoriaLogs. This brings 12 months of stability to VictoriaLogs Enterprise users.

But the big one was the [launch of VictoriaTraces on VictoriaMetrics Cloud](https://victoriametrics.com/blog/announcing-victoriatraces-in-cloud/). Finally, we can offer a complete managed observability stack with metrics, logs, and traces. We run it, you use it.

This ecosystem update includes:
- Releases of VictoriaMetrics [1.153.0](#vm11530) and [1.152.0](#vm11520)
- Updates for [vmestimator 0.1.16](#vmestimator), [VictoriaLogs 1.53.0](#vl1530), [VictoriaLogs LTS 1.52.1](#vl1521), and [VictoriaTraces 0.12.0](#vt0120).
- Releases of vmanomaly [1.30.7](#vmanomaly1307), [1.30.6](#vmanomaly1306), [1.30.5](#vmanomaly1305).

## VictoriaMetrics 1.153.0 is out: OIDC single sign-on for vmauth {#vm11530}

> **Caution:** A [known bug](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/10278#issuecomment-5862232663) in vmauth SSO was [fixed](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/11656) after the release cutoff. We recommend waiting for the next release before using SSO in production.

- New OIDC-based single sign-on on vmauth: [vmauth](https://docs.victoriametrics.com/victoriametrics/vmauth/) now supports [Single sign-on (SSO)](https://docs.victoriametrics.com/victoriametrics/vmauth/#single-sign-on-sso) via OpenID Connect.
- Skip pending state on vmalert restarts: [vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/) now moves alerts straight to firing when their state can be [restored after a restart](https://docs.victoriametrics.com/victoriametrics/vmalert/#alerts-state-on-restarts), instead of going through an unnecessary pending phase first.
- Improved page cache efficiency in storage: [vmsingle](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/) and vmstorage now support opt-in `FADV_RANDOM` and `MADV_RANDOM` hints for data part files. We *believe* this setting can improve query performance. If you are interested in testing, enable the `-fs.disableAdviseRandomRead=false` flag and, should you find it helpful, please share your results and feedback in [issue #11461](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/11461). Thank you!
- HTTP headers in vmauth access logs: vmauth can now print selected HTTP headers in its [access logs](https://docs.victoriametrics.com/victoriametrics/vmauth/#access-log). Only headers listed in the `access_log.headers` setting are logged, which makes it easier to troubleshoot pipelines that rely on HTTP headers.
- Includes 11 bug fixes and 3 security fixes

[Read the changelog](https://docs.victoriametrics.com/victoriametrics/changelog/#v11530)

## VictoriaMetrics 1.152.0 is out: Accurate unused metrics reporting in cluster {#vm11520}

- Unused metrics reporting fix: In [VictoriaMetrics cluster](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/), vmselect now applies the `le` filter on the `/api/v1/status/metric_names_stats` endpoint only after merging request counters from all vmstorage nodes. Previously, the endpoint could incorrectly report actively queried metrics as unused when a node's local request counter was zero. Thanks to [@missusk](https://github.com/missusk) for their contribution.

[Read the changelog](https://docs.victoriametrics.com/victoriametrics/changelog/#v11520)

## vmestimator 0.1.16 is out: Churn ratio estimation {#vmestimator}

- New churn ratio metric: vmestimator can now estimate [churn](https://docs.victoriametrics.com/victoriametrics/vmestimator/#churn-calculation) as well as cardinality. Set the new `churn_interval` option on a [stream](https://docs.victoriametrics.com/victoriametrics/vmestimator/#configuration), and vmestimator will expose a `cardinality_churn_ratio` metric per group, with values from `0` (a stable series set) to `1` (the series were completely replaced). This makes it easier to find jobs that put [high churn](https://docs.victoriametrics.com/victoriametrics/faq/#what-is-high-churn-rate) indexing pressure on storage.
- Improves the Cardinality Estimations dashboard: The [VictoriaMetrics - Cardinality Estimations](https://play-grafana.victoriametrics.com/d/mktd5g9/victoriametrics-cardinality-estimations) is available in our Grafana playground and now shows churn alongside cardinality estimates. See the [dashboards docs](https://docs.victoriametrics.com/victoriametrics/vmestimator/#dashboards) for details.

![Screenshot of the vmestimator Global Cardinality Trend and Churn Rate panels on Grafana](/blog/victoriametrics-september-2026-ecosystem-updates/vmestimator.webp)
<figcaption style="text-align: center; font-style: italic;">vmestimator can now estimate churn rate.</figcaption>

[Read the changelog](https://github.com/VictoriaMetrics/vmestimator/blob/main/docs/vmestimator/CHANGELOG.md#v0116)

## VictoriaLogs 1.53.0 is out: Faster latest logs queries in cluster mode {#vl1530}

> **Caution:** If you run v1.52.0, we recommend upgrading since that version could panic or drop logs with duplicate [stream field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) names at ingestion, and panic when querying such logs written by earlier releases.
>
> **Endpoint changes**: This release has two security-related changes that may break existing scripts or automation.
> - The `/internal/force_merge`, `/internal/force_flush`, `/internal/log_new_streams`, and `/internal/partition/*` HTTP endpoints now accept only the POST method Update any scripts that call these endpoints with GET.
> - Requests to `/select/vmalert/*` path ending with `/config` or `/reload` now require the `-httpAuth.*` credentials.

- Faster **latest logs** queries in cluster mode: Queries that return the newest N logs, which the [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui) runs by default, now run on each `vlstorage` node instead of going back and forth through `vlselect`. [Cluster](https://docs.victoriametrics.com/victorialogs/cluster/) users get noticeably faster results on large time ranges.
- Adds separate key for log deletion: The new `-deleteAuthKey` flag protects the `/delete/*` endpoints of the [deletion API](https://docs.victoriametrics.com/victorialogs/#how-to-delete-logs) with their own key, independent of `-httpAuth.*`.
- Wait time in slow query logs: Each slow query log line emitted according to `-search.logSlowQueryDuration` now includes `waitDuration`. It shows whether a query waited in the queue for a free concurrency slot or was expensive to run, so you know whether to add capacity or improve the query.
- Added slow query logs on vlstorage nodes: `-search.logSlowQueryDuration` now also applies to vlstorage nodes. In a [cluster](https://docs.victoriametrics.com/victorialogs/cluster/), you can see which storage node was slow, not only that the overall query was slow.
- Added listener for Unix domain sockets: VictoriaLogs can now accept HTTP requests on a local socket with `-httpListenAddr=unix:/path/to/socket`. Sidecars and reverse proxies on the same host can reach it without a TCP port, and file permissions control access. See the [Unix domain socket docs](https://docs.victoriametrics.com/victorialogs/security-and-lb/#unix-domain-socket).
- New fsync latency panel: The [single-node](https://grafana.com/grafana/dashboards/22084), [cluster](https://grafana.com/grafana/dashboards/23274), and [vlagent](https://grafana.com/grafana/dashboards/24513) Grafana dashboards have a new `Fsync avg duration` panel in the Troubleshooting section. Slow disks that delay writes to storage show up there before they cause ingestion problems.
- New compression ratio over time panel: The single-node and cluster dashboards now show how the compression ratio changes over time. You can see how changes such as a new [stream field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) or log parsing at ingestion time affect disk usage.
- Logchef integration: You can now use [Logchef](https://docs.victoriametrics.com/victorialogs/integrations/logchef/) as another UI to query and explore your VictoriaLogs data.
- Web UI Improvements
  - Gradual loading of the hits chart: For slow queries, the [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui) loads the hits chart step by step and shows a progress indicator. You can narrow the time range as soon as the first data appears instead of waiting for the whole query to finish.
  - Better stream fields sidebar: The `Stream fields` sidebar now lets you search field names and values, sort by hits or names, and pick an `Any value` option. Selected items are pinned to the top. Finding the right [stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) in large setups is much quicker.
  - Quicker field actions: In an expanded log entry, `Exclude` is now right next to field values, so narrowing down noisy logs takes fewer clicks. Stream fields are also visually marked, and include/exclude actions on them automatically use [stream filters](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter).
  - Field names in the Group view: Each value now has its field name in front of it (except `_msg`), so rows with many fields are easier to read.
  - Back button on the hits chart: After you zoom or pan the hits chart, one click takes you back to the previous time range.
  - Auto-refresh moved to the Execute button: The auto-refresh settings are now in the Execute button dropdown, next to the action they affect. This frees up space in the header.
  - Custom favicon color: In the settings, you can pick a color for the browser tab icon. This makes it easy to tell prod, staging, and test tabs apart.
  - Cleaner hits chart controls: The `Stacked` toggle is now hidden when `Group by` is set to `none`, because it has no effect there.

[Read the changelog](https://docs.victoriametrics.com/victorialogs/changelog/#v1530)

## VictoriaLogs v1.52.1: the first LTS release for VictoriaLogs {#vl1521}

This is our **first LTS release for VictoriaLogs**, bringing Enterprise users 12 months of support.

> - If you run v1.52.0, we recommend upgrading since that version could panic or drop logs with duplicate [stream field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) names at ingestion, and panic when querying such logs written by earlier releases.
> - The misspelled `vl_udp_reqests_total` metric is renamed to [`vl_udp_requests_total`](https://docs.victoriametrics.com/victorialogs/metrics/#vl_udp_requests_total). Update any dashboards or alerts that use the old name.

Here are the highlights:

- Safer log deletion: The `/delete/run_task` endpoint of the [deletion API](https://docs.victoriametrics.com/victorialogs/#how-to-delete-logs) now accepts only `POST` requests, which blocks some [SSRF](https://en.wikipedia.org/wiki/Server-side_request_forgery)-based deletion attacks. The Go builder is also upgraded to Go 1.26.8 for upstream security fixes.
- No more lost logs on vlinsert shutdown: In [cluster](https://docs.victoriametrics.com/victorialogs/cluster/) mode, `vlinsert` now drains buffered logs to `vlstorage` nodes on graceful shutdown instead of dropping them. The new `-insert.drainTimeout` flag (default `5s`) limits how long it waits.
- Even load when storage nodes are down: Data rerouted away from unavailable `vlstorage` nodes is now spread evenly across healthy nodes. Before, nodes next to the failed one in the `-storageNode` list could get much more data than the others.
- Canceled queries actually stop: When a query is canceled or times out on `vlselect`, the matching work on `vlstorage` is now stopped too, instead of wasting CPU and disk I/O.
- Correct ordering for "latest logs" queries: Queries such as `| sort by (_time desc) limit N` now return correct results when other pipes rewrite `_time` (for example `math`, `replace`, `extract`, or `unpack_json`). Before, they could fail with a `missing _time field` error or return logs out of order.
- vlagent file collection fixes: The [File Collector](https://docs.victoriametrics.com/victorialogs/vlagent/#collect-logs-from-files) no longer panics when a log file is deleted, and it no longer fails at startup when it can't read neighboring files. It sets the `file` field correctly with `-fileCollector.extraFields`. Incomplete last lines are dropped with a warning instead of being joined to the first line of the rotated file.

[Read the release notes](https://github.com/VictoriaMetrics/VictoriaLogs/releases/tag/v1.52.1)

## VictoriaTraces 0.12.0 is out: A new built-in UI for trace exploration {#vt0120}

- This release introduces a new built-in UI (VTUI) for exploring traces. You can now find a trace, drill into its spans, and see what’s slowing down your application without any dependency on an external visualization platform.

    Try it out yourself in the playground: [https://play-vtraces.victoriametrics.com/select/vmui/](https://play-vtraces.victoriametrics.com/select/vmui/)

- This release also ships 2 bug fixes and 2 security changes.

![Screenshot of VictoriaTraces UI](/blog/victoriametrics-september-2026-ecosystem-updates/vtui-animation.webp)
<figcaption style="text-align: center; font-style: italic;">VictoriaTraces now ships with built-in Trace Explorer UI</figcaption>

[Read the changelog](https://docs.victoriametrics.com/victoriatraces/changelog/#v0120)

## vmanomaly 1.30.7 is out: Jump from a firing alert to its anomaly results {#vmanomaly1307}

> Before a stateful upgrade or rollback, check the [directional compatibility matrix](https://docs.victoriametrics.com/anomaly-detection/migration/#compatibility-matrix) and run [`--dryRun`](https://docs.victoriametrics.com/anomaly-detection/migration/#dry-run) to preview how persisted model state and on-disk artifacts will be migrated.

- Go from a firing alert to its data: Add an [investigation link](https://docs.victoriametrics.com/anomaly-detection/ui/#investigate-a-firing-alert) to your [vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/) rules to open persisted anomaly results in the UI, with the configured model, queries, and incident time range already selected.
- Simplified model choices in the UI: Use the [`server.ui_allowed_models`](https://docs.victoriametrics.com/anomaly-detection/components/server/#parameters) parameter to choose which models appear in the UI model selectors and Auto tuning targets. This only curates what the UI shows; it doesn't restrict which models can be executed through the API.
- Activity and AI usage at a glance: Revision 10 of the [self-monitoring dashboard](https://docs.victoriametrics.com/anomaly-detection/self-monitoring/#grafana-dashboard) ([Grafana ID 22337](https://grafana.com/grafana/dashboards/22337)) adds successful-run counts, corrected queries, per-model views and [AI Copilot](https://docs.victoriametrics.com/anomaly-detection/ui/#ai-assistance) usage by instance, including tokens and estimated cost. Missing telemetry is shown as Unknown, and cost estimates exclude unpriced responses and tool/server compute.
- Reduced reader overhead: The [reader](https://docs.victoriametrics.com/anomaly-detection/components/reader/) does less work when parsing datapoints, preparing columns, and loading single-chunk Parquet data.
- More accurate alert rules: Alert rules generated from the UI now use a strict anomaly score threshold (`> 1` by default instead of `>= 1`), so a score exactly at the threshold no longer fires. [Preset](https://docs.victoriametrics.com/anomaly-detection/presets/) version suffixes are now preserved in metric selectors.
- Updated documentation for AI-assisted investigations: [MCP server v0.4.2](https://github.com/VictoriaMetrics/mcp-vmanomaly/releases/tag/v0.4.2) bundles the v1.30.7 documentation and adds guidance for investigating saved incident results with AI assistants.

[Read the changelog](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1307)

## vmanomaly 1.30.6 is out: Searchable server queries in the UI and experimental hardened images {#vmanomaly1306}

- Find and manage server queries more easily: The bundled [vmanomaly UI](https://docs.victoriametrics.com/anomaly-detection/ui/) is updated from [v1.9.0](https://docs.victoriametrics.com/anomaly-detection/ui/#v190) to [v1.9.1](https://docs.victoriametrics.com/anomaly-detection/ui/#v191). The Server tab in the Queries sheet now offers a searchable list with model filters, model badges, and consistent actions for individual queries or selected queries. Selected queries stay selected even when the active filter hides them, and the UI help lists keyboard shortcuts.
- Try the experimental hardened image: [Docker Hardened Images](https://docs.victoriametrics.com/anomaly-detection/quickstart/#experimental-hardened-image) are now available under the `v1.30.6-dhi` and `v1.30.6-enterprise-dhi` tags on Docker Hub and Quay, for both amd64 and arm64. Both tags point to the same image, which ships with fewer OS packages and no runtime shell or package manager. Standard tags, including `latest`, keep their existing base image. Thanks to Zakhar and Andrew Chubatiuk for their contributions.

> The hardened image runs as a non-root user and has no runtime shell or package manager. Before switching, check the permissions of mounted files (such as configs and license files) and make sure the container user can write to the state directory used for [state restoration](https://docs.victoriametrics.com/anomaly-detection/components/settings/#state-restoration).

- More stable estimates for intermittent data: Univariate [Temporal Envelope](https://docs.victoriametrics.com/anomaly-detection/components/models/#temporal-envelope) models no longer produce implausible current-time estimates on intermittent data, in both batch and streaming learning. In such cases, predictions fall back to the learned median when valid observations are available. Explicit future forecasts keep their extrapolation behavior, and the existing model state remains loadable.
- Consistent scheduling across containers: [Periodic schedulers](https://docs.victoriametrics.com/anomaly-detection/components/scheduler/#periodic-scheduler) now consistently use their configured timezone (UTC by default) for interval triggers and for ISO `start_from` values without an offset. ISO `start_from` values with an explicit offset keep their exact instant, so schedules no longer depend on the container's local timezone.
- Updated documentation for AI-assisted investigations: The [MCP server v0.4.1](https://github.com/VictoriaMetrics/mcp-vmanomaly/releases/tag/v0.4.1) now includes the v1.30.6 documentation, so AI assistants get up-to-date guidance when helping you configure and investigate vmanomaly.

[Read the changelog](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1306)

## vmanomaly 1.30.5 is out: Multivariate investigation and per-query business policies in the UI {#vmanomaly1305}

> Are you running per-query business settings on v1.30.2-v1.30.4 with `settings.n_workers` > 1? **Please upgrade to this version or greater.**
> This release fixes how query-level data ranges, detection directions, and minimum deviations are applied in [multiprocessing mode](https://docs.victoriametrics.com/anomaly-detection/components/settings/#parallelization). Explicit query settings now consistently take precedence over model defaults.

- Fixes per-query business settings in multiprocessing mode: Query-level [`data_range`, `detection_direction`, `min_dev_from_expected`, and `min_rel_dev_from_expected`](https://docs.victoriametrics.com/anomaly-detection/components/reader/#per-query-parameters) are now applied consistently when `settings.n_workers` > 1. Explicit query policies take precedence over model-level fallbacks, and observations outside the configured range keep their [out-of-range anomaly score](https://docs.victoriametrics.com/anomaly-detection/components/settings/#anomaly-score-outside-data-range).
- Investigate related signals together: The bundled [UI v1.9.0](https://docs.victoriametrics.com/anomaly-detection/ui/#v190) adds an **experimental** [multivariate investigation workspace](https://docs.victoriametrics.com/anomaly-detection/ui/#multivariate-investigation) for [multivariate models](https://docs.victoriametrics.com/anomaly-detection/components/models/#multivariate-models). You can compare aligned signals, inspect individual channels, and see a joint anomaly score for each group, for example, CPU and memory grouped by service or instance.

    ![Screenshot of Multivariate Investigation for a service](/blog/victoriametrics-september-2026-ecosystem-updates/vmanomaly-ui-multivariate-investigation.webp)
    <figcaption style="text-align: center; font-style: italic;">New multivariate investigation UI.</figcaption>

- Keep queries and business intent together: You can now name your queries, manage reusable query sets, and give each query its own data range, detection direction, and minimum deviations. Inheritance controls make it clear which values come from model defaults. See the [query explorer](https://docs.victoriametrics.com/anomaly-detection/ui/#query-explorer) docs, or try the new UX in the [playgrounds](https://docs.victoriametrics.com/anomaly-detection/ui/#playgrounds).

    ![Screenshot of improved UX: naming a query and setting parameters](/blog/victoriametrics-september-2026-ecosystem-updates/vmanomaly-ui-sections-explore.webp)
    <figcaption style="text-align: center; font-style: italic;">New UI for managing and saving queries.</figcaption>

- Reduce manual configuration with AI Copilot: [AI Copilot](https://docs.victoriametrics.com/anomaly-detection/ui/#ai-assistance) can now propose complete named-query sets, along with their policies, for you to review and approve.
- Shared autotune across queries: [Shared autotune](https://docs.victoriametrics.com/anomaly-detection/components/server/#time-series-analysis-and-autotune-api) accepts named queries and evaluates multiple queries and aligned multivariate groups in a single study, returning one shared model configuration.
- Improved visibility into AI usage: Copilot now shows reported tokens, the model used, and approximate costs where pricing is available. Duplicate context is reduced, context and output budgets are configurable, and new [AI Copilot self-monitoring metrics](https://docs.victoriametrics.com/anomaly-detection/components/monitoring/#ai-copilot-metrics) track request size, reported tokens, finish reasons, and local budget refusals.
- Ships with companion tooling updates: The [vmanomaly MCP server v0.4.0](https://github.com/VictoriaMetrics/mcp-vmanomaly/releases/tag/v0.4.0) and the [vmanomaly skills plugin v1.1.0](https://github.com/VictoriaMetrics/skills/tree/main/plugins/vmanomaly) are released alongside this version.

[Read the changelog](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1305)

## Until next month

A busy month, but a good one: sharper tooling, faster queries, easier investigations, and traces now available in the cloud. We hope something in this update makes your observability setup a little smoother. Until next time.

