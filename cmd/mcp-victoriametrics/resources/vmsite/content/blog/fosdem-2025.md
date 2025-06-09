---
draft: false
page: blog blog_post
authors:
  - Jose Gomez-Selles
  - Roman Khavronenko
date: 2025-02-14
title: "FOSDEM 2025 recap"
summary: "FOSDEM (Free and Open Source Software Developers' European Meeting) is a huge, free, gathering for open-source software enthusiasts that happens every February in Brussels, Belgium. It's a non-profit event put together by the community, and it's one of the biggest of its kind. See more in our recap!"
enableComments: true
categories:
  - Open Source Tech
  - Monitoring
tags:
  - open source
  - monitoring
  - fosdem
images:
  - /blog/fosdem-2025/preview.webp
---

## VictoriaMetrics FOSDEM 2025 recap 

In case you haven't heard about it yet, [FOSDEM](https://fosdem.org/) (Free and Open Source Software Developers' European Meeting)
is a huge, free, gathering for open-source software enthusiasts that happens every February in Brussels, Belgium.
It's a non-profit event put together by the community, and it's one of the biggest of its kind - we're talking about 
around 10,000 people from all over the world coming to hang out and talk about all things open source.

This means that this conference is a big deal for open source! It's where developers, users, and fans can get together 
and talk about their work and ideas. Open source projects can show off what they're doing and find new people to help out.

For VictoriaMetrics, FOSDEM is an especially important event because it is a chance to connect with communities 
like OpenTelemetry and Prometheus, as well as other open source projects that VictoriaMetrics integrates with. 
It is also an opportunity to learn about the latest trends in open source monitoring and observability.

{{<image href="/blog/fosdem-2025/fosdem-campus-morning.webp" alt="Entrance to the Université Libre de Bruxelles, where FOSDEM takes place" >}}

## FOSDEM 2025 - the anniversary edition

This year's FOSDEM was particularly special because it was the 25th anniversary of the event. The conference was 
more crowded than ever, with many attendees having to queue to get into the talks they wanted to attend.
This is a testament to the vibrancy and enthusiasm of the open source community.

FOSDEM hosts so many devrooms, it's like a city within a city. People are walking everywhere, meeting old friends 
and colleagues, attending community parties and events. There's so much to see and do that one could write a book about it all. 
Have you seen the [number of tracks](https://fosdem.org/2025/schedule/)!?

In this article, we'll focus on what the VictoriaMetrics team experienced at the Monitoring and Observability dev room.

## The Monitoring and Observability DevRoom

Observability seems to continue gaining traction these days. At least it looks like it raises as much interest as the 
[Go programming language](https://go.dev/), since the venue for the 
[Monitoring and Observability](https://fosdem.org/2025/schedule/track/monitoring/) track this year was the same that
received all popular [Go talks](https://fosdem.org/2025/schedule/track/go/) the day before.

{{<image href="/blog/fosdem-2025/the-state-of-go.webp" alt="Opening of Golang DevRoom. “The state of Go” by Maartje Eyskens" >}}

The Sunday started with a great opening by [Richard "RichiH" Hartmann](https://fosdem.org/2025/schedule/speaker/richard_richih_hartmann/)
who set the ground for the day as the room was still being populated by those who didn't get lost in Brussels' Saturday night.

After a small break to let more people in, [Jose Gomez-Selles](https://fosdem.org/2025/schedule/speaker/jose_gomez-selles/),
Product Lead for Cloud at VictoriaMetrics, started his talk ["Discovering the Magic Behind OpenTelemetry Instrumentation"](https://fosdem.org/2025/schedule/event/fosdem-2025-4146-discovering-the-magic-behind-opentelemetry-instrumentation/).
Here, Jose tried to demystify how the internals of OpenTelemetry instrumentation work, making it easier for end users 
to produce high quality data that is both useful and avoids waste. Despite the automatic (or [zero-code](https://opentelemetry.io/docs/zero-code/))
instrumentation capabilities that OpenTelemetry brings to the table are super appealing in terms of the simplicity 
of activation and operations, it comes with some caveats (that were very well explained later by 
[James Belchamber](https://fosdem.org/2025/schedule/speaker/james_belchamber/) in his talk: [The performance impact of auto-instrumentation](https://fosdem.org/2025/schedule/event/fosdem-2025-5502-the-performance-impact-of-auto-instrumentation/)).

That's why, as Jose explains, we always need to take care of our code and provide useful instrumentation. 
And, for that, we need to spend time and learn how to do it.

{{<image href="/blog/fosdem-2025/jose-intro.webp">}}

By explaining how OpenTelemetry defines a specification that is common to every language, we were able to get a better
understanding of the components that we need in our code, regardless of our favorite stack. A brief demo with a testing
client and a dummy server was also enough to demonstrate how a simple stack based on VictoriaMetrics and Jaeger can be
enough to monitor and observe distributed applications.

{{<image href="/blog/fosdem-2025/jose-slide-1.webp">}}

Don't hesitate to take a look at the [slides](https://fosdem.org/2025/events/attachments/fosdem-2025-4146-discovering-the-magic-behind-opentelemetry-instrumentation/slides/237238/FOSDEM_20_mlr45ST.pdf) 
or [recording](https://video.fosdem.org/2025/ud2120/fosdem-2025-4146-discovering-the-magic-behind-opentelemetry-instrumentation.av1.webm)
if you are struggling with instrumenting your application!

After this talk, the day advanced with other great sessions full of insightful takeaways: from the mentioned performance
impact of auto-instrumentation, to learning what's new in [Prometheus Version 3](https://fosdem.org/2025/schedule/event/fosdem-2025-6571-prometheus-version-3/),
understanding how [Lorenzo Nicora](https://fosdem.org/2025/schedule/speaker/lorenzo_nicora/) and [Hong Teoh](https://fosdem.org/2025/schedule/speaker/hong_teoh/)
scaled their Observability set up with [Apache Flink](https://fosdem.org/2025/schedule/event/fosdem-2025-5726-apache-flink-and-prometheus-better-together-to-improve-the-efficiency-of-your-observability-platform-at-scale/),
just to name a few. The day was packed with great experiences and information!

It was nearly at the end of the day, when [Roman Khavronenko](https://fosdem.org/2025/schedule/speaker/roman_khavronenko/),
Co-Founder at VictoriaMetrics and contributor to many Open Source projects such as ClickHouse and Grafana, took the stage
to talk about a highly important topic, its relevance we sometimes only understand in the later stages of setting up our 
production environments: ["How to Monitor the Monitoring"](https://fosdem.org/2025/schedule/event/fosdem-2025-5388-how-to-monitor-the-monitoring/).

{{<image href="/blog/fosdem-2025/roman-slide-1.webp">}}

In this talk, Roman explained how the VictoriaMetrics Open Source project spends time and efforts to help users understand
how the time series database is performing in production, so they can spend more time on their own code instead of trying
to understand the internals of a component designed to observe their own systems.

By deep-diving into the use of features in VictoriaMetrics combined with integrations with other projects like Grafana, 
the audience could learn which signals and information are relevant for the end user, but also, which strategies can be 
applied, as an Open Source project maintainer to ease operability in production.

{{<image href="/blog/fosdem-2025/roman-slide-2.webp">}}

We navigated through monitoring-related questions, shedding light on effective strategies to improve our monitoring 
practices. All in all, by sharing past experiences in many support cases encountered with VictoriaMetrics, we learnt
what's important in Monitoring of Monitoring via engineer-friendly Grafana dashboard creation, alert optimization, 
and troubleshooting guide compilation.

After Roman's talk, we still had some energy left to learn about [Effortless, standardised homelab observability with eBPF](https://fosdem.org/2025/schedule/event/fosdem-2025-4680-effortless-standardised-homelab-observability-with-ebpf/)
by [Goutham Veeramachaneni](https://fosdem.org/2025/schedule/speaker/goutham_veeramachaneni/).

## Conclusion

It's amazing the number of talks and insights that can be taken from this event. Simply summarizing them wouldn't do FOSDEM justice.
You just need to live it! This time we want to focus more on sharing our experience.

At a personal level, it was a great opportunity for old friends in the VictoriaMetrics team to meet up, and also meet 
some team members in person for the first time.

On a more Community and Business oriented note, it was also awesome to have discussions with both Open Source users and
customers, discussing their use cases and learn about the many ways they use VictoriaMetrics in production combined with
other Open Source projects.

About the future, we confirmed the importance of integrating in this vast and beautiful ecosystem of projects, and how 
the enterprise world helps to boost it to the next level. There's never a one-size fits all solution, and the diversity 
in the landscape is what brings value. Open Source solutions are mixed with managed and enterprise solutions, and our team
is working on helping make that happen on many fronts.

{{<image href="/blog/fosdem-2025/victoria-metrics-team.webp" alt="VictoriaMetrics representatives at FOSDEM. From left to right: Alexander Marshalov, Jose Gomez-Selles, Dima Kozlov, Aliaksandr Valialkin and Roman Khavronenko" >}}

Thanks, FOSDEM 2025!