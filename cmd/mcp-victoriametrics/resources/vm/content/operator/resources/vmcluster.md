---
weight: 6
title: VMCluster
menu:
  docs:
    identifier: operator-cr-vmcluster
    parent: operator-cr
    weight: 6
aliases:
  - /operator/resources/vmcluster/
  - /operator/resources/vmcluster/index.html
tags:
  - kubernetes
  - metrics
---
`VMCluster` represents a high-available and fault-tolerant version of VictoriaMetrics database.
The `VMCluster` CRD defines a [cluster version VM](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/).

For each `VMCluster` resource, the Operator creates:

- `VMStorage` as `StatefulSet`,
- `VMSelect` as `StatefulSet`
- and `VMInsert` as deployment.

For `VMStorage` and `VMSelect` headless  services are created. `VMInsert` is created as service with clusterIP.

There is a strict order for these objects creation and reconciliation:

1. `VMStorage` is synced - the Operator waits until all its pods are ready;
1. Then it syncs `VMSelect` with the same manner;
1. `VMInsert` is the last object to sync.

All [statefulsets](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/) are created 
with [OnDelete](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#on-delete) update type. 
It allows to manually manage the rolling update process for Operator by deleting pods one by one and waiting for the ready status.

Rolling update process may be configured by the operator env variables.
The most important is `VM_PODWAITREADYTIMEOUT=80s` - it controls how long to wait for pod's ready status.

## Specification

You can see the full actual specification of the `VMCluster` resource in the **[API docs -> VMCluster](https://docs.victoriametrics.com/operator/api/#vmcluster)**.

If you can't find necessary field in the specification of the custom resource,
see [Extra arguments section](./#extra-arguments).

Also, you can check out the [examples](#examples) section.

## Requests Load-Balancing

 Operator provides enhanced load-balancing mechanism for `vminsert` and `vmselect` clients. By default, operator uses built-in Kubernetes [service]() with `clusterIP` type for clients connection. It's good solution for short lived connections. But it acts poorly with long-lived TCP sessions and leads to the uneven resources utilization for `vmselect` and `vminsert` components.

 Consider the following example:

![CR](vmcluster_default_balancer.webp)

 In this case clients could establish multiple connections to the same `pod` via `service`. And client requests will be served only by subset of `pods`.

 Operator allows to tweak this behaviour with enabled [requestsLoadBalancer](https://docs.victoriametrics.com/operator/api/#vmclusterspec-requestsloadbalancer):

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: with-balancer
spec:
  retentionPeriod: "4"
  replicationFactor: 1
  vminsert:
    replicaCount: 1
  vmselect:
    replicaCount: 1
  vmstorage:
    replicaCount: 1
  requestsLoadBalancer:
    enabled: true
    spec:
      replicaCount: 2
```

 Operator will deploy `VMAuth` deployment with 2 replicas. And update vminsert and vmselect services to point to `vmauth`.
 In addition, operator will create 3 additional services with the following pattern:

- vminsertinternal-CLUSTER_NAME - needed for vmselect pod discovery
- vmselectinternal-CLUSTER_NAME - needed for vminsert pod discovery
- vmclusterlb-CLUSTER_NAME - needed for metrics collection and exposing `vmselect` and `vminsert` components via `VMAuth` balancer.

 Network scheme with load-balancing:
 ![CR](vmcluster_with_balancer.webp)

The `requestsLoadBalancer` feature works transparently and is managed entirely by the `VMCluster` operator, 
with no direct access to the underlying [VMAuth](https://docs.victoriametrics.com/victoriametrics/vmauth/) configuration. 
If you need more control over load balancing behavior, 
or want to combine request routing with authentication or (m)TLS, 
consider deploying a standalone [VMAuth](https://docs.victoriametrics.com/operator/resources/vmauth/) resource instead of enabling `requestsLoadBalancer`.

## High availability

The cluster version provides a full set of high availability features - metrics replication, node failover, horizontal scaling.

First, we recommend familiarizing yourself with the high availability tools provided by "VictoriaMetrics Cluster" itself:

- [High availability](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#high-availability),
- [Cluster availability](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#cluster-availability),
- [Replication and data safety](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#replication-and-data-safety).

`VMCluster` supports all listed in the above-mentioned articles parameters and features:

- `replicationFactor` - the number of replicas for each metric.
- for every component of cluster (`vmstorage` / `vmselect` / `vminsert`):
  - `replicaCount` - the number of replicas for components of cluster.
  - `affinity` - the affinity (the pod's scheduling constraints) for components pods. See more details in [kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity).
  - `topologySpreadConstraints` - controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains. See more details in [kubernetes docs](https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/).

In addition, operator:

- uses k8s services or vmauth for load balancing between `vminsert` and `vmselect` components,
- uses health checks for to determine the readiness of components for work after restart,
- allows to horizontally scale all cluster components just by changing `replicaCount` field.

Here is an example of a `VMCluster` resource with HA features:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: example-persistent
spec:
  replicationFactor: 2       
  vmstorage:
    replicaCount: 10
    storageDataPath: "/vm-data"
    affinity:
      podAntiAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
        - labelSelector:
            matchExpressions:
            - key: "app.kubernetes.io/name"
              operator: In
              values:
              - "vmstorage"
          topologyKey: "kubernetes.io/hostname"
    storage:
      volumeClaimTemplate:
        spec:
          resources:
            requests:
              storage: 10Gi
    resources:
      limits:
        cpu: "2"
        memory: 2048Mi
  vmselect:
    replicaCount: 3
    cacheMountPath: "/select-cache"
    affinity:
      podAntiAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
        - labelSelector:
            matchExpressions:
            - key: "app.kubernetes.io/name"
              operator: In
              values:
              - "vmselect"
          topologyKey: "kubernetes.io/hostname"
    storage:
      volumeClaimTemplate:
        spec:
          resources:
            requests:
              storage: 2Gi
    resources:
      limits:
        cpu: "1"
        memory: "500Mi"
  vminsert:
    replicaCount: 4
    affinity:
      podAntiAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
        - labelSelector:
            matchExpressions:
            - key: "app.kubernetes.io/name"
              operator: In
              values:
              - "vminsert"
          topologyKey: "kubernetes.io/hostname"
    resources:
      limits:
        cpu: "1"
        memory: "500Mi"
```

## Version management

For `VMCluster` you can specify tag name from [releases](https://github.com/VictoriaMetrics/VictoriaMetrics/releases) and repository setting per cluster object:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: example
spec:
  vmstorage:
    replicaCount: 2
    image:
      repository: victoriametrics/vmstorage
      tag: v1.110.13-cluster
      pullPolicy: Always
  vmselect:
    replicaCount: 2
    image:
      repository: victoriametrics/vmselect
      tag: v1.110.13-cluster
      pullPolicy: Always
  vminsert:
    replicaCount: 2
    image:
      repository: victoriametrics/vminsert
      tag: v1.110.13-cluster
      pullPolicy: Always
```

or for all cluster components all together, using `clusterVersion` property:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: example
spec:
  clusterVersion: v1.110.13-cluster
```

Also, you can specify `imagePullSecrets` if you are pulling images from private repo, 
but `imagePullSecrets` is global setting for all `VMCluster` specification:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: example
spec:
  vmstorage:
    replicaCount: 2
    image:
      repository: victoriametrics/vmstorage
      tag: v1.110.13-cluster
      pullPolicy: Always
  vmselect:
    replicaCount: 2
    image:
      repository: victoriametrics/vmselect
      tag: v1.110.13-cluster
      pullPolicy: Always
  vminsert:
    replicaCount: 2
    image:
      repository: victoriametrics/vminsert
      tag: v1.110.13-cluster
      pullPolicy: Always
  imagePullSecrets:
    - name: my-repo-secret
  # ...
```

## Resource management

You can specify resources for each component of `VMCluster` resource in the `spec` section of the `VMCluster` CRD.

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: resources-example
spec:
    # ...
    vmstorage:
      resources:
          requests:
            memory: "16Gi"
            cpu: "4"
          limits:
            memory: "16Gi"
            cpu: "4"
    # ...
    vmselect:
      resources:
        requests:
          memory: "16Gi"
          cpu: "4"
        limits:
          memory: "16Gi"
          cpu: "4"
    # ...
    vminsert:
      resources:
        requests:
          memory: "16Gi"
          cpu: "4"
        limits:
          memory: "16Gi"
          cpu: "4"
  # ...
```

If these parameters are not specified, then,
by default all `VMCluster` pods have resource requests and limits from the default values of the following [operator parameters](https://docs.victoriametrics.com/operator/configuration/):

- `VM_VMCLUSTERDEFAULT_VMSTORAGEDEFAULT_RESOURCE_LIMIT_MEM` - default memory limit for `VMCluster/vmstorage` pods,
- `VM_VMCLUSTERDEFAULT_VMSTORAGEDEFAULT_RESOURCE_LIMIT_CPU` - default memory limit for `VMCluster/vmstorage` pods,
- `VM_VMCLUSTERDEFAULT_VMSTORAGEDEFAULT_RESOURCE_REQUEST_MEM` - default memory limit for `VMCluster/vmstorage` pods,
- `VM_VMCLUSTERDEFAULT_VMSTORAGEDEFAULT_RESOURCE_REQUEST_CPU` - default memory limit for `VMCluster/vmstorage` pods,
- `VM_VMCLUSTERDEFAULT_VMSELECTDEFAULT_RESOURCE_LIMIT_MEM` - default memory limit for `VMCluster/vmselect` pods,
- `VM_VMCLUSTERDEFAULT_VMSELECTDEFAULT_RESOURCE_LIMIT_CPU` - default memory limit for `VMCluster/vmselect` pods,
- `VM_VMCLUSTERDEFAULT_VMSELECTDEFAULT_RESOURCE_REQUEST_MEM` - default memory limit for `VMCluster/vmselect` pods,
- `VM_VMCLUSTERDEFAULT_VMSELECTDEFAULT_RESOURCE_REQUEST_CPU` - default memory limit for `VMCluster/vmselect` pods,
- `VM_VMCLUSTERDEFAULT_VMINSERTDEFAULT_RESOURCE_LIMIT_MEM` - default memory limit for `VMCluster/vmselect` pods,
- `VM_VMCLUSTERDEFAULT_VMINSERTDEFAULT_RESOURCE_LIMIT_CPU` - default memory limit for `VMCluster/vmselect` pods,
- `VM_VMCLUSTERDEFAULT_VMINSERTDEFAULT_RESOURCE_REQUEST_MEM` - default memory limit for `VMCluster/vmselect` pods,
- `VM_VMCLUSTERDEFAULT_VMINSERTDEFAULT_RESOURCE_REQUEST_CPU` - default memory limit for `VMCluster/vmselect` pods.

These default parameters will be used if:

- `VM_VMCLUSTERDEFAULT_USEDEFAULTRESOURCES` is set to `true` (default value),
- `VMCluster/*` CR doesn't have `resources` field in `spec` section.

Field `resources` in `VMCluster/*` spec have higher priority than operator parameters.

If you set `VM_VMCLUSTERDEFAULT_USEDEFAULTRESOURCES` to `false` and don't specify `resources` in `VMCluster/*` CRD,
then `VMCluster/*` pods will be created without resource requests and limits.

Also, you can specify requests without limits - in this case default values for limits will not be used.

## Enterprise features

VMCluster supports following features 
from [VictoriaMetrics Enterprise](https://docs.victoriametrics.com/victoriametrics/enterprise/#victoriametrics-enterprise):

- [Downsampling](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#downsampling)
- [Multiple retentions / Retention filters](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#retention-filters)
- [Advanced per-tenant statistic](https://docs.victoriametrics.com/victoriametrics/pertenantstatistic/)
- [mTLS for cluster components](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#mtls-protection)
- [Backup automation](https://docs.victoriametrics.com/victoriametrics/vmbackupmanager/)

VMCluster doesn't support yet feature 
[Automatic discovery for vmstorage nodes](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#automatic-vmstorage-discovery).

For using Enterprise version of [vmcluster](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/) you need to:
 - specify license at [`spec.license.key`](https://docs.victoriametrics.com/operator/api/#license-key) or at [`spec.license.keyRef`](https://docs.victoriametrics.com/operator/api/#license-keyref).
 - change version of `vmcluster` to version with `-enterprise-cluster` suffix using [Version management](#version-management).

### Downsampling

After that you can pass [Downsampling](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#downsampling)
flag to `VMCluster/vmselect` and `VMCluster/vmstorage` with [extraArgs](./#extra-arguments) too.

Here are complete example for [Downsampling](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#downsampling):

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: ent-example
spec:
  # enabling enterprise features
  license:
    keyRef:
      name: k8s-secret-that-contains-license
      key: key-in-a-secret-that-contains-license
  clusterVersion: v1.110.13-enterprise-cluster
  vmselect:
    # enabling enterprise features for vmselect
    extraArgs:
      # using enterprise features: Downsampling
      # more details about downsampling you can read on https://docs.victoriametrics.com/victoriametrics/cluster-victoriaMetrics/#downsampling
      downsampling.period: 30d:5m,180d:1h,1y:6h,2y:1d
  vmstorage:
    # enabling enterprise features for vmstorage
    extraArgs:
      # using enterprise features: Downsampling
      # more details about downsampling you can read on https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#downsampling
      downsampling.period: 30d:5m,180d:1h,1y:6h,2y:1d

  # ...other fields...
```

### Retention filters

You can pass [Retention filters](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#retention-filters)
flag to  `VMCluster/vmstorage` with [extraArgs](./#extra-arguments).

Here are complete example for [Retention filters](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#retention-filters):

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: ent-example
spec:
  # enabling enterprise features
  license:
    keyRef:
      name: k8s-secret-that-contains-license
      key: key-in-a-secret-that-contains-license
  clusterVersion: v1.110.13-enterprise-cluster
  vmstorage:
    extraArgs:
      # using enterprise features: Retention filters
      # more details about retention filters you can read on https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#retention-filters
      retentionFilter: '{vm_account_id="5",env="dev"}:5d,{vm_account_id="5",env="prod"}:5y'

  # ...other fields...
```

### Advanced per-tenant statistic

For using [Advanced per-tenant statistic](https://docs.victoriametrics.com/victoriametrics/pertenantstatistic/)
you only need to [enable Enterprise version of vmcluster components](#enterprise-features) 
and operator will automatically create 
[Scrape objects](https://docs.victoriametrics.com/operator/resources/vmagent/#scraping) for cluster components.

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: ent-example
spec:
  # enabling enterprise features
  license:
    keyRef:
      name: k8s-secret-that-contains-license
      key: key-in-a-secret-that-contains-license
  clusterVersion: v1.110.13-enterprise-cluster

  # ...other fields...
```

After that [VMAgent](https://docs.victoriametrics.com/operator/resources/vmagent/) will automatically
scrape [Advanced per-tenant statistic](https://docs.victoriametrics.com/victoriametrics/pertenantstatistic/) for cluster components.

### mTLS protection

You can pass [mTLS protection](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#mtls-protection)
flags to `VMCluster/vmstorage`, `VMCluster/vmselect` and `VMCluster/vminsert` with [extraArgs](./#extra-arguments) and mount secret files
with `extraVolumes` and `extraVolumeMounts` fields.

Here are complete example for [mTLS protection](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#mtls-protection)

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: ent-example
spec:
  # enabling enterprise features
  license:
    keyRef:
      name: k8s-secret-that-contains-license
      key: key-in-a-secret-that-contains-license
  clusterVersion: v1.110.13-enterprise-cluster
  vmselect:
    extraArgs:
      # using enterprise features: mTLS protection
      # more details about mTLS protection you can read on https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#mtls-protection
      cluster.tls: true
      cluster.tlsCAFile: /etc/mtls/ca.crt
      cluster.tlsCertFile: /etc/mtls/vmselect.crt
      cluster.tlsKeyFile: /etc/mtls/vmselect.key
    extraVolumes:
      - name: mtls
        secret:
          secretName: mtls
    extraVolumeMounts:
      - name: mtls
        mountPath: /etc/mtls
      
  vminsert:
    extraArgs:
      # using enterprise features: mTLS protection
      # more details about mTLS protection you can read on https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#mtls-protection
      cluster.tls: true
      cluster.tlsCAFile: /etc/mtls/ca.crt
      cluster.tlsCertFile: /etc/mtls/vminsert.crt
      cluster.tlsKeyFile: /etc/mtls/vminsert.key
    extraVolumes:
      - name: mtls
        secret:
          secretName: mtls
    extraVolumeMounts:
      - name: mtls
        mountPath: /etc/mtls
      
  vmstorage:
    env:
      - name: POD
        valueFrom:
          fieldRef:
            fieldPath: metadata.name
    extraArgs:
      # using enterprise features: mTLS protection
      # more details about mTLS protection you can read on https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/#mtls-protection
      cluster.tls: true
      cluster.tlsCAFile: /etc/mtls/ca.crt
      cluster.tlsCertFile: /etc/mtls/$(POD).crt
      cluster.tlsKeyFile: /etc/mtls/$(POD).key
    extraVolumes:
      - name: mtls
        secret:
          secretName: mtls
    extraVolumeMounts:
      - name: mtls
        mountPath: /etc/mtls

  # ...other fields...

---

apiVersion: v1
kind: Secret
metadata:
  name: mtls
  namespace: default
stringData:
  ca.crt: |
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
  mtls-vmstorage-0.crt: |
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
  mtls-vmstorage-0.key: |
    -----BEGIN PRIVATE KEY-----
    ...
    -----END PRIVATE KEY-----
  mtls-vmstorage-1.crt: |
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
  mtls-vmstorage-1.key: |
    -----BEGIN PRIVATE KEY-----
    ...
    -----END PRIVATE KEY-----
  vminsert.crt: |
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
  vminsert.key: |
    -----BEGIN PRIVATE KEY-----
    ...
    -----END PRIVATE KEY-----
  vmselect.crt: |
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
  vmselect.key: |
    -----BEGIN PRIVATE KEY-----
    ...
    -----END PRIVATE KEY-----

```

Example commands for generating certificates you can read 
on [this page](https://gist.github.com/f41gh7/76ed8e5fb1ebb9737fe746bae9175ee6#generate-self-signed-ca-with-key).

### Backup automation

You can check [vmbackupmanager documentation](https://docs.victoriametrics.com/victoriametrics/vmbackupmanager/) for backup automation.
It contains a description of the service and its features. This section covers vmbackumanager integration in vmoperator.

`VMCluster` has built-in backup configuration, it uses `vmbackupmanager` - proprietary tool for backups.
It supports incremental backups (hourly, daily, weekly, monthly) with popular object storages (aws s3, google cloud storage).

Here is a complete example for backup configuration:

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: ent-example
spec:
  vmstorage:
    vmBackup:
      # this feature is only available in Victoriametrics Enterprise
      # more details about backup automation you can read on https://docs.victoriametrics.com/victoriametrics/vmbackupmanager/
      destination: "s3://your_bucket/folder"
      # Read the object storage credentials from a secret
      credentialsSecret:
        name: remote-storage-keys
        key: credentials
      # customS3Endpoint: 'https://s3.example.com' # uncomment and adjust if you using s3 compatible storage instead of AWS s3
      # uncomment and adjust to fit your backup schedule
      # disableHourly: false
      # disableDaily: false
      # disableWeekly: false
      # disableMonthly: false
  # ...other fields...

---

apiVersion: v1
kind: Secret
metadata:
  name: remote-storage-keys
type: Opaque
stringData:
  credentials: |-
    [default]
    aws_access_key_id = your_access_key_id
    aws_secret_access_key = your_secret_access_key
```

**NOTE**: for cluster version operator adds suffix for destination: `"s3://your_bucket/folder"`, it becomes `"s3://your_bucket/folder/$(POD_NAME)"`.
It's needed to make consistent backups for each storage node.

You can read more about backup configuration options and mechanics [here](https://docs.victoriametrics.com/victoriametrics/vmbackupmanager/)

Possible configuration options for backup crd can be found at [link](https://docs.victoriametrics.com/operator/api/#vmbackup)

**Using VMBackupmanager for restoring backups** in Kubernetes environment is described [here](https://docs.victoriametrics.com/victoriametrics/vmbackupmanager/#how-to-restore-in-kubernetes).

Also see VMCluster example spec [here](https://github.com/VictoriaMetrics/operator/blob/master/config/examples/vmcluster_with_backuper.yaml).

## Examples

### Minimal example without persistence

```yaml
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMCluster
metadata:
  name: example-minimal
spec:
  # ...
  retentionPeriod: "1"
  vmstorage:
    replicaCount: 2
  vmselect:
    replicaCount: 2
  vminsert:
    replicaCount: 2
```

### With persistence

```yaml
kind: VMCluster
metadata:
  name: example-persistent
spec:
  # ...
  retentionPeriod: "4"
  replicationFactor: 2
  vmstorage:
    replicaCount: 2
    storageDataPath: "/vm-data"
    storage:
      volumeClaimTemplate:
        spec:
          storageClassName: standard
          resources:
            requests:
              storage: 10Gi
    resources:
      limits:
        cpu: "0.5"
        memory: 500Mi
  vmselect:
    replicaCount: 2
    cacheMountPath: "/select-cache"
    storage:
      volumeClaimTemplate:
        spec:
          resources:
            requests:
              storage: 2Gi
    resources:
      limits:
        cpu: "0.3"
        memory: "300Mi"
  vminsert:
    replicaCount: 2
```
