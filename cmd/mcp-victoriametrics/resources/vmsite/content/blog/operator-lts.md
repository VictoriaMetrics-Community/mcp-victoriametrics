---
draft: false
page: blog blog_post
authors:
  - Vadim Rutkovsky
date: 2026-05-26
title: "Operator now has Long-Term Support (LTS) version"
enableComments: true
description: "Learn how VictoriaMetrics Operator's Kubernetes long term support release improves upgrade stability for production monitoring teams."
summary: "VictoriaMetrics Operator introduces Long-Term Support (LTS) releases starting with v0.68.x, ensuring stability and a predictable upgrade path for users."
categories:
  - Product News
tags:
  - kubernetes
  - operator
  - LTS
  - long-term support release
images:
  - /blog/operator-lts/preview.webp
---

> **Summary**: Introducing a Kubernetes long term support release model for the VictoriaMetrics Kubernetes operator to give production teams a more predictable, stable upgrade path. In this article, we explain how the new LTS release cycle works, including one year of support for each LTS line, a new LTS version every six months, and an overlap period so teams have time to test and migrate safely. We also cover why this matters for Kubernetes monitoring in production, how it reduces the risk of disruptive operator upgrades, and how dry-run mode helps teams preview changes before applying them to live environments.

VictoriaMetrics Operator has been developing at a neck-breaking pace, bringing numerous improvements, features, and fixes to our community. We usually make at least a single release every two weeks. While this rapid iteration cycle is great for delivering fixes and improvements quickly, it can be challenging for administrators managing critical production environments.

Frequent updates mean that changes in the operator may sometimes cause unexpected pod rollouts, restarts, or configuration modifications. For large-scale deployments, managing these frequent updates requires significant time and effort to test. To address these concerns and provide a more predictable, stable path for our enterprise and production users, we are would like to introduce Long-Term Support (LTS) releases for the VictoriaMetrics Operator.

## **The current LTS release: v0.68.x**

The v0.68 release is the current LTS version and brings many stability improvements and features, including:
- Introduction of the `VMDistributed` Custom Resource (CR) to easily deploy and manage VictoriaMetrics Distributed cluster components.
- Vertical Pod Autoscaler (VPA) support for `vmcluster`, `vtcluster`, `vlcluster` and `vmauth`.
- The `statefulRollingUpdateStrategyBehavior` property for managing `VMAgent` update strategy in a stateful mode.
- Improved logging with prettified reconcile diffs showing only changed JSON fields.

See latest v0.68.x release notes [here](https://github.com/VictoriaMetrics/operator/releases/tag/v0.68.4) and LTS-specific changelog is available [here](https://github.com/VictoriaMetrics/operator/blob/release-0.68/docs/CHANGELOG.md).

LTS releases include not only bug fixes and security patches, but also new features and improvements that are considered stable and suitable for production use. An example of such feature is dry-run mode.

## **Dry run mode**

One of the features that help with operator maintenance is the ability to see what changes will be applied without actually making them. Changes in the operator may sometimes be destructive or unexpected, so it is important to have a mode where diffs are printed, but not applied.

With the new dry-run mode, you can run the operator using the `--dry-run` argument or the `DRY_RUN` environment variable. In this mode, the operator uses a fake client that logs API calls instead of executing them against the Kubernetes API server. These calls also have a diff to be applied printed first. This is especially useful during operator upgrades to ensure that the upgrade won't cause any unexpected mutating changes.

## **Our Commitment**

We are committed to supporting every LTS release for a year and marking one of our releases as LTS every 6 months.

This allows our users, who prefer stability over new features, to stay on LTS releases. There is a 6 months overlap between LTS releases, which gives enough time for the upgrade to the next LTS release.

## Frequently Asked Questions

### What is a dry run in Kubernetes?

A dry run in Kubernetes lets users preview what would happen if a resource were created, updated, or changed, without actually applying that change to the cluster. It is useful because teams can validate manifests, inspect generated changes, check for potential issues, and reduce the risk of breaking production workloads. For operator upgrades, dry-run mode helps teams see what the operator intends to modify before those changes affect live Kubernetes resources.

### Why is the dry-run mode useful for Kubernetes Operator upgrades?

The dry-run mode is useful for Kubernetes Operator upgrades because operators can modify managed resources automatically. Without a preview, teams may not know which objects will change until after the upgrade is applied. Dry-run mode makes the process safer by showing intended changes first, helping teams catch risky updates, unexpected mutations, or configuration issues before they reach production.

### What is the difference between a standard release and an LTS release?

A standard release usually provides the latest features and improvements, but it may have a shorter support window. An LTS release focuses on stability and longer-term maintenance. For Kubernetes monitoring, this distinction matters because production teams often prioritize predictable behavior, compatibility, and safer upgrade planning over adopting every new feature immediately.

### How should teams safely upgrade a Kubernetes Operator?

Teams should safely upgrade a Kubernetes Operator by reviewing release notes, testing in a non-production environment, using dry-run mode where available, checking which resources the operator will modify, and scheduling upgrades during controlled maintenance windows. For production monitoring systems, using an LTS release can also reduce the need for frequent upgrades.

### How does the VictoriaMetrics Kubernetes operator LTS cycle work?

The VictoriaMetrics Kubernetes operator LTS model provides one year of support for each LTS release line, with new LTS versions planned every six months. This creates an overlap period between LTS releases, giving teams time to test, plan, and migrate without being forced into immediate upgrades.
