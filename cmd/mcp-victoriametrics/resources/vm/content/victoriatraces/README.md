> VictoriaTraces is currently under active development and not ready for production use. It is built on top of VictoriaLogs and therefore shares some flags and APIs. These will be fully separated once VictoriaTraces reaches a stable release. Until then, features may change or break without notice.

VictoriaTraces is an open-source, user-friendly database designed for storing and querying distributed [tracing data](https://en.wikipedia.org/wiki/Tracing_(software)), 
built by the [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) team.

VictoriaTraces provides the following features:
- It is resource-efficient and fast. It uses up to [**3.7x less RAM and up to 2.6x less CPU**](https://victoriametrics.com/blog/dev-note-distributed-tracing-with-victorialogs/) than other solutions such as Grafana Tempo.
- VictoriaTraces' capacity and performance scales linearly with the available resources (CPU, RAM, disk IO, disk space).
- It accepts trace spans in the popular [OpenTelemetry protocol](https://opentelemetry.io/docs/specs/otel/protocol/)(OTLP).
- It provides [Jaeger Query Service JSON APIs](https://www.jaegertracing.io/docs/2.6/apis/#internal-http-json) 
  to integrate with [Grafana](https://grafana.com/docs/grafana/latest/datasources/jaeger/) or [Jaeger Frontend](https://www.jaegertracing.io/docs/2.6/frontend-ui/).

## Quick Start

The easiest way to get started with VictoriaTraces is by using the pre-built Docker Compose file.
It launches VictoriaTraces, Grafana, and HotROD (a sample application that generates tracing data).
Everything is preconfigured and connected out of the box, so you can start exploring distributed tracing within minutes.

Clone the repository:
```bash 
git clone https://github.com/VictoriaMetrics/VictoriaTraces.git;
cd  VictoriaTraces;
```

Run VictoriaTraces with Docker Compose:
```bash
make docker-vt-single-up;
```

Now you can open HotROD at [http://localhost:8080](http://localhost:8080) and click around to generate some traces.
Then, you can open Grafana at [http://localhost:3000/explore](http://localhost:3000/explore) and explore the traces using the Jaeger data source.

To stop the services, run:
```bash
make docker-vt-single-down;
```

You can read more about docker compose and what's available there in the [Docker compose environment for VictoriaTraces](https://github.com/VictoriaMetrics/VictoriaMetrics/blob/deployment/victoriatraces/deployment/docker/README.md#victoriatraces-server).

### How to build from sources

Building from sources is reasonable when developing additional features specific to your needs or when testing bugfixes.

{{% collapse name="How to build from sources" %}}

Clone VictoriaTraces repository: 
```bash 
git clone https://github.com/VictoriaMetrics/VictoriaTraces.git;
cd  VictoriaTraces;
```

#### Build binary with go build

1. [Install Go](https://golang.org/doc/install).
1. Run `make victoria-traces` from the root folder of [the repository](https://github.com/VictoriaMetrics/VictoriaTraces).
   It builds `victoria-traces` binary and puts it into the `bin` folder.

#### Build binary with Docker

1. [Install docker](https://docs.docker.com/install/).
1. Run `make victoria-traces-prod` from the root folder of [the repository](https://github.com/VictoriaMetrics/VictoriaTraces).
   It builds `victoria-traces-prod` binary and puts it into the `bin` folder.

#### Building docker images

Run `make package-victoria-traces`. It builds `victoriametrics/victoria-traces:<PKG_TAG>` docker image locally.
`<PKG_TAG>` is auto-generated image tag, which depends on source code in the repository.
The `<PKG_TAG>` may be manually set via `PKG_TAG=foobar make package-victoria-traces`.

The base docker image is [alpine](https://hub.docker.com/_/alpine) but it is possible to use any other base image
by setting it via `<ROOT_IMAGE>` environment variable.
For example, the following command builds the image on top of [scratch](https://hub.docker.com/_/scratch) image:

```sh
ROOT_IMAGE=scratch make package-victoria-traces
```

{{% /collapse %}}

### Configure and run VictoriaTraces

VictoriaTraces can be run with:
```shell
/path/to/victoria-traces
```

or with Docker:
```shell
docker run --rm -it -p 10428:10428 -v ./victoria-traces-data:/victoria-traces-data \
  docker.io/victoriametrics/victoria-traces:latest
```

VictoriaTraces is configured via command-line flags. 
All the command-line flags have sane defaults, so there is no need in tuning them in general case. 
VictoriaTraces runs smoothly in most environments without additional configuration.

Pass `-help` to VictoriaTraces in order to see the list of supported command-line flags with their description and default values:

```bash
/path/to/victoria-traces -help
```

The following command-line flags are used the most:

* `-storageDataPath` - VictoriaTraces stores all the data in this directory. The default path is `victoria-traces-data` in the current working directory.
* `-retentionPeriod` - retention for stored data. Older data is automatically deleted. Default retention is 7 days.

Once it's running, it will listen to port `10428` (`-httpListenAddr`) and provide the following APIs:
1. for ingestion:
```
http://<victoria-traces>:10428/insert/opentelemetry/v1/traces
```
2. for querying:
```
http://<victoria-traces>:10428/select/jaeger/<endpoints>
```

See [data ingestion](https://docs.victoriametrics.com/victoriatraces/data-ingestion/) and [querying](https://docs.victoriametrics.com/VictoriaTraces/querying/) for more details.

## How does it work

VictoriaTraces was initially built on top of [VictoriaLogs](https://docs.victoriametrics.com/victorialogs/), a log database.
It receives trace spans in OTLP format, transforms them into structured logs, and provides [Jaeger Query Service JSON APIs](https://www.jaegertracing.io/docs/2.6/apis/#internal-http-json) for querying.

For detailed data model and example, see: [Key Concepts](https://docs.victoriametrics.com/victoriatraces/keyConcepts).

![How does VictoriaTraces work](how-does-it-work.webp)

Building VictoriaTraces in this way enables it to scale easily and linearly with the available resources, like VictoriaLogs.

## Multitenancy

VictoriaTraces supports multitenancy. A tenant is identified by `(AccountID, ProjectID)` pair, where `AccountID` and `ProjectID` are arbitrary 32-bit unsigned integers.
The `AccountID` and `ProjectID` fields can be set during [data ingestion](https://docs.victoriametrics.com/victoriatraces/data-ingestion/)
and [querying](https://docs.victoriametrics.com/victoriatraces/querying/) via `AccountID` and `ProjectID` request headers.

If `AccountID` and/or `ProjectID` request headers aren't set, then the default `0` value is used.

VictoriaTraces has very low overhead for per-tenant management, so it is OK to have thousands of tenants in a single VictoriaTraces instance.

VictoriaTraces doesn't perform per-tenant authorization. Use [vmauth](https://docs.victoriametrics.com/victoriametrics/vmauth/) or similar tools for per-tenant authorization.

## List of command-line flags

```shell
  -blockcache.missesBeforeCaching int
        The number of cache misses before putting the block into cache. Higher values may reduce indexdb/dataBlocks cache size at the cost of higher CPU and disk read usage (default 2)
  -defaultMsgValue string
        Default value for _msg field if the ingested log entry doesn't contain it; see https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field (default "missing _msg field; see https://docs.victoriametrics.com/victorialogs/keyconcepts/#message-field")
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
        Flag value can be read from the given file when using -flagsAuthKey=file:///abs/path/to/file or -flagsAuthKey=file://./relative/path/to/file . Flag value can be read from the given http/https url when using -flagsAuthKey=http://host/path or -flagsAuthKey=https://host/path
  -forceFlushAuthKey value
        authKey, which must be passed in query string to /internal/force_flush . It overrides -httpAuth.* . See https://docs.victoriametrics.com/victoriatraces/#forced-flush
        Flag value can be read from the given file when using -forceFlushAuthKey=file:///abs/path/to/file or -forceFlushAuthKey=file://./relative/path/to/file . Flag value can be read from the given http/https url when using -forceFlushAuthKey=http://host/path or -forceFlushAuthKey=https://host/path
  -forceMergeAuthKey value
        authKey, which must be passed in query string to /internal/force_merge . It overrides -httpAuth.* . See https://docs.victoriametrics.com/victoriatraces/#forced-merge
        Flag value can be read from the given file when using -forceMergeAuthKey=file:///abs/path/to/file or -forceMergeAuthKey=file://./relative/path/to/file . Flag value can be read from the given http/https url when using -forceMergeAuthKey=http://host/path or -forceMergeAuthKey=https://host/path
  -fs.disableMmap
        Whether to use pread() instead of mmap() for reading data files. By default, mmap() is used for 64-bit arches and pread() is used for 32-bit arches, since they cannot read data files bigger than 2^32 bytes in memory. mmap() is usually faster for reading small data chunks than pread()
  -futureRetention value
        Log entries with timestamps bigger than now+futureRetention are rejected during data ingestion; see https://docs.victoriametrics.com/victoriatraces/#retention
        The following optional suffixes are supported: s (second), h (hour), d (day), w (week), y (year). If suffix isn't set, then the duration is counted in months (default 2d)
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
  -http.header.frameOptions string
        Value for 'X-Frame-Options' header
  -http.header.hsts string
        Value for 'Strict-Transport-Security' header, recommended: 'max-age=31536000; includeSubDomains'
  -http.idleConnTimeout duration
        Timeout for incoming idle http connections (default 1m0s)
  -http.maxGracefulShutdownDuration duration
        The maximum duration for a graceful shutdown of the HTTP server. A highly loaded server may require increased value for a graceful shutdown (default 7s)
  -http.pathPrefix string
        An optional prefix to add to all the paths handled by http server. For example, if '-http.pathPrefix=/foo/bar' is set, then all the http requests will be handled on '/foo/bar/*' paths. This may be useful for proxied requests. See https://www.robustperception.io/using-external-urls-and-proxies-with-prometheus
  -http.shutdownDelay duration
        Optional delay before http server shutdown. During this delay, the server returns non-OK responses from /health page, so load balancers can route new requests to other servers
  -httpAuth.password value
        Password for HTTP server's Basic Auth. The authentication is disabled if -httpAuth.username is empty
        Flag value can be read from the given file when using -httpAuth.password=file:///abs/path/to/file or -httpAuth.password=file://./relative/path/to/file . Flag value can be read from the given http/https url when using -httpAuth.password=http://host/path or -httpAuth.password=https://host/path
  -httpAuth.username string
        Username for HTTP server's Basic Auth. The authentication is disabled if empty. See also -httpAuth.password
  -httpListenAddr array
        TCP address to listen for incoming http requests. See also -httpListenAddr.useProxyProtocol
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -httpListenAddr.useProxyProtocol array
        Whether to use proxy protocol for connections accepted at the given -httpListenAddr . See https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt . With enabled proxy protocol http server cannot serve regular /metrics endpoint. Use -pushmetrics.url for metrics pushing
        Supports array of values separated by comma or specified via multiple flags.
        Empty values are set to false.
  -inmemoryDataFlushInterval duration
        The interval for guaranteed saving of in-memory data to disk. The saved data survives unclean shutdowns such as OOM crash, hardware reset, SIGKILL, etc. Bigger intervals may help increase the lifetime of flash storage with limited write cycles (e.g. Raspberry PI). Smaller intervals increase disk IO load. Minimum supported value is 1s (default 5s)
  -insert.concurrency int
        The average number of concurrent data ingestion requests, which can be sent to every -storageNode (default 2)
  -insert.disable
        Whether to disable /insert/* HTTP endpoints
  -insert.disableCompression
        Whether to disable compression when sending the ingested data to -storageNode nodes. Disabled compression reduces CPU usage at the cost of higher network usage
  -insert.maxFieldsPerLine int
        The maximum number of log fields per line, which can be read by /insert/* handlers; see https://docs.victoriametrics.com/victorialogs/faq/#how-many-fields-a-single-log-entry-may-contain (default 1000)
  -insert.maxQueueDuration duration
        The maximum duration to wait in the queue when -maxConcurrentInserts concurrent insert requests are executed (default 1m0s)
  -internStringCacheExpireDuration duration
        The expiry duration for caches for interned strings. See https://en.wikipedia.org/wiki/String_interning . See also -internStringMaxLen and -internStringDisableCache (default 6m0s)
  -internStringDisableCache
        Whether to disable caches for interned strings. This may reduce memory usage at the cost of higher CPU usage. See https://en.wikipedia.org/wiki/String_interning . See also -internStringCacheExpireDuration and -internStringMaxLen
  -internStringMaxLen int
        The maximum length for strings to intern. A lower limit may save memory at the cost of higher CPU usage. See https://en.wikipedia.org/wiki/String_interning . See also -internStringDisableCache and -internStringCacheExpireDuration (default 500)
  -internalinsert.disable
        Whether to disable /internal/insert HTTP endpoint
  -internalinsert.maxRequestSize size
        The maximum size in bytes of a single request, which can be accepted at /internal/insert HTTP endpoint
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 67108864)
  -internalselect.disable
        Whether to disable /internal/select/* HTTP endpoints
  -logIngestedRows
        Whether to log all the ingested log entries; this can be useful for debugging of data ingestion; see https://docs.victoriametrics.com/victoriatraces/data-ingestion/ ; see also -logNewStreams
  -logNewStreams
        Whether to log creation of new streams; this can be useful for debugging of high cardinality issues with log streams; see https://docs.victoriametrics.com/victoriatraces/keyconcepts/#stream-fields ; see also -logIngestedRows
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
        The maximum number of concurrent insert requests. Set higher value when clients send data over slow networks. Default value depends on the number of available CPU cores. It should work fine in most cases since it minimizes resource usage. See also -insert.maxQueueDuration (default 48)
  -memory.allowedBytes size
        Allowed size of system memory VictoriaMetrics caches may occupy. This option overrides -memory.allowedPercent if set to a non-zero value. Too low a value may increase the cache miss rate usually resulting in higher CPU and disk IO usage. Too high a value may evict too much data from the OS page cache resulting in higher disk IO usage
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 0)
  -memory.allowedPercent float
        Allowed percent of system memory VictoriaMetrics caches may occupy. See also -memory.allowedBytes. Too low a value may increase cache miss rate usually resulting in higher CPU and disk IO usage. Too high a value may evict too much data from the OS page cache which will result in higher disk IO usage (default 60)
  -metrics.exposeMetadata
        Whether to expose TYPE and HELP metadata at the /metrics page, which is exposed at -httpListenAddr . The metadata may be needed when the /metrics page is consumed by systems, which require this information. For example, Managed Prometheus in Google Cloud - https://cloud.google.com/stackdriver/docs/managed-prometheus/troubleshooting#missing-metric-type
  -metricsAuthKey value
        Auth key for /metrics endpoint. It must be passed via authKey query arg. It overrides -httpAuth.*
        Flag value can be read from the given file when using -metricsAuthKey=file:///abs/path/to/file or -metricsAuthKey=file://./relative/path/to/file . Flag value can be read from the given http/https url when using -metricsAuthKey=http://host/path or -metricsAuthKey=https://host/path
  -opentelemetry.traces.maxRequestSize size
        The maximum size in bytes of a single OpenTelemetry trace export request.
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 67108864)
  -pprofAuthKey value
        Auth key for /debug/pprof/* endpoints. It must be passed via authKey query arg. It overrides -httpAuth.*
        Flag value can be read from the given file when using -pprofAuthKey=file:///abs/path/to/file or -pprofAuthKey=file://./relative/path/to/file . Flag value can be read from the given http/https url when using -pprofAuthKey=http://host/path or -pprofAuthKey=https://host/path
  -pushmetrics.disableCompression
        Whether to disable request body compression when pushing metrics to every -pushmetrics.url
  -pushmetrics.extraLabel array
        Optional labels to add to metrics pushed to every -pushmetrics.url . For example, -pushmetrics.extraLabel='instance="foo"' adds instance="foo" label to all the metrics pushed to every -pushmetrics.url
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -pushmetrics.header array
        Optional HTTP request header to send to every -pushmetrics.url . For example, -pushmetrics.header='Authorization: Basic foobar' adds 'Authorization: Basic foobar' header to every request to every -pushmetrics.url
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -pushmetrics.interval duration
        Interval for pushing metrics to every -pushmetrics.url (default 10s)
  -pushmetrics.url array
        Optional URL to push metrics exposed at /metrics page. See https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#push-metrics . By default, metrics exposed at /metrics page aren't pushed to any remote storage
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -retention.maxDiskSpaceUsageBytes size
        The maximum disk space usage at -storageDataPath before older per-day partitions are automatically dropped; see https://docs.victoriametrics.com/victoriatraces/#retention-by-disk-space-usage ; see also -retentionPeriod
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 0)
  -retentionPeriod value
        Log entries with timestamps older than now-retentionPeriod are automatically deleted; log entries with timestamps outside the retention are also rejected during data ingestion; the minimum supported retention is 1d (one day); see https://docs.victoriametrics.com/victoriatraces/#retention ; see also -retention.maxDiskSpaceUsageBytes
        The following optional suffixes are supported: s (second), h (hour), d (day), w (week), y (year). If suffix isn't set, then the duration is counted in months (default 7d)
  -search.maxConcurrentRequests int
        The maximum number of concurrent search requests. It shouldn't be high, since a single request can saturate all the CPU cores, while many concurrently executed requests may require high amounts of memory. See also -search.maxQueueDuration (default 16)
  -search.maxQueryDuration duration
        The maximum duration for query execution. It can be overridden to a smaller value on a per-query basis via 'timeout' query arg (default 30s)
  -search.maxQueueDuration duration
        The maximum time the search request waits for execution when -search.maxConcurrentRequests limit is reached; see also -search.maxQueryDuration (default 10s)
  -search.traceMaxDurationWindow duration
        The window of searching for the rest trace spans after finding one span.It allows extending the search start time and end time by -search.traceMaxDurationWindow to make sure all spans are included.It affects both Jaeger's /api/traces and /api/traces/<trace_id> APIs. (default 10m0s)
  -search.traceMaxServiceNameList uint
        The maximum number of service name can return in a get service name request. This limit affects Jaeger's /api/services API. (default 1000)
  -search.traceMaxSpanNameList uint
        The maximum number of span name can return in a get span name request. This limit affects Jaeger's /api/services/*/operations API. (default 1000)
  -search.traceSearchStep duration
        Splits the [0, now] time range into many small time ranges by -search.traceSearchStep when searching for spans by trace_id. Once it finds spans in a time range, it performs an additional search according to -search.traceMaxDurationWindow and then stops. It affects Jaeger's /api/traces/<trace_id> API. (default 24h0m0s)
  -search.traceServiceAndSpanNameLookbehind duration
        The time range of searching for service name and span name. It affects Jaeger's /api/services and /api/services/*/operations APIs. (default 168h0m0s)
  -select.disable
        Whether to disable /select/* HTTP endpoints
  -select.disableCompression
        Whether to disable compression for select query responses received from -storageNode nodes. Disabled compression reduces CPU usage at the cost of higher network usage
  -storage.minFreeDiskSpaceBytes size
        The minimum free disk space at -storageDataPath after which the storage stops accepting new data
        Supports the following optional suffixes for size values: KB, MB, GB, TB, KiB, MiB, GiB, TiB (default 10000000)
  -storageDataPath string
        Path to directory where to store VictoriaTraces data; see https://docs.victoriametrics.com/victoriatraces/#storage (default "victoria-traces-data")
  -storageNode array
        Comma-separated list of TCP addresses for storage nodes to route the ingested logs to and to send select queries to. If the list is empty, then the ingested logs are stored and queried locally from -storageDataPath
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.bearerToken array
        Optional bearer auth token to use for the corresponding -storageNode
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.bearerTokenFile array
        Optional path to bearer token file to use for the corresponding -storageNode. The token is re-read from the file every second
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.password array
        Optional basic auth password to use for the corresponding -storageNode
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.passwordFile array
        Optional path to basic auth password to use for the corresponding -storageNode. The file is re-read every second
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.tls array
        Whether to use TLS (HTTPS) protocol for communicating with the corresponding -storageNode. By default communication is performed via HTTP
        Supports array of values separated by comma or specified via multiple flags.
        Empty values are set to false.
  -storageNode.tlsCAFile array
        Optional path to TLS CA file to use for verifying connections to the corresponding -storageNode. By default, system CA is used
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.tlsCertFile array
        Optional path to client-side TLS certificate file to use when connecting to the corresponding -storageNode
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.tlsInsecureSkipVerify array
        Whether to skip tls verification when connecting to the corresponding -storageNode
        Supports array of values separated by comma or specified via multiple flags.
        Empty values are set to false.
  -storageNode.tlsKeyFile array
        Optional path to client-side TLS certificate key to use when connecting to the corresponding -storageNode
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.tlsServerName array
        Optional TLS server name to use for connections to the corresponding -storageNode. By default, the server name from -storageNode is used
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -storageNode.username array
        Optional basic auth username to use for the corresponding -storageNode
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tls array
        Whether to enable TLS for incoming HTTP requests at the given -httpListenAddr (aka https). -tlsCertFile and -tlsKeyFile must be set if -tls is set. See also -mtls
        Supports array of values separated by comma or specified via multiple flags.
        Empty values are set to false.
  -tlsCertFile array
        Path to file with TLS certificate for the corresponding -httpListenAddr if -tls is set. Prefer ECDSA certs instead of RSA certs as RSA certs are slower. The provided certificate file is automatically re-read every second, so it can be dynamically updated. See also -tlsAutocertHosts
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tlsCipherSuites array
        Optional list of TLS cipher suites for incoming requests over HTTPS if -tls is set. See the list of supported cipher suites at https://pkg.go.dev/crypto/tls#pkg-constants
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tlsKeyFile array
        Path to file with TLS key for the corresponding -httpListenAddr if -tls is set. The provided key file is automatically re-read every second, so it can be dynamically updated. See also -tlsAutocertHosts
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -tlsMinVersion array
        Optional minimum TLS version to use for the corresponding -httpListenAddr if -tls is set. Supported values: TLS10, TLS11, TLS12, TLS13
        Supports an array of values separated by comma or specified via multiple flags.
        Value can contain comma inside single-quoted or double-quoted string, {}, [] and () braces.
  -version
        Show VictoriaMetrics version
```
