---
weight: 10
title: VMRule
menu:
  docs:
    identifier: operator-cr-vmrule
    parent: operator-cr
    weight: 10
aliases:
  - /operator/resources/vmrule/
  - /operator/resources/vmrule/index.html
tags:
  - kubernetes
  - metrics
  - logs
---
`VMRule` represents [alerting](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/)
or [recording](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/) rules 
for [VMAlert](https://docs.victoriametrics.com/operator/resources/vmalert/) instances.

The `VMRule` CRD declaratively defines a desired Prometheus rule to be consumed by one or more VMAlert instances.

`VMRule` object generates [VMAlert](https://docs.victoriametrics.com/operator/resources/vmalert/) 
with ruleset defined at `VMRule` spec.

Alerts and recording rules can be saved and applied as YAML files, and dynamically loaded without requiring any restart.

See more details about rule configuration in [VMAlert docs](https://docs.victoriametrics.com/victoriametrics/vmalert/#quickstart).

## Specification

You can see the full actual specification of the `VMRule` resource in
the **[API docs -> VMRule](https://docs.victoriametrics.com/operator/api/#vmrule)**.

Also, you can check out the [examples](#examples) section.

## Enterprise features

Custom resource `VMRule` supports feature [Multitenancy](https://docs.victoriametrics.com/victoriametrics/vmalert/#multitenancy)
from [VictoriaMetrics Enterprise](https://docs.victoriametrics.com/victoriametrics/enterprise/#victoriametrics-enterprise).

### Multitenancy

For using [Multitenancy](https://docs.victoriametrics.com/victoriametrics/vmalert/#multitenancy) in `VMRule`
you need to **[enable VMAlert Enterprise](https://docs.victoriametrics.com/operator/resources/vmalert/#enterprise-features)**.

After that you can add `tenant` field for groups in `VMRule`:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMRule
metadata:
  name: ent-example
spec:
  groups:
    - name: vmalert-1
      rules:
        # using enterprise features: Multitenancy
        # more details about multitenancy you can read on https://docs.victoriametrics.com/operator/resources/vmalert/#multitenancy
        - tenant: 1
          alert: vmalert config reload error
          expr: delta(vmalert_config_last_reload_errors_total[5m]) > 0
          for: 10s
          labels:
            severity: major
            job:  "{{ $labels.job }}"
          annotations:
            value: "{{ $value }}"
            description: 'error reloading vmalert config, reload count for 5 min {{ $value }}'
```

## Examples

### Alerting rule

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMRule
metadata:
  name: alerting-example
spec:
  groups:
    - name: vmalert
      rules:
        - alert: vmalert config reload error
          expr: delta(vmalert_config_last_reload_errors_total[5m]) > 0
          for: 10s
          labels:
            severity: major
            job:  "{{ $labels.job }}"
          annotations:
            value: "{{ $value }}"
            description: 'error reloading vmalert config, reload count for 5 min {{ $value }}'
```

### Recording rule

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMRule
metadata:
  name: recording-example
spec:
  groups:
    - name: vmrule_recording_groupname
      interval: 1m
      rules:
        - record: vm_http_request_errors_total:sum_by_cluster_namespace_job:rate:5m
          expr: |-
            sum by (cluster, namespace, job) (
              rate(vm_http_request_errors_total[5m])
            )
```
