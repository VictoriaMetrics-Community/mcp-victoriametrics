---
draft: false
page: blog blog_post
authors:
 - Roman Khavronenko
date: 2025-04-18
title: "Don’t default to microservices: You’ll thank us later!"
summary: "We believe microservices shouldn’t be the default and that companies should start with monoliths until reality actually demands they scale and shift from one node to a cluster of nodes. As a result, we offer clustering on an open source basis because we want to support realistic growth. Read on for details!"
enableComments: true
categories:
 - Monitoring
 - Open Source Tech
 - Observability
tags:
 - monitoring
 - microservices
 - monoliths
 - clustering
 - open source
 - victoriametrics
 - victorialogs
 - prometheus
images:
 - /blog/dont-default-to-microservices-you-will-thank-us-later/dont-default-to-microservices-you-will-thank-us-later-preview.webp
---

Donald Knuth, professor emeritus at Stanford University and "father" of algorithm analysis, once said – now quite famously – that “Premature optimization is the root of all evil.”

It’s one of those sayings that all engineers know, most understand, and many struggle to follow through on consistently. What Knuth misses in this pithy, memorable quote is the fact that evil is *tempting*. Engineers frequently fall for this temptation because, like the devil, the greatest trick premature optimization ever pulled was convincing the world it didn't exist.

The result of this trick is most visible in microservices architectures and, downstream of that, distributed systems. The biggest companies in the world – Netflix, Spotify, Amazon – use microservices, and this architecture approach enables the scalability that these companies need. But do you?

If you don’t, the costs can outweigh the benefits. Once you commit to a microservices architecture, you’re also committing to distributed systems, which require you to cluster your system’s nodes instead of sticking with one node – even if one node is really all you need.

Using a microservices approach prepares companies for a scale they haven’t achieved yet, forcing them to incur costs before they can even reap the (potential) benefits. As they grow toward that scale, complexity costs compound, leaving them inflexible and weighed down by overhead. Using microservices, especially at startups, is like laying the groundwork for a skyscraper before you even have a hut.

VictoriaMetrics has a different perspective: We believe microservices shouldn’t be the default and that companies should start with monoliths until reality actually demands they scale and shift from one node to a cluster of nodes. As a result, we offer clustering on an open source basis because we want to support realistic growth.

## Clustering should not be the default

There’s a very real and, to be fair, compelling reason companies tend to treat microservices as the go-to, default option: The best of the best are doing it.

