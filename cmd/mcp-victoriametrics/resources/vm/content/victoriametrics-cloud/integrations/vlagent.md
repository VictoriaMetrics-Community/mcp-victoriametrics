---
title : "VLAgent"
menu:
  docs:
    parent: "integrations"
---

[VLAgent](https://docs.victoriametrics.com/victorialogs/vlagent/) is VictoriaMetrics' native log
collection agent for Kubernetes. Deployed as a **DaemonSet** via Helm chart or the VictoriaMetrics
operator, it runs on every node, automatically discovers and collects logs from all pod containers on
that node, enriches them with Kubernetes metadata (pod name, namespace, labels), and forwards them to
VictoriaLogs using the native protocol.

This integration is useful for collecting Kubernetes pod logs cluster-wide without per-application
configuration.

## Integrating with VLAgent

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access**, since VLAgent pushes data), follow the steps
in the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To set up VLAgent to forward logs to VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/vlagent),
or follow this interactive guide:

{{< cloud-console-iframe "vlagent" >}}
