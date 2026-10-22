---
build:
  list: never
  publishResources: false
  render: never
sitemap:
  disable: true
---
This section contains a collection of quick start guides for integrating various software and tools
with VictoriaMetrics Cloud. The guides are organized by integration type to help you quickly find
what you need. Some integrations may apply only to specific `Deployment` types, depending on their
nature.

In the [VictoriaMetrics Cloud Console](https://console.victoriametrics.cloud/integrations), you can
explore interactive integration guides tailored to your actual deployments. These guides provide
personalized assistance, including pre-filled configuration snippets, relevant URLs, and secure
access tokens.

> [!NOTE] Tip
> Generally, all integrations in VictoriaMetrics Cloud just require a URL and an
> [Access Token](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/) to work.

In this section, the same principle is applied so, in each integration, interactive guides and
examples are provided for example deployments. Don't forget that you can always find interactive
integration guides in [VictoriaMetrics Cloud Console](https://cloud.victoriametrics.com/integrations/),
which generate steps and configurations specifically for your real cloud deployments.

If there's an integration you'd like to see here but it's currently missing, feel free to [contact us](https://docs.victoriametrics.com/victoriametrics-cloud/support/#how-to-reach-support).

## Ingestion

- [CloudWatch - Agentless AWS monitoring](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/cloudwatch/)
- [CURL](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/curl/)
- [Fluent Bit](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/fluentbit/)
- [Grafana Alloy](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/alloy/)
- [Kubernetes](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/kubernetes/)
- [Logstash](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/logstash/)
- [OpenTelemetry](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/opentelemetry/)
- [Prometheus (remote write)](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/prometheus/)
- [Promtail](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/promtail/)
- [Telegraf](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/telegraf/)
- [Vector](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/vector/)
- [VLAgent](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/vlagent/)
- [VMAgent](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/vmagent/)
- [VTAgent](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/vtagent/)

## Data Visualization

- [AI Assistants](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/ai-assistants/)
- [CURL](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/curl/)
- [Grafana](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/grafana/)
- [Perses](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/perses/)

## Notifications

- [Cloud Alertmanager](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/cloud-alertmanager/)
- [Custom Alertmanager](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/custom-alertmanager/)

## Alerting Presets

- [Alerting Rule Presets](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/alerting-rule-presets/)

Ready-to-use sets of alerting and recording rules for common exporters, tools and runtimes - apply
them to a deployment from **Integrations** in the cloud console instead of writing your own.

- [Node Exporter](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=node-exporter)
- [Kube State Metrics](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=kube-state-metrics)
- [cAdvisor](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=cadvisor)
- [ArgoCD](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=argocd)
- [Clickhouse](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=clickhouse)
- [Consul](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=consul)
- [CoreDNS](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=coredns)
- [etcd](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=etcd)
- [FluxCD](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=fluxcd)
- [Istio](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=istio)
- [OpenTelemetry Collector](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=opentelemetry-collector)
- [PostgreSQL](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=postgresql)
- [Prometheus Self-Monitoring](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=prometheus-self-rules)
- [Thanos](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=thanos-rules)
- [Cilium](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=cilium-rules)
- [HAProxy](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=haproxy-rules)
- [Traefik](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=traefik-rules)
- [cert-manager](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=cert-manager-rules)
- [Grafana Mimir](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=grafana-mimir-rules)
- [Envoy](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=envoy-rules)
- [Grafana Loki](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=loki-rules)
- [Grafana Tempo](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=grafana-tempo-rules)
- [Cortex](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=cortex-rules)
- [HashiCorp Vault](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=hashicorp-vault-rules)
- [MinIO](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=minio-rules)
- [Jaeger](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=jaeger-rules)
- [Linkerd](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=linkerd-rules)
- [Grafana Alloy](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=grafana-alloy-rules)
- [Promtail](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=promtail-rules)
- [HashiCorp Nomad](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=nomad-rules)
- [Caddy](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=caddy-rules)
- [Spinnaker](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=spinnaker-rules)
- [JVM](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=jvm-rules)
- [Golang applications](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=golang-rules)
- [Python applications](https://console.victoriametrics.cloud/integrations/alerting-rule-presets?integration=python-rules)

## Infrastructure Management

- [Terraform](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/terraform/)
