---
draft: false
page: blog blog_post
authors:
 - Aliaksandr Valialkin
date: 2023-03-08
enableComments: true
title: "Save network costs with VictoriaMetrics remote write protocol"
summary: "Save network bandwidth costs when transferring data to VictoriaMetrics starting from v1.88"
categories: 
 - Performance
 - Monitoring
 - Product News
tags:
 - prometheus remote write protocol
 - save network costs
 - victoriametrics remote write protocol
images:
 - /blog/victoriametrics-remote-write/preview-1.webp
---
#### Prometheus remote write protocol

Prometheus remote write protocol is used by Prometheus for sending data to remote storage systems
such as VictoriaMetrics. See [these docs](https://docs.victoriametrics.com/#prometheus-setup) on how to setup Prometheus
to send the data to VictoriaMetrics. This protocol is very simple - it writes the collected [raw samples](https://docs.victoriametrics.com/keyConcepts.html#raw-samples)
into [WriteRequest protobuf message](https://github.com/prometheus/prometheus/blob/35026fb26d30fb63dbe6557058fb336c1bd400fa/prompb/remote.proto#L22),
then compresses the message with [Snappy compression algorithm](https://github.com/google/snappy) and sends it to the remote storage
in an [HTTP POST](https://en.wikipedia.org/wiki/POST_(HTTP)) request.

[vmagent](https://docs.victoriametrics.com/vmagent.html) uses Prometheus remote write protocol for transferring the collected samples
to remote storage specified via `-remoteWrite.url`.

The Prometheus remote write protocol serves well in most cases. But it isn't optimized for low network bandwidth usage.
So it can consume big amounts of network traffic when millions of samples per second must be transferred to the remote storage.
Is is OK when the remote storage is located in the same network as Prometheus, and this network has no limits on bandwidth
and the network transfer is free. In reality the remote storage may be located in another datacenter, availability zone or region.
In this case there may be some limits on network bandwidth and/or on network transfer costs for transferring the data
from Prometheus to the remote storage.

Let's calculate monthly costs for transferring the data from Prometheus located in Google Cloud to the remote storage located in AWS
at a rate of 1 million of samples/sec. Our internal stats show that an average data sample in production costs around 50 bytes to transfer via Prometheus remote write protocol. So 1 million samples/sec requires 50MB/sec of network bandwidth. This transforms to
`50MB/sec * 3600sec * 24h * 30d = 129600GB` of network traffic per month. 1GB of egress network traffic costs $0.08 at Google Cloud
according to [this pricing](https://cloud.google.com/vpc/network-pricing). So the monthly network transfer costs will be around `129600GB * $0.08 = $10K`.

The issue with high network costs when transferring the data from [vmagent](https://docs.victoriametrics.com/vmagent.html)
to Prometheus-compatible remote storage is quite common. That's why we started exploring on how to resolve it.

Prometheus remote write protocol transfers all the labels with each [sample](https://docs.victoriametrics.com/keyConcepts.html#raw-samples).
Real-world samples usually contain a big number of labels. For example:


```
process_cpu_seconds_total{
  job="foo",
  instance="bar",
  env="prod",
  namespace="default",
  container="qwerty",
  pod="abcdef",
  ...
} <value> <timestamp>
```

The average length of metric name plus all the labels per each sample in production is around 200 bytes according to our stats.
When the sample is encoded into [TimeSeries protobuf message](https://github.com/prometheus/prometheus/blob/35026fb26d30fb63dbe6557058fb336c1bd400fa/prompb/types.proto#L123),
its size becomes even bigger than the plaintext representation shown above. The average per-sample size on the wire is reduced to 50 bytes
thanks to [Snappy compression](https://github.com/google/snappy). But 50 bytes is still too big of a value compared to 0.4-0.8 bytes of disk space
needed per each sample stored by VictoriaMetrics.

#### Possible solutions for reducing network bandwidth costs

A single [time series](https://docs.victoriametrics.com/keyConcepts.html#time-series) usually consists of many samples.
These samples are sent to the remote storage with some interval. This interval is known as `scrape_interval` in the Prometheus ecosystem.
So we can assign a small id per each new time series seen on the wire and then send the series id together with `(value, timestamp)`
to the remote storage instead of sending the metric name plus labels next time. This allows sending `<4-byte sample id> + <8-byte value> + <8-byte timestamp> = 20 bytes`
instead of 200 bytes per each sample. This is `50/20 = 2.5x` better than Prometheus remote write protocol does.

We can go further and send [varint-encoded](https://stackoverflow.com/questions/24614553/why-is-varint-an-efficient-data-representation) difference (aka delta)
between the current value and the previous value per each sample. The same encoding technique can be applied to timestamps as well.
This allows reducing the on-the-wire sample size to ~10 bytes according to our tests.

Then a block of encoded samples can be compressed with [zstd compression](https://github.com/facebook/zstd) in order to reduce per-sample size to ~5 bytes.

This data transfer protocol allows reducing network bandwidth usage by `50/5 = 10x` comparing to Prometheus remote write protocol!

Unfortunately, this protocol has the following issues:

- It requires maintaining non-trivial amounts of state on both sender and receiver side. The sender must maintain a map for locating the series id
  by time series name plus labels. The receiver must maintain a map for locating the time series name plus labels by series ID.
  Both maps can be quite big when samples for millions of time series are transferred over the network.

- The maps can grow indefinitely when old time series are constantly substituted by new time series
  (aka [high churn rate](https://docs.victoriametrics.com/FAQ.html#what-is-high-churn-rate)).

- The state must be maintained individually per each connection (or session) between the sender and the receiver.
  This means that the memory usage at the receiver may go out of control when many independent senders transfer data
  to a single remote storage.
  This also means that it may be hard to use this protocol over [HTTP](https://en.wikipedia.org/wiki/HTTP).
  You need either to pass the same session id between multiple HTTP requests or to stream the data in
  a single request body via [chunked transfer encoding](https://en.wikipedia.org/wiki/Chunked_transfer_encoding).
  This may make hard load balancing and horizontal scalability at the receiver side.

- It requires additional CPU time for the encoding and decoding comparing to Prometheus remote write protocol.

That's why we decided to use a much simpler approach.

#### VictoriaMetrics remote write protocol

The first version of the VictoriaMetrics remote write protocol is almost identical to the Prometheus remote write protocol.
It writes the collected raw samples into [WriteRequest protobuf message](https://github.com/prometheus/prometheus/blob/35026fb26d30fb63dbe6557058fb336c1bd400fa/prompb/remote.proto#L22), but then compresses it with [zstd compression](https://github.com/facebook/zstd) instead of [snappy compression](https://github.com/google/snappy)
before sending it to the remote storage in an [HTTP POST](https://en.wikipedia.org/wiki/POST_(HTTP)) request.

The implementation of this protocol is very simple at both sender and receiver side - just replace `snappy` compression with `zstd` compression
and change `Content-Encoding: snappy` to `Content-Encoding: zstd` in HTTP request headers.

The VictoriaMetrics remote write protocol allows reducing network traffic costs by 2x-4x comparing to the Prometheus remote write protocol
at the cost of slightly higher CPU usage (+10% according to our production stats). Put it another way,
this reduces monthly network transfer costs from $10K to $2.5K when transferring a million of samples per second
between different cloud providers.

#### How to enable VictoriaMetrics remote write protocol?

Just [upgrade](https://docs.victoriametrics.com/#how-to-upgrade-victoriametrics) VictoriaMetrics components (including [vmagent](https://docs.victoriametrics.com/vmagent.html))
to [v1.88](https://docs.victoriametrics.com/CHANGELOG.html) or to any version beyond that.

`vmagent` automatically detects whether the configured remote storage supports VictoriaMetrics remote write protocol
and uses this protocol instead of Prometheus remote write protocol for data transfer.
See [these docs](https://docs.victoriametrics.com/vmagent.html#victoriametrics-remote-write-protocol) for more details.

#### Future work

We are going to explore and implement more advanced algorithms in the future versions of VictoriaMetrics remote write protocol,
in order to gain bigger cost savings.

It would be great if Prometheus itself and [other Prometheus-compatible remote storage systems](https://prometheus.io/docs/operating/integrations/#remote-endpoints-and-storage)
would support VictoriaMetrics the remote write protocol as well. The whole Prometheus ecosystem would benefit from this!
