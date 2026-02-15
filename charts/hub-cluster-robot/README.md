# Hub Cluster Robot

[Hub Cluster Robot](https://github.com/kluster-manager) - Hub Cluster Robot

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/hub-cluster-robot --version=v2026.2.16
$ helm upgrade -i hub-cluster-robot appscode/hub-cluster-robot -n open-cluster-management --create-namespace --version=v2026.2.16
```

## Introduction

This chart deploys a Hub Cluster Robot on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `hub-cluster-robot`:

```bash
$ helm upgrade -i hub-cluster-robot appscode/hub-cluster-robot -n open-cluster-management --create-namespace --version=v2026.2.16
```

The command deploys a Hub Cluster Robot on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `hub-cluster-robot`:

```bash
$ helm uninstall hub-cluster-robot -n open-cluster-management
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `hub-cluster-robot` chart and their default values.

|    Parameter     | Description |     Default     |
|------------------|-------------|-----------------|
| nameOverride     |             | <code>""</code> |
| fullnameOverride |             | <code>""</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i hub-cluster-robot appscode/hub-cluster-robot -n open-cluster-management --create-namespace --version=v2026.2.16 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i hub-cluster-robot appscode/hub-cluster-robot -n open-cluster-management --create-namespace --version=v2026.2.16 --values values.yaml
```
