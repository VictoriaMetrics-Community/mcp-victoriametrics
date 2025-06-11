---
draft: false
page: blog blog_post
authors:
  - Artem Navoiev
  - Aliaksandr Valialkin
date: 2023-02-17
enableComments: true
title: "VictoriaMetrics Long-Term Support (LTS): Commitment, Current and Next LTS Versions"
summary: "Overview of LTS releases, commitment from the VictoriaMetrics team about them, and migration from one LTS to another"
categories: 
  - Product News
tags:
  - victoriametrics
  - LTS
  - long-term support release
images:
  - /blog/lts-status-h1-2023/preview-lts.webp
---
{{< image href="/blog/lts-status-h1-2023/roadmap1.webp" class="wide-img">}}

VictoriaMetrics is always improving, with frequent updates adding new features, performance improvements and bug fixes listed at the [CHANGELOG page](https://docs.victoriametrics.com/CHANGELOG.html). We usually make at least a single release every month. All the new features and bug fixes go to the latest release. That’s why we recommend [periodically upgrading VictoriaMetrics components to the latest available release](https://docs.victoriametrics.com/#how-to-upgrade-victoriametrics). But the latest release may also contain bugs in the latest features. So we decided to start publishing [Long-term support releases (LTS)](https://en.wikipedia.org/wiki/Long-term_support) on top of usual releases, which contain only bug fixes without new features and performance improvements.

## **Our Commitment**

We are committed to supporting every LTS release for a year and marking one of our releases as LTS every 6 months.

This allows our users, who prefer stability over new features and performance improvements, to stay on LTS releases. There is a 6 months overlap between LTS releases, which gives enough time for the upgrade to the next LTS release.


## **The current LTS release: v1.79.x**

In July 2022, we cut the LTS version, which is v1.79.x release. This release was explicitly marked as LTS on Github, and its latest version is [v1.79.8](https://docs.victoriametrics.com/CHANGELOG.html#v1798). This release includes all the important bug fixes from the subsequent releases listed at the [CHANGELOG](https://docs.victoriametrics.com/CHANGELOG.html).

The v1.79.x line of releases will stop receiving new bug fixes after July 2023.

## **The next LTS release: v1.87.x**

We are going to mark the v1.87.x release as the next LTS. It will receive bug fixes until January 2024. Both v1.87.x and v1.79.x releases will receive bug fixes during the next 6 months.

## **Upgrade from v1.79.x to v1.87.x**

It is recommended to read the [CHANGELOG](https://docs.victoriametrics.com/CHANGELOG.html) between these releases before the upgrade in order to be prepared for new features or issues. It is OK to upgrade from v1.79.x to v1.87.x according to [these docs](https://docs.victoriametrics.com/#how-to-upgrade-victoriametrics). If you are going to upgrade the VictoriaMetrics cluster, then please read [these docs](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#updating--reconfiguring-cluster-nodes). In case of unlikely issues it is OK to downgrade from v1.87.x to v1.79.x.

We suggest upgrading the [Grafana dashboards](https://docs.victoriametrics.com/#monitoring) and [alerting rules](https://docs.victoriametrics.com/#monitoring) for VictoriaMetrics components, since they contain various improvements with each release.
