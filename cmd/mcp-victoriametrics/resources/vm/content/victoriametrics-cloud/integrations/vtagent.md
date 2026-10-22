---
title : "VTAgent"
menu:
  docs:
    parent: "integrations"
---

VTAgent is VictoriaMetrics' native trace collection agent for Kubernetes. Deployed as a
**StatefulSet** via Helm chart or the VictoriaMetrics operator, it accepts trace spans over OTLP and
forwards them to VictoriaTraces using the native protocol.

This integration is useful for collecting OpenTelemetry trace spans from applications and routing
them to VictoriaMetrics Cloud for storage and analysis.

## Integrating with VTAgent

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, since VTAgent pushes data), follow the steps
in the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To set up VTAgent to forward trace spans to VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/vtagent),
or follow this interactive guide:

{{< cloud-console-iframe "vtagent" >}}
