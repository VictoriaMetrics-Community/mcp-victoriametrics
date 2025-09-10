---
draft: false
page: blog blog_post
authors:
  - Roman Khavronenko
  - Mathias Palmersheim
date: 2025-08-22
enableComments: true
title: "Alerting Best Practices"
summary: "Proper alerting is an art. It is all about foreseeing bad scenarios before they happen, so you can prepare for them. In this article, we go through practical recommendations of approaching the alerting to reduce alerting fatigue and avoid false positives."
categories: 
 - Observability
 - Monitoring
tags:
 - monitoring
 - observability
 - alerting
images:
 - /blog/alerting-best-practices/preview.webp
---

A firing alert is like someone ringing your doorbell - it demands your immediate attention, interrupting whatever else you're doing. 
It requires focus and a quick response.

But imagine trying to live in an apartment where the doorbell never stops ringing. You could put in earplugs to block the noise,
but that only masks the problem - it doesn't solve it.

On the other hand, disconnecting the doorbell entirely isn't a solution either. You still want to know when your food or 
a package arrives.

A doorbell that's always silent or always ringing is equally useless. 
The goal is to find the right balance - distinguishing between what truly matters and what doesn't.

## **Every alert should be actionable**

If you're receiving alert notifications and consistently ignoring them, then those alerts shouldn't have been triggered
in the first place. Why go through the trouble of setting a "trap" only to ignore it when it springs?

<p style="max-width: 451px; margin: 1rem auto;">
<img src="/blog/alerting-best-practices/trap.webp" style="width:100%" alt="Trap!">
</p>

As engineers, we don't appreciate the work automated alerting does. It tirelessly checks the conditions we asked it to
check, daily and nightly. Only so we can get upset when it sends us notifications.

Imagine asking a colleague to monitor a server and let you know if something breaks. You gave clear instructions,
and when they follow through - you ignore them. That colleague wouldn't stay motivated for long.

So if you find yourself drowning in alerts or simply tuning them out, it's a signal in itself: something needs to change.
**It's time to take action.**

