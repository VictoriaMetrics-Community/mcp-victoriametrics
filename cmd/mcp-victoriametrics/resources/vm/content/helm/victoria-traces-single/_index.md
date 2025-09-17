

---
weight: 15
title: VictoriaTraces Single
menu:
  docs:
    parent: helm
    weight: 15
    identifier: helm-victoria-traces-single
url: /helm/victoria-traces-single/
aliases:
  - /helm/victoriatraces-single/
tags:
  - traces
  - kubernetes
---

![Version](https://img.shields.io/badge/0.0.1-gray?logo=Helm&labelColor=gray&link=https%3A%2F%2Fdocs.victoriametrics.com%2Fhelm%2Fvictoria-traces-single%2Fchangelog%2F%23001)
![ArtifactHub](https://img.shields.io/badge/ArtifactHub-informational?logoColor=white&color=417598&logo=artifacthub&link=https%3A%2F%2Fartifacthub.io%2Fpackages%2Fhelm%2Fvictoriametrics%2Fvictoria-traces-single)
![License](https://img.shields.io/github/license/VictoriaMetrics/helm-charts?labelColor=green&label=&link=https%3A%2F%2Fgithub.com%2FVictoriaMetrics%2Fhelm-charts%2Fblob%2Fmaster%2FLICENSE)
![Slack](https://img.shields.io/badge/Join-4A154B?logo=slack&link=https%3A%2F%2Fslack.victoriametrics.com)
![X](https://img.shields.io/twitter/follow/VictoriaMetrics?style=flat&label=Follow&color=black&logo=x&labelColor=black&link=https%3A%2F%2Fx.com%2FVictoriaMetrics)
![Reddit](https://img.shields.io/reddit/subreddit-subscribers/VictoriaMetrics?style=flat&label=Join&labelColor=red&logoColor=white&logo=reddit&link=https%3A%2F%2Fwww.reddit.com%2Fr%2FVictoriaMetrics)

The VictoriaTraces single Helm chart deploys VictoriaTraces database in Kubernetes.

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

The chart deploys VictoriaTraces database in StatefulSet mode.
For a quick start, install `victoria-traces-single` chart using the following commands.
Make sure to replace the environment variables with your own values:

```sh
export RETENTION=30d
export PVC_SIZE=10Gi
export NAMESPACE=tracing

kubectl create namespace $NAMESPACE

# Install victoria-traces-single chart
helm install vts vm/victoria-traces-single --namespace $NAMESPACE --wait \
    --set "server.retentionPeriod=$RETENTION" --set "server.persistentVolume.size=$PVC_SIZE"

For finer control and easier maintenance, it is recommended to set these
values in a separate `values` file and use it during installation.
See [how to install victoria-traces-single](https://docs.victoriametrics.com/helm/victoria-traces-single/#install-victoria-traces-single-chart) for an example.
You can do this later if you want to configure more settings than shown in the example.

## Chart configuration

The chart provides the following configuration options:

- Persistent volume size for storing the ingested traces. See `server.persistentVolume.size` config option at the [minimal configuration example](#minimal-configuration).
- Data retention (how long to keep traces in the database). It can be configurated via time-based retention (with units of days, weeks, or years)
  or via disk space-based retention limits. See `server.retentionPeriod` and `server.retentionDiskSpaceUsage` config options in the [minimal configuration example](#minimal-configuration).
  See [retention docs](https://docs.victoriametrics.com/victoriatraces/#retention) for details.

### Minimal Configuration

You can install the chart right away without any configuration (see [how to install](#how-to-install)), but here is the minimal configuration that is commonly customized:

```yaml
# VictoriaTraces server
server:

  # Time-based retention
  retentionPeriod: 7d

  # Or disk-based retention
  # retentionDiskSpaceUsage: 5GB

  # Storage size for the ingested traces
  persistentVolume:
    size: 20Gi
```

## How to install

Access a Kubernetes cluster.

### Setup chart repository (can be omitted for OCI repositories)

Add a chart helm repository with follow commands:

```console
helm repo add vm https://victoriametrics.github.io/helm-charts/

helm repo update
```
List versions of `vm/victoria-traces-single` chart available to installation:

```console
helm search repo vm/victoria-traces-single -l
```

### Install `victoria-traces-single` chart

Export default values of `victoria-traces-single` chart to file `values.yaml`:

  - For HTTPS repository

    ```console
    helm show values vm/victoria-traces-single > values.yaml
    ```
  - For OCI repository

    ```console
    helm show values oci://ghcr.io/victoriametrics/helm-charts/victoria-traces-single > values.yaml
    ```

Change the values according to the need of the environment in ``values.yaml`` file.

> Consider setting `.Values.nameOverride` to a small value like `vts` to avoid hitting resource name limits of 63 characters

Test the installation with command:

  - For HTTPS repository

    ```console
    helm install vts vm/victoria-traces-single -f values.yaml -n NAMESPACE --debug
    ```

  - For OCI repository

    ```console
    helm install vts oci://ghcr.io/victoriametrics/helm-charts/victoria-traces-single -f values.yaml -n NAMESPACE --debug
    ```

Install chart with command:

  - For HTTPS repository

    ```console
    helm install vts vm/victoria-traces-single -f values.yaml -n NAMESPACE
    ```

  - For OCI repository

    ```console
    helm install vts oci://ghcr.io/victoriametrics/helm-charts/victoria-traces-single -f values.yaml -n NAMESPACE
    ```

Get the pods lists by running this commands:

```console
kubectl get pods -A | grep 'vts'
```

Get the application by running this command:

```console
helm list -f vts -n NAMESPACE
```

See the history of versions of `vts` application with command.

```console
helm history vts -n NAMESPACE
```

## How to uninstall

Remove application with command.

```console
helm uninstall vts -n NAMESPACE
```

## Documentation of Helm Chart

Install ``helm-docs`` following the instructions on this [tutorial](https://docs.victoriametrics.com/helm/requirements/).

Generate docs with ``helm-docs`` command.

```bash
cd charts/victoria-traces-single

helm-docs
```

The markdown generation is entirely go template driven. The tool parses metadata from charts and generates a number of sub-templates that can be referenced in a template file (by default ``README.md.gotmpl``). If no template file is provided, the tool has a default internal template that will generate a reasonably formatted README.

## Parameters

The following tables lists the configurable parameters of the chart and their default values.

Change the values according to the need of the environment in ``victoria-traces-single/values.yaml`` file.

<table class="helm-vars">
  <thead>
    <th class="helm-vars-key">Key</th>
    <th class="helm-vars-description">Description</th>
  </thead>
  <tbody>
    <tr id="extraobjects">
      <td><a href="#extraobjects"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">extraObjects</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Add extra specs dynamically to this chart</p>
</td>
    </tr>
    <tr id="global-cluster-dnsdomain">
      <td><a href="#global-cluster-dnsdomain"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">global.cluster.dnsDomain</span><span class="p">:</span><span class="w"> </span><span class="l">cluster.local.</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>K8s cluster domain suffix, uses for building storage pods&rsquo; FQDN. Details are <a href="https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/" target="_blank">here</a></p>
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
      <td><a href="#nameoverride"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">nameOverride</span><span class="p">:</span><span class="w"> </span><span class="l">vt-single</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Override chart name</p>
</td>
    </tr>
    <tr id="poddisruptionbudget">
      <td><a href="#poddisruptionbudget"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">podDisruptionBudget</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">false</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">extraLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>See <code>kubectl explain poddisruptionbudget.spec</code> for more. Details are <a href="https://kubernetes.io/docs/tasks/run-application/configure-pdb/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="poddisruptionbudget-extralabels">
      <td><a href="#poddisruptionbudget-extralabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">podDisruptionBudget.extraLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>PodDisruptionBudget extra labels</p>
</td>
    </tr>
    <tr id="printnotes">
      <td><a href="#printnotes"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">printNotes</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Print chart notes</p>
</td>
    </tr>
    <tr id="server-affinity">
      <td><a href="#server-affinity"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.affinity</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Pod affinity</p>
</td>
    </tr>
    <tr id="server-containerworkingdir">
      <td><a href="#server-containerworkingdir"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.containerWorkingDir</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Container workdir</p>
</td>
    </tr>
    <tr id="server-deployment">
      <td><a href="#server-deployment"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.deployment</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">spec</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">        </span><span class="nt">strategy</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">            </span><span class="nt">type</span><span class="p">:</span><span class="w"> </span><span class="l">Recreate</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p><a href="https://kubernetes.io/docs/concepts/workloads/controllers/deployment/" target="_blank">K8s Deployment</a> specific variables</p>
</td>
    </tr>
    <tr id="server-emptydir">
      <td><a href="#server-emptydir"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.emptyDir</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em></td>
    </tr>
    <tr id="server-enabled">
      <td><a href="#server-enabled"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Enable deployment of server component. Deployed as StatefulSet</p>
</td>
    </tr>
    <tr id="server-env">
      <td><a href="#server-env"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.env</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Additional environment variables (ex.: secret tokens, flags). Details are <a href="https://docs.victoriametrics.com/victoriametrics/#environment-variables" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-envfrom">
      <td><a href="#server-envfrom"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.envFrom</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Specify alternative source for env variables</p>
</td>
    </tr>
    <tr id="server-extraargs">
      <td><a href="#server-extraargs"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.extraArgs</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">envflag.enable</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">envflag.prefix</span><span class="p">:</span><span class="w"> </span><span class="l">VM_</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">http.shutdownDelay</span><span class="p">:</span><span class="w"> </span><span class="l">15s</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">httpListenAddr</span><span class="p">:</span><span class="w"> </span><span class="p">:</span><span class="m">10428</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">loggerFormat</span><span class="p">:</span><span class="w"> </span><span class="l">json</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Extra command line arguments for container of component</p>
</td>
    </tr>
    <tr id="server-extracontainers">
      <td><a href="#server-extracontainers"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.extraContainers</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Extra containers to run in a pod with VictoriaTraces container</p>
</td>
    </tr>
    <tr id="server-extrahostpathmounts">
      <td><a href="#server-extrahostpathmounts"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.extraHostPathMounts</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Additional hostPath mounts</p>
</td>
    </tr>
    <tr id="server-extralabels">
      <td><a href="#server-extralabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.extraLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>StatefulSet/Deployment additional labels</p>
</td>
    </tr>
    <tr id="server-extravolumemounts">
      <td><a href="#server-extravolumemounts"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.extraVolumeMounts</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Extra Volume Mounts for the container</p>
</td>
    </tr>
    <tr id="server-extravolumes">
      <td><a href="#server-extravolumes"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.extraVolumes</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Extra Volumes for the pod</p>
</td>
    </tr>
    <tr id="server-fullnameoverride">
      <td><a href="#server-fullnameoverride"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.fullnameOverride</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Overrides the full name of server component</p>
</td>
    </tr>
    <tr id="server-image-pullpolicy">
      <td><a href="#server-image-pullpolicy"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.image.pullPolicy</span><span class="p">:</span><span class="w"> </span><span class="l">IfNotPresent</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Image pull policy</p>
</td>
    </tr>
    <tr id="server-image-registry">
      <td><a href="#server-image-registry"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.image.registry</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Image registry</p>
</td>
    </tr>
    <tr id="server-image-repository">
      <td><a href="#server-image-repository"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.image.repository</span><span class="p">:</span><span class="w"> </span><span class="l">victoriametrics/victoria-traces</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Image repository</p>
</td>
    </tr>
    <tr id="server-image-tag">
      <td><a href="#server-image-tag"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.image.tag</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Image tag</p>
</td>
    </tr>
    <tr id="server-image-variant">
      <td><a href="#server-image-variant"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.image.variant</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Image tag suffix, which is appended to <code>Chart.AppVersion</code> if no <code>server.image.tag</code> is defined</p>
</td>
    </tr>
    <tr id="server-imagepullsecrets">
      <td><a href="#server-imagepullsecrets"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.imagePullSecrets</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Image pull secrets</p>
</td>
    </tr>
    <tr id="server-ingress-annotations">
      <td><a href="#server-ingress-annotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.ingress.annotations</span><span class="p">:</span><span class="w"> </span><span class="kc">null</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Ingress annotations</p>
</td>
    </tr>
    <tr id="server-ingress-enabled">
      <td><a href="#server-ingress-enabled"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.ingress.enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">false</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Enable deployment of ingress for server component</p>
</td>
    </tr>
    <tr id="server-ingress-extralabels">
      <td><a href="#server-ingress-extralabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.ingress.extraLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Ingress extra labels</p>
</td>
    </tr>
    <tr id="server-ingress-hosts">
      <td><a href="#server-ingress-hosts"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.ingress.hosts</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span>- <span class="nt">name</span><span class="p">:</span><span class="w"> </span><span class="l">vtraces.local</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">      </span><span class="nt">path</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">        </span>- <span class="l">/</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">      </span><span class="nt">port</span><span class="p">:</span><span class="w"> </span><span class="l">http</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Array of host objects</p>
</td>
    </tr>
    <tr id="server-ingress-ingressclassname">
      <td><a href="#server-ingress-ingressclassname"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.ingress.ingressClassName</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Ingress controller class name</p>
</td>
    </tr>
    <tr id="server-ingress-pathtype">
      <td><a href="#server-ingress-pathtype"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.ingress.pathType</span><span class="p">:</span><span class="w"> </span><span class="l">Prefix</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Ingress path type</p>
</td>
    </tr>
    <tr id="server-ingress-tls">
      <td><a href="#server-ingress-tls"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.ingress.tls</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Array of TLS objects</p>
</td>
    </tr>
    <tr id="server-initcontainers">
      <td><a href="#server-initcontainers"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.initContainers</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Init containers for VictoriaTraces Pod</p>
</td>
    </tr>
    <tr id="server-lifecycle">
      <td><a href="#server-lifecycle"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.lifecycle</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Specify pod lifecycle</p>
</td>
    </tr>
    <tr id="server-mode">
      <td><a href="#server-mode"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.mode</span><span class="p">:</span><span class="w"> </span><span class="l">statefulSet</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>VictoriaTraces mode: deployment, statefulSet</p>
</td>
    </tr>
    <tr id="server-nodeselector">
      <td><a href="#server-nodeselector"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.nodeSelector</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Pod&rsquo;s node selector. Details are <a href="https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-persistentvolume-accessmodes">
      <td><a href="#server-persistentvolume-accessmodes"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.accessModes</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span>- <span class="l">ReadWriteOnce</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Array of access modes. Must match those of existing PV or dynamic provisioner. Details are <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-persistentvolume-annotations">
      <td><a href="#server-persistentvolume-annotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.annotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Persistent volume annotations</p>
</td>
    </tr>
    <tr id="server-persistentvolume-enabled">
      <td><a href="#server-persistentvolume-enabled"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Create/use Persistent Volume Claim for server component. Use empty dir if set to false</p>
</td>
    </tr>
    <tr id="server-persistentvolume-existingclaim">
      <td><a href="#server-persistentvolume-existingclaim"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.existingClaim</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Existing Claim name. If defined, PVC must be created manually before volume will be bound</p>
</td>
    </tr>
    <tr id="server-persistentvolume-matchlabels">
      <td><a href="#server-persistentvolume-matchlabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.matchLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Bind Persistent Volume by labels. Must match all labels of targeted PV.</p>
</td>
    </tr>
    <tr id="server-persistentvolume-mountpath">
      <td><a href="#server-persistentvolume-mountpath"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.mountPath</span><span class="p">:</span><span class="w"> </span><span class="l">/storage</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Mount path. Server data Persistent Volume mount root path.</p>
</td>
    </tr>
    <tr id="server-persistentvolume-name">
      <td><a href="#server-persistentvolume-name"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.name</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Override Persistent Volume Claim name</p>
</td>
    </tr>
    <tr id="server-persistentvolume-size">
      <td><a href="#server-persistentvolume-size"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.size</span><span class="p">:</span><span class="w"> </span><span class="l">10Gi</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Size of the volume. Should be calculated based on the traces you send and retention policy you set.</p>
</td>
    </tr>
    <tr id="server-persistentvolume-storageclassname">
      <td><a href="#server-persistentvolume-storageclassname"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.storageClassName</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>StorageClass to use for persistent volume. Requires server.persistentVolume.enabled: true. If defined, PVC created automatically</p>
</td>
    </tr>
    <tr id="server-persistentvolume-subpath">
      <td><a href="#server-persistentvolume-subpath"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.persistentVolume.subPath</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Mount subpath</p>
</td>
    </tr>
    <tr id="server-podannotations">
      <td><a href="#server-podannotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.podAnnotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Pod&rsquo;s annotations</p>
</td>
    </tr>
    <tr id="server-podlabels">
      <td><a href="#server-podlabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.podLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Pod&rsquo;s additional labels</p>
</td>
    </tr>
    <tr id="server-podsecuritycontext">
      <td><a href="#server-podsecuritycontext"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.podSecurityContext</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">fsGroup</span><span class="p">:</span><span class="w"> </span><span class="m">2000</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">runAsNonRoot</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">runAsUser</span><span class="p">:</span><span class="w"> </span><span class="m">1000</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Pod&rsquo;s security context. Details are <a href="https://kubernetes.io/docs/tasks/configure-pod-container/security-context/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-priorityclassname">
      <td><a href="#server-priorityclassname"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.priorityClassName</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Name of Priority Class</p>
</td>
    </tr>
    <tr id="server-probe-liveness">
      <td><a href="#server-probe-liveness"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.probe.liveness</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">failureThreshold</span><span class="p">:</span><span class="w"> </span><span class="m">10</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">initialDelaySeconds</span><span class="p">:</span><span class="w"> </span><span class="m">30</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">periodSeconds</span><span class="p">:</span><span class="w"> </span><span class="m">30</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">tcpSocket</span><span class="p">:</span><span class="w"> </span>{}<span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">timeoutSeconds</span><span class="p">:</span><span class="w"> </span><span class="m">5</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Indicates whether the Container is running. If the liveness probe fails, the kubelet kills the Container, and the Container is subjected to its restart policy. If a Container does not provide a liveness probe, the default state is Success.</p>
</td>
    </tr>
    <tr id="server-probe-readiness">
      <td><a href="#server-probe-readiness"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.probe.readiness</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">failureThreshold</span><span class="p">:</span><span class="w"> </span><span class="m">3</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">httpGet</span><span class="p">:</span><span class="w"> </span>{}<span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">initialDelaySeconds</span><span class="p">:</span><span class="w"> </span><span class="m">5</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">periodSeconds</span><span class="p">:</span><span class="w"> </span><span class="m">5</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">timeoutSeconds</span><span class="p">:</span><span class="w"> </span><span class="m">5</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Indicates whether the Container is ready to service requests. If the readiness probe fails, the endpoints controller removes the Pod&rsquo;s IP address from the endpoints of all Services that match the Pod. The default state of readiness before the initial delay is Failure. If a Container does not provide a readiness probe, the default state is Success.</p>
</td>
    </tr>
    <tr id="server-probe-startup">
      <td><a href="#server-probe-startup"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.probe.startup</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Indicates whether the Container is done with potentially costly initialization. If set it is executed first. If it fails Container is restarted. If it succeeds liveness and readiness probes takes over.</p>
</td>
    </tr>
    <tr id="server-replicacount">
      <td><a href="#server-replicacount"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.replicaCount</span><span class="p">:</span><span class="w"> </span><span class="m">1</span></span></span></code></pre>
</a></td>
      <td><em><code>(int)</code></em><p>Replica count</p>
</td>
    </tr>
    <tr id="server-resources">
      <td><a href="#server-resources"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.resources</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Resource object. Details are <a href="https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-retentiondiskspaceusage">
      <td><a href="#server-retentiondiskspaceusage"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.retentionDiskSpaceUsage</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Data retention max capacity. Default unit is GiB. See these <a href="https://docs.victoriametrics.com/victoriatraces/#retention-by-disk-space-usage" target="_blank">docs</a></p>
</td>
    </tr>
    <tr id="server-retentionperiod">
      <td><a href="#server-retentionperiod"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.retentionPeriod</span><span class="p">:</span><span class="w"> </span><span class="m">1</span></span></span></code></pre>
</a></td>
      <td><em><code>(int)</code></em><p>Data retention period. Possible units character: h(ours), d(ays), w(eeks), y(ears), if no unit character specified - month. The minimum retention period is 24h. See these <a href="https://docs.victoriametrics.com/victoriatraces/#retention" target="_blank">docs</a></p>
</td>
    </tr>
    <tr id="server-schedulername">
      <td><a href="#server-schedulername"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.schedulerName</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Use an alternate scheduler, e.g. &ldquo;stork&rdquo;. Check details <a href="https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-securitycontext">
      <td><a href="#server-securitycontext"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.securityContext</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">allowPrivilegeEscalation</span><span class="p">:</span><span class="w"> </span><span class="kc">false</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">capabilities</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">        </span><span class="nt">drop</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">            </span>- <span class="l">ALL</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">readOnlyRootFilesystem</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span></span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Security context to be added to server pods</p>
</td>
    </tr>
    <tr id="server-service-annotations">
      <td><a href="#server-service-annotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.annotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Service annotations</p>
</td>
    </tr>
    <tr id="server-service-clusterip">
      <td><a href="#server-service-clusterip"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.clusterIP</span><span class="p">:</span><span class="w"> </span><span class="l">None</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Service ClusterIP</p>
</td>
    </tr>
    <tr id="server-service-externalips">
      <td><a href="#server-service-externalips"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.externalIPs</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Service external IPs. Details are <a href="https://kubernetes.io/docs/concepts/services-networking/service/#external-ips" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-service-externaltrafficpolicy">
      <td><a href="#server-service-externaltrafficpolicy"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.externalTrafficPolicy</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Service external traffic policy. Check <a href="https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip" target="_blank">here</a> for details</p>
</td>
    </tr>
    <tr id="server-service-extraports">
      <td><a href="#server-service-extraports"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.extraPorts</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Extra service ports</p>
</td>
    </tr>
    <tr id="server-service-healthchecknodeport">
      <td><a href="#server-service-healthchecknodeport"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.healthCheckNodePort</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Health check node port for a service. Check <a href="https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip" target="_blank">here</a> for details</p>
</td>
    </tr>
    <tr id="server-service-ipfamilies">
      <td><a href="#server-service-ipfamilies"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.ipFamilies</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>List of service IP families. Check <a href="https://kubernetes.io/docs/concepts/services-networking/dual-stack/#services" target="_blank">here</a> for details.</p>
</td>
    </tr>
    <tr id="server-service-ipfamilypolicy">
      <td><a href="#server-service-ipfamilypolicy"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.ipFamilyPolicy</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Service IP family policy. Check <a href="https://kubernetes.io/docs/concepts/services-networking/dual-stack/#services" target="_blank">here</a> for details.</p>
</td>
    </tr>
    <tr id="server-service-labels">
      <td><a href="#server-service-labels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.labels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Service labels</p>
</td>
    </tr>
    <tr id="server-service-loadbalancerip">
      <td><a href="#server-service-loadbalancerip"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.loadBalancerIP</span><span class="p">:</span><span class="w"> </span><span class="s2">&#34;&#34;</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Service load balancer IP</p>
</td>
    </tr>
    <tr id="server-service-loadbalancersourceranges">
      <td><a href="#server-service-loadbalancersourceranges"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.loadBalancerSourceRanges</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Load balancer source range</p>
</td>
    </tr>
    <tr id="server-service-serviceport">
      <td><a href="#server-service-serviceport"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.servicePort</span><span class="p">:</span><span class="w"> </span><span class="m">10428</span></span></span></code></pre>
</a></td>
      <td><em><code>(int)</code></em><p>Service port</p>
</td>
    </tr>
    <tr id="server-service-targetport">
      <td><a href="#server-service-targetport"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.targetPort</span><span class="p">:</span><span class="w"> </span><span class="l">http</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Target port</p>
</td>
    </tr>
    <tr id="server-service-type">
      <td><a href="#server-service-type"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.service.type</span><span class="p">:</span><span class="w"> </span><span class="l">ClusterIP</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Service type</p>
</td>
    </tr>
    <tr id="server-servicemonitor-annotations">
      <td><a href="#server-servicemonitor-annotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.serviceMonitor.annotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Service Monitor annotations</p>
</td>
    </tr>
    <tr id="server-servicemonitor-basicauth">
      <td><a href="#server-servicemonitor-basicauth"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.serviceMonitor.basicAuth</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Basic auth params for Service Monitor</p>
</td>
    </tr>
    <tr id="server-servicemonitor-enabled">
      <td><a href="#server-servicemonitor-enabled"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.serviceMonitor.enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">false</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Enable deployment of Service Monitor for server component. This is Prometheus operator object</p>
</td>
    </tr>
    <tr id="server-servicemonitor-extralabels">
      <td><a href="#server-servicemonitor-extralabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.serviceMonitor.extraLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>Service Monitor labels</p>
</td>
    </tr>
    <tr id="server-servicemonitor-metricrelabelings">
      <td><a href="#server-servicemonitor-metricrelabelings"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.serviceMonitor.metricRelabelings</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Service Monitor metricRelabelings</p>
</td>
    </tr>
    <tr id="server-servicemonitor-relabelings">
      <td><a href="#server-servicemonitor-relabelings"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.serviceMonitor.relabelings</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Service Monitor relabelings</p>
</td>
    </tr>
    <tr id="server-servicemonitor-targetport">
      <td><a href="#server-servicemonitor-targetport"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.serviceMonitor.targetPort</span><span class="p">:</span><span class="w"> </span><span class="l">http</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Service Monitor target port</p>
</td>
    </tr>
    <tr id="server-statefulset">
      <td><a href="#server-statefulset"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.statefulSet</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">    </span><span class="nt">spec</span><span class="p">:</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">        </span><span class="nt">podManagementPolicy</span><span class="p">:</span><span class="w"> </span><span class="l">OrderedReady</span><span class="w">
</span></span></span><span class="line"><span class="cl"><span class="w">        </span><span class="nt">updateStrategy</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p><a href="https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/" target="_blank">K8s StatefulSet</a> specific variables</p>
</td>
    </tr>
    <tr id="server-statefulset-spec-podmanagementpolicy">
      <td><a href="#server-statefulset-spec-podmanagementpolicy"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.statefulSet.spec.podManagementPolicy</span><span class="p">:</span><span class="w"> </span><span class="l">OrderedReady</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>Deploy order policy for StatefulSet pods</p>
</td>
    </tr>
    <tr id="server-statefulset-spec-updatestrategy">
      <td><a href="#server-statefulset-spec-updatestrategy"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.statefulSet.spec.updateStrategy</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>StatefulSet update strategy. Check <a href="https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#update-strategies" target="_blank">here</a> for details.</p>
</td>
    </tr>
    <tr id="server-terminationgraceperiodseconds">
      <td><a href="#server-terminationgraceperiodseconds"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.terminationGracePeriodSeconds</span><span class="p">:</span><span class="w"> </span><span class="m">60</span></span></span></code></pre>
</a></td>
      <td><em><code>(int)</code></em><p>Pod&rsquo;s termination grace period in seconds</p>
</td>
    </tr>
    <tr id="server-tolerations">
      <td><a href="#server-tolerations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.tolerations</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Node tolerations for server scheduling to nodes with taints. Details are <a href="https://kubernetes.io/docs/concepts/configuration/assign-pod-node/" target="_blank">here</a></p>
</td>
    </tr>
    <tr id="server-topologyspreadconstraints">
      <td><a href="#server-topologyspreadconstraints"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.topologySpreadConstraints</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Pod topologySpreadConstraints</p>
</td>
    </tr>
    <tr id="server-vmservicescrape-annotations">
      <td><a href="#server-vmservicescrape-annotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.vmServiceScrape.annotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em></td>
    </tr>
    <tr id="server-vmservicescrape-enabled">
      <td><a href="#server-vmservicescrape-enabled"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.vmServiceScrape.enabled</span><span class="p">:</span><span class="w"> </span><span class="kc">false</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Enable deployment of VMServiceScrape for server component. This is Victoria Metrics operator object</p>
</td>
    </tr>
    <tr id="server-vmservicescrape-extralabels">
      <td><a href="#server-vmservicescrape-extralabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.vmServiceScrape.extraLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em></td>
    </tr>
    <tr id="server-vmservicescrape-metricrelabelings">
      <td><a href="#server-vmservicescrape-metricrelabelings"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.vmServiceScrape.metricRelabelings</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em></td>
    </tr>
    <tr id="server-vmservicescrape-relabelings">
      <td><a href="#server-vmservicescrape-relabelings"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.vmServiceScrape.relabelings</span><span class="p">:</span><span class="w"> </span><span class="p">[]</span></span></span></code></pre>
</a></td>
      <td><em><code>(list)</code></em><p>Commented. TLS configuration to use when scraping the endpoint    tlsConfig:      insecureSkipVerify: true</p>
</td>
    </tr>
    <tr id="server-vmservicescrape-targetport">
      <td><a href="#server-vmservicescrape-targetport"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">server.vmServiceScrape.targetPort</span><span class="p">:</span><span class="w"> </span><span class="l">http</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>target port</p>
</td>
    </tr>
    <tr id="serviceaccount-annotations">
      <td><a href="#serviceaccount-annotations"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">serviceAccount.annotations</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>ServiceAccount annotations</p>
</td>
    </tr>
    <tr id="serviceaccount-automounttoken">
      <td><a href="#serviceaccount-automounttoken"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">serviceAccount.automountToken</span><span class="p">:</span><span class="w"> </span><span class="kc">true</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Mount API token to pod directly</p>
</td>
    </tr>
    <tr id="serviceaccount-create">
      <td><a href="#serviceaccount-create"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">serviceAccount.create</span><span class="p">:</span><span class="w"> </span><span class="kc">false</span></span></span></code></pre>
</a></td>
      <td><em><code>(bool)</code></em><p>Create service account.</p>
</td>
    </tr>
    <tr id="serviceaccount-extralabels">
      <td><a href="#serviceaccount-extralabels"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">serviceAccount.extraLabels</span><span class="p">:</span><span class="w"> </span>{}</span></span></code></pre>
</a></td>
      <td><em><code>(object)</code></em><p>ServiceAccount labels</p>
</td>
    </tr>
    <tr id="serviceaccount-name">
      <td><a href="#serviceaccount-name"><pre class="chroma"><code><span class="line"><span class="cl"><span class="nt">serviceAccount.name</span><span class="p">:</span><span class="w"> </span><span class="kc">null</span></span></span></code></pre>
</a></td>
      <td><em><code>(string)</code></em><p>The name of the service account to use. If not set and create is true, a name is generated using the fullname template</p>
</td>
    </tr>
  </tbody>
</table>

