---
draft: true
page: blog blog_post
authors: 
  - Phuong Le
date: 2025-03-11
title: "gRPC Guide for Go"
summary: ""
enableComments: true
toc: true
categories:
  - Go @ VictoriaMetrics
  - Open Source Tech
tags:
  - go
  - golang
  - http2
  - grpc
images:
  - /blog/go-http2/go-http2-preview.webp
---

This article is part of the ongoing gRPC communication protocol series:

- [From net/rpc to gRPC in Go Applications](/blog/go-net-rpc/)
- [How HTTP/2 Works and How to Enable It in Go](/blog/go-http2/)
- [Practical Protobuf - From Basic to Best Practices](/blog/go-protobuf-basic/)
- [How Protobuf Works—The Art of Data Encoding](/blog/go-protobuf/)
- [gRPC Guide for Go](/blog/go-grpc/)

#### Synchronous vs Asynchronous

In gRPC, the `Send` and `Recv` methods are _synchronous operations_, but they interact with an underlying asynchronous transport layer.

- `Recv()`: blocks while waiting for a message to arrive from the network. It only returns when either: A complete message has been received and decoded; an error occurs (including `io.EOF` when the stream ends), or the context is canceled.
- `Send()`: blocks blocks until the message is encoded and written to the transport layer, or until an error occurs.

While the API methods are synchronous, the underlying transport layer in gRPC operates _asynchronously_. 

When a gRPC connection is established, it (either client or server) spins up two background goroutines: reader and writer. The message doesn't immediately get transmitted over the network when you call `Send()`. Instead, the message is serialized, compressed (if enabled), wrapped in a DATA frame, and placed into a _control buffer queue_.

That said, your `Send()` call only blocks until the message is successfully sitting in this buffer.

The writer goroutine is responsible for periodically processing items from this buffer and writes them to the network connection. The writer may even batch multiple frames together before flushing them to the network.

On the other side, the reader goroutine reads an HTTP/2 frame from the network. For data frames, it finds the corresponding stream for the frame, copies the data, and writes the data to the _stream's receive buffer_. The application code calls `Recv()` to get messages from this buffer.

When you call `Send`, it doesn't immediately send the message over the network. Instead, it queues the message to be sent later. The actual sending happens asynchronously in the background.

Similarly, when you call `Recv`, it doesn't block waiting for a response. Instead, it returns a stream of messages that you can read from. The actual receiving happens asynchronously in the background.

#### Intercept Every Message in the Stream?

To intercept every `Send()` or `Recv()` call, we need to wrap the `grpc.ServerStream` and override its SendMsg and RecvMsg methods.

Here’s how you wrap the ServerStream to log every message sent and received:

```go
type wrappedStream struct {
    grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
    fmt.Printf("Intercepted server Recv: %v\n", m)
    return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
    fmt.Printf("Intercepted server Send: %v\n", m)
    return w.ServerStream.SendMsg(m)
}
```

Now modify the streaming interceptor to wrap the stream before passing it to the actual gRPC method:

```go
func streamLoggingInterceptor(
    ctx context.Context,
    stream grpc.ServerStream,
    info *grpc.StreamServerInfo,
    handler grpc.StreamHandler,
) error {
    fmt.Printf("Streaming RPC started: %s\n", info.FullMethod)

    // Wrap the original stream
    wrapped := &wrappedStream{ServerStream: stream}

    err := handler(ctx, wrapped) // Call the actual gRPC method

    fmt.Printf("Streaming RPC completed: %s\n", info.FullMethod)
    return err
}
```

Now, what happens?

The interceptor still runs only once per streaming RPC. But now, every single message sent and received within the stream gets logged! This is because RecvMsg and SendMsg are now overridden to log the message before forwarding it.