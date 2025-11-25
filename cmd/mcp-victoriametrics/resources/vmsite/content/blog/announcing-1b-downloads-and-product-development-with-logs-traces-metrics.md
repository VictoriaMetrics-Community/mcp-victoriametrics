---
draft: false
page: blog blog_post
authors:
  - Jean-Jerome Schmidt-Soisson
date: 2025-11-13
enableComments: true
featuredPost: true
title: "Announcing 1B+ Downloads & Product Development With Logs, Traces, Metrics"
summary: "This year saw us blast past the 1 billion downloads on Docker, which is fueled by our customer-centric approach to software development and the introduction of new open source projects, such as our new database for traces. Read this blog post to learn more about our 2025 milestones."
categories:
 - Company News
 - Product News
tags:
 - open source
 - innovation
 - monitoring
 - observability
 - docker
 - victoriametrics
 - victorialogs
 - victoriatraces
images:
  - /blog/announcing-1b-downloads-and-product-development-with-logs-traces-metrics/preview.webp
---

We're currently at KubeCon + CloudNativeCon North America 2025 in Atlanta, and it’s a great opportunity to connect with the community and share some of the progress we've made this year.

It's been a busy period of development, new releases, and community engagement, all guided by our focus on delivering simple, reliable, and efficient monitoring & observability solutions.

## New products and customer-centric approach drive adoption & expansion: Hitting one billion downloads

This year saw us blast through the one billion downloads milestone on [Docker Hub](https://hub.docker.com/u/victoriametrics). This adoption metric is a great Indicator of the trust that users place in our observability platform for the developer and enterprise communities.

Here are some of our main highlights this year:

* Hit one billion software downloads on Docker Hub 
* Released the much-requested [cluster version of VictoriaLogs](https://docs.victoriametrics.com/victorialogs/cluster/)
* Introduced our new open source database for traces, [VictoriaTraces](https://docs.victoriametrics.com/victoriatraces/)
* Passed 10,000 PRs in the [VictoriaMetrics repository](https://github.com/VictoriaMetrics/VictoriaMetrics/pulls): Thanks to all our users!
* Expanded global headcount by 50%
* New customer & user stories by [Spotify R&D](https://www.youtube.com/watch?v=87koDlpKDR4), [Prezi](https://www.infoq.com/news/2025/02/prezi-prometheus-victoriametrics/), [Dreamhost](https://victoriametrics.com/blog/dreamhost-case-study/), [Hudson River Trading](https://www.hudsonrivertrading.com/hrtbeat/heraclesql-a-python-dsl-for-writing-alerts/), and more…

## When our users & customers talk about our solutions

**Spotify R&D** [successfully migrated](https://www.youtube.com/watch?v=87koDlpKDR4) from their internal observability system - Heroic - to VictoriaMetrics. This resulted in significant cost savings and improved user experience.

We've seen significant adoption this year, and we measure our success by the real-world problems our users solve.

Another nice example is our recent <a href="https://victoriametrics.com/blog/dreamhost-case-study/"><b>DreamHost</b></a>  case study: When they migrated to VictoriaMetrics, their goal was to solve scaling and resource consumption challenges.

The results demonstrate the efficiency of our products: 
* 80% reduction in memory usage compared to their previous solution
* Scaled to 76 million active time series
* Handles ingestion of over 450,000 data points per second

As Jordan Tardif, Distinguished Engineer at DreamHost, stated, VictoriaMetrics is "Prometheus, that scales with way less effort & resources." 

Our Co-Founder and CEO, Artem Navoiev explains, "This year has been a significant step forward for our company. Staying self-funded thanks to our amazing customers, we continued to focus on creating simple yet efficient products for engineers. In 2026, we’re determined to go further: Triple down on open source, further improve our metrics, logging and tracing solutions, and stick to our commitment to solve the main observability problem."

## More product development and key releases

Our engineering efforts have been focused on both long-term stability and the integration of new capabilities.

Here are a few additional key product-related announcements from this year:

* <a href="https://victoriametrics.com/blog/victoriametrics-long-term-support-h2-2025-update/"><b>New LTS Release</b></a>: We've shipped new Long-Term Support (LTS) versions, which provide the stability many enterprises require for critical production workloads.
* <a href="https://victoriametrics.com/blog/q3-2025-whats-new-victoriametrics-cloud/"><b>VictoriaMetrics Cloud Expansion</b></a>: To support our global users, our cloud platform has expanded into new regions, including a new deployment in Asia (ap-southeast-1). VictoriaMetrics Cloud scales effortlessly while reducing monitoring costs by 5x.
* **MCP Server for AI Integration**: We released the new VictoriaMetrics MCP server. This tool uses the Model Context Protocol to allow AI and LLM tools to use your VictoriaMetrics deployment as a data source, enabling natural language querying of your observability data.
* <a href="https://docs.victoriametrics.com/"><b>AI-Powered Documentation</b></a>: To improve user experience, we’ve integrated a new chatbot into our official documentation. It's designed to help users find technical answers and navigate our resources more efficiently.
* **Dedicated Grafana data sources for both** <a href="https://grafana.com/grafana/plugins/victoriametrics-metrics-datasource/"><b>VictoriaMetrics</b></a> **and** <a href="https://grafana.com/grafana/plugins/victoriametrics-logs-datasource/"><b>VictoriaLogs</b></a>: These new data sources provide seamless integration and enhanced monitoring capabilities, allowing engineers to leverage the powerful features of Grafana with the efficiency and reliability of VictoriaMetrics and VictoriaLogs.

## Welcoming new talent to drive innovation

To keep pace with the soaring demand for our time-series data, logs, and traces solutions, we’ve expanded our team by 50%. New additions include colleagues in Customer Success, Developer Advocacy, and engineers bolstering both open-source and enterprise offerings. This strategic growth underscores our dedication to delivering top-notch support and innovative solutions to our expanding and diverse customer base.

## Community and events

Connecting with the open-source community remains a priority. This year, our team has participated in numerous events, including LinuxFest Northwest, KubeCon Europe, SREday Cologne, PromCon EU, and many KubeCon CloudNativeDay (KCD) gatherings.

We've continued that at KubeCon Atlanta, which is happening this week. We were there to discuss technical challenges related to high-cardinality, multi-tenancy, and long-term storage, or to demonstrate how VictoriaMetrics can help simplify your observability stack.

<p style="max-width: 800px; margin: 1rem auto;">
    <img src="/blog/announcing-1b-downloads-and-product-development-with-logs-traces-metrics/victoriametrics-team-at-kubecon-2025-atlanta.webp" style="width:100%" alt="VictoriaMetrics' Team at the KubeCon 2025 in Atlanta, Georgia">
</p>

Next week we’ll be at the [Open Source Monitoring Conference (OSMC)](https://osmc.de/) in Nuremberg, Germany with four talks. We’ll also have a booth at the conference, so please do come by. We’d love to talk to you and help answer any questions you may have on anything related to open source monitoring and observability. See you there!