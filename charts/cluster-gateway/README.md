# cluster-gateway

[cluster-gateway](https://github.com/kluster-manager/cluster-gateway) - Installs cluster-gateway in managed clusters

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/cluster-gateway --version=v2026.2.16
$ helm upgrade -i cluster-gateway appscode/cluster-gateway -n open-cluster-management --create-namespace --version=v2026.2.16
```

## Introduction

This chart deploys a cluster-gateway on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-gateway`:

```bash
$ helm upgrade -i cluster-gateway appscode/cluster-gateway -n open-cluster-management --create-namespace --version=v2026.2.16
```

The command deploys a cluster-gateway on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-gateway`:

```bash
$ helm uninstall cluster-gateway -n open-cluster-management
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-gateway` chart and their default values.

|         Parameter          |              Description               |                               Default                               |
|----------------------------|----------------------------------------|---------------------------------------------------------------------|
| image                      | Image of the cluster-gateway instances | <code>ghcr.io/kluster-manager/cluster-gateway</code>                |
| tag                        |                                        | <code></code>                                                       |
| replicas                   | Number of replicas                     | <code>1</code>                                                      |
| clusterProxy.enabled       |                                        | <code>true</code>                                                   |
| clusterProxy.endpoint.host |                                        | <code>proxy-entrypoint.open-cluster-management-cluster-proxy</code> |
| clusterProxy.endpoint.port |                                        | <code>8090</code>                                                   |
| featureGate.healthiness    |                                        | <code>false</code>                                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-gateway appscode/cluster-gateway -n open-cluster-management --create-namespace --version=v2026.2.16 --set image=ghcr.io/kluster-manager/cluster-gateway
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-gateway appscode/cluster-gateway -n open-cluster-management --create-namespace --version=v2026.2.16 --values values.yaml
```
