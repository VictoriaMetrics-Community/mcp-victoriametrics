---
draft: false
page: blog blog_post
authors:
 - Mathias Palmersheim
 - Denys Holius
date: 2024-06-19
title: "Monitoring Proxmox VE via VictoriaMetrics Cloud"
enableComments: true
summary: "Monitoring Proxmox hypervisor via VictoriaMetrics and Proxmox's built-in metric server"
categories: 
 - Monitoring
 - Time Series Database
tags:
 - time series database
 - victoriametrics cloud
 - cloud
 - open source
 - monitoring
 - proxmox
 - PVE
images:
 - /blog/proxmox-monitoring-with-vmcloud/proxmox-grafana-dashboard.webp
aliases:
 - /blog/proxmox-monitoring-with-dbaas/
 - /blog/proxmox-monitoring-with-dbaas/index.html
---

_This Post was updated in June of 2024 to remove the requirement to install VMAgent on each Proxmox VE node, and update the screenshots to reflect updates in VictoriaMetrics Cloud and Grafana._

---------

# Monitoring Proxmox VE via VictoriaMetrics CLoud 

In this blog post we’re going to walk you through how to monitor Proxmox VE via [VictoriaMetrics Cloud](https://victoriametrics.com/products/cloud/), including a step by step guide on how to setup, configure, and visualize this environment.


Proxmox VE is a complete, open-source server management platform for enterprise virtualization. It tightly integrates the [KVM hypervisor](https://www.linux-kvm.org/page/Main_Page), [Linux Containers (LXC)](https://linuxcontainers.org), software-defined storage, and, software defined networking in a single platform. All of these features can be managed via a web GUI, CLI, and infrastructure as code tools such as [Ansible](https://docs.ansible.com/ansible/latest/collections/community/general/proxmox_module.html) and [Terraform](https://registry.terraform.io/providers/bpg/proxmox/latest/docs). All of these interfaces support managing virtual machines, containers, high availability for clusters, and configuring disaster recovery.

VictoriaMetrics Cloud allows users to run VictoriaMetrics on [AWS](https://aws.amazon.com) without the need to perform typical DevOps tasks such as proper configuration, monitoring, logs collection, access protection, software updates, backups, etc.

**The guide covers:**
* How to setup a deployment on [VictoriaMetrics Cloud](https://console.victoriametrics.cloud/signUp)
* How to view data in VictoriaMetrics Cloud
* How to add VictoriaMetrics Cloud as a datasource in Grafana
* How to visualize data from VictoriaMetrics Cloud in Grafana

**Preconditions:**
* Proxmox VE version 7.0+
* Grafana 11.0+
* VictoriaMetrics Cloud account
* Installed [Grafana](https://grafana.com/)
* [Grafana dashboard for Proxmox VE](https://grafana.com/grafana/dashboards/16060)


## 1. Setup VictoriaMetrics Cloud deployment

If you don't have the VictoriaMetrics Cloud account yet, just [sign up here](https://console.victoriametrics.cloud/signUp) – it's free.

To read more about VictoriaMetrics Cloud see the [announcement blog post](https://victoriametrics.com/blog/managed-victoriametrics-announcement/).

Open https://console.victoriametrics.cloud/deployments and click `start sending metrics` if you don't already have an existing deployment.

Configure the deployment with parameters that best suit your case and click `Create`.
For this demo I chose the starter instance, in the region closest to me with 13 months of retention.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/create-instance.webp" alt="New deployment creation">}}

Once the deployment is created and provisioned, you will get an email notification. Click on the created deployment to see configuration details, and the status should say `running`: 

{{<image href="/blog/proxmox-monitoring-with-vmcloud/choose-instance.webp" alt="Deployment is in Running state">}}

Now we need to generate a token that will allow us to write data to our VictoriaMetrics Cloud instance.
Go to the "Access" tab, type `proxmox` in the name box and set the token type to write, and click generate.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/remote-write-params.webp" alt="Generating Write Token">}}

Then we will need to copy the hostname of your VictoriaMetrics Cloud instance, and the token you just generated.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/remote-write-creds.webp" alt="Getting Hostname and Token">}}

_Please, do not share the generated access tokens with untrusted parties._

## 2. Configure Proxmox VE to send metrics to VictoriaMetrics Cloud

Login as a `root@PAM` or `user@pve` with Administrator permissions to Proxmox VE:

{{<image href="/blog/proxmox-monitoring-with-vmcloud/proxmox-login.webp" alt="Login to Proxmox Web UI">}}

Click on `Datacenter` in the Proxmox UI, then click on `Metric Server`, click add, and click InfluxDB:

{{<image href="/blog/proxmox-monitoring-with-vmcloud/influxdb-metric-server.webp" alt="Adding new Metric Server on Proxmox PVE">}}

Set parameters to match this screenshot replacing the `Token` with token you generated in step 1 and `Server` with the hostname mentioned in step 1:

{{<image href="/blog/proxmox-monitoring-with-vmcloud/influxdb-metric-server-parameters.webp" alt="Configure parameters for Metric Server">}}

## 3. Confirm Data is Being Sent to VictoriaMetrics Cloud

Go to your VictoriaMetrics Cloud deployment and click explore and run the following query `system_uptime{object="nodes"}`.

You should see 1 time series per ProxmoxVE node in your cluster.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/vmui-test.webp" alt="Victoriametrics Data in VictoriaMetrics Cloud">}}

## 4. View the Data in Grafana 

Before adding our VictoriaMetrics Cloud to Grafana we need a read only access token. 

Go to the Access section of your VictoriaMetrics Cloud deployment, type `grafana` in the name box, set the `Token Type` to read, and click generate.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/dbaas-grafana-access.webp" alt="Generating Grafana Token">}}

Click `Show Examples` on the token we just created and copy the values for Grafana

{{<image href="/blog/proxmox-monitoring-with-vmcloud/dbaas-grafana-creds.webp" alt="Getting credential for VictoriaMetrics Cloud">}}

To add VictoriaMetrics Cloud as a datasource in Grafana, login to Grafana as an admin, click the hamburger menu in the top left of the screen, click `Connections`, then click `Data sources`, and click `Add new data source`.

On the next page select `Prometheus` as the datasource type.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/add-datasource.webp" alt="Adding Connection in Grafana">}}

To configure the datasource change the following options:

1. Set the `Name` to `vm-dbaas`
2. Set the `Prometheus server URL` to the Data Source URL shown in example above.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/grafana-datasource-1.webp" alt="Adding VictoriaMetrics datasource in Grafana">}}

3. Expand the HTTP Headers section, click Add Header, set the value to Authorization and the value to the `Header Value` in the example above.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/grafana-datasource-2.webp" alt="Setting Auth Header on Grafana Datasource">}}

After changing these settings click save and test at the bottom.
If You see a Green check mark saying `Successfully queried the Prometheus API` then everything is working.

To import a [Proxmox dashboard](https://grafana.com/grafana/dashboards/16060) click the hamburger menu in the upper left, click dashboards, click the blue button that says `New` on the right side of the screen, and press `Import` on the drop down:

{{<image href="/blog/proxmox-monitoring-with-vmcloud/import-proxmox-dahsboard1.webp" alt="Adding new dashboard to Grafana">}}

Enter dashboard's **ID 16060** and press the `Load` button on the right:

{{<image href="/blog/proxmox-monitoring-with-vmcloud/import-proxmox-dahsboard2.webp" alt="Importing Proxmox dashboard">}}

Choose `vm-dbaas` as a datasource and press the `Import` button.

{{<image href="/blog/proxmox-monitoring-with-vmcloud/import-proxmox-dahsboard3.webp" alt="Importing Proxmox dashboard">}}

The expected result will resemble the following screenshot:

{{<image href="/blog/proxmox-monitoring-with-vmcloud/proxmox-grafana-dashboard.webp" alt="Full Proxmox dashboard">}}

If this is what you see, then everything is working and you can observe the data on the dashboard!

## 5. Final thoughts

* We have set up a deployment of [VictoriaMetrics Cloud](https://console.victoriametrics.cloud?utm_source=blog&utm_campaign=proxmox)
* We have configured Proxmox VE to send metrics to VictoriaMetrics Cloud over HTTPS
* We have configured VictoriaMetrics Cloud as a datasource in Grafana
* We visualized resource usage in Proxmox VE via VictoriaMetrics Cloud and Grafana

Please comment below if you have any questions, or feel free to [contact us](https://victoriametrics.com/contact-us/). 
