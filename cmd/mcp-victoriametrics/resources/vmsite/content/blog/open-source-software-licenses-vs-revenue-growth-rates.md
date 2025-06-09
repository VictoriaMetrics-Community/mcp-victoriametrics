---
draft: false
page: blog blog_post
authors:
 - Aliaksandr Valialkin
date: 2024-08-30
title: "Open Source Software Licenses vs Revenue Growth Rates"
enableComments: true
summary: "A software license change may have a short term impact on revenue, but the long-term damage can be consequential and take time to fix. Read our CTO’s take on open source software licenses vs revenue growth rates."
categories:
 - Company News
tags:
 - open source
 - proprietary software
 - elasticsearch
 - apache2
 - cloud provider
 - AWS
 - victoriametrics
 - community
 - enterprise software
images:
 - /blog/open-source-software-licenses-vs-revenue-growth-rates/preview.webp
---

# Open Source Software Licenses vs Revenue Growth Rates

I don't understand why pure open-source licenses, such as Apache2, MIT or BSD, should be replaced with a source available license in order to increase profits from enterprise support contracts.

- In most cases, the license change won't force cloud companies to sign an enterprise agreement with you. If they didn't want to pay you before the license change, why would they change their mind after the license change? It is better from a cost and freedom perspective to fork an open source version of your product, and use it for free, like Amazon did with Elasticsearch.
- The license change leads to user base fragmentation. Some of your users will switch to forks run by cloud companies, while others will start searching for alternative open source products. Therefore, after the license change, you’ll start losing users and market share.
- The license change doesn't bring you new beefy enterprise contracts, since it doesn't include any incentives for your users to sign such contracts.

That's why we at VictoriaMetrics [aren't going to change the Apache2 license for our products](https://victoriametrics.com/blog/bsl-is-short-term-fix-why-we-choose-open-source/). Our main goal is to provide good products to users, and to help users use these products in the most efficient way. See also the goals that we have set for the development of our products: https://docs.victoriametrics.com/goals/

## How Do You Compete? 

A question that we hear regularly is: What if a cloud provider, such as AWS, decided to take your open source code and launch their own version of your product, hosting it on their cloud? How do you compete with that?

Should a company such as Amazon make a product on top of open-source VictoriaMetrics, then we'll thank them for it, since this would be good marketing: More people will be made aware of the great products provided by VictoriaMetrics!

There is close to zero probability that a company such as Amazon would ever pay us for such a product. Companies like these will never sign long-term contracts with open-source product vendors, so there is no sense in changing the license of our software from Apache2, to some BSL-like license.

Despite the fact that open source software has been around for quite some time, the status quo on a lot of people’s minds is still proprietary software (although it’s interesting to note that this term only exists because of open source software). For example, we just recently spoke with an analyst, who couldn’t come to terms with the fact that we’re generating revenue based on open source software (which is inherently free) without having taken any investment funds. Our company is entirely fueled by the revenue we generate working with our users and customers, and as long as we continue to provide software that is useful to people, we should be able to continue in that direction.

## How Does a For-profit Company Develop True Open Source Software?

Companies often have justified concerns when they consider choosing proprietary vs open source software for their (often) mission-critical business needs. Why would they start building their projects using open source software, only to find that the owner of that software decides to convert to some restricted  license further down the line? Is it not preferable then to stick with proprietary software from the outset?

Many open source users are aware of what the stakes are for companies such as ours that run a business around the software they’ve built. Their (gracious) concern if they were to use a version of our product hosted by someone such as Amazon is: How do we contribute to your business if we go to a cloud provider who offers support for your software, in addition to the infrastructure?

Our take is that if a user goes to Amazon directly, this is great. They continue using our products,recommend them to their friends, and there’s a chance that next time they'll become our customer. For example, they may not be satisfied with the support from Amazon, or find there are some missing features, or they may switch departments or companies. User and Customer LifeTime Values start early in the process!

We develop open source products, we are profitable, and we have a good revenue growth rate. We make money mostly on high-quality enterprise technical support for our open-source products. Some of our products have enterprise-only features, but many of our paid customers continue using open-source versions of VictoriaMetrics products.

## License Change Does Not Necessarily Equate to Revenue Growth

With this blog post, we’re trying to provide good reasons for why changing a license from truly open source, to some source-available license, makes little sense from a business perspective (in our opinion and experience). Of course, something may change in the future, which could force us to reconsider the decision of sticking with the Apache2 license. However, we currently don't see any reason to change the license, and we’re sure there will be no such reason in the next 10 years.

Finally, and reading between the lines of recent announcements we’ve seen, the main reason to change the license at CockroachDB, Redis, Elasticsearch, MongoDB, TimescaleDB, Grafana and other products, is likely to be [weak revenue growth rate](https://www.techzine.eu/news/analytics/123922/elastic-stocks-take-25-percent-hit-despite-positive-quarter/) combined with investor pressure. 

"Elastic recently presented positive quarterly figures, but also had to issue a revenue warning. The value of Elastic shares fell 25 percent as a result."

Shareholders falsely think that the license change may help increase the revenue growth rate, but I don't understand why…

A software license change may have [a short term impact on revenue](https://victoriametrics.com/blog/bsl-is-short-term-fix-why-we-choose-open-source/), but the long-term damage, for example to community trust, can be consequential and take time to fix.
