---
draft: false
page: blog blog_post
authors:
 - Jean-Jerome Schmidt-Soisson
 - Marc Sherwood
date: 2025-06-19
enableComments: true
title: "VictoriaLogs Unleashed: Cluster Version Now Available for Exceptional, Linear Scaling"
summary: "We're thrilled to announce the release of the VictoriaLogs Cluster version – one of the most requested and anticipated updates from our user community. If you've been pushing the boundaries of vertical scaling, the solution for horizontal scalability is now here."
categories: 
 - Product News
tags:
 - victorialogs
 - cluster
 - open source
 - logging
 - logs database
 - scalability
images:
 - /blog/victorialogs-unleashed-cluster-version-now-available/victorialogs-unleashed-cluster-version-now-available-for-exceptional-linear-scaling.webp
---
![VictoriaLogs Cluster Version Now Available for Exceptional, Linear Scaling](/blog/victorialogs-unleashed-cluster-version-now-available/victorialogs-unleashed-cluster-version-now-available-for-exceptional-linear-scaling.webp)

You asked, and we listened!

We're thrilled to announce the release of the [VictoriaLogs Cluster version](https://docs.victoriametrics.com/victorialogs/cluster/) – one of the most requested and anticipated updates from our user community. This marks a significant leap forward for VictoriaLogs, empowering users to handle log volumes and ingestion rates far beyond the limits of a single node. 
VictoriaLogs is the open source, user-friendly database for logs that is resource-efficient and fast, using up to 30x less RAM and up to 15x less disk space than comparable solutions. 
If you've been pushing the boundaries of vertical scaling, the solution for horizontal scalability is now here!

## Why Cluster Mode? Addressing Scalability Limits

While single-node VictoriaLogs offers exceptional efficiency and simplicity (and remains the preferred choice if it meets your needs!), we know many of you operate at a scale where the CPU, RAM, storage, or I/O limits of a single host become a bottleneck. 
VictoriaLogs Cluster directly addresses this by providing robust horizontal scaling capabilities. You can now distribute your logging infrastructure across multiple nodes, ensuring performance and capacity keep pace with your growing demands.

## [Simple, Powerful Architecture](https://docs.victoriametrics.com/victorialogs/cluster/#architecture)

VictoriaLogs Cluster introduces a straightforward, distributed architecture built on the foundation of the single-node version. Getting started is designed to be simple. You can begin with a minimal cluster where the `vlstorage` node is your existing single-node instance, adding just the lightweight `vlinsert` and `vlselect` components.
1. **vlinsert**: Handles log ingestion, accepting data through all supported protocols and intelligently distributing it across the available storage nodes.
2. **vlstorage**: The workhorse for storing log data. **Crucially, each vlstorage node is essentially a single-node VictoriaLogs instance**, making migration seamless.
3. **vlselect**: Manages incoming queries, fanning them out to the relevant vlstorage nodes, processing the results, and returning them to the user.
Each of these components can be scaled independently, allowing you to tailor your cluster resources precisely to your workload demands – whether you need more ingestion power, query capacity, or storage.

## Key Benefits of VictoriaLogs Cluster:

* **Linear Scalability**: Break free from single-server limitations. The performance and capacity of VictoriaLogs Cluster scale linearly with the number of nodes added, allowing predictable expansion as your needs grow. 
* **Seamless Migration**: Upgrading from a single-node setup is remarkably straightforward. Existing single-node instances can directly become vlstorage nodes in your new cluster – just upgrade the executable and update the configuration!
* **Global Query View**: Need to see logs across different sites or stages? `vlselect` enables querying across multiple `vlstorage` nodes or even entirely separate VictoriaLogs clusters (or single-node instances), consolidating results into a single view.
* **Architectural Flexibility**: Scale ingestion, querying, and storage independently. You can even build multi-level cluster setups for complex environments.
* **Performance Focused**: Designed for efficient distributed operation, including automatic performance tuning and options for managing network traffic via compression settings.
* **Operational Simplicity**: Despite being a distributed system, we've focused on maintaining the ease of use VictoriaMetrics products are known for, including clear configuration and leveraging the familiar single-node instance as the core storage component. By using HTTP, users are able to easily integrate tools to help secure, route, and modify data in transit. 

## When to Choose Cluster vs. Single-Node

It's important to reiterate: If your logging needs are comfortably met by a single-node VictoriaLogs instance (perhaps by scaling up the hardware of that single node), that remains the simpler and often more performant option due to the lack of network overhead. And thanks to [storage mode duality feature](https://docs.victoriametrics.com/victorialogs/cluster/#single-node-and-cluster-mode-duality) it is very easy to upgrade to cluster mode when needed.
The cluster version is specifically designed for scenarios where **horizontal scaling across multiple machines is necessary** because you're hitting those single-host limits.

## Getting Started

Ready to scale? Setting up a basic cluster is easy. Our documentation includes:
* Cluster Architecture Overview: [Documentation](https://docs.victoriametrics.com/victorialogs/cluster/#single-node-and-cluster-mode-duality)
* Quick Start Guide: [Follow these steps to get a multi-node cluster running locally in minutes!](https://docs.victoriametrics.com/victorialogs/cluster/#quick-start)
* Deployment with Container Tools: For easy local development and testing, we provide examples using Docker Compose. For production deployments on Kubernetes, official Helm charts are available to streamline your setup.
* Security Best Practices: [Learn how to secure communication between cluster components](https://docs.victoriametrics.com/victorialogs/cluster/#security).

## Share Your Feedback

The arrival of the cluster version is a major milestone for VictoriaLogs, driven directly by the needs and feedback of our incredible user base. We're incredibly excited to see how you leverage this new capability to manage your logs at scale.

Download the [latest release](https://docs.victoriametrics.com/victorialogs/quickstart/), explore the documentation, and join the conversation in our [community channels](https://victoriametrics.com/community/) to share your feedback!
