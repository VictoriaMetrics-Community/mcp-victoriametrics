---
weight: 7
title: VMNodeScrape
menu:
  docs:
    identifier: operator-cr-vmnodescrape
    parent: operator-cr
    weight: 7
aliases:
  - /operator/resources/vmnodescrape/
  - /operator/resources/vmnodescrape/index.html
tags:
  - kubernetes
  - metrics
---
The `VMNodeScrape` CRD provides discovery mechanism for scraping metrics kubernetes nodes,
it is useful for node exporters monitoring.

`VMNodeScrape` object generates part of [VMAgent](https://docs.victoriametrics.com/victoriametrics/vmagent/) configuration.
It has various options for scraping configuration of target (with basic auth,tls access, by specific port name etc.).

By specifying configuration at CRD, operator generates config 
for [VMAgent](https://docs.victoriametrics.com/victoriametrics/vmagent/) and syncs it. It's useful for cadvisor scraping,
node-exporter or other node-based exporters. `VMAgent` `nodeScrapeSelector` must match `VMNodeScrape` labels.

More information about selectors you can find in [this doc](https://docs.victoriametrics.com/operator/resources/vmagent/#scraping).

## Specification

You can see the full actual specification of the `VMNodeScrape` resource in
the **[API docs -> VMNodeScrape](https://docs.victoriametrics.com/operator/api/#vmnodescrape)**.

Also, you can check out the [examples](#examples) section.

## Examples

### Cadvisor scraping

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMNodeScrape
metadata:
  name: cadvisor-metrics
spec:
  scheme: "https"
  tlsConfig:
    insecureSkipVerify: true
    caFile: "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
  bearerTokenFile: "/var/run/secrets/kubernetes.io/serviceaccount/token"
  relabelConfigs:
    - action: labelmap
      regex: __meta_kubernetes_node_label_(.+)
    - targetLabel: __address__
      replacement: kubernetes.default.svc:443
    - sourceLabels: [__meta_kubernetes_node_name]
      regex: (.+)
      targetLabel: __metrics_path__
      replacement: /api/v1/nodes/$1/proxy/metrics/cadvisor
```
