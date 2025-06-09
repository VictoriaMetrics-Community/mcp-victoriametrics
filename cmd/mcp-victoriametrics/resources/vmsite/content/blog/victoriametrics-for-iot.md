---
draft: true
page: blog blog_post
authors:
  - Phuong Le
date: 2025-02-02
title: "Use VictoriaMetrics as Database and Solution for IoT"
summary: "IndexDB acts as vmstorage's memory - it remembers which numbers (TSIDs) belong to which metrics, making sure your queries get answered fast. This article walks through how this system works, from the way it organizes data to how it keeps track of millions of timeseries."
enableComments: true
categories:
  - Open Source Tech
  - Monitoring
  - Time Series Database
tags:
  - vmstorage
  - indexdb
  - open source
  - database
  - monitoring
  - high-availability
  - time series
images:
  - /blog/vmstorage-how-it-handles-query-requests/vmstorage-how-indexdb-works-preview.webp
---

Here is the short answer: Yes, VictoriaMetrics is a good choice for IoT and many users have already used it.

In this discussion, I will discuss why choose VictoriaMetrics and the problem our users are facing can be fixed by VictoriaMetrics.

Use two distinct VictoriaLogs instances for server logs and IoT events. This will simplify maintenance and management of these instances - you can specify different retention policies per every VictoriaLogs instance, you can use different backup strategies for these instances, you can run them on different hardware (CPU, RAM, disk space, disk IO), which fits better for the particular workload.

VictoriaLogs should be much better than TimescaleDB for IoT events:

it should be easier to setup and operate - you don't need to provide fine-tunes configs to VictoriaLogs, and you don't need to create any database schemas

it should use less RAM, CPU, disk space and disk IO

it should perform typical queries at much faster speed than TimescaleDB.

Try ingesting IoT events into both VictoriaLogs and TimescaleDB in parallel - and then choose the system, which is better suited for your workload. I bet VictoriaLogs will win with a high margin.