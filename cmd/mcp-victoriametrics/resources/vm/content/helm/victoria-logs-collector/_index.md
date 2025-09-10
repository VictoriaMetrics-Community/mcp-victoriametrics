

---
weight: 2
title: VictoriaLogs Collector
menu:
  docs:
    parent: helm
    weight: 2
    identifier: helm-victorialogs-collector
url: /helm/victorialogs-collector
aliases:
  - /helm/victorialogs-collector/changelog/index.html
tags:
  - logs
  - kubernetes
---

![Version](https://img.shields.io/badge/0.0.4-gray?logo=Helm&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fhelm%2Fvictoria-logs-collector%2Fchangelog%2F%23004)
![ArtifactHub](https://img.shields.io/badge/ArtifactHub-informational?logoColor=white&color=417598&logo=artifacthub&link=https%3A%2F%2Fartifacthub.io%2Fpackages%2Fhelm%2Fvictoriametrics%2Fvictoria-logs-collector)
![License](https://img.shields.io/github/license/VictoriaMetrics/helm-charts?labelColor=green&label=&link=https%3A%2F%2Fgithub.com%2FVictoriaMetrics%2Fhelm-charts%2Fblob%2Fmaster%2FLICENSE)
![Slack](https://img.shields.io/badge/Join-4A154B?logo=slack&link=https%3A%2F%2Fslack.victoriametrics.com)
![X](https://img.shields.io/twitter/follow/VictoriaMetrics?style=flat&label=Follow&color=black&logo=x&labelColor=black&link=https%3A%2F%2Fx.com%2FVictoriaMetrics)
![Reddit](https://img.shields.io/reddit/subreddit-subscribers/VictoriaMetrics?style=flat&label=Join&labelColor=red&logoColor=white&logo=reddit&link=https%3A%2F%2Fwww.reddit.com%2Fr%2FVictoriaMetrics)

VictoriaLogs Collector - collects logs from Kubernetes containers and stores them to VictoriaLogs

## Prerequisites

Before installing this chart, ensure your environment meets the following requirements:

* **Kubernetes cluster** - A running Kubernetes cluster with sufficient resources
* **Helm** - Helm package manager installed and configured

Additional requirements depend on your configuration:

* **Persistent storage** - Required if you enable persistent volumes for data retention (enabled by default)
* **kubectl** - Needed for cluster management and troubleshooting

For installation instructions, refer to the official documentation:
* [Installing Helm](https://helm.sh/docs/intro/install/)
* [Installing kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Quick start

This Helm chart deploys a log collection agent as a DaemonSet.
It collects logs from all containers in a Kubernetes cluster and forwards them to the configured VictoriaLogs destinations.
If more than one destination is specified, then the collected logs are replicated among the configured destinations.

This chart will expand its functionality as the corresponding features are added to [vlagent](https://docs.victoriametrics.com/victorialogs/vlagent/).

- To quickly install single-node version of VictoriaLogs and `victoria-logs-collector`, see [these docs](https://docs.victoriametrics.com/helm/victoria-logs-single/#quick-start).
- To start with a VictoriaLogs cluster and `victoria-logs-collector`, see [these docs](https://docs.victoriametrics.com/helm/victoria-logs-cluster/#quick-start).

## Chart configuration

The simplest working configuration includes specifying the `remoteWrite` array and setting CPU and memory resources for the chart.

Example of a minimal working configuration:

```yaml
remoteWrite:
  - url: http://victoria-logs:9428

resources:
  limits:
    cpu: 100m
    memory: 128Mi
  requests:
    cpu: 100m
    memory: 128Mi
```

If multiple `remoteWrite` entries are defined, logs are replicated to all the specified destinations.

### Basic auth

If you need to use basic auth, define the secrets via environment variables and fill out the `basicAuth` object as shown below:

```yaml
remoteWrite:
  - url: http://victoria-logs:9428
    basicAuth:
      passwordEnvKey: 'VL_PASSWORD'
      usernameEnvKey: 'VL_USERNAME'

env:
  - name: VL_PASSWORD
    valueFrom:
      secretKeyRef:
        name: auth-secret
        key: VL_PASSWORD
  - name: VL_USERNAME
    valueFrom:
      secretKeyRef:
        name: auth-secret
        key: VL_USERNAME
```

### Multitenancy

To define [tenant](https://docs.victoriametrics.com/victorialogs/#multitenancy), use `projectID` and `accountID` as shown below:

```yaml
remoteWrite:
  - url: http://localhost:9428
    projectID: 12
    accountID: 42
```

### TLS

To enable TLS verification for the remoteWrite target, you can specify the `tls` block inside each remoteWrite entry.

At a minimum, you should provide the `caFile` path so that the collector can verify the server's TLS certificate.
This is useful when the target endpoint uses a certificate signed by a custom or self-signed Certificate Authority (CA).

```yaml
remoteWrite:
  - url: https://victoria-logs:9428
    tls:
      caFile: "/etc/tls/ca.crt"

extraVolumes:
  - name: tls-certs
    secret:
      secretName: tls-secret

extraVolumeMounts:
  - name: tls-certs
    mountPath: /etc/tls
    readOnly: true
```

If you want to disable TLS certificate verification (not recommended in production), you can set `insecureSkipVerify` to true.

This will skip verification of the server's certificate and allow connecting to targets with self-signed or invalid certificates.

```yaml
remoteWrite:
  - url: https://victoria-logs:9428
    tls:
      insecureSkipVerify: true
```

### Ignore fields

VictoriaLogs efficiently compresses repeated values, such as pod and node labels.
However, if you prefer not to store certain fields, you can ignore them using the `ignoreFields` option.
For example:

```yaml
remoteWrite:
  - url: http://victoria-logs:9428
    ignoreFields:
      - file
      - kubernetes.container_id
      - kubernetes.pod_annotations*
      - kubernetes.node_labels*
      - kubernetes.namespace_labels*
      - kubernetes.pod_labels*
```

This allows you to exclude unnecessary or sensitive fields from being ingested.
If sensitive data has already been ingested, see how to
[exclude logs from search result](https://docs.victoriametrics.com/victorialogs/security-and-lb/#access-control-inside-a-single-tenant).

### Extra fields

You can add custom fields to your logs by including the `extraFields` section in your configuration.
For example:

```yaml
remoteWrite:
  - url: http://victoria-logs:9428/
    extraFields:
      zone: us-east1-c
      source: victoria-logs-collector
```

This feature lets you attach metadata to every log entry,
making it easier to filter, group, or analyze logs based on these additional attributes.

## How to install

Access a Kubernetes cluster.

### Setup chart repository (can be omitted for OCI repositories)

Add a chart helm repository with follow commands:

```console
helm repo add vm https://victoriametrics.github.io/helm-charts/

helm repo update
```
List versions of `vm/victoria-logs-collector` chart available to installation:

```console
helm search repo vm/victoria-logs-collector -l
```

### Install `victoria-logs-collector` chart

Export default values of `victoria-logs-collector` chart to file `values.yaml`:

  - For HTTPS repository

    ```console
    helm show values vm/victoria-logs-collector > values.yaml
    ```
  - For OCI repository

    ```console
    helm show values oci://ghcr.io/victoriametrics/helm-charts/victoria-logs-collector > values.yaml
    ```

Change the values according to the need of the environment in ``values.yaml`` file.

> Consider setting `.Values.nameOverride` to a small value like `vlc` to avoid hitting resource name limits of 63 characters

Test the installation with command:

  - For HTTPS repository

    ```console
    helm install vlc vm/victoria-logs-collector -f values.yaml -n NAMESPACE --debug
    ```

  - For OCI repository

    ```console
    helm install vlc oci://ghcr.io/victoriametrics/helm-charts/victoria-logs-collector -f values.yaml -n NAMESPACE --debug
    ```

Install chart with command:

  - For HTTPS repository

    ```console
    helm install vlc vm/victoria-logs-collector -f values.yaml -n NAMESPACE
    ```

  - For OCI repository

    ```console
    helm install vlc oci://ghcr.io/victoriametrics/helm-charts/victoria-logs-collector -f values.yaml -n NAMESPACE
    ```

Get the pods lists by running this commands:

```console
kubectl get pods -A | grep 'vlc'
```

Get the application by running this command:

```console
helm list -f vlc -n NAMESPACE
```

See the history of versions of `vlc` application with command.

```console
helm history vlc -n NAMESPACE
```

## How to uninstall

Remove application with command.

```console
helm uninstall vlc -n NAMESPACE
```

## Documentation of Helm Chart

Install ``helm-docs`` following the instructions on this [tutorial](https://docs.victoriametrics.com/helm/requirements/).

Generate docs with ``helm-docs`` command.

```bash
cd charts/victoria-logs-collector

helm-docs
```

The markdown generation is entirely go template driven. The tool parses metadata from charts and generates a number of sub-templates that can be referenced in a template file (by default ``README.md.gotmpl``). If no template file is provided, the tool has a default internal template that will generate a reasonably formatted README.

## Parameters

The following tables lists the configurable parameters of the chart and their default values.

Change the values according to the need of the environment in ``victoria-logs-collector/values.yaml`` file.

<table class="helm-vars">
  <thead>
    <th class="helm-vars-key">Key</th>
    <th class="helm-vars-description">Description</th>
  </thead>
  <tbody>
    <tr id="affinity">
      <td><a href="#affinity"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">affinity</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Pod affinity</p>
</td>
    </tr>
    <tr id="annotations">
      <td><a href="#annotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">annotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Annotations to be added to the deployment</p>
</td>
    </tr>
    <tr id="env">
      <td><a href="#env"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">env</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Environment variables (ex.: secret tokens).</p>
</td>
    </tr>
    <tr id="extravolumemounts">
      <td><a href="#extravolumemounts"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">extraVolumeMounts</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Extra Volume Mounts for the container</p>
</td>
    </tr>
    <tr id="extravolumes">
      <td><a href="#extravolumes"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">extraVolumes</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Extra Volumes for the pod</p>
</td>
    </tr>
    <tr id="fullnameoverride">
      <td><a href="#fullnameoverride"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">fullnameOverride</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Override resources fullname</p>
</td>
    </tr>
    <tr id="global-cluster-dnsdomain">
      <td><a href="#global-cluster-dnsdomain"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">global.cluster.dnsDomain</span><span class="p">:</span><span class="w"> </span><span class="l">cluster.local.</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>K8s cluster domain suffix, used for building storage pods&rsquo; FQDN. Details are <a href="https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="global-compatibility">
      <td><a href="#global-compatibility"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">global.compatibility</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">openshift</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">        </span><span class="nt">adaptSecurityContext</span><span class="p">:</span><span class="w"> </span><span class="l">auto</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Openshift security context compatibility configuration</p>
</td>
    </tr>
    <tr id="global-image-registry">
      <td><a href="#global-image-registry"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">global.image.registry</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Image registry, that can be shared across multiple helm charts</p>
</td>
    </tr>
    <tr id="global-imagepullsecrets">
      <td><a href="#global-imagepullsecrets"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">global.imagePullSecrets</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Image pull secrets, that can be shared across multiple helm charts</p>
</td>
    </tr>
    <tr id="nameoverride">
      <td><a href="#nameoverride"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">nameOverride</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Override chart name</p>
</td>
    </tr>
    <tr id="nodeselector">
      <td><a href="#nodeselector"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">nodeSelector</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Pod&rsquo;s node selector. Details are <a href="https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="podannotations">
      <td><a href="#podannotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">podAnnotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Annotations to be added to pod</p>
</td>
    </tr>
    <tr id="podlabels">
      <td><a href="#podlabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">podLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Extra labels for Pods only</p>
</td>
    </tr>
    <tr id="podsecuritycontext">
      <td><a href="#podsecuritycontext"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">podSecurityContext</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Security context to be added to pod</p>
</td>
    </tr>
    <tr id="priorityclassname">
      <td><a href="#priorityclassname"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">priorityClassName</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Priority class to be assigned to the pod(s)</p>
</td>
    </tr>
    <tr id="remotewrite">
      <td><a href="#remotewrite"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">remoteWrite</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>List of log destinations. Logs will be replicated to all listed destinations.  If using a proxy (e.g., vmauth, nginx) in front of VictoriaLogs, make sure /insert/jsonline and /internal/insert endpoints are properly routed.</p>
</td>
    </tr>
    <tr id="resources">
      <td><a href="#resources"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">resources</span><span class="p">:</span><span class="w"> </span><span class="kc">null</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em></td>
    </tr>
    <tr id="securitycontext">
      <td><a href="#securitycontext"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">securityContext</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Security context to be added to pod&rsquo;s containers</p>
</td>
    </tr>
    <tr id="serviceaccount">
      <td><a href="#serviceaccount"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">serviceAccount</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">annotations</span><span class="p">:</span><span class="w"> </span>{}<span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">automount</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">name</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Service account is needed to enrich logs with pod metadata using Kubernetes API</p>
</td>
    </tr>
    <tr id="tolerations">
      <td><a href="#tolerations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">tolerations</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Node tolerations for server scheduling to nodes with taints. Details are <a href="https://kubernetes.io/docs/concepts/configuration/assign-pod-node/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="topologyspreadconstraints">
      <td><a href="#topologyspreadconstraints"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">topologySpreadConstraints</span><span class="p">:</span><span class="w"> </span><span class="kc">null</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Pod topologySpreadConstraints</p>
</td>
    </tr>
  </tbody>
</table>

