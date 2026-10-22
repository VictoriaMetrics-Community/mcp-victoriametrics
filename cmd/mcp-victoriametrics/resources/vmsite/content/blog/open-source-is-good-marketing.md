---
draft: false
page: blog blog_post
authors:
  - Jean-Jerome Schmidt-Soisson
date: 2026-08-24
enableComments: true
title: "Open Source Is Good Marketing"
summary: "Open source and commercial marketing often seem at odds, but open source often needs marketing, and open source itself is often good marketing. This blog post explains how open source supports our marketing goals and how we use marketing to support open source growth."
description: "Open source can be great marketing: VictoriaMetrics explains how community, content, and discoverability drive growth without compromising OSS values."
toc: true
categories:
  - Company News
tags:
  - victoriametrics
  - open source
  - marketing
images:
  - /blog/open-source-is-good-marketing/preview-4.webp
---

> **Summary**: Jean-Jérôme Schmidt-Soisson, marketing lead at VictoriaMetrics, shows that open source and marketing are not opposites. Open source itself can be a powerful growth engine through trust, community, content, documentation, and discoverability. The article explains VictoriaMetrics’ pull-based marketing strategy, anchored in open-source values, to support adoption and business growth.

This is our second article on the relationship between open source and marketing. The [first article](https://victoriametrics.com/blog/creating-a-sustainable-open-source-business-model/) in this series talks about sustainable open-source business models. Here, I lay out our strategy to turn VictoriaMetrics into a recognized standard in the open-source observability space.

There has always been tension between commercial and non-commercial in the open source space. There is the notion that open source should stay "pure". That it shouldn't have profit or growth goals, or any commercial interests at all. Marketing often draws the line: true open-source projects don't do marketing at all.

This is false. After all, open source is good marketing in itself. Think about it: the primary benefits of open source are trust, community contribution, and distribution. Any good marketer would appreciate any of those.

Many people don't connect marketing with community and clear information sharing. This makes marketing feel disconnected from open-source values. A developer might ask: How can you really have open-source values if you're pushing your project via ads or targeting open-source users with enterprise sales pitches?

Done right, open source can be an engine for growth for the community and business. Done wrong, it can be everything open-source purists worry about. So, how can we do it right?

## Unbundling "marketing"

Marketing is everywhere. It's on television, in podcasts, on billboards, and all over public transit. Sometimes, it's obnoxious, like a pop-up interrupting your browsing. Often, it's insidious: a "free" e-book that triggers dozens of emails and phone calls from a vendor's sales team.

In 2025, views on the advertising and PR industry [hit new lows](https://www.marketingcharts.com/industries/government-and-politics-236989). Only 29% of US adults view the industry positively, while 37% view it negatively.

![Line graph showing net positive rating of advertising industry: 2022 marked a steep decline in perception with all-time lows in 2024 at -17% rating](/blog/open-source-is-good-marketing/perception-of-marketing.webp)

A 2025 Survey by We Are Social found that [32.7% of Americans use an ad-blocker](https://datareportal.com/reports/digital-2025-global-overview-report). It's no surprise, then, that many engineers, who are even more privacy-minded, have developed a particular aversion to marketing. They need to stay up to date on new tools and technologies, and often have some level of purchasing power. So, they're more likely to feel pushed than others.

![Slide of report shows: 71.6% of Americans express concerns about what is real vs. fake, 43.6% worry about how companies might use their data, 48.2% decline cookies at least some of the time, 32.7% of Americans use a tool to block ads at least some of the time, 24.6% use a VPN at least some of the time](/blog/open-source-is-good-marketing/ad-block-usage.webp)
<figcaption style="text-align: center; font-style: italic;">Slide 152 of the <a href="https://indd.adobe.com/view/c5c21264-0430-49b8-b0ad-07f4349b3fda?allowFullscreen=true">We Are Social report</a></figcaption>

Part of the issue is that "marketing" covers a lot. Engineers, who may not be that interested in marketing, might even say straight up that they don’t like it. But they often enjoy tutorials, events, swag, podcasts, essays, and more. But what are these if not yet another form of marketing?

When we say open source can be good marketing, we don't mean turning projects into billboards. 

We're saying that, if you unbundle all the activities that comprise marketing, open source can support many of the growth channels we typically associate with marketing. 

### What's the difference between push and pull marketing?

The primary divide in this bundle is between "push" marketing and "pull" marketing. 

Push marketing is what most people associate with overly commercial marketing, including:

- Billboards, commercials, and ads (whether offline or online).	
- Gated content that leads to contact from salespeople.
- Demos are only available by scheduling a meeting with the vendor representative.

Pull marketing often doesn't feel like marketing because it isn't being pushed at you. Instead, you’re pulling it toward you because the value of the material doesn't need to be pushed. This includes:

- Community spaces (virtual, such as Discord, or in person, such as meetups).
- Content (announcements, benchmarks, technical articles, tutorials, and guides).
- High-quality documentation.
- Branding (e.g., branded templates for user group presentations, etc.).

In my experience, pull marketing is often compatible with open source while push marketing rarely is. At VictoriaMetrics, we felt this early on because going open source was the spark that started growth. In the [original announcement post](https://valyala.medium.com/open-sourcing-victoriametrics-f31e34485c2b), Aliaksandr Valialkin, CTO at VictoriaMetrics, wrote: "We believe that open-sourcing will help VictoriaMetrics gain traction, so it will become a prominent player in the world of time series databases".

And so it did:

- We hit 1B+ downloads at the end of 2025.
- The [VictoriaMetrics repository](https://github.com/VictoriaMetrics/VictoriaMetrics/) has over 17k stars on GitHub.
- VictoriaLogs hit [general availability](https://victoriametrics.com/blog/victoriametrics-efficiently-simplifies-log-complexity-with-victorialogs/).
- [VictoriaTraces](/products/victoriatraces/) was added to the stack, completing the three signals: metrics, logs, and traces.

![Star history chart for VictoriaMetrics/VictoriaMetrics repository.](/blog/open-source-is-good-marketing/star-history.webp)
<figcaption style="text-align: center; font-style: italic;">Star count as of 2026/07: 17373</figcaption>

The early choice to go open source was the best marketing decision we've ever made. Our long-term goal is to make VictoriaMetrics a recognized standard for open source observability. Building a project that's good enough to be a regarded standard in the space is good for the community, and becoming one is great marketing. The incentives align.

## **The open source funnel: Product-led growth ahead of its time?**

Open source puts the product at every stage of the funnel, not just the bottom; that's the difference from traditional marketing.

The marketing funnel, at its simplest, is a three-step process that maps how a company builds an audience and converts part of them into buyers.

![Marketing funnel diagram. From top to bottom: awareness -> consideration -> purchase](/blog/open-source-is-good-marketing/marketing-funnel.webp)

Push marketing blankets websites with ads, betting that a few will click and lead to a sales follow-up. Pull marketing publishes high-value blog posts, tutorials, demos, and documentation that entice users to try the product on their own terms.

If open source is good marketing, how does it fit into the funnel? 

Traditionally, the product comes into play only at the bottom of the funnel, after prospects convert into customers. With open source, the product is featured throughout the funnel. It is the core driver of growth and acquisition at every stage.

This might sound familiar if you've heard of product-led growth (PLG). [Coined in 2016](https://openviewpartners.com/blog/inventing-product-led-growth/), PLG emphasizes self-service models as the primary drivers of customer growth. It is why so many companies have freemium models, such as Slack and Calendly.

PLG wasn't the first to drive growth through the product itself. Open source projects have used this model for a long time. The link between product-led growth and open source values is even stronger than in PLG.

When a PLG company wants you to try its product, it often offers a freemium version with limited functionality. In contrast, open-source projects give you the product without limitations or strings attached (though that can also depend on the open-source license in play). 

At VictoriaMetrics, we take open source seriously and run under the Apache-2.0 license without gates. We don't track who downloads and uses the software, and we rely solely on aggregated GitHub, Docker, and similar download metrics, which provide only general adoption data.

When a PLG company wants to grow, it waits for purchasing signals and sends in the sales team. When open source projects want to grow, we wait. We wait for users to contact us when, and only when, they need enterprise features, hosting, or support.

The benefit for us is that when users reach out, they are very motivated. They've already tried the software. At VictoriaMetrics, for example, the majority of our paid users were running the open-source version before contacting us. That's another reason our business model is as transparent as the open-source code itself.

Open source is important in the marketing funnel. It raises awareness by showcasing our projects. It also helps users learn before they buy, as they can try it as often as they want.

The funnel looks simple, but it is hard to execute in practice. Product quality in an open-source project matters far more than in commercial software.

This is how open source fits into the marketing funnel: It increases awareness by putting our product out there and educates our users before consideration and purchase by giving them the chance to try it as much as they want. 

The funnel looks simple, but it is hard to execute in practice. Product quality in an open-source project matters much more than it does in commercial software. 

If the product isn't good, there's no hiding it. Consider that:

- There's an immediate and public feedback cycle; everyone can see the bugs and how you deal with them.
- Users can easily test before committing. You can't hide flaws until the contract is signed.
- Users can switch whenever they want. If the product isn't good, users can pack up and move on to the next project.

Building a great product is a fundamental precondition for success. But fulfilling this precondition isn't enough. Once you understand the funnel, you can fold marketing in without sacrificing what makes open source special.

## Amplifying the marketing benefits of open source

Many people want to believe that "If you build it, they will come," but they won't if nobody knows it exists in the first place. That's the reason marketing exists.

Open source can be great marketing on its own. But how can marketing help us without losing the open-source spirit? We break open-source marketing down into three core strategies:

- Encouraging word of mouth
- Developing great content
- Maintaining discoverability

### Encouraging word of mouth

Open source is the ideal engine for word-of-mouth distribution. If you like our project, you can recommend it without feeling like you're selling something. You can also try it immediately without demos, sales conversations, credit card sign-ups, or newsletter sign-ups.

To encourage this, we host community spaces on [the VictoriaMetrics Slack](https://slack.victoriametrics.com/) and [Telegram](https://t.me/VictoriaMetrics_en) so users can interact with each other. Bringing users with similar challenges and questions together helps them share solutions. We also encourage users to contribute (via pull requests, comments, and issues) and [share their experiences working with the project](https://victoriametrics.com/blog/observability-lessons-from-openai/) (via blog posts, social media, etc.).

We also like to show ourselves where developers gather. The product may speak for itself, but we have to be visible and show up at events and conferences to get them into the conversation.

### Developing great content

We rely almost exclusively on inbound organic growth (also known as "pull" marketing). For us, that means creating content so valuable that users and customers come to us rather than us pushing content to them.

At VictoriaMetrics, the two primary content channels are blog posts and documentation.

Blog posts include:

- Announcements
- Benchmarks
- Technical articles
- Tutorials
- Guides

In each content type, we can contribute to the surrounding ecosystem. For example, in a post about [alerting best practices](https://victoriametrics.com/blog/alerting-best-practices/), we can teach people how to design alerts more effectively and help readers realize we know what we're talking about. By [benchmarking OpenTelemetry, Prometheus, and vmagent](https://victoriametrics.com/blog/opentelemetry-prometheus-and-more/), we gather valuable and objective data. This gives us visibility into more recognized projects.

If the blog sits near the top of the funnel, the docs sit closer to the bottom. In [our development goals](https://docs.victoriametrics.com/victoriametrics/goals/), docs are explicitly called out:

- On the one hand, "Improving docs for VictoriaMetrics components, so users can find answers to their questions using Google or their favorite AI platforms without the need to ask these questions at our support channels" is a critical goal.
- On the other hand, "Breaking links to VictoriaMetrics docs, since users will be unhappy seeing a 404-page or unexpected results after they click some old link somewhere on the Internet or in the internal knowledge base" is a critical non-goal.

High-quality, well-maintained docs are often more important to us than new features. Docs help users find and use products and features. This makes them happy and encourages them to share our product with others.

### Maintaining discoverability

Great content and strong word of mouth don't mean much if users can't find us. A key aspect of open-source marketing is ensuring visibility. This means maintaining an active, visible presence on platforms such as GitHub, Docker, Quay.io, and other relevant repositories. When someone searches for us in the usual places, we want to ensure that we appear at the top of the results.

Similarly, we ensure our docs and other content are optimized for SEO and LLMs. If you search for a concept we've written about, we want Google and other search engines (as well as ChatGPT and other AI products) to know we've covered that topic.

Across online and offline channels, we also try to maintain a consistent brand voice, tone, and visual style to ensure people know they're seeing VictoriaMetrics when they see us.

## Traditional marketing as a complement

All this isn't to say we never do traditional "push" marketing. We do.

Our core marketing goal is to build and maintain credibility. Credibility, however, is earned; you can't buy or force it. Paid advertising won't make us credible.

That's why these traditional marketing methods support only our commercial products and never our open-source projects. The line between them is clear. So far, we have limited ourselves to:

- Run outbound advertising campaigns for our enterprise products.
- Create LinkedIn ads targeting specific customer segments.
- Operate a sales team that manages enterprise contracts.
- Keep the traditional marketing lean. Open source is the foundation; inbound marketing amplifies open source; and traditional marketing fills in the gaps.

With the rise of AI search and LLMs, other traditional marketing functions are also on the rise: Events, conferences, press & analyst relations … This year we’re increasing our presence at conferences and events by 100%; and we'll be working with press and analyst relations experts again. All with the view to help amplify our organic messaging and presence. 

## Open source first and always

Open source is always at the center of everything we do and runs through all our communication.

We have an advantage here that most companies don't: we are self-funded. That means we have to (and get to) think long term. We're not under pressure to burn marketing dollars and convert as many customers as possible. We can focus on our community first and grow our customer base by attracting contributors and users genuinely interested in the product.

We need to make every marketing dollar count and spend it wisely: On supporting and amplifying open source.

If you want to see this in practice:

- All the code is on [GitHub](https://github.com/VictoriaMetrics)
- Documentation is on [docs.victoriametrics.com](https://docs.victoriametrics.com)
- You can try VictoriaMetrics online in [playgrounds](https://docs.victoriametrics.com/playgrounds/)
- Join the community on [Slack](https://slack.victoriametrics.com/) and [Telegram](https://t.me/VictoriaMetrics_en).

## Frequently Asked Questions

**What does "pull marketing" mean?** Pull marketing means attracting users with useful content, strong docs, community spaces, and visible products, rather than pushing them with ads or gated sales tactics.

**Why is push marketing viewed more skeptically?** Push marketing is associated with ads, gated content, and sales follow-ups, which often feel misaligned with open-source values and less trustworthy to developers.

**How does open source support product-led growth?** Open source puts the product in front of users throughout the funnel, so people can try it freely, learn from it, and later contact the company only when they need enterprise features or support.

**Why is product quality especially important in open source?** Because users can test the software immediately, see bugs publicly, and switch to alternatives easily if the product is not good.

**What content does VictoriaMetrics publish?** VictoriaMetrics focuses on blog posts and documentation, including announcements, benchmarks, technical articles, tutorials, and guides.

**Why are docs so important in this strategy?** Because docs help users find answers through search and AI tools, reduce support burden, and help people successfully use the product.

**What does "maintaining discoverability" mean?** It means making sure VictoriaMetrics is easy to find on GitHub, Docker, search engines, and AI platforms, while keeping branding and messaging consistent.

**Does VictoriaMetrics completely avoid traditional marketing?** No. VictoriaMetrics still uses traditional marketing for enterprise products, such as outbound campaigns, LinkedIn ads, and sales, but keeps it lean and separate from open source.

---

*I'm Jean-Jérôme Schmidt-Soisson, and I first got involved with open source at MySQL, where I was in sales operations. I then held marketing lead positions (amongst others) at Pentaho, Severalnines & MariaDB, MySQL's sister open-source database. I've had the privilege of leading marketing at VictoriaMetrics since 2021, where we're on a mission to become a standard in the open-source observability space.*

By engineers, for engineers.
