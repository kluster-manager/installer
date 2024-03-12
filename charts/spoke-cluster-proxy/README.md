# Spoke Cluster Proxy

[Spoke Cluster Proxy](https://github.com/kluster-manager/cluster-proxy) - Spoke Cluster Proxy

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/spoke-cluster-proxy --version=v2024.2.25
$ helm upgrade -i spoke-cluster-proxy appscode/spoke-cluster-proxy -n c1 --create-namespace --version=v2024.2.25
```

## Introduction

This chart deploys a Spoke Cluster Proxy on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `spoke-cluster-proxy`:

```bash
$ helm upgrade -i spoke-cluster-proxy appscode/spoke-cluster-proxy -n c1 --create-namespace --version=v2024.2.25
```

The command deploys a Spoke Cluster Proxy on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `spoke-cluster-proxy`:

```bash
$ helm uninstall spoke-cluster-proxy -n c1
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `spoke-cluster-proxy` chart and their default values.

|      Parameter       | Description |                      Default                       |
|----------------------|-------------|----------------------------------------------------|
| nameOverride         |             | <code>""</code>                                    |
| fullnameOverride     |             | <code>""</code>                                    |
| kubeconfigSecretName |             | <code>""</code>                                    |
| kubectl.image        |             | <code>ghcr.io/appscode/kubectl-nonroot:1.25</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i spoke-cluster-proxy appscode/spoke-cluster-proxy -n c1 --create-namespace --version=v2024.2.25 --set kubectl.image=ghcr.io/appscode/kubectl-nonroot:1.25
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i spoke-cluster-proxy appscode/spoke-cluster-proxy -n c1 --create-namespace --version=v2024.2.25 --values values.yaml
```
