---
draft: false
page: blog blog_post
authors:
  - Jose Gomez-Selles
  - Alexander Marshalov
date: 2025-06-10
title: "Integrations made easy with VictoriaMetrics Cloud"
enableComments: true
summary: "Discover the latest improvements to integrations in VictoriaMetrics Cloud, including interactive guides, streamlined Kubernetes monitoring, and our commitment to full-stack observability without vendor lock-in."
categories:
  - Product News
tags:
  - victoriametrics
  - integrations
  - cloud
  - observability
  - kubernetes
  - opentelemetry
images:
  - /blog/integrations-made-easy-with-victoriametrics-cloud/preview.webp
---

VictoriaMetrics Cloud continues to evolve as the most efficient, scalable and open platform
in the observability landscape. In our last [Q1 update blogpost](https://victoriametrics.com/blog/q1-2025-whats-new-victoriametrics-cloud/),
we shared new features such as seamless OpenTelemetry integrations, new Organizations support, and
improvements in the Explore UI and APIs.

This time we wanted to take a minute to showcase how we’re taking the **interoperability** journey
very seriously.

{{<image href="/blog/integrations-made-easy-with-victoriametrics-cloud/integrations.webp" alt="Integrations in VictoriaMetrics Cloud"  width="{{ 35 }}">}}

>[!TIP] Haven’t tried VictoriaMetrics Cloud yet?
> [Sign up for free](https://console.victoriametrics.cloud/signup) — no credit card required — and get $200 in credits for one month.


## Integrations in VictoriaMetrics Cloud

VictoriaMetrics has always focused on **openness and interoperability**. Whether you’re collecting
data from Prometheus, OpenTelemetry, Graphite, or pushing metrics through any observability
stack — we don’t force any tools on you. This unopinionated approach truly ensures you **avoid vendor
lock-in** and fit VictoriaMetrics into your existing workflows.

The newly published [integrations documentation](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/) demonstrates just how easy it is to integrate your systems. Once you select a deployment, all the instructions — from sending metrics to visualizing them — are generated for you with the correct URL and token.

>[!IMPORTANT] Integrating with VictoriaMetrics Cloud
> To get started, all you need is a **URL and an Access Token**.

## Integrate With Everything

All integrations come with **interactive, step-by-step guides** available in the [VictoriaMetrics Cloud Console](https://console.victoriametrics.cloud/integrations/),
tailored for your real deployments. This includes **copy-paste-ready snippets** with all required settings in place.

You can check it out directly in our revamped [integrations documentation](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/)
and experience how, with a couple of clicks in the [VictoriaMetrics Cloud Console](https://console.victoriametrics.cloud/integrations/)
you can get started with real-world integration snippets customized for your deployments.

VictoriaMetrics Cloud supports integration across the **entire observability lifecycle**:

### Ingestion
- [**CloudWatch - Agentless AWS monitoring**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/cloudwatch/) integration allows to **forward metrics from AWS services** (like EC2, RDS, Lambda, etc.) to VictoriaMetrics Cloud without deploying extra collectors or agents.
- [**CURL**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/curl/) can be used to interact with VictoriaMetrics Cloud for **pushing metrics using HTTP API endpoints**.
- [**Kubernetes**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/kubernetes/): collect metrics from **cluster, nodes, and workloads**, and forward them to VictoriaMetrics Cloud using [vmagent](https://docs.victoriametrics.com/victoriametrics/vmagent/).
- [**OpenTelemetry**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/opentelemetry/) Collector (using either the **Helm chart or the Operator**) to collect, process, and forward observability data from a wide variety of sources into VictoriaMetrics Cloud.
- [**Prometheus (remote write)**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/prometheus/), allows forwarding metrics collected by Prometheus to VictoriaMetrics Cloud for **long-term storage and advanced querying**.
- [**Telegraf**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/telegraf/) integration is ideal for environments where Telegraf is already used to gather **system, application, or custom metrics**.
- If you are using [Vector](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/vector/), this integration is useful to **route metrics** to VictoriaMetrics Cloud for storage and analysis.
- [**vmagent**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/vmagent/) is a lightweight agent designed to collect metrics from various sources, apply relabeling and filtering rules, and forward the data to storage systems. It supports both the Prometheus remote_write protocol and the [VictoriaMetrics remote_write protocol](https://docs.victoriametrics.com/victoriametrics/vmagent/#victoriametrics-remote-write-protocol) for sending data. This makes [vmagent](https://docs.victoriametrics.com/victoriametrics/vmagent/) **ideal for centralized metric collection and forwarding in a resource-efficient way**.

### Visualization
- VictoriaMetrics Cloud can easily be added as a datasource to [**Grafana**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/grafana/), via the built-in **Prometheus or VictoriaMetrics datasources**. This integration allows you to build powerful, customizable dashboards and monitor your systems in real time using VictoriaMetrics as the backend.
- [**Perses**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/perses/): VictoriaMetrics Cloud can be used as a data source in [Perses](https://perses.dev) via the Prometheus-compatible query API, allowing you to **create dashboards and monitor time series** data with a modern and lightweight interface.
- [**CURL**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/curl/), can also be used for querying stored data using HTTP API endpoints. This makes it a simple and flexible option for testing or **basic integrations**.

### Alerts & Notifications
- [**Cloud AlertManager**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/cloud-alertmanager/): VictoriaMetrics Cloud comes with a fully managed AlertManager based on [vmalert](https://docs.victoriametrics.com/vmalert/), which can be used to send notifications. This integration provides a seamless way to **trigger alerts based on Prometheus-compatible queries** and route them to your preferred notification channels, such as email, PagerDuty, Slack, MS Teams or webhooks.
- If you already have a [**Custom AlertManager**](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/custom-alertmanager/) VictoriaMetrics Cloud allows you to define and manage alerting rules using to your instance. This integration provides **full flexibility for organizations that already operate their own AlertManager setup** and want to connect it to VictoriaMetrics Cloud’s alerting engine.

## Example: Integrate Kubernetes with VictoriaMetrics Kubernetes Stack

As already mentioned, if you are curious about how easy it is to integrate with VictoriaMetrics Cloud,
we encourage you to visit our [docs](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/)
to experience by yourself. However, here's a brief example of what it takes for one of the most used integrations:
**Monitoring Kubernetes via the [VictoriaMetrics Kubernetes Stack](https://docs.victoriametrics.com/helm/victoriametrics-k8s-stack/)**:
a Helm chart that brings together all the key components to collect and push metrics to VictoriaMetrics
Cloud if you are starting from zero.

This stack includes:

- **vmagent** to collect and forward metrics.
- Prebuilt **dashboards and alerts** to get started quickly in observability.
- Easy setup using just your `<DEPLOYMENT_ENDPOINT_URL>` and `<YOUR_ACCESS_TOKEN>`.

{{<image href="/blog/integrations-made-easy-with-victoriametrics-cloud/Kubernetes.svg" alt="Integrating Kubernetes with VictoriaMetrics Cloud"  width="{{ 35 }}">}}


### Three steps to observe:

1. **Create a Secret with your VictoriaMetrics Cloud Access Token**:
```bash
kubectl create secret generic vmauth-creds \
  --from-literal=VMAUTH_TOKEN='<YOUR_ACCESS_TOKEN>'
```

2. **Update the `values.yaml`** with your endpoint and credentials:
```yaml
vmagent:
  remoteWrite:
    - url: https://<DEPLOYMENT_ENDPOINT_URL>/api/v1/write
      headers:
        Authorization: "Bearer $VMAUTH_TOKEN"
```

3. **Install the Helm chart**:
```bash
helm repo add victoria-metrics https://victoriametrics.github.io/helm-charts/
helm upgrade --install vm-stack victoria-metrics/victoria-metrics-k8s-stack \
  -f values.yaml
```

>[!TIP] How do I pick endpoint and Access Token?
> Both the endpoint and Access Token are already filled for you in VictoriaMetrics Cloud at the [interactive Kubernetes integration guide](https://console.victoriametrics.cloud/integrations/kubernetes).

After that, your cluster metrics will be flowing into VictoriaMetrics Cloud, with dashboards and alerts ready to go!

## Many Ways to Adapt to Your Needs

But this is not the only way to ingest Kubernetes metrics! As highlighted in our [Q1 blog post](https://victoriametrics.com/blog/q1-2025-whats-new-victoriametrics-cloud/),
VictoriaMetrics is fully committed to **OpenTelemetry** as a first-class citizen in our observability stack.

You can monitor your Kubernetes cluster using OpenTelemetry as well, either via the Helm chart or the OpenTelemetry Operator. Explore our [OpenTelemetry integration docs](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/opentelemetry/) or try the [guided setup in the Cloud Console](https://console.victoriametrics.cloud/integrations/opentelemetry).

{{<image href="/blog/integrations-made-easy-with-victoriametrics-cloud/OpenTelemetryHelm.svg" alt="Integrating Kubernetes using OpenTelemetry with VictoriaMetrics Cloud"  width="{{ 35 }}">}}


Last reminder (promised!): as every integration with VictoriaMetrics Cloud, you’ll only need a URL and an [Access Token](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/)
to connect your data source or visualization tool — no guesswork needed!

## Help Us Improve

We’re always looking to make VictoriaMetrics Cloud better. Please take a moment to [fill out our quick survey](https://docs.google.com/forms/d/e/1FAIpQLSfNsqFiyVgWXlLsDpgfpYpeZKdsVjjXOaZnEV0HjC5lLo82Bg/viewform?usp=sharing&ouid=109672915073950352502) and share your feedback.

Thanks for being part of our community! We hope these improvements make your integration journey smoother than ever. As always, [sign up](https://console.victoriametrics.cloud/signup) to try everything for free with $200 credits — **no credit card required**.
