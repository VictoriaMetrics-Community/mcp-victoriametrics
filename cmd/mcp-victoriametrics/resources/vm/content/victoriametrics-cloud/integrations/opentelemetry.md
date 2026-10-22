---
title : "OpenTelemetry"
menu:
  docs:
    parent: "integrations"
---

VictoriaMetrics Cloud supports integration with the [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/)
for ingesting metrics, logs and traces using the OpenTelemetry Protocol (OTLP).

The endpoint to use depends on the deployment type:

| Deployment type | OTLP endpoint |
|-----------------|---------------|
| VictoriaMetrics | `<DEPLOYMENT_ENDPOINT_URL>/opentelemetry` |
| VictoriaLogs    | `<DEPLOYMENT_ENDPOINT_URL>/insert/opentelemetry/v1/logs` |
| VictoriaTraces  | `<DEPLOYMENT_ENDPOINT_URL>/insert/opentelemetry/v1/traces` |

> [!NOTE]
> VictoriaMetrics Cloud exposes OTLP over **HTTP** only. Use the `otlphttp` exporter, since the OTLP
> gRPC endpoint (port `4317`) is not available.

You can deploy the OpenTelemetry Collector using either the **Helm chart** or the **Operator**,
depending on your preference and environment, to collect, process, and
forward observability data from a wide variety of sources into VictoriaMetrics Cloud.

## Integrating with OpenTelemetry

All VictoriaMetrics Cloud integrations, including this one, require an access token for
authentication. The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>`
and `<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, as data is pushed to VictoriaMetrics Cloud),
follow the steps in the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To integrate OpenTelemetry Collector with VictoriaMetrics Cloud, visit the
[cloud console](https://console.victoriametrics.cloud/integrations/opentelemetry), or follow this interactive guide:

{{< cloud-console-iframe "opentelemetry" >}}
