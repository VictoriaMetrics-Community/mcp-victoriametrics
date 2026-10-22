---
title : "Logstash"
menu:
  docs:
    parent: "integrations"
---

[Logstash](https://www.elastic.co/logstash/) is an open-source server-side data processing pipeline
that ingests data, transforms it, and sends it to your favorite "stash". It can be configured to send
logs to VictoriaMetrics Cloud using the [Elasticsearch bulk API](https://docs.victoriametrics.com/victorialogs/data-ingestion/logstash/).

This integration is useful for collecting and transforming logs before routing them to
VictoriaMetrics Cloud for storage and analysis.

## Integrating with Logstash

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, since Logstash pushes data), follow the steps
in the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To set up Logstash to forward logs to VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/logstash),
or follow this interactive guide:

{{< cloud-console-iframe "logstash" >}}
