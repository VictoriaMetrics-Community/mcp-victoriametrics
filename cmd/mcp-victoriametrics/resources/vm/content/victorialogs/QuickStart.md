---
weight: 1
title: Quick Start
menu:
  docs:
    parent: victorialogs
    identifier: vl-quick-start
    weight: 1
    title: Quick Start
tags:
  - logs
  - guide
aliases:
- /victorialogs/QuickStart.html
- /victorialogs/quick-start.html
- /victorialogs/quick-start/
---
It is recommended to read [README](https://docs.victoriametrics.com/victorialogs/)
and [Key Concepts](https://docs.victoriametrics.com/victorialogs/keyconcepts/)
before you start working with VictoriaLogs.

## How to install and run VictoriaLogs

There are the following options exist:

- [To run pre-built binaries](#pre-built-binaries)
- [To run Docker image](#docker-image)
- [To run in Kubernetes with Helm charts](#helm-charts)
- [To build VictoriaLogs from source code](#building-from-source-code)

### Pre-built binaries

Pre-built binaries for VictoriaLogs are available at the [releases](https://github.com/VictoriaMetrics/VictoriaLogs/releases/) page.
Just download archive for the needed Operating system and architecture, unpack it and run `victoria-logs-prod` from it.

For example, the following commands download VictoriaLogs archive for Linux/amd64, unpack and run it:

```sh
curl -L -O https://github.com/VictoriaMetrics/VictoriaLogs/releases/download/v1.26.0/victoria-logs-linux-amd64-v1.26.0.tar.gz
tar xzf victoria-logs-linux-amd64-v1.26.0.tar.gz
./victoria-logs-prod -storageDataPath=victoria-logs-data
```

VictoriaLogs is ready for [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/)
and [querying](https://docs.victoriametrics.com/victorialogs/querying/) at the TCP port `9428` now!
It has no any external dependencies, so it may run in various environments without additional setup and configuration.
VictoriaLogs automatically adapts to the available CPU and RAM resources. It also automatically setups and creates
the needed indexes during [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/).

See also:

- [How to configure VictoriaLogs](#how-to-configure-victorialogs)
- [How to ingest logs into VictoriaLogs](https://docs.victoriametrics.com/victorialogs/data-ingestion/)
- [How to query VictoriaLogs](https://docs.victoriametrics.com/victorialogs/querying/)

### Docker image

You can run VictoriaLogs in a Docker container. It is the easiest way to start using VictoriaLogs.
Here is the command to run VictoriaLogs in a Docker container:

```sh
docker run --rm -it -p 9428:9428 -v ./victoria-logs-data:/victoria-logs-data \
  docker.io/victoriametrics/victoria-logs:v1.26.0 -storageDataPath=victoria-logs-data
```

See also:

- [How to configure VictoriaLogs](#how-to-configure-victorialogs)
- [How to ingest logs into VictoriaLogs](https://docs.victoriametrics.com/victorialogs/data-ingestion/)
- [How to query VictoriaLogs](https://docs.victoriametrics.com/victorialogs/querying/)

### Helm charts

You can run VictoriaLogs in Kubernetes environment
with [VictoriaLogs single](https://docs.victoriametrics.com/helm/victorialogs-single/)
or [cluster](https://docs.victoriametrics.com/helm/victorialogs-cluster) helm charts.

### Building from source code

Follow the following steps in order to build VictoriaLogs from source code:

- Checkout VictoriaLogs source code:

  ```sh
  git clone https://github.com/VictoriaMetrics/VictoriaLogs
  cd VictoriaLogs
  ```

- Build VictoriaLogs:

  ```sh
  make victoria-logs
  ```

- Run the built binary:

  ```sh
  bin/victoria-logs -storageDataPath=victoria-logs-data
  ```

VictoriaLogs is ready for [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/)
and [querying](https://docs.victoriametrics.com/victorialogs/querying/) at the TCP port `9428` now!
It has no any external dependencies, so it may run in various environments without additional setup and configuration.
VictoriaLogs automatically adapts to the available CPU and RAM resources. It also automatically setups and creates
the needed indexes during [data ingestion](https://docs.victoriametrics.com/victorialogs/data-ingestion/).

See also:

- [How to configure VictoriaLogs](#how-to-configure-victorialogs)
- [How to ingest logs into VictoriaLogs](https://docs.victoriametrics.com/victorialogs/data-ingestion/)
- [How to query VictoriaLogs](https://docs.victoriametrics.com/victorialogs/querying/)

## How to configure VictoriaLogs

VictoriaLogs is configured via command-line flags. All the command-line flags have sane defaults,
so there is no need in tuning them in general case. VictoriaLogs runs smoothly in most environments
without additional configuration.

Pass `-help` to VictoriaLogs in order to see the list of supported command-line flags with their description and default values:

```sh
/path/to/victoria-logs -help
```

VictoriaLogs stores the ingested data to the `victoria-logs-data` directory by default. The directory can be changed
via `-storageDataPath` command-line flag. See [these docs](https://docs.victoriametrics.com/victorialogs/#storage) for details.

By default, VictoriaLogs stores [log entries](https://docs.victoriametrics.com/victorialogs/keyconcepts/) with timestamps
in the time range `[now-7d, now]`, while dropping logs outside the given time range.
E.g. it uses the retention of 7 days. Read [these docs](https://docs.victoriametrics.com/victorialogs/#retention) on how to control the retention
for the [ingested](https://docs.victoriametrics.com/victorialogs/data-ingestion/) logs.

It is recommended setting up monitoring of VictoriaLogs according to [these docs](https://docs.victoriametrics.com/victorialogs/#monitoring).

See also:

- [How to ingest logs into VictoriaLogs](https://docs.victoriametrics.com/victorialogs/data-ingestion/)
- [How to query VictoriaLogs](https://docs.victoriametrics.com/victorialogs/querying/)

## Docker demos

Docker-compose demos for single-node and cluster version of VictoriaLogs that include logs collection,
monitoring, alerting and Grafana are available [here](https://github.com/VictoriaMetrics/VictoriaLogs/tree/master/deployment/docker#readme).

Docker-compose demos that integrate VictoriaLogs and various log collectors:

- [Filebeat demo](https://github.com/VictoriaMetrics/VictoriaLogs/tree/master/deployment/docker/victorialogs/filebeat)
- [Fluentbit demo](https://github.com/VictoriaMetrics/VictoriaLogs/tree/master/deployment/docker/victorialogs/fluentbit)
- [Logstash demo](https://github.com/VictoriaMetrics/VictoriaLogs/tree/master/deployment/docker/victorialogs/logstash)
- [Vector demo](https://github.com/VictoriaMetrics/VictoriaLogs/tree/master/deployment/docker/victorialogs/vector)
- [Promtail demo](https://github.com/VictoriaMetrics/VictoriaLogs/tree/master/deployment/docker/victorialogs/promtail)

You can use [VictoriaLogs single](https://docs.victoriametrics.com/helm/victorialogs-single/)
or [cluster](https://docs.victoriametrics.com/helm/victorialogs-cluster) helm charts as a demo for running Vector
in Kubernetes with VictoriaLogs.
