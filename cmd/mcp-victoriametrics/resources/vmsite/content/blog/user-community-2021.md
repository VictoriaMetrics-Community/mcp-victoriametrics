---
draft: false
page: blog blog_post
authors:
 - Roman Khavronenko
date: 2022-01-07
title: "Vicky User Community 2021: Thank You for the Contributions!"
enableComments: true
summary: 'In this post, we want to say "Thank you!" to all the people who helped VictoriaMetrics  become what it is today and that we appreciate their contributions via this overview of the most interesting user blog posts of the year and a shortlist of top community contributors.'
categories: 
 - Company News
 - Community
tags:
 - victoriametrics 
 - open source 
 - community 
 - 2021 
 - monitoring solutions 
 - user contributions
---
2021 was a great year for VictoriaMetrics! We delivered a lot of [new features](/blog/features-roundup-2021/), our team doubled in size, and so did the list of [public case studies](https://docs.victoriametrics.com/CaseStudies.html#case-studies-and-talks) written by VictoriaMetrics users as well as the community contributions to the product. See our [2021 Momentum blog post](/blog/momentum-2021/) for details on all our achievements last year.

All this wouldn't be possible without our supportive community, their help, patience and creativity.

In this post, we want to say "Thank you!" to all the people who helped VictoriaMetrics  become what it is today and that we appreciate their contributions via this overview of the most interesting user blog posts of the year and a shortlist of top community contributors.
<p>&nbsp;</p>

# Top Community Blogs

We'll start with the post by Perconian [Steve Hoffman](https://www.percona.com/blog/author/steve-hoffman/) in December 2020 [Foiled by the Firewall: A Tale of Transition From Prometheus to VictoriaMetrics](https://www.percona.com/blog/2020/12/01/foiled-by-the-firewall-a-tale-of-transition-from-prometheus-to-victoriametrics/). It is not just a blog post, it is a story filled with emotions and plot twists which I personally very enjoyed reading. If you ever wondered why someone would prefer the Push model to the Pull model, how to run monitoring in resource-constrained environments and how complicated networking could be - this article will satisfy your curiosity.

The next blog post is [Multi-tenancy monitoring system for Kubernetes cluster using VictoriaMetrics and operators](https://blog.kintone.io/entry/2021/03/31/175256) by [UMEZAWA Takeshi](https://github.com/umezawatakeshi) in March 2021. The article covers one of the most popular problems in monitoring: high availability, multi-tenancy and long-term storing. The author goes through the details of running highly available monitoring solutions in Kubernetes, his experience with [VictoriaMetrics Kubernetes operator](https://github.com/VictoriaMetrics/operator) and storage durability via [Ceph](https://docs.ceph.com/) and [TopoLVM](https://github.com/topolvm/topolvm).

In May 2021 [Thomas Ptacek](https://twitter.com/tqbf) published a [Fly's Prometheus Metrics](https://fly.io/blog/measuring-fly/) blog post about fly-proxy, [Borgmon](https://research.google/pubs/pub43438/), [Prometheus](https://prometheus.io/), [Thanos](https://thanos.io/) and, of course, Vicky (that's how they call VictoriaMetrics :-)). This is a serious, lengthy technical post with a lot of details and architectural decisions, full of references to industry standards and solutions. Nevertheless, I read it in one breath and would  recommend it!

[Scaling to trillions of metric data points](https://engineering.razorpay.com/scaling-to-trillions-of-metric-data-points-f569a5b654f2) - you just can't get past of such a title! The blog post by [Razorpay](https://razorpay.com/) engineers [Venkat Vaidhyanathan](https://medium.com/u/e9b3bbbc82dd) and Vaibhav Khurana in July 2021. It contains a lot of things: how [Prometheus](https://prometheus.io/) replaced [Nagios](https://www.nagios.org/) and [Icinga](https://icinga.com/) for [Kubernetes](https://kubernetes.io/) and applications monitoring; scalability issues with [Prometheus](https://prometheus.io/); experience of running [Thanos](https://thanos.io/) in production; comparison of [Cortex](https://grafana.com/oss/cortex/) and VictoriaMetrics architectures; the cost benefits of switching to VictoriaMetrics in the end. The article is well-designed, illustrated, full of technical details, architectural comparisons between projects based on personal experience. A big "Thank you!" to the authors for their hard work!

The "high cardinality" problem isn't something new to people who are familiar with modern monitoring. This September blog post [Choosing a Time Series Database for High Cardinality Aggregations](https://abiosgaming.com/press/high-cardinality-aggregations/) by [Patrik Karlström](https://abiosgaming.com/press/author/patrikabios-se/) tells a story about how an increase in metrics cardinality pushed them to look for an alternative to Prometheus. It contains a performance comparison between VictoriaMetrics and TimescaleDB based not on synthetic benchmarks, but on real production scenarios. The post itself has a great structure and is very well-written - definitely a must-read!
<p>&nbsp;</p>

# Top Community Contributors

In no particular order :-)

- [Faceair](https://github.com/faceair), for various performance optimizations
- [Belm0](https://github.com/Belm0), for being thorough and thoughtful in MetricsQL details
- [Patsevanton](https://github.com/Patsevanton), for maintaining RPM package for VictoriaMetrics
- [Johnseekins](https://github.com/Johnseekins), for OpenTSDB support in vmctl and being an active community member

We also want to thank all of our 80+ contributors!

We once again want to say ‘thank you’ to all our community members for being honest and supportive!

We hope that 2022 becomes even better for us all!
