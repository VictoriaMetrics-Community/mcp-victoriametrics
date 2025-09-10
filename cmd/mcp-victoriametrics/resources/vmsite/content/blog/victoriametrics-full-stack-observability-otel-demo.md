---
draft: false
page: blog blog_post
authors: 
 - Diana Todea
date: 2025-09-10
title: "Full-Stack Observability with VictoriaMetrics in the OTel Demo"
summary: "The OpenTelemetry Astronomy Shop demo has long served as a reference environment for exploring observability in distributed systems, but until now it shipped with only a Prometheus datasource. VictoriaMetrics forked the demo and extended it with VictoriaMetrics, VictoriaLogs, and VictoriaTraces, providing insights into VictoriaMetrics’ observability stack where metrics, logs, and traces flow into a unified backend."
enableComments: true
categories: 
 - OpenTelemetry
 - Observability
 - Monitoring
tags: 
 - opentelemetry
 - observability
 - monitoring
 - victoriametrics
 - victoriatraces
 - victorialogs
images:
 - /blog/vm-full-stack-observability-otel-demo/preview.webp
---
The [OpenTelemetry Astronomy Shop](https://github.com/open-telemetry/opentelemetry-demo) is a widely used demonstration environment designed to illustrate the concepts and practical implementation of observability in distributed systems. Built as a microservice-based e-commerce application, the demo provides developers with a near real-world environment where they can explore how telemetry data—metrics, logs, and traces—can be collected, processed, and visualized. Its value lies in offering a hands-on way to understand [OpenTelemetry (OTel)](https://opentelemetry.io/docs/what-is-opentelemetry/) as a standard for instrumenting modern applications.

[OTLP](https://opentelemetry.io/docs/specs/otel/protocol/) works and offers seamless integration with any other compatible solution, with no vendor lock-in. However, until now, the demo shipped with only a [Prometheus](https://grafana.com/docs/grafana/latest/datasources/prometheus/) datasource. To enable developers to experiment with the [VictoriaMetrics](https://victoriametrics.com/) databases, our engineers forked the OTel demo repository and extended it with the **VictoriaMetrics observability ecosystem**. Our contribution adds [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) for metrics, [VictoriaLogs](https://github.com/VictoriaMetrics/VictoriaLogs) for logs, [VictoriaTraces](https://github.com/VictoriaMetrics/VictoriaTraces) for distributed traces, thus making it possible to experiment with all 3 signals.

Getting started with this enhanced demo is straightforward. Developers only need to clone the forked demo: [VictoriaMetrics-Community/opentelemetry-demo](https://github.com/VictoriaMetrics-Community/opentelemetry-demo) and launch the environment with a single command. The services start automatically, and Grafana is preconfigured with [VictoriaMetrics datasource](https://docs.victoriametrics.com/victoriametrics-datasource/), [VictoriaLogs datasource](https://docs.victoriametrics.com/victorialogs/victorialogs-datasource/), [VictoriaTraces in Grafana](https://docs.victoriametrics.com/victoriatraces/querying/grafana/) (via [Grafana Jaeger datasource](https://grafana.com/docs/grafana/latest/datasources/jaeger/)).

![Grafana dashboard OTel Demo](/blog/vm-full-stack-observability-otel-demo/service-otel-demo.webp)  

## Exploring Telemetry

Metrics can be explored using [MetricsQL](https://docs.victoriametrics.com/victoriametrics/metricsql/) query language, logs can be searched through [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/) , and traces can be visualized with VictoriaTraces—all accessible from the Grafana interface. In this way, the demo not only illustrates how [OpenTelemetry instrumentation](https://opentelemetry.io/docs/concepts/instrumentation/) works but also demonstrates how a single, scalable backend can handle diverse observability signals.

![VictoriaMetrics datasource](/blog/vm-full-stack-observability-otel-demo/vm-grafana.webp)

![VictoriaLogs datasource](/blog/vm-full-stack-observability-otel-demo/vlogs-grafana.webp)

![VictoriaTraces datasource](/blog/vm-full-stack-observability-otel-demo/vtraces-grafana.webp)

## Benefits to the ecosystem

The benefits of this integration extend beyond convenience. By replacing the Prometheus-only setup with a full-stack VictoriaMetrics backend, the demo now mirrors production-ready observability architectures more closely. 
[VictoriaMetrics](https://docs.victoriametrics.com/quick-start/) is recognized for its scalability, cost efficiency, and [performance](https://victoriametrics.com/blog/opentelemetry-prometheus-and-more/), which means that the lessons developers learn in the demo can be directly applied to real-world environments. The unified stack also reduces complexity, since there is no longer a need to maintain separate systems for metrics, logs, and traces.

To showcase the new functionality, we have provided several examples within the demo. Developers can run MetricsQL queries in Grafana to analyze application performance, inspect structured logs to debug service behavior, and follow distributed traces across microservices to understand request flows. These features together create a richer, more holistic learning experience than was previously possible.

![Span Metrics](/blog/vm-full-stack-observability-otel-demo/spanmetrics.webp)

## OpenTelemetry Collector Dashboard

An important component of the enhanced demo is the **OpenTelemetry Collector Grafana dashboard**. This dashboard provides a consolidated view of key performance indicators for the system by presenting RED metrics—rate, errors, and duration—derived from the spanmetrics connector in the OpenTelemetry Collector. Developers can select a service of interest and immediately examine its RED metrics in order to evaluate performance and reliability. If the chosen service emits logs, these logs are displayed alongside the metrics, offering valuable context for troubleshooting and analysis. In addition, custom metrics produced by individual services are integrated into the same view, further enriching the observability landscape. It is worth noting that, due to the nature of telemetry aggregation, chart panels may require several minutes after the demo is first launched before meaningful data becomes visible. 

This dashboard thus serves as a crucial entry point for developers, bringing together metrics and logs in a single interface and demonstrating how the [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/) can generate actionable insights when paired with the [VictoriaMetrics backend](https://docs.victoriametrics.com/guides/getting-started-with-opentelemetry/).

![Receivers](/blog/vm-full-stack-observability-otel-demo/vm-receivers.webp)

![Exporters](/blog/vm-full-stack-observability-otel-demo/vm-exporters.webp)

![Signal Flows](/blog/vm-full-stack-observability-otel-demo/signal-flows.webp)

## Conclusion

The VictoriaMetrics fork of the OpenTelemetry Astronomy Shop demo provides insights into the VictoriaMetrics observability stack through the integration of VictoriaMetrics, VictoriaLogs, and VictoriaTraces. We invite the community to explore the [demo](https://github.com/VictoriaMetrics-Community/opentelemetry-demo), experiment with the extended capabilities and contribute to further improvements.

By doing so, developers will not only deepen their understanding of OpenTelemetry but also experience firsthand how **simple and effective unified observability** with VictoriaMetrics can be.

Check out our VictoriaMetrics, VictoriaLogs, VictoriaTraces for more details and get involved with our [community](https://victoriametrics.com/community/). If you spot anything outdated or have questions, don’t hesitate to reach out. You can drop me a DM on [LinkedIn](https://www.linkedin.com/in/diana-todea-b2a79968/).

## Related Articles

- [OTel me more about Traces: Introducing VictoriaMetrics’ Trace Analyzer](https://medium.com/@dianatodea/otel-me-more-on-traces-introducing-victoriametrics-trace-analyzer-3db76d8a2bb2)
- [I played with OTel and I liked it](https://medium.com/@dianatodea/i-played-with-otel-and-i-liked-it-a69af868568b)
