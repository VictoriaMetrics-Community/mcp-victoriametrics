---
draft: false
page: blog blog_post
authors:
  - Marc Sherwood
date: 2025-08-20
title: "vmanomaly Deep Dive: Smarter Alerting with AI (Tech Talk Companion)"
summary: "Tech Talk: In this post, we explore vmanomaly through the eyes of its creators. Learn how this AI-powered alerting system helps cut through noise, avoid static rule spaghetti, and deliver actionable insights directly from your monitoring data."
enableComments: true
categories:
  - Tech Talk
tags:
  - VictoriaMetrics
  - Tech Talk
  - vmanomaly
  - Events
  - Video
images:
  - /blog/tech-talk-streems/vm-anomaly-deep-dive/preview.webp
---

## Our vmanomaly Deep Dive: My Favorite Takeaways from the Tech Talk

![VM Anomaly Deep Dive: Smarter Alerting with AI (Tech Talk Companion)](/blog/tech-talk-streems/vm-anomaly-deep-dive/preview.webp)

I was thrilled to host our [latest tech talk](https://www.youtube.com/watch?v=Uuix_glPfjM), where we got to do a deep dive into vmanomaly with the best possible guests: Fred Navruzov, the actual team lead for the product, and Co-Host, Matthias Palmersheim.

We covered a ton of ground, from high-level concepts to the nitty-gritty of configuration. For everyone who couldn't make it, I wanted to share my personal recap of the most important technical takeaways from our conversation.

### The "Why": Moving Beyond Brittle, Static Alerts

A topic that always comes up is alert fatigue. We've all seen those alerting rule sets that become pure "spaghetti code" — so complex and interconnected that nobody wants to touch them. The core of the problem is that traditional static thresholds just don't have enough context.

As Fred explained, these rules fail when faced with:

* **Contextual Anomalies**: Imagine a spike in server load. Is it a problem? Well, it depends. If it's Tuesday at 2 p.m., probably not. If it's Sunday at 3 a.m., that’s a different story. Static rules can't tell the difference.
* **Collective Anomalies**: This is a subtle but critical one. Sometimes, a series of events are individually fine—no single data point crosses a threshold—but together they form a problematic pattern.
* **The scale issue**: - e.g. your query returns a set of timeseries of completely different magnitudes, which you can’t craft threshold in advance (you don’t even know the magnitude), unless changing raw scale and losing interpretability (by adding some offsets, complicate queries and calculations, etc.)

This is the problem vmanomaly was built to solve. It uses ML to learn what "normal" looks like for your systems, including all their seasonal quirks.

### The Core Mechanism: It's an "ML-Powered Recording Rule"

I love this distinction. vmanomaly doesn't replace your alerting engine; it supercharges it.

Think of it this way:

1. vmanomaly reads your time series data from VictoriaMetrics.
2. It applies a machine-learning model to that data.
3. It then writes a brand new, simple metric back into VictoriaMetrics: the `anomaly_score`.

This means your complex, hard-to-maintain alerting rules can be replaced with one beautifully simple expression in [vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/): `anomaly_score > 1`. That’s it. Now you're alerting on a true deviation from the norm, not just an arbitrary number.

### Let's Get Technical: Architectural Updates

Fred walked us through some recent architectural enhancements that make vmanomaly ready for serious production workloads.

* **Stateful Mode & [Hot Reloads](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1250)**: This was a huge one. Previously, you had to restart the service to apply config changes, forcing a full model retrain. If you've ever had to wait on that, you'll love this. Now, vmanomaly can be configured to be stateful, so models persist across restarts. Plus, with hot reloads, you can tweak your YAML config on the fly and the changes are applied automatically. It makes backtesting and fine-tuning incredibly seamless.
* [**Scalability & High Availability**](https://docs.victoriametrics.com/anomaly-detection/scaling-vmanomaly/): For large environments, you can now run vmanomaly in a sharded configuration. This lets you scale horizontally across multiple instances, with each one handling a partition of the workload for performance and redundancy, all without a complex leader election process.

### Fine-Tuning is Where the Magic Happens

While the models are smart, the real power comes when you apply your own business logic. We had a great discussion about how you, the engineer, can fine-tune the output.

Your main toolkit includes parameters like:

* [**detection_direction**](https://docs.victoriametrics.com/anomaly-detection/components/models/#detection-direction): Only care if latency goes *up*? Set the direction to above. This alone cuts out a massive amount of noise.
* [**min_deviation_from_expected**](https://docs.victoriametrics.com/anomaly-detection/components/models/#minimal-deviation-from-expected): This is your noise filter. It tells the model to ignore small, insignificant deviations and only generate a score when something is *meaningfully* out of line.
* [**clip_predictions**](https://docs.victoriametrics.com/anomaly-detection/components/models/#clip-predictions): You can tell the model that a metric has a known valid range (like CPU usage being 0-100%), which keeps its predictions grounded in reality.

### Handling Missing Data (A Great Question from the Audience!)

We got a fantastic question about how to handle gaps in data — for instance, if a device goes offline. The consensus was a two-part strategy:

1. **Use vmalert for the definitive answer**: The best way to know if data is missing is with the lag() function in MetricsQL. A simple alert on this gives you a clear signal that an endpoint is down.
2. **Monitor vmanomaly itself**: vmanomaly is [self-aware](https://docs.victoriametrics.com/anomaly-detection/self-monitoring/)! If it tries to run a prediction and finds no data, it increments a missing_infer counter. You can set up a warning on this to know that your anomaly detection has a blind spot. We have also created [pre-made alerting rules](https://docs.victoriametrics.com/anomaly-detection/self-monitoring/#alerting-rules) available.

This was one of my favorite talks to host so far. It’s clear that vmanomaly is an incredibly powerful tool for adding an intelligent layer to your monitoring strategy.

To get started, I highly recommend checking out the official docs, especially the pages on the [**self-monitoring dashboard**](https://docs.victoriametrics.com/anomaly-detection/self-monitoring/) and the [**Grafana dashboard presets**](https://docs.victoriametrics.com/anomaly-detection/presets/#grafana-dashboard).

Thanks so much to Fred, Matthias, and everyone who joined us live. We'll see you at the end of August for the next one!

