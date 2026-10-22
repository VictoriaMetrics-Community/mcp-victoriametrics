

# crds

![Version: 0.0.0](https://img.shields.io/badge/Version-0.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.0](https://img.shields.io/badge/AppVersion-0.0.0-informational?style=flat-square)

A subchart stores victoriametrics operator CRDs.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| upgrade.busybox.image.pullPolicy | string | `"IfNotPresent"` |  |
| upgrade.busybox.image.repository | string | `"busybox"` |  |
| upgrade.busybox.image.tag | string | `"latest"` |  |
| upgrade.enabled | bool | `false` |  |
| upgrade.forceConflicts | bool | `false` |  |
| upgrade.kubectl.image.pullPolicy | string | `"IfNotPresent"` |  |
| upgrade.kubectl.image.repository | string | `"rancher/kubectl"` |  |
| upgrade.kubectl.image.tag | string | `""` |  |
| upgrade.serviceAccount.create | bool | `true` |  |
| upgrade.serviceAccount.labels | object | `{}` |  |

