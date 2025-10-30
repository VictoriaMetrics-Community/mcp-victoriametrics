---
draft: false
page: blog blog_post
authors:
  - Zhu Jiekun
date: 2025-10-27
title: "Discarding gRPC-Go: The Story Behind OTLP/gRPC Support in VictoriaTraces"
summary: "Why did VictoriaTraces build gRPC support without using gRPC-Go? And what are the benefits of adopting a simple HTTP/2 server and easyproto?"
enableComments: true
categories:
- Distributed Tracing
- Performance
- Benchmark
- Observability
- OpenTelemetry
- OTLP
tags:
 - distributed tracing
 - performance
 - benchmark
 - observability
 - opentelemetry
 - otlp
images:
  - /blog/opentelemetry-without-grpc-go/cover.webp
---

Let's begin with the results we achieved by discarding the use of [gRPC-Go](https://github.com/grpc/grpc-go) to build the gRPC server for OTLP/gRPC:
- Binary size: **-25%**
- CPU usage: **-36%**

## Background

The OpenTelemetry protocol (OTLP) is very popular for exchanging telemetry data between any OpenTelemetry instrumented applications and OpenTelemetry (compatible) collectors/backends.

Assume you have an application/collector which want to export data to another collector, you can config it to send data with:
- [ ] OTLP/gRPC exporter
- [x] OTLP/HTTP exporter, with protobuf payloads encoded in:
    - [x] binary format
    - [x] JSON format

Currently, VictoriaTraces only exposes an HTTP endpoint to receive data via the latter 2 formats: OTLP/HTTP binary & JSON.
There are a lot of applications out there that **only support sending data via OTLP/gRPC**, one typical example could be [kube-apiserver](https://kubernetes.io/docs/concepts/cluster-administration/system-traces/#kube-apiserver-traces).
So it's important to cover these cases as well.

## Supporting OTLP/gRPC

### The Goal

Our goal is to **provide a gRPC server** that can serve as a `TraceService` and **provide the `Export` method** for invocation,
as defined in [the OpenTelemetry's proto](https://github.com/open-telemetry/opentelemetry-proto/blob/v1.8.0/opentelemetry/proto/collector/trace/v1/trace_service.proto).

```proto
// Service that can be used to push spans between one Application instrumented with
// OpenTelemetry and a collector, or between a collector and a central collector (in this
// case spans are sent/received to/from multiple Applications).
service TraceService {
  rpc Export(ExportTraceServiceRequest) returns (ExportTraceServiceResponse) {}
}
```

What makes us consider not using gRPC-Go, or to be more specific, not using the `protoc` toolchain?

### Problem 1: The `protoc` Toolchain Isn't User-Friendly

The common way to build a gRPC server is:
- Use `protoc` and `protoc-gen-go` to generate the `struct`s of the **messages** defined in `.proto`.
- Use `protoc` and `protoc-gen-go-grpc` to generate **the service interface** defined in `.proto`.

But, wait. Let's recall how it could be done. Assume I (who is new to the Protobuf) just cloned the project, and want to add new messages/methods to the `.proto`:

1. Notice that the `protoc` toolchain is not part of my Ubuntu/MacOS.
2. Download the latest version of `protoc` from [the release page](https://github.com/protocolbuffers/protobuf/releases).
3. Ooops, `protoc` can't run solely if I want to generate go code. I need to `go install protoc-gen-go` and `go install protoc-gen-go-grpc`.
4. All set, what's the commands and flags to generate them? Google ["gRPC compile Go"](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code) for the tutorial.
5. Finally, I prepared the commands and run it locally. Bomb, it said dependency errors, because there are some `import`s in my `.proto` and I have to specify the paths.
6. After fixing them, rerun the commands and successfully generate (thousands of lines of) Go code.

But don't jump for joy too soon, because `protoc` may have unexpected surprises in store for you. 

"Why do the new code look different from the previous ones?"

The new one is like:

```go
type TracesData struct {
	state protoimpl.MessageState `protogen:"open.v1"`

	ResourceSpans []*ResourceSpans `protobuf:"bytes,1,rep,name=resource_spans,json=resourceSpans,proto3" json:"resource_spans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
```

While the previous one is like:

```go
type TracesData struct {
    ResourceSpans        []*ResourceSpans `protobuf:"bytes,1,opt,name=resource_spans,proto3" json:"resource_spans,omitempty"`
    XXX_NoUnkeyedLiteral struct{}         `json:"-"`
    XXX_unrecognized     []byte           `json:"-"`
    XXX_sizecache        int32            `json:"-"`
}
```

7. Google for the [reason](https://github.com/golang/protobuf/issues/276), then redo steps 2-6 using a different version of the `protoc` toolchain. (Sometimes you may need to `git log` to find out who the author of this code.)

All these steps show that compiling protobuf-related stuff is not as straightforward as writing simple HTTP JSON APIs.

It could become easy if you:

1. Write commands into the `Makefile`.
2. Use [Buf CLI](https://buf.build/product/cli) to compile without the `protoc` toolchain.

But many developers still opt for the HTTP JSON APIs.

That said, this hardly suffices to persuade us to discard the `protoc` toolchain. What else?

### Problem 2: The Existing Use of `easyproto` Instead of `golang/protobuf`

> [!NOTE]
> Problems 2 may provide further insights, though it should be clarified that it **only apply to VictoriaTraces** given its **unique context**.

In VictoriaMetrics, VictoriaLogs and VictoriaTraces, we use `easyproto` to marshal and unmarshal protobuf messages. The reasons are written in its [`README`](https://github.com/VictoriaMetrics/easyproto):

> - `easyproto` **doesn't require `protoc` or `go generate`**.
> - `easyproto` **doesn't increase the binary size** unlike traditional protoc-compiled code may do.
> - `easyproto` **allows writing zero-alloc code**.

However, to add OTLP/gRPC support, we need to consider the following:

1. If we simply build a gRPC server with code generated by `protoc`, how much will the **binary size** of the application increase?
2. Can we combine `easyproto` with gRPC? This way, `protoc` would only need to generate code for the gRPC service, and not for protobuf message `struct`s.
   - Note that this still requires importing gRPC-related packages, which weakens the second reason of using `easyproto` (aimed at reducing binary size).
3. Are there any other solutions that reuse `easyproto` **without importing new packages**?

## Unorthodox Way: An HTTP/2 Server

### The Theory

gRPC is a protocol that uses HTTP/2. So it's possible to implement an HTTP/2 server to serve requests at specific endpoints.

> gRPC can also use HTTP/3 (QUIC) and HTTP/1.1, but let’s avoid getting too deep into that for now.
> Just note that the current implementation also supports gRPC over HTTP/1.1.
> What’s more, thanks to its highly straightforward design, the OTLP/gRPC JSON format could also be supported very easily. 
> But as the JSON format currently only works with OTLP/HTTP, we haven’t put extra effort into it.
> 
> We'll leave it to readers to explore further.

According to [gRPC over HTTP2](https://grpc.github.io/grpc/core/md_doc__p_r_o_t_o_c_o_l-_h_t_t_p2.html), the data frame format is like:

```go
// +------------+---------------------------------------------+
// |   1 byte   |                 4 bytes                     |
// +------------+---------------------------------------------+
// | Compressed |               Message Length                |
// |   Flag     |                 (uint32)                    |
// +------------+---------------------------------------------+
// |                                                          |
// |                   Message Data                           |
// |                 (variable length)                        |
// |                                                          |
// +----------------------------------------------------------+
```

And the HTTP endpoint for `Export` method in `TraceService` is: `/opentelemetry.proto.collector.trace.v1.TraceService/Export`. 

The following code block shows you how the implementation looks like:

```go
// Init initializes an HTTP server.
func Init() {
	logger.Infof("starting OTLP gPRC server at :4317...")
	go httpserver.Serve(
		[]string{":4317"},
		OTLPGRPCRequestHandler,
		httpserver.ServeOptions{UseProxyProtocol: nil, DisableBuiltinRoutes: true, EnableHTTP2: true},
	)
}

// OTLPGRPCRequestHandler is the router of gRPC requests.
func OTLPGRPCRequestHandler(r *http.Request, w http.ResponseWriter) bool {
	switch r.URL.Path {
	case `/opentelemetry.proto.collector.trace.v1.TraceService/Export`:
		otlpExportTracesHandler(r, w)
	default:
		grpc.WriteErrorGrpcResponse(w, grpc.StatusCodeUnimplemented, fmt.Sprintf("gRPC method not found: %s", r.URL.Path))
	}
	return true
}

// otlpExportTracesHandler handles OTLP export traces requests.
func otlpExportTracesHandler(r *http.Request, w http.ResponseWriter) {
	// decompression with gzip
    ...

    // verify headers (5 bytes), and unmarshalling the rest bytes with easyproto
    ...

    // presisting data, and more
    ...

	writeExportTraceServiceResponse()
}
```

The complete code could be found in [VictoriaTraces #59](https://github.com/VictoriaMetrics/VictoriaTraces/pull/59).

### What’s the Cost

While the implementation looks straightforward and simple, there must be a cost.

So far this approach has only been tested with [the unary RPC](https://grpc.io/docs/what-is-grpc/core-concepts/#unary-rpc). For streaming RPC, we have no scenarios or motivation for further testing.

This approach can cover what we need for OTLP/gRPC, but it might not work for other cases. If you know more about that, feel free to leave a comment!

### Comparison

We conduct the benchmark against binary size and resource usage between different approaches of OTLP/gRPC support in VictoriaTraces:

1. Write an HTTP/2 server, and unmarshal with `easyproto`.
2. Compile an gRPC server with `protoc`, and unmarshal with the native gRPC decoder.

Additionally, the compiled gRPC server does support customizing encoder and decoder via the following code example. We add this to the comparison as well.

```go
import (
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterCodec(&easyProtoCodec{})
}
```

And here's the result.

For build size:

<div style="display: flex; gap: 10px;">
	<div class="" style="max-width: 50%;max-height: 50%;display: inline-block;">
 		<img src="/blog/opentelemetry-without-grpc-go/release_size.webp">
  	</div>
	<div class="" style="max-width: 50%;max-height: 50%;display: inline-block;">
  		<img src="/blog/opentelemetry-without-grpc-go/build_size.webp">
  	</div>
</div>

And regarding the performance of request handling (CPU usage, no-op: requests are responded immediately after decompressing and unmarshalling):

<div style="max-width: 50%; max-height: 50%; margin: 0 auto;">
 	<img src="/blog/opentelemetry-without-grpc-go/cpu_usage.webp" style="max-width: 100%; height: auto;">
</div>

The monitoring snapshot can be found [here](https://snapshots.raintank.io/dashboard/snapshot/9U4rWHfXx91AHwiaatmmF75XkmtOyiUH?orgId=0). CPU and memory profiles are available [here](/blog/opentelemetry-without-grpc-go/profiles.zip).

Based on the benchmark results, it is evident that **HTTP/2 combined with easyproto does demonstrate a clear advantage**.

## Conclusion

This blog shares the story **why VictoriaTraces implements gRPC server for OTLP/gRPC in the HTTP/2 + `easyproto` way**.
The core implementation was done by [@JayiceZ](https://github.com/JayiceZ), with the initial idea coming from [@makasim](https://github.com/makasim).

There are certain contextual reasons behind this, we're not try to persuade you to do so. But we see great potential and value, and better developer experience in this approach. 

As the VictoriaMetrics Stack aims for high performance and cost efficiency, every bit of saved CPU, memory, and network traffic matters significantly. 
And the same holds true for binary sizes, Docker image sizes, and other aspects,
just as mentioned in [this blog](https://valyala.medium.com/stripping-dependency-bloat-in-victoriametrics-docker-image-983fb5912b0d) by [Aliaksandr Valialkin](https://github.com/valyala) (founder of VictoriaMetrics),
and they remain as critical today as they were then.
