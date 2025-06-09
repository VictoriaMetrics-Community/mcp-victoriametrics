---
draft: false
page: blog blog_post
authors:
 - Artem Navoiev
date: 2022-08-04
title: "How to Decommission a vmstorage Node from a VictoriaMetrics Cluster"
summary: "An article about how to remove a storage node from a VictoriaMetrics cluster gracefully"
enableComments: true
categories:
 - Time Series Database
tags:
 - victoriametrics
 - cluster
keywords: 
 - cluster
 - decommission
 - performance
images:
 - /blog/storage-node-decommission/decommission-node-3.webp
---

## **Problem**

We need to remove a vmstorage node from [VictoriaMetrics cluster](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html) gracefully. Every vmstorage node contains its own portion of data and removing the vmstorage node from the cluster creates gaps in the graph (because replication is out of scope).

## **Setup example**

We have a VictoriaMetrics cluster with 2 vminsert, 2 vmselect and 3 vmstorage nodes. We want to gracefully remove `vmstorage A` from the cluster.
{{< image href="/blog/storage-node-decommission/decommission-node-1.webp" >}}

### **Solution One**

{{< image href="/blog/storage-node-decommission/decommission-node-2.webp" >}}
<p></p>

1. Remove `vmstorage A` from the vminsert list
2. Wait for the retention period
3. Remove `vmstorage A` from the cluster

**Note**: please expect higher resource usage on the existing vmstorage nodes (`vmstorage B` and `vmstorage C`), as they now need to handle all the incoming data.

**Pros**: Simple implementation

**Cons**: You may need to wait for a long period of time


### **Solution Two**

{{< image href="/blog/storage-node-decommission/decommission-node-3.webp" >}}
<p></p>

1. Remove `vmstorage A` from the vminsert list (same as in Solution One).
2. Set up a dedicated vmselect node that knows only about the vmstorage node that we want to remove (vmstorage A). We need this vmselect node for migration data from vmstorage A to other vmstorage nodes in the cluster.
3. Using [vmctl native import/export](https://docs.victoriametrics.com/vmctl.html#migrating-data-from-victoriametrics) reads data from vmselect for `vmstorage A` and writes data back to vminsert nodes. 4. This process creates duplicates.
5. Turn on [deduplication on vmselect nodes](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#deduplication).
6. Remove `vmstorage A` from the cluster.

**Note**: Please expect higher resource usage on the existing nodes (`vmstorage B` and `vmstorage C`), as they now need to handle all the incoming data.

**Pros**: Faster way to decommission a vmstorage node.

**Cons**: The process is more complex compared to solution One. The vmctl import/export process may require tuning if you migrate hundreds GB of data (or more).

*Hint*	: [downsampling](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html#downsampling) reduces the amount of data in a cluster; after downsampling, the vmctl migration requires less data to transfer and less time.

We trust that this is helpful!

Please let us know how you get on or if you have any questions by submitting a comment below.
