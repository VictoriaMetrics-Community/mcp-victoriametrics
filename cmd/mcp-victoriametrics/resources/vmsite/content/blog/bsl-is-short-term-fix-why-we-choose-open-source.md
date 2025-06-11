---
draft: false
page: blog blog_post
authors:
  - Aliaksandr Valialkin
date: 2023-09-08
enableComments: true
title: "The BSL is a short-term fix: Why we choose open source"
summary: "In this blog post, we’ll explain the controversy over the BSL, and why we believe remaining open source helps businesses remain sustainable for the long-term."
categories: 
 - Community
tags:
 - BSL
 - business source licence
 - open source
 - hashicorp
 - terraform
 - community
images:
 - /blog/bsl-is-short-term-fix-why-we-choose-open-source/preview.webp
---
On August 13 2023, users of HashiCorp’s Terraform forked the software under the name [OpenTF](https://opentf.org/). This was a strong and rapid community reaction to HashiCorp [switching the license](https://www.hashicorp.com/blog/hashicorp-adopts-business-source-license) on their products merely three days before. The list of companies and individuals pledging their support to the new fork has been overwhelming.

The new license that HashiCorp has chosen for its products, the [Business Source License (BSL)](https://www.hashicorp.com/bsl), is no longer open source, but instead source-available. While the distinction may seem subtle, it will have dramatic effects on the open source community and usage of HashiCorp’s products.

We, along with the OpenTF foundation, are disappointed in the announcement from HashiCorp. Many people depend on Terraform, Vagrant, and other products from HashiCorp and are right to feel betrayed by this decision.

VictoriaMetrics is a proudly open source company, and always will be. We don’t believe the BSL is a good thing for open source products, or the longevity of the businesses built around them.

In this blog post, we’ll explain the controversy over the BSL, and why we believe remaining open source helps businesses remain sustainable for the long-term.

## **What is the BSL?**

Traditionally, software licensing has fallen into two categories, open source and closed-source. There are a lot of subtleties here, but in general:

* Open source software makes source code available to anyone for any purpose
* Closed-source software doesn’t make source code available, giving creator(s) more control over its use

Recently, however, a third category of software license has appeared. This is best described as source-available, where the source code is available for everyone to read but there are restrictions on its use for certain companies or commercial use cases.

The Business Source License (BSL) is a recent source-available license that is gaining popularity among some previously open source companies. It has roughly the following properties:

* The source code is publicly available
* Use of the software is free for some use cases but requires a commercial license for others
* BSL licensed code reverts to an open source license after a certain period of time (e.g., 4 years)

Famous projects using it include HashiCorp’s suite of products (Terraform, Vagrant, etc.), Couchbase, CockroachDB, Sentry, and MariaDB.

## **Where did the BSL come from?**

The BSL was [first announced](http://monty-says.blogspot.com/2013/06/business-source-software-license-with.html) in 2013 and first applied to MariaDB. The license was developed by the creator of MySQL and MariaDB, Michael “Monty” Widenius, alongside Linus Nyman, an economist. The identity of the license’s creator [shocked](https://www.adventuresinoss.com/2016/09/12/open-core-returns-from-the-dead-sigh/) quite a few people due to his open source credentials.

Monty hoped for [a few things with](https://monty-says.blogspot.com/2016/08/applying-business-source-licensing-bsl.html) the BSL:

* Commercial companies that depended on open source software would be required to fund continued development of the software
* Companies would see the BSL as a viable alternative to closed source and more source code would be made available
* Open source companies would be able to use this license to generate a comparable amount of revenue to closed source companies and compete on a more even footing

Building a sustainable business is hard, especially in open source, and this license was intended to help open source companies become sustainable by generating more revenue. It allowed software owners to force companies to pay for their software, while also ensuring that the code would eventually be fully open source.

Since MariaDB first started using the BSL, more open source companies have switched to it. Most, if not all companies switching to the BSL are investor-funded, meaning they are under pressure to maximize investor returns.

It’s easy to see why companies may be tempted to use the BSL. It promises a solution to the cash-flow problem in open source. However, this temptation may be short-sighted.

## **Why is the BSL a short-term fix?**

The problem is that switching an open source license to the BSL erodes trust in your product and your company. Clients of your software may be suddenly required to pay under the BSL and are prompted to seek alternatives. This applies to open source projects too, since the BSL cannot be combined with popular licenses like the GPL and requires forks of your code to be [BSL licensed as well](https://www.hashicorp.com/license-faq#mixing-bsl-with-other-licenses).

One of the great strengths of open source licensing is that people want to contribute to your product. Unfortunately, switching to the BSL is likely to reduce your pool of contributors considerably. Many developers will either disagree with the license (see this [HN thread](https://news.ycombinator.com/item?id=37081306), and this [Register comments section](https://forums.theregister.com/forum/all/2023/08/11/hashicorp_bsl_licence/)) or just hear negative things about it and be put off contributing to your project.

Existing contributors that have invested into your product can also end up feeling betrayed. People contribute to open source in part because it’s open source. When code stops being open source, contributors might [become disillusioned](https://news.ycombinator.com/item?id=37082876).

The BSL does revert to an open source license after a period of time, which will be effective at preventing abandonware. However, this leaves anyone that wants to use the software with an open source license with an old and likely vulnerable version. Given the speed at which software development moves, this does not give all the benefits that open source licensing from the start does.

A switch to the BSL after previously being open source can be [seen as taking the marketing benefits](https://news.ycombinator.com/item?id=37082801) of having an open source product, without benefiting the community in return. This can harm your company’s reputation and appear greedy.

## **How else can you make a sustainable business on open source?**

Making a sustainable business means making one that can survive and achieve its goals indefinitely. For most companies, this should be their top priority as it’s very rare that a company can achieve its purpose in a short space of time.

Unfortunately, the goal of sustainability doesn’t always align with investor incentives. The founders and employees of a company have a goal they want to achieve, while the investors in a company are looking to maximize the return on their investment, something largely driven by company profit. Attempting to maximize profit can lead to decisions like implementing the BSL, which harms the long-term sustainability of the company.

The misalignment of incentives between founders and investors implies that one strategy for building a sustainable business is to refuse investor funding. In fact, this has been VictoriaMetrics’ strategy from the start. Bootstrapping our business has allowed us greater freedom than other methods of funding.

Our goal is to make VictoriaMetrics the biggest player in time-series databases and observability. On the face of it, this might sound like a revenue-driven goal, but remember that we give our software away for free! Thanks to not having investors that expect returns from us, we are free to do what’s best for our product and our long-term sustainability.

Our strategy of using a permissive open source license on our code allows us to maximize our customer base and market share. No company, person or project is excluded from using our code which maximizes our potential customer base.

To generate enough revenue to sustain ourselves, VictoriaMetrics provides an [Enterprise version](https://victoriametrics.com/products/enterprise/) of our product and a [cloud solution](https://victoriametrics.com/products/cloud/). Our engineers are the world experts in the software we build, which means that we can provide value to our customers in ways that other companies simply couldn’t.

Other open source companies have chosen to offer hosted versions of their offerings, for example [Gatsby](https://www.gatsbyjs.com/) and [Grafana](https://grafana.com/). Others still have dual licensed their source code, for example with the GPL alongside a commercial paid license as [Qt](https://www.qt.io/) does. There are [many great ways](https://en.wikipedia.org/wiki/Business_models_for_open-source_software) to make a sustainable business while keeping your products fully open source.

## **VictoriaMetrics' position**

VictoriaMetrics believes that open source licensing is the best way for your product and business to survive long-term. Companies that change their license to a source-available or proprietary one may see a short-term boost in their revenue, but customers will lose trust and seek alternatives. Forks and increased competition are inevitable after a switch to the BSL.

We believe licenses like the BSL are short-term fixes to investor pressure to generate revenue. All the companies we listed above on the BSL are investor-funded and therefore under pressure to generate returns.

Our software is all licensed under the [Apache 2.0 license](https://choosealicense.com/licenses/apache-2.0/), which is very permissive. You are free to use our software for whatever you like, and that is part of our products’ appeal. What’s more, we don’t require a contributor agreement. This means that your contributions remain your own copyright and we couldn’t change our license even if we wanted to (which we don’t!).

VictoriaMetrics was founded as, operates as, and will always be an open source company. We fully believe that it is not only possible to build a sustainable business on open source, but the best thing you can do for your software.

Do you need efficient and reliable open source observability tools? Check out our [offerings here](https://victoriametrics.com/products/).

Are you a software engineer interested in monitoring and observability? In case it wasn’t already obvious, we love open source and welcome contributions. Check out our GitHub organization [here](https://github.com/victoriametrics).
