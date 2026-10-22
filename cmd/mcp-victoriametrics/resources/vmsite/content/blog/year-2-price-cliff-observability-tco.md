---
draft: false
page: blog blog_post
authors:
  - Michelle Sebek
date: 2026-10-05
title: "The Year-2 Price Cliff: What Your Observability Stack Really Costs Over 3 Years"
summary: "Why aggressive year-1 discounts in observability often lead to 2–3x renewal bills, plus a simple 3-year TCO framework and five questions to ask before signing."
description: "Learn how to evaluate observability pricing over three years, account for renewal increases and usage growth, and compare vendors by total cost of ownership."
enableComments: true
toc: true
categories:
  - Observability
  - Monitoring
tags:
  - observability
  - monitoring
  - FinOps
  - pricing
  - TCO
images:
  - /blog/year-2-price-cliff-observability-tco/preview.webp
---

I'm not the one whose phone lights up at 3 a.m. when production breaks. But I've spent years working alongside the engineers who are, and I've noticed that observability migrations happen for two reasons. Either engineering needed one, or, far more often, a quote landed that looked too good to refuse. The engineers rarely regret the first kind. The second kind they tell me about in year two, usually with a renewal notice in hand. This post is about what happens in year two of those deals, and how to evaluate the next "too good to refuse" quote that lands in your inbox.

One thing before we start: this isn't about any particular vendor, and I'm not going to argue that discounts are bad. They're normal. My point is smaller and more specific. An unusually large year-1 discount is information. If you can't explain why you're getting it, the explanation is probably waiting for you in year two.

## The pattern: land low, reprice high

The deal tends to look like this. A vendor quotes a first-year price 60 to 90 percent below list price. The reasoning is vague. Strategic partnership pricing, a migration credit, some end-of-quarter flexibility that happens to require a signature by Friday. You sign, the product works, everyone moves on. Twelve months later the renewal shows up at two or three times what you paid, and you're negotiating from a very different chair, because by now your dashboards, alerts, runbooks, and everyone's muscle memory live inside their platform.

I want to be careful here, because none of this requires bad faith. It's arithmetic plus incentives. Vendors spend a lot to win an account and plan to earn it back over years, not in year one. In [Zylo's 2026 SaaS Management Index](https://zylo.com/blog/saas-pricing-trends), 79% of IT leaders reported a price increase at renewal in the past year, with typical hikes of 8–12% and aggressive vendors pushing 15–25%. Only 29% of SaaS contracts cap renewal increases at all, per [Common Paper's contract benchmark data](https://commonpaper.com/resources/contract-benchmark-2024-q1/); where a cap exists, 5–8% a year is the standard band, and [Salesforce's own MSA bakes in 7%](https://www.saastr.com/whats-a-typical-price-increase-i-can-expect-when-renewing-my-saas-subscriptions/). Then there's the quiet one: expansion usage billed at *then-current list price* rather than your negotiated rate. Stack all of that on top of a discount that was always going to expire, and the year-2 number isn't an insult. It's the plan, working.

> **The reframe that matters:** the vendor isn't pricing year one. They're pricing the relationship. A year-1 number that looks irrational as a price makes perfect sense as an acquisition cost.

## Why observability bills grow even when nothing changes

There's a second mechanism underneath, and it hits observability harder than almost any other software category: your usage grows even when your business doesn't.

- **Telemetry outgrows infrastructure.** Engineers add labels, custom metrics, traces, log fields. Cardinality routinely grows two or three times faster than the fleet it describes, so per-GB and per-series pricing compounds on its own. No headcount or traffic growth required.
- **Defaults are generous.** Agents auto-discover everything and ship it all. No alert fires when telemetry volume creeps up. The first notification anyone gets is the invoice.
- **Growth gets billed at list.** Discounts usually cover the committed volume in the order form. The growth on top, the part you can least predict, lands at on-demand or then-current rates.

Put those together. The discount expires, the list escalates, and your unit count went up 30 or 50 percent on its own. A $50k year turns into a $150k year and nobody lied to anyone.

## A simple 3-year TCO framework

You don't need a procurement department to protect yourself here. One spreadsheet, four inputs per vendor.

1. **Year-1 price.** The number on the quote. The easy one.
2. **Renewal basis.** What is year 2 calculated from, your price or list? What does the discount decay to, and is there a cap in writing? If the vendor won't commit on paper, model year 2 at list, because that's what "we'll take care of you at renewal" converts to.
3. **Growth at the margin.** Your realistic usage growth (30%+ a year is normal for telemetry, not pessimistic) times the marginal rate for growth. Not the blended rate on the quote. The marginal one.
4. **Switching cost, both directions.** What it costs to get in, and what it would cost to get out. The second number is your negotiating position at every renewal for the life of the relationship. If it rounds to infinity, expect pricing that reflects that.

Then compare 3-year totals and only 3-year totals. Here's the arithmetic with two made-up vendors on the same workload:

<div class="table-responsive">

| | Vendor A ("the deal") | Vendor B (transparent) |
|---|---:|---:|
| List price | $150,000 / yr | $95,000 / yr (published) |
| Year 1 | $52,500 (65% off) | $95,000 |
| Year 2 | $129,600 (discount → 20%, +8% list escalation) | $99,750 |
| Year 3 | $139,968 (same terms, compounded) | $104,738 (5% cap) |
| **3-year total** | **$322,068** | **$299,488** |
| Avg. effective / yr | $107,356 | $99,829 |

</div>

Vendor A's first-year price looks dramatically cheaper. Over three years, it costs about $22K more. And that's before additional usage growth or migration costs in either direction. The lesson isn't that discounts are bad. It's that year-one pricing tells you very little unless you understand how the economics change over the full term. The cheapest quote and the lowest total cost are often two different things.

*Numbers are illustrative. Actual multi-year pricing depends on committed usage, expected scale, contract term, and negotiated total contract value. Run the model using the terms in front of you.*

## Questions to ask before you sign

If you take one thing from this post, take these five questions. Ask them in writing. "Let's discuss at renewal" is an answer. Just not a reassuring one.

- **What's the list price, and why am I getting this discount?** Discounts with a structural reason behind them, like a multi-year commit or committed volume, tend to survive. Discounts with no reason tend to expire.
- **If I change nothing, what does year 2 cost?** Same workload, exact number, on paper.
- **What will this cost over the full contract term?** Don't evaluate year one in isolation. Ask for the assumptions behind years two and three, including expected scale, committed usage, and the total contract value.
- **What rate applies above my commit?** This is where observability budgets actually die, so don't skip it.
- **If I leave in three years, what does it cost to get my data and my config out?** However they answer, that's your future leverage, priced.

## What good looks like

The fix isn't cynicism about every vendor. It's symmetry. A vendor with published pricing, boring renewal mechanics, and a product that doesn't hold your data hostage has voluntarily given up the whole year-2 playbook, and that's worth real money. Ask every vendor on your shortlist to match it and watch what happens.

For what it's worth, that's the standard we try to hold ourselves to at VictoriaMetrics. The core is open source, so the exit door stays visible. Enterprise pricing is built to be predictable at renewal, not just attractive at signature. And you can see how the plans are structured on our [Plans & Features](/plans-features/) page instead of finding out in year two. By engineers, for engineers, including the one who has to defend the budget line next year.
