# Spoke Cluster Addons

[Spoke Cluster Addons](https://github.com/kluster-manager/installer) - Spoke Cluster Addons

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/spoke-cluster-addons --version=v2025.8.31
$ helm upgrade -i spoke-cluster-addons appscode/spoke-cluster-addons -n c1 --create-namespace --version=v2025.8.31
```

## Introduction

This chart deploys a Spoke Cluster Addons on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `spoke-cluster-addons`:

```bash
$ helm upgrade -i spoke-cluster-addons appscode/spoke-cluster-addons -n c1 --create-namespace --version=v2025.8.31
```

The command deploys a Spoke Cluster Addons on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `spoke-cluster-addons`:

```bash
$ helm uninstall spoke-cluster-addons -n c1
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `spoke-cluster-addons` chart and their default values.

|                Parameter                 | Description |      Default       |
|------------------------------------------|-------------|--------------------|
| nameOverride                             |             | <code>""</code>    |
| fullnameOverride                         |             | <code>""</code>    |
| addons.cluster-proxy.enabled             |             | <code>false</code> |
| addons.kubeslice.enabled                 |             | <code>false</code> |
| addons.kubeslice.values.projectNamespace |             | <code>""</code>    |
| addons.kubeslice.values.networkInterface |             | <code>""</code>    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i spoke-cluster-addons appscode/spoke-cluster-addons -n c1 --create-namespace --version=v2025.8.31 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i spoke-cluster-addons appscode/spoke-cluster-addons -n c1 --create-namespace --version=v2025.8.31 --values values.yaml
```
