# managed-serviceaccount

[managed-serviceaccount](https://github.com/kluster-manager/managed-serviceaccount) - Installs managed-serviceaccount in managed clusters

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/managed-serviceaccount --version=v2024.2.25
$ helm upgrade -i managed-serviceaccount appscode/managed-serviceaccount -n open-cluster-management --create-namespace --version=v2024.2.25
```

## Introduction

This chart deploys a managed-serviceaccount on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `managed-serviceaccount`:

```bash
$ helm upgrade -i managed-serviceaccount appscode/managed-serviceaccount -n open-cluster-management --create-namespace --version=v2024.2.25
```

The command deploys a managed-serviceaccount on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `managed-serviceaccount`:

```bash
$ helm uninstall managed-serviceaccount -n open-cluster-management
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `managed-serviceaccount` chart and their default values.

|           Parameter            |                                                                                    Description                                                                                     |                               Default                               |
|--------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| image                          | Image of the managed service-account instances                                                                                                                                     | <code>quay.io/open-cluster-management/managed-serviceaccount</code> |
| tag                            |                                                                                                                                                                                    | <code></code>                                                       |
| agentInstallAll                |                                                                                                                                                                                    | <code>true</code>                                                   |
| replicas                       | Number of replicas                                                                                                                                                                 | <code>1</code>                                                      |
| enableAddOnDeploymentConfig    |                                                                                                                                                                                    | <code>false</code>                                                  |
| featureGates.ephemeralIdentity |                                                                                                                                                                                    | <code>false</code>                                                  |
| agentImagePullSecret           |                                                                                                                                                                                    | <code>""</code>                                                     |
| hubDeployMode                  | Hub deploy mode: AddOnTemplate or Deployment                                                                                                                                       | <code>Deployment</code>                                             |
| kubeconfigSecretName           | Name of the managed service-account addon template, only used when hubDeployMode is AddOnTemplate addOnTemplateName: managed-serviceaccount required for multicluster controlplane | <code>""</code>                                                     |
| kubectl.image                  |                                                                                                                                                                                    | <code>ghcr.io/appscode/kubectl:1.23</code>                          |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i managed-serviceaccount appscode/managed-serviceaccount -n open-cluster-management --create-namespace --version=v2024.2.25 --set image=quay.io/open-cluster-management/managed-serviceaccount
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i managed-serviceaccount appscode/managed-serviceaccount -n open-cluster-management --create-namespace --version=v2024.2.25 --values values.yaml
```
