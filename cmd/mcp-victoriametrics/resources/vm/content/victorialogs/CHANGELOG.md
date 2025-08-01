---
weight: 7
title: CHANGELOG
menu:
  docs:
    identifier: "victorialogs-changelog"
    parent: "victorialogs"
    weight: 7
    title: CHANGELOG
tags:
  - logs
aliases:
- /victorialogs/CHANGELOG.html
---

The following `tip` changes can be tested by building VictoriaLogs from the latest commit of [VictoriaLogs](https://github.com/VictoriaMetrics/VictoriaLogs/) repository
according to [these docs](https://docs.victoriametrics.com/victorialogs/quickstart/#building-from-source-code)

## tip

* SECURITY: upgrade base docker image (Alpine) from 3.22.0 to 3.22.1. See [Alpine 3.22.1 release notes](https://www.alpinelinux.org/posts/Alpine-3.19.8-3.20.7-3.21.4-3.22.1-released.html).

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix broken "Collapse all" button in Group view. See [#509](https://github.com/VictoriaMetrics/VictoriaLogs/issues/509). The bug has been introduced in [v1.26.0](https://github.com/VictoriaMetrics/VictoriaLogs/releases/tag/v1.26.0).
* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): prevent from possible crash when ingesting logs for dates, which are concurrently removed because of [the configured retention](https://docs.victoriametrics.com/victorialogs/#retention). See [#505](https://github.com/VictoriaMetrics/VictoriaLogs/issues/505).

## [v1.26.0](https://github.com/VictoriaMetrics/VictoriaLogs/releases/tag/v1.26.0)

Released at 2025-07-18

* FEATURE: [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/): add ability to configure auth options and TLS options for connections to the `-datasource.url`. See [auth options docs](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/#auth-options) and [TLS options docs](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/#tls-options). See [this feature request](https://github.com/VictoriaMetrics/VictoriaLogs/issues/54). Thanks to @thom-vend for [the initial pull request](https://github.com/VictoriaMetrics/VictoriaLogs/pull/457).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add the ability to hide the logs panel to view only the graph. When the logs panel is hidden, the `/query` request is not executed.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): improve autocomplete functionality with enhanced quick autocomplete via hotkey support and removed special characters from autocomplete suggestions. See [this comment](https://github.com/VictoriaMetrics/VictoriaLogs/issues/70#issuecomment-3043591443) for details.

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): prevent groups from automatically expanding on list updates if all groups were previously collapsed. See [#92](https://github.com/VictoriaMetrics/VictoriaLogs/issues/92).
* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): restore logging of too long ingested lines in order to simplify debugging of such cases. See [#430](https://github.com/VictoriaMetrics/VictoriaLogs/issues/430). The regression has been introduced in [v1.24.0-victorialogs](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.24.0-victorialogs).
* BUGFIX: properly persist newly created data on unclean shutdown such as power off, hardware crash or operating system crash. See [#505](https://github.com/VictoriaMetrics/VictoriaLogs/issues/505).

## [v1.25.1](https://github.com/VictoriaMetrics/VictoriaLogs/releases/tag/v1.25.1)

Released at 2025-07-14

* SECURITY: upgrade Go builder from Go1.24.4 to Go1.24.5. See [the list of issues addressed in Go1.24.5](https://github.com/golang/go/issues?q=milestone%3AGo1.24.5+label%3ACherryPickApproved).

* FEATURE: [VictoriaLogs cluster](https://docs.victoriametrics.com/victorialogs/cluster/): expose `vl_insert_remote_send_errors_total`, `vl_select_remote_send_errors_total` and `vl_insert_remote_is_reachable` metrics at the [`/metrics` page](https://docs.victoriametrics.com/victorialogs/#monitoring) for monitoring failed remote insert/select attempts per `vlstorage` node and detecting temporarily unavailable `vlstorage` nodes.

* BUGFIX: properly return VictoriaLogs version at `./victoria-logs -version` and at the `vm_app_version` metric exposed via [`/metrics` page](https://docs.victoriametrics.com/victorialogs/#monitoring). See [#409](https://github.com/VictoriaMetrics/VictoriaLogs/issues/409). The bug has been introduced in [v1.25.0](https://github.com/VictoriaMetrics/VictoriaLogs/releases/tag/v1.25.0).

## [v1.25.0](https://github.com/VictoriaMetrics/VictoriaLogs/releases/tag/v1.25.0)

Released at 2025-07-07

**VictoriaLogs source code has been moved from [VictoriaMetrics repository](https://github.com/VictoriaMetrics/VictoriaMetrics/) to its own [github.com/VictoriaMetrics/VictoriaLogs](https://github.com/VictoriaMetrics/VictoriaLogs/) repository. All the future development will be performed in the new repository. The `-victorialogs` suffix is removed from the release tag and from docker image tags starting from v1.25.0 release**

* FEATURE: [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe): add support for Unix timestamps in seconds, milliseconds and microseconds additionally to nanoseconds when using `<time:field_name>` formatting. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8659).

* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): properly parse fractional Unix timestamps with millisecond, microsecond and nanosecond precision in the ingested logs. Previously the precision loss in the parsed timestamps could occur. See [this comment](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/8767#discussion_r2051657518) for details.
* BUGFIX: [`rate_sum` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#rate_sum-stats): fix inconsistent per-second rate calculation when time filters are specified via HTTP query parameters instead of LogsQL expression. This affects recording rule results. See [#9303](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9303).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): disabled opening of autocomplete popup on initial page load.

## [v1.24.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.24.0-victorialogs)

Released at 2025-06-20

* FEATURE: add `-http.disableKeepAlive` to disable HTTP keep-alives for incoming connections. The flag could improve load balancing among replicas behind HTTP load balancers. See [#9125](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/9125) and [#2395](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/2395) for details.
* FEATURE: [`delete` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#delete-pipe): allow deleting all the fields with common prefix via `... | delete prefix*` syntax.
* FEATURE: [`fields` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#fields-pipe): allow keeping all the fields with common prefix via `... | fields prefix*` syntax.
* FEATURE: [`copy` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#copy-pipe): allow copying all the fields with common prefix to fields with another common prefix via `... | copy old_prefix* as new_prefix*` syntax.
* FEATURE: [`rename` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#rename-pipe): allow renaming all the fields with common prefix to fields with another common prefix via `... | rename old_prefix* as new_prefix*` syntax.
* FEATURE: [`unpack_json` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_json-pipe): allow unpacking JSON fields with common prefix via `... fields (prefix*)` syntax.
* FEATURE: [`unpack_logfmt` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_logfmt-pipe): allow unpacking JSON fields with common prefix via `... fields (prefix*)` syntax.
* FEATURE: [`avg` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#avg-stats): allow calculating the average value over all the fields with common prefix via `avg(prefix*)` syntax.
* FEATURE: [`max` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#max-stats): allow calculating the maximum value over all the fields with common prefix via `max(prefix*)` syntax.
* FEATURE: [`min` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#min-stats): allow calculating the minimum value over all the fields with common prefix via `min(prefix*)` syntax.
* FEATURE: [`median` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#median-stats): allow calculating the median value over all the fields with common prefix via `median(prefix*)` syntax.
* FEATURE: [`quantile` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#quantile-stats): allow calculating the maximum value over all the fields with common prefix via `quantile(prefix*)` syntax.
* FEATURE: [`sum` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#sum-stats): allow calculating the sum for all the fields with common prefix via `sum(prefix*)` syntax.
* FEATURE: [`sum_len` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#sum_len-stats): allow calculating the sum of byte lengths for all the fields with common prefix via `sum_len(prefix*)` syntax.
* FEATURE: [`count` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#count-stats): allow calculating the number of logs with at least a single non-empty field across fields with common prefix via `count(prefix*)` syntax.
* FEATURE: [`count_empty` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#count_empty-stats): allow calculating the number of logs with empty fields with common prefix via `count_empty(prefix*)` syntax.
* FEATURE: [`rate_sum` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#avg-stats): allow calculating the per-second rate over the sum of all the fields with common prefix via `rate_sum(prefix*)` syntax.
* FEATURE: [`row_any` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#avg-stats): allow returning all the fields with common prefix via `row_any(prefix*)` syntax.
* FEATURE: [`row_max` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#avg-stats): allow returning all the fields with common prefix via `row_max(max_field, prefix*)` syntax.
* FEATURE: [`row_min` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#avg-stats): allow returning all the fields with common prefix via `row_min(min_field, prefix*)` syntax.
* FEATURE: [`uniq_values` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#uniq_values-stats): allow fetching unique values for all the fields with common prefix via `uniq_values(prefix*)` syntax.
* FEATURE: [`values` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#values-stats): allow fetching values for all the fields with common prefix via `values(prefix*)` syntax.
* FEATURE: [`json_values` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#json_values-stats): allow fetching values for all the fields with common prefix via `json_values(prefix*)` syntax.
* FEATURE: [`-insert.maxLineSizeBytes`](https://docs.victoriametrics.com/victorialogs/faq/#what-length-a-log-record-is-expected-to-have): add logging of the number of bytes skipped for oversize lines.
* FEATURE: add `-insert.disable` and `-select.disable` command-line flags for disabling both public and internal HTTP endpoints (`/insert/*` + `/internal/insert` and `/select/*` + `/internal/select/*` respectively). See [#9061](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9061).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): enhance autocomplete with parsed field suggestions from unpack pipe. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8806).
* FEATURE: [Journald data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/journald/): automatically add `level` log field according to [the `PRIORITY` field value](https://wiki.archlinux.org/title/Systemd/Journal#Priority_level). This enables [log level highlighting in Grafana](https://grafana.com/docs/grafana/latest/explore/logs-integration/#log-level). See [#8535](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8535).
* FEATURE: [Syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): automatically add `level` log field according to [the `severity` field value](https://en.wikipedia.org/wiki/Syslog#Severity_level). This enables [log level highlighting in Grafana](https://grafana.com/docs/grafana/latest/explore/logs-integration/#log-level).
* FEATURE: [Journald data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/journald/): use `(_MACHINE_ID, _HOSTNAME, _SYSTEMD_UNIT)` fields as [log stream fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) by default. See [#9143](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9143).
* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): optimize `(any_filter or *)` filters to `*`. This avoids executing the `any_filter`. Such filters are frequently generated by Grafana.
* FEATURE: [`/select/logsql/query` endpoint](https://docs.victoriametrics.com/victorialogs/querying/#querying-logs): optimize the input query after adding the `limit` to it. This improves performance and reduces memory usage for queries ending with [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe). See [#9200](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9200). Thanks to @vadimalekseev for [the fix](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/9201).

* BUGFIX: [query API](https://docs.victoriametrics.com/victorialogs/querying/#querying-logs): properly set storage node authorization in cluster mode when [Basic Auth](https://docs.victoriametrics.com/victorialogs/cluster/#security) is enabled. See [#9080](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9080).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): live tailing tab automatically reconnects when the connection is lost. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9129).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix issue with hits chart ignoring selected AccountID and ProjectID. See [#9157](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9157).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix missing field values in auto-complete. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8749)
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): remove the compact mode of the table tab and add field sorting capabilities to the JSON tab. See [#7047](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7047).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix errors in console about loading of `manifest.json` when accessing UI through vmauth with Basic Auth enabled.
* BUGFIX: [Journald data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/journald/): properly read log timestamp from `__REALTIME_TIMESTAMP` field according to [the docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/journald/#time-field). See [#9144](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9144). The bug has been introduced in [v1.22.0-victorialogs](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.22.0-victorialogs).
* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): support `-` as a timestamp value, as described in [RFC5424](https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.3).
* BUGFIX: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): properly handle quotes inside quoted strings such as `"\""`. Previously this could lead to panics. See [#9219](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/9219).
* BUGFIX: [LogsQL regexp filter](https://docs.victoriametrics.com/victorialogs/logsql/#regexp-filter): properly parse unquoted filter ending with `*`, such as `foo:~bar.*`. It must be parsed as `foo:~"bar.*"`, while previously it was incorrectly parsed as `foo:~"bar."`.
* BUGFIX: [Journald data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/journald/): properly read large Journald requests in streaming manner. See [#9070](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9070). Deprecate `-journald.maxRequestSize` command-line flag, since it is no longer used.

## [v1.23.3](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.23.3-victorialogs)

Released at 2025-06-02

* BUGFIX: [live tailing API](https://docs.victoriametrics.com/victorialogs/querying/#live-tailing): properly return live tailing results. Previously some of these results could be missing, while others could be returned out of order (e.g. improperly sorted by [`_time` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field)). The issue has been introduced in [v1.18.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.18.0-victorialogs).
* BUGFIX: [query API](https://docs.victoriametrics.com/victorialogs/querying/#querying-logs): properly return the last `limit` logs on the selected time range if the `limit` query arg is passed to `/select/logsql/query`. Previously logs could be returned without proper sorting because of the bug, which has been introduced in [v1.18.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.18.0-victorialogs).
* BUGFIX: [querying HTTP APIs](https://docs.victoriametrics.com/victorialogs/querying/#http-api): properly drop LogsQL pipes from the provided `query` before obtaining field names and values. This is needed in order to properly implement auto-suggestion for log field names and values. See [#9068](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9068#issuecomment-2931275012).

## [v1.23.2](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.23.2-victorialogs)

Released at 2025-05-30

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): correctly handle `sort` pipe in queries — UI now respects the server-defined sort order instead of always sorting by time. See [#8660](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8660).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add alphabetical sorting for record fields in selectors and table view. See [#8438](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8438).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix an issue where queries were not triggered when relative time was selected and the chart was hidden. See [#8983](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8983).

## [v1.23.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.23.1-victorialogs)

Released at 2025-05-29

* BUGFIX: [multi-level cluster setup](https://docs.victoriametrics.com/victorialogs/cluster/#multi-level-cluster-setup): properly calculate [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) functions when a `vlselect` node queries other `vlselect` nodes, which, in turn, query `vlstorage` nodes. Previously such setup resulted in the `unexpected non-empty tail left` error for all the queries with `stats` pipe (explicit or implicit). See [#8815](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8815).

## [v1.23.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.23.0-victorialogs)

Released at 2025-05-28

* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add "Live" tab that allows monitoring logs in real-time as they arrive. This feature helps users to observe the most recent log entries without manual refreshing, making troubleshooting and monitoring more efficient. See [#7046](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7046).
* FEATURE: [dashboards/cluster](https://grafana.com/grafana/dashboards/23274) and [dashboards/single](https://grafana.com/grafana/dashboards/22084): add panels for [Pressure Stall Information (PSI)](https://docs.kernel.org/accounting/psi.html) metrics to dashboards. They could help to identify shortage of resources for VictoriaLogs components.
* FEATURE: [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe): add `now()` function, which returns the current [Unix timestamp](https://en.wikipedia.org/wiki/Unix_time) in nanoseconds.

* BUGFIX: [OpenTelemetry](https://docs.victoriametrics.com/victorialogs/data-ingestion/opentelemetry/): properly handle nested attributes by expanding them into separate top-level fields. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8862).
* BUGFIX: [Datadog](https://docs.victoriametrics.com/victorialogs/data-ingestion/datadog-agent/): respond HTTP 202 instead of HTTP 200 on successful Datadog endpoint ingestion as it's strictly required by Datadog agent. See [#8956](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8956).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): properly escape special characters in field values shown in autocomplete suggestions. See [#8925](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8925).
* BUGFIX: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): Properly handle time filters when querying vlstorage directly or through vlselect. See [#8985](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8985).
* BUGFIX: Self-healing from OOM interruption during the creation of a daily partition. See [#8873](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8873).

## [v1.22.2](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.22.2-victorialogs)

Released at 2025-05-10

* BUGFIX: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): properly calculate [`min`](https://docs.victoriametrics.com/victorialogs/logsql/#min-stats) and [`max`](https://docs.victoriametrics.com/victorialogs/logsql/#max-stats) stats for numeric log fields.

## [v1.22.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.22.1-victorialogs)

Released at 2025-05-09

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix issue where log entries did not update after modifying the query. See [#8912](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8912).

## [v1.22.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.22.0-victorialogs)

Released at 2025-05-08

* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): add [`decolorize` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#decolorize-pipe) for dropping [ANSI color codes](https://en.wikipedia.org/wiki/ANSI_escape_code) from the given log fields.
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): add the ability to remove [ANSI color codes](https://en.wikipedia.org/wiki/ANSI_escape_code) from the ingested log fields before storing them in the database. This can be done by passing comma-separated list of field names to remove ANSI color codes from, to `decolorize_fields` HTTP query arg or to `VL-Decolorize-Fields` HTTP header. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-parameters) for details.
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): support for comma-separated list of fields at the `_time_field` HTTP query arg and at `VL-Time-Field` HTTP header. Then the first non-empty log field from the list is used as a timestamp field. This is useful when the ingested logs may contain log timestamp across different fields. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-parameters) for details.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add sorting of fields by key in the Group tab. See [#8438](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8438).
* FEATURE: [JSON lines data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/#json-stream-api): return `400 Bad Request` if no logs were successfully processed. This improves error visibility for malformed requests. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8818).

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix log entry sorting in group mode (newest logs appear first). See [#8726](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8726).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix page freeze after timeline zooming and panning. See [#8655](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8655).
* BUGFIX: [replace_regexp](https://docs.victoriametrics.com/victorialogs/logsql/#replace_regexp-pipe): fixed infinite loop when regex pattern matches empty strings (e.g. `\d*`, `()`, `\b`). Also fixed incorrect behavior with anchors (`^`) due to repeated string slicing. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8625).

## [v1.21.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.21.0-victorialogs)

Released at 2025-04-25

* FEATURE: [querying](https://docs.victoriametrics.com/victorialogs/querying/): support for multiple `extra_filters` and `extra_stream_filters` args. This should simplify applying multiple independent filters to the `query`. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#extra-filters) and [this comment](https://github.com/VictoriaMetrics/victorialogs-datasource/issues/293#issuecomment-2827285583).
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): add an ability to force flush the recently ingested logs, so they become available for querying. See [these docs](https://docs.victoriametrics.com/victorialogs/#forced-flush).
* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): add [`sample` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sample-pipe), which returns `1/Nth` random sample for the selected logs.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add a toggle to handle ANSI escape sequences in log messages. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6614).

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix the Group tab to display raw JSON when `_msg` is missing. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8205).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix display of `Query history` icon in mobile view. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8788).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix zoom behavior on logs chart. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8558).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): preserve zoom selection in the logs chart when auto-refresh is enabled. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8557).

## [v1.20.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.20.0-victorialogs)

Released at 2025-04-22

* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): reduce CPU usage by up to 10x when ingesting small number of log entries per request.
* FEATURE: [dashboards/cluster](https://grafana.com/grafana/dashboards/23274): add Grafana dashboard for cluster version of VictoriaLogs. The source of the dashboard is available [here](https://github.com/VictoriaMetrics/VictoriaMetrics/blob/master/dashboards/victorialogs-cluster.json) and will continue improving with time.
* FEATURE: [deployment/docker](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/docker#readme): add [docker-compose environment for VictoriaLogs cluster](https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/deployment/docker#victorialogs-cluster) deployment. It is intended for demonstrative purposes, includes alerting and Grafana dashboards.

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix high CPU usage when hovering over log entries due to complex button styles. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8135).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add 0 label to the Y axis. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8409).
* BUGFIX: [OpenTelemetry](https://docs.victoriametrics.com/victorialogs/data-ingestion/opentelemetry/): add missed WARN severity levels. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8647). Thanks to @090809 for [the bugfix](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/8648).

## [v1.19.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.19.0-victorialogs)

Released at 2025-04-17

* FEATURE: [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe): add an ability to format [duration values](https://docs.victoriametrics.com/victorialogs/logsql/#duration-values) as floating-point seconds via `<duration_seconds:field_with_duration_value>` syntax.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add button for downloading displayed logs. It supports downloading in the following formats: csv, json. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8604). Thanks to @arturminchukov .
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): optimize vmui for mobile layout to use space more efficiently. See [this pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/8679). Thanks to @arturminchukov .
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add query history for quick access to previously executed queries, in the way similar to VictoriaMetrics. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8500). Thanks to @arturminchukov .

* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): properly store the contents of the [`_msg` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field) for newly ingested logs. It was incorrectly replaced with `missing _msg field; see https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field` because of the bug in [v1.18.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.18.0-victorialogs). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8707).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix incorrect table sorting for numeric columns. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8606).

## [v1.18.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.18.0-victorialogs)

Released at 2025-04-10

* FEATURE: add [cluster mode](https://docs.victoriametrics.com/victorialogs/cluster/), which allow scaling VictoriaLogs horizontally to multiple nodes. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8223), [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/5077), [this question](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7950).
* FEATURE: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): add [`json_values` function](https://docs.victoriametrics.com/victorialogs/logsql/#json_values-stats) for JSON-encoding logs into JSON arrays.
* FEATURE: [stream filters](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter): support `{field in (*)}` and `{field not_in (*)}` to be consistent with [`in(*)`](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter). The `{field in (*)}` matches any [log stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields), while the `{field not_in (*)}` matches zero log streams. This is needed for [this feature request for VictoriaLogs datasource in Grafana](https://github.com/VictoriaMetrics/victorialogs-datasource/issues/238).

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix endless group expansion loop bug. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8347).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): respect nanosecond precision when sorting logs. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8346).

## [v1.17.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.17.0-victorialogs)

Released at 2025-03-16

**Update note: this release changes data storage format in backwards-incompatible way, so it is impossible to downgrade to the previous releases after upgrading to this release.
It is safe upgrading to this release and all the future releases from older releases.**

* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): support zstd compression for all HTTP-based ingestion protocols. See [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8380) and [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8300) issues.
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): allow specifying prefixes for log fields, which must be ignored during data ingestion, at `ignore_fields` HTTP query arg and at `VL-Ignore-Fields` HTTP header. For example, `ignore_fields=foo.*,bar.baz.*` will ignore all the [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model), which start from `foo.` or from `bar.baz.`. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-parameters) for details.
* FEATURE: [stream filter](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter): support `{label in ("v1", ..., "vN")}` and `{label not_in ("v1", ..., "vN")}` syntax. It is equivalent to `{label=~"v1|...|vN"}` and `{label!~"v1|...|vN"}` respectively, where `v1`, ... , `vN` are properly escaped inside regexp. For example, `{app in ("foo.bar","baz")}` is equivalent to `{app=~"foo\\.bar|baz"}` - note that the `.` char is properly escaped inside the regexp.
* FEATURE: improve performance when processing constant log fields with length exceeding 256 bytes. For example, repeated stack traces.

* BUGFIX: [Loki data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/): return `204 No Content` HTTP response code from `/insert/loki/api/v1/push` endpoint. Previously `200 Success` code was sent. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8505).
* BUGFIX: [OpenTelemetry data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/opentelemetry/): properly parse `trace_id` and `span_id` [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) as hex numbers. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8502) for details. Thanks to @forgethub for [the pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/8511).

## [v1.16.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.16.0-victorialogs)

Released at 2025-03-12

* FEATURE: [Loki data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/): automatically parse JSON-encoded log fields from the plaintext log message and store them as separate [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). This behavior allows achieving lower disk space usage and higher query performance comparing to the case when the JSON-encoded log fields were stored in VictoriaLogs as a plaintext log message, which should be parsed at query time with [`unpack_json` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_json-pipe). The previous behaviour can be restored if needed by passing `-loki.disableMessageParsing` command-line flag to VictoriaLogs (the previous behavior isn't recommended because it is less efficient). See [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8486).

* BUGFIX: [querying](https://docs.victoriametrics.com/victorialogs/querying/): properly parse floating-point numbers with leading zeroes in fractional part (for example, `12.03` or `1.0002`). Parsing for these numbers has been broken in [v1.15.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.15.0-victorialogs). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8464).
* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): support floating-point timestamps for [Elasticsearch data ingestion protocol](https://docs.victoriametrics.com/victorialogs/data-ingestion/#elasticsearch-bulk-api). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8470).
* BUGFIX: [OpenTelemetry data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/opentelemetry/): properly convert nested OpenTelemetry attributes into JSON. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8384).

## [v1.15.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.15.0-victorialogs)

Released at 2025-02-27

* FEATURE: [`pack_json` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#pack_json-pipe): allow packing fields, which start with the given prefixes. For example, `pack_json fields (foo.*, bar.*)` creates a JSON containing all the fields, which start with either `foo.` or `bar.`.
* FEATURE: [`pack_logfmt` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#pack_logfmt-pipe): allow packing fields, which start with the given prefixes. For example, `pack_logfmt fields (foo.*, bar.*)` creates [logfmt](https://brandur.org/logfmt) message containing all the fields, which start with either `foo.` or `bar.`.
* FEATURE: expose `vl_http_request_duration_seconds` [summaries](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#summary) for [select APIs](https://docs.victoriametrics.com/victorialogs/querying/#http-api) at the [/metrics](https://docs.victoriametrics.com/victorialogs/#monitoring) page.
* FEATURE: allow passing `*` as a subquery inside [`in(*)`, `contains_any(*)` and `contains_all(*)` filters](https://docs.victoriametrics.com/victorialogs/logsql/#subquery-filter). Such filters are treated as `match all` aka `*`. This is going to be used by [Grafana plugin for VictoriaLogs](https://docs.victoriametrics.com/victorialogs/victorialogs-datasource/). See [this issue](https://github.com/VictoriaMetrics/victorialogs-datasource/issues/238#issuecomment-2685447673).
* FEATURE: [victorialogs dashboard](https://grafana.com/grafana/dashboards/22084): add panels to display amount of ingested logs in bytes, latency of [select APIs](https://docs.victoriametrics.com/victorialogs/querying/#http-api) calls, troubleshooting panels.
* FEATURE: provide alternative registry for all VictoriaLogs components at [Quay.io](https://quay.io/organization/victoriametrics): [VictoriaLogs](https://quay.io/repository/victoriametrics/victoria-logs?tab=tags) and [vlogscli](https://quay.io/repository/victoriametrics/vlogscli?tab=tags).

* BUGFIX: do not treat a string containing leading zeros as a number during data ingestion and querying. For example, `00123` string shouldn't be treated as `123` number. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8361).

## [v1.14.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.14.0-victorialogs)

Released at 2025-02-25

* FEATURE: add [`lt_field` filter](https://docs.victoriametrics.com/victorialogs/logsql/#lt_field-filter), which can be used for obtaining logs where the given [field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) value doesn't exceed the other field value.
* FEATURE: add [`le_field` filter](https://docs.victoriametrics.com/victorialogs/logsql/#le_field-filter), which can be used for obtaining logs where the given [field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) value is smaller or equal to the other field value.

* BUGFIX: [elasticsearch data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/#elasticsearch-bulk-api): support health-check endpoint requested by Jaeger v2. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8353).
* BUGFIX: [`stats_query_range` HTTP endpoint](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-range-stats): fix inconsistent result of `stats_query_range` API. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8312).

## [v1.13.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.13.0-victorialogs)

Released at 2025-02-22

* FEATURE: add [`contains_all` filter](https://docs.victoriametrics.com/victorialogs/logsql/#contains_all-filter), which matches logs containing all the given words / phrases in the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: add [`contains_any` filter](https://docs.victoriametrics.com/victorialogs/logsql/#contains_any-filter), which matches logs containing at least one of the given words / phrases in the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: add [`eq_field` filter](https://docs.victoriametrics.com/victorialogs/logsql/#eq_field-filter), which can be used for obtaining logs with identical values at the given [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: add [`unpack_words` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_words-pipe) for unpacking individual [words](https://docs.victoriametrics.com/victorialogs/logsql/#word) from the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) into the given destination field as JSON array.
* FEATURE: add [`json_array_len` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#json_array_len-pipe) for calculating the length of JSON array stored in the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).

## [v1.12.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.12.0-victorialogs)

Released at 2025-02-20

* FEATURE: [`_time` filter](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter): allow using `>`, `>=`, `<` and `<=`. For example, `_time:<2025-02-24Z` selects all the logs with [`_time` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field) less than `2025-02-24` by UTC. Another example: `_time:>1d` selects all the logs with `_time` field older than one day from the current time. This simplifies querying VictoriaLogs.
* FEATURE: [`top` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe): allow specifying the list of log fields without parens. For example, `top 5 foo` is equivalent to `top 5 by (foo)`.
* FEATURE: [`uniq` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#uniq-pipe): allow specifying the list of log fields without parens. For example, `uniq foo` is equivalent to `uniq (foo)`.
* FEATURE: [`unroll` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unroll-pipe): allow specifying the list of log fields without parens. For example, `unroll foo` is equivalent to `unroll (foo)`.

## [v1.11.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.11.0-victorialogs)

Released at 2025-02-19

**Update note: this release changes data storage format in backwards-incompatible way, so it is impossible to downgrade to the previous releases after upgrading to this release.
It is safe upgrading to this release and all the future releases from older releases.**

* FEATURE: improve per-field data locality on disk. This reduces overhead related to reading data from unrelated fields during queries. This improves query performance over structured logs with big number of fields (aka [wide events](https://jeremymorrell.dev/blog/a-practitioners-guide-to-wide-events/)) when only a small portion of fields are used in the query.
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): reduce memory usage by up to 4x when ingesting [wide events](https://jeremymorrell.dev/blog/a-practitioners-guide-to-wide-events/) at high rate into VictoriaLogs.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add ability to limit the number of logs per page in the Group view (client-side). See [this pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/8334).

* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): Fixed journald ingestion to support entities with single-character names. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8314).

## [v1.10.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.10.1-victorialogs)

Released at 2025-02-14

* BUGFIX: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): fix possible `concurrent map writes` panic when using [`count_uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq-stats) and [`count_uniq_hash`](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq_hash-stats) functions over big number of unique `stats by (...)` groups. The bug has been introduced in [v1.9.0-victorialogs](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.9.0-victorialogs).
* BUGFIX: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): properly calculate `<q> | len(f) f_len | min(f_len)`. Previously it incorrectly returned `0` if the minimum value length for `f` field is bigger than `0`.
* BUGFIX: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): properly calculate [stats functions with additional filters](https://docs.victoriametrics.com/victorialogs/logsql/#stats-with-additional-filters). Previously some [stats functions](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe-functions) could return incorrect results if the additional filter filters out all the values in the processed data block.

## [v1.10.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.10.0-victorialogs)

Released at 2025-02-12

* FEATURE: [OpenTelemetry data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/opentelemetry/): support parsing of [`span_id`](https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md#field-spanid) and [`trace_id`](https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md#field-traceid) log fields. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8255) for details.
* FEATURE: [Journald data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/journald/): add support for data compression at journald side. See [this pull request](https://github.com/systemd/systemd/pull/34822).
* FEATURE: [`stats` by time buckets](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-time-buckets), [`stats` by field buckets](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-field-buckets), [`stats` by IPv4 buckets](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-ipv4-buckets): improve performance for large buckets.

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): respect the selected tenant for autocomplete suggestions. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8042).
* BUGFIX: [`stats` by time buckets](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-time-buckets): properly align start of the week to Monday at `stats by (_time:week) ...`. Previously the week was started on Wednesday. The start of the week can be adjusted with [offset](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-time-buckets-with-timezone-offset). For example, the following query aligns the start of the week to Sunday: `_time:30d | stats by (_time:week offset -1d) count() | sort by (_time)`.
* BUGFIX: [`stats` by field value buckets](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-field-buckets): properly calculate buckets for negative values.
* BUGFIX: [syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): correctly parse rows with allowed in rfc5424 `\]` character. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8282).

## [v1.9.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.9.1-victorialogs)

Released at 2025-02-10

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): properly apply [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8152) and [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8153) bugfixes. They weren't applied to [`v1.9.0-victorialogs`](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.9.0-victorialogs) by an accident.

## [v1.9.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.9.0-victorialogs)

Released at 2025-02-10

* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): improve performance for [`stats by (...) ...`](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) by up to 30% when it is applied to big number of `by (...)` groups.
* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): improve performance for [`top` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe) by up to 30% when it is applied to big number of unique values.
* FEATURE: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): improve performance for [`count_uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq-stats) and [`count_uniq_hash`](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq_hash-stats) functions by up to 30% when they are applied to big number of unique values.
* FEATURE: [`block_stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#block_stats-pipe): return the path to the part where every data block is stored. The path to the part is returned in the `part_path` field. This allows investigating the distribution of data blocks among parts.
* FEATURE: reduce VictoriaLogs startup time by multiple times when it opens a large datastore with big [retention](https://docs.victoriametrics.com/victorialogs/#retention).
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): accept timestamps with microsecond and nanosecond precision at [`_time` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field).
* FEATURE: [JSON lines data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/#json-stream-api): continue parsing lines after encountering parse errors for some lines. Previously the input JSON lines' stream was closed after the first parse error.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add the `_msg` field to the list of fields for the group view, allowing users to select multiple fields, including `_msg`, for log display.

* BUGFIX: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): properly limit [`concurrency` query option](https://docs.victoriametrics.com/victorialogs/logsql/#query-options) for [`stats`](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe), [`uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#uniq-pipe) and [`top`](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe). This prevents from `runtime error: index out of range` panic. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8201).
* BUGFIX: [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe): properly sort [RFC3339 timestamps](https://www.rfc-editor.org/rfc/rfc3339) with variable sub-second precision. Previously such timestamps were sorted using [natural sorting](https://en.wikipedia.org/wiki/Natural_sort_order), and this could lead to unexpected results. For example, `2025-02-21T10:20:30.9Z` was incorrectly considered smaller than `2025-02-21T10:20:30.012Z`, since the last one had higher decimal value after the last dot.
* BUGFIX: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): drop log entries with too long field names and log the dropped log entries with the `ignoring log entry with too long field name` message, so human operators could notice and fix the ingestion of incorrect logs ASAP. Previously too long field names were silently truncated to shorter values. This isn't what most users expect. See [why VictoriaLogs has a limit on the field name length](https://docs.victoriametrics.com/victorialogs/faq/#what-is-the-maximum-supported-field-name-length).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix transparency for bars in the hits bar chart to improve visibility. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8152).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix `Group by field` dropdown menu not displaying any options in Group View settings. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8153).

## [v1.8.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.8.0-victorialogs)

Released at 2025-01-24

* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): add an ability to limit query concurrency for the `<q>` [query](https://docs.victoriametrics.com/victorialogs/logsql/#query-syntax) via `options(concurrency=N) <q>` syntax. This may be needed for reducing RAM and CPU usage at the cost of longer query execution times. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#query-options) for details.
* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): allow overriding the global time range filter at subqueries inside [`in` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter), [`join` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#join-pipe) and [`union` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#union-pipe) via `options(ignore_global_time_filter=true) <q>` syntax. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#query-options) for details.
* FEATURE: add [`hash` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#hash-pipe) for calculating hashes over the selected log fields. This may be useful for splitting the selected logs into distinct buckets. For example, the following query splits `user_id` fields into 4 buckets with the help of `hash` pipe: `_time:5m | hash(user_id) as h | math h%4 as bucket | stats by (bucket) count()`.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): reflect column settings for the table view in URL, so the table view can be shared via link. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7662).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): added context menu for legend items with options to copy and filter streams and fields. See this [PR](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/7750).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): improved legend functionality with consistent click behavior across all vmui charts. See this [PR](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/7750).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): added sorting for group view by record count in descending order. See this [PR](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/7750).

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): removed the ability to filter by other in the legend, as other represents an aggregated series of all streams not included in the top results. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7552).

* BUGFIX: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): properly apply optimizations to all the subqueries. Previously only the top-level query was optimized, while subqueries inside [`in` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter), [`join` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#join-pipe) and [`union` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#union-pipe) weren't optimized.
* BUGFIX: [HTTP querying API](https://docs.victoriametrics.com/victorialogs/querying/#http-api): properly apply [extra filters](https://docs.victoriametrics.com/victorialogs/querying/#extra-filters) to all the subqueries. Previously extra filters were applied only to the top-level query and weren't applied to sub-queries inside [`in` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter), [`join` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#join-pipe) and [`union` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#union-pipe).
* BUGFIX: [HTTP querying API](https://docs.victoriametrics.com/victorialogs/querying/#http-api): properly apply time range filter from `start` and `end` args to all the subqueries. Previously the time range filters were applied only to the top-level query and weren't applied to sub-queries inside [`in` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter), [`join` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#join-pipe) and [`union` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#union-pipe). The global time range filter can be ignored on a per-subquery basis via the `ignore_global_time_filter=true` option - see [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#query-options) for details.

## [v1.7.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.7.0-victorialogs)

Released at 2025-01-20

* FEATURE: [`join` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#join-pipe): add an ability to execute `INNER JOIN` by adding `inner` suffix to the `join` pipe.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): updated the default graph type in the `hits panel` to bars with color fill. Removed options for `lines`, `stepped lines`, and `points`. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7101).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): reduce logs text size and improved styles in grouped view. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7479).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add the ability to select fields for display instead of the `_msg` field. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7419).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add various display configuration settings for the grouped view. See [this pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/7815)

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix an issue where pressing the "Enter" key in the query editor did not execute the query. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8058).

## [v1.6.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.6.1-victorialogs)

Released at 2025-01-16

* BUGFIX: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): properly apply multiple [log stream filters](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter). For example, `{foo="bar"} AND {baz="x"}` must correctly return logs with `foo="bar"` and `baz="x"` [log stream fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/8037).

## [v1.6.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.6.0-victorialogs)

Released at 2025-01-15

* FEATURE: add [`union` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#union-pipe), which can be used for returning results from multiple independent LogsQL queries.
* FEATURE: add [`histogram` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#histogram-stats) for calculating [VictoriaMetrics histogram buckets](https://valyala.medium.com/improving-histogram-usability-for-prometheus-and-grafana-bc7e5df0e350) over the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). They will be used for building heatmaps at the [built-in Web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui) and [VictoriaLogs plugin for Grafana](https://docs.victoriametrics.com/victorialogs/victorialogs-datasource/).
* FEATURE: [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe): add `rand()` function, which can be used for generating random numbers in the range `[0 ... 1)`.

## [v1.5.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.5.0-victorialogs)

Released at 2025-01-13

* FEATURE: [`count_uniq` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq-stats): improve performance by up to 50% and reduce memory usage by up to 4x when this function is applied to fields with big number of unique integer values.
* FEATURE: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): improve performance and reduce memory usage by up to 50% for `log_field` with big number of unique values at `stats by (log_field) ...`.
* FEATURE: [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe): improve performance by up to 8x.
* FEATURE: [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe): add ability to apply URL encoding / decoding (aka [percent encoding](https://en.wikipedia.org/wiki/Percent-encoding)) to the formatted log fields with `<urlencode:field_name>` and `<urldecode:field_name>` syntax.
* FEATURE: [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe): add ability to [base64-encode](https://en.wikipedia.org/wiki/Base64) / base64-decode log fields with `<base64encode:field_name>` and `<base64decode:field_name>` syntax.
* FEATURE: [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe): add ability to hex-encode / hex-decode log fields with `<hexencode:field_name>` and `<hexdecode:field_name>` syntax.
* FEATURE: [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe): add ability to convert integers to hexadecimal numbers via `<hexnumencode:field_name>` and `<hexnumdecode:field_name>` syntax.
* FEATURE: add [`value_type` filter](https://docs.victoriametrics.com/victorialogs/logsql/#value_type-filter), which can be useful during exploration of the stored logs.
* FEATURE: [Datadog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/datadog-agent/): added `-datadog.streamFields` and `-datadog.ignoreFields` flags to configured default stream and ignore fields. Useful for Datadog serverless plugin, which doesn't allow to provide extra headers of query args.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add support for autocomplete in LogsQL queries. This feature provides suggestions for field names, field values, and pipe names.

* BUGFIX: [Datadog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/datadog-agent/): accepts `message` field as both string and object type to fix compatibility with Datadog serverless extension, which sends logs data in format, which is not documented. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7761).
* BUGFIX: [vlinsert](https://docs.victoriametrics.com/victorialogs/): order of `VL-Msg-Field` fields now defines a priority of these fields.

## [v1.4.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.4.0-victorialogs)

Released at 2024-12-22

* FEATURE: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): allow non-numeric field values at [`median`](https://docs.victoriametrics.com/victorialogs/logsql/#median-stats) and [`quantile`](https://docs.victoriametrics.com/victorialogs/logsql/#quantile-stats) stats functions.
* FEATURE: improve performance of [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) and [`top` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe) by up to 2x when these pipes are applied to logs with millions of unique `by (...)` groups.
* FEATURE: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): add [`count_uniq_hash`](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq_hash-stats) function, which counts the number of unique value hashes. This number is usually a good approximation to the number of unique values, so the `count_uniq_hash` can be used as a faster alternative to [`count_uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq-stats).
* FEATURE: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): improve performance of [`count_uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq-stats) and [`uniq_values`](https://docs.victoriametrics.com/victorialogs/logsql/#uniq_values-stats) functions when they are applied to fields with big number of unique values.
* FEATURE: [`facets` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#facets-pipe): add an ability to return log fields with the same values across all the selected logs by adding `keep_const_fields` option. Such log fields aren't interesting in most cases, so they aren't returned by default.
* FEATURE: [`in` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter): improve performance for `in(<query>)` when the `<query>` returns big number of values.
* FEATURE: [HTTP querying APIs](https://docs.victoriametrics.com/victorialogs/querying/#http-api): allow passing arbitrary [LogsQL filters](https://docs.victoriametrics.com/victorialogs/logsql/#filters) to `extra_filters` and `extra_stream_filters` query args. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#extra-filters) and [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/5542) for details.
* FEATURE: [Grafana Loki data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/): add support of Loki healthcheck `/insert/ready` endpoint. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7824).
* FEATURE: [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe): return an error as soon as too many logs and/or [log streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) are passed to this pipe. This prevents from excess resource usage by the `stream_context` pipe when it is improperly used. It is expected that the results of this pipe are investigated by humans, who cannot inspect surrounding logs for millions of the logs passed to `stream_context`. This change addresses [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7903) and [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7766) issues.

* BUGFIX: [syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): correctly parse rows with multiple consecutive spaces between fields. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7776).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix cursor reset in query input field. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7288).
* BUGFIX: [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe): fix improper sorting of numeric fields in some cases.
* BUGFIX: properly return an empty minimum value from [`min` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#min-stats).

## [v1.3.2](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.3.2-victorialogs)

Released at 2024-12-09

* FEATURE: [`collapse_nums` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#collapse_nums-pipe): add an ability to prettify some patterns across collapsed numbers. For example, `<N>.<N>.<N>.<N>` is replaced with `<IP4>` when executing `collapse_nums prettify` pipe.

* BUGFIX: [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe): fix `index out of range [0] with length 0` panic, which has been introduced in [v1.3.0-victorialogs](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.3.0-victorialogs). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7762).

## [v1.3.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.3.1-victorialogs)

Released at 2024-12-08

* BUGFIX: [`facets` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#facets-pipe): fix `assignment to entry in nil map` panic, which has been introduced in [v1.3.0-victorialogs](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.3.0-victorialogs).

## [v1.3.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.3.0-victorialogs)

Released at 2024-12-08

* FEATURE: add [`collapse_nums` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#collapse_nums-pipe), which replaces all the decimal and hexadecimal numbers with `<N>` in the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). This can be useful for locating the most frequently seen log message patterns if log messages differ only by decimal and hexadecimal numbers (this is very frequent case). For example, the following query returns top 5 log message patterns seen over the last hour: `_time:1h | collapse_nums | top 5 by (_msg)`.
* FEATURE: improve performance for [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe) over [log streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) with big number of logs (millions and more). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7637).
* FEATURE: [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe) allow changing the time window for search for surrounding logs via `time_window` option. For example, the following query searches for surrounding [log stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) logs on the one week window: `_time:5m error | stream_context before 10 time_window 1w`. Thanks to @worker24h for [the idea](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7637#issuecomment-2523313740).

## [v1.2.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.2.0-victorialogs)

Released at 2024-12-06

* FEATURE: add [`rate`](https://docs.victoriametrics.com/victorialogs/logsql/#rate-stats) and [`rate_sum`](https://docs.victoriametrics.com/victorialogs/logsql/#rate_sum-stats) stats functions, which can be used for calculating the average per-second rate of matching logs and the average per-second rate of sum over the given numeric [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). See [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7415).
* FEATURE: add [`facets` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#facets-pipe), which can be used for returning the most frequent values across all the fields seen in the selected logs. This pipe simplifies logs' exploration.
* FEATURE: add [`/select/logsql/facets` HTTP endpoint](https://docs.victoriametrics.com/victorialogs/querying/#querying-facets), which returns the most frequent values across all the fields seen in the selected logs. This endpoint is going to be used for building faceted search over logs in the [VictoriaLogs web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui).

* BUGFIX: [`/select/logsql/stats_query`](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-stats): properly apply `limit` at [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe). The bug was introduced in the release [v1.1.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.1.0-victorialogs) when fixing [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7699).
* BUGFIX: [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe): properly format expressions with multiple binary operations with the same priority. For example, `x / (y * z)` was improperly formatted as `x / y * z`, while `x - (y + z)` was improperly formatted as `x - y + z`. This could lead to incorrect query results in [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/).

## [v1.1.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.1.0-victorialogs)

Released at 2024-12-05

* FEATURE: add [`first`](https://docs.victoriametrics.com/victorialogs/logsql/#first-pipe) and [`last`](https://docs.victoriametrics.com/victorialogs/logsql/#last-pipe) pipes for returning the first `N` and the last `N` logs after sorting them by the given set of [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). For example, the following query returns up to 5 logs with the biggest value for `request_duration` over the last hour: `_time:1h | last 5 by (request_duration)`.
* FEATURE: [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe): add an ability to apply `limit` individually per group of logs via `partition by (...)` syntax. For example, the following query returns up to 3 logs with the smallest `request_duration` individually per each `host`: `_time:5m | sort by (request_duration) limit 3 partition by (host)`.
* FEATURE: [format pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe): allow formatting log fields in lowercase and uppercase via `<uc:field_name>` and `<lc:field_name>` syntax. This can be useful when some fields must be consistently transformed to the same case during querying. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7620#issuecomment-2502170924).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add frontend-only pagination for table view.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): improve memory consumption during data processing. This enhancement reduces the overall memory footprint, leading to better performance and stability.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): reduce memory usage across all tabs for improved performance and stability. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7185).
* FEATURE: [Grafana Loki data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/): use Loki [stream labels](https://grafana.com/docs/loki/latest/get-started/labels/) as VictoriaLogs [stream fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) by default. The set of stream fields can be overridden via `_stream_fields` query arg or via `VL-Stream-Fields` header as described [here](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-parameters).
* FEATURE: [OpenTelemetry data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/opentelemetry/): use [resource labels](https://opentelemetry.io/docs/concepts/resources/) as VictoriaLogs [stream fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) by default. The set of stream fields can be overridden via `_stream_fields` query arg or via `VL-Stream-Fields` header as described [here](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-parameters).
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): expose `vl_bytes_ingested_total` [counter](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#counter) at `/metrics` page. This counter tracks an estimated number of bytes processed when parsing the ingested logs. This counter is exposed individually per every [supported data ingestion protocol](https://docs.victoriametrics.com/victorialogs/data-ingestion/) - the protocol name is exposed in the `type` label. For example, `vl_bytes_ingested_total{type="jsonline"}` tracks an estimated number of bytes processed when reading the ingested logs via [json line protocol](https://docs.victoriametrics.com/victorialogs/data-ingestion/#json-stream-api). Thanks to @tenmozes for the idea and [the initial implementation](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/7682).
* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): expose `vl_too_long_lines_skipped_total` [counter](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#counter) at `/metrics` page. This counter tracks the number of the ingested lines with the length bigger than the value of `-insert.maxLineSizeBytes` command-line flag. Such lines are ignored.

* BUGFIX: [`/select/logsql/stats_query_range` API](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-range-stats): properly handle [`limit` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#limit-pipe) after [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe). Previously the `limit` was applied globally across all the calculated stats, while it must be applied individually per each `step` on the `start ... end` time range. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7699).
* BUGFIX: [vmui](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#vmui): fix for `showLegend` and `alias` flags in predefined panels. [See this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7565)
* BUGFIX: fix `too big number of columns detected in the block` panic when the ingested logs contain more than 2000 fields with different names per every [log stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7568) for details.
* BUGFIX: properly parse lines after too long [JSON lines](https://docs.victoriametrics.com/victorialogs/data-ingestion/#json-stream-api) and [Elasticsearch lines](https://docs.victoriametrics.com/victorialogs/data-ingestion/#elasticsearch-bulk-api) with the length exceeding `-insert.maxLineSizeBytes`. Previously all the lines after the too long line in the stream were ignored.

## [v1.0.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v1.0.0-victorialogs)

Released at 2024-11-12

This release is identical to [v0.42.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.42.0-victorialogs).

VictoriaLogs gained all [the planned features](https://docs.victoriametrics.com/victorialogs/)
since the [initial v0.1.0 release](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.1.0-victorialogs) 1.5 years ago,
and is ready for production!

## [v0.42.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.42.0-victorialogs)

Released at 2024-11-08

* FEATURE: [`join` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#join-pipe): add an ability to add prefix to all the log field names from the joined query, by using `| join by (<by_fields>) (<query>) prefix "some_prefix"` syntax.
* FEATURE: [`_time` filter](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter): allow specifying offset without time range. For example, `_time:offset 1d` matches all the logs until `now-1d` in the [`_time` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field). This is useful when building graphs for time ranges with some offset in the past.
* FEATURE: [`/select/logsql/tail` HTTP endpoint](): support for `offset` query arg, which can be used for delayed emission of matching logs during live tailing. Thanks to @Fusl for the initial idea and implementation in [this pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/7428).
* FEATURE: [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/): allow enabling and disabling wrapping of long lines, which do not fit screen width, with `\wrap_long_lines` command.
* FEATURE: [syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): allow overriding default [stream fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) with the given [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) during data ingestion. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#stream-fields). See [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7480).
* FEATURE: [syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): allow adding arbitrary [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) via `[label1=value1 ... labelN=valueN]` syntax inside Syslog messages. For example, `<165>1 2024-06-03T17:42:00.000Z example.com appname 12345 ID47 [field1=value1 field2=value2] some message`.
* FEATURE: [syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): allow dropping the specified [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) during data ingestion. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#dropping-fields).
* FEATURE: [syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): allow adding the specified [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) during data ingestion. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#adding-extra-fields). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7354).
* FEATURE: [Loki data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/#loki-json-api): show the original request body on parse errors. This should simplify debugging. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7490).

* BUGFIX: [`values` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#values-stats): fix a bug, which could lead to corrupted results. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7458).
* BUGFIX: [`uniq_values`](https://docs.victoriametrics.com/victorialogs/logsql/#uniq_values-stats), [`min`](https://docs.victoriametrics.com/victorialogs/logsql/#min-stats), [`max`](https://docs.victoriametrics.com/victorialogs/logsql/#max-stats), [`row_min`](https://docs.victoriametrics.com/victorialogs/logsql/#row_min-stats) and [`row_max`](https://docs.victoriametrics.com/victorialogs/logsql/#row_max-stats) stats functions: fix a bug, which could return non-matching field values for these functions. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7458).
* BUGFIX: [HTTP querying APIs](https://docs.victoriametrics.com/victorialogs/querying/#http-api): properly take into account the `end` query arg when calculating time range for [`_time:duration` filter](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter). Previously the `_time:duration` filter was treated as `_time:[now-duration, now)`, while it should be treated as `_time:[end-duration, end)`.

## [v0.41.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.41.0-victorialogs)

Released at 2024-11-06

* FEATURE: support [structured metadata](https://grafana.com/docs/loki/latest/get-started/labels/structured-metadata/) when ingesting logs with [Grafana Loki ingestion protocol](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7431).
* FEATURE: add [`join` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#join-pipe), which can be used for performing SQL-like joins.
* FEATURE: support returning historical logs from [live tailing API](https://docs.victoriametrics.com/victorialogs/querying/#live-tailing) via `start_offset` query arg. For example, request to `/select/logsql/tail?query=*&start_offset=5m` returns logs for the last 5 minutes before starting returning live tailing logs for the given `query`.
* FEATURE: add an ability to specify extra fields for logs ingested via [HTTP-based data ingestion protocols](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-apis). See `extra_fields` query arg and `VL-Extra-Fields` HTTP header in [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-parameters).
* FEATURE: add [`block_stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#block_stats-pipe) for returning various per-block stats. This pipe is useful for debugging.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add sorting of logs by groups and within each group by time in desc order. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7184) and [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7045).
* FEATURE: add support for receiving DataDog logs over network. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6632).

* BUGFIX: properly sort fields with floating-point numbers by [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe). Previously floating-point numbers could be improperly sorted because they were treated as strings, and [natural sorting](https://en.wikipedia.org/wiki/Natural_sort_order) was incorrectly applied to them. For example, `0.123` was treated as bigger than `0.9`.

## [v0.40.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.40.0-victorialogs)

Released at 2024-10-31

* FEATURE: add support for extra filters across all the [HTTP querying APIs](https://docs.victoriametrics.com/victorialogs/querying/#http-api). See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#extra-filters) for details. This is needed for implementing quick filtering on field values at [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7365).

* BUGFIX: properly apply [`replace`](https://docs.victoriametrics.com/victorialogs/logsql/#replace-pipe) and [`replace_regexp`](https://docs.victoriametrics.com/victorialogs/logsql/#replace_regexp-pipe) pipes to identical values in adjacent log entries. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7162).
* BUGFIX: properly apply [`extract`](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe) and [`extract_regexp`](https://docs.victoriametrics.com/victorialogs/logsql/#extract_regexp-pipe) pipe with additional `if (...)` filter (aka [conditional extract](https://docs.victoriametrics.com/victorialogs/logsql/#conditional-extract) and [conditional extract_regexp](https://docs.victoriametrics.com/victorialogs/logsql/#conditional-extract_regexp)).

## [v0.39.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.39.0-victorialogs)

Released at 2024-10-30

* FEATURE: allow specifying a list of log fields, which may contain log message, via `_msg_field` query arg and via `VL-Msg-Field` HTTP request header. For example, `_msg_field=message,event.message` instructs obtaining [message field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field) from the first non-empty field out of the `message` and `event.message` fields. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/#http-parameters) for details.
* FEATURE: accept logs without [`_msg` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field). In this case the `_msg` field is automatically set to the value specified in the `-defaultMsgValue` command-line flag.

* BUGFIX: fix `runtime error: index out of range [0] with length 0` panic during low-rate data ingestion. The panic has been introduced in [v0.38.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.38.0-victorialogs). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7391).

## [v0.38.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.38.0-victorialogs)

Released at 2024-10-29

* FEATURE: added the ability to receive systemd (journald) logs over network. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4618).
* FEATURE: improve performance for queries over large volume of logs with big number of [fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) (aka `wide events`).
* FEATURE: improve performance for [`/select/logsql/field_values` HTTP endpoint](https://docs.victoriametrics.com/victorialogs/querying/#querying-field-values).
* FEATURE: improve performance for [`field_values` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#field_values-pipe) when it is applied directly to [log filter](https://docs.victoriametrics.com/victorialogs/logsql/#filters).
* FEATURE: add an ability to return `rank` field from [`top` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe). For example, the following query returns `1..5` rank per each returned `ip` with the biggest number of logs over the last 5 minute: `_time:5m | top 5 by (ip) rank`.

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix various glitches with updating query responses. The issue was introduced in [v0.36.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.36.0-victorialogs). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7279).

## [v0.37.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.37.0-victorialogs)

Released at 2024-10-18

* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add ability to hide hits chart. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7117).
* FEATURE: add basic [alerting rules](https://github.com/VictoriaMetrics/VictoriaMetrics/blob/master/deployment/docker/rules/alerts-vlogs.yml) for VictoriaLogs process. See details at [monitoring docs](https://docs.victoriametrics.com/victorialogs/#monitoring).
* FEATURE: improve [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) performance on systems with many CPU cores when `by(...)` fields contain big number of unique values. For example, `_time:1d | stats by (user_id) count() x` should be executed much faster when `user_id` field contains millions of unique values.
* FEATURE: improve performance for [`top`](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe), [`uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#uniq-pipe) and [`field_values`](https://docs.victoriametrics.com/victorialogs/logsql/#field_values-pipe) pipes on systems with many CPU cores when it is applied to [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) with big number of unique values. For example, `_time:1d | top 5 (user_id)` should be executed much faster when `user_id` field contains millions of unique values.
* FEATURE: improve performance for [`field_names` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#field_names-pipe) when it is applied to logs with hundreds of [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix display of hits chart. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7133).

## [v0.36.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.36.0-victorialogs)

Released at 2024-10-16

* FEATURE: optimize [LogsQL queries](https://docs.victoriametrics.com/victorialogs/logsql/), which need to scan big number of logs with big number of [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) (aka `wide events`). The performance for such queries is improved by 10x and more depending on the number of log fields in the scanned logs. The performance improvement is visible when querying logs ingested after the upgrade to this release.
* FEATURE: add support for forced merge. See [these docs](https://docs.victoriametrics.com/victorialogs/#forced-merge).
* FEATURE: skip empty log fields in query results, since they are treated as non-existing fields in [VictoriaLogs data model](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). This should reduce the level of confusion for end users when they see empty log fields.
* FEATURE: allow using [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe) for creating output labels from existing log fields at [`/select/logsql/stats_query`](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-stats) and [`/select/logsql/stats_query_range`](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-range-stats) endpoints.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add the ability to cancel running queries. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7097).

* BUGFIX: avoid possible panic when logs for a new day are ingested during execution of concurrent queries.
* BUGFIX: avoid panic at `lib/logstorage.(*blockResultColumn).forEachDictValue()` when the query contains [stats with additional filters](https://docs.victoriametrics.com/victorialogs/logsql/#stats-with-additional-filters). The panic has been introduced in [v0.33.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.33.0-victorialogs) in [this commit](https://github.com/VictoriaMetrics/VictoriaMetrics/commit/a350be48b68330ee1a487e1fb09b002d3be45163).
* BUGFIX: add more checks for [stats query APIs](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-stats) to avoid invalid results.
* BUGFIX: [vmui](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#vmui): fix error messages rendering from overflowing the screen with long messages. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7207).

## [v0.35.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.35.0-victorialogs)

Released at 2024-10-09

* FEATURE: [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/): add ability to live tail query results - see [these docs](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/#live-tailing).
* FEATURE: [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/): add compact output mode for query results. It can be enabled by typing `\c` and then pressing `enter`. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/#output-modes).
* FEATURE: [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/): add `-accountID` and `-projectID` command-line flags for setting `AccountID` and `ProjectID` values when querying the specific [tenants](https://docs.victoriametrics.com/victorialogs/#multitenancy).

## [v0.34.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.34.0-victorialogs)

Released at 2024-10-08

* FEATURE: [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/): add ability to display results in `logfmt` mode, single-line and multi-line JSON modes according [these docs](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/#output-modes).
* FEATURE: [vlogscli](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/): preserve `less` output after the exit from scrolling mode. This should help reusing previous query results in subsequent queries.
* FEATURE: add [`len` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#len-pipe) for calculating the length for the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) value in bytes.

## [v0.33.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.33.0-victorialogs)

Released at 2024-10-01

* FEATURE: add interactive command-line tool for querying VictoriaLogs - [`vlogscli`](https://docs.victoriametrics.com/victorialogs/querying/vlogscli/).

* BUGFIX: [`count_uniq` stats function](https://docs.victoriametrics.com/victorialogs/logsql/#count_uniq-stats): do not count field values, which aren't matched by the used filters. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7152).

## [v0.32.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.32.1-victorialogs)

Released at 2024-09-30

* BUGFIX: do not return field values with zero matching logs from [`field_values`](https://docs.victoriametrics.com/victorialogs/logsql/#field_values-pipe), [`top`](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe) and [`uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#uniq-pipe) pipes. See [this issue](https://github.com/VictoriaMetrics/victorialogs-datasource/issues/72#issuecomment-2352078483).

## [v0.32.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.32.0-victorialogs)

Released at 2024-09-29

* FEATURE: [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/): accept Unix timestamps in seconds in the ingested logs. This simplifies integration with systems, which prefer Unix timestamps over text-based representation of time.
* FEATURE: [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe): allow using `order` alias instead of `sort`. For example, `_time:5s | order by (_time)` query works the same as `_time:5s | sort by (_time)`. This simplifies [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/) transition from SQL-like query languages.
* FEATURE: [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe): allow using multiple identical [stats functions](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe-functions) with distinct [filters](https://docs.victoriametrics.com/victorialogs/logsql/#stats-with-additional-filters) and automatically generated result names. For example, `_time:5m | count(), count() if (error)` query works as expected now, e.g. it returns two results over the last 5 minutes: the total number of logs and the number of logs with `error` [word](https://docs.victoriametrics.com/victorialogs/logsql/#word). Previously this query couldn't be executed because the `if (...)` condition wasn't included in the automatically generate result name, so both results had the same name - `count(*)`.

* BUGFIX: properly calculate [`uniq`](https://docs.victoriametrics.com/victorialogs/logsql/#uniq-pipe) and [`top`](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe) pipes. Previously they could return invalid results in some cases.

## [v0.31.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.31.0-victorialogs)

Released at 2024-09-27

* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): improved readability of staircase graphs and tooltip usability. See [this comment](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6545#issuecomment-2336805237).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): simplify query input by adding only the label name when `ctrl`+clicking the line legend. See [this comment](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6545#issuecomment-2336805237).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): keep selected columns in table view on page reloads. Before, selected columns were reset on each update. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7016).
* FEATURE: allow skipping `_stream:` prefix in [stream filters](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter). This simplifies writing queries with stream filters. Now `{foo="bar"}` is the recommended format for stream filters over the `_stream:{foo="bar"}` format.
* FEATURE: allow using `-` instead of `!` as `NOT` operator shorthand in [logical filters](https://docs.victoriametrics.com/victorialogs/logsql/#logical-filter). For example, `-info -warn` query is equivalent to `!info !warn`. This simplifies transition from other query languages with full-text search support, which usually use `-` as `NOT` operator.

## [v0.30.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.30.1-victorialogs)

Released at 2024-09-27

* BUGFIX: consistently return matching log streams sorted by time from [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe). Previously log streams could be returned in arbitrary order with every request. This could complicate using `stream_context` pipe.
* BUGFIX: [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe): add missing `_msg="---"` delimiter between stream contexts belonging to different [log streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields). This should simplify investigating `stream_context` output for multiple matching log streams.

## [v0.30.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.30.0-victorialogs)

Released at 2024-09-27

* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add button for enabling auto refresh, similarly to VictoriaMetrics vmui. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7017).
* FEATURE: drop logs without [`_msg`](https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field) field or with empty `_msg` field, since this field is required to be non-empty in [VictoriaLogs data model](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6785).
* FEATURE: improve performance of analytical queries, which do not need reading the `_time` field. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7070).
* FEATURE: add [`blocks_count` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#blocks_count-pipe), which can be used for counting the number of matching blocks for the given query. For example, `_time:5m | blocks_count` returns the number of blocks with logs for the last 5 minutes. This pipe can be useful for debugging purposes.
* FEATURE: support [ingesting logs](https://docs.victoriametrics.com/victorialogs/data-ingestion/) with `_time` field, which doesn't contain timezone information. For example, `2024-09-20T10:20:30`. In this case the local timezone of the host where VictoriaLogs runs is used. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6721).
* FEATURE: reduce memory usage when [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe) is applied to [log streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) with big number of messages. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6730).

* BUGFIX: fix Windows build, which has been broken in [v0.29.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.29.0-victorialogs). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6973).
* BUGFIX: properly return logs from [`/select/logsql/tail` endpoint](https://docs.victoriametrics.com/victorialogs/querying/#live-tailing) if the query contains [`_time:some_duration` filter](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter) like `_time:5m`. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7028). The bug has been introduced in [v0.29.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.29.0-victorialogs).
* BUGFIX: properly return logs without [`_msg`](https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field) field when `*` query is passed to [`/select/logsql/query` endpoint](https://docs.victoriametrics.com/victorialogs/querying/#querying-logs) together with positive `limit` arg. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6785). Thanks to @jiekun for identifying the root cause of the issue.
* BUGFIX: support [ingesting logs](https://docs.victoriametrics.com/victorialogs/data-ingestion/) with `_time` field containing whitespace delimiter between the date and time instead of `T` delimiter. For example, `2024-09-20 10:20:30`. This is valid [ISO8601 format](https://en.wikipedia.org/wiki/ISO_8601) aka `SQL datetime` format, which sometimes is used in production. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6721).
* BUGFIX: return all the requested surrounding logs for [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe). Previously only logs matching the [`_time` filter](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter) were returned. This is needed for [this feature](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/7063).

## [v0.29.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.29.0-victorialogs)

Released at 2024-09-08

* FEATURE: add [`/select/logsql/stats_query` HTTP API](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-stats), which is going to be used by [vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/) for executing alerting and recording rules against VictoriaLogs. See [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6942) for details.
* FEATURE: add [`/select/logsql/stats_query_range` HTTP API](https://docs.victoriametrics.com/victorialogs/querying/#querying-log-range-stats), which is going to be used by [VictoriaLogs plugin for Grafana](https://docs.victoriametrics.com/victorialogs/victorialogs-datasource/) for building time series panels. See [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6943) for details.
* FEATURE: optimize [multi-exact queries](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter) with many phrases to search. For example, `ip:in(path:="/foo/bar" | keep ip)` when there are many unique values for `ip` field among log entries with `/foo/bar` path.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add support for displaying the top 5 log streams in the hits graph. The remaining log streams are grouped into an "other" label. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6545).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add the ability to customize the graph display with options for bar, line, stepped line, and points.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add fields for setting AccountID and ProjectID. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6631).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add a toggle button to the "Group" tab that allows users to expand or collapse all groups at once.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): introduce the ability to select a key for grouping logs within the "Group" tab.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): display the number of entries within each log group.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): move the Markdown toggle to the general settings panel in the upper left corner.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add search functionality to the column display settings in the table. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6668).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add the ability to select all columns in the column display settings of the table. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6668). Thanks to @yincongcyincong for the [pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/6680).
* FEATURE: Allow to define ingestion parameters via headers. Supported headers - `VL-Msg-Field`,`VL-Stream-Fields`,`VL-Ignore-Fields`,`VL-Time-Field`, `VL-Debug`. See this [PR](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/6443) for details.
* FEATURE: [vlinsert](https://docs.victoriametrics.com/victorialogs/): added OpenTelemetry logs ingestion support. See this [PR](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/6218) for details.

* BUGFIX: properly handle Logstash requests for Elasticsearch configuration when using `outputs.elasticsearch` in Logstash pipelines. Previously, the requests could be rejected with `400 Bad Request` response. Updates [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4750).
* BUGFIX: [vmui](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#vmui): fix `not found index.js` error when loading vmui in VictoriaLogs. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6764). Thanks to @yincongcyincong for the [pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/6770).
* BUGFIX: properly execute queries with `OR` [filters](https://docs.victoriametrics.com/victorialogs/logsql/#logical-filter) for distinct [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). For example, `field1:foo OR field2:bar`. Previously logs matching these filters may be skipped during querying. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6554) for details. Thanks to @yincongcyincong for the [pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/6556).

## [v0.28.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.28.0-victorialogs)

Released at 2024-07-10

* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): show a spinner on top of bar chart until user's request is finished. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6558).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): use compact representation of JSON lines at `JSON` tab if only a single [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) is queried. See [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6559).
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): properly show the number of matching logs on the selected time range at bar chart for queries with arbitrary [pipes](https://docs.victoriametrics.com/victorialogs/logsql/#pipes), including [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) and [`top` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe).

## [v0.27.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.27.1-victorialogs)

Released at 2024-07-05

* BUGFIX: properly JSON-encode strings with special chars in [HTTP querying API](https://docs.victoriametrics.com/victorialogs/querying/#http-api) responses. This fixes the `error decode response: invalid character 'x' in string escape code` error in [VictoriaLogs datasource for Grafana](https://github.com/VictoriaMetrics/victorialogs-datasource/). See [this issue](https://github.com/VictoriaMetrics/victorialogs-datasource/issues/24). The issue has been introduced in the release [v0.9.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.9.0-victorialogs).

## [v0.27.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.27.0-victorialogs)

Released at 2024-07-02

* FEATURE: add `-syslog.useLocalTimestamp.tcp` and `-syslog.useLocalTimestamp.udp` command-line flags, which could be used for using the local timestamp as [`_time` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field) for the logs ingested via the corresponding `-syslog.listenAddr.tcp` / `-syslog.listenAddr.udp`. By default, the timestamp from the syslog message is used as [`_time` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field). See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/).

* BUGFIX: make slowly ingested logs visible for search as soon as they are ingested into VictoriaLogs. Previously slowly ingested logs could remain invisible for search for long time.

## [v0.26.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.26.1-victorialogs)

Released at 2024-07-01

* BUGFIX: return the proper surrounding logs for [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe) when additional [pipes](https://docs.victoriametrics.com/victorialogs/logsql/#pipes) are put after the `stream_context` pipe. This has been broken in [v0.26.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.26.0-victorialogs).

## [v0.26.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.26.0-victorialogs)

Released at 2024-07-01

* FEATURE: add ability to return log position (aka rank) after sorting logs with [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe). This can be done by adding `rank as <fieldName>` to the end of `| sort ...` pipe. For example, `_time:5m | sort by (_time) rank as position` instructs storing position of every sorted log line into `position` [field name](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: add delimiter log with `---` message between log chunks returned by [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe). This should simplify investigation of the returned logs.
* FEATURE: reduce memory usage when big number of context logs are requested from [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe).

## [v0.25.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.25.0-victorialogs)

Released at 2024-06-28

* FEATURE: add ability to select surrounding logs in front and after the selected logs via [`stream_context` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stream_context-pipe). This functionality may be useful for investigating stacktraces, panics or some correlated log messages. This functionality is similar to `grep -A` and `grep -B`.
* FEATURE: add ability to return top `N` `"fields"` groups from  [`/select/logsql/hits` HTTP endpoint](https://docs.victoriametrics.com/victorialogs/querying/#querying-hits-stats), by specifying `fields_limit=N` query arg. This query arg is going to be used in [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6545).

* BUGFIX: fix `runtime error: index out of range [0] with length 0` panic when empty lines are ingested via [Syslog format](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/) by Cisco controllers. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6548).

## [v0.24.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.24.0-victorialogs)

Released at 2024-06-27

* FEATURE: add `/select/logsql/tail` HTTP endpoint, which can be used for live tailing of [LogsQL query](https://docs.victoriametrics.com/victorialogs/logsql/) results. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#live-tailing) for details.
* FEATURE: add `/select/logsql/stream_ids` HTTP endpoint, which can be used for returning [`_stream_id` values](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) with the number of hits for the given [LogsQL query](https://docs.victoriametrics.com/victorialogs/logsql/). See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#querying-stream_ids) for details.
* FEATURE: add `-retention.maxDiskSpaceUsageBytes` command-line flag, which allows limiting disk space usage for [VictoriaLogs data](https://docs.victoriametrics.com/victorialogs/#storage) by automatic dropping the oldest per-day partitions if the storage disk space usage becomes bigger than the `-retention.maxDiskSpaceUsageBytes`. See [these docs](https://docs.victoriametrics.com/victorialogs/#retention-by-disk-space-usage).

* BUGFIX: properly take into account query timeout specified via `-search.maxQueryDuration` command-line flag and/or via `timeout` query arg. Previously these timeouts could be ignored during query execution.
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix the update of the relative time range when `Execute Query` is clicked. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6345).

## [v0.23.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.23.0-victorialogs)

Released at 2024-06-25

* FEATURE: [syslog data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/): parse [STRUCTURED-DATA](https://datatracker.ietf.org/doc/html/rfc5424#section-6.3) into `SD-ID.field1=value1`, `SD-ID.field2=value2`, ..., `SD-ID.fieldN=valueN` [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). Previously the `STRUCTURED-DATA` was parsed into a single log field with the `SD-ID` name and `field1=value1 field2=value2 ... fieldN=valueN` value. This could complicate querying of such data.

* BUGFIX: properly parse timestamps with timezones during [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/) and [querying](https://docs.victoriametrics.com/victorialogs/querying/). This has been broken in [v0.20.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.20.0-victorialogs). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6508).

## [v0.22.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.22.0-victorialogs)

Released at 2024-06-24

* FEATURE: allow specifying multiple `_stream_id` values in [`_stream_id` filter](https://docs.victoriametrics.com/victorialogs/logsql/#_stream_id-filter) via `_stream_id:in(id1, ..., idN)` syntax.
* FEATURE: allow specifying subquery for searching for `_stream_id` values inside [`_stream_id` filter](https://docs.victoriametrics.com/victorialogs/logsql/#_stream_id-filter). For example, `_stream_id:in(_time:5m error | fields _stream_id)` returns logs for [logs streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) with the `error` word across logs for the last 5 minutes.

## [v0.21.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.21.0-victorialogs)

Released at 2024-06-20

* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add a bar chart displaying the number of log entries over a time range. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6404).
* FEATURE: expose `_stream_id` field, which uniquely identifies [log streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields). This field can be used for quick obtaining of all the logs belonging to a particular stream via [`_stream_id` filter](https://docs.victoriametrics.com/victorialogs/logsql/#_stream_id-filter).

## [v0.20.2](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.20.2-victorialogs)

Released at 2024-06-18

* BUGFIX: properly parse timestamps with nanosecond precision for logs ingested via [jsonline format](https://docs.victoriametrics.com/victorialogs/data-ingestion/#json-stream-api). The bug has been introduced in [v0.20.0 release](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.20.0-victorialogs).

## [v0.20.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.20.1-victorialogs)

Released at 2024-06-18

* FEATURE: allow configuring multiple receivers with distinct configs for syslog messages. See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#multiple-configs).

* BUGFIX: properly read syslog messages over TCP and TLS connections according to [RFC5425](https://datatracker.ietf.org/doc/html/rfc5425) when [data ingestion for syslog protocol](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/) is enabled.

## [v0.20.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.20.0-victorialogs)

Released at 2024-06-17

* FEATURE: add ability to accept logs in [Syslog format](https://en.wikipedia.org/wiki/Syslog). See [these docs](https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/).
* FEATURE: add ability to specify timezone offset when parsing [rfc3164](https://datatracker.ietf.org/doc/html/rfc3164) syslog messages with [`unpack_syslog` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_syslog-pipe).
* FEATURE: add [`top` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#top-pipe) for returning top N sets of the given fields with the maximum number of matching log entries.

## [v0.19.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.19.0-victorialogs)

Released at 2024-06-11

* FEATURE: do not allow starting the [filter](https://docs.victoriametrics.com/victorialogs/logsql/#filters) with [pipe names](https://docs.victoriametrics.com/victorialogs/logsql/#pipes) and [stats function names](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe-functions). This prevents from unexpected results returned by incorrect queries, which miss mandatory [filter](https://docs.victoriametrics.com/victorialogs/logsql/#query-syntax).
* FEATURE: treat unexpected syslog message as [RFC3164](https://datatracker.ietf.org/doc/html/rfc3164) containing only the `message` field when using [`unpack_syslog` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_syslog-pipe).
* FEATURE: allow using `where` prefix instead of `filter` prefix in [`filter` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#filter-pipe).
* FEATURE: disallow unescaped `!` char in [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/) queries, since it permits writing incorrect query, which may look like correct one. For example, `foo!:bar` instead of `foo:!bar`.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): add markdown support to the `Group` view. See [this pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/6292).

* BUGFIX: return back the improved performance for queries with `*` filters (aka `SELECT *`). This has been broken in [v0.16.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.16.0-victorialogs).

## [v0.18.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.18.0-victorialogs)

Released at 2024-06-06

* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): improve displaying of logs. See [this pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/6419) and the following issues: [6408](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6408), [6405](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6405), [6406](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6406) and [6407](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6407).
* FEATURE: add support for [day range filter](https://docs.victoriametrics.com/victorialogs/logsql/#day-range-filter) and [week range filter](https://docs.victoriametrics.com/victorialogs/logsql/#week-range-filter). These filters allow selecting logs on a particular time range per every day or on a particular day of the week.
* FEATURE: allow using `eval` instead of `math` keyword in [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe).

## [v0.17.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.17.0-victorialogs)

Released at 2024-06-05

* FEATURE: add [`pack_logfmt` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#pack_logfmt-pipe) for formatting [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) into [logfmt](https://brandur.org/logfmt) messages.
* FEATURE: allow using IPv4 addresses in [range comparison filters](https://docs.victoriametrics.com/victorialogs/logsql/#range-comparison-filter). For example, `ip:>'12.34.56.78'` is valid filter now.
* FEATURE: add `ceil()` and `floor()` functions to [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe).
* FEATURE: add support for bitwise `and`, `or` and `xor` operations at [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe).
* FEATURE: add support for automatic conversion of [RFC3339 time](https://www.rfc-editor.org/rfc/rfc3339) and IPv4 addresses into numeric representation at [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe).
* FEATURE: add ability to format numeric fields into string representation of time, duration and IPv4 with [`format` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe).
* FEATURE: set `format` field to `rfc3164` or `rfc5424` depending on the [Syslog format](https://en.wikipedia.org/wiki/Syslog) parsed via [`unpack_syslog` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_syslog-pipe).

* BUGFIX: always respect the limit set in [`limit` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#limit-pipe). Previously the limit could be exceeded in some cases.

## [v0.16.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.16.0-victorialogs)

Released at 2024-06-04

* FEATURE: add [`unpack_syslog` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_syslog-pipe) for unpacking [syslog](https://en.wikipedia.org/wiki/Syslog) messages from [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: parse timestamps in [`_time` filter](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter) with nanosecond precision.
* FEATURE: return the last `N` matching logs from [`/select/logsql/query` HTTP API](https://docs.victoriametrics.com/victorialogs/querying/#querying-logs) with the maximum timestamps if `limit=N` query arg is passed to it. Previously a random subset of matching logs could be returned, which could complicate investigation of the returned logs.
* FEATURE: add [`drop_empty_fields` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#drop_empty_fields-pipe) for dropping [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) with empty values.

## [v0.15.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.15.0-victorialogs)

Released at 2024-05-30

* FEATURE: add [`row_any`](https://docs.victoriametrics.com/victorialogs/logsql/#row_any-stats) function for [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe). This function returns a sample log entry per every calculated [group of results](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-fields).
* FEATURE: add `default` operator to [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe). It allows overriding `NaN` results with the given default value.
* FEATURE: add `exp()` and `ln()` functions to [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe).
* FEATURE: allow omitting result name in [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe) expressions. In this case the result name is automatically set to string representation of the corresponding math expression. For example, `_time:5m | math duration / 1000` is equivalent to `_time:5m | math (duration / 1000) as "duration / 1000"`.
* FEATURE: allow omitting result name in [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe). In this case the result name is automatically set to string representation of the corresponding [stats function expression](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe-functions). For example, `_time:5m | count(*)` is valid [LogsQL query](https://docs.victoriametrics.com/victorialogs/logsql/) now. It is equivalent to `_time:5m | stats count(*) as "count(*)"`.

* BUGFIX: properly calculate the number of matching rows in `* | field_values x | stats count() rows` and in `* | unroll (x) | stats count() rows` queries.

## [v0.14.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.14.0-victorialogs)

Released at 2024-05-29

* FEATURE: allow specifying fields, which must be packed into JSON in [`pack_json` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#pack_json-pipe) via `pack_json fields (field1, ..., fieldN)` syntax.

* BUGFIX: properly apply `if (...)` filters to calculated results in [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) when [grouping by fields](https://docs.victoriametrics.com/victorialogs/logsql/#stats-by-fields) is enabled. For example, `_time:5m | stats by (host) count() logs, count() if (error) errors` now properly calculates per-`host` `errors`.

## [v0.13.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.13.0-victorialogs)

Released at 2024-05-28

* FEATURE: add [`extract_regexp` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#extract_regexp-pipe) for extracting arbitrary substrings from [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) with [RE2 regular expressions](https://github.com/google/re2/wiki/Syntax).
* FEATURE: add [`math` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#math-pipe) for mathematical calculations over [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: add [`field_values` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#field_values-pipe), which returns unique values for the given [log field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: allow omitting `stats` prefix in [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe). For example, `_time:5m | count() rows` is a valid query now. It is equivalent to `_time:5m | stats count() as rows`.
* FEATURE: allow omitting `filter` prefix in [`filter` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#filter-pipe) if the filter doesn't clash with [pipe names](#https://docs.victoriametrics.com/victorialogs/logsql/#pipes). For example, `_time:5m | stats by (host) count() rows | rows:>1000` is a valid query now. It is equivalent to `_time:5m | stats by (host) count() rows | filter rows:>1000`.
* FEATURE: allow [`head` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#limit-pipe) without number. For example, `error | head`. In this case 10 first values are returned as `head` Unix command does by default.
* FEATURE: allow using [comparison filters](https://docs.victoriametrics.com/victorialogs/logsql/#range-comparison-filter) with strings. For example, `some_text_field:>="foo"` matches [log entries](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) with `some_text_field` field values bigger or equal to `foo`.

## [v0.12.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.12.1-victorialogs)

Released at 2024-05-26

* FEATURE: add support for comments in multi-line LogsQL queries. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#comments).

* BUGFIX: properly apply [`in(...)` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter) inside `if (...)` conditions at various [pipes](https://docs.victoriametrics.com/victorialogs/logsql/#pipes). This bug has been introduced in [v0.12.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.12.0-victorialogs).

## [v0.12.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.12.0-victorialogs)

Released at 2024-05-26

* FEATURE: add [`pack_json` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#pack_json-pipe), which packs all the [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) into a JSON object and stores it into the given field.
* FEATURE: add [`unroll` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unroll-pipe), which can be used for unrolling JSON arrays stored in [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: add [`replace_regexp` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#replace_regexp-pipe), which allows updating [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) with regular expressions.
* FEATURE: improve performance for [`format`](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe) and [`extract`](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe) pipes.
* FEATURE: improve performance for [`/select/logsql/field_names` HTTP API](https://docs.victoriametrics.com/victorialogs/querying/#querying-field-names).

* BUGFIX: prevent from panic in [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe) when VictoriaLogs runs on a system with one CPU core.
* BUGFIX: do not return referenced fields if they weren't present in the original logs. For example, `_time:5m | format if (non_existing_field:"") "abc"` could return empty `non_exiting_field`, while it shouldn't be returned because it is missing in the original logs.
* BUGFIX: properly initialize values for [`in(...)` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter) inside [`filter` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#filter-pipe) if the `in(...)` contains other [filters](https://docs.victoriametrics.com/victorialogs/logsql/#filters). For example, `_time:5m | filter ip:in(user_type:admin | fields ip)` now works correctly.

## [v0.11.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.11.0-victorialogs)

Released at 2024-05-25

* FEATURE: add [`replace` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#replace-pipe), which allows replacing substrings in [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model).
* FEATURE: support [comparing](https://docs.victoriametrics.com/victorialogs/logsql/#range-filter) log field values with [special numeric values](https://docs.victoriametrics.com/victorialogs/logsql/#numeric-values). For example, `duration:>1.5s` and `response_size:<15KiB` are valid filters now.
* FEATURE: properly sort [durations](https://docs.victoriametrics.com/victorialogs/logsql/#duration-values) and [short numeric values](https://docs.victoriametrics.com/victorialogs/logsql/#short-numeric-values) in [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe). For example, `10s` goes in front of `1h`, while `10KB` goes in front of `1GB`.
* FEATURE: add an ability to preserve the original non-empty field values when executing [`extract`](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe), [`unpack_json`](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_json-pipe), [`unpack_logfmt`](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_logfmt-pipe) and [`format`](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe) pipes.
* FEATURE: add an ability to preserve the original field values if the corresponding unpacked values are empty when executing [`extract`](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe), [`unpack_json`](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_json-pipe), [`unpack_logfmt`](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_logfmt-pipe) and [`format`](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe) pipes.

## [v0.10.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.10.0-victorialogs)

Released at 2024-05-24

* FEATURE: return the number of matching log entries per returned value in [HTTP API](https://docs.victoriametrics.com/victorialogs/querying/#http-api) results. This simplifies detecting [field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) / [stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) values with the biggest number of logs for the given [LogsQL query](https://docs.victoriametrics.com/victorialogs/logsql/).
* FEATURE: improve performance for [regexp filter](https://docs.victoriametrics.com/victorialogs/logsql/#regexp-filter) in the following cases:
  * If the regexp contains just a phrase without special regular expression chars. For example, `~"foo"`.
  * If the regexp starts with `.*` or ends with `.*`. For example, `~".*foo.*"`.
  * If the regexp contains multiple strings delimited by `|`. For example, `~"foo|bar|baz"`.
  * If the regexp contains multiple [words](https://docs.victoriametrics.com/victorialogs/logsql/#word). For example, `~"foo bar baz"`.
* FEATURE: allow disabling automatic unquoting of the matched placeholders in [`extract` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe). See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#format-for-extract-pipe-pattern).

* BUGFIX: properly parse `!` in front of [exact filter](https://docs.victoriametrics.com/victorialogs/logsql/#exact-filter), [exact-prefix filter](https://docs.victoriametrics.com/victorialogs/logsql/#exact-prefix-filter) and [regexp filter](https://docs.victoriametrics.com/victorialogs/logsql/#regexp-filter). For example, `!~"some regexp"` is properly parsed as `not ="some regexp"`. Previously it was incorrectly parsed as `'~="some regexp"'` [phrase filter](https://docs.victoriametrics.com/victorialogs/logsql/#phrase-filter).
* BUGFIX: properly sort results by [`_time` field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field) when [`limit` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#limit-pipe) is applied. For example, `_time:5m | sort by (_time) desc | limit 10` properly works now.

## [v0.9.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.9.1-victorialogs)

Released at 2024-05-22

* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix loading web UI, which has been broken in [v0.9.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.9.0-victorialogs).

## [v0.9.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.9.0-victorialogs)

Released at 2024-05-22

* FEATURE: allow using `~"some_regexp"` [regexp filter](https://docs.victoriametrics.com/victorialogs/logsql/#regexp-filter) instead of `re("some_regexp")`.
* FEATURE: allow using `="some phrase"` [exact filter](https://docs.victoriametrics.com/victorialogs/logsql/#exact-filter) instead of `exact("some phrase")`.
* FEATURE: allow using `="some prefix"*` [exact prefix filter](https://docs.victoriametrics.com/victorialogs/logsql/#exact-prefix-filter) instead of `exact("some prefix"*)`.
* FEATURE: add ability to generate output fields according to the provided format string. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#format-pipe).
* FEATURE: add ability to extract fields with [`extract` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe) only if the given condition is met. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#conditional-extract).
* FEATURE: add ability to unpack JSON fields with [`unpack_json` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_json-pipe) only if the given condition is met. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#conditional-unpack_json).
* FEATURE: add ability to unpack [logfmt](https://brandur.org/logfmt) fields with [`unpack_logfmt` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_logfmt-pipe) only if the given condition is met. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#conditional-unpack_logfmt).
* FEATURE: add [`row_min`](https://docs.victoriametrics.com/victorialogs/logsql/#row_min-stats) and [`row_max`](https://docs.victoriametrics.com/victorialogs/logsql/#row_max-stats) functions for [`stats` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe), which allow returning all the [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) for the log entry with the minimum / maximum value at the given field.
* FEATURE: add `/select/logsql/streams` HTTP endpoint for returning [streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) from results of the given query. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#querying-streams) for details.
* FEATURE: add `/select/logsql/stream_field_names` HTTP endpoint for returning [stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) field names from results of the given query. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#querying-stream-field-names) for details.
* FEATURE: add `/select/logsql/stream_field_values` HTTP endpoint for returning [stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) field values for the given label from results of the given query. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#querying-stream-field-values) for details.
* FEATURE: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): change time range limitation from `_time` in the expression to `start` and `end` query args.

* BUGFIX: fix `invalid memory address or nil pointer dereference` panic when using [`extract`](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe), [`unpack_json`](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_json-pipe) or [`unpack_logfmt`](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_logfmt-pipe) pipes. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6306).
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): fix an issue where logs with long `_msg` values might not display. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6281).
* BUGFIX: properly handle time range boundaries with millisecond precision. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6293).

## [v0.8.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.8.0-victorialogs)

Released at 2024-05-20

* FEATURE: add ability to extract JSON fields from [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_json-pipe).
* FEATURE: add ability to extract [logfmt](https://brandur.org/logfmt) fields from [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#unpack_logfmt-pipe).
* FEATURE: add ability to extract arbitrary text from [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) into the output fields. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#extract-pipe).
* FEATURE: add ability to put arbitrary [queries](https://docs.victoriametrics.com/victorialogs/logsql/#query-syntax) inside [`in()` filter](https://docs.victoriametrics.com/victorialogs/logsql/#multi-exact-filter).
* FEATURE: add support for post-filtering of query results with [`filter` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#filter-pipe).
* FEATURE: allow applying individual [filters](https://docs.victoriametrics.com/victorialogs/logsql/#filters) per each [stats function](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe-functions). See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#stats-with-additional-filters).
* FEATURE: allow passing string values to [`min`](https://docs.victoriametrics.com/victorialogs/logsql/#min-stats) and [`max`](https://docs.victoriametrics.com/victorialogs/logsql/#max-stats) functions. Previously only numeric values could be passed to them.
* FEATURE: speed up [`sort ... limit N` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe) for typical cases.
* FEATURE: allow using more convenient syntax for [`range` filters](https://docs.victoriametrics.com/victorialogs/logsql/#range-filter) if upper or lower bound isn't needed. For example, it is possible to write `response_size:>=10KiB` instead of `response_size:range[10KiB, inf)`, or `temperature:<42` instead of `temperature:range(-inf, 42)`.
* FEATURE: add `/select/logsql/hits` HTTP endpoint for returning the number of matching logs per the given time bucket over the selected time range. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#querying-hits-stats) for details.
* FEATURE: add `/select/logsql/field_names` HTTP endpoint for returning [field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) names from results of the given query. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#querying-field-names) for details.
* FEATURE: add `/select/logsql/field_values` HTTP endpoint for returning unique values for the given [field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) obtained from results of the given query. See [these docs](https://docs.victoriametrics.com/victorialogs/querying/#querying-field-values) for details.

* BUGFIX: properly take into account `offset` at [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe) when it already has `limit`. For example, `_time:5m | sort by (foo) offset 20 limit 10`.

## [v0.7.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.7.0-victorialogs)

Released at 2024-05-15

* FEATURE: add support for optional `start` and `end` query args to [HTTP querying API](https://docs.victoriametrics.com/victorialogs/querying/#http-api), which can be used for limiting the time range for [LogsQL query](https://docs.victoriametrics.com/victorialogs/logsql/).
* FEATURE: add ability to return the first `N` results from [`sort` pipe](#https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe). This is useful when `N` biggest or `N` smallest values must be returned from large amounts of logs.
* FEATURE: add [`quantile`](https://docs.victoriametrics.com/victorialogs/logsql/#quantile-stats) and [`median`](https://docs.victoriametrics.com/victorialogs/logsql/#median-stats) [stats functions](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe).

## [v0.6.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.6.1-victorialogs)

Released at 2024-05-14

* FEATURE: use [natural sort order](https://en.wikipedia.org/wiki/Natural_sort_order) when sorting logs via [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe).

* BUGFIX: properly return matching logs in [streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) with small number of entries. Previously they could be skipped. The issue has been introduced in [the release v0.6.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.6.0-victorialogs).
* BUGFIX: fix `runtime error: index out of range` panic when using [`sort` pipe](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe) like `_time:1h | sort by (_time)`. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6258).

## [v0.6.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.6.0-victorialogs)

Released at 2024-05-12

* FEATURE: return all the log fields by default in query results. Previously only [`_stream`](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields), [`_time`](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field) and [`_msg`](https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field) fields were returned by default.
* FEATURE: add support for returning only the requested log [fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#fields-pipe).
* FEATURE: add support for calculating various stats over [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model). Grouping by arbitrary set of [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) is supported. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#stats-pipe) for details.
* FEATURE: add support for sorting the returned results. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#sort-pipe).
* FEATURE: add support for returning unique results. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#uniq-pipe).
* FEATURE: add support for limiting the number of returned results. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#limiters).
* FEATURE: add support for copying and renaming the selected log fields. See [these](https://docs.victoriametrics.com/victorialogs/logsql/#copy-pipe) and [these](https://docs.victoriametrics.com/victorialogs/logsql/#rename-pipe) docs.
* FEATURE: allow using `_` inside numbers. For example, `score:range[1_000, 5_000_000]` for [`range` filter](https://docs.victoriametrics.com/victorialogs/logsql/#range-filter).
* FEATURE: allow numbers in hexadecimal and binary form. For example, `response_size:range[0xff, 0b10001101101]` for [`range` filter](https://docs.victoriametrics.com/victorialogs/logsql/#range-filter).
* FEATURE: allow using duration and byte size suffixes in numeric values inside LogsQL queries. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#numeric-values).
* FEATURE: improve data ingestion performance by up to 50%.
* FEATURE: optimize performance for [LogsQL query](https://docs.victoriametrics.com/victorialogs/logsql/), which contains multiple filters for [words](https://docs.victoriametrics.com/victorialogs/logsql/#word-filter) or [phrases](https://docs.victoriametrics.com/victorialogs/logsql/#phrase-filter) delimited with [`AND` operator](https://docs.victoriametrics.com/victorialogs/logsql/#logical-filter). For example, `foo AND bar` query must find [log messages](https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field) with `foo` and `bar` words at faster speed.

* BUGFIX: prevent from possible corruption of short [log fields](https://docs.victoriametrics.com/victorialogs/keyconcepts/#data-model) during data ingestion.
* BUGFIX: prevent from additional CPU usage for up to a few seconds after canceling the query.
* BUGFIX: prevent from returning log entries with empty `_stream` field in the form `"_stream":""` in [search query results](https://docs.victoriametrics.com/victorialogs/querying/). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6042).

## [v0.5.2](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.5.2-victorialogs)

Released at 2024-04-11

* BUGFIX: properly register new [log streams](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) under high data ingestion rate. The issue has been introduced in [v0.5.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.5.0-victorialogs).

## [v0.5.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.5.1-victorialogs)

Released at 2024-04-04

* BUGFIX: properly apply time range filter for queries containing [`OR` operators](https://docs.victoriametrics.com/victorialogs/logsql/#logical-filter). See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/5920).
* BUGFIX: do not log debug lines `DEBUG: start trimLines` and `DEBUG: end trimLines`. This bug has been introduced in [v0.5.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.5.0-victorialogs) in [this commit](https://github.com/VictoriaMetrics/VictoriaMetrics/commit/0514091948cf8e00e42f44318c0e5e5b63b6388f).

## [v0.5.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.5.0-victorialogs)

Released at 2024-03-01

* FEATURE: support the ability to limit the number of returned log entries from [HTTP querying API](https://docs.victoriametrics.com/victorialogs/querying/#http-api) by passing `limit` query arg. Previously all the matching log entries were returned until closing the response stream. See [this feature request](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/5674). Thanks to @dmitryk-dk for [the pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/5778).

* BUGFIX: do not panic on incorrect regular expression in [stream filter](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter). Thanks to @XLONG96 for [the bugfix](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/5897).
* BUGFIX: properly determine when the assisted merge is needed. Previously the logs for determining whether the assisted merge is needed was broken. This could lead to too big number of parts under high data ingestion rate. Thanks to @lujiajing1126 for [the fix](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/5447).
* BUGFIX: properly stop execution of aborted query when the query doesn't contain [`_stream` filter](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter). Previously such a query could continue consuming resources after being aborted by the client. Thanks to @z-anshun for [the fix](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/5400).

## [v0.4.2](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.4.2-victorialogs)

Released at 2023-11-15

* BUGFIX: properly locate logs for the [requested streams](https://docs.victoriametrics.com/victorialogs/logsql/#stream-filter). Previously logs for some streams may be missing in query results. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4856). Thanks to @XLONG96 for [the fix](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/5295)!
* BUGFIX: [web UI](https://docs.victoriametrics.com/victorialogs/querying/#web-ui): properly sort found logs by time. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/5300).

## [v0.4.1](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.4.1-victorialogs)

Released at 2023-10-04

* BUGFIX: fix the free space verification process in VictoriaLogs that was erroneously shifting to read-only mode, despite there being sufficient free space available. See [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/5112) issue.

## [v0.4.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.4.0-victorialogs)

Released at 2023-10-03

* FEATURE: add `-elasticsearch.version` command-line flag, which can be used for specifying Elasticsearch version returned by VictoriaLogs to Filebeat at [elasticsearch bulk API](https://docs.victoriametrics.com/victorialogs/data-ingestion/#elasticsearch-bulk-api). This helps resolving [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4777).
* FEATURE: expose the following metrics at [/metrics](https://docs.victoriametrics.com/victorialogs/#monitoring) page:
  * `vl_data_size_bytes{type="storage"}` - on-disk size for data excluding [log stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) indexes.
  * `vl_data_size_bytes{type="indexdb"}` - on-disk size for [log stream](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) indexes.
* FEATURE: add `-insert.maxFieldsPerLine` command-line flag, which can be used for limiting the number of fields per line in logs sent to VictoriaLogs via ingestion protocols. This helps to avoid issues like [this](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4762).
* FEATURE: expose `vl_http_request_duration_seconds` histogram at the [/metrics](https://docs.victoriametrics.com/victorialogs/#monitoring) page. Thanks to @crossoverJie for [this pull request](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/4934).
* FEATURE: add support of `-storage.minFreeDiskSpaceBytes` command-line flag to allow switching to read-only mode when running out of disk space at `-storageDataPath`. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4737).

* BUGFIX: fix possible panic when no data is written to VictoriaLogs for a long time. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4895). Thanks to @crossoverJie for filing and fixing the issue.
* BUGFIX: add `/insert/loki/ready` endpoint, which is used by Promtail for healthchecks. This should remove `unsupported path requested: /insert/loki/ready` warning logs. See [this comment](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4762#issuecomment-1690966722).
* BUGFIX: prevent from panic during background merge when the number of columns in the resulting block exceeds the maximum allowed number of columns per block. See [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/4762).

## [v0.3.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.3.0-victorialogs)

Released at 2023-07-20

* FEATURE: add support for data ingestion via Promtail (aka default log shipper for Grafana Loki). See [these](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/) and [these](https://docs.victoriametrics.com/victorialogs/data-ingestion/#loki-json-api) docs.

## [v0.2.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.2.0-victorialogs)

Released at 2023-07-17

* FEATURE: support short form of `_time` filters over the last X minutes/hours/days/etc. For example, `_time:5m` is a short form for `_time:(now-5m, now]`, which matches logs with [timestamps](https://docs.victoriametrics.com/victorialogs/keyconcepts/#time-field) for the last 5 minutes. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter) for details.
* FEATURE: add ability to specify offset for the selected time range. For example, `_time:5m offset 1h` is equivalent to `_time:(now-5m-1h, now-1h]`. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#time-filter) for details.
* FEATURE: [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/): replace `exact_prefix("...")` with `exact("..."*)`. This makes it consistent with [i()](https://docs.victoriametrics.com/victorialogs/logsql/#case-insensitive-filter) filter, which can accept phrases and prefixes, e.g. `i("phrase")` and `i("phrase"*)`. See [these docs](https://docs.victoriametrics.com/victorialogs/logsql/#exact-prefix-filter).

## [v0.1.0](https://github.com/VictoriaMetrics/VictoriaMetrics/releases/tag/v0.1.0-victorialogs)

Released at 2023-06-21

Initial release
