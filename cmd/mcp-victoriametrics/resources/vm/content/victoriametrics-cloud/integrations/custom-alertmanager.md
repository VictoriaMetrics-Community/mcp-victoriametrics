---
title : "Custom Alertmanager"
menu:
  docs:
    parent: "integrations"
---

VictoriaMetrics Cloud allows you to define and manage alerting rules using
[vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/), and send notifications to an
**external Alertmanager instance** that you host and control in your own environment.

This integration provides full flexibility for organizations that already operate their own
Alertmanager setup and want to connect it to VictoriaMetrics Cloud’s alerting engine.

## Integrating with Custom Alertmanager

To configure alerting rules and route notifications to your self-hosted Alertmanager, simply follow
this guide:

{{< cloud-console-iframe "custom-alertmanager" >}}
