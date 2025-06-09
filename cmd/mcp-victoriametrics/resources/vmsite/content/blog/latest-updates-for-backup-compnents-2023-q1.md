---
draft: false
page: blog blog_post
authors:
 - Zakhar Bessarab
date: 2023-01-09
title: "Latest updates about backup components of VictoriaMetrics"
summary: "An overview of the latest features added to VictoriaMetrics backup components."
enableComments: true
categories: 
 - Company News
 - Product News
tags:
 - open source
 - victoriametrics
 - kubernetes
 - new features
 - backups
images:
 - /blog/latest-updates-for-backup-components-2023-q1/preview.webp
---

VictoriaMetrics is proud to announce that we consider [vmbackup](https://docs.victoriametrics.com/vmbackup.html) and [vmbackupmanager](https://docs.victoriametrics.com/vmbackupmanager.html) to be feature-complete solutions as of release [1.85.3](https://docs.victoriametrics.com/CHANGELOG.html#v1853). These backup components are essential for ensuring the safety and integrity of your data, and we have made a number of improvements in recent releases to make them even more reliable and user-friendly.

Some key updates in the last few releases include:
- Support for Azure Blob Storage, added in [1.82.0](https://docs.victoriametrics.com/CHANGELOG.html#v1820)
- Enhanced automation for [Kubernetes](https://docs.victoriametrics.com/vmbackupmanager.html#how-to-restore-in-kubernetes) and [CLI mode](https://docs.victoriametrics.com/vmbackupmanager.html#cli), introduced in [1.83.0](https://docs.victoriametrics.com/CHANGELOG.html#v1830)
- Improved observability and the release of an official [metrics dashboard](https://github.com/VictoriaMetrics/VictoriaMetrics/blob/master/dashboards/backupmanager.json) in [1.85.3](https://docs.victoriametrics.com/CHANGELOG.html#v1853)

We hope these updates will help you make the most of these powerful backup solutions and keep your data safe and secure.

## Support of Azure Blob Storage

Supporting Azure Blob Storage was a final piece to support all 3 major cloud providers. Azure storage backend is available for both vmbackup and vmbackupmanager.

In order to use newly supported storage backend provide credentials and URL for backup location in form `azblob://<container>/<path/to/backup>`.
Credentials for Azure Blob must be provided as environment variables. You can either use `AZURE_STORAGE_ACCOUNT_NAME` and `AZURE_STORAGE_ACCOUNT_KEY`, or `AZURE_STORAGE_ACCOUNT_CONNECTION_STRING`.

vmbackupmanager simplifies and streamlines the process of regularly creating backups. It runs as a separate service alongside the [single-node](https://docs.victoriametrics.com/Single-server-VictoriaMetrics.html) or vmstorage node of a [cluster](https://docs.victoriametrics.com/Cluster-VictoriaMetrics.html) setup, and uses command-line flags to define the types of backups to be created and retention policies to be applied.

vmbackupmanager also employs a smart upload strategy to save bandwidth when creating multiple backups. This is achieved by first uploading changed partitions to the latest backup, and then using server-side copy to create a full backup from already-uploaded partitions.

## Improved Kubernetes automation for vmbackupmanager

Previously, vmbackupmanager was only responsible for creating backups. In order to restore from the backup, it was required to use [vmrestore](https://docs.victoriametrics.com/vmrestore.html) with its own configuration flags (even though they're similar).
Starting from [1.83.0](https://docs.victoriametrics.com/CHANGELOG.html#v1830) it is possible to use vmbackupmanager for full backup lifecycle.

Before this feature was added, in order to restore backup data in k8s it was required to do a quite complex and error-prone process which required:
- To scale down existing storage nodes
- To attach PVCs to pods with vmrestore and perform restore
- Scale storage nodes to previous state

Or an easier option such as:
- To add vmrestore as an init container (this step triggers pods restart and restores the data)
- To remove init container (otherwise after pod restart vmrestore will restore backup again)

In order to improve this process we have added CLI mode for vmbackupmanager which will allow:
- To see all backups available for restore at the current storage node
- To configure which backup should be restored on the next pod restart

Now in order to restore from the backup, assuming vmbackupmanager is already used for creating the backup, follow the procedure below:
- Add an init container for vmstorage or single-node pod.
  - For the [VM Operator](https://docs.victoriametrics.com/operator/VictoriaMetrics-Operator.html) you just need to add the following:
    ```yaml
    vmBackup:
      restore:
        onStart:
          enabled: true
    ```
  - For [official Helm charts](https://github.com/VictoriaMetrics/helm-charts) users:
    ```yaml
    vmbackupmanager:
      restore:
        onStart:
          enabled: true
    ```
- Exec into the pod
- Run `/vmbackupmanager backup list` and pick a backup you need to restore
  ```console
  $ /vmbackupmanager backup list
  ["daily/2022-10-06","daily/2022-10-10","hourly/2022-10-04:13","hourly/2022-10-06:12","hourly/2022-10-06:13","hourly/2022-10-10:14","hourly/2022-10-10:16","monthly/2022-10","weekly/2022-40","weekly/2022-41"]
  ```
- Run `/vmbackupmanager restore create {backup_name}`
  ```console
  $ /vmbackupmanager restore create daily/2022-10-10
  ```
- Restart pod

The whole process now seems much easier and less error-prone.

By using CLI it is also easy to restore backup into a separate cluster, read more about this [in our documentation](https://docs.victoriametrics.com/vmbackupmanager.html#restore-cluster-into-another-cluster).

But that's not all - the CLI commands for vmbackupmanager actually use the vmbackupmanager API to perform operations. This means that you can use the vmbackupmanager APIs to build custom workflows for more specialized use cases.
You can find more info about using API [here](https://docs.victoriametrics.com/vmbackupmanager.html#api-methods).

## Improved observability

The last step was to add better observability for the backups process. Release [1.85.3](https://docs.victoriametrics.com/CHANGELOG.html#v1853) added metrics which made it possible to get more information about vmbackupmanager status.
Having new data in place allowed us to build an official dashboard for vmbackupmanager. You can find it on [Grafana website](https://grafana.com/grafana/dashboards/17798-victoriametrics-backupmanager/) and [Github](https://github.com/VictoriaMetrics/VictoriaMetrics/blob/master/dashboards/backupmanager.json).
{{<image href="/blog/latest-updates-for-backup-components-2023-q1/grafana-vmbackupmanager-dashboard.webp" alt="Grafana dashboard for vmbackupmanager" >}}

Make sure to configure [monitoring](https://docs.victoriametrics.com/vmbackupmanager.html#monitoring) of vmbackupmanager to see the status of backups.
