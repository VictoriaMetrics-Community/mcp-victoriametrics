---
draft: false
page: blog blog_post
authors:
 - Zakhar Bessarab
date: 2024-10-22
enableComments: true
title: "Monitoring Azure AKS & Azure Linux with VictoriaMetrics"
summary: "Learn how to monitor Azure AKS and Azure Linux with VictoriaMetrics. This blog post covers the setup process for environments with high security requirements and how to monitor them with VictoriaMetrics."
description: ""
categories:
 - Observability
 - Monitoring
 - Kubernetes
tags:
 - open source
 - kubernetes
 - monitoring
 - azure linux
 - azure AKS
keywords: 
 - open source
 - kubernetes
 - monitoring
 - azure linux
 - azure AKS
images:
 - /blog/monitoring-azure-aks-azure-linux-with-victoriametrics/preview.webp
---

# What is Azure Linux?

[Azure linux](https://github.com/microsoft/azurelinux) is a Linux distribution built for Microsoft's cloud infrastructure. 
It can be used as a base OS when creating node pools in Azure Kubernetes Service (AKS) clusters. Using Azure linux as a base OS for AKS node pools
has several benefits, such as lower resources footprint, faster boot times, and better security.

# Using VictoriaMetrics to monitor services running in AKS with Azure Linux

VictoriaMetrics is a high-performance, cost-effective, and scalable open source monitoring solution that can be used to monitor 
services running in AKS with Azure Linux. It can be used in order to monitor the applications running in AKS with Azure Linux, 
as well as the underlying infrastructure.

# How to deploy VictoriaMetrics in AKS

Pre-requisites:

- An Azure account
- An AKS cluster with Azure Linux node pools
- kubectl installed on your local machine and configured to connect to the AKS cluster
- [helm](https://helm.sh/) installed on your local machine

In order to deploy VictoriaMetrics by using a Helm chart in AKS with Azure Linux, you can follow these steps:

1. Prepare configuration values for the Helm chart. You can use the following `values.yaml` file as a starting point:
     
    ```yaml
    victoria-metrics-operator:
      podSecurityContext:
        seccompProfile:
          type: RuntimeDefault
      securityContext:
        runAsUser: 1001
        runAsNonRoot: true
        readOnlyRootFilesystem: true
        allowPrivilegeEscalation: false
        capabilities:
          drop:
            - ALL
      env:
        - name: VM_ENABLESTRICTSECURITY
          value: "true"
    ```

    This configuration is an example of how to enable strict security settings for the VictoriaMetrics operator. It sets strict security settings for the operator&rsquo;s pod, such as running as a non-root user, using a read-only root filesystem, and dropping all capabilities. This configuration should be suitable for most use cases, but you can adjust it according to your needs.

    You can adjust other parameters in the `values.yaml` file according to your requirements. See the [Helm chart documentation](https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-k8s-stack#configuration).

1. Add the VictoriaMetrics Helm repository to your Helm client:
    ```shell
        helm repo add vm https://victoriametrics.github.io/helm-charts/
        helm repo update
    ```

1. Install the VictoriaMetrics Kubernetes stack by using the Helm chart:

    ```shell
        helm install vm-k8s-stack vm/victoria-metrics-k8s-stack -f values.yaml
    ```

    This command installs the VictoriaMetrics operator and the VictoriaMetrics single-node in your AKS cluster with Azure Linux nodes. 
    It also deploys resources for basic monitoring of the cluster and the applications running in it, such as node exporter, kube-state-metrics,
    Grafana and Alertmanager.

## Accessing Grafana dashboards

Once this will be completed you can access Grafana dashboard by using port-forwarding:
```shell
kubectl port-forward svc/vm-k8s-stack-grafana 3000:80
```

Alternatively, it is possible to use an Ingress or LoadBalancer service to expose Grafana UI to the public internet.
Note that setting up Microsoft Entra ID authentication for Grafana requires an endpoint with HTTPS enabled.

Default password can be obtained by using the following command:
```shell
kubectl get secret vm-k8s-stack-grafana -o jsonpath="{.data.admin-password}" | base64 --decode
```
Default administrator account is `admin` and password is the one you've obtained in the previous step.

By default, victoria-metrics-k8s-stack Helm chart deploys a set of dashboards for monitoring Kubernetes cluster and VictoriaMetrics itself.
Once the deployment is completed you can navigate to Grafana UI and start exploring the dashboards.

At this point you can use VictoriaMetrics to monitor the AKS Cluster and the applications running on it.
To Collect Metrics for other applications from other clusters, please refer to our documentation.

- example configuration for [VMServiceScrape](https://docs.victoriametrics.com/operator/quick-start/#vmservicescrape)
- operator CRDs [reference](https://docs.victoriametrics.com/operator/api/)
The VictoriaMetrics operator will automatically configure vmagent instances to scrape metrics from your applications and services based on the custom resources you define.

# Hardening the security of your monitoring setup with VictoriaMetrics Enterprise

VictoriaMetrics Enterprise provides additional features for securing your monitoring setup, such as OIDC authentication and access control.
The next section will cover how to set up OIDC for authentication with VictoriaMetrics using [vmgateway](https://docs.victoriametrics.com/vmgateway/).

You can request a free trial access to VictoriaMetrics Enterprise by using this [form](https://victoriametrics.com/products/enterprise/trial/).

# Setting up OIDC for Azure Entra with VictoriaMetrics

In order to improve security of your monitoring setup, you can use OIDC for authentication with VictoriaMetrics.
[VictoriaMetrics Enterprise](https://victoriametrics.com/products/enterprise/) provides a component which can be used as a reverse proxy for authentication purposes - [vmgateway](https://docs.victoriametrics.com/vmgateway/).
It allows to authenticate users before they access VictoriaMetrics and enforce access control policies.

Microsoft Entra ID is a cloud-based identity and access management service that can be used to authenticate users.
You can use Microsoft Entra ID as an authentication provider for VictoriaMetrics by following these steps:

1.  Create an Application in Entra admin center. See this guide for step-by-step instructions: https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app

2.  Configure Grafana to use Microsoft Entra ID as an authentication provider. See this guide for step-by-step instructions: https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/azuread/

    Use the following configuration values in Grafana for reference:
```shell
grafana:
  env:
    GF_AUTH_AZUREAD_CLIENT_ID: <tenant-id>
    GF_AUTH_AZUREAD_CLIENT_SECRET: <client-secret>
  grafana.ini:
    server:
      domain: "<grafana-domain>"
      root_url: "https://<grafana-domain>" # Note that HTTPS is required for OIDC
    auth.azuread:
      enabled: true
      allow_sign_up: true
      scopes: "openid email profile"
      auth_url: https://login.microsoftonline.com/<tenant_id>/oauth2/v2.0/authorize
      token_url: https://login.microsoftonline.com/<tenant_id>/oauth2/v2.0/token
      allowed_organizations: <tenant_id>
```

## Set up authentication for VictoriaMetrics access by using vmgateway

1. Create a secret with your VictoriaMetrics Enterprise license key:
    ```shell
    kubectl create secret generic vm-license --from-literal=license=<license-key>
    ```

1. Deploy vmgateway by using a [Helm chart](https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-gateway). 
   Save the following as a `values-vmgateway.yaml` file:
    ```yaml
    license:
      secret:
        name: vm-license
        key: license
    
    image:
      tag: v1.104.0-enterprise
    
    auth:
      enabled: true
    
    clusterMode: "<cluster-mode>"
    
    read:
      url: "<victoriametrics-read-url>"
    
    write:
      url: "<victoriametrics-write-url>"
    
    extraArgs:
      envflag.enable: "true"
      envflag.prefix: VM_
      loggerFormat: json
      auth.oidcDiscoveryEndpoints: "https://login.microsoftonline.com/<tenant_id>/v2.0/.well-known/openid-configuration"
      auth.httpHeader: "X-Id-Token"
      auth.httpHeaderAllowWithoutPrefix: "false"
    ```
   
    Where `<victoriametrics-read-url>` and `<victoriametrics-write-url>` are the URLs of your VictoriaMetrics instances for read and write operations.
    For single-node type of deployment the URL will be the same for both options, it should be in the following format: `http://vmsingle-vm-victoria-metrics-k8s-stack.vm.svc:8428`.

    Cluster type of deployment will have different URLs for read and write operations, see the [following docs](https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-gateway) for the details.
    
    `<cluster-mode>` needs to be set to `false` for single-node deployment and `true` for cluster deployment.

    Perform the installation by using the following command:
    ```shell
    helm install vm-gateway vm/victoria-metrics-gateway -f values-gateway.yaml
    ```

1. Add Grafana datasource configuration to query VictoriaMetrics via vmgateway.
    Update Grafana deployment configuration to add the following:
    ```yaml
    grafana:
      datasources:
        datasources.yaml:
          apiVersion: 1
          datasources:
          - name: VictoriaMetrics-vmgateway
            type: prometheus
            url: http://vm-gateway-victoria-metrics-gateway:8431
            access: proxy
            isDefault: false
            jsonData:
              oauthPassThru: true
    ```
    Using `oauthPassThru` instructs Grafana to send authentication token from Microsoft Entra ID to the datasource endpoint. 
    vmgateway will use these tokens to verify if user is allowed to access VictoriaMetrics.

1. Set up attribute mapping for `vm_access` field.
    In order to enforce restricted access to data stored in VictoriaMetrics it is possible to provide additional filtering configuration via access token. See [this docs](https://docs.victoriametrics.com/vmgateway/#access-control) for the details on `vm_access` field format.
    See these docs in order to configure attribute mapping for Microsoft Entra ID:
    - https://learn.microsoft.com/en-us/entra/identity-platform/reference-claims-customization
    - https://learn.microsoft.com/en-us/entra/fundamentals/custom-security-attributes-overview
    - https://learn.microsoft.com/en-us/entra/identity-platform/custom-claims-provider-reference

Note that when changing the attribute configuration mapping in Microsoft Entra ID it is required to log out and log in again in to get a token with the new attributes.
After that you can navigate to Grafana log-in page, authenticate by using a newly created Microsoft Entra ID option and use `VictoriaMetrics-vmgateway` datasource for querying.

# Conclusion

In this blog post, we have covered how to monitor Azure AKS and Azure Linux with VictoriaMetrics. 
We have shown how to deploy VictoriaMetrics in AKS with Azure Linux and how to set up OIDC for authentication with VictoriaMetrics using vmgateway. 
By following these steps, you can monitor your services running in AKS with Azure Linux in a secure and efficient way.
