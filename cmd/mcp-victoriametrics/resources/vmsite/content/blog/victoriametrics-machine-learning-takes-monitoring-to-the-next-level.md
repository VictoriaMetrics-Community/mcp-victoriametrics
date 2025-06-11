---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2024-03-19
enableComments: true
title: "VictoriaMetrics Machine Learning takes monitoring to the next level"
summary: "Announcing  VictoriaMetrics Anomaly Detection solution, which harnesses machine learning to make database alerts more relevant, accurate and actionable for enterprise customers."
categories: 
 - Company News
tags:
 - anomaly detection
 - root cause analysis
 - machine learning
 - victoriametrics
 - monitoring
 - timeseries
 - alerting
 - alerts
images:
 - /blog/victoriametrics-machine-learning-takes-monitoring-to-the-next-level/preview-image-announcement-anomaly.webp
---

## Anomaly Detection empowers Enterprise IT teams overwhelmed by ‘sea of red’ alerts

Today we’re happy to announce our new <a href="/products/enterprise/anomaly-detection/" target="_blank">VictoriaMetrics Anomaly Detection</a> solution, which harnesses machine learning to make database alerts more relevant, accurate and actionable for enterprise customers.

VictoriaMetrics Anomaly Detection lightens the load on overworked data engineers, focusing their scarce resources on the alerts that matter most to their organization.

By unifying anomalies under a simple scoring system, VictoriaMetrics continues its mission to make monitoring even the most complex data sets simpler, more reliable, and more efficient.

### Conquering alert fatigue

As monitoring has spiralled in complexity, with databases becoming more interconnected and co-dependent, engineering teams can quickly become overwhelmed with alerts. Simplistic alerting is unable to distinguish minor performance concerns from potentially mission-critical outages.

To simplify monitoring for the very large datasets enterprises rely on today, we’ve developed <a href="/products/enterprise/anomaly-detection/" target="_blank">VictoriaMetrics Anomaly Detection</a> from scratch to identify data trends and alert only for signals that matter. Now, for the first time using neural networks, it is possible to set alerts which ‘understand’ data and so can draw conclusions from the data’s context. This turns the challenges of time-series data into a strength.

<p><i>
"Machine Learning is famously energy intensive, as a company that prides itself on efficiency we had to balance energy usage with the value it created for businesses. VictoriaMetrics Anomaly Detection is designed to be as efficient as the rest of our product range once calibrated to make sure businesses see a clear return on their investment"
</i></p>
<p style="text-align:right"><i>- Roman Khavronenko, co-founder VictoriaMetrics </i></p>

### Using intelligent analysis

Most alert systems use threshold alerting, where an alert is sent only if a value exceeds, or falls below, a predetermined range to signal systems operating outside normal tolerances. With the scalability and seasonality of modern real-time and distributed systems, alert thresholds need to be more complex if they are to offer control of an ever-evolving and scaling database.

Instead, VictoriaMetrics Anomaly Detection analyzes historical data and attaches an anomaly score to each data point indicating how far a signal deviates from the expected value or pattern. For engineers, it couldn't be simpler, whenever the anomaly value exceeds 1, an alert can be generated, taking the cognitive load off of engineers so they can focus on what matters.

### Trained in minutes

VictoriaMetrics monitoring tools are already used by some of the largest databases on the planet counting [Grammarly](/case-studies/grammarly/), [Wix](https://docs.victoriametrics.com/casestudies/#wixcom) and [CERN](https://docs.victoriametrics.com/casestudies/#cern) among their users. Anomaly detection can ingest the historic data businesses are already generating to calibrate itself, with minimal oversight.

As part of its commitment to efficiency, the VictoriaMetrics team designs all new technologies with the aim of reducing database workloads. Previously, database monitoring at this level required engineering teams continuously on-call; now, monitoring teams are augmented with an AI-like tool continuously observing the system. 

VictoriaMetrics Anomaly Detection can account for:

#### Seasonality

Machine learning understands context and intelligently adapts to changing data dynamics. 

#### Contextual anomalies

Anomaly Detection trained with historic data, allowing it to identify anomalies that would otherwise require an engineer familiar with the data set.

#### Collective anomalies

In isolation, concerning signals can go under the radar, continuously analyzing entire datasets to detect patterns all but senior engineers would miss.

#### Novelties

Anomaly Detection can detect ‘novelties’, or significant changes in the underlying system, intelligently adjusting  to a ‘New Normal’.

## Getting Started with VictoriaMetrics Anomaly Detection

Getting started is simple: Follow the [QuickStart guide](https://docs.victoriametrics.com/anomaly-detection/quickstart/), where you can find instructions on how to run VictoriaMetrics Anomaly Detection in Docker or Kubernetes.

* [Request a trial license](/products/enterprise/trial/)
* [Visit the Docker Hub Repository](https://hub.docker.com/r/victoriametrics/vmanomaly)
