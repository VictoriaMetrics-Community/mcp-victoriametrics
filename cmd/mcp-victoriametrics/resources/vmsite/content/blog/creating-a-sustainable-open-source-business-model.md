---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
date: 2025-09-24
enableComments: true
title: "Creating a Sustainable Open Source Business Model - Introduction"
summary: "Open source defies everything you've ever heard or learned about business before. This blog post is an introduction to how we’re creating a sustainable business model rooted in open source."
categories: 
 - Company News
tags:
 - open source
 - business models
 - licensing
 - victoriametrics
images:
 - /blog/creating-a-sustainable-open-source-business-model/preview.webp
---

Open source defies everything you've ever heard or learned about business before (author’s quote).

Yes, open source software has been around since the 90s, but there’s still little else like it. If anything, as time has gone on, we’ve added adjacent concepts like "open core" and "source available" that have added complexity to a model that isn’t that straight forward to grasp to begin with.

VictoriaMetrics is an open source company. As our team grows, new colleagues join us who sometimes have some to no experience with open source, which can result in confusion because open source can be so unintuitive at first glance.

Open source, however, is not a bit of jargon. We need everyone on the team to adopt the open source mindset, and we need the industry to understand our position in the open source ecosystem, if we're going to be as successful as we can be.

## Open source beyond "free software"

Open source doesn't just mean "free software." Open source is a complex ecosystem that includes projects with thousands of maintainers and millions of users, as well as projects with individual maintainers and single-digit user bases.

Open source also defies what we know about traditional businesses. In a pure business sense, open source doesn't have anything to sell, which can be counterintuitive considering how many open source businesses make millions and millions of dollars.

Our mission is to make VictoriaMetrics a standard for modern observability stacks over the next three years. We aim to offer an open source monitoring and observability solution that is simple, reliable, and efficient for metrics, logs, and traces.

This mission is ambitious and reflects our commitment to becoming a standard in our industry, rather than just the best or the number one. Open source, and the mindset it gives us, is essential to accomplishing this mission.

## VictoriaMetrics and the open source mindset

VictoriaMetrics was founded as an open source company; we operate as one, and we will always be one. Open source permeates everything we do, and we map our most important company principles back to open source principles.

