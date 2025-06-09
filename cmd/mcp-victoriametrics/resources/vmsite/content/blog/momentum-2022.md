---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2023-01-25
enableComments: true
title: "Monitoring the Universe & Beyond: Our 2022 in Review"
summary: "This ‘VictoriaMetrics 2022 Momentum Milestones’ blog post provides a summary of this year’s main achievements with our top features, blogs and talks."
categories: 
 - Company News
tags:
 - victoriametrics
 - momentum
 - 2022
 - achievements
 - open source
 - database
 - monitoring
 - timeseries
images:
 - /blog/momentum-2022/twitter-asset-100M.webp
---

When we posted our first ever Momentum blog about a year ago detailing our 2021 achievements, we were just weeks away from Russia’s renewed attack of Ukraine. While the war isn’t won yet and we’re approaching the one year anniversary of the attack, it’s heartening to see how much has changed around the world and that almost everyone now knows the expression: Slava Ukraini!

So if we had to choose one word to best describe 2022 it might be: Resilience.

What’s your word for 2022? Please feel free to share in the comments below!

## **Our 2022: 100M Downloads!**

Overall, the past year has taught us to keep going and that everything is possible: even monitoring the universe - and reaching 100M downloads of our software!

As our co-founder, Roman recently tweeted: “It was a tough year. But I'm glad VictoriaMetrics continues to accelerate!”

{{< image href="/blog/momentum-2022/tweet.webp">}}