Please read this outstanding article [Prometheus Alerting 101: Rules, Recording Rules, and Alertmanager](https://victoriametrics.com/blog/alerting-recording-rules-alertmanager/)
by [Phuong Le](https://victoriametrics.com/authors/phuong-le/) to get the basics of the alerting in VictoriaMetrics ecosystem. 
The rest of the article will be dedicated to practical tips on improving the alerting experience.

## **Defining an alerting rule**

The alerting rule consists of multiple fields. Let's start from the most important ones:

![Defining alert](/blog/alerting-best-practices/defining_alert.webp)

Here, we define the alert `RemoteWriteConnectionIsSaturated`, which is supposed to notify us when metrics collector 
is unable to push data fast enough.

The alerting rule name should be descriptive, as it's the first thing an on-call engineer will see.
It should convey a basic understanding of the issue at a glance, before the engineer even reads the rest of the alert message.

### **Rule expression**

A rule’s `expr` should satisfy the following criteria:
1. **It must describe a problematic system state that genuinely requires action from the on-call engineer.** 
  Test the expression against real data to see if it "catches" that problematic state.
2. **Verify that expression gives the expected results in more than one situation.** 
  Try it on longer time intervals, apply it to different environments.
3. **Make sure that expression returns labels that you actually need.** 
  For example, if you don't care about a specific pod experiencing connection issues, then modify the query expression 
  to produce alert per-job by wrapping it with `max(...) by(job) > 0.9`. This approach helps reduce alert noise when 
  multiple pods within the same `job` are affected.

There are a bunch of [common mistakes](https://docs.victoriametrics.com/victoriametrics/vmalert/#common-mistakes) users 
make when configuring alerting rules. But we want to draw attention to the lookbehind window importance.

vmalert executes [instant queries](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#instant-query) for rules.
Instant queries are limited in how far VictoriaMetrics will look back when retrieving data points. For example, a simple
rule like `config_reload_error == 1` will only search for data points within a 5-minute window (controlled by `-datasource.queryStep`).
So if `config_reload_error` scrape interval is `>= 5 minutes`, this query might miss valid data and produce false negatives,
since the expected datapoint might fall just outside the query's lookbehind window.

![Lookbehind window](/blog/alerting-best-practices/lookbehind_window.webp)

In this case, lookbehind window can be extended globally by setting `-datasource.queryStep=15m` (to always look behind for 15min)
or by modifying the query to look for more than 5 minutes:

![Lookbehind window 15m](/blog/alerting-best-practices/lookbehind_window_15m.webp)

> Note: even if scrape_interval is <=5min, you should always account for the possibility of data delivery delay. 
> See more details about data delay [here](https://docs.victoriametrics.com/victoriametrics/vmalert/#data-delay).

The same issue applies to [rollup functions](https://docs.victoriametrics.com/victoriametrics/metricsql/#rollup-functions)
with too short lookbehind window, like `rate(http_request_errors_total[1m])`. If `http_request_errors_total` 
scrape_interval is `1 minute`, then this expression makes no sense as it needs to capture at least **2 data points** to calculate the rate. 

A good rule of thumb is to set the lookbehind window to at least **4x the scrape interval**. This helps ensure accuracy
and accounts for potential delays or missed scrapes.

### **The for param**

The `for` param defines how long the `expr` returns data for a time series before the alert actually fires. Its primary
purpose is to prevent [alerts flapping](https://docs.victoriametrics.com/victoriametrics/vmalert/#flapping-alerts) 
caused by short-lived or transient issues. 

For example, it's normal for a vmagent connection to become temporarily saturated while the remote destination is restarting.
But if the saturation persists for more than 15 minutes, it likely indicates a real problem that won't resolve on its own.

The `for` parameter is one of the most effective tools for reducing noisy alerts. Some metrics - like CPU usage - are 
naturally spiky and prone to short bursts of high values. By increasing the `for` duration, you can filter out these 
harmless spikes and focus on sustained issues. For example, it helps distinguish between a CPU that occasionally handles
heavy workloads and one that remains saturated over an extended period of time.

> Note: the longer the for duration, the more time it takes for the alert to fire. Some alerts are too important to wait 
> for 15 or 30 minutes. Choosing the right for value requires a good understanding of the signal you're monitoring - how
> it behaves and how quickly you need to react when things go wrong.

The `for` param is also related to lookbehind window. For example, `increase(http_request_errors_total[5m])` counts 
the number of errors over the **last 5 minutes**. If there's even a single increment in that time, the expression will
evaluate as true for the full 5-minute window, because the data point remains within the range.

In this case, setting `for: 5m` doesn't add much value, since the alert will likely always remain active for at least that long. 
To make `for` meaningful in such cases, it should be set to a value greater than the lookbehind window. E.g., 
`for: 10m` when using `[5m]`, to ensure you're capturing a sustained condition, not just a single event.

### **The keep_firing_for param**

The opposite of the `for` param is `keep_firing_for`. This setting delays alert resolution by keeping the alert active 
for a specified duration, even if the expr stops returning results.

By default, vmalert waits for the full `for` interval before firing an alert. However, it only needs **one empty evaluation** 
to resolve it. This can lead to alerts resolving too quickly in cases of brief data gaps or missing samples. 
For example, an alerting rule for CPU utilization gets enough above the threshold to become firing:

![Volatile CPU signal](/blog/alerting-best-practices/volatile_signal.webp)

Now imagine that CPU usage drops slightly below the threshold once every 30 minutes—just enough to resolve the alert. 
A few minutes later, it rises above the threshold again and triggers a new alert.

This results in unnecessary alert noise and constant flapping. By setting a `keep_firing_for` interval, you can smooth 
out these fluctuations and avoid repetitive notifications for the same underlying issue.

### **Labels**

Labels are metadata attached to each alert generated by a rule. They serve two primary purposes:
1. **Categorization** – labels help classify the alert (e.g., by severity, team, or environment), allowing it to be 
  properly routed to the right destination or on-call rotation.
1. **Enrichment** – labels can add extra context that isn't available in the original metric, such as static identifiers 
  or tags useful for downstream processing.

![Labels](/blog/alerting-best-practices/labels.webp)

Categorizing alerting notifications is useful for [routing](https://prometheus.io/docs/alerting/latest/configuration/#label-matchers). 
For example, routing by alert's `severity` label will notify the on-call person about `warning`-type alerts, 
while `critical`-type alerts will ping the Engineering Manager that something out of the ordinary is happening.

Another example is routing by department. Having labels `team: platform` and `team: engineering` can help send 
application-related alerts to developers, while alerts related to the platform will be sent to platform engineers.

Enriching alerts with additional information is especially useful when the same set of alerting rules is deployed across
multiple environments. For example, if an alerting rule is running in the EMEA region, you can attach a label like `region="EMEA"`.
This allows the on-call engineer to immediately identify which region is affected, without needing to dig into the metric data.

> Note: one of the common mistakes is setting label values to something dynamically changing, like `$value`. 
> Since `$value` is changing on every rule evaluation, it will change the alert's label set and reset the `for` duration.

### **Annotations**

Annotations are a great way to provide more context about the alert or link to helpful resources.

![Annotations](/blog/alerting-best-practices/annotations.webp)

In the example above, the `summary` and `description` annotations serve as a simplified runbook. 
The reason for using annotations for this information instead of labels is that annotations are not stored in VictoriaMetrics.
They're only stored as part of the alert, making it an ideal place for detailed messages, dashboard links, and other long strings
that would be challenging to store in VictoriaMetrics.

Ideally, alerts should include clear, actionable instructions directly in the notification, so engineers don't need to
look up an external runbook. If you can briefly explain how to respond to an alert, include that guidance in the annotations.

Another good example is `dashboard` annotation. It contains a link to a specific panel on VictoriaMetrics Grafana dashboard.
When clicked, it takes the on-call engineer directly to a visual overview of the issue, showing historical context, 
related metrics, and other signals that can help diagnose and resolve the problem more effectively.

As you can see, we heavily use [templating](https://docs.victoriametrics.com/victoriametrics/vmalert/#templating) 
in annotations to enrich each unique alerting notification with personalized information. 

> It's OK to use templates like `$value` or `$labels` in annotations, as annotations aren't taken into account during `for` checks.

Annotations can be additionally enriched by executing arbitrary MetricsQL queries via `query()` template function:

```yaml
annotations:
 message: |
   The configuration of the instances of the Alertmanager cluster `{{ $labels.namespace }}/{{ $labels.service }}` are out of sync.
   {{ range printf "alertmanager_config_hash{namespace=\"%s\",service=\"%s\"}" $labels.namespace $labels.service | query }}
   Configuration hash for pod {{ .Labels.pod }} is "{{ printf "%.f" .Value }}"
   {{ end }}
```

The `message` annotation above makes an extra query call to fetch `alertmanager_config_hash` metric for the triggered 
alert and prints it in the annotation text.

## **Improving user experience**

Additional information, such as links to Alertmanager, links to silence the alert, and a link to view the alerting rule
that generated the alert, are added automatically to all alerts by vmalert and Alertmanager. However, these URLs need 
to be changed from the defaults in most cases. The `-external.url` and `-external.alert.source` command-line flags in vmalert
will change the external link users see in Alertmanager and in the notifications it sends. However, these will usually 
default to internal service URLs that users do not have access to. To make these links more useful, they should be 
configured to point to something like Grafana that users will have access to. 

Configuring the `-external.url` allows to use the `$externalURL` variable in annotations and makes it easier to share 
rules across environments. For example:
```yaml
- alert: Empty Alert Rules found
  expr: 'max(vmalert_alerting_rules_last_evaluation_series_fetched) by(group, alertname) == 0'
  annotations: 
    summary: empty alerting rules found
    description: "{{ $labels.alertname }} in {{ $labels.group }} does not match any series"
    dashboard: '{{ $externalURL }}/d/LzldHAVnz_vm/victoriametrics-vmalert-vm'
```

The rule above could be applied to multiple environments without any changes, even if the URL to the dashboard there is different.

## **Alerts history**

_"If you want to know the future, look at the past."_

During alerting rules evaluation, vmalert persists [alerts state changes](https://docs.victoriametrics.com/victoriametrics/vmalert/#alerts-state-on-restarts)
in the form of time series with names **ALERTS** and **ALERTS_FOR_STATE**. Using these metrics we can see the history 
of alerts state changes. For this purpose, we (originally attributed to [Alexander Marshalov](https://github.com/amper))
have built a [Grafana dashboard for alerts statistics](https://github.com/VictoriaMetrics/VictoriaMetrics/pull/9427):

![Alerts history](/blog/alerting-best-practices/history.webp)

With the help of the dashboard, we can see which alerts were too noisy, or which alerts have never fired. Both cases are suspicious.

When dealing with alerting fatigue, use this dashboard to find the noisiest alerting rules and inspect their configurations 
for possible optimizations. Remember, **every alert should be actionable**. If there is no action to take on firing alert - it shouldn't exist.

See the [Grafana dashboard here](https://github.com/VictoriaMetrics/VictoriaMetrics/blob/master/dashboards/alert-statistics.json).

## Reducing noise

Usually, the first alerts are defined for relatively slow workloads. For example, the first alert we created for our service looked like this:
```yaml
- alert: RequestErrorsToAPI
  expr: increase(http_request_errors_total[5m]) > 0
```

It catches an unwanted state of the application generating errors for client requests. 

This alert was very helpful when we ran one or two replicas of the application. But once we scaled to hundreds in many 
regions, receiving an alert for each overloaded replica might be overkill. Instead, we can modify the expression to notify
us only about **specific region** experiencing issues:
```yaml
- alert: RequestErrorsToAPI
  expr: sum(increase(vm_http_request_errors_total[5m])) by(region) > 0
```

With the updated expression, we will receive only one alert per region. So even if many replicas within the region 
will start serving errors - we will receive only one firing alert. It still will be actionable, it can contain links to
the dashboard that would show us a more detailed situation. But we won't be overwhelmed with too many notifications.

Maybe even a better approach would be to [define error budget](https://sre.google/workbook/implementing-slos/) and send
alerts only when this budget is burning too fast. This approach assumes that errors are acceptable to some level and 
expects notifying engineers only if the promised service level objective is heading to be breached. 

Sometimes, alerts can start firing because of some incidents that are out of our control: power outage, datacenter failure, etc.
These events could start a cascade of alerting notifications because monitored services depend on connectivity. 
It could be overwhelming to receive thousands of alerts at once, so we recommend [configuring rules inhibition](https://prometheus.io/docs/alerting/latest/configuration/#inhibition-related-settings)
in Alertmanager. It effectively allows muting a set of alerts based on the presence of another set of alerts.

## **Testing alerts**

Above, we recommended testing alerting rule expressions before applying them. But just running them in Grafana Explore 
or vmui could be not indicative, as such query doesn't account for `for` or `keep_firing_for` params. 

As a better approach, we recommend using [vmalert-tool](https://docs.victoriametrics.com/victoriametrics/vmalert-tool/)
for unit-testing rules. Writing tests gives confidence and verification of the expression correctness. It is also a good
practice to include such tests in Continuous Integration (CI) when changing rule definitions.

vmalert also supports a backfilling mechanism called [replay](https://docs.victoriametrics.com/victoriametrics/vmalert/#rules-backfilling).
Via replay, it is possible to run alerting rules on the production data in the past just to see when alerts will or won't trigger.
Results of replay can be verified via `Alerts history` dashboard.

## **Summary**

Proper alerting is an art. It is all about foreseeing bad scenarios before they happen, so you can prepare for them.

![img.webp](/blog/alerting-best-practices/preparing.webp)

VictoriaMetrics ecosystem provides all required tools for defining, testing and monitoring alerting processes. Please refer to the following resources:
1. [Prometheus Alerting 101: Rules, Recording Rules, and Alertmanager](https://victoriametrics.com/blog/alerting-recording-rules-alertmanager/)
2. [VictoriaMetrics Monitoring](https://victoriametrics.com/blog/victoriametrics-monitoring/)
3. [Never-firing alerts: What are they and how to deal with them](https://victoriametrics.com/blog/never-firing-alerts/)
4. https://docs.victoriametrics.com/victoriametrics/vmalert
5. https://docs.victoriametrics.com/victoriametrics/vmalert-tool/ 