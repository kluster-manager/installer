# cluster-proxy

[cluster-proxy](https://github.com/kluster-manager/cluster-proxy) - Installs cluster-proxy in managed clusters

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/cluster-proxy --version=0.3.0
$ helm upgrade -i cluster-proxy appscode/cluster-proxy -n open-cluster-management --create-namespace --version=0.3.0
```

## Introduction

This chart deploys a cluster-proxy on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-proxy`:

```bash
$ helm upgrade -i cluster-proxy appscode/cluster-proxy -n open-cluster-management --create-namespace --version=0.3.0
```

The command deploys a cluster-proxy on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-proxy`:

```bash
$ helm uninstall cluster-proxy -n open-cluster-management
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-proxy` chart and their default values.

|             Parameter              |              Description               |                       Default                        |
|------------------------------------|----------------------------------------|------------------------------------------------------|
| registry                           | Image registry                         | <code>ghcr.io/kluster-manager</code>                 |
| image                              | Image of the cluster-gateway instances | <code>cluster-proxy</code>                           |
| tag                                | Image tag                              | <code>latest</code>                                  |
| replicas                           | Number of replicas                     | <code>1</code>                                       |
| spokeAddonNamespace                |                                        | <code>"open-cluster-management-cluster-proxy"</code> |
| proxyServerImage                   |                                        | <code>ghcr.io/kluster-manager/cluster-proxy</code>   |
| proxyAgentImage                    |                                        | <code>ghcr.io/kluster-manager/cluster-proxy</code>   |
| proxyServer.entrypointLoadBalancer |                                        | <code>false</code>                                   |
| proxyServer.entrypointAddress      |                                        | <code>""</code>                                      |
| proxyServer.port                   |                                        | <code>8091</code>                                    |
| kubeconfigSecretName               | required for multicluster controlplane | <code>""</code>                                      |
| kubectl.image                      |                                        | <code>ghcr.io/appscode/kubectl:1.23</code>           |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-proxy appscode/cluster-proxy -n open-cluster-management --create-namespace --version=0.3.0 --set registry=ghcr.io/kluster-manager
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-proxy appscode/cluster-proxy -n open-cluster-management --create-namespace --version=0.3.0 --values values.yaml
```
