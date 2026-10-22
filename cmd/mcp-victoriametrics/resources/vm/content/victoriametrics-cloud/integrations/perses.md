---
title : "Perses"
menu:
  docs:
    parent: "integrations"
---
[Perses](https://perses.dev/) is an open-source visualization and dashboarding tool, designed for
simplicity, performance, and scalability.

VictoriaMetrics Cloud can be used as a data source in Perses, allowing you to create dashboards and
monitor your data with a modern and lightweight interface. The data source to configure depends on
the deployment type:

| Deployment type | Perses data source | URL | Panel and query type |
|-----------------|--------------------|-----|----------------------|
| VictoriaMetrics | Prometheus | `<DEPLOYMENT_ENDPOINT_URL>` | Time Series Chart, Prometheus Time Series Query |
| VictoriaLogs    | VictoriaLogs | `<DEPLOYMENT_ENDPOINT_URL>` | Logs Table, VictoriaLogs Log Query |
| VictoriaTraces  | Tempo | `<DEPLOYMENT_ENDPOINT_URL>/select/tempo` | Trace Table, Tempo Trace Query |

For VictoriaTraces deployments, the **Service Name** and **Span Name** fields of the Tempo Trace
Query can be left empty to fetch all traces.

## Integrating with Perses

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and `<YOUR_ACCESS_TOKEN>`.
These need to be replaced with your actual access token.

To generate your access token (with **read access**, for querying data), follow the steps in the
[Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To connect Perses with VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/perses),
or follow this interactive guide:

{{< cloud-console-iframe "perses" >}}