We believe it is not only possible to build a sustainable business on open source but that [it's the best thing anyone can do for their software](https://tech.eu/2024/05/03/how-victoriametrics-open-source-approach-led-to-mass-industry-adoption/). We're fully committed to open source because open source software helps us achieve our primary goal: providing good products and helping users use those products efficiently.

Open source is our path not just because we think it will work, but because the open source values align with our own. The principles that underlie open source, for example, emphasize open exchange, collaborative participation, transparency, and community-oriented development – much like we do. Similarly, open source prizes community, even as it creates business opportunities.

On the one hand, each open source project aims to grow its own community of users; on the other, the problems that open source projects solve can be opportunities for businesses willing to take on the work of balancing licensing and ownership with profit and control. These tensions are not without controversy (see Redis and HashiCorp), but there are many more examples of open source businesses benefiting all involved (see Linux and Red Hat, Git and GitHub, and more).

Open source, then, is an extension of our values, not a distribution channel or a mere technology choice. These values are just as applicable to our company as to open source in general. We exchange information openly, collaborate effectively, release updates rapidly, and maintain transparency internally and externally. Our community of users and contributors is vital to our success, and our company uses a flat organizational structure.

With this open source mindset framing the way we think and operate, you’ll hear numerous things said here, inside and outside of the company, that scramble traditional business logic, including:

* "If a company like Amazon makes a product on top of open source VictoriaMetrics, then we'll thank them for it. It would be good marketing."
* "If they decide to continue using our open source product, that's also a win, especially if they talk about us to others."
* "Open source is good marketing."

This mindset is foundational, and everything we do emerges from it.

## The licensing question

Licensing is a nuanced issue in open source. Without understanding how licensing works, it can be hard to navigate the variety of open source projects and businesses out there.

Organizations like The Open Source Initiative even offer "Certified Open Source" and "Open Source Approved License" certifications that demonstrate a "[consensus on what constitutes Open Source](https://opensource.org/programs)" that people can gather around and anchor to.

Different licenses have varying levels of permissiveness and restrictions. For example, the Apache 2 license, [which we use](https://www.theregister.com/2023/12/11/victoriametrics_interview/), is one of the most permissive licenses, and it allows users to modify and distribute software freely. Other licenses, like GPL, have different conditions and restrictions.

We prioritize using the most permissive licenses with the fewest restrictions. Users are free to use our software for whatever they like, which is part of our product's appeal. We don't require a contributor agreement, which means user contributions remain their own copyright, and it disallows us from changing our license even if we wanted to.

Other well-known licenses include:

* Mozilla Public License
* PHP License
* MIT License
* Business Source License (BSL)
 
BSL is particularly interesting. MariaDB Corp. Ab. originally created BSL and renamed it to BUSL after it was coined the “Bullshit (BS) License” by the community following its introduction. BUSL is a software license that publishes source code but limits the right to use the software to certain classes of users.

BUSL is not an open source license but a source-available license that mandates an eventual transition to an open source license. HashiCorp recently moved to this type of license in a controversial move.

Redis and MongoDB have made similarly controversial licensing decisions, which is why licensing is frequently an elephant in the room when it comes to figuring out how open source business models can be sustainable, if not profitable.

The truth is, there’s no one way to run an open source business. Organizations like the Open Source Initiative (OSI), Free Software Foundation (FSF), and the Linux Foundation all have different approaches to open source, and new businesses take new angles on open source every day.

## Business model diversity in open source

Open source is foundational to our business model. Understanding our business model, however, requires understanding how our choice of model fits into the many available ways to monetize open source. 

### The many business models for open source software

Open source can be extremely profitable. Just take a look at some of the most well-known projects and the market values of the leading companies that provide them.

<p style="max-width: 451px; margin: 1rem auto;">
<img src="/blog/creating-a-sustainable-open-source-business-model/oss-projects-ranking-market-value.webp" style="width:100%" alt="Battery Open Source Software Index (BOSS)"><figcaption style="text-align: center; font-style: italic;">Battery Open Source Software Index (BOSS)</figcaption>
</p>

Source: [Wikipedia](https://en.wikipedia.org/wiki/Open_source)

The business models across these ten businesses vary, and they vary even further across the rest of the open source ecosystem. You can think of these models as running along a spectrum from less commercial to more commercial.
* **Open source purist**: The software is open source, and there are no commercial elements. Revenue is purely based on services.
* **Open source with commercial elements**: The software is open source and has some commercial elements. Revenue is based on commercial features and services.
* **Open source services**: The company sells services for open source software and has some open source tools to use.
* **Open source front**: The software is presented as open source, but it has some form of restrictive license.
* **Commercial business that "gives back"**: The company has a proprietary software model, and it makes some of its innovations available as open source software.

### Examples of different businesses and their models

Examples of these businesses in action include:

* **Grafana**: Grafana operates on an open core business model, where the core software is open source and free to use, while advanced features and enterprise plugins are available for a fee. They offer both self-managed and cloud-based solutions, including Grafana Enterprise and Grafana Cloud, which provide additional scalability, security, and support features for enterprise customers.
* **ClickHouse**: ClickHouse offers a fast, open source column-oriented database management system. They provide a cloud service, ClickHouse Cloud, which is optimized for development use cases and offers a serverless experience for managing large-scale data analytics.
* **Red Hat**: Red Hat's business model is based on providing support, services, and training for their open source software. They do not sell the software itself but offer subscriptions for support and additional services.
* **HashiCorp**: HashiCorp used an open-core model, where its core tools are open source. It offers enterprise versions with additional features and managed services. HashiCorp also provides a cloud platform, HashiCorp Cloud Platform (HCP). They recently [announced a move to a BSL license](https://www.hashicorp.com/en/blog/hashicorp-adopts-business-source-license).
* **Percona**: Percona focuses on providing enterprise-class support, consulting, and managed services for open source databases like MySQL, PostgreSQL, and MongoDB - amongst others, as well as its own open source software. 
* **MariaDB**: MariaDB offers an open source database with additional enterprise features available through a subscription. They use a combination of open source licensing and enterprise offerings, including the Business Source License (BSL) for some products.
* **MongoDB**: MongoDB used to operate on a dual licensing model, offering a free open source version and a paid enterprise version with additional features and support. They also provide a cloud service, MongoDB Atlas, which is a fully managed database service. The open source version mutated to a 'free community version' over time.
* **Oracle/MySQL**: Oracle offers MySQL under a dual licensing model, where the community version is open source, and the enterprise version includes additional features and support. Oracle also provides cloud services and other proprietary software products.

There are many different business models that all gravitate around the concept of open source, and if we’re looking at it from a business perspective only, companies will create a model that they think works best for them and for their business objectives. Within the broad categories above, business models can get [even more granular](https://en.wikipedia.org/wiki/Business_models_for_open-source_software). They could also run on:

* Professional services
* Voluntary donations and crowdsourcing
* Training and certification
* Partnerships with funding organizations
* Bounty-driven development and reverse-bounties
* Open core or dual-licensing
* Selling proprietary additives

This variety can seem confusing, but the variety is the point. As Nadia Eghbal, author of *Working in Public: The Making and Maintenance of Open Source Software*, writes, "The term 'open source' refers only to how code is distributed and consumed. It says nothing about how code is produced. 'Open source' projects have nothing more in common with one another than 'companies' do. All companies, by definition, produce something of value that is exchanged for money, but we don’t assume that every company has the same business model."

Similarly, we don’t expect similar open source projects to perform the same way, even if they’re in the same category.

### The 1% Rule

Despite this diversity, almost all open source business models depend on a core pattern: The (unwritten) 1% Rule. The 1% Rule refers to the conversion rate of users who eventually become paying customers. This rate generally falls around 1% and is based on the number of unique downloads.

While this might seem low, the value of these conversions increases over time as the user base grows. Our business model, for example, is designed to capitalize on this conversion rate by offering professional services and enterprise features to our customers.

## The VictoriaMetrics business model

Open source runs through everything we do.

### Open source gives our users confidence

Prospective customers want to know: "What happens if VictoriaMetrics Inc. disappears (for whatever reason)?" Because our products are open source with an active user community, customers can feel safe in the knowledge that they're investing in technology that will last.

### Open source helps us position our solutions
When team members ask things like, "Can we add a 'sign up for cloud' action button on our homepage?" or "Can we highlight enterprise and cloud in our documentation?", our answers are informed by open source values.

### Open source shapes the customer journey 
The path from user or prospect to customer goes from open source to enterprise. Future customers will make their decisions based on:

* The quality of our open source products.
* The size of our open source user community.
* How active our user community is.
* How many outside contributors we have.
* Interactions they’ve had previously with our engineers and developers.

As our user community grows, so will our business opportunities, which is why we treat all of our users with care. When potential users approach our booth at a conference, for example, we don’t just automatically collect their contact details and drop them in our funnel. We always ask whether they actually want to be contacted after the event. The open source mindset comes before sheer lead generation.

### Open structures our funnel

We have an organically growing inbound funnel that includes the following steps:

1. Users look for a solution for their monitoring and observability challenges.
2. They typically find resources online that point them to one of our open source solutions or hear recommendations from others.
3. They start working with our open source solutions and are typically happy with the result.
4. We may never hear from them or know they're using our products.
5. Some users work for organizations that need something specific, such as SLAs, professional support, security features, or additional scalability.
6. They will look at our enterprise offering and see whether that's what they need. Some will find ways to carry on by themselves. Some will contact us.
7. Organizations that end up contacting us about our commercial offering tend to be well-established, medium to large-sized companies with complex and/or large sets of data and infrastructure.

We sell on-premises licensing deals ([VictoriaMetrics Enterprise](https://victoriametrics.com/products/enterprise/)), which, for the most part, consist of providing professional services. We also sell additive proprietary features and operate a SaaS model through [VictoriaMetrics Cloud](https://victoriametrics.com/products/cloud/).

### Open source makes us stand out

Open source makes us unique. VictoriaMetrics is:

* Fully open source (with an enterprise offering).
* By engineers for engineers.
* Self-funded/customer-funded.
* Led by four founders.
* Growing organically.
* Inbound-driven.
* Marketed by product-led growth, word of mouth, and content.

Open source makes community a vital part of our business model. Our user community contributes to our products, provides valuable feedback, and helps spread the word about our solutions.

The size and activity of our community are essential indicators of our success and long-term viability. By fostering a strong community, we can ensure that our products continue to meet the needs of our users and remain relevant in the market.

## Open source: How we started and how we’ll grow

Understanding open source is crucial for our success as a company. By adopting the open source mindset, embracing the technological advantages open source offers, navigating the licensing landscape, and fostering a strong community, we can achieve our mission of becoming a standard in modern observability stacks.

Our business model, which combines open source software with commercial offerings, allows us to generate revenue while staying true to our open source roots. As we continue to grow and evolve, our commitment to open source will remain a cornerstone of our identity and success.

How exactly will we achieve our mission? This will be the premise of a follow up to this blog post. Please let me know if you have any questions or comments in reaction to this initial post. I’d love to include elements of discussion resulting from this into the next one.

*I'm Jean-Jérôme Schmidt-Soisson and I first got involved with open source at MySQL, where I was in sales operations. I then held marketing lead positions (amongst others) at Pentaho, Severalnines & MariaDB, MySQL’s sister open source database. I’ve had the privilege of leading marketing at VictoriaMetrics since 2021, where we’re on a mission to become a standard in the open source observability space.*