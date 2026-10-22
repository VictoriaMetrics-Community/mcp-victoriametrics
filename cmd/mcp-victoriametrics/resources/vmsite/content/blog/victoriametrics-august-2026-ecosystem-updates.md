---
draft: false
page: blog blog_post
authors:
  - Pablo Fernandez
date: 2026-09-08
enableComments: true
title: "VictoriaMetrics August 2026 Ecosystem Updates"
summary: "The VictoriaMetrics team shipped two core releases with important security fixes. The vmanomaly team shipped two updates with improved scaling and more efficient models. Plus, Tech Talks are back, this time demoing AI-assisted incident troubleshooting."
toc: true
categories:
  - Product News
  - Company News
tags:
  - victoriametrics
  - vmanomaly
  - observability
  - new features
images:
  - /blog/victoriametrics-august-2026-ecosystem-updates/preview-5.webp
---

August is a month when things typically quiet down. At VictoriaMetrics, however, we kept up the pace. This month saw two VictoriaMetrics releases, vmanomaly updates, a new Tech Talk, and great blog posts from the VictoriaMetrics team you might have missed:

- Jean-Jerome Schmidt-Soisson [made a case for open source being a good marketing strategy](https://victoriametrics.com/blog/open-source-is-good-marketing/).
- Phuong Lee continued the series on handling concurrency in Go with a post [on sync.Map and its hash trie design](https://victoriametrics.com/blog/go-sync-map-hash-trie/).
- Diana Todea looked at how vmanomaly, paired with MCP and natural-language AI skills, [lets engineers describe a monitoring goal in plain English and get a validated anomaly-detection config back](https://victoriametrics.com/blog/observabilitys-sixth-sense-grounding-anomaly-detection-in-reality/).
- Jesus Espino published the [Go interactive tour for v1.27](https://victoriametrics.com/blog/go-1-27/) with the main features and changes in the language's latest release.

## VictoriaMetrics 1.151.0 is out: Critical HTTP Basic Auth bypass

> This release addresses five security advisories. We recommend upgrading if you're running a vulnerable version, especially for deployments using `-httpAuth.username` and `-httpAuth.password` flags for access control.
> Affected versions are:
> - v1.136.16 and lower (including EOL LTS v1.122)
> - v1.137.0 up to v1.148.2
> - v1.149.0 up to v1.150.0

- Fixes an HTTP Basic Auth bypass: The vulnerability affected authentication controlled by the `-httpAuth.username` and `-httpAuth.password` flags. Versions v1.151.0, v1.148.3, and v1.136.17 **are not affected by this vulnerability**.
- Prevents vmselect panics on low disk space: vmselect now fails the query request when there isn't enough disk space to store temporary search results (instead of panicking as it did before).
- Adds a clean shutdown indicator: The new `vm_app_prev_shutdown_unclean` metric reports whether the previous shutdown was clean. It's exposed by vmsingle, vmstorage, vmselect, and vmagent, making it easier to spot crashes or forced restarts.
- Improves Kafka consumer troubleshooting: New metrics for the vmagent Kafka consumer help identify whether a slowdown originates from the Kafka broker or from the vmagent consumer itself. See the [Kafka slow consumption docs](https://docs.victoriametrics.com/victoriametrics/integrations/kafka/#slow-consumption) for details.
- Fixes uneven metrics metadata distribution: Metadata is now distributed more evenly among storage nodes, resolving an issue where users could see unexpected partial metadata results.

[Read the changelog](https://docs.victoriametrics.com/victoriametrics/changelog/#v11510)

## VictoriaMetrics 1.150.0 released: multi-tenancy by default and faster relabeling

> This version is impacted by an HTTP Basic Auth bypass vulnerability. We recommend upgrading to v.1.151.0.

- Multi-tenancy via headers is enabled by default: [`--enableMultitenancyViaHeaders`](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#multitenancy-via-headers). VictoriaMetrics now accepts tenant IDs through HTTP headers without extra configuration, aligning multi-tenancy configuration with VictoriaLogs and VictoriaTraces.
- Faster relabeling: relabeling rules that use multiple exact metric-name matchers run up to 30% faster.
- Automatic scaling of vmselect concurrency: `-search.maxConcurrentRequests` now scales with the number of available CPU cores instead of stopping at a fixed maximum of 16.
- Consistent error codes for invalid queries: `/api/v1/query` and `/api/v1/query_range` return 400 Bad Request instead of 422 for invalid parameters, matching Prometheus behavior.
- More efficient caching during shutdown: prevents storage from starting cold when a rotation occurs just before restart, avoiding unnecessary cache rebuilds.
- Linode support added to vmagent: vmagent now supports Linode as a service discovery source, letting users automatically discover and scrape targets running on Linode infrastructure.

[Read the changelog](https://docs.victoriametrics.com/victoriametrics/changelog/#v11500)

## vmanomaly v1.30.3-v1.30.4 is out with more stable sharding and reliable model-state restoration

The new vmanomaly release adds to the [new features shipped during the 2nd quarter of 2026](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2026-q2/).

> We recommend upgrading directly to [v1.30.4](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1304), as v1.30.3 can fail inference for Temporal Envelope checkpoints created by v1.130.0 to v1.30.2. vmanomaly v1.30.4 fixes compatibility issues in all multivariate Temporal Envelope checkpoints. Before a stateful rollout, run `--dryRun` or use the [MCP/skills compatibility checks](https://docs.victoriametrics.com/ai-tools/) to catch issues early. See the [migration guide](https://docs.victoriametrics.com/anomaly-detection/migration/) for details.

**v1.30.4 fixes two regressions in v1.30.3**

- Idle instances stay alive: Shards with no assigned work remain live and observable instead of shutting down. With hot reload enabled, they can pick up newly assigned work without restarting.
- Temporal Envelope checkpoints restored: Multivariate Temporal Envelope checkpoints written by v1.30.0-v1.30.2 load correctly again, so upgrading does not discard trained model state.

**v1.30.3 highlights**

- Less disruptive configuration changes: When you change a shard's configuration, rendezvous sharding keeps unrelated workloads on their current shards as configurations evolve. Changing the shard count can still move work. [Learn more about horizontal scalability](https://docs.victoriametrics.com/anomaly-detection/scaling-vmanomaly/#horizontal-scalability).
- More stable Temporal Envelope baselines: The [Temporal Envelope](https://docs.victoriametrics.com/anomaly-detection/components/models/#temporal-envelope) model produces more reliable baselines when your configured seasonal calendar covers a superset of the patterns actually present in the data. Temporal Envelope handles partially present seasonal patterns more reliably, with more consistent behavior between cold starts and initialization from historical data
- More efficient high-cardinality output: The [VM writer](https://docs.victoriametrics.com/anomaly-detection/components/writer/#vm-writer) now supports `batch_max_series` and `batch_max_bytes` to split large forecast outputs, reducing HTTP and serialization overhead.
- Quicker iterations: Temporal Envelope is now the starting point for new investigations, online models appear first in model lists, and AI-suggested business settings stay in sync with the model wizard and the generated configuration.
- More resilient UI under failure conditions: An invalid ad-hoc request or an unreachable datasource now fails only that specific task instead of shutting down `/vmui`, `/metrics`, and other query-server endpoints.

> **Custom model users**: Automated compatibility checks don't cover custom Python code, dependencies, or serialized state. If you maintain a custom many-to-one model that relies only on `is_multivariate = True`, you now need to explicitly declare `topology = ModelTopology.MANY_TO_ONE`. Built-in multivariate models remain compatible without changes.

## vmanomaly 1.30.2 is out: flexible per-query policies and leaner multivariate models

> If you're running v1.29.1 or newer, you can upgrade directly. For older versions, use the [migration guide](https://docs.victoriametrics.com/anomaly-detection/migration/) and run a [dry run](https://docs.victoriametrics.com/anomaly-detection/migration/#dry-run) first, or let the [skills](https://docs.victoriametrics.com/ai-tools/#agent-skills) and [MCP server](https://docs.victoriametrics.com/ai-tools/#vmanomaly-mcp-server) automate backward-compatibility checks.

- Improved multitenancy in VMUI: [VMUI v1.8.2](https://docs.victoriametrics.com/anomaly-detection/ui/#v182) makes tenant discovery and switching easier when working with multiple tenants.
- Per-query anomaly policies: detection direction and deviation thresholds can now be set per query in the [reader config](https://docs.victoriametrics.com/anomaly-detection/components/reader/#per-query-parameters), so each KPI policy stays consistent across models without duplicating settings ([model-to-query example](https://docs.victoriametrics.com/anomaly-detection/components/reader/#per-query-config-example)).
- New reader concurrency control: the [`reader.workers` setting](https://docs.victoriametrics.com/anomaly-detection/components/reader/#config-parameters) gives you more predictable resource usage during data reads.
- Native-thread controls for workers: the [`settings.native_threads_per_worker` setting](https://docs.victoriametrics.com/anomaly-detection/components/settings/#parallelization) reduces CPU contention, memory pressure, and fit latency.
- Automatic fit-data cleanup: fit data is now cleaned as soon as all dependent models finish and commit. This especially benefits [online-model](https://docs.victoriametrics.com/anomaly-detection/components/models/#online-models) scheduling, where fit data no longer sits around waiting for a distant `fit_every` cycle.
- More efficient grouped processing: disk-based [grouped](https://docs.victoriametrics.com/anomaly-detection/components/models/#group-by) multivariate workloads now require substantially less memory to run. Fitted models now accept reordered input channels as long as the channel set stays the same.
- Per-channel guardrails with Temporal Envelope: the multivariate [Temporal Envelope](https://docs.victoriametrics.com/anomaly-detection/components/models/#temporal-envelope) model can apply direction, data range, and deviation guardrails independently per channel. For example, targeting a GPU throttling pattern where temperature rises, power rises, and clock speed drops.
- Faster fitting with less repeated disk I/O: Temporal Envelope uses less overhead when no seasonality or holiday features are needed. In addition, vmanomaly reuses disk-backed queries across models during fitting.

> **vmoperator compatibility ahead**: the next vmoperator release adds typed `spec.reader.workers` support and recognizes query-level detection direction and deviation guardrails in [VMAnomalyConfig](https://docs.victoriametrics.com/operator/resources/vmanomalyconfig/). Note that `settings.native_threads_per_worker` stays part of the raw vmanomaly configuration and isn't exposed as a typed operator field.

## Community News

In August, Tech Talk is back with an AI-assisted issue troubleshooting demo. This time, Palmersheim and co-founder Roman Khavronenko ran AI agents on VictoriaMetrics, VictoriaLogs, VictoriaTraces, and vmanomaly using the VictoriaMetrics [Agent Skills](https://docs.victoriametrics.com/ai-tools/#agent-skills) and [MCP servers](https://docs.victoriametrics.com/ai-tools/#mcp-servers) to troubleshoot incidents.

<div style="text-align: center;">
<iframe width="560" height="315" src="https://www.youtube.com/embed/VqgM3pNyI-A?si=5rj6Oss16wRegfdu" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
</div>

JetBrains organized a launch party for Go v1.27, and Jesus Espino was invited to give a talk on the Go Select Statement. Check out the recording below if you missed the talk.

<div style="text-align: center;">
<iframe width="560" height="315" src="https://www.youtube.com/embed/fH5hoZJL_tQ?si=VZwIpKCfskZ-de1b" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
</div>

That's it for August. Back with more next month.


