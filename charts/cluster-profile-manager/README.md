# Cluster Profile Manager

[Cluster Profile Manager](https://github.com/kluster-manager/cluster-profile) - Cluster Profile Manager

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/cluster-profile-manager --version=v2026.2.16
$ helm upgrade -i cluster-profile-manager appscode/cluster-profile-manager -n open-cluster-management-addon --create-namespace --version=v2026.2.16
```

## Introduction

This chart deploys an Cluster Profile Manager on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-profile-manager`:

```bash
$ helm upgrade -i cluster-profile-manager appscode/cluster-profile-manager -n open-cluster-management-addon --create-namespace --version=v2026.2.16
```

The command deploys an Cluster Profile Manager on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-profile-manager`:

```bash
$ helm uninstall cluster-profile-manager -n open-cluster-management-addon
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-profile-manager` chart and their default values.

|                 Parameter                 |                                    Description                                     |                                                                                             Default                                                                                             |
|-------------------------------------------|------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                              |                                                                                    | <code>""</code>                                                                                                                                                                                 |
| fullnameOverride                          |                                                                                    | <code>""</code>                                                                                                                                                                                 |
| manager.image                             |                                                                                    | <code>ghcr.io/kluster-manager/cluster-profile</code>                                                                                                                                            |
| manager.tag                               |                                                                                    | <code>""</code>                                                                                                                                                                                 |
| manager.imagePullPolicy                   |                                                                                    | <code>Always</code>                                                                                                                                                                             |
| manager.securityContext                   |                                                                                    | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"privileged":false,"readOnlyRootFilesystem":true,"runAsNonRoot":true,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| manager.envFrom                           | List of sources to populate environment variables in the container                 | <code>[]</code>                                                                                                                                                                                 |
| manager.env                               | List of environment variables to set in the container                              | <code>[]</code>                                                                                                                                                                                 |
| placement.create                          |                                                                                    | <code>true</code>                                                                                                                                                                               |
| placement.name                            |                                                                                    | <code>global</code>                                                                                                                                                                             |
| registryFQDN                              | ace values Docker registry fqdn used to pull cluster-profile-manager docker images | <code>ghcr.io</code>                                                                                                                                                                            |
| offlineInstaller                          |                                                                                    | <code>false</code>                                                                                                                                                                              |
| image.proxies.appscode                    | r.appscode.com                                                                     | <code>r.appscode.com</code>                                                                                                                                                                     |
| image.proxies.dockerHub                   | company/bin:tag                                                                    | <code>""</code>                                                                                                                                                                                 |
| image.proxies.dockerLibrary               | alpine, nginx etc.                                                                 | <code>""</code>                                                                                                                                                                                 |
| image.proxies.ghcr                        | ghcr.io/company/bin:tag                                                            | <code>ghcr.io</code>                                                                                                                                                                            |
| image.proxies.quay                        | quay.io/company/bin:tag                                                            | <code>quay.io</code>                                                                                                                                                                            |
| image.proxies.kubernetes                  | registry.k8s.io/bin:tag                                                            | <code>registry.k8s.io</code>                                                                                                                                                                    |
| image.proxies.microsoft                   |                                                                                    | <code>mcr.microsoft.com</code>                                                                                                                                                                  |
| image.proxies.weaviate                    |                                                                                    | <code>cr.weaviate.io</code>                                                                                                                                                                     |
| registry.credentials                      |                                                                                    | <code>{}</code>                                                                                                                                                                                 |
| registry.certs                            | username: "abc" password: "xyz"                                                    | <code>{}</code>                                                                                                                                                                                 |
| registry.imagePullSecrets                 | ca.crt: "***"                                                                      | <code>[]</code>                                                                                                                                                                                 |
| helm.createNamespace                      |                                                                                    | <code>true</code>                                                                                                                                                                               |
| helm.repositories.appscode-charts-oci.url |                                                                                    | <code>oci://ghcr.io/appscode-charts</code>                                                                                                                                                      |
| helm.releases.kube-ui-server.enabled      |                                                                                    | <code>true</code>                                                                                                                                                                               |
| helm.releases.kube-ui-server.version      |                                                                                    | <code>"v2024.8.21"</code>                                                                                                                                                                       |
| helm.releases.opscenter-features.enabled  |                                                                                    | <code>true</code>                                                                                                                                                                               |
| helm.releases.opscenter-features.version  |                                                                                    | <code>"v2024.8.21"</code>                                                                                                                                                                       |
| platform.baseURL                          |                                                                                    | <code>""</code>                                                                                                                                                                                 |
| platform.token                            |                                                                                    | <code>""</code>                                                                                                                                                                                 |
| platform.caBundle                         |                                                                                    | <code>""</code>                                                                                                                                                                                 |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-profile-manager appscode/cluster-profile-manager -n open-cluster-management-addon --create-namespace --version=v2026.2.16 --set manager.image=ghcr.io/kluster-manager/cluster-profile
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-profile-manager appscode/cluster-profile-manager -n open-cluster-management-addon --create-namespace --version=v2026.2.16 --values values.yaml
```