We were delighted to be able to share our stories of how VictoriaMetrics is being used by organizations such [Open Cosmos](https://www.siliconrepublic.com/start-ups/victoriametrics-data-monitoring-satellites) and [CERN’s CMS team](https://dcnnmagazine.com/data/victoriametrics-monitoring-cern/), where our technology is literally being used to monitor the universe - as well as many other similar stories.

Thanks to everyone who’s used and/or has contributed to VictoriaMetrics in 2022 - we’re looking forward to many more such interactions  this year.

And whether you are a longtime or new member of the VictoriaMetrics Community, please share in last  year’s success with our main highlights and stats.

## **2022 Momentum Milestones**

- 100M+ Downloads
- 10K+ GitHub Stars for [VictoriaMetrics overall](https://coderstats.net/github/#victoriametrics)
- Triple-digit Enterprise growth
- New customer and user stories with Ably, Semrush, CERN, Open Cosmos, and more
- [Launched Managed VictoriaMetrics](https://www.theee.ai/2022/10/28/21495-monitoring-with-fully-managed-service-on-aws-is-simplified-says-victoriametrics/)
- Featured in Forbes, The Times, Silicon Republic, Diginomica, TechTarget, and many more

### **VictoriaMetrics Highlights**

- 650+ issues
- 600+ PRs
- 189 contributors
- 31 releases, from 1.72 to 1.85:
  - 315 Features
  - 215 Bug fixes


We have a very short and prolific release cycle, which can be followed on our GitHub page: [https://github.com/VictoriaMetrics](https://github.com/VictoriaMetrics)

Thanks again to everyone who’s contributed to VictoriaMetrics in the past year!

We are proudly a self-funded startup that generates profitability from [services that we offer](https://victoriametrics.com/support/) in support to VictoriaMetrics as well as its [Enterprise version](https://victoriametrics.com/products/enterprise/). Our team is laser-focused on solving our customer and community user needs, while constantly perfecting and enhancing our software.

This blog post provides a summary of our main achievements in 2022 with our top features, blogs and talks.


### **[Top New VictoriaMetrics Features](https://www.youtube.com/watch?v=Mesc6JBFNhQ&t=660s)**

#### MetricsQL Features

- Support for @ modifier
- keep_metric_names modifier
- Advanced label filters’ propagation
- Automatic label filters’ propagation
- Support for short numeric constants
- Distributed query tracing!
- New functions

See here for all MetricsQL features details: [https://docs.victoriametrics.com/metricsql/](https://docs.victoriametrics.com/metricsql/)

#### vmui Features

- Cardinality explorer!
- Top queries
- Significantly improved usability and stability!

Visit our vmui playground for details: [vmui playground](https://play.victoriametrics.com/select/accounting/1/6a716b0f-38bc-4856-90ce-448fd713e3fe/prometheus/graph/?g0.expr=100%20*%20sum(rate(process_cpu_seconds_total))%20by%20(job)&g0.range_input=1d&_gl=1*1wt77yi*_ga*MjA3MDAxOTQwMy4xNjYxNDMyMTYy*_ga_N9SVT8S3HK*MTY3NDY0NDQ5NC4zNzYuMC4xNjc0NjQ0NDk0LjAuMC4w)

#### vmagent Features

- Fetch target response on behalf of vmagent
- Filter targets by url and labels
- /service-discovery page
- Relabel debugging!
- support for absolute `__address__`
- New service discovery mechanisms
- Multi-tenant support
- Performance improvements

See here for all vmagent features details: [https://docs.victoriametrics.com/vmagent.html](https://docs.victoriametrics.com/vmagent.html)

#### Relabeling Features

- Conditional relabeling
- Named label placeholders
- Graphite-style relabeling

See here for all relabeling features details: [https://docs.victoriametrics.com/vmagent.html#relabeling](https://docs.victoriametrics.com/vmagent.html#relabeling)

#### vmalert Features

- Better integration with Grafana alerts
- Reusable templates for annotations
- Debugging of alerting rules
- Improved compatibility with Prometheus

See here for all vmalert features details: [https://docs.victoriametrics.com/vmalert.html](https://docs.victoriametrics.com/vmalert.html)

#### vmctl Features

- Migrate all the data between clusters
- Data migration via Prometheus remote_read protocol

See here for all vmctl features details: [https://docs.victoriametrics.com/vmctl.html](https://docs.victoriametrics.com/vmctl.html)


For a complete review of all our new features in 2022, please [watch the presentation by Aliaksandr Valialkin](https://www.youtube.com/watch?v=Mesc6JBFNhQ&t=660s) recorded at our December 2022 Meet Up.

### **Top New VictoriaMetrics Enterprise Features**

- mTLS support
- vmgateway JWT token enhancements
- Automatic restore from backups
- Automatic vmstorage discovery
- Multiple retentions

See here for all VictoriaMetrics Enterprise features details: [https://docs.victoriametrics.com/enterprise.html](https://docs.victoriametrics.com/enterprise.html)


### **Top 3 Blogs**

- [Grafana Mimir and VictoriaMetrics: Performance Tests](https://victoriametrics.com/blog/mimir-benchmark/)
- [How to Choose a Scalable Open Source Time Series Database: The Cost of Scale](https://victoriametrics.com/blog/the-cost-of-scale/)
- [Pricing comparison for Managed Prometheus](https://victoriametrics.com/blog/managed-prometheus-pricing/)


### **Top 3 Talks**

- [OSMC 2022 | VictoriaMetrics: scaling to 100 million metrics per second](https://www.youtube.com/watch?v=xfed9_Q0_qU&list=LLZikvGfLwcOoapkibw5AcOw)
  - By Aliaksandr Valialkin
- [OSA Con 2022: Specifics of Data Analysis in Time Series Databases](https://www.youtube.com/watch?v=_zORxrgLtec&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj)
  - By Roman Khavronenko
- [CNCF Paris Meetup 2022-09-15 - VictoriaMetrics - The cost of scale in Prometheus ecosystem](https://www.youtube.com/watch?v=gcZYHpri2Hw&list=PLXT8DSiuv5ylmEbeWptT-512GpOF8_Ppj&index=7)
  - By Aliaksandr Valialkin

### **2023 Roadmap Highlights**

- [Grafana datasource plugin](https://github.com/VictoriaMetrics/grafana-datasource)
  - Completed in January 2023
- [Streaming aggregation](https://docs.victoriametrics.com/stream-aggregation.html)
  - Completed in January 2023
- vmalert: UI for rules management
- vmalert hysteresis support
- [vmui explore tab](https://play.victoriametrics.com/select/accounting/1/6a716b0f-38bc-4856-90ce-448fd713e3fe/prometheus/graph/?g0.range_input=1d&g0.end_input=2023-01-25T11%3A09%3A10&g0.step_input=6m40s&g0.relative_time=none&size=medium#/metrics)
  - Completed in January 2023
- VictoriaLogs

View Roman Khavronenko’s [2023 Roadmap presentation](https://www.youtube.com/watch?v=Mesc6JBFNhQ&t=3300s) for all the details!

Thanks again for your support this past year and have a successful 2023!

PS.: If you would be interested in learning more about our Enterprise features or getting more personalized support, [please click here](https://victoriametrics.com/products/).




