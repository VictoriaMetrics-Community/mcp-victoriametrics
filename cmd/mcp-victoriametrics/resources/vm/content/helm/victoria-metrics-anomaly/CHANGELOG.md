## Next release

- TODO

## 1.11.2

**Release date:** 30 Jul 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.25.2](https://img.shields.io/badge/v1.25.2-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1252)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.25.2](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1252).

## 1.11.1

**Release date:** 24 Jul 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.25.1](https://img.shields.io/badge/v1.25.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1251)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.25.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1251).

## 1.11.0

**Release date:** 17 Jul 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.25.0](https://img.shields.io/badge/v1.25.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1250)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.25.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1250).
- Remove config has from pod annotations, when anomaly detection version is <1.25.0 to support hot-reload

## 1.10.2

**Release date:** 07 Jul 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.24.1](https://img.shields.io/badge/v1.24.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1241)

- Support `.Values.topologySpreadConstraints` property. See [#2219](https://github.com/VictoriaMetrics/helm-charts/issues/2219)
- Add support of using [`VMPodScrape`](https://docs.victoriametrics.com/operator/resources/vmpodscrape/) for monitoring configuration. See `.Values.podMonitor.vm`. 

## 1.10.1

**Release date:** 20 Jun 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.24.1](https://img.shields.io/badge/v1.24.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1241)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.24.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1241). Upgrading from [1.23.0] - [1.24.0] **is recommended to a critical bug fixed**, see [v1.23.0 note](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1230) for details.

## 1.10.0

**Release date:** 19 Jun 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.24.0](https://img.shields.io/badge/v1.24.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1240)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.24.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1240).
- Automatically enable [restoring state](http://docs.victoriametrics.com/anomaly-detection/components/settings/#state-restoration) when persistence is enabled via `.Values.persistentVolume.enabled`.

## 1.9.6

**Release date:** 13 Jun 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.23.3](https://img.shields.io/badge/v1.23.3-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1233)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.23.3](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1233).


## 1.9.5

**Release date:** 09 Jun 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.23.2](https://img.shields.io/badge/v1.23.2-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1232)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.23.2](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1232).

## 1.9.4

**Release date:** 09 Jun 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.23.1](https://img.shields.io/badge/v1.23.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1231)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.23.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1231).

## 1.9.3

**Release date:** 05 Jun 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.23.0](https://img.shields.io/badge/v1.23.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1230)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.23.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1230).

## 1.9.2

**Release date:** 11 May 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.22.1](https://img.shields.io/badge/v1.22.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1221)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.22.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1221).

## 1.9.1

**Release date:** 18 Apr 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.22.0-experimental](https://img.shields.io/badge/v1.22.0--experimental-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1220)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.22.0-experimental](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1220-experimental).

## 1.9.0

**Release date:** 19 Mar 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.21.0](https://img.shields.io/badge/v1.21.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1210)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.21.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1210).
- Add support of sharding and replication via `.Values.shardsCount` and `.Values.replicationFactor`. See [these docs](https://docs.victoriametrics.com/anomaly-detection/faq/index.html#scaling-vmanomaly) for the details
- updated common dependency 0.0.39 -> 0.0.42

## 1.8.1

**Release date:** 16 Mar 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.20.1](https://img.shields.io/badge/v1.20.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1201)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.20.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1201).

## 1.8.0

**Release date:** 04 Mar 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.20.0](https://img.shields.io/badge/v1.20.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1200)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.20.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1200).
- updated common dependency 0.0.37 -> 0.0.39

## 1.7.2

**Release date:** 27 Jan 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.19.2](https://img.shields.io/badge/v1.19.2-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1192)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.19.2](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1192).

## 1.7.1

**Release date:** 21 Jan 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.19.1](https://img.shields.io/badge/v1.19.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1191)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.19.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1191).

## 1.7.0

**Release date:** 21 Jan 2025

![Helm: v3](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0) ![AppVersion: v1.19.0](https://img.shields.io/badge/v1.19.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%23v1190)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.19.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1190).
- updated common dependency 0.0.34 -> 0.0.37
- Exclude markdown files from package
- support templating in `.Values.extraObjects`

## 1.6.11

**Release date:** 2024-12-03

![AppVersion: v1.18.8](https://img.shields.io/badge/v1.18.8-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1188)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.8](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1188).

## 1.6.10

**Release date:** 2024-12-02

![AppVersion: v1.18.7](https://img.shields.io/badge/v1.18.7-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1187)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.7](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1187).

## 1.6.9

**Release date:** 2024-12-01

![AppVersion: v1.18.6](https://img.shields.io/badge/v1.18.6-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1186)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.6](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1186).
- updated common dependency 0.0.32 -> 0.0.33

## 1.6.8

**Release date:** 2024-11-27

![AppVersion: v1.18.5](https://img.shields.io/badge/v1.18.5-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1185)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.5](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1185).
- updated common dependency 0.0.28 -> 0.0.32
- fixed app.kubernetes.io/version tag override if custom tag is set. See [this issue](https://github.com/VictoriaMetrics/helm-charts/issues/1766).

## 1.6.7

**Release date:** 2024-11-18

![AppVersion: v1.18.4](https://img.shields.io/badge/v1.18.4-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1184)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.4](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1184).
- updated common dependency 0.0.23 -> 0.0.28

## 1.6.6

**Release date:** 2024-11-14

![AppVersion: v1.18.3](https://img.shields.io/badge/v1.18.3-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1183)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded ['vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.3](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1183). This is a patch release that fixes a service crash during parallelized data processing with [VmReader](https://docs.victoriametrics.com/anomaly-detection/components/reader/#vm-reader).


## 1.6.5

**Release date:** 2024-11-13

![AppVersion: v1.18.2](https://img.shields.io/badge/v1.18.2-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1182)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded [`vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.2](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1182)

## 1.6.4

**Release date:** 2024-11-12

![AppVersion: v1.18.1](https://img.shields.io/badge/v1.18.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1181)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded [`vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1181)
- use common templates
- set default DNS domain to `cluster.local.`
- added podLabels and podAnnotations to add extra pod labels and annotations
- updated common dependency 0.0.19 -> 0.0.23

## 1.6.3

**Release date:** 2024-10-28

![AppVersion: v1.18.0](https://img.shields.io/badge/v1.18.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1180)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded [`vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.18.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1180)

## 1.6.2

**Release date:** 2024-10-22

![AppVersion: v1.17.2](https://img.shields.io/badge/v1.17.2-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1172)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded [`vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.17.2](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1172)

## 1.6.1

**Release date:** 2024-10-18

![AppVersion: v1.17.1](https://img.shields.io/badge/v1.17.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1171)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded [`vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.17.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1171)

## 1.6.0

**Release date:** 2024-10-17

![AppVersion: v1.17.0](https://img.shields.io/badge/v1.17.0-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1170)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded [`vmanomaly`](https://docs.victoriametrics.com/anomaly-detection/) to [1.17.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1170)

## 1.5.2

**Release date:** 2024-10-11

![AppVersion: v1.16.1](https://img.shields.io/badge/v1.16.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1161)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Human-readable error about Helm version requirement

## 1.5.1

**Release date:** 2024-10-04

![AppVersion: v1.16.1](https://img.shields.io/badge/v1.16.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1161)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded common chart dependency

## 1.5.0

**Release date:** 2024-10-03

![AppVersion: v1.16.1](https://img.shields.io/badge/v1.16.1-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1161)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Upgraded vmanomaly to [1.16.1](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1161)
- Added the ability to enable persistence for models and data via `.Values.persistentVolume.dumpModels` and `.Values.persistentVolume.dumpData` variables respectively.
- Fix default `podSecurityContext` configuration to ensure fs group matches container user.
- Fix passing empty `tenant_id` in case tenant is not defined in values.

## 1.4.6

**Release date:** 2024-09-16

![AppVersion: v1.15.9](https://img.shields.io/badge/v1.15.9-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1159)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Added the ability to add annotations to the configMap using `values.configMapAnnotations`
- Fixed license file flag name

## 1.4.5

**Release date:** 2024-09-12

![AppVersion: v1.15.9](https://img.shields.io/badge/v1.15.9-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1159)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Added ability to override deployment namespace using `namespaceOverride` and `global.namespaceOverride` variables
- Removed vmanomaly not existing `loggerFormat` extra arg. See [this issue](https://github.com/VictoriaMetrics/helm-charts/issues/1476)

## 1.4.4

**Release date:** 2024-09-03

![AppVersion: v1.15.9](https://img.shields.io/badge/v1.15.9-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1159)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Fixed PVC in StatefulSet

## 1.4.3

**Release date:** 2024-08-27

![AppVersion: v1.15.9](https://img.shields.io/badge/v1.15.9-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1159)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Removed `eula` support
- Disable PodMonitor, when pull port is not defined
- Upgraded application version to 1.15.9
- Fixed default podDisruptionBudget configuration

## 1.4.2

**Release date:** 2024-08-26

![AppVersion: v1.15.6](https://img.shields.io/badge/v1.15.6-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1156)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Set minimal kubernetes version to `1.25`
- Added `.Values.global.imagePullSecrets` and `.Values.global.image.registry`
- Fixed volume template. See [this issue](https://github.com/VictoriaMetrics/helm-charts/issues/1280)
- Fixed image pull secrets. See [this issue](https://github.com/VictoriaMetrics/helm-charts/issues/1285)
- Renamed `.Values.persistentVolume.storageClass` to `.Values.persistentVolume.storageClassName`
- Removed necessity to set `.Values.persistentVolume.existingClaim` when it should be created by chart. See [this issue](https://github.com/VictoriaMetrics/helm-charts/issues/189)
- Added PDB, PodMonitor, extra volumes and extra volumeMounts

## 1.4.1

**Release date:** 2024-08-15

![AppVersion: v1.15.4](https://img.shields.io/badge/v1.15.4-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1154)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Update vmanomaly to [v1.15.4](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1154).

## 1.4.0

**Release date:** 2024-08-14

![AppVersion: v1.15.3](https://img.shields.io/badge/v1.15.3-success?logo=VictoriaMetrics&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fanomaly-detection%2Fchangelog%2F%23v1153)
![Helm: v3.14](https://img.shields.io/badge/Helm-v3.14%2B-informational?color=informational&logo=helm&link=https%3A%2F%2Fgithub.com%2Fhelm%2Fhelm%2Freleases%2Ftag%2Fv3.14.0)

- Update vmanomaly to [v1.15.3](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1153).
- Update configuration example format to match the latest version of vmanomaly.

## 1.3.4

**Release date:** 2024-07-19

![AppVersion: v1.13.3](https://img.shields.io/static/v1?label=AppVersion&message=v1.13.3&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- leave empty `schedulers` and `models` section to fix aliases error

## 1.3.3

**Release date:** 2024-07-17

![AppVersion: v1.13.2](https://img.shields.io/static/v1?label=AppVersion&message=v1.13.2&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- fix default value for `.Values.config.schedulers.class`.

## 1.3.2

**Release date:** 2024-07-17

![AppVersion: v1.13.2](https://img.shields.io/static/v1?label=AppVersion&message=v1.13.2&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- changes made for vmanomaly [v1.13.2](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1132)

## 1.3.1

**Release date:** 2024-07-08

![AppVersion: v1.13.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.13.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- add missing API version and kind for volumeClaimTemplates, see [this issue](https://github.com/VictoriaMetrics/helm-charts/issues/1092).

## 1.3.0

**Release date:** 2024-06-11

![AppVersion: v1.13.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.13.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- Add ability to configure persistent volume for vmanomaly models storage.
- Fix `.Values.podSecurityContext` not being applied to the pod.
- Update vmanomaly to [v1.13.0](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1130).

## 1.2.4

**Release date:** 2024-05-16

![AppVersion: v1.12.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.12.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- fix lost customized securityContext when introduced new default behavior for securityContext in [pull request](https://github.com/VictoriaMetrics/helm-charts/pull/995).

## 1.2.3

**Release date:** 2024-05-10

![AppVersion: v1.12.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.12.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- support disabling default securityContext to keep compatible with platform like openshift, see this [pull request](https://github.com/VictoriaMetrics/helm-charts/pull/995) by @Baboulinet-33 for details.

## 1.2.2

**Release date:** 2024-04-02

![AppVersion: v1.12.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.12.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- apply [v1.12](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1120) as a default (no config changes).

## 1.2.1

**Release date:** 2024-03-20

![AppVersion: v1.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.11.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- Add support of passing preset configuration.

## 1.2.0

**Release date:** 2024-02-26

![AppVersion: v1.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.11.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- apply [v1.11](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1110) change in [schedulers section](https://docs.victoriametrics.com/anomaly-detection/components/scheduler/): add configuration for using multiple schedulers at once via `schedulers`. Old `scheduler` field is deprecated and will be automatically converted to `schedulers` definition starting from [v1.11](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1110).
- docs fixes

## 1.1.1

**Release date:** 2024-02-20

![AppVersion: v1.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.10.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- Fix passing path to license file location when using `license.secret` mount.

## 1.1.0

**Release date:** 2024-02-19

![AppVersion: v1.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.10.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- apply [v1.10](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1100) change in [models section](https://docs.victoriametrics.com/anomaly-detection/components/models/): add configuration for using multiple models at once via `models`. Old `model` field is deprecated and will be automatically converted to `models` definition starting from [v1.10](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1100).
- docs fixes

## 1.0.0

**Release date:** 2024-02-05

![AppVersion: v1.9.2](https://img.shields.io/static/v1?label=AppVersion&message=v1.9.2&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- Breaking change: passing [full vmanomaly config](https://docs.victoriametrics.com/anomaly-detection/components/) via `config` parameter.
- vmanomaly image moving to DockerHub

## 0.5.0

**Release date:** 2023-10-31

![AppVersion: v1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.6.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- Add options to use `bearer_token` for reader and writer authentication.
- Add `verify_tls` option to bypass TLS verification for reader and writer.
- Add `extra_filters` option to supply additional filters to enforce for reader queries.

## 0.4.1

**Release date:** 2023-10-10

![AppVersion: v1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

- Add an options to override default `metric_format` for remote write configuration of vmanomaly.

## 0.4.0

**Release date:** 2023-08-21

![AppVersion: v1.93.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.93.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add ability to provide license key

## 0.3.5

**Release date:** 2023-06-22

![AppVersion: v1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.1.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* bump version of vmanomaly
* charts/victoria-metrics-anomaly: fix monitoring config indentation (#567)

## 0.3.4

**Release date:** 2023-06-22

![AppVersion: v1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.1.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* bump vmanomaly remove tricky make command
* charts/victoria-metrics-anomaly: make monitoring config more configurable (#562)

## 0.3.3

**Release date:** 2023-06-07

![AppVersion: v1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.1.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* bump anomaly chart, make package make merge

## 0.3.2

**Release date:** 2023-06-06

![AppVersion: v1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.1.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Anomaly: change defaults (#518)
* charts/operator: update version to 0.30.4 adds extraArgs and serviceMonitor options for operator
* vmanomaly re-release

## 0.3.1

**Release date:** 2023-01-26

![AppVersion: v1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.1.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* vmanomaly: fix monitoring part of config (#457)

## 0.3.0

**Release date:** 2023-01-24

![AppVersion: v1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.1.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* release vmanomaly v1.1.0 (#454)
* vmanomaly: fix config for pull-based monitoring (#446)
