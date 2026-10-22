---
title : "Grafana"
menu:
  docs:
    parent: "integrations"
    identifier: "cloud-integrations-grafana"
---

[Grafana](https://grafana.com/) is a popular open-source visualization and dashboarding tool. You can
use Grafana to query and visualize data stored in VictoriaMetrics Cloud, with a different data source
per deployment type:

| Deployment type | Grafana data source | URL |
|-----------------|---------------------|-----|
| VictoriaMetrics | Prometheus (built-in) | `<DEPLOYMENT_ENDPOINT_URL>` |
| VictoriaLogs    | [VictoriaLogs](https://grafana.com/grafana/plugins/victoriametrics-logs-datasource/) | `<DEPLOYMENT_ENDPOINT_URL>` |
| VictoriaTraces  | Tempo (built-in) | `<DEPLOYMENT_ENDPOINT_URL>/select/tempo` |

For VictoriaTraces deployments, the built-in **Jaeger** data source can be used as an alternative,
pointing it to `<DEPLOYMENT_ENDPOINT_URL>/select/jaeger`.

This integration allows you to build powerful, customizable dashboards and monitor your systems in
real time using VictoriaMetrics as the backend.

## Integrating with Grafana

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **read access**, for querying data), follow the steps in the
[Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To connect Grafana with VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/grafana),
or follow this interactive guide:

{{< cloud-console-iframe "grafana" >}}
