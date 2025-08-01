---
weight: 8
title: VMPodScrape
menu:
  docs:
    identifier: operator-cr-vmpodscrape
    parent: operator-cr
    weight: 8
aliases:
  - /operator/resources/vmpodscrape/
  - /operator/resources/vmpodscrape/index.html
tags:
  - kubernetes
  - metrics
---
The `VMPodScrape` CRD allows to declaratively define how a dynamic set of pods should be monitored.
Use label selections to match pods for scraping. This allows an organization to introduce conventions
for how metrics should be exposed. Following these conventions new services will be discovered automatically without
need to reconfigure.

`VMPodScrape` object generates part of [VMAgent](https://docs.victoriametrics.com/victoriametrics/vmagent/) configuration with
[kubernetes service discovery](https://docs.victoriametrics.com/victoriametrics/sd_configs/#kubernetes_sd_configs) role `pod` having specific labels and ports.
It has various options for scraping configuration of target (with basic auth,tls access, by specific port name etc.).

A `Pod` is a collection of one or more containers which can expose Prometheus metrics on a number of ports.

The `VMPodScrape` object discovers pods and generates the relevant scraping configuration.

The `PodMetricsEndpoints` section of the `VMPodScrapeSpec` is used to configure which ports of a pod are going to be
scraped for metrics and with which parameters.

Both `VMPodScrape` and discovered targets may belong to any namespace. It is important for cross-namespace monitoring
use cases, e.g. for meta-monitoring. Using the `podScrapeNamespaceSelector` of the `VMAgentSpec`
one can restrict the namespaces from which `VMPodScrape`s are selected from by the respective [VMAgent](https://docs.victoriametrics.com/operator/resources/vmagent/) server.
Using the `namespaceSelector` of the `VMPodScrapeSpec` one can restrict the namespaces from which `Pods` are discovered from.
To discover targets in all namespaces the `namespaceSelector` has to have value `any: true` specified:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMPodScrape
metadata:
  name: example
spec:
  namespaceSelector:
    any: true
```

More information about selectors you can find in [this doc](https://docs.victoriametrics.com/operator/resources/vmagent/#scraping).

## Specification

You can see the full actual specification of the `VMPodScrape` resource in
the **[API docs -> VMPodScrape](https://docs.victoriametrics.com/operator/api/#vmpodscrape)**.

Also, you can check out the [examples](#examples) section.

## Migration from Prometheus

The `VMPodScrape` CRD from VictoriaMetrics Operator is a drop-in replacement
for the Prometheus `PodMonitor` from prometheus-operator.

More details about migration from prometheus-operator you can read in [this doc](https://docs.victoriametrics.com/operator/migration/).

## Examples

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMPodScrape
metadata:
  name: example
spec:
  podMetricsEndpoints:
    - port: web
      scheme: http
  selector:
    matchLabels:
      owner: dev
```
