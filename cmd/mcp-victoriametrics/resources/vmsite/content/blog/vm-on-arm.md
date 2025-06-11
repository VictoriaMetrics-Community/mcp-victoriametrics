---
draft: false
page: blog blog_post
authors:
 - Nikolay Khramchikhin
date: 2022-03-11
title: "Running VictoriaMetrics on ARM-based processors"
enableComments: true
summary: "VictoriaMetrics has new production ready builds for ARM"
categories: 
 - Company News
 - Product News
 - Performance
tags:
 - open source
 - monitoring
 - time series database
 - victoriametrics
 - ARM
images:
 - /blog/vm-on-arm/cpu_optimized.webp
---

## The future is now and it's ARM

ARM processors become more popular and more cost-effective according to many benchmarks. 
One of them was made by [Percona for MySQL](https://www.percona.com/blog/comparing-graviton-performance-to-arm-and-intel-for-mysql/).

Some of our users reported issues with VictoriaMetrics at [AWS Graviton](https://aws.amazon.com/pm/ec2-graviton) instances.
The main concerns were higher CPU and disk IO usage compared to x86 instances of the same size and for the same workload.
  
By that time, we verified that VictoriaMetrics works fine for raspberry and IoT devices,
but didn't do any optimizations for ARM builds.

## Benchmarks

The main difference between x86 and ARM builds is the library we use for data encoding.
x86-build uses [gozstd](https://github.com/valyala/gozstd) library, a wrapper over
[Facebook's zstd](https://github.com/facebook/zstd) written in C.

Cross-compiled ARM64 build uses [compress](https://github.com/klauspost/compress) library 
by [Klaus Post](https://github.com/klauspost) written in native Go.

So in many aspects, the performance difference between x86-build and cross-compiled ARM64 build
heavily depends on the performance of these libraries.

That's why, in order to improve performance of the ARM builds we added [CGO](https://pkg.go.dev/cmd/cgo) 
build with [some](https://github.com/VictoriaMetrics/VictoriaMetrics/commit/a8acad7453365254ae4e59a91bfcd66556d8dfdd)
[tweaks](https://github.com/valyala/gozstd/commit/d4c2028fade890a6f5ee1ac6bd0650d688faf8f0).


### Testing env

For my test I created 3 instances at AWS:
1. `m5.4xlarge` - intel x86 based instances with **16 CPU** and **64 RAM** $0.768 hour
to run x86 build of VictoriaMetrics.
2. `m6g.4xlarge` - graviton2 based instance with **16 CPU** and **64 RAM** $0.61 hour
to run cross-compiled ARM64 build without [CGO](https://pkg.go.dev/cmd/cgo).
3. `m6g.4xlarge` - graviton2 based instance with **16 CPU** and **64 RAM** $0.616 hour
to run [CGO](https://pkg.go.dev/cmd/cgo) build for ARM64.

And installed [vmsingle](https://docs.victoriametrics.com/Single-server-VictoriaMetrics.html) 
[v1.72.0](https://docs.victoriametrics.com/CHANGELOG.html#v1720) on each of the instances.


### Workload

For workload generation, I've used our [benchmark suite](https://github.com/VictoriaMetrics/prometheus-benchmark)
and set up a separate [vmsingle](https://docs.victoriametrics.com/Single-server-VictoriaMetrics.html) node for metrics collection.
 

### Results

Initial CPU profiling proves this theory and shows performance improvements with [gozstd](https://github.com/valyala/gozstd) lib:

{{< image href="/blog/vm-on-arm/cpu_before_optimize.webp" class="wide-img" alt="CPU usage by VictoriaMetrics before optimizations" >}}
{{< image href="/blog/vm-on-arm/cpu_optimized.webp" class="wide-img" alt="CPU usage by VictoriaMetrics after optimizations" >}}


Optimized ARM and x86 versions show almost the same result for disk IO usage:

{{< image href="/blog/vm-on-arm/disk_usage.webp" class="wide-img" alt="Disk writes/reads during the benchmark" >}}


Query performance for x86 version outperforms optimized ARM by ~10% and unoptimized ARM by ~25%:

{{< image href="/blog/vm-on-arm/query_performance.webp" class="wide-img" alt="Query latency during the benchmark" >}}
 

## Building ARM64 golang with CGO

One of major challenges was to add CGO build into our cross-compilation pipeline. 
We are using [musl](https://www.musl-libc.org/how.html) based builds and the default musl compiler 
isn't aware of how to build code for ARM. Instead, special [aarch64-musl-gcc compiler](https://musl.cc/) must be used:
```
CC=/path_to_folder/bin/aarch64-linux-musl-gcc \
GOOS=linux GOARCH=arm64 CGO_ENABLED=1 go build main.go
```

Important note, your C-lang dependencies must be built with the same compiler. 
In my case, I had to rebuild [gozstd](https://github.com/valyala/gozstd) lib.


## Conclusions

 - VictoriaMetrics for ARM has better cost-performance compared with x86 machines. 
ARM instances are ~ 20% cheaper than x86 with the same performance.
 - Read queries latency at x86 system is better - x86 instance has ~10% lower query duration.
 - VictoriaMetrics has production-ready builds for ARM, prebuilt binaries, docker images 
since [v1.73.0 release](https://docs.victoriametrics.com/CHANGELOG.html#v1730).

