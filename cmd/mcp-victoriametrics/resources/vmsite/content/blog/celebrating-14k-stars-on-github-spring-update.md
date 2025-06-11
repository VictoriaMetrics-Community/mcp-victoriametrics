---
draft: false
page: blog blog_post
authors:
 - Denys Holius
date: 2025-05-21
enableComments: true
title: "Celebrating 14K Stars on GitHub: Spring Update"
summary: "Seeing that VictoriaMetrics products are this popular with engineers worldwide is fantastic: Just a little over a year ago, we hit 10K stars, and with the adoption of VictoriaLogs, the star count now went beyond 14K. Read about most recent achievements in this blog post. "
categories: 
 - Company News
tags:
 - github
 - stars
 - victoriametrics
 - victorialogs
 - achievements
images:
 - /blog/celebrating-14k-stars-on-github-spring-update/preview.webp
---

![A 4K star's growth in one year!](/blog/celebrating-14k-stars-on-github-spring-update/vm-starhistory-4k-in-one-year.webp)

Seeing that VictoriaMetrics products are this popular with engineers worldwide is fantastic: Just a little over a year ago, we hit 10K stars, and with the adoption of VictoriaLogs, the star count now went beyond 14K. We don’t take these GitHub Stars milestones for granted: It’s amazing to see these stats grow organically thanks to the community of users out there who use our products. Thank you so much!

## What other interesting things happened this spring?

### Cool new features
* Want to understand your data complexity aligned with the reading pattern? VictoriaMetrics [Cardinality Explorer](https://docs.victoriametrics.com/#cardinality-explorer) is a useful tool for understanding stored data and its impact on resource usage. With the latest updates, we combine the functionality of [tracking metrics names](https://docs.victoriametrics.com/#track-ingested-metrics-usage) (aka unused metrics) with your cardinality. Now, you can see the metrics with the higher number of time series along with the number of times they are a part of the queries.
* VictoriaMetrics and VictoriaLogs have become more friendly for newbies since we released the two MCP servers: [mcp-victoriametrics](https://github.com/VictoriaMetrics-Community/mcp-victoriametrics) and [mcp-victorialogs](https://github.com/VictoriaMetrics-Community/mcp-victorialogs). This provides a comprehensive interface for managing logs, observability, and debugging tasks in your VictoriaLogs instances while enabling advanced automation and interaction capabilities for engineers and tools, or those who want to explore them deeply.
* Check out our new features in VictoriaMetrics Cloud: A revamped Organizations feature for better collaboration, seamless OpenTelemetry integration, a powerful new Explore tab, as well as additional API endpoints to enhance automation and control, together with many improvements based on your invaluable feedback. [Read this post for details](https://victoriametrics.com/blog/q1-2025-whats-new-victoriametrics-cloud/).

### Our team was on the road & online meeting users
* VictoriaMetrics sponsored SCaLE22x, which is the largest community-run open-source and free software conference in North America. We enjoyed talking to the community about open source, monitoring, observability, golang, ... our favorite topics! No one was left without stickers 😎

  ![VictoriaMetrics at SCaLE22x](/blog/celebrating-14k-stars-on-github-spring-update/SCaLE22x.webp)

* We held the first VictoriaMetrics Online Community Meet Up of the year where we talked about new and cool features implemented in our products for the first quarter.
The recording is available on our [YouTube channel](https://www.youtube.com/@victoriametrics/videos), so don't miss the opportunity to watch an exciting demo by Natan Yellin: [Using AI + VictoriaMetrics to Answer Developer Questions](https://www.youtube.com/watch?v=33z8e6ZEeWk&list=PLXT8DSiuv5yl47-tv4Tl1Uty6nLNX6ar1&index=14)!

  ![VictoriaMetrics Virtual Meet Up March 2025](/blog/celebrating-14k-stars-on-github-spring-update/victoriametrics-virtual-meet-up-march-2025.webp) 

* Our products have a lot of interesting features that not everyone knows about, so this year [we have launched and held four VictoriaMetrics Tech Talks series](https://www.youtube.com/watch?v=deDo_keTxjs&list=PLXT8DSiuv5ymWbr02i0rqcGimDppvdre5&index=4) - online sessions, where you can see how VictoriaMetrics works for blackbox monitoring, explore how VictoriaLogs makes log management effortless, and more.
* It was a pleasure meeting everyone at KubeCon 2025, hearing your feedback, and discussing open source observability! We were thrilled to share the latest innovations in VictoriaMetrics and VictoriaLogs, and to reconnect with the community—especially following our time at KubeCon Salt Lake City 2024. KubeCon continues to be a fantastic opportunity to engage with the community about our open source solutions and explore how we can support your monitoring and observability goals. [Watch the highlights](https://youtu.be/0pZ1kEIupI4) and see you at the next KubeCone NA this fall!

  ![VictoriaMetrics Team at the KubeCon 2025](/blog/celebrating-14k-stars-on-github-spring-update/victoriametrics-team-at-the-kubecon-2025.webp)

## Looking Ahead
There are several interesting events lined up for the near future.
1. Don't miss the next #5 Tech Talks: “[Mastering vmalert: Best Practices for Effective Alerting](https://www.youtube.com/watch?v=zpjBSZ8TkGU)”, which will be held this week.
2. For the first time, we'll be holding “[Features & Community Call](https://www.youtube.com/watch?v=yfNa9cvUAVQ)”. We'll talk about new features and improvements as well as answer interesting and frequently asked questions in our communities.
  ![Feature and Community Call](/blog/celebrating-14k-stars-on-github-spring-update/feature-and-community-call.webp)
  
3. Our traditional [VictoriaMetrics' Virtual Meet Up will be held on June 19th](https://youtube.com/live/Y8OG1JnEKA0), where we will discuss products and roadmap updates. Don't miss this opportunity to hear Eric Deleforterie's story of adopting VictoriaMetrics products in his company and see a demo from Alexander Marshalov about using MCP server for VictoriaMetrics!
  ![Feature and Community Call](/blog/celebrating-14k-stars-on-github-spring-update/q2-virtual-meetup-june-2025.webp)

We're excited for the next chapters of our journey—and, of course, the next 10,000 stars on GitHub!
