---
title : "Grafana Alloy"
menu:
  docs:
    identifier: victoriametrics-cloud-integrations-alloy
    parent: "integrations"
---

[Grafana Alloy](https://grafana.com/docs/alloy/latest/) is a vendor-neutral distribution of the
OpenTelemetry Collector, with built-in Prometheus pipelines support. It can collect and forward
metrics, logs and traces to VictoriaMetrics Cloud.

The component and endpoint to configure depend on the deployment type:

| Deployment type | Alloy component | Endpoint |
|-----------------|-----------------|----------|
| VictoriaMetrics | `prometheus.remote_write` | `<DEPLOYMENT_ENDPOINT_URL>/prometheus/api/v1/write` |
| VictoriaMetrics | `otelcol.exporter.otlphttp` | `<DEPLOYMENT_ENDPOINT_URL>/opentelemetry` |
| VictoriaLogs    | `loki.write` | `<DEPLOYMENT_ENDPOINT_URL>/insert/loki/api/v1/push` |
| VictoriaLogs    | `otelcol.exporter.otlphttp` | `<DEPLOYMENT_ENDPOINT_URL>/insert/opentelemetry` |
| VictoriaTraces  | `otelcol.exporter.otlphttp` | `<DEPLOYMENT_ENDPOINT_URL>/insert/opentelemetry` |

> [!NOTE]
> `otelcol.exporter.otlphttp` appends the signal path (`/v1/metrics`, `/v1/logs`, `/v1/traces`) to the
> configured `endpoint`, so the base URL is all that needs to be set. VictoriaMetrics Cloud exposes
> OTLP over **HTTP** only, which is why `otelcol.exporter.otlphttp` is used instead of
> `otelcol.exporter.otlp`.

## Integrating with Grafana Alloy

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, since Alloy pushes data), follow the steps in
the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To set up Grafana Alloy to forward data to VictoriaMetrics Cloud, visit the
[cloud console](https://console.victoriametrics.cloud/integrations/alloy), or follow this interactive guide:

{{< cloud-console-iframe "alloy" >}}
