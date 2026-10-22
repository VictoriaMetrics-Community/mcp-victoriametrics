---
title : "Alerting Rule Presets"
menu:
  docs:
    parent: "integrations"
---

Ready-to-use sets of alerting and recording rules for common exporters, tools and runtimes, so you
don't have to write and maintain your own.

## Available presets

Pick an integration below to see its presets and the rules each one contains:

{{< cloud-console-iframe "alerting-rule-presets" >}}

## Applying a preset

Presets are applied to a deployment in the [cloud console](https://console.victoriametrics.cloud/integrations/alerting-rule-presets):

1. Open **Integrations** and choose an alerting rules integration, or open the deployment's
   **Alerting** settings and choose **Add from presets**.
1. Select the target deployment and the preset.
1. Disable the rules you don't need and click **Apply**.

Applied presets are listed under **Installed presets**, where they can be edited or deleted.

To get notified when these alerts fire, connect a notifier - see
[Setup Alertmanager & VMAlert for VictoriaMetrics Cloud](https://docs.victoriametrics.com/victoriametrics-cloud/alertmanager-setup-for-deployment/).