Netflix will happily tell you about how microservices support its [TimeSeries Data Abstraction Layer](https://netflixtechblog.com/introducing-netflix-timeseries-data-abstraction-layer-31552f6326f8). Uber will explain why it needed clusters to support its [schema-agnostic log analytics platform](https://www.uber.com/blog/logging/). The list goes on.

In the vast majority of cases, companies can simply use Prometheus, but companies like Netflix and Uber have unique challenges that Prometheus just can’t support. If you look under the hood and read through the engineering journeys, even Netflix and Uber used a different solution before switching to a custom, hyper-scale one. They didn’t optimize for scale prematurely; they solved the problem when they actually faced it.

Engineers tend to default to microservices because they want to prepare for a level of scale that is only hypothetical but that they hope is inevitable. This optimism feels like pragmatism, which can lead to more incentives to cluster.

* Don’t you want to build an ideal system? You want a system that can run for the next five years, especially if the proverbial graph climbs up and to the right.
* Don’t you want credibility with company leaders? They want to grow more than anyone, and a microservices architecture gives engineering credibility because you can show them you’re ready to scale the workload 10x once the company grows 10x.
* Don’t you want to avoid worrying about scaling a monolith? Just build microservices now, so scaling – even when the company is big or growing quickly – only takes a button press.

Ten years ago, Joe Hellerstein, a professor and database systems expert, ranted to his students, [saying](https://blog.bradfieldcs.com/you-are-not-google-84912cf44afb), “There’s like five companies in the world that run jobs that big. People got kinda Google mania in the 2000s: ‘We’ll do everything the way Google does because we also run the world’s largest internet data service.’”

Google mania is still in effect, but the mania is dispersed among all the big companies and embedded in all the best practices we take as fact. And that makes it all the more difficult to notice them.

## Why you should be less afraid of using a monolith (and more afraid of microservices)

Engineers tend to overestimate the limitations of using a monolith and underestimate the complexity costs incurred by building microservices before necessary. As a result, engineers can end up building distributed systems and using database clustering long before they need to – even when a monolithic system relying on a single node would have scaled well enough.

Distributed systems are **fragile**. By nature, they have many moving parts and those parts are all in different conditions at different times. Moving parts are hard to manage and easy to break, making the entire system prone to cascading breakages.

Distributed systems are **less efficient**. When a system runs in one instance, it uses fewer resources because everything is in memory and processed by one CPU. Adding even one instance requires resources to support communication between components. Networking, encoding messages, decoding messages – all result in more CPU and memory consumption.

Distributed systems **make configuration difficult**. When you have to change the hardware or the network, a distributed system can become difficult to predict and control, which often results in failure. As Richard Cook [writes](https://how.complexsystems.fail/), “The potential for catastrophic outcomes is a hallmark of complex systems. It is impossible to eliminate the potential for such catastrophic failure; the potential for such failure is always present by the system’s own nature.” Distributed systems are more complex and are, as a result, more prone to failure.

Monolithic systems using a single node, in contrast, present advantages across every disadvantage:

* Monoliths are less complex, so they incur less complexity cost by nature.
* Monoliths provide better performance due to a lack of network overhead and not needing to encode and decode communication between components.
* Monoliths are much easier to maintain because they’re much simpler.

Of course, a different array of tradeoffs means little if monoliths can’t scale, but this is where engineers tend to underestimate them the most. Monoliths, in practice, have the capacity for surprisingly sizable workloads. You can thrive with a single node for much longer than you might think, meaning you can also benefit from avoiding the complexity costs of distributed systems for much longer than you might think.

## Why VictoriaMetrics provides clustering in open source

Of course, we’re not against microservices and database clustering – far from it. Once you can account for the tradeoffs, clustering can be really effective – so effective that we, unlike many other vendors, offer clustering as part of [our open source project](https://docs.victoriametrics.com/cluster-victoriametrics/).

As engineers, we’re often surprised when other engineers choose the clustered version of database products, even when they have small workloads that could be easily handled by a single node. 

We work hard to make VictoriaMetrics as simple as possible, and VictoriaMetrics single-node is both simple and capable. We try to make [clustering](https://docs.victoriametrics.com/cluster-victoriametrics/) simple, too, but the microservices approach is more complex by nature. Wanting the best for our users, and for users more broadly, means wanting them to choose the simple solution when it’s all they need. 

We also offer clustering as part of our open source offering because we’ve seen what happens when vendors make companies suffer from success.

InfluxDB, for example, limits access to product versions that include clustering so that companies that adopt it will eventually grow, hit a limit, and be forced to either upgrade or migrate to a new solution. 

When InfluxDB announced this change, members of the developer community had largely the same reaction, as [represented by one HackerNews user](https://news.ycombinator.com/item?id=11264157), who says, “It feels like a bait n' switch for all the people who evaluated/used InfluxDB in single-node operation as a temporary measure while giving the team ample time to work out the clustering kinks.”

Of course, InfluxDB is not alone. In March of 2024, Netdata announced [new pricing plans](https://www.netdata.cloud/blog/netdata-unified-plans/#:~:text=A%3A%20The%20Community%20Plan%20has,visualize%20at%20a%20given%20time) that repeated a similar plan. This “new unified plan structure” meant customers using the community plan would be limited to five concurrently monitored nodes. And if users didn’t pick a new plan, they would be “automatically moved to the Netdata Community Plan.”

InfluxDB and Netdata are part of a much larger trend. Many formerly and truly open source companies have introduced licensing changes (and had to backtrack somewhat following user pressure - such as Elasticsearch) that conflict with the open source spirit - just look at Redis, MongoDB, Hashicorp. etc.

Prometheus and VictoriaMetrics stand counter to this trend.

Prometheus doesn’t offer clustering, but they don’t hide this fact. It’s a deliberate design decision, not a bait and switch. We respect the decision and the transparency, and that’s one reason, among many, why we stand behind the project.

At VictoriaMetics, we build our business around supporting long-term growth, and this principle extends throughout our offering.

As a result, VictoriaMetrics provides both [single](https://docs.victoriametrics.com/single-server-victoriametrics/) and [cluster](https://docs.victoriametrics.com/cluster-victoriametrics/) versions out-of-the-box. Instead of extracting payment through must-have features, we provide as much functionality as possible on an open source basis. Instead, we offer technical support and architectural guidance as premium offerings – a strategy that aligns us closely with our customers’ interests.

There’s still work to be done to ensure this strategy is fully reflected in how our customers use and grow through our products. The migration to the clustered version requires more steps than we’d like, and we haven’t been able to simplify the process yet because it’s almost impossible to change the data layout without breaking backward compatibility.

When we announced the [general availability of VictoriaLogs](https://victoriametrics.com/blog/victoriametrics-efficiently-simplifies-log-complexity-with-victorialogs/) in November of 2024, however, we were able to put this lesson into practice. VictoriaLogs will offer seamless migration between the single node and its upcoming cluster versions as soon as you’re ready to scale.

## Cluster for actual reasons, not theoretical ones

If you look beyond the Googles, Netflixes, and Spotifys of the world, many companies are fine with single nodes.

We have customers who run just a single node and are very happy with scaling vertically. And, as indicated by the lack of support tickets, the maintenance is much easier.

Prometheus, which doesn’t offer clustering, has many happy users – all good examples of the success that can come from single-node systems. That’s why we often recommend starting with Prometheus and configuring Prometheus to push data to VictoriaMetrics once you need to cluster. Or just getting started with VictoriaMetrics Single from the outset.

It all comes down to the time interval between running a distributed system without needing it and when you can finally make use of it. The longer the time interval, the greater the cost, and there’s no assurance that the moment when you need a distributed system will ever come. A single-node system can minimize the time interval and its costs or avoid it completely.

To minimize the time interval to its shortest duration, use VictoriaMetrics. We won’t stand in the way of your growth and never will.