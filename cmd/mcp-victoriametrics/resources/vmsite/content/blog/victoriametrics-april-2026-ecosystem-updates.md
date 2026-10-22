---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-05-06
enableComments: true
title: "VictoriaMetrics April 2026 Ecosystem Updates"
summary: "VictoriaMetrics April 2026 release round‑up: heads up about critical VictoriaMetrics bugfixes in v1.141.0–v1.142.0, and explore new VictoriaLogs features including Splunk ingestion and advanced LogsQL tooling."
toc: true
categories:
  - Product News
  - Company News
tags:
  - victoriametrics
  - victorialogs
  - observability
  - open source
  - new features
  - community
  - KubeCon
  - talks
images:
  - /blog/victoriametrics-april-2026-ecosystem-updates/preview.webp
---

We're excited to learn that our [vmagent](https://docs.victoriametrics.com/victoriametrics/vmagent/) helped Airbnb migrate its high-volume metrics pipeline from StatsD and Veneur to OpenTelemetry. Airbnb is now handling **100 million samples per second**. You can read more about the migration in these articles:
- The Airbnb Tech Blog: [Building a high-volume metrics pipeline with OpenTelemetry and vmagent](https://medium.com/airbnb-engineering/building-a-high-volume-metrics-pipeline-with-opentelemetry-and-vmagent-c714d6910b45)
- InfoQ: [Airbnb Migrates High-Volume Metrics Pipeline to OpenTelemetry](https://www.infoq.com/news/2026/04/airbnd-opentelemetry-vmagent/)

In other news, April saw releases across the VictoriaMetrics Observability Stack. We have released several **important bugfixes for VictoriaMetrics** and many new features in VictoriaLogs.

This release round-up covers updates for:
- [VictoriaMetrics](#vm)
- [VictoriaLogs](#vl)

## VictoriaMetrics v1.141.0 and v1.142.0 provide important bugfixes {#vm}

These releases are mostly focused on bug fixes, including two critical issues. Users should avoid affected versions or upgrade to the latest release.

### Critical issues

**OpenTelemetry ingestion issues**

There’s a bug in [OpenTelemetry ingestion](https://docs.victoriametrics.com/victoriametrics/integrations/opentelemetry/) when the [`-opentelemetry.usePrometheusNaming` flag](https://docs.victoriametrics.com/victoriametrics/integrations/opentelemetry/#label-sanitization) is enabled. Metric names can be incorrectly modified based on the previously processed metric. For example, if a metric `http_requests` has `Unit: seconds`, and the next metric has no unit (e.g., `cpu_usage`), then the latter could be ingested incorrectly using the last unit (e.g., `cpu_usage_seconds`). For more information, see [#10889](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/10889).

> [!NOTE] Bugfixes
> The bug was introduced in [v1.132.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2025/#v11320). All versions starting with v1.132.0, including LTS v1.136.0-v1.136.6, are affected. LTS v1.122.x was not affected by the bug.
>
> If you’re running any of these versions, use OpenTelemetry and have the flag set, upgrade to a fixed release:
> - [v1.142.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11420)
> - [v1.136.8](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11368)

**Incorrect binary operation precedence**

There is a known serious issue in **v1.140.0, v1.136.4, and v1.122.19** versions where binary operation precedence may be evaluated incorrectly, leading to unexpected query results. For example, expressions like `10 - (3 + 3 + 4)` are evaluated as `10 - 3 + 3 + 4`.

**Please avoid upgrading or using these versions in production**. Instead, use a fixed version such as [v1.136.8](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11368) or [v1.142.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11420).

For more details, read the [#10856 on GitHub](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/10856).

### Notable improvements in v1.141.0 and v1.142.0

- Configurable future sample retention: [Added `-futureRetention` flag](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#retention). Previously, future-dated samples were capped at 2 days; now this window is configurable. Useful for workloads that ingest forecasts or predictions.
- More accurate alert debugging in VMUI: The UI now preserves alert parameters (e.g., group `interval`, `evaluationTime`) when clicking **Run Query** on Alert view and going to Query execution tab. This makes the Query tab behave much closer to actual alert evaluation and reduces confusion.
- Metadata visibility in cluster Grafana dashboards: Metadata ingestion [is now reflected in cluster dashboard](https://github.com/VictoriaMetrics/VictoriaMetrics/blob/master/dashboards/victoriametrics-cluster.json). Since metadata processing can take up considerable resources, this information lets you spot potential issues before they impact your cluster.
- Correct tenant attribution for metadata: Metadata ingested via multitenant endpoints or scraped with `__tenant_id__` is now properly associated with the correct tenant.

Read the changelogs:
- [v1.141.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11410)
- [v1.142.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11420)

## VictoriaLogs v1.50.0 now supports Splunk Ingestion {#vl}

- Splunk HEC ingestion: VictoriaLogs supports Splunk HEC-compatible endpoints such as `/services/collector/event` and `/insert/splunk/services/collector/event`, allowing Splunk HEC clients (for example, [Docker's Splunk logging driver](https://docs.docker.com/engine/logging/drivers/splunk/)) to [send logs to VictoriaLogs with minimal integration changes](https://docs.victoriametrics.com/victorialogs/data-ingestion/splunk/).
- File‑based collection in vlagent: [vlagent](https://docs.victoriametrics.com/victorialogs/vlagent/) can collect logs from files matched by `-fileCollector.glob`, store read checkpoints, attach `file` and `hostname` fields, and support the common create/rename log rotation strategy, making it convenient to collect application logs from paths such as `/var/log/nginx/*.log` without a separate log shipper.
- Multi‑field LogsQL search: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/) now supports field‑prefix filters such as `kubernetes.*:nginx` and `*:error`, which simplifies searching across dynamic or nested field groups without listing every field name, while still allowing precise field filters when performance is important.
- Loki JSON message field prefixing: [Loki ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/#parsing-log-message) supports `-loki.messageFieldsPrefix` and the `message_fields_prefix` query arg; JSON fields extracted from the log message can be stored with a prefix such as `msg.`, reducing name collisions with stream labels or existing fields.
- Native Datadog Agent paths: VictoriaLogs accepts Datadog's default `/api/v2/logs` and `/api/v1/validate` paths directly, in addition to `/insert/datadog/…`, which simplifies [Datadog Agent and serverless Datadog setups](https://docs.victoriametrics.com/victorialogs/data-ingestion/datadog-agent/) since path rewriting through vmauth is no longer required.

[Read the changelog](https://docs.victoriametrics.com/victorialogs/changelog/#v1500)

## VictoriaLogs v1.49.0 highlights

- CSV export: Logs from `/select/logsql/query` can now be [returned in CSV format](https://docs.victoriametrics.com/victorialogs/querying/#querying-logs-in-csv-format), and the [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui) supports downloading logs as CSV for use in Excel, Google Sheets, sharing with teams, or as an attachment in tickets.
- Better `_stream` ingestion: VictoriaLogs can [ingest logs](https://docs.victoriametrics.com/victorialogs/data-ingestion/) with pre-formatted `_stream` fields, making re-import tasks easier. In addition, VictoriaLogs now validates that `_stream` matches row fields to catch incorrect metadata early.
- New multitenant native endpoint: Added `/insert/multitenant/native` for sending native-format logs with [mixed tenants from vlagent](https://docs.victoriametrics.com/victorialogs/vlagent/#multitenancy); recommended over `/internal/insert`.
- Global filter propagation: `options(global_filter=(...))` allows defining a common filter applied across [subqueries](https://docs.victoriametrics.com/victorialogs/logsql/#subqueries), reducing repetition in joins, unions, and nested queries.
- Enhanced conditional stats: Added `switch(case ..., default ...)` in [stats pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) to simplify grouped counters like 2xx, 4xx, and 5xx classifications.
- Standard deviation support: Introduced [`stddev(...)` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#stddev-stats) for analyzing variability in values such as latency or payload size.
- Static data in LogsQL: Queries can now use `join by (...) rows(...)` to [enrich logs](https://docs.victoriametrics.com/victorialogs/logsql/#enriching-logs-with-static-fields) with predefined static fields or [append static rows](https://docs.victoriametrics.com/victorialogs/logsql/#adding-static-logs) with `union rows(...)`. Handy for small lookup data, labels, or quick comparisons without preparing extra data outside VictoriaLogs.
- Field discovery improvements: `field_names` and `field_values` pipes [now support substring filtering](https://docs.victoriametrics.com/victorialogs/querying/#querying-stream-field-names), and the field/stream discovery HTTP APIs also support the filter query arg.
- Easier filtering in Web UI: In log details, you can now include or exclude a field value directly from the UI. That saves time during investigation because you do not need to manually edit the query for every small drill-down step.
- LogsQL examples in Web UI: A query example modal was added with categories, so beginners can pick a pattern and run it right away.
    ![Query example modal in Web UI](vt-modal-examples.webp)
    <figcaption style="text-align: center; font-style: italic;">New query examples modal in VictoriaLogs VMUI</figcaption>
- The hits chart was also improved: Clicking legend items now focuses a series or shows all series, while Cmd/Ctrl-click toggles hide/show. Additional actions moved to a hover menu.

[Read the changelog](https://docs.victoriametrics.com/victorialogs/changelog/#v1490)

## Changelogs

- [VictoriaMetrics v1.142.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11420)
- [VictoriaMetrics v1.141.0](https://docs.victoriametrics.com/victoriametrics/changelog/changelog_2026/#v11410)
- [VictoriaLogs v1.50.0](https://docs.victoriametrics.com/victorialogs/changelog/#v1500)
- [VictoriaLogs v1.49.0](https://docs.victoriametrics.com/victorialogs/changelog/#v1490)

## Community Highlights

During April, VictoriaMetrics participated in several community events:

- KubeCon + CloudNativeCon Europe Amsterdam: these events brought together 13,500 attendees, the largest turnout yet. We were happy to talk with our users at our booth, present talks, and participate in panels. 
  - We also hosted our very first after party: "After Deploy". Check out the talks and pictures [on our blog](https://victoriametrics.com/blog/victoriametrics-at-kubecon-amsterdam/).
  - Roman's and Jiekun's talk at KubeCon [on cutting 90% Distributed Tracing Bandwidth on OpenTelemetry](https://www.youtube.com/watch?v=ehKzt_kGEeM) has been published as a blog post. Read it here: [Optimizing Tail Sampling in OpenTelemetry with Retroactive Sampling](https://victoriametrics.com/blog/kubecon-eu-2026-sampling/).
- VictoriaMetrics Virtual Meetup: this year's Virtual Meetup took place in April. We presented roadmap updates for Q1, news on Anomaly Detection, updates on VictoriaMetrics Cloud and VictoriaLogs, and even had time for an AMA. [Watch the recording on YouTube](https://www.youtube.com/watch?v=v5VT0hthH7I).
- LinuxFest Northwest (April 24–26, Bellingham): As Gold sponsor, delivered talks on using VictoriaMetrics for observability in SIEM systems, homelab monitoring setups, intelligent routing optimizations, and cost-optimized Kubernetes observability
- StackConf (April 28–29, Munich): VictoriaMetrics was a Bronze sponsor featuring sessions on aggregating metrics in-flight to reduce cardinality and building full OSS observability stacks combining VictoriaMetrics with OpenTelemetry Collector.

That's a wrap for April. Here's to even more developments and exciting events for May. Thank you for reading!
