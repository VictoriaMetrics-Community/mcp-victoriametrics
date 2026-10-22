---
weight: 7
title: Notifications in VictoriaMetrics Cloud
menu:
  docs:
    parent: "cloud"
    weight: 7
    name: Notifications
tags:
  - metrics
  - logs
  - cloud
  - enterprise
aliases:
  - /victoriametrics-cloud/setup-notifications/index.html
  - /managed-victoriametrics/setup-notifications/index.html
---
The section explains the `Notifications` concept in VictoriaMetrics Cloud.

## What are notifications?

The purpose of notifications is to help teams stay informed about critical events, product updates,
and account activities. They are divided into different interest groups:
* `System alerts`: Critical updates about platform health and infrastructure
* `Billing`: Notifications about invoices, usage limits, and payment issues
* `Product news`: Updates on new features, improvements, and releases
* `Activity`: Events related to user actions and deployments

## Default notification recipients

By default, organization administrators are automatically subscribed to all VictoriaMetrics Cloud
notification categories.

> [!TIP]
> Dedicated documentation about [User Management](https://docs.victoriametrics.com/victoriametrics-cloud/account-management/organizations/)
> is available, including roles and permissions, user management and organizations.

## Configuring notifications

Organization administrators can manage notification settings by using the pencil (✏️) icon next to
each category to edit recipients. Changes done apply across the whole organization.
At a user level, it's also possible to opt-out from available notifications.

Notifications can be received in the following channels:
* Email addresses can be defined for all notification types
* Additionally, Slack channels can be defined to receive System Alerts.

![Editing System Alerts](https://docs.victoriametrics.com/victoriametrics-cloud/notifications-emails.webp)
<figcaption style="text-align: center; font-style: italic;">Editing System Alerts Recipients</figcaption>

> [!TIP]
> Not sure if everything is properly defined? Click on `Send test message` and click `Save` to
> make a fast check.

### Configuring Slack notifications

Slack channel notifications may be configured for System Alerts via webhooks. In order to define
where in Slack notifications should show up, VictoriaMetrics Cloud needs to know:
* The Slack webhook url
* The Channels to receive notifications

To learn about Slack webhook urls please check this information: https://api.slack.com/messaging/webhooks.
In a nutshell, it's needed to create a Slack app, enable incoming webhook, and create an incoming
webhook.

>[!TIP]
> The webhook url to paste in VictoriaMetrics Cloud will have the following format:
> `https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX`


{{% collapse name="Expand to see notification message examples" %}}

![Test Slack Notifications](https://docs.victoriametrics.com/victoriametrics-cloud/notifications_slack_test.webp)
<figcaption style="text-align: center; font-style: italic;">Test message in Slack notifications</figcaption>

<br>

![Test Email Notifications](https://docs.victoriametrics.com/victoriametrics-cloud/notifications_email_test.webp)
<figcaption style="text-align: center; font-style: italic;">Test message via email</figcaption>

{{% /collapse %}}

<br>

## Setup Webhook notifications

Webhook notifications allow you to send alerts to any HTTP endpoint that accepts POST requests.
This is useful for integrating with custom alert handlers or third-party services such as
[Grafana Alerting](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/webhook-notifier/).

### Authentication

Choose one of the following authentication methods:

- **Basic Auth** – Enter a username and password.
- **Bearer Token** – Enter a bearer token to be sent in the `Authorization` header.
- **No Authentication** – Leave authentication fields empty if your endpoint does not require it.

### Configuration

1. Enter the destination URL of your HTTP endpoint.
1. Select the desired authentication method and fill in the required credentials.
1. Press **Save** to store your settings, or **Save and Test** to save and immediately send a test notification to the configured endpoint.

> For a list of services that support Alertmanager webhook receiver integrations,
> see [Alertmanager Webhook Receiver Integrations](https://prometheus.io/docs/operating/integrations/#alertmanager-webhook-receiver).

<br>