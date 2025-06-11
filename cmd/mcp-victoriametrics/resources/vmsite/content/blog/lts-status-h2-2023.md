---
draft: false
page: blog blog_post
authors:
 - Artem Navoiev
date: 2023-10-15
enableComments: true
title: "VictoriaMetrics Long-Term Support (LTS): Current State"
summary: "Overview of LTS releases, deprecation of 1.79, review of the most recent LTS 1.93"
categories: 
 - Product News
tags:
 - victoriametrics
 - LTS
 - long-term support release
images:
 - /blog/lts-status-h2-2023/preview-lts.webp
---
{{< image href="/blog/lts-status-h2-2023/lts-state-h2-2023.webp" class="wide-img" alt="The state of VictoriaMetrics LTS releases" >}}

We release VictoriaMetrics several times a month, including at least one major update. However, because these new releases often introduce new features, they may be less stable. That's why we also regularly publish [Long-term support releases (LTS)](https://en.wikipedia.org/wiki/Long-term_support)  alongside our regular releases. These LTS versions focus exclusively on bug fixes without new features and performance improvements.

We [committed](/blog/lts-status-h1-2023/) to publishing LTS versions every six months and supporting them for one year.
Let me provide an overview of the most recent changes in LTS releases.

## **LTS release: v1.79.x. Deprecated**

We initially released this version in July 2022 and support for v1.79.14 ended in July 2024. 

We recommend upgrading to the recent releases if you still use v1.79.14.   

## **The previous LTS release: v1.87.x**

The previous LTS release, [v1.87.x](https://docs.victoriametrics.com/CHANGELOG.html#v1879) will continue to receive bug fixes until January 2023.  

## **The current LTS release: v1.93.x**

The most recent [LTS release](https://docs.victoriametrics.com/CHANGELOG.html#v1935) of VictoriaMetrics brings substantial performance improvements and features compared to 1.87.x. 

We will continue to support it until July 2024.

## **Upgrade from v1.87.x to v1.93.x**

We recommend reviewing the [CHANGELOG](https://docs.victoriametrics.com/CHANGELOG.html) between these releases before performing the upgrade to be prepared for new features or potential issues. You can upgrade from v1.87.x to v1.93.x following the guidelines provided in [these docs](https://docs.victoriametrics.com/#how-to-upgrade-victoriametrics). If you plan to upgrade the VictoriaMetrics cluster, please refer to [these instructions](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#updating--reconfiguring-cluster-nodes).

Version 1.93.x has been thoroughly tested and is used by many people in production. However, please be aware that there have been some non-backward compatible changes introduced between v1.87.x and v1.93.x.
 If you encounter any issues with v1.93.x, we recommend downgrading to v1.91.3.

We strongly suggest upgrading your [Grafana dashboards](https://docs.victoriametrics.com/#monitoring) and [alerting rules](https://docs.victoriametrics.com/#monitoring) or VictoriaMetrics components, as each release incorporates various improvements.
