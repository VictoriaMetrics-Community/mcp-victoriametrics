---
draft: false
page: blog blog_post
authors:
  - Jose Gomez-Selles
date: 2025-10-07
title: "What's New in VictoriaMetrics Cloud Q3 2025? From new region in Asia to proactive alerts"
enableComments: true
featuredPost: true
summary: "Discover the latest updates in VictoriaMetrics Cloud that make observability smoother, smarter, and more secure."
categories:
  - Product News
tags:
  - victoriametrics
  - cloud
  - monitoring
  - observability
  - aws
  - roadmap
images:
  - /blog/q3-2025-whats-new-victoriametrics-cloud/preview.webp
---

The third quarter of 2025 has been a busy one for [VictoriaMetrics Cloud](https://victoriametrics.com/products/cloud/)!
We expanded globally, polished the user experience, introduced new enterprise debugging tools, and
delivered smarter alerts to help users make the most of their observability data.

If you missed our **Quarterly Live Update**, don’t worry! You can watch the full recording here:

{{< youtube n4lHqQJ0zmk >}}

Let’s recap what’s new in VictoriaMetrics Cloud this quarter.

## VictoriaMetrics Cloud Expands to Asia!

As recently announced, VictoriaMetrics Cloud is now available in **Asia-Pacific (ap-southeast-1)**.
This means faster performance and lower latency for users in the region, with the same reliability
and scalability you already know from VictoriaMetrics Cloud.

Read the full announcement [in our dedicated blog post](https://victoriametrics.com/blog/cloud-new-region-ap-southeast-1/)
or directly launch your VictoriaMetrics Cloud deployment:

<br></br>
<img src="/blog/q3-2025-whats-new-victoriametrics-cloud/regions.webp"
     style="width:65%; display:block; margin:auto;" />
<br></br>

This is important for our users that have infrastructure, applications, (or even their own users!) in Asia.
We've just removed unnecessary network round-trips and data transit overheads, helping:
- Metrics ingestion and query responsiveness
- Users with regulatory or compliance constraints over where data must be stored

If you're already a VictoriaMetrics Cloud user, you can now launch new deployments in ap-southeast-1 (Singapore)
via the console (region selector).
If you’re evaluating VictoriaMetrics Cloud, you can just [Sign Up](https://console.victoriametrics.cloud/signUp)
and test our performance during a month free of charges.

Of course, if you’d like us to expand into additional regions, [contact us](https://victoriametrics.com/contact-us/)
and let us know! Your request helps guide our roadmap.

## “Deployed on AWS”: Helping You Maximize Cloud Commitments

Since May 1st, AWS customers can only retire Private Pricing spend commitments with software
**validated as “100% built on AWS”**. VictoriaMetrics Cloud now proudly carries the
**“Deployed on AWS”** badge.

<br></br>
<img src="/blog/q3-2025-whats-new-victoriametrics-cloud/deployed-on-aws.webp"
     style="width:75%; display:block; margin:auto;" />
<br></br>

This means you can now use your AWS spend commitments on VictoriaMetrics Cloud. We totally
sympathize with the need to wisely use expenses at the end of the year, and we'll be happy to help
in this tedious task.


Read more on [AWS’s official announcement](https://aws.amazon.com/blogs/awsmarketplace/aws-marketplace-announces-upcoming-expansion-to-saas-product-catalog/).

## New Navigation Menu

We’ve rolled out a completely new **left navigation menu** designed to make everything just a click
away. Now you can easily jump between all VictoriaMetrics Cloud functionality: from monitoring your
deployments usage, change their settings, or manage access tokens.

<br></br>
<img src="/blog/q3-2025-whats-new-victoriametrics-cloud/navigation.webp"
     style="width:85%; display:block; margin:auto;" />
<br></br>

We’ve also added **breadcrumbs** so you can quickly navigate between deployments without going back
and forth.

This results in a further **step towards a more modern observability experience**, as we evolve
the _VictoriaStack_ into the Cloud, especially with **Logs** coming soon.

## Improved Deployment Logs Navigation

We know logs can be… a lot.
That’s why we’ve improved the **deployment logs view** to make it easier to read and search through.

Deployment logs are now more **responsive** and accessible, **color-coded** and **structured by columns**,
while they can still be **filtered** by date, service, or even **regex** (if that's your coup of tea!).

<br></br>
<img src="/blog/q3-2025-whats-new-victoriametrics-cloud/logs.webp"
     style="width:85%; display:block; margin:auto;" />
<br></br>

In brief: finding what you need in long lines is now much simpler (and your eyes will thank you).


## Debugging Tool for Downsampling and Retention Filters

We’ve added two powerful features from VictoriaMetrics UI directly into VictoriaMetrics Cloud:
the debugging tool for both Downsampling and Retention Filters.

These tools are now available in the Explore section, and you can now use them to easily understand
or validate how these configurations affect your metrics before applying them to production.

### But what is Downsampling?

In a nutshell, **downsampling** is the process of **reducing data resolution over time into broader
intervals**. This helps in many ways:
- Saving storage
- Improving query performance without losing visibility for longer retention times
- Reduce overall load in your deployments

For example, if you want to **keep raw data at one-second resolution for the past few days, while
storing one-minute granularity for data older than a month**, you just need to add downsampling
configuration flags for the desired time intervals and metrics.

### What about Retention Filters?

In general, we use retention configuration to **determine how long time series data is kept before it’s deleted**.
But reality tells us that not all metrics are _that relevant_ as the time passes.
With **retention filters**, you can specify different policies per metric, label, or pattern: for
example, **keeping infrastructure metrics for 30 days and business KPIs for 180 days**.
This level of granularity allows you to fine-tune your storage.

### So what's this tool about?

In theory this is great, but in reality it can be scary. That's where this Debugging Tool shines: it gives
you an **interactive way to simulate how downsampling and retention rules work**. By playing around,
you can **check how these rules and filters will impact your metrics without changing your live setup**.
It works in 3 steps:
1. Enter a list of flags
2. Select one or more metrics to which the configuration should apply
3. And click “Apply” to generate a detailed preview

<br></br>
<img src="/blog/q3-2025-whats-new-victoriametrics-cloud/retention-debug.webp"
     style="width:85%; display:block; margin:auto;" />
<br></br>

The result displays a comprehensive list of how each metric is affected.

If this is still a bit overwhelming for you, don't forget that you can just click on the **“Try Example”**
button to load sample configurations and metrics to see how the system behaves, as you learn without risks.

>[!TIP] You can find more info about [retention filters](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#retention-filters) and [downsampling](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#downsampling) in the docs.


## Alerts and Limits: Better together

VictoriaMetrics Cloud automatically manages your deployment, but **staying within your chosen tier** is
important to maintain stability and avoid throttling.

Now, you’ll receive **notifications** when you’re approaching or exceeding:
- Ingestion rate
- Active time series
- New series over 24h (churn rate)
- Datapoint read speed
- Series read per query

These alerts include **actionable recommendations**, helping you decide whether to scale up or optimize your workload.

> [!INFO] Remember that you can always [contact support](https://docs.victoriametrics.com/victoriametrics-cloud/support/) if you need help or guidance.

## Other Improvements

Apart from all of these improvements, the new VictoriaMetrics Cloud version includes:
- **Prometheus 3.0 MetricsQL support** – including quoted label names
- **New styles for compact screen resolutions**
- **API for VM Single:** enabled `log_new_series` to track new series and identify churn sources
- **Regular security updates** including **TLS 1.0 and 1.1 deprecation** for improved security, as [previously announced](https://victoriametrics.com/blog/q2-2025-whats-new-victoriametrics-cloud/#important-update-tls-versions-support-change).

<br></br>

---

## What’s Next?

We’re not slowing down! Here’s what’s coming soon:

- **New capacity tiers**
  Stay tuned, because we are revamping our tiers. Based on real-world usage, new benchmarks and feedback, we’re providing a new list of VictoriaMetrics Cloud tiers to ensure smoother scalability and clearer progression between deployments.

- **Logs are coming to VictoriaMetrics Cloud!**
  As we advance in our path to deliver VictoriaLogs in VictoriaMetrics Cloud, we can announce that the integration with initial support for **OpenTelemetry**, **Vector**, and **FluentBit** is done (and **Perses** warming up on the bench).
  The new UI already includes selectors and integrations by type, and early PoCs are looking great. **I want to personally take this opportunity to thank our great community of users and customers for their [feedback](https://forms.gle/aXMTEijVQgiEVPqu5)** (still open!).

As always, thanks for spending this time with us and for being part of the VictoriaMetrics community ❤️.
Your feedback keeps shaping the future of VictoriaMetrics Cloud.

Ready to try the latest features?
[Sign up for VictoriaMetrics Cloud](https://console.victoriametrics.cloud/signUp) and get started today!
