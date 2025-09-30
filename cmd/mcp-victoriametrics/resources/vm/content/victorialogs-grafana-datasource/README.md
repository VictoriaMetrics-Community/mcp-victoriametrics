

The [VictoriaLogs Grafana plugin](https://grafana.com/grafana/plugins/victoriametrics-logs-datasource/) allows Grafana 
to query, visualize, and interact with [VictoriaLogs](https://docs.victoriametrics.com/victorialogs/),
a high-performance log storage and processing system.

<img alt="Grafana Dashboard Screenshot" width="100%" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/dashboard.png?raw=true">

## Capabilities

1. Use [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/) to filter, aggregate, and transform logs data to gain insights into application behavior.
1. Use Explore mode with Grafana.
1. Show live-streaming logs.
1. Build dashboards and setup alerts.
1. Use Ad Hoc filters.

Try it at [VictoriaMetrics playground](https://play-grafana.victoriametrics.com/d/be5zidev72m80f/k8s-logs-demo)!

## Installation

For detailed instructions on how to install the plugin on Grafana Cloud or locally, please checkout the [Plugin installation docs](https://grafana.com/docs/grafana/latest/plugins/installation/).
For installation options in Docker or Kubernetes refer to [these docs](https://github.com/VictoriaMetrics/victorialogs-datasource?tab=readme-ov-file#installation).

### Manual configuration via UI

Once the plugin is installed on your Grafana instance, follow [these instructions](https://grafana.com/docs/grafana/latest/datasources/add-a-data-source/)
to add a new VictoriaLogs data source, and enter configuration options.

### Configuration via file

Provisioning of Grafana plugin requires creating [datasource config file](http://docs.grafana.org/administration/provisioning/#datasources):

```yaml
apiVersion: 1
datasources:
  - name: VictoriaLogs
    type: victoriametrics-logs-datasource
    access: proxy
    url: http://victorialogs:9428
    isDefault: true
```

## Building queries

VictoriaLogs query language is [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/).
Queries can be built using raw LogsQL or via QueryBuilder.

See panels examples at [VictoriaMetrics playground](https://play-grafana.victoriametrics.com/d/be5zidev72m80f/k8s-logs-demo)
and LogsQL examples [here](https://docs.victoriametrics.com/victorialogs/logsql-examples/).

### Logs panel

For using [Logs panel](https://grafana.com/docs/grafana/latest/panels-visualizations/visualizations/logs/)
switch to `Raw Logs` query type:

<img alt="Logs panel" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/panel_logs.png?raw=true">

### Time series panel

For using [Time series panel](https://grafana.com/docs/grafana/latest/panels-visualizations/visualizations/time-series/)
switch to `Range` query type:

<img alt="Time series panel" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/panel_time_series.png?raw=true">

### Stats panel

For using [Stats panel](https://grafana.com/docs/grafana/latest/panels-visualizations/visualizations/stat/)
switch to `Instant` query type:

<img alt="Stats panel" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/panel_stat.png?raw=true">

For enabling background visualization switch to `Range` query type.

### Table panel

For using [Table panel](https://grafana.com/docs/grafana/latest/panels-visualizations/visualizations/table/)
switch to `Raw Logs` query type:

<img alt="Table panel" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/panel_table.png?raw=true">

And apply `Transformations` by labels:

<img alt="Transformations" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/panel_table_transformation.png?raw=true">

### Log level rules

The **Log level rules** section in the datasource configuration allows you to assign log levels based on custom field conditions. This helps classify logs dynamically (e.g., as `error`, `info`, `debug`, etc.) using rules you define.

#### How to use

1. Open the datasource settings.

2. Scroll to the **Log level rules** section.

3. Click **"Add rule"** to define a new rule.

4. For each rule, configure the following:

    * **Enable switch** – enable or disable the rule.
    * **Field name** – the log field the condition will evaluate.
    * **Operator** – choose from: `Equals`, `Not equal`, `Matches regex`, `Less than`, `Greater than`
    * **Value** – the value to compare the field against.
    * **Log level** – level to assign if the condition matches: `critical`, `warning`, `error`, `info`, `debug`, `trace`, `unknown`
    * **Delete button** – remove the rule.

5.  After adding or editing rules, click **"Save & test"** to apply the changes.

**Rule priority**: If multiple rules match a log entry, the **first matching rule** (top to bottom) takes precedence.

6. To define rules via the provision file, use the following format of the provision file:

```yaml
apiVersion: 1
datasources:
  # <string, required> Name of the VictoriaLogs datasource
  # displayed in Grafana panels and queries.
  - name: VictoriaLogs
    # <string, required> Sets the data source type.
    type: victoriametrics-logs-datasource
    # <string, required> Sets the access mode, either
    # proxy or direct (Server or Browser in the UI).
    access: proxy
    # <string> Sets URL for sending queries to VictoriaLogs server.
    # see https://docs.victoriametrics.com/victorialogs/querying/
    url: https://play-vmlogs.victoriametrics.com
    # <string> Sets the pre-selected datasource for new panels.
    # You can set only one default data source per organization.
    isDefault: true
    jsonData:
      logLevelRules:
        - field: "_stream_id"
          value: "123123"
          level: "error"
          operator: "regex"
          enabled: true
```
Where:
* `field` is the name of the log field to evaluate (e.g. `_stream_id`, `status_code`, `message`).
* `value` is the value to compare against. Can be a string or a number, depending on the field.
* `level` is the log level to assign if the condition matches. Valid values: `critical`, `error`, `warning`, `info`, `debug`, `trace`, `unknown`.
* `operator` is the comparison operator to use, such as `equals`, `notEquals`, `regex`, `lessThan`, `greaterThan`.
* `enabled` is a boolean flag to enable or disable the rule. Defaults to `true` if omitted.

## Correlations

### Trace to logs

Tempo, Jaeger, and Zipkin data sources support [Trace to logs](https://grafana.com/docs/grafana/latest/explore/trace-integration/#trace-to-logs)
feature for navigating from a span in a trace directly to logs relevant for that span. _(Supported since Grafana v12.2.0)_.

<img alt="Derived fields configuration" width="100%" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/trace_to_logs.png?raw=true">

An example of the correlation query in traces datasource is the following:
```
trace_id:="${__trace.traceId}" AND span_id:="${__trace.spanId}"
```

### Log to metrics

Log to traces correlation is possible via Derived Fields functionality. But for it to work log entries and time series
in metrics datasource should share common labels that could be used as filters. See example of building logs to metrics
correlation in [this issue](https://github.com/VictoriaMetrics/VictoriaMetrics/issues/9592#issuecomment-3202104607).

### Log to traces

Log to traces correlation is possible via Derived Fields functionality. See its description in the sections below.

## Derived Fields

In VictoriaLogs datasource settings, you can configure rules of extracting values from a log message to create a link with that value.

For example, if log entries have field `trace_id` then we can configure a Derived Field to make a link to Jaeger datasource
for viewing an associated trace:

<img alt="Derived fields configuration" width="100%" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/derived_fields_cfg.png?raw=true">

Once configured, in Explore mode or in Logs panel log entries with field `trace_id` will also get a link that would
open a Jaeger datasource and search for the `trace_id` value:

<img alt="Derived fields explore" src="https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/src/img/derived_fields_explore.png?raw=true">

If the trace ID is not stored in a separate field but in a log message itself, then use `Regex in log line` option
to specify a regex expression for extracting the trace value from log message.

Learn more about [Derived Fields in Grafana](https://grafana.com/docs/grafana/next/datasources/loki/configure-loki-data-source/#derived-fields).

## License

This project is licensed under
the [Apache 2.0 license](https://github.com/VictoriaMetrics/victorialogs-datasource/blob/main/LICENSE).
