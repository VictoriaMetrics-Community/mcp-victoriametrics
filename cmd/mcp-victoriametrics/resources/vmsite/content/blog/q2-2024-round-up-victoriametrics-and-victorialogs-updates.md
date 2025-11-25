---
raft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2024-07-05
title: "Q2 2024 Round Up: VictoriaMetrics & VictoriaLogs Updates"
summary: "Read about our Q2 2024 achievements, the roadmap for VictoriaMetrics, the latest news on VictoriaLogs, and more!"
enableComments: true
categories:
 - Company News
 - Product News
tags:
 - victoriametrics
 - roadmap
 - achievements
 - open source
 - database
 - monitoring
 - timeseries
 - victorialogs
 - logs management
 - observability
images:
 - /blog/q2-2024-round-up-victoriametrics-and-victorialogs-updates/preview.webp
---
Many thanks to everyone who joined us for our recent [virtual meetup](https://www.youtube.com/watch?v=hzlMA_Ae9_4), during which we discussed some of our Q2 2024 highlights, including features highlights, the 2024 roadmap for VictoriaMetrics and all the latest news on VictoriaLogs!

In this blog post, we’d like to share a summary of these highlights.

## Latest Community Stats

* Downloads: 650+ million
* [Slack](https://slack.victoriametrics.com/) users: 3.4K
* [Telegram](https://t.me/VictoriaMetrics_en) users: 2K
* Contributors: 254
* 20 (!) [VictoriaLogs](https://victoriametrics.com/products/victorialogs/) releases 😎
* 15K stars on GitHub

## What's new in VictoriaMetrics at Q2 2024

### New Features Highlights

**Let's Encrypt support**

All the VictoriaMetrics Enterprise components support automatic issuing of TLS certificates for public HTTPS servers via Let’s Encrypt service. [Read more](https://docs.victoriametrics.com/#automatic-issuing-of-tls-certificates).

**VictoriaMetrics k8s operator**

Operator simplifies VictoriaMetrics cluster installation, upgrading and managing. [See all the latest features](https://github.com/VictoriaMetrics/operator/releases).

**[Helm chart: victoriametrics-distributed](https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-distributed)**

This chart sets up multiple VictoriaMetrics cluster instances on multiple Availability Zones:

* Improved reliability
* Faster read queries
* Easy maintenance

See all the latest features updates [here](https://docs.victoriametrics.com/changelog/).

## VictoriaMetrics in the Cloud: Update

* [VictoriaMetrics Cloud](https://victoriametrics.com/products/cloud/) trial period extended to 30 days, which starts when you create your first deployment
* New API Keys
  + Available in global and per-deployment scope
  + Alerting and recording rules management
* New Alerting and Recording Rules
  + Upload recording and alerting rules via UI and API
  + No additional cost
  + Internal Alertmanager
  + Dedicated installation for vmalert and alertmanager per deployment
  + Overview of all alerts
  + Deep overview of single alert
  + Silences and full control of AM configuration
* Update of Monitoring tab
  + Better UI and navigation
  + Explanation for every panel
* Updates re: integrations
  + Managed Prometheus
  + Compatibility With Popular Ingestion Protocols
    - OpenTelemetry, InfluxDB, DataDog, NewRelic, OpenTSDB & Graphite
  + Grafana Cloud
  + AWS PrivateLink

## The Latest on VictoriaLogs: Happy 1 Year Anniversary

* [Open source database for logs](https://victoriametrics.com/products/victorialogs/)
* Easy to setup and operate - just a single executable with sane default configs
* Works great with both structured and plaintext logs
* Uses up to 30x less RAM and up to 15x disk space than Elasticsearch
* Provides simple yet powerful query language for VictoriaLogs - [LogsQL](https://victoriametrics.com/products/logsql/)

{{<image href="/blog/q2-2024-round-up-victoriametrics-and-victorialogs-updates/victorialogs-one-year-anniversary.webp">}}

## New Features Highlights

* Improved querying HTTP API
* Add `/select/logsql/tail` HTTP endpoint, which can be used for live tailing of [LogsQL query results](https://docs.victoriametrics.com/victorialogs/logsql/)
* Data ingestion via Syslog protocol
* Automatic parsing of Syslog fields
* Supported transports:
  + UDP
  + TCP
  + TCP+TLS
* Gzip and deflate compression support
* Ability to configure distinct TCP and UDP ports with distinct settings
* Automatic log streams with (hostname, app_name, app_id) fields

[See all the latest VictoriaLogs features](https://docs.victoriametrics.com/victorialogs/changelog/)

## LogsQL Improvements

The latest set of features is focused on analytics enhancement:

We’re happy to share that our query language LogsQL can now be used for logs analysis when data is wrapped in different ways (json, syslog), extracting subfields from messages, and support for enhancement filters.

* Filtering shorthands
* week_range and day_range filters
* Limiters
* Log analytics
* Data extraction & transformation
* Additional filtering
* Sorting

For more details, visit the [LogsQL documentation](https://docs.victoriametrics.com/victorialogs/logsql/)

## VictoriaLogs Roadmap

* Accept logs via OpenTelemetry protocol
* VMUI improvements based on HTTP querying API
* Improve Grafana plugin for VictoriaLogs
* Cluster version
  + Try single-node VictoriaLogs: It can replace 30-node Elasticsearch cluster in production
* Transparent historical data migration to object storage such as S3
  + Try single-node VictoriaLogs with persistent volumes: It compresses 1TB of production logs from Kubernetes to 20GB

View the [VictoriaLogs roadmap](https://docs.victoriametrics.com/victorialogs/roadmap/)

Also check the new [VictoriaLogs PlayGround ](https://play-vmlogs.victoriametrics.com/)

## [Q2 2024 VictoriaMetrics Virtual Meetup](https://www.youtube.com/watch?v=hzlMA_Ae9_4&t=207s)

For a more detailed description and discussion of the topics covered in this blog, please watch the recording of our second virtual meetup this year.

Many thanks in particular to Alexis Ducastel from Blackswift & Calum Miller from Millersoft for their guest talks!

Thanks to all of you who participated in the meetup - you can watch the recording here, and [Subscribe to our YouTube channel](https://bit.ly/VictoriaMetrics-Youtube).

[VictoriaMetrics Meetup June 2024 - VictoriaLogs Update](https://www.youtube.com/watch?v=hzlMA_Ae9_4&t=207s)

## Community Podcasts & Talks

* [Measure PromQL / MetricsQL Expression Complexity](https://www.youtube.com/watch?v=lDyxBqoC_ww&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj&index=2) | Roman Khavronenko | Conf42 Observability 2024
* [Monitoring the Future: Insights from VictoriaMetrics](https://www.youtube.com/watch?v=FUDUCfG-oBo&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj&index=1) with Aliaksandr Valialkin | Live On Production Podcast
* Roman Khavronenko - [How to monitor the monitoring in open source projects](https://www.youtube.com/watch?v=CJcKUIoD-gs&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj&index=4) | Schrödinger Hat
* Roman Khavronenko: [Grafana Mimir and VictoriaMetrics: Performance Tests](https://www.youtube.com/watch?v=tsNbDdGEjoo&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj&index=5) | Data Miner
* [VictoriaMetrics Internals with Alex and Roman @victoriametrics](https://www.youtube.com/watch?v=umEw0PODDjs&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj&index=3) | The Geek Narrator
* [Solving monitoring scalability challenges with VictoriaMetrics](https://www.youtube.com/watch?v=-jzNpa-3JMQ) | ITGix
* In French: [Observabilité : dépoussiérer Prometheus avec VictoriaMetrics](https://www.youtube.com/watch?v=bzLtWjUj2k0) | Julien Briault | Devoxx France
* In French (new episodes): [Série sur VictoriaMetrics, une technologie prometteuse pour le stockage, la collecte et l'exploitation des métriques](https://www.youtube.com/watch?v=5256wndpyOI&t=35s) | xavki

## Join a VictoriaMetrics User Meetup Group

We’re starting to organize in-person meetups as well, please do join us for one of these groups. If you’d like to start a group in a city that is not listed here yet, please let us know in the comments.

* [Berlin](https://www.meetup.com/monitoring-observability-victoriametrics-berlin/)
* [Kraków](https://www.meetup.com/monitoring-observability-victoriametrics-krakow/)
* [London](https://www.meetup.com/monitoring-and-observability-victoriametrics-london/)
* [Paris](https://www.meetup.com/open-source-monitoring-observability-victoriametrics/)
* [San Francisco](https://www.meetup.com/monitoring-observability-victoriametrics-san-francisco/)
* Any other cities? Let us know!

## Where to Meet Our Team: Upcoming Talks & Events

* [DevConf.US, Boston](https://pretalx.com/devconf-us-2024/talk/XPTRR3/), 14-16/08
* [KubeCon China](https://events.linuxfoundation.org/kubecon-cloudnativecon-open-source-summit-ai-dev-china/), Hong Kong, 21-23/08
* [PromCon Europe](https://promcon.io/2024-berlin/), Berlin, 11&12/09
* [KubeCon North America](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/), Salt Lake City, 12-15/11
* Look out for dates of upcoming user meetups

## VictoriaMetrics in the News!

We’ve had some nice press coverage in the past few months as well as some community articles that have been published - thank you for these, and please let us know if you’d like to talk to us about what we do, or if you have any suggestions on articles that we could help with also.

* [KubeCon24: VictoriaMetrics' Simpler Alternative to Prometheus](https://thenewstack.io/kubecon24-victoriametrics-simpler-alternative-to-prometheus/)
* [VictoriaMetrics Machine Learning takes monitoring to the next level](https://digitalisationworld.com/news/67266/victoriametrics-machine-learning-takes-monitoring-to-the-next-level)
* [OpenTelemetry Is Too Complicated, VictoriaMetrics Says](https://www.datanami.com/2024/04/01/opentelemetry-is-too-complicated-victoriametrics-says/)
* [AI and Machine Learning will not save the planet (yet)](https://www.techradar.com/pro/ai-and-machine-learning-will-not-save-the-planet-yethttps:/)
* [How VictoriaMetrics' open source approach led to mass industry adoption](https://tech.eu/2024/05/03/how-victoriametrics-open-source-approach-led-to-mass-industry-adoption/https:/)
* [Green coding - VictoriaMetrics: The efficiency vs complexity trade-off](https://www.computerweekly.com/blog/CW-Developer-Network/Green-coding-VictoriaMetrics-The-efficiency-vs-complexity-trade-off)
* [How will energy prices affect data center costs in 2024 and beyond?](https://www.itpro.com/infrastructure/data-centres/how-will-energy-prices-affect-data-center-costs-in-2024-and-beyond)
* [VictoriaMetrics slashes data storage bills by 90%](https://msp-channel.com/news/67792/victoriametrics-slashes-data-storage-bills-by-90)
* [Millions forced to use brain as OpenAI's ChatGPT takes morning off](https://www.theregister.com/2024/06/04/openai_chatgpt_outage/)

## Our Recent Blog Posts

* [Monitoring Proxmox VE via Managed VictoriaMetrics](https://victoriametrics.com/blog/proxmox-monitoring-with-vmcloud/)
* [Introduction to Managed Monitoring](https://victoriametrics.com/blog/introduction-to-managed-monitoring/)
* [How to reduce expenses on monitoring: be smarter about data](https://victoriametrics.com/blog/reducing-costs-p2/)
* [VictoriaMetrics slashes data storage bills by 90% with world’s most cost-efficient monitoring](https://victoriametrics.com/blog/victoriametrics-slashes-data-storage-bills-with-worlds-most-cost-efficient-monitoring/)
* [How ilert Can Help Enhance Your Monitoring With Its VictoriaMetrics Integration](https://victoriametrics.com/blog/using-victoriametrics-and-ilert/)

This sums up our Q2 2024!

As always, we welcome your feedback and questions, so feel free to use the comments box below!
