---
weight: 1
title: "Tiers limits"
menu:
  docs:
    parent: "deployments"
    weight: 8
    name: "Tiers limits"
tags:
  - metrics
  - cloud
  - enterprise
---

## Exceeding Limits

If your usage exceeds the limits of your current tier, you may experience throttling or errors.
You will be notified via the alert system when you are approaching or have reached a limit.
Consider upgrading to a higher tier or [contacting support](https://console.victoriametrics.cloud/contact_support) for custom options.

Most relevant tier limits are available in the VictoriaMetrics Cloud deployment overview page.

![Deployment overview](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/tiers-and-types-deployment-overview.webp)

You can also check your current tier limits and usage in the **Monitoring** panel of your deployment within the VictoriaMetrics Cloud dashboard.
This helps you proactively monitor your resource consumption and avoid unexpected issues.

![Monitoring panel](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/tiers-and-types-monitoring-example.webp)

When a limit is approached or exceeded, a system alert will be generated to notify you of the situation.

The system alert will appear in the **Alerts** section of the VictoriaMetrics Cloud dashboard.
![Alerts section](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/tiers-and-types-alert-section.webp)

This alert will also be sent to your email address or via the Slack integration if
configured in the [**Notifications**](https://console.victoriametrics.cloud/notifications) section of the VictoriaMetrics Cloud dashboard.

All system alerts are visible in the overview page of your deployment.
![Deployment overview with alert](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/tiers-and-types-deployment-overview-with-alert.webp)
