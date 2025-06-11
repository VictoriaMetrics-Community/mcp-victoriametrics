---
draft: false
page: blog blog_post
authors:
 - Denys Holius
 - Jean-Jerome Schmidt-Soisson
date: 2023-12-21
enableComments: true
title: "5 Year Anniversary Celebrations"
summary: "We’re celebrating 5 years of VictoriaMetrics and this blog post shares details on our top 5 stats, contributors, commentators and more! Happy anniversary!"
categories: 
 - Company News
tags:
 - victoriametrics
 - anniversary
images:
 - /blog/5-year-anniversary-celebrations/preview.webp
---

Happy 5th Anniversary, VictoriaMetrics!

What better way to wrap up the year than with a celebration!

5 years of VictoriaMetrics is a bit of a milestone for our team and - by extension - for all of our users, who have made our company and software what they are today. Many thanks to all of you.

We celebrated in style during [last week’s virtual meetup](https://www.youtube.com/watch?v=8qdvynaWDgQ&t=4985s), where we were lucky enough to be joined by guest speakers for the first time, who helped us make the occasion extra special. 

The number 5 is (unsurprisingly) the main theme of this anniversary blog post in which you’ll find: 

* Details about our [5 Year anniversary meet up](https://www.youtube.com/watch?v=8qdvynaWDgQ&t=4985s) (including links to the recording)
* Some of our top stats achieved in the past 5 years
* Our Top 5
  * Contributors
  * Commentators
  * Blog posts
  * 2024 predictions

Thanks to everyone who’s contributed and is still contributing to the VictoriaMetrics story - here’s to many happy more years!

## [5 Year Anniversary Meetup](https://www.youtube.com/watch?v=8qdvynaWDgQ&t=6s)

What's new in VictoriaMetrics - Roman Khavronenko

* VictoriaMetrics Cloud - Ivan Yatskevich
* Anomaly Detection - Daria Karavaieva / Fred Navruzov

Community Update: Our Vicky Story

* Egor Pronin (WEDOS)
* Andrey Tyulenev (Deutsche Bank)
* Carsten Aulbert (Max Planck Institute for Gravitational Physics, Albert Einstein Institute)
* Haley Wang: From Vicky contributor to Vicky team member

Community News - Denys Holius
* Latest VictoriaMetrics News - Jean-Jérôme Schmidt-Soisson

Party Time: Celebrating 5 Years of VictoriaMetrics BYOD (Bring Your Own Drink) for a toast 

{{< image href="/blog/5-year-anniversary-celebrations/celebrating.webp" alt="Party Time: Celebrating 5 Years of VictoriaMetrics BYOD (Bring Your Own Drink) for a toast" >}}

View the meetup: [https://www.youtube.com/watch?v=8qdvynaWDgQ&t=4985s](https://www.youtube.com/watch?v=8qdvynaWDgQ&t=4985s)

## 5 Years of VictoriaMetrics in Stats

* &#x23; of downloads: 338 million
* &#x23; of Slack users: 2,841
* &#x23; of contributors: 219
* &#x23; of issues 876 + 772 PRs
* &#x23; of releases, from 1.86 to 1.95:
  * 262 FEATURES
  * 334 BUG FIXES

5 Years of VictoriaMetrics Popularity

{{< image href="/blog/5-year-anniversary-celebrations/ranking.webp" alt="VictoriaMetrics DB Engines ranking" >}}

Source & more insights: [https://db-engines.com/en/ranking_trend/system/VictoriaMetrics](https://db-engines.com/en/ranking_trend/system/VictoriaMetrics)

## 5 Years of VictoriaMetrics on GitHub

{{< image href="/blog/5-year-anniversary-celebrations/stars-history.webp" alt="VictoriaMetrics GitHub stars history" >}}

Source & more insights: [https://ossinsight.io/collections/time-series-database/](https://ossinsight.io/collections/time-series-database/)

## Who’s In the VictoriaMetrics Galaxy

{{< image href="/blog/5-year-anniversary-celebrations/stargazers.webp" alt="VictoriaMetrics GitHub Stargazer's Companies" >}}

Source & more insights: [https://ossinsight.io/collections/time-series-database/](https://ossinsight.io/collections/time-series-database/)

## The VictoriaMetrics Top 5s

### Top 5 Commentators

* belm0 - 127
* n4mine - 67
* faceair - 65
* dxtrzhang - 55
* jiangxinlingdu - 53

### Top 5 Contributors

* faceair - 16
* belm0 - 12
* jiangxinlingdu - 8
* rodrigc - 6
* michal-kralik - 6

### Top 5 (Most Read) Blog Posts

**[Grafana Mimir & VictoriaMetrics: Performance Tests](/blog/mimir-benchmark/)**

In this blogpost, we compare the performance and resource usage of VictoriaMetrics and Grafana Mimir clusters running under moderate workload on the same hardware. In the first round Mimir and VictoriaMetrics were running under the same load and on the same hardware. Benchmark results were the following:

* VictoriaMetrics uses x1.7 less CPU for the same workload;
* VictoriaMetrics uses x5 less RAM for the same amount of active series;
* VictoriaMetrics uses x3 less storage space for the 24h of data collected during the benchmark.

[Read more!](/blog/mimir-benchmark/)

**[VictoriaMetrics Monitoring](/blog/victoriametrics-monitoring/)**

The recommendations in this post provide information and tools for maintaining a healthy and performant VictoriaMetrics installation. For enterprise users, we provide a [Monitoring of Monitoring](/products/mom/) service, where VictoriaMetrics team looks after installations, notifies about potential issues, and helps to build performant and reliable setups.

**[Pricing Comparison for Managed Prometheus](/blog/managed-prometheus-pricing/)**

In this post, we compare the cost of using managed services for Prometheus, which de-facto became a standard for modern monitoring. We look at what would be the cost of serving the same workload at Amazon Managed Service for Prometheus, Google Cloud Managed Service for Prometheus and [VictoriaMetrics Cloud](/blog/managed-victoriametrics-announcement/).

**[Cardinality Explorer](/blog/cardinality-explorer/)**

In monitoring, the term [cardinality](https://docs.victoriametrics.com/keyConcepts.html#cardinality) defines the number of unique time series stored in a time series database (TSDB). The higher the cardinality, the more resources are usually required for metrics processing and querying. According [to our observations], more than 50% of stored metrics are never used. So having an insight to the stored data and its structure can significantly improve the reliability and resource usage of the monitoring solution.

**[Save Network Costs with VictoriaMetrics Remote Write Protocol](/blog/victoriametrics-remote-write/)**

Prometheus remote write protocol is used by Prometheus for sending data to remote storage systems such as VictoriaMetrics. It serves well in most cases. But it isn’t optimized for low network bandwidth usage. This blog post outlines how the VictoriaMetrics remote write protocol allows reducing network traffic costs by 2x-4x compared to the Prometheus remote write protocol at the cost of slightly higher CPU usage (+10% according to our production stats).

## Our Top 5 Predictions for 2024

**Prediction 1:** AI will hide more problems than it will solve

AI powered, low code interfaces, will become popular additions to most data ops platforms next year. The goal will be to reduce increased operations complexity by asking GPT to write queries or control infrastructure. However, for most GPT outputs the user can assess the quality of the text generated. If an AI is used to translate and interpret between the user and the database, neither party can be sure that the answers are correct.

If the industry is serious about simplifying something, AI should be not only an input interface, but output as well. It should be able to tell whether the result of the query is what was intended. Expect a business to have an issue with an AI powered system misinterpreting a critical task in 2024.

**Prediction 2:** Data engineers will teach AI how to do their jobs

AI developers will incorporate tools to quickly fine tune custom models, model maintainers need to have an easy interface for re-training/correcting their models. Companies that intend to use these tools will need a faster feedback loop to acclimatize models to their new jobs.
Rather than large batches of corrections models can be calibrated output by output:

1. User asks question
1. Model  responds
1. User marks it as erroneous
1. Model owner receives the report and generates a correct response
1. Model owner feeds the response back to the model as correction step
1. Model improves

This is one path to fine-tune the models that will take over some aspects of data ops management. 

**Prediction 3:** Clarity on observability

The term observability is becoming too broad, either the category will split into newly defined roles or consolidate to better reflect the services covered.  Complexity can't grow all the time, because at some point it becomes unsustainable. So it is natural for simpler solutions to emerge and reduce the complexity.

It could be AI, why not? Or a new standard with generalised telemetry signals to reduce variability and complexity.

**Prediction 4:** Sustainable SaaS will create sustainable sales

Data ops solutions that prioritise sustainability will outperform competitors by being better value for money. Sustainability includes prioritising efficiency which will control costs. Providers that can achieve greater efficiency will out compete services that have become bloated with features and redundancies especially as legislation targets data centre energy consumption which currently accounts for up to 1.5% of global emissions.

**Prediction 5:** We will continue to make simple, reliable and cost-efficient open source monitoring software; and help users and organizations solve their monitoring and observability set ups, no matter the scale.

Here’s to a successful 2024 for us all!
