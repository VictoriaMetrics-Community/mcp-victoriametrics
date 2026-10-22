`vmestimator` measures metrics cardinality across arbitrary label dimensions in real time and exposes the results as metrics.

## Why measure?

Consider a setup where metrics are pushed or scraped from hundreds of targets.
One day, a team deploys a new version of their service with a `trace_id` or `user_id` label.
Overnight, that job's [cardinality](https://www.robustperception.io/cardinality-is-key/) explodes from 500 to 500,000 time series.
Suddenly, Time Series DB that stores these metrics starts consuming 100x more memory and disk.
Ingestion slows down, storage struggles to keep up, and in the worst case becomes unavailable.

By the time someone gets paged, the damage is already done: indexes are bloated, caches are oversized, and observability
across the entire system is affected.

`vmestimator` continuously tracks cardinality in real time and exposes the estimation results as [metrics](https://github.com/VictoriaMetrics/vmestimator/blob/main/README.md#cardinality-metrics).
This allows alerting on cardinality spikes within minutes and identifying the offending label directly from the alert.
Instead of discovering the problem after it impacts the infrastructure, it becomes possible to react before it turns into an outage.

`vmestimator` can measure cardinality across arbitrary label dimensions, enabling use cases such as per-tenant usage analysis,
long-term trend tracking, and capacity planning. See more [use cases](https://docs.victoriametrics.com/victoriametrics/vmestimator/#use-cases).

## Key features

* real-time in-flight cardinality tracking without impacting the metrics source or storage
* no vendor-lock: ingest telemetry via [Prometheus RW 1.x](https://prometheus.io/docs/specs/prw/remote_write_spec/) protocol, scrape cardinality estimations in [Prometheus exposition format](https://prometheus.io/docs/instrumenting/exposition_formats/#prometheus-text-format)
* [configurable measurement windows, labels grouping](https://docs.victoriametrics.com/victoriametrics/vmestimator/#configuration)
* low resource usage: can handle millions of samples/s and hundreds of millions unique time series on a single machine
* [horizontally scalable](https://docs.victoriametrics.com/victoriametrics/vmestimator/#cluster)
* ships with a preset set of [alerting rules](https://docs.victoriametrics.com/victoriametrics/vmestimator/#alerting) and [Grafana dashboards](https://docs.victoriametrics.com/victoriametrics/vmestimator/#dashboards)

## Design

`vmestimator` accepts traffic for analysis via [Prometheus Remote-Write v1](https://prometheus.io/docs/specs/prw/remote_write_spec/) protocol. It then computes the cardinality of
the ingested metrics according to the given [configuration](https://docs.victoriametrics.com/victoriametrics/vmestimator/#configuration)
and exposes [cardinality metrics](https://docs.victoriametrics.com/victoriametrics/vmestimator/#cardinality-metrics)
in Prometheus exposition format on `/metrics` endpoint, so they can be [scraped](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#pull-model) by any Prometheus-compatible collector.

![design-1](/victoriametrics/vmestimator/design-1.webp)

`vmestimator` is heavily optimized for high-throughput processing. It approximately requires 1 CPU core for processing
800K metric samples/s, and around 150MiB of memory per each configured [stream](https://docs.victoriametrics.com/victoriametrics/vmestimator/#configuration)
(tracking 3Mil unique time series via 2 streams is using 200MiB on the testing stand).
It is expected that one `vmestimator` instance can easily handle traffic for millions of samples/s and hundreds of millions of unique time series. It is also possible
to [scale vmestimator horizontally](https://docs.victoriametrics.com/victoriametrics/vmestimator/#cluster).

We recommend deploying `vmestimator` close to the metrics source, ideally alongside `vmagent` instances that scrape or forward metrics.
Configure each `vmagent` to [replicate](https://docs.victoriametrics.com/vmagent/index.html#replication-and-high-availability) all metrics into the vmestimator.

Run vmestimator:

```bash
/path/to/vmestimator -config=streams.yaml # -httpListenAddr=:8490
```

Run vmagent:

```bash
/path/to/vmagent \
  -remoteWrite.url=http://127.0.0.1:8428/api/v1/write \             # main remote destination
  -remoteWrite.url=http://127.0.0.1:8490/cardinality/api/v1/write \ # replicate workload to vmestimator
  -remoteWrite.disableOnDiskQueue=false,true \                      # disable queueing for vmestimator
  -remoteWrite.disableMetadata=false,true                           # disable metadata for vmestimator
```

> `-remoteWrite.disableOnDiskQueue` and `-remoteWrite.disableMetadata` reduce resource overhead for vmagent, so it can skip persisting on-disk queue or metadata processing for vmestimator.

When `vmestimator` receives metrics, it starts estimating the cardinality according to the configuration. For example, the following config calculates cardinality per unique metric name across the last 5 minutes:
```yaml
- interval: '5m'
  group_by: ['__name__']
```

The output of this estimation is available on vmestimator's [`/metrics` endpoint](https://docs.victoriametrics.com/victoriametrics/vmestimator/#cardinality-metrics):
```text
cardinality_estimate{interval="5m0s",filter="",group_by_keys="__name__",group_by_values="vm_app_version",by__name__="vm_app_version"} 4
cardinality_estimate{interval="5m0s",filter="",group_by_keys="__name__",group_by_values="process_cpu_seconds_total",by__name__="process_cpu_seconds_total"} 4
...
```

The next step is to collect cardinality estimates and to deliver to the monitoring system, so we can visualize or alert on them.
We can re-use the same `vmagent` instance for this purpose, but it is recommended to run a dedicated pipeline for cardinality estimation. This allows isolating impact: cardinality estimation won't affect the main observability pipeline, and if the main observability pipeline becomes unavailable - cardinality estimations will remain available:

![design-2](/victoriametrics/vmestimator/design-2.webp)

In this architecture, `vmestimator` metrics are isolated from the main observability pipeline, ensuring cardinality visibility remains available even during incidents affecting the primary monitoring system.

## Install

### Start with Docker

Create a `streams.yaml` from [example config](https://github.com/VictoriaMetrics/vmestimator/blob/main/streams.yaml).
Run the Docker image from [Docker Hub](https://hub.docker.com/r/victoriametrics/vmestimator) or [Quay](https://quay.io/repository/victoriametrics/vmestimator), mounting your config file:

```bash
docker run --rm \
  -p 8490:8490 \
  -v /path/to/streams.yaml:/streams.yaml \
  docker.io/victoriametrics/vmestimator:latest \
  -config=/streams.yaml
```

Ingest metrics into `vmestimator` via `http://127.0.0.1:8490/cardinality/api/v1/write` path, observe cardinality estimation metrics at [http://127.0.0.1:8490/metrics](http://127.0.0.1:8490/metrics).

See [Use Cases](https://docs.victoriametrics.com/victoriametrics/vmestimator/#use-cases) for more configuration examples and
[Command-line flags](https://docs.victoriametrics.com/victoriametrics/vmestimator/#command-line-flags) for all available options.

To build from sources, see [How to build from sources](https://docs.victoriametrics.com/victoriametrics/vmestimator/#how-to-build-from-sources).

### Start with Docker Compose

To run vmestimator in docker-compose environment, clone vmestimator repository and run the following command:
```sh
git clone https://github.com/VictoriaMetrics/vmestimator.git && cd vmestimator
make docker-up
```

Visit [http://127.0.0.1:3000](http://127.0.0.1:3000) for Grafana (_login with admin:admin_) to see `vmestimator` dashboards. Or observe cardinality estimation metrics at [http://127.0.0.1:8490/metrics](http://127.0.0.1:8490/metrics).

To stop the environment, run the following command:
```
make docker-down
```

See the full compose file [here](https://github.com/VictoriaMetrics/vmestimator/blob/main/deployment/docker/compose.yml).

### Start with Binary

See [How to build binary from sources](https://github.com/VictoriaMetrics/vmestimator#how-to-build-from-sources) or download pre-built binaries from [GitHub Releases](https://github.com/VictoriaMetrics/vmestimator/releases).

Create a `streams.yaml` from [example config](https://github.com/VictoriaMetrics/vmestimator/blob/main/streams.yaml). Then run `vmestimator` with the following command:
```sh
./vmestimator-prod -config=/path/to/streams.yaml
```

Ingest metrics into `vmestimator` via `http://127.0.0.1:8490/cardinality/api/v1/write` path, observe cardinality estimation metrics at [http://127.0.0.1:8490/metrics](http://127.0.0.1:8490/metrics).

See [Use Cases](https://docs.victoriametrics.com/victoriametrics/vmestimator/#use-cases) for more configuration examples and
[Command-line flags](https://docs.victoriametrics.com/victoriametrics/vmestimator/#command-line-flags) for all available options.

## Configuration

To run `vmestimator` a `streams.yaml` config has to be provided (see [example config](https://github.com/VictoriaMetrics/vmestimator/blob/main/streams.yaml)):

```bash
/path/to/vmestimator -config=streams.yaml # -httpListenAddr=:8490
```

Config reference:

```yaml
streams:
  -
    # The measurement window: how long unique series are retained before the HLL sketch resets.
    # Increases are always reflected immediately. Interval only controls how fast the estimate
    # drops after previously seen series disappear.
    #
    # default: 5m
    interval: <duration>

    # Optional. When set to true, vmestimator emits a cardinality_churn_ratio metric measuring
    # how quickly the series set changes within the measurement window.
    # vmestimator retains a snapshot of the HLL sketch and replaces it every interval.
    # The churn ratio in [0, 1] is computed as:
    #   (2*|A∪B| - |A| - |B|) / |A∪B|
    # where A is the retained snapshot and B is the current sketch.
    # A ratio of 0 means the series set is completely stable; 1 means completely replaced.
    #
    # default: false (churn ratio metric is not emitted)
    calculate_churn: <bool>

    # Optional. MetricsQL label selector used to pre-filter time series before counting.
    # Only series matching all matchers are counted. Supports equality (=), negative equality (!=),
    # regexp (=~), and negative regexp (!~) matchers.
    # Examples:
    #  - '{job="api"}'                       — only series with job="api"
    #  - '{env!="dev"}'                       — exclude dev environment
    #  - '{job=~"api|worker",env!~"dev|staging"}' — regexp and negative regexp combined
    #
    # default: none (all series are counted)
    filter: '<MetricsQL selector>'

    # Optional. Label names used to split the cardinality estimate into per-combination groups.
    # Each distinct combination of values for these labels gets its own estimate metric.
    # Omit entirely for a single global estimate across all series.
    # Examples:
    #  - ["job"]
    #  - ["__name__"]
    #  - ["vm_account_id","vm_project_id"]
    #
    # The special pseudo-label "__label__" may appear as the last element.
    # It estimates the number of unique values per label name instead of a whole time series.
    # Can be combined with explicit keys: ["job", "__label__"] gives unique values per label name, per job.
    # "__label__" may appear at most once and must be the last element.
    #
    # default: none (single global estimate)
    group_by: '[<string>, ...]'

    # Optional. Maximum number of distinct groups (HLL sketches) to track.
    # Once the limit is reached, excess groups are counted in a single shared "rejected" sketch
    # rather than getting their own entry. Acts as a memory cap and a safeguard against OOM
    # when the group_by label values grow unboundedly.
    # Memory upper bound per stream:
    #   group_limit * 2^hll_precision bytes.
    #
    # default: 10000
    group_limit: <integer>

    # Optional. Minimum cardinality estimate to expose for group_by streams.
    # When positive, overrides the global -cardinalityMetrics.minCardinality flag for this stream only.
    # When zero, the global flag value is used.
    # When negative, no threshold is enforced for this stream.
    #
    # default: 0
    min_cardinality: <integer>

    # Optional. Number of shards used to reduce lock contention during parallel ingestion.
    # Slightly increases memory for global streams (no group_by); negligible otherwise.
    # Leave at the default unless you have profiled lock contention or have a specific reason to change it.
    #
    # default: min(64, 2*availableCPUs)
    buckets:  <integer>

    # Optional. HyperLogLog precision p, in range [4..18].
    # Determines the number of registers m = 2^p and the relative error 1.04 / sqrt(m):
    #   p=14 → m=16 384, error ~0.81%, memory ~16 KB per sketch  (default, suits most cases)
    #   p=18 → m=262 144, error ~0.20%, memory ~256 KB per sketch (billing-grade accuracy)
    #   p=10 → m=1 024,   error ~3.25%, memory ~1 KB per sketch   (thousands of groups, memory-tight)
    # See more in https://research.google.com/pubs/archive/40671.pdf
    #
    # default: 14
    hll_precision: <integer>

    # Optional. Whether to use the sparse HyperLogLog representation for low-cardinality groups.
    # Sparse mode uses far less memory until a group's cardinality reaches ~2^(p-1),
    # at which point it automatically promotes to the dense representation.
    # See more in # See more in https://research.google.com/pubs/archive/40671.pdf
    #
    # default: true
    hll_sparse: <bool>

    # Optional. Static labels attached to every output metric produced by this stream entry.
    # Useful when multiple vmestimator instances feed the same storage and you need
    # to distinguish their estimates in dashboards and alerts.
    labels:
      <labelname>: <labelvalue> ...
```

### Config reload

`vmestimator` reloads the `-config` file without restart on:

* `SIGHUP` signal, e.g. `kill -HUP $(pidof vmestimator)`.
* HTTP request to `/-/reload`. Protect it with `-reloadAuthKey`, passed via the `authKey` query arg.
* Every `-configCheckInterval`, if the flag is set.

If the file content is unchanged, nothing happens. Otherwise the new config is applied stream by stream:

* Streams with unchanged settings keep their accumulated estimates and churn state. The order of streams in the file doesn't matter.
* Added streams and streams with any changed setting (e.g. `group_limit` or `hll_precision`) start with empty estimates.
* Removed streams are stopped and their estimates are no longer exposed.

If any stream is added or changed, the deduplication state is reset, so the new streams receive all the series.
If the new config is invalid, the error is logged and the previous config stays active.

Reload status is exposed via `vmestimator_config_reloads_total`, `vmestimator_config_reloads_errors_total`,
`vmestimator_config_last_reload_successful` and `vmestimator_config_last_reload_success_timestamp_seconds` metrics.
Changes in command-line flags need a restart.

## Cardinality Metrics

Cardinality estimates are exposed as the `cardinality_estimate` metric.
All metrics include `interval`, `filter`, `group_by_keys`, `group_by_values`, and any static labels defined in the stream config.

For global estimates (no `group_by` configured), `group_by_keys` is `__global__` and `group_by_values` is omitted:

```
cardinality_estimate{interval="1h0m0s",filter="",group_by_keys="__global__"} 142300
```

For grouped estimates, one summary line shows the total number of distinct groups `group_by_keys="__group__"`, followed by one line per distinct label value combination.
Each per-group line also includes individual `by_{key}="{val}"` labels:

```
cardinality_estimate{interval="5m0s",filter="",group_by_keys="__group__",group_by_values="instance,job"} 2
cardinality_estimate{interval="5m0s",filter="",group_by_keys="instance,job",group_by_values="host1:9090,prometheus",by_instance="host1:9090",by_job="prometheus"} 312
cardinality_estimate{interval="5m0s",filter="",group_by_keys="instance,job",group_by_values="host2:9100,node",by_instance="host2:9100",by_job="node"} 87
```

For `__label__` streams, the group value is the label name and `by__label__` holds it:

```
cardinality_estimate{interval="5m0s",filter="",group_by_keys="__group__",group_by_values="__label__"} 3
cardinality_estimate{interval="5m0s",filter="",group_by_keys="__label__",group_by_values="__name__",by__label__="__name__"} 450
cardinality_estimate{interval="5m0s",filter="",group_by_keys="__label__",group_by_values="job",by__label__="job"} 12
cardinality_estimate{interval="5m0s",filter="",group_by_keys="__label__",group_by_values="instance",by__label__="instance"} 87
```

> Note: the total distinct group count in the summary line may exceed the number of per-group lines when `group_limit` is reached
and excess groups are counted in a single shared "rejected" sketch rather than getting their own entry.

By default, cardinality estimates are merged with the estimator's operational metrics and exposed at `/metrics`.
This is controlled by the `-cardinalityMetrics.exposeAt` flag:

- `-cardinalityMetrics.exposeAt=/metrics` (default): cardinality metrics merged with operational metrics at `/metrics`
- `-cardinalityMetrics.exposeAt=/cardinality/metrics`: cardinality metrics exposed at separate path
- `-cardinalityMetrics.exposeAt=`: cardinality metrics not exposed via HTTP

Use `-cardinalityMetrics.minCardinality=N` to suppress group-by cardinality estimates below `N`.
This is useful when `group_by` produces many low-cardinality groups that add noise without actionable signal.
When groups are suppressed, the `vmestimator_cardinality_estimates_dropped` metric reports the number of dropped groups per stream.

The global flag applies the same threshold to every stream. To override it per stream — for example to
keep low-cardinality estimates from a `__label__` stream while suppressing noise from high-cardinality
streams — set the `min_cardinality` field on individual streams:

```yaml
streams:
  - group_by: [__name__, job, region]
    interval: 5m
    min_cardinality: 1000

  - group_by: [__label__]
    interval: 5m
    # min_cardinality unset → falls back to the global flag
```

When `min_cardinality` is set to a positive value on a stream, it overrides the global flag for that stream only.
When unset or `0`, the global flag value is used.
When negative, no threshold is enforced for that stream, even if the global flag is set.
In cluster mode, `min_cardinality` is taken from the storage nodes' stream configs; if they differ, the select node uses the largest value.

Computing cardinality estimates is expensive, so results are cached.
Cache duration is controlled by `-cardinalityMetrics.cacheTTL` (default: `30s`).
Set to `0` to disable caching entirely.

## Use Cases

### Basic

Global cardinality:

```yaml
# streams.yaml

- interval: '5m'
```

Per metric name cardinality:

```yaml
# streams.yaml

- interval: '5m'
  group_by: ['__name__']
```

Per job label cardinality:

```yaml
# streams.yaml

- interval: '5m'
  group_by: ['job']
```

Per tenant cardinality:

```yaml
# streams.yaml

- interval: '5m'
  group_by: ['vm_account_id', 'vm_project_id']
```

Cardinality scoped to a single job:

```yaml
# streams.yaml

- interval: '5m'
  filter: '{job="api"}'
  group_by: ['__name__']
```

Cardinality excluding non-production environments:

```yaml
# streams.yaml

- interval: '5m'
  filter: '{env!~"dev|staging"}'
```

Per label unique values:

```yaml
# streams.yaml

- interval: '5m'
  group_by: ['__label__']
```

This estimates the number of unique values for every label name seen in the ingested traffic.
Useful for identifying high-cardinality labels (e.g. `trace_id`, `user_id`) before they cause storage problems.
Can be narrowed to a specific job: `['job', '__label__']`.

### Churn calculation

[Churn ratio](https://valyala.medium.com/prometheus-storage-technical-terms-for-humans-4ab4de6c3d48#churn-rate) measures how quickly time series are created and disappear.
[High churn](https://docs.victoriametrics.com/victoriametrics/faq/#what-is-high-churn-rate) means many series appear briefly and are replaced by new ones.
This puts pressure on storage, because each new series must be indexed regardless of how short its lifetime is.

To measure churn, set `calculate_churn: true` on a stream:

```yaml
# streams.yaml

- interval: '15m'
  calculate_churn: true
  group_by: ['job']
```

vmestimator will emit a `cardinality_churn_ratio` metric per group, comparing the series set against a snapshot taken one `interval` earlier. The value is in `[0, 1]`:
- `0` — the series set is completely stable; the same series were active across both windows.
- `1` — complete churn; entirely different series appeared in the latest window.

```
cardinality_churn_ratio{interval="15m0s",filter="",group_by_keys="job",group_by_values="api",by_job="api"} 0.1200
```

This helps identify jobs that create the most indexing pressure on storage, even when their current active cardinality appears moderate.

### Alerting

Pre-built alert rules for cardinality monitoring are available in
[deployment/docker/rules/alerts-cardinality.yml](https://github.com/VictoriaMetrics/vmestimator/blob/main/deployment/docker/rules/alerts-cardinality.yml).

They require the following stream configuration to also support churn detection:

```yaml
# streams.yaml
# or use example config:
# https://github.com/VictoriaMetrics/vmestimator/blob/main/streams.yaml

# Global stream — required for GlobalChurnTooHigh and GlobalCardinalityTooHigh.
- interval: '15m'
  calculate_churn: true

# Per-job stream — required for JobChurnTooHigh and JobCardinalityTooHigh.
- interval: '15m'
  calculate_churn: true
  group_by: ['job']
```

The included alerts are:

- **GlobalChurnTooHigh** — fires when global series churn exceeds 10% for 15 minutes.
  High global churn can cause unexpected OOM or slow queries across the whole database.

- **JobChurnTooHigh** — fires when more than 20% of a job's series churned for 30 minutes.
  Catches jobs that generate continuous indexing pressure even when their active series count looks moderate.

- **GlobalCardinalityTooHigh** — fires when global cardinality exceeds the weekly mean by more than 3 standard deviations
  and has grown at least 20% above that mean, sustained for 30 minutes. The guard baseline is 100,000 series.

- **JobCardinalityTooHigh** — fires when a job's cardinality exceeds its weekly mean by more than 3 standard deviations
  and has grown at least 20% above that mean, sustained for 30 minutes. The guard baseline is 10,000 series per job.
  Uses 3-sigma anomaly detection rather than a fixed threshold, so it adapts to the expected scale of each job.

- **CardinalityGroupLimitNearlyReached** — fires when the number of tracked groups exceeds 80% of the configured `group_limit`.
  Acts as an early warning that some label value combinations may soon be dropped from individual tracking.

- **CardinalityGroupLimitReached** — fires when groups are actively rejected because `group_limit` is full.
  At this point, some label combinations are being counted in a shared "rejected" sketch rather than tracked individually.

All alerts link to the [Cardinality Explorer dashboard](https://play-grafana.victoriametrics.com/d/mktd5h8/).

## Alternative solutions

### PromQL

Cardinality can be estimated with PromQL.

Global cardinality:

```
count({__name__=~".*"})
```

Top ten metric names by cardinality:

```
topk(10, count({__name__=~".*"}) by (__name__))
```

Top ten jobs by cardinality:

```
topk(10, count({__name__=~".*"}) by (job))
```

This approach works for small setups but does not scale well, because these queries scan the entire time series set.
Most critically, if the storage is overloaded or unavailable, these queries could not be executed.

### Cardinality Explorer

VictoriaMetrics includes a built-in [cardinality explorer](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#cardinality-explorer).
It provides per-metric detail beyond raw series counts: query frequency, last access time, day-over-day change, and share of total cardinality.
It is well suited for in-depth, ad-hoc investigation.
For example, finding metrics that are high-cardinality but rarely queried,
so they can be [dropped via relabeling](https://docs.victoriametrics.com/victoriametrics/relabeling/#how-to-drop-metrics-during-scrape) or reduce cardinality with [stream aggregation](https://docs.victoriametrics.com/victoriametrics/stream-aggregation/).

Both tools serve different purposes and work well together.
Use `vmestimator` for continuous monitoring, alerting, and cross-cluster cardinality tracking.
Use the cardinality explorer when you need to drill into a specific metric or label and understand what is driving its cardinality.

## Cluster

`vmestimator` supports a clustered deployment for high availability or when CPU on a single instance becomes a limiting factor.

Instances are split into two roles: **storage nodes** accept Prometheus remote write and maintain local HyperLogLog sketches; **selector nodes** query all storage nodes, merge their sketches, and expose a unified cardinality estimate. Cardinality estimate results should be scraped from selector nodes.

![cluster](/victoriametrics/vmestimator/cluster.webp)

**Storage nodes:**

```
vmestimator -config=streams.yaml -httpListenAddr=:8491 -cardinalityMetrics.exposeAt=/cardinality/metrics
vmestimator -config=streams.yaml -httpListenAddr=:8492 -cardinalityMetrics.exposeAt=/cardinality/metrics
vmestimator -config=streams.yaml -httpListenAddr=:8493 -cardinalityMetrics.exposeAt=/cardinality/metrics
```

**Selector nodes:**

```
vmestimator -storageNode=http://vmestimator-storage-1:8491 \
            -storageNode=http://vmestimator-storage-2:8492 \
            -storageNode=http://vmestimator-storage-3:8493 \
            -httpListenAddr=:8490
```

Setting `-cardinalityMetrics.exposeAt=/cardinality/metrics` on storage nodes keeps per-node estimates off the default `/metrics` path. The `/metrics` endpoint then returns only operational metrics, while `/cardinality/metrics` exposes the node's local estimate — useful for inspecting or debugging a specific node.

A selector with `-storageNode` flags and no `-config` runs without local estimators and only merges remote data.

When multiple selector nodes are scraped, each returns a fully merged estimate.
Deduplicate at query time to avoid overcounting:

```
max(cardinality_estimate) without (instance)
```

## Operational metrics

Each stream with a non-empty `group_by` exposes the following operational metrics at `/metrics`:

* `vmestimator_estimator_group_size{interval, filter, group_by_keys}` — the current number of unique `group_by_keys` held by the group.
* `vmestimator_estimator_group_rejected_size{interval, filter, group_by_keys}` — the number of unique `group_by_keys` rejected after the group reached its `group_limit`.
* `vmestimator_estimator_group_limit{interval, filter, group_by_keys}` — the maximum number of unique `group_by_keys` the group can hold, as configured by `group_limit`.

Additionally, every stream (including non-grouped ones) exposes:

- `vmestimator_estimator_insert_total{group_by_keys, interval, filter}` — total number of samples inserted into this stream's estimator (only counts series that passed the `filter`)

## Deduplication

Processing each time series is CPU-intensive: `vmestimator` must iterate over all labels, resolve group-by values, and call HLL insert for every estimator.
For `__label__` streams it is even more expensive, as HLL insert is called once per label per series.

Scraped time series are mostly identical between scrapes on short intervals.
Once a series has been seen, re-processing it within the same window does not change the cardinality estimate.
Deduplication exploits this by dropping already-seen series before they reach the estimators, reducing CPU usage without affecting accuracy.

Enable deduplication with `-deduplication.interval`:

```bash
/path/to/vmestimator -config=streams.yaml -deduplication.interval=2m
```

When enabled, `vmestimator` tracks which time series it has already seen within the configured window using bloom filters.
Series seen within the window are dropped before reaching the estimators.
Set the interval 3–5× lower than the shortest estimator interval to ensure estimates remain accurate.

**Memory budget**

By default, the deduplicator uses 5% of the allowed memory (see `-memory.allowedBytes` / `-memory.allowedPercent`).
To set an explicit limit use `-deduplication.maxSizeBytes` (e.g. `100MB`) or `-deduplication.maxSize` (number of unique series).
When both flags are set, the stricter limit applies.

**Overflow behaviour**

When the true number of unique series exceeds 80% of the configured limit, the deduplicator gradually switches to pass-through mode for the excess series.
This prevents bloom filter saturation (which would cause false positives) and keeps memory usage bounded.
The pass-through ratio is logged whenever it changes:

A ratio of `0/1000` means full deduplication; `1000/1000` means all series pass through unchanged.

## Dashboards

Two Grafana dashboards are available in the [dashboards](https://github.com/VictoriaMetrics/vmestimator/tree/main/dashboards) directory:

- [VictoriaMetrics - vmestimator](https://play-grafana.victoriametrics.com/d/mkv22l4/victoriametrics-vmestimator) — application health: CPU, memory, ingestion rates, concurrent inserts, and group key saturation.
  ![dashboard-vmestimator](/victoriametrics/vmestimator/dashboard-vmestimator.webp)

- [VictoriaMetrics - Cardinality Estimations](https://play-grafana.victoriametrics.com/d/mktd5g9/victoriametrics-cardinality-estimations) — cardinality estimations: shows estimations, churn over configured label dimensions.
  ![dashboard-cardinality](/victoriametrics/vmestimator/dashboard-cardinality.webp)

## Limitations

For performance reasons, `vmestimator` doesn't decode [metric samples](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#raw-samples) in received Remote Write requests. It only parses [labels](https://docs.victoriametrics.com/victoriametrics/keyconcepts/#labels). This optimization introduces the following limitations:
1. Cardinality is estimated based on time of arrival of the request, not the actual timestamps of the samples within the request. If `vmestimator` receives backfilling traffic with month old timestamps, the cardinality estimation will be estimated as of now.
2. [Staleness markers](https://prometheus.io/docs/prometheus/latest/querying/basics/#staleness) aren't decoded. A time series carrying a staleness marker is counted like any other received time series.
3. Consistent order of labels in time series matters. `vmestimator` expects that clients already send time series with deterministic orders of labels, so it skips sorting for performance reasons. If time series arrive with unsorted order of labels within time series, then cardinality estimations can become inflated. For example, `metric{foo=bar, baz=qux}` and `metric{baz=qux, foo=bar}` are considered as two distinct time series.

## How to build from sources

It is recommended to use the [docker images](https://hub.docker.com/r/victoriametrics/vmestimator).

Development build:

1. [Install Go](https://golang.org/doc/install).
1. Run `make vmestimator` from the root folder of [the repository](https://github.com/VictoriaMetrics/vmestimator).
   It builds `vmestimator` binary and places it into the `bin` folder.

Production build:

1. [Install docker](https://docs.docker.com/install/).
1. Run `make vmestimator-prod` from the root folder of [the repository](https://github.com/VictoriaMetrics/vmestimator).
   It builds `vmestimator-prod` binary and puts it into the `bin` folder.

Building docker images:

Run `make package-vmestimator`. It builds `victoriametrics/vmestimator:<PKG_TAG>` docker image locally.
`<PKG_TAG>` is auto-generated image tag, which depends on source code in the repository.
The `<PKG_TAG>` may be manually set via `PKG_TAG=foobar make package-vmestimator`.

The base docker image is [alpine](https://hub.docker.com/_/alpine) but it is possible to use any other base image by setting it via `<ROOT_IMAGE>` environment variable.
For example, the following command builds the image on top of [scratch](https://hub.docker.com/_/scratch) image:

```sh
ROOT_IMAGE=scratch make package-vmrestore
```

You can build and publish to your own registry and namespace:

```
DOCKER_REGISTRIES=ghcr.io DOCKER_NAMESPACE=foo make publish-vmagent
```

## Command-line flags

Run `vmestimate -help` in order to see all the available options:

```
Usage of ./bin/vmestimator:
  -cardinalityMetrics.cacheTTL duration
        Duration for caching cardinality metrics response (default 30s)
  -cardinalityMetrics.exposeAt string
        HTTP path for exposing cardinality metrics. If set to the default /metrics, cardinality metrics are merged with regular metrics and exposed together. If set to a different path, only cardinality metrics are exposed at that endpoint. If set to an empty value, cardinality metrics are not exposed via HTTP at all. (default "/metrics")
  -cardinalityMetrics.minCardinality uint
        Minimum cardinality estimate to expose. Estimates below this threshold are suppressed to reduce metric volume and noise. Set to 0 to expose all estimates (default 0).
  -config string
        Path to YAML configuration file. Must be set unless -storageNode is specified. See https://github.com/VictoriaMetrics/vmestimator/blob/main/streams.yaml for config example
  -configCheckInterval duration
        Interval for checking for changes in '-config' file. By default, the checking is disabled. Send SIGHUP signal or request /-/reload in order to force config check for changes
  -deduplication.interval duration
        Time window for deduplication. When set, time series already seen within the window are dropped before being forwarded to estimators. Disabled when set to 0.
  -deduplication.maxSize int
        Maximum number of unique time series the deduplicator tracks. Pass-through begins at 80% of this limit so bloom filters never saturate and start producing false positives. Defaults to a value derived from 5 percent of allowed memory. When both -deduplication.maxSizeBytes and -deduplication.maxSize are set, the stricter limit is used.
  -deduplication.maxSizeBytes size
        Memory budget for deduplication bloom filters, e.g. 100MB. Controls how many unique time series can be tracked before the deduplicator gradually switches to pass-through mode. Pass-through starts at 80% of the limit to keep bloom filters well below saturation. Defaults to 5 percent of allowed memory. When both -deduplication.maxSizeBytes and -deduplication.maxSize are set, the stricter limit is used.
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 0)
  -enableTCP6
        Whether to enable IPv6 for listening and dialing. By default, only IPv4 TCP and UDP are used
  -envflag.enable
        Whether to enable reading flags from environment variables in addition to the command line. Command line flag values have priority over values from environment vars. Flags are read only from the command line if this flag isn't set. See https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#environment-variables for more details
  -envflag.prefix string
        Prefix for environment variables if -envflag.enable is set
  -filestream.disableFadvise
        Whether to disable fadvise() syscall when reading large data files. The fadvise() syscall prevents from eviction of recently accessed data from OS page cache during background merges and backups. In some rare cases it is better to disable the syscall if it uses too much CPU
  -flagsAuthKey value
        Auth key for /flags endpoint. It must be passed via authKey query arg. It overrides -httpAuth.*
        Flag value can be read from the given file when using -flagsAuthKey=file:///abs/path/to/file or -flagsAuthKey=file://./relative/path/to/file.
        Flag value can be read from the given http/https url when using -flagsAuthKey=http://host/path or -flagsAuthKey=https://host/path
  -fs.maxConcurrency int
        The maximum number of concurrent goroutines to work with files; smaller values may help reducing Go scheduling latency on systems with small number of CPU cores; higher values may help reducing data ingestion latency on systems with high-latency storage such as NFS or Ceph (default 160)
  -http.connTimeout duration
        Incoming connections to -httpListenAddr are closed after the configured timeout. This may help evenly spreading load among a cluster of services behind TCP-level load balancer. Zero value disables closing of incoming connections (default 2m0s)
  -http.disableCORS
        Disable CORS for all origins (*)
  -http.disableKeepAlive
        Whether to disable HTTP keep-alive for incoming connections at -httpListenAddr
  -http.disableResponseCompression
        Disable compression of HTTP responses to save CPU resources. By default, compression is enabled to save network bandwidth
  -http.header.csp string
        Value for 'Content-Security-Policy' header, recommended: "default-src 'self'"
  -http.header.disableServerHostname
        Whether to disable 'X-Server-Hostname' header in HTTP responses
  -http.header.frameOptions string
        Value for 'X-Frame-Options' header
  -http.header.hsts string
        Value for 'Strict-Transport-Security' header, recommended: 'max-age=31536000; includeSubDomains'
  -http.idleConnTimeout duration
        Timeout for incoming idle http connections (default 1m0s)
  -http.maxGracefulShutdownDuration duration
        The maximum duration for a graceful shutdown of the HTTP server. During this period the server stops accepting new connections, but it will continue serving existing connections. The remaining in-flight requests are canceled before the deadline, so the shutdown can finish within this duration. A highly loaded server may require increased value for a graceful shutdown (default 7s)
  -http.pathPrefix string
        An optional prefix to add to all the paths handled by http server. For example, if '-http.pathPrefix=/foo/bar' is set, then all the http requests will be handled on '/foo/bar/*' paths. This may be useful for proxied requests. See https://www.robustperception.io/using-external-urls-and-proxies-with-prometheus
  -http.shutdownDelay duration
        Optional delay before http server shutdown. During this delay, the server returns non-OK responses from /health page, so load balancers can route new requests to other servers
  -httpAuth.password value
        Password for HTTP server's Basic Auth. The authentication is disabled if -httpAuth.username is empty
        Flag value can be read from the given file when using -httpAuth.password=file:///abs/path/to/file or -httpAuth.password=file://./relative/path/to/file.
        Flag value can be read from the given http/https url when using -httpAuth.password=http://host/path or -httpAuth.password=https://host/path
  -httpAuth.username string
        Username for HTTP server's Basic Auth. The authentication is disabled if empty. See also -httpAuth.password
  -httpListenAddr array
        TCP address to listen for incoming HTTP requests
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -insert.maxQueueDuration duration
        The maximum duration to wait in the queue when -maxConcurrentInserts concurrent insert requests are executed (default 1m0s)
  -internStringCacheExpireDuration duration
        The expiry duration for caches for interned strings. See https://en.wikipedia.org/wiki/String_interning . See also -internStringMaxLen and -internStringDisableCache (default 6m0s)
  -internStringDisableCache
        Whether to disable caches for interned strings. This may reduce memory usage at the cost of higher CPU usage. See https://en.wikipedia.org/wiki/String_interning . See also -internStringCacheExpireDuration and -internStringMaxLen
  -internStringMaxLen int
        The maximum length for strings to intern. A lower limit may save memory at the cost of higher CPU usage. See https://en.wikipedia.org/wiki/String_interning . See also -internStringDisableCache and -internStringCacheExpireDuration (default 500)
  -loggerDisableTimestamps
        Whether to disable writing timestamps in logs
  -loggerErrorsPerSecondLimit int
        Per-second limit on the number of ERROR messages. If more than the given number of errors are emitted per second, the remaining errors are suppressed. Zero values disable the rate limit
  -loggerFormat string
        Format for logs. Possible values: default, json (default "default")
  -loggerJSONFields string
        Allows renaming fields in JSON formatted logs. Example: "ts:timestamp,msg:message" renames "ts" to "timestamp" and "msg" to "message". Supported fields: ts, level, caller, msg
  -loggerLevel string
        Minimum level of errors to log. Possible values: INFO, WARN, ERROR, FATAL, PANIC (default "INFO")
  -loggerMaxArgLen int
        The maximum length of a single logged argument. Longer arguments are replaced with 'arg_start..arg_end', where 'arg_start' and 'arg_end' is prefix and suffix of the arg with the length not exceeding -loggerMaxArgLen / 2 (default 5000)
  -loggerOutput string
        Output for the logs. Supported values: stderr, stdout (default "stderr")
  -loggerTimezone string
        Timezone to use for timestamps in logs. Timezone must be a valid IANA Time Zone. For example: America/New_York, Europe/Berlin, Etc/GMT+3 or Local (default "UTC")
  -loggerWarnsPerSecondLimit int
        Per-second limit on the number of WARN messages. If more than the given number of warns are emitted per second, then the remaining warns are suppressed. Zero values disable the rate limit
  -maxConcurrentInserts int
        The maximum number of concurrent insert requests. Set higher value when clients send data over slow networks. Default value depends on the number of available CPU cores. It should work fine in most cases since it minimizes resource usage. See also -insert.maxQueueDuration (default 20)
  -maxInsertRequestSize size
        The maximum size in bytes of a single Prometheus remote_write API request
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 33554432)
  -memory.allowedBytes size
        Allowed size of system memory VictoriaMetrics caches may occupy. This option overrides -memory.allowedPercent if set to a non-zero value. Too low a value may increase the cache miss rate usually resulting in higher CPU and disk IO usage. Too high a value may evict too much data from the OS page cache resulting in higher disk IO usage. The process may behave unexpectedly if this flag is set too small (e.g., 1 byte).
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 0)
  -memory.allowedPercent float
        Allowed percent of system memory VictoriaMetrics caches may occupy. See also -memory.allowedBytes. Too low a value may increase cache miss rate usually resulting in higher CPU and disk IO usage. Too high a value may evict too much data from the OS page cache which will result in higher disk IO usage (default 60)
  -metrics.exposeMetadata
        Whether to expose TYPE and HELP metadata at the /metrics page, which is exposed at -httpListenAddr . The metadata may be needed when the /metrics page is consumed by systems, which require this information. For example, Managed Prometheus in Google Cloud - https://cloud.google.com/stackdriver/docs/managed-prometheus/troubleshooting#missing-metric-type
  -metricsAuthKey value
        Auth key for /metrics endpoint. It must be passed via authKey query arg. It overrides -httpAuth.*
        Flag value can be read from the given file when using -metricsAuthKey=file:///abs/path/to/file or -metricsAuthKey=file://./relative/path/to/file.
        Flag value can be read from the given http/https url when using -metricsAuthKey=http://host/path or -metricsAuthKey=https://host/path
  -pprofAuthKey value
        Auth key for /debug/pprof/* endpoints. It must be passed via authKey query arg. It overrides -httpAuth.*
        Flag value can be read from the given file when using -pprofAuthKey=file:///abs/path/to/file or -pprofAuthKey=file://./relative/path/to/file.
        Flag value can be read from the given http/https url when using -pprofAuthKey=http://host/path or -pprofAuthKey=https://host/path
  -pushmetrics.disableCompression
        Whether to disable request body compression when pushing metrics to every -pushmetrics.url
  -pushmetrics.extraLabel array
        Optional labels to add to metrics pushed to every -pushmetrics.url . For example, -pushmetrics.extraLabel='instance="foo"' adds instance="foo" label to all the metrics pushed to every -pushmetrics.url
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -pushmetrics.header array
        Optional HTTP request header to send to every -pushmetrics.url . For example, -pushmetrics.header='Authorization: Basic foobar' adds 'Authorization: Basic foobar' header to every request to every -pushmetrics.url
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -pushmetrics.interval duration
        Interval for pushing metrics to every -pushmetrics.url (default 10s)
  -pushmetrics.url array
        Optional URL to push metrics exposed at /metrics page. See https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#push-metrics . By default, metrics exposed at /metrics page aren't pushed to any remote storage
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -reloadAuthKey value
        Auth key for /-/reload http endpoint. It must be passed via authKey query arg. It overrides -httpAuth.*
        Flag value can be read from the given file when using -reloadAuthKey=file:///abs/path/to/file or -reloadAuthKey=file://./relative/path/to/file.
        Flag value can be read from the given http/https url when using -reloadAuthKey=http://host/path or -reloadAuthKey=https://host/path
  -secret.flags array
        Comma-separated list of flag names with secret values. Values for these flags are hidden in logs and on /metrics page
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode array
        HTTP URLs of remote vmestimator nodes to query for cardinality snapshots, e.g. http://vmestimator-2:8490
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tls array
        Whether to enable TLS for incoming HTTP requests at the given -httpListenAddr (aka https). -tlsCertFile and -tlsKeyFile must be set if -tls is set. See also -mtls
        Supports array of values separated by comma or specified via multiple flags.
        Empty values are set to false.
  -tlsCertFile array
        Path to file with TLS certificate for the corresponding -httpListenAddr if -tls is set. Prefer ECDSA certs instead of RSA certs as RSA certs are slower. The provided certificate file is automatically re-read every second, so it can be dynamically updated. See also -tlsAutocertHosts
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tlsCipherSuites array
        Optional list of TLS cipher suites for incoming requests over HTTPS if -tls is set. See the list of supported cipher suites at https://pkg.go.dev/crypto/tls#pkg-constants
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tlsKeyFile array
        Path to file with TLS key for the corresponding -httpListenAddr if -tls is set. The provided key file is automatically re-read every second, so it can be dynamically updated. See also -tlsAutocertHosts
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tlsMinVersion array
        Optional minimum TLS version to use for the corresponding -httpListenAddr if -tls is set. Supported values: TLS10, TLS11, TLS12, TLS13
        Supports an array of values separated by comma or specified via multiple flags.
        Each array item can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -version
        Show VictoriaMetrics version
  -workers int
        The number of workers for processing time series insertions concurrently. Each worker handles insertMany calls for a single estimator. Defaults to 2x the number of available CPUs. (default 20)
```
