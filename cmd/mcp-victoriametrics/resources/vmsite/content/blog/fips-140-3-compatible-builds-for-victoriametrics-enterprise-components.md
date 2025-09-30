---
draft: false
page: blog blog_post
authors:
  - Artem Navoiev
  - Zakhar Bessarab
date: 2025-06-24
title: "FIPS 140-3 Compatible Builds for VictoriaMetrics Enterprise Components"
enableComments: true
summary: "VictoriaMetrics Enterprise now offers FIPS 140-3 compatible builds, leveraging the BoringCrypto module. This enables organizations in regulated sectors, such as the federal government, finance, and healthcare, to meet stringent cryptographic requirements using VictoriaMetrics."
categories:
  - Product News
tags:
  - FIPS 140-3 
  - victoriaMetrics enterprise
  - observability
  - FedRamp
images:
  - /blog/fips-140-3-compatible-builds-for-victoriametrics-enterprise-components/preview.webp
---

VictoriaMetrics introduces FIPS 140-3 compatible builds for its components, starting with version `1.117.0`.  These builds utilize Google’s [FIPS 140-3](https://go.dev/doc/security/fips140) validated BoringCrypto module.

This is critical for customers in regulated environments (federal government, finance, healthcare) to meet FIPS 140-3 cryptographic requirements for data encryption, TLS, and secure communications.

While VictoriaMetrics itself is not a FIPS-certified cryptographic module, these builds ensure all cryptographic operations are handled by the validated BoringCrypto module. This simplifies compliance for system integrators and customers undergoing FedRAMP, HIPAA, or similar audits.

## Availability

FIPS-compatible binaries are on our GitHub repository; images are on DockerHub or Quay registries.

* *Binaries*: Attached to enterprise assets, alongside regular versions. Example: `victoria-metrics-darwin-amd64-v1.117.0-enterprise.tar.gz` (There is a `victoriametrics-prod-fips binary` in the archive).
* *Container Images*: `scratch`-based, with a `-fips` suffix in tags. Example: `v1.117.1-enterprise-cluster-fips`.
* *Architecture Support*: Available for `arm64` and `amd64`.


## Performance

Our internal tests show no measurable impact on resource consumption or read/write request efficiency with FIPS-compatible builds.

CGO support for `vmagent`'s Kafka integration was disabled for simplification. This only affects the underlying library choice (Go-native vs. C-libraries) and does not impact functionality.

## Please Note

These VictoriaMetrics builds use FIPS 140-3 validated cryptographic modules (BoringCrypto).

However, VictoriaMetrics as a complete software solution has not undergone formal FIPS 140-3 certification under the CMVP.


## Get Started & Try VictoriaMetrics Enterprise

Ready to enhance your observability stack with FIPS 140-3 compatible components?

Request a VictoriaMetrics Enterprise [Trial License](https://victoriametrics.com/products/enterprise/trial/), or contact our [Sales team](https://victoriametrics.com/products/enterprise/) to get more information. 