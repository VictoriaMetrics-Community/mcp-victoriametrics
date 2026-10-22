---
title : "Vector"
menu:
  docs:
    identifier: victoriametrics-cloud-integrations-vector
    parent: "integrations"
---

[Vector](https://vector.dev/) is an observability data pipeline that can collect,
transform, and route metrics, logs and traces. For metrics, Vector can be configured to send data to
VictoriaMetrics Cloud using the **Prometheus remote write** protocol, or OTLP. Logs can be sent with
the Elasticsearch, Loki or OTLP sinks, and trace spans with the `opentelemetry` sink.

This integration is useful to route observability data to VictoriaMetrics Cloud for storage and
analysis.

> [!NOTE]
> For VictoriaTraces deployments, configure the `opentelemetry` sink with `protocol.type: http` and
> point its `uri` to `<DEPLOYMENT_ENDPOINT_URL>/insert/opentelemetry/v1/traces`. The OTLP gRPC
> endpoint is not exposed in VictoriaMetrics Cloud.

## Integrating with Vector

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, since Vector pushes data), follow the steps in
the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To set up Vector to forward data to VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/vector),
or follow this interactive guide:

{{< cloud-console-iframe "vector" >}}
