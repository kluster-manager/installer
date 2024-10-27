# managed-serviceaccount-manager

[managed-serviceaccount-manager](https://github.com/kluster-manager/managed-serviceaccount) - Installs managed-serviceaccount in managed clusters

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/managed-serviceaccount-manager --version=v2024.7.10
$ helm upgrade -i managed-serviceaccount-manager appscode/managed-serviceaccount-manager -n open-cluster-management-addon --create-namespace --version=v2024.7.10
```

## Introduction

This chart deploys a managed-serviceaccount addon manager on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `managed-serviceaccount-manager`:

```bash
$ helm upgrade -i managed-serviceaccount-manager appscode/managed-serviceaccount-manager -n open-cluster-management-addon --create-namespace --version=v2024.7.10
```

The command deploys a managed-serviceaccount addon manager on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `managed-serviceaccount-manager`:

```bash
$ helm uninstall managed-serviceaccount-manager -n open-cluster-management-addon
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `managed-serviceaccount-manager` chart and their default values.

|           Parameter            |                  Description                   |                                                                                            Default                                                                                             |
|--------------------------------|------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                   |                                                | <code>""</code>                                                                                                                                                                                |
| fullnameOverride               |                                                | <code>""</code>                                                                                                                                                                                |
| image                          | Image of the managed service-account instances | <code>ghcr.io/kluster-manager/managed-serviceaccount</code>                                                                                                                                    |
| tag                            |                                                | <code>""</code>                                                                                                                                                                                |
| replicas                       | Number of replicas                             | <code>1</code>                                                                                                                                                                                 |
| securityContext                |                                                | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| featureGates.ephemeralIdentity |                                                | <code>false</code>                                                                                                                                                                             |
| agentImagePullSecret           |                                                | <code>""</code>                                                                                                                                                                                |
| kubeconfigSecretName           | required for multicluster controlplane         | <code>""</code>                                                                                                                                                                                |
| addonManagerNamespace          |                                                | <code>open-cluster-management-managed-serviceaccount</code>                                                                                                                                    |
| placement.create               |                                                | <code>true</code>                                                                                                                                                                              |
| placement.name                 |                                                | <code>global</code>                                                                                                                                                                            |
| kubectl.image                  |                                                | <code>ghcr.io/appscode/kubectl-nonroot:1.31</code>                                                                                                                                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i managed-serviceaccount-manager appscode/managed-serviceaccount-manager -n open-cluster-management-addon --create-namespace --version=v2024.7.10 --set image=ghcr.io/kluster-manager/managed-serviceaccount
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i managed-serviceaccount-manager appscode/managed-serviceaccount-manager -n open-cluster-management-addon --create-namespace --version=v2024.7.10 --values values.yaml
```
