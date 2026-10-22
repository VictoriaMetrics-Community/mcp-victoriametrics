---
draft: false
page: blog blog_post
authors:
 - Pablo Fernandez
date: 2026-07-14
enableComments: true
title: "Announcing vmestimator: Real-time Cardinality Estimations for VictoriaMetrics and Prometheus"
summary: "vmestimator is a new open-source tool that estimates metric cardinality in real time. It can alert to high-cardinality spikes before they slow queries and increase resource usage. Works with any Prometheus-compatible setups, Grafana, and vmalert"
categories:
 - Performance
 - Time Series Database
 - Monitoring
tags:
 - performance
 - prometheus
 - Grafana
 - VictoriaMetrics
 - alerting
 - cardinality
 - monitoring tools
 - scalability
 - vmagent
 - vmalert
images:
 - /blog/vmestimator-announcement/preview.webp
---

Cardinality problems usually begin with a small change that looks harmless: you add a label, and suddenly one metric turns into thousands of unique series. Cardinality explosions are often caught only after performance degrades. And at that point, your observability stack may be degraded and painful to troubleshoot.

[vmestimator](https://docs.victoriametrics.com/victoriametrics/vmestimator/) is a new project specifically designed to follow cardinality trends in real time and send you alerts before they turn into a real problem.

<iframe width="560" height="315" src="https://www.youtube.com/embed/o-XGkZqrg-Q?si=Wvg5Vharyt0JZnND" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

## When Cardinality Becomes a Problem

Metric cardinality is the number of unique time series your monitoring system stores across all combinations of metric names and labels. For example, `http_requests_total{path="/login"}` and `http_requests_total{path="/checkout"}` are two distinct series because the `path` label differs.

You can see how highly variable labels such as `user_id` or `pod_name` can drastically increase the cardinality of your time-series database. When labels are unbounded or change frequently, your system has to ingest and index more series, increasing pressure on the monitoring stack.

## Announcing vmestimator

[vmestimator](https://docs.victoriametrics.com/victoriametrics/vmestimator/) is a **new open-source tool** (Apache-2.0) from VictoriaMetrics that continuously tracks cardinality. When coupled with vmalert, it can send you alerts identifying offending jobs within minutes of a spike in cardinality.

Getting alerts is the [main use case](https://docs.victoriametrics.com/victoriametrics/vmestimator/#use-cases), but it can also be useful to:

- Track global cardinality across all active series.
- Measure cardinality per metric name to see which metrics create the most series.
- Drill down into cardinality per job to identify the services behind cardinality growth.
- Measure per-tenant cardinality for usage analysis in multi-tenant setups.
- Estimate churn rate to find workloads that continuously create and replace series.
- Follow long-term cardinality trends over time.
- Support capacity planning by showing how cardinality evolves across jobs, metrics, or tenants.

## How vmestimator Works

vmestimator runs at the collection point, deriving cardinality estimates from the incoming stream and exposing them as Prometheus‑compatible metrics for scraping and alerting. It's usually deployed alongside [vmagent](https://docs.victoriametrics.com/victoriametrics/vmagent/) and observes ingested data without interfering with the ingestion path.

![Basic vmestimator topology diagram](arch_basic_1.webp)

vmestimator output fits naturally into any Prometheus-compatible observability system. You can store the results in VictoriaMetrics, [query them with the API](https://docs.victoriametrics.com/victoriametrics/url-examples/#apiv1query), trigger alerts with vmalert, and visualize cardinality estimates in Grafana to spot trends.

![Topology diagram using vmalert and Grafana](arch_vmalert_grafana_1.webp)

## Get Started with vmestimator

To get started with vmestimator, grab the [example config](https://github.com/VictoriaMetrics/vmestimator/blob/main/streams.yaml) and save it as `streams.yaml`. Then [download the binary for your architecture](https://github.com/VictoriaMetrics/vmestimator/releases) or run it as a Docker image:

```sh
docker run --rm \
  -p 8490:8490 \
  -v /path/to/streams.yaml:/streams.yaml \
  docker.io/victoriametrics/vmestimator:latest \
  -config=/streams.yaml
```

Behavior is controlled in `streams.yaml`, where you define parameters such as the estimation interval and arbitrary label dimensions. For instance, to estimate cardinality globally with a 15-minute interval:

```yaml
# streams.yaml
- interval: '15m'
```

If you want to group cardinality per label name, add `group_by`. For example:

```yaml
# streams.yaml
- interval: '15m'
  group_by: ['__name__']
```

To view all options, see the [configuration reference](https://docs.victoriametrics.com/victoriametrics/vmestimator/#configuration).

Once configured, vmestimator exposes its estimates and operational metrics at `/metrics` so that vmagent or Prometheus can scrape them like any other target.

The key metric is called `cardinality_estimate`, which shows the current estimate over a given interval. For example, the following query shows the global estimate over a 15-minute interval (you can try it directly in the [VictoriaMetrics Playground](https://play.victoriametrics.com/select/0/vmui/?#/?g0.range_input=2d&g0.end_input=2026-07-05T13%3A52%3A53&g0.relative_time=last_2_days&g0.tab=0&g0.expr=cardinality_estimate{interval%3D%2215m0s%22%2Cgroup_by_keys%3D%22__global__%22})).

```promql
cardinality_estimate{interval="15m0s",group_by_keys="__global__"}
```

![Screenshot of VMUI showing cardinality trends](cardinality_estimate_global_15m.webp)

vmestimator ships with [two Grafana dashboards](https://github.com/VictoriaMetrics/vmestimator/tree/main/dashboards):

1. [vmestimator metrics](https://play-grafana.victoriametrics.com/d/mkv22l4/victoriametrics-vmestimator?orgId=1&from=now-1h&to=now&timezone=utc&var-ds=P4169E866C3094E38&var-job=$__all&var-instance=$__all): shows metrics for the vmestimator application itself. Includes CPU, memory, ingestion rates, concurrent inserts, and group key saturation.
2. [Cardinality Estimations](https://play-grafana.victoriametrics.com/d/mktd5g9/victoriametrics-cardinality-estimations?orgId=1&from=now-3h&to=now&timezone=utc&var-ds=P4169E866C3094E38&var-job=$__all&var-instance=$__all&var-interval=$__all&var-group_by_keys=$__all&var-group_by_values=$__all&var-churn_short=15m0s&var-churn_long=30m0s): shows global cardinality, churn rate, and top 10 metrics and jobs.

![Screenshot of the Grafana cardinality estimations dashboard](grafana_cardinality_estimations.webp)

## Alerts with vmalert

By plugging in [vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/), you can get alerted when cardinality exceeds a predefined threshold.

Get started with the [examples provided in the vmestimator repository](https://github.com/VictoriaMetrics/vmestimator/blob/main/deployment/docker/rules/alerts-cardinality.yml). For example, the first alert fires when the per-job cardinality exceeds 20,000 over a 30-minute interval:


```yaml
groups:
  - name: 'cardinality'
    rules:
      - alert: 'JobTooHighCardinality'
        expr: 'max (cardinality_estimate{group_by_keys="job",interval="30m0s"}) without(instance) > 20000'
        for: '15m'
        labels:
          severity: 'warning'
        annotations:
          dashboard: "{{ $externalURL }}/d/mktd5g9?var-group_by_keys=job&var-interval={{ $labels.interval }}"
          summary: "High cardinality for job {{ $labels.group_by_values }}"
          description: "Job {{ $labels.group_by_values }} has {{ $value }} estimated series over the last {{ $labels.interval }}, exceeding the 20000 series threshold."

...
```

The second alert compares two streams at different intervals and triggers when the churn rate is too high. A value of 1 means a 100% churn in active series:

```yaml
...
      - alert: 'JobTooHighChurnRate'
        expr: |
          (
              max(cardinality_estimate{group_by_keys="job",interval="30m0s"}) without (instance)
              -
              max(cardinality_estimate{group_by_keys="job",interval="15m0s"}) without (instance)
          )
          /
          max(cardinality_estimate{group_by_keys="job",interval="30m0s"}) without (instance) * 100 > 10
        for: '15m'
        labels:
          severity: 'warning'
        annotations:
          dashboard: "{{ $externalURL }}/d/mktd5g9?var-group_by_keys=job&var-churn_short=15m0s&var-churn_long=30m0s"
          summary: "High churn rate for job {{ $labels.group_by_values }}"
          description: 'Job {{ $labels.group_by_values }} has {{ printf "%.0f%%" $value }} of series churned between the 15m and 30m cardinality windows, exceeding the 10% threshold.'

...
```

You can view cardinality alerts online in our [VictoriaMetrics Grafana Playground](https://play-grafana.victoriametrics.com/alerting/list). Look inside the "cardinality" and "cardinality.recording" sections.

![](/blog/vmestimator-announcement/grafana-managed-alerts-playground.webp)
<figcaption style="text-align: center; font-style: italic;">Grafana <a href="https://docs.victoriametrics.com/guides/vmalert-datasource-managed-alerts-grafana/">datasource-managed alerts</a> in the VictoriaMetrics Grafana Playground.</figcaption>


## Conclusion

If you already use a Prometheus-compatible pipeline, vmestimator slots into that workflow with minimal extra moving parts. It gives you an early warning on runway cardinality and helps with long-term capacity planning.

## Frequently Asked Questions

### Is vmestimator compatible with Prometheus?

Yes. vmestimator exports Prometheus-compatible metrics and is compatible with any client that can send data in Prometheus remote write v1 format. So, you can scrape it like any other target and use the results in PromQL-based dashboards and alerts.

### Do I need vmalert and VictoriaMetrics to use vmestimator?

No. vmestimator can run on its own and export metrics over HTTP, but vmalert and VictoriaMetrics make it much more useful by letting you store, query, visualize, and alert on cardinality estimates.

### What problem does vmestimator solve?

It helps you detect cardinality growth early by estimating it in real time at the source, across arbitrary label dimensions. This makes it useful for tracking trends and triggering alerts on sudden changes.

### What can I monitor with vmestimator?

You can track cardinality across arbitrary label dimensions, including globally, per metric name, per job, per tenant, and over time through churn estimates.

### How do I run vmestimator?

You can [download the binary](https://github.com/VictoriaMetrics/vmestimator/releases) from GitHub, build it from source, or run it as a Docker container. vmestimator runs from a `streams.yaml` file, where you choose the estimation interval, grouping labels, and related options. You can run it with the example config, scrape the exported `cardinality_estimate` metrics, and pair it with Grafana and vmalert to visualize trends and alert on spikes.

### What's the difference between vmestimator and the Cardinality Explorer in VMUI?

A [cardinality explorer](https://docs.victoriametrics.com/victoriametrics/vmestimator/#cardinality-explorer) view is included with VMUI; vmestimator is meant to supplement rather than replace it.

vmestimator estimates cardinality only for predefined label groups, but it does so efficiently and in real time, making it suitable for monitoring and alerting. Cardinality Explorer, on the other hand, can analyze cardinality across arbitrary label combinations, but those queries are more expensive and don't provide historical trends. That makes it better suited for ad-hoc investigations and in-depth analysis.

They measure cardinality at different points in the pipeline. vmestimator estimates it at the source, before ingestion, while Cardinality Explorer reflects the cardinality of data already stored in a database. Source-side cardinality can be aggregated, whereas cardinality across cluster storage nodes cannot.



