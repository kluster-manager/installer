# Cluster Auth Manager

[Cluster Auth Manager](https://github.com/kluster-manager/cluster-auth) - Cluster Auth Manager

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/cluster-auth-manager --version=v2024.10.7
$ helm upgrade -i cluster-auth-manager appscode/cluster-auth-manager -n open-cluster-management-addon --create-namespace --version=v2024.10.7
```

## Introduction

This chart deploys an Cluster Auth Manager on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-auth-manager`:

```bash
$ helm upgrade -i cluster-auth-manager appscode/cluster-auth-manager -n open-cluster-management-addon --create-namespace --version=v2024.10.7
```

The command deploys an Cluster Auth Manager on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-auth-manager`:

```bash
$ helm uninstall cluster-auth-manager -n open-cluster-management-addon
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-auth-manager` chart and their default values.

|       Parameter       |                             Description                             |                      Default                       |
|-----------------------|---------------------------------------------------------------------|----------------------------------------------------|
| nameOverride          |                                                                     | <code>""</code>                                    |
| fullnameOverride      |                                                                     | <code>""</code>                                    |
| registryFQDN          | Docker registry fqdn used to pull license-proxyserver docker images | <code>ghcr.io</code>                               |
| image                 |                                                                     | <code>ghcr.io/kluster-manager/cluster-auth</code>  |
| tag                   |                                                                     | <code>""</code>                                    |
| imagePullPolicy       |                                                                     | <code>Always</code>                                |
| kubeconfigSecretName  |                                                                     | <code>""</code>                                    |
| addonManagerNamespace |                                                                     | <code>open-cluster-management-cluster-auth</code>  |
| placement.create      |                                                                     | <code>true</code>                                  |
| placement.name        |                                                                     | <code>global</code>                                |
| kubectl.image         |                                                                     | <code>ghcr.io/appscode/kubectl-nonroot:1.31</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-auth-manager appscode/cluster-auth-manager -n open-cluster-management-addon --create-namespace --version=v2024.10.7 --set registryFQDN=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-auth-manager appscode/cluster-auth-manager -n open-cluster-management-addon --create-namespace --version=v2024.10.7 --values values.yaml
```
