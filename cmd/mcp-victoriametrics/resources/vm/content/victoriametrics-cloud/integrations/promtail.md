---
title : "Promtail"
menu:
  docs:
    parent: "integrations"
---

[Promtail](https://grafana.com/docs/loki/latest/send-data/promtail/) is a logs collector agent from
Grafana Labs that ships logs to Loki-compatible endpoints. It can be configured to send logs to
VictoriaMetrics Cloud using the [Loki push API](https://docs.victoriametrics.com/victorialogs/data-ingestion/promtail/).

This integration is useful for collecting logs from files and routing them to VictoriaMetrics Cloud
for storage and analysis.

## Integrating with Promtail

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, since Promtail pushes data), follow the steps
in the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To set up Promtail to forward logs to VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/promtail),
or follow this interactive guide:

{{< cloud-console-iframe "promtail" >}}
