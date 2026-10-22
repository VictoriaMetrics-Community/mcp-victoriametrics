---
title : "Fluent Bit"
menu:
  docs:
    parent: "integrations"
---

[Fluent Bit](https://fluentbit.io/) is a super fast, lightweight, and highly scalable logging,
metrics, and traces processor and forwarder. It can be configured to send metrics to
VictoriaMetrics Cloud using the **Prometheus remote write** protocol or OTLP, logs using the
VictoriaLogs JSON Line format, the Loki push API, or OTLP, and trace spans using OTLP.

This integration is useful for collecting container and host logs, metrics, and traces and routing
them to VictoriaMetrics Cloud for storage and analysis.

## Integrating with Fluent Bit

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, since Fluent Bit pushes data), follow the steps
in the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To set up Fluent Bit to forward data to VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/fluentbit),
or follow this interactive guide:

{{< cloud-console-iframe "fluentbit" >}}
