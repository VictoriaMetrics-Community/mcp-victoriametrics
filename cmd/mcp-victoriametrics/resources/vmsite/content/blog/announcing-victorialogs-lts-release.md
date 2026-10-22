---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-10-13
title: "VictoriaLogs Has Now a Long Term Support (LTS) Release"
enableComments: true
featuredPost: false
summary: "VictoriaLogs v1.52.1 is our first Long Term Support release. Enterprise users get 12 months of guaranteed security and bug fix backports, with no new features or breaking changes."
description: "VictoriaLogs LTS is here: v1.52.1 is the first Long Term Support release for VictoriaLogs Enterprise, with 12 months of security fixes and bug fix backports."
categories:
  - Product News
tags:
  - victorialogs
  - lts
  - observability
  - logs
  - roadmap
images:
- /blog/announcing-victorialogs-lts-release/preview-1.webp
---

VictoriaLogs now runs in production everywhere, and many teams running it have told us the same thing: everyone wants new features, but it's hard to keep upgrading their logging backend every few weeks. They want the mix of stability and up-to-date security that only LTS releases can satisfy.

We're happy to announce the first Long Term Support (LTS) release of VictoriaLogs: **v1.52.1**. This is the start of a new v1.52.x LTS line.

VictoriaLogs follows the same LTS model we've run with VictoriaMetrics since 2022: every LTS has guaranteed support for security backports for 12 months since its first release.

## What are LTS Releases?

[Long Term Support](https://docs.victoriametrics.com/victorialogs/lts-releases/) (LTS) releases are stable release lines that receive bug fixes and security fixes for 12 months after the initial release, without new features or breaking changes. 

If you've already used LTS releases for VictoriaMetrics, you know what to expect. LTS releases are published for [VictoriaLogs Enterprise](https://docs.victoriametrics.com/victoriametrics/enterprise/) users. A new LTS line is created every 6 months, so we support two LTS lines at any time, and Enterprise users always have up to 6 months to move from one line to the next.

If you run the open source version, you won't be left behind. VictoriaLogs Enterprise is built on the open source code, and every bug fix and security fix included in an LTS release is also in the latest community release, currently v1.53.0. We recommend that community users upgrade to the [latest community release](https://github.com/VictoriaMetrics/VictoriaLogs/releases/tag/v1.53.0) regularly to stay protected.

## VictoriaLogs LTS v1.52.1

VictoriaLogs v1.52 makes a good base for a long-lived line. It switched the Docker images from Alpine to distroless to shrink the attack surface. This is the first VictoriaLogs LTS line, so for now it's the only one. The next LTS line is planned about 6 months from now, after which two lines will be supported side by side.

### Update Notes

> **Caution:** If you run v1.52.0, we recommend upgrading since that version could panic or drop logs with duplicate [stream field](https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields) names at ingestion, and panic when querying such logs written by earlier releases.

- Go toolchain update: The Go builder moves from Go 1.26.5 to Go 1.26.8, picking up the fixes from the 1.26.6, 1.26.7, and 1.26.8 patch releases.
- Deletion API hardening: The `/delete/run_task` endpoint of the deletion API now accepts only POST requests, which blocks some SSRF-based deletion attacks.
- vlinsert shutdown: On graceful shutdown, vlinsert now sends buffered logs to vlstorage instead of dropping them. This is limited by a new `-insert.drainTimeout` flag, which defaults to 5s.
- Cluster rerouting: Rerouted data is now spread evenly across available vlstorage nodes. Before, nodes next to unavailable ones in the list could get overloaded.
- Query cancellation: Queries running on vlstorage are now canceled when the matching vlselect query is canceled or times out. Before, they could keep running and waste CPU and disk IO.
- Sort order: `sort by (_time) limit N` no longer returns logs out of order when pipes like `unpack_json` overwrite `_time`.
- vlagent collection fixes: The File Collector no longer panics when a log file is deleted, and it no longer fails at startup when it can't read neighboring files. Incomplete last lines are dropped with a warning instead of being joined to the first line of the rotated file.

[Read the complete changelog](https://docs.victoriametrics.com/victorialogs/changelog/#v1521)

## Our commitment

For VictoriaLogs Enterprise users on the v1.52.x line, we guarantee 12 months of:

- Security backports: Every security fix, including Go toolchain updates, is backported to v1.52.x and shipped as a new patch release.
- Bug fix backports: Important bug fixes land in the LTS line without the new features, flags, or behavior changes that come with regular releases.
- Enterprise support: Our team supports the v1.52.x line for the full 12 months, with a predictable path to the next LTS line.

At the same time, we'll keep shipping new VictoriaLogs features at the usual pace in community releases. LTS doesn't slow the open source version down. Instead, it gives teams that need stability a supported place to stand as we move forward.

Want LTS for your VictoriaLogs deployment? [Get in touch](https://victoriametrics.com/contact-us/) or start a [free Enterprise trial](https://victoriametrics.com/products/enterprise/).

## Frequently Asked Questions

**Which version should I run?** If you're a VictoriaLogs Enterprise user who values stability, run the latest v1.52.x LTS release, currently v1.52.1. If you use the open source version, or want the newest features, run the latest community release.

**How long is v1.52.x supported?** For at least 12 months from the v1.52.0 release on July 16, 2026.

**Will LTS releases get new features?** No. LTS lines receive only security and bug fixes. New features ship in regular releases and arrive in LTS when the next line is created.

**I'm on the open source version. Am I missing security fixes?** No, as long as you stay up to date. Every fix in an LTS release is also in the latest community release. Security patches are only guaranteed for LTS lines and the latest community release, so upgrade regularly.

**When is the next LTS line?** New LTS lines are published every 6 months. Once the next one ships, v1.52.x and the new line will be supported in parallel, giving you up to 6 months to migrate.

**Where can I learn more?** See the [LTS documentation](https://docs.victoriametrics.com/victorialogs/lts-releases/), [contact us](https://victoriametrics.com/contact-us/), or ask in our [public Slack](https://inviter.co/victoriametrics).
