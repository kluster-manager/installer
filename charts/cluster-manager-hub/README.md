# Cluster Manager Hub

[Cluster Manager Hub](https://github.com/kluster-manager/installer) - Create a Open Cluster Manager Hub

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/cluster-manager-hub --version=v2024.12.26
$ helm upgrade -i cluster-manager-hub appscode/cluster-manager-hub -n open-cluster-management --create-namespace --version=v2024.12.26
```

## Introduction

This chart deploys a Cluster Manager Hub on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-manager-hub`:

```bash
$ helm upgrade -i cluster-manager-hub appscode/cluster-manager-hub -n open-cluster-management --create-namespace --version=v2024.12.26
```

The command deploys a Cluster Manager Hub on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-manager-hub`:

```bash
$ helm uninstall cluster-manager-hub -n open-cluster-management
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-manager-hub` chart and their default values.

|               Parameter                | Description  |                                                                              Default                                                                              |
|----------------------------------------|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| hub.useBootstrapToken                  |              | <code>false</code>                                                                                                                                                |
| hub.tokenID                            | randAlpha 6  | <code>""</code>                                                                                                                                                   |
| hub.tokenSecret                        | randAlpha 16 | <code>""</code>                                                                                                                                                   |
| hub.registry                           |              | <code>quay.io/open-cluster-management</code>                                                                                                                      |
| hub.securityContext                    |              | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"privileged":false,"runAsNonRoot":true,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| bundleVersion.registrationImageVersion |              | <code>""</code>                                                                                                                                                   |
| bundleVersion.placementImageVersion    |              | <code>""</code>                                                                                                                                                   |
| bundleVersion.workImageVersion         |              | <code>""</code>                                                                                                                                                   |
| bundleVersion.operatorImageVersion     |              | <code>""</code>                                                                                                                                                   |
| bundleVersion.addonManagerImageVersion |              | <code>""</code>                                                                                                                                                   |
| autoApprove                            |              | <code>false</code>                                                                                                                                                |
| registrationFeatures                   |              | <code>[{"feature":"DefaultClusterSet","mode":"Enable"}]</code>                                                                                                    |
| workFeatures                           |              | <code>[{"feature":"ManifestWorkReplicaSet","mode":"Enable"}]</code>                                                                                               |
| addonFeatures                          |              | <code>[{"feature":"AddonManagement","mode":"Enable"}]</code>                                                                                                      |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-manager-hub appscode/cluster-manager-hub -n open-cluster-management --create-namespace --version=v2024.12.26 --set hub.registry=quay.io/open-cluster-management
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-manager-hub appscode/cluster-manager-hub -n open-cluster-management --create-namespace --version=v2024.12.26 --values values.yaml
```
