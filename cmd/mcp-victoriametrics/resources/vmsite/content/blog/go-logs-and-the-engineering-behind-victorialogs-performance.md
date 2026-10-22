---
draft: false
page: blog blog_post
authors:
  - Michelle Sebek
date: 2026-10-16
title: "Go, Logs and the Engineering Behind VictoriaLogs Performance"
summary: "Tech Talk #15: Jesús Espino explains how architecture, data locality, Bloom filters, and Go help VictoriaLogs deliver efficient log storage and fast queries."
description: "Discover how VictoriaLogs uses LSM trees, Bloom filters, columnar storage, and Go to reduce work and improve performance in this Tech Talk #15 recap."
enableComments: true
categories:
  - Tech Talk
tags:
  - VictoriaLogs
  - Go
  - Performance
  - Tech Talk
  - Video
images:
  - /blog/go-logs-and-the-engineering-behind-victorialogs-performance/preview.webp
---

What actually makes a logging system fast?

That was the question behind Tech Talk #15.

And if you were expecting the answer to be a collection of clever Go optimizations, Jesús Espino, Principal Engineer at VictoriaMetrics, made an important point pretty early in our conversation:

**A huge part of VictoriaLogs' performance comes from architecture and data structures. The micro-optimizations come later.**

Now *that* gave my Tech Talk partner in crime, Mathias Palmersheim, and me plenty to dig into.

We went from LSM trees and data locality to Bloom filters, columnar storage, compression, Go's garbage collector, memory allocations and profiling. We even opened up a VictoriaLogs deployment to see how some of this is actually represented on disk.

We also managed to fit in foster dogs, Brazilian Jiu-Jitsu, a dead ZFS pool, and the fact that Jesús is technically a Lord of Sealand.

Because apparently that's what we do on Tech Talks now.

But underneath all of it was one engineering idea that kept coming back:

**Performance isn't only about doing work faster. Good architecture reduces how much work you need to do in the first place.**

Let's get into it.

## First: don't read data you don't need

One of the things I liked about this conversation was that the architecture started making sense once we connected it to a pretty straightforward goal:

**Read as little data as possible to answer the query.**

VictoriaLogs uses an LSM-tree architecture. Incoming logs are organized into immutable parts, and those parts are merged into progressively larger ones in the background. Logs are sorted and grouped around stream and time, improving data locality as parts are merged.

Mathias gave us a pretty good real-world analogy: moving.

At first, you're getting boxes into the house. Then you start unpacking and putting the kitchen stuff in the kitchen, bedroom stuff in the bedroom, and everything else where it belongs.

Keep moving into bigger houses—or eventually Jesús' castle—and you're still putting related things together.

That's essentially what those background merges are helping accomplish with the data.

Why does that matter?

Because most investigations already give the system useful information.

You're rarely asking:

*"Please search every log we've ever collected."*

You're much more likely asking:

*"What happened to this service during this time window?"*

When related data is stored together, VictoriaLogs has less irrelevant data to touch to answer that question.

That's where the architecture starts doing the performance work before the micro-optimizations even enter the conversation.

## Time + stream = a very good place to start

VictoriaLogs takes the idea of data locality further by organizing data into daily partitions. Within those partitions, it keeps logs from the same stream together.

So time isn't just something attached to a log entry. It's part of how VictoriaLogs narrows down where it needs to look.

If your query covers a particular time period, VictoriaLogs doesn't need to search partitions outside that period. Add a stream filter, and the search becomes even more targeted because the data is organized internally by **stream and time**.

That led to one of my favorite practical takeaways from the discussion:

**The fastest data to search is the data you don't have to search.**

There are plenty of times when you need to run a broad query. But asking any database to search terabytes of data without narrowing things down is going to require work.

Give it useful context, and suddenly a lot of that work disappears.

## Bloom filters: another way to avoid unnecessary work

Then we got into Bloom filters.

A Bloom filter doesn't tell VictoriaLogs with absolute certainty that a word *is* in a particular block.

What it can tell VictoriaLogs is that the word **definitely isn't** there.

That's incredibly useful.

If you're searching for something like `timeout`, VictoriaLogs can check the Bloom filter associated with a block. If the filter determines that `timeout` cannot be there, VictoriaLogs doesn't need to read that block.

Skip it. Move on.

If the Bloom filter says the term *could* be present, then VictoriaLogs may need to inspect the underlying data.

And this is where the live part of Tech Talks makes these conversations better.

**Mr. Alpaca in the live chat** asked a great question: are Bloom filters built for all the words in a block, or only a subset?

Jesús did something I always appreciate in a technical conversation—he said he wanted to verify it rather than guess.

So, we checked.

VictoriaLogs builds Bloom filters from the words seen in log fields and uses them to quickly eliminate blocks that cannot contain the searched terms.

**Live question asked. Documented answer found.**

And that's exactly why I like doing these Tech Talks live. Sometimes the best questions aren't the ones Mathias and I walked in with.

## Structured logs give us another advantage

Next came columnar storage and encoding.

With structured logs, you might have a timestamp, log level, IP address, user ID, message, and any number of other fields.

Those values don't all need to be treated as arbitrary strings.

VictoriaLogs stores fields using a column-oriented approach, which opens up some interesting opportunities for encoding.

Jesús gave several examples during the session.

Timestamps can be represented as differences between values rather than repeatedly storing full timestamps. IP addresses can be represented more efficiently than storing the full text each time. Repetitive values like log levels can use dictionary encoding rather than storing `info`, `warning`, or `error` over and over again.

Then compression gets applied on top of that.

The exact compression ratio is going to depend on your logs—and Mathias shared some pretty different examples from deployments he's seen—but the underlying engineering idea is what matters here.

Encode the data efficiently.

Compress it.

Store less.

And when it's time to query it, try to read only the pieces you actually need.

There is that theme again.

## A note on object storage

Object storage came up during the live conversation, particularly around organizations that need long retention periods.

That discussion raises an important distinction worth clarifying.

VictoriaLogs stores data in independent per-day partitions, and snapshots of those partitions can be backed up to object storage such as S3 or GCS.

That is **not** the same thing as VictoriaLogs using object storage directly as its primary data storage.

Primary object-storage support remains on the VictoriaLogs roadmap.

It's exactly the kind of distinction that's easy to lose in a live technical conversation, so if you're planning an architecture around object storage today, check the current VictoriaLogs documentation rather than relying on our whiteboard conversation.

## Then we switched gears: why Go?

About halfway through the Tech Talk we moved from:

**How does VictoriaLogs go fast?**

to:

**How does Go help us build fast software?**

And this part took the conversation somewhere I wasn't expecting.

Jesús talked about two kinds of performance.

There's obviously the performance users experience: how efficiently and quickly the software runs.

But there's also **developer performance**.

Go is intentionally a relatively small language. Its compiler is fast. It has a strong standard library, built-in testing and profiling capabilities, straightforward tooling and the ability to compile applications into static binaries.

That creates a fast feedback loop for engineers.

And engineering time matters.

If you're spending less time waiting for builds, dealing with tooling, or writing unnecessary boilerplate, you have more time to work on the architecture, features, and performance of the software itself.

There's also an operational benefit.

VictoriaMetrics and VictoriaLogs are good examples of what Go makes possible: compile the application into a single binary, download it, and run it.

There can be a lot of sophisticated engineering underneath something that's deliberately simple to operate.

That's kind of the point.

## Go is fast. That doesn't mean you stop measuring.

This is where another question from the live audience took us deeper.

**Wol forward APL** asked: since Go is garbage collected, how careful does VictoriaMetrics need to be about memory allocations and GC behavior to keep performance fast?

Jesús' answer was basically:

**Very careful—but where it matters.**

You don't need to obsess over every allocation everywhere.

Find the hot path first.

If a piece of code runs constantly and generates many allocations, reducing them can ease pressure on the garbage collector and improve performance.

If a function barely runs?

There are probably better places to spend your engineering time.

Go gives engineers profiling tools to measure CPU and memory behavior and figure out where resources are actually being spent.

Measure first.

Then optimize.

Seems obvious when you write it down.

We all know how often "obvious" engineering advice gets ignored.

And again, this is exactly why I like the live questions. We went from *“Go is fast”* into a much more useful conversation about what engineers actually have to think about when they're building software where performance matters.

## And then The Karate Kid entered the chat

This might have been my favorite part of the entire conversation.

Jesús explained that after identifying an expensive hot path, his first recommendation isn't necessarily to dive straight into optimizing the function.

Back up.

Look at the bigger architecture.

Could you cache something?

Could you reduce the amount of data going through the function?

Could you change the approach?

**Could you avoid doing that work altogether?**

And Mathias summed it up with *The Karate Kid*:

**“Best block, no be there.”**

There you have it.

Nearly an hour of Go and database performance distilled into Mr. Miyagi engineering advice.

Sometimes the best optimization isn't making something faster.

It's making sure you don't need to do it.

## Do newer versions of Go just make everything faster?

Before we wrapped, Mathias asked one more question I thought was worth keeping:

Go releases pretty frequently. When the Go runtime gets a new garbage collector or another performance improvement, does VictoriaMetrics automatically get faster?

And we got the most engineer answer possible:

**It depends.**

Different applications have different usage patterns. A runtime improvement might make a meaningful difference for one workload and very little difference for another.

Jesús also pointed out that VictoriaMetrics already pays close attention to things like garbage collection, so a runtime improvement doesn't automatically translate into the same percentage improvement at the application level.

There may be improvements.

But you measure them.

Which brings us right back to the point from earlier: **performance claims need data.**

I can get behind that.

## So, what actually makes VictoriaLogs fast?

By the end of the conversation, we'd covered LSM trees, streams, partitions, Bloom filters, columnar storage, encoding, compression, memory allocation, garbage collection, profiling, and Go itself.

But Jesús came back to where we started.

Yes, micro-optimizations matter.

At the scale VictoriaMetrics operates, engineers absolutely pay attention to them.

But as Jesús put it, those optimizations are the **cherry on top**.

The bigger performance story starts with architecture and design.

And I think that's what stuck with me most from this Tech Talk.

We talk a lot at VictoriaMetrics about **Simple, Efficient, Reliable Observability.**

Those aren't things you add at the end.

Efficiency starts with how you structure the data.

Simplicity starts with the engineering decisions underneath the product.

Reliability comes from thinking about how all of those pieces behave together.

Then you measure.

Then you optimize.

And apparently, somewhere along the way, you talk about foster dogs.

## Want to go deeper? Watch Tech Talk #15.

There's a lot I intentionally didn't try to reproduce here.

Mathias actually opens up his VictoriaLogs deployment during the session, and Jesús walks through what we're seeing on disk—including the daily partitions, parts, Bloom filters, timestamps, values, and metadata.

So if you're the person who wants to see what this looks like under the hood rather than just read about it, **[that's the part I'd watch.](https://www.youtube.com/live/2YQVgB3iUyM)**

▶️ **[Watch Tech Talk #15: Go, Logs and Performance](https://www.youtube.com/live/2YQVgB3iUyM)**

And if you join us live next time, bring questions.

Seriously.

Mr. Alpaca and Wol forward APL proved the point this month: some of the best parts of these conversations come from someone in the community asking something we didn't have on our list.

That's what I want Tech Talks to be.

Not another polished webinar where we talk *at* you for an hour.

Engineers talking to engineers, asking questions, opening things up, occasionally saying *“I need to check that”* and figuring things out together.

Until next month...

**Go, logs and dogs.**

Mic drop.
