---
title: "Welcome to VictoriaMetrics Docs"
layout: home
description: "Landing page for docs visitors from marketing site."

intro:
  heading: "Welcome to the VictoriaMetrics Documentation"
  text: "Find here all the relevant, technical information about our open source and enterprise observability solutions that you’ll need to efficiently query and visualize your metrics, logs, and traces."

sections:
  - title: "Open Source"
    items:
      - title: "VictoriaMetrics"
        url: "/victoriametrics/"
        text: "The fast, cost-effective, scalable monitoring solution and time series database."
        links:
          - label: "Quick start"
            url: "/victoriametrics/quick-start/"
          - label: "Best practices"
            url: "/victoriametrics/bestpractices/"
          - label: "Integrations"
            url: "/victoriametrics/integrations/"
          - label: "See all"
            url: "/victoriametrics/"
        upcoming: false

      - title: "VictoriaLogs"
        url: "/victorialogs/"
        text: "The simple, resource-efficient and fast logs database that scales."
        links:
          - label: "Quick start"
            url: "/victorialogs/quickstart/"
          - label: "Key concepts"
            url: "/victorialogs/keyconcepts/"
          - label: "See all"
            url: "/victorialogs/"
        upcoming: false

      - title: "VictoriaTraces"
        url: "/victoriatraces/"
        text: "Preview: Our new database designed for storing and querying distributed tracing data."
        links:
          - label: "Quick Start"
            url: "/victoriatraces/#quick-start"
          - label: "Key concepts"
            url: "/victoriatraces/keyconcepts/"
          - label: "See all"
            url: "/victoriatraces/"
        upcoming: false

  - title: "Enterprise"
    items:
      - title: "VictoriaMetrics Enterprise"
        url: "/victoriametrics/enterprise/"
        text: "Reliable, secure and cost-efficient monitoring for enterprises."
        links:
          - label: "Enterprise components"
            url: "/victoriametrics/enterprise/"
          - label: "Contact Us"
            url: "https://victoriametrics.com/contact-us/"
        upcoming: false

      - title: "VictoriaMetrics Cloud"
        url: "/victoriametrics-cloud/"
        text: "The managed, easy-to-use monitoring solution that integrates seamlessly with other tools and frameworks."
        links:
          - label: "Get started"
            url: "/victoriametrics-cloud/get-started/"
          - label: "Integrations"
            url: "/victoriametrics-cloud/integrations/"
          - label: "See all"
            url: "/victoriametrics-cloud/"
        upcoming: false

      - title: "VictoriaMetrics Anomaly Detection"
        url: "/anomaly-detection/"
        text: "The 'observability with AI' tool that automates the detection of anomalies in time-series data."
        links:
          - label: "Quick start"
            url: "/anomaly-detection/quickstart/"
          - label: "Components"
            url: "/anomaly-detection/components/"
          - label: "See all"
            url: "/anomaly-detection/"
        upcoming: false

resources_heading: "Tools & Resources"
resources:
  - title: "Guides"
    icon: "book"
    text: "From monitoring Kubernetes with VictoriaMetrics via how to use OpenTelemetry with our solutions and more, our guides provide the insight needed."
    url: "/guides/"

  - title: "Kubernetes Operator"
    icon: "kubernetes"
    text: "Run VictoriaMetrics applications on top of Kubernetes while preserving Kubernetes-native configuration options."
    url: "/operator/"

  - title: "Helm Charts"
    icon: "helm"
    text: "This repository provides all our helm charts for VictoriaMetrics and VictoriaLogs."
    url: "/helm/"
---

{{< ds-homepage >}}
