# Cluster Manager Spoke

[Cluster Manager Spoke](https://github.com/kluster-manager/installer) - Join a Open Cluster Manager Hub

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/cluster-manager-spoke --version=v2024.2.25
$ helm upgrade -i cluster-manager-spoke appscode/cluster-manager-spoke -n open-cluster-management --create-namespace --version=v2024.2.25
```

## Introduction

This chart deploys a Cluster Manager Spoke on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-manager-spoke`:

```bash
$ helm upgrade -i cluster-manager-spoke appscode/cluster-manager-spoke -n open-cluster-management --create-namespace --version=v2024.2.25
```

The command deploys a Cluster Manager Spoke on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-manager-spoke`:

```bash
$ helm uninstall cluster-manager-spoke -n open-cluster-management
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-manager-spoke` chart and their default values.

|               Parameter                |                                                      Description                                                       |                      Default                       |
|----------------------------------------|------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| clusterMetadata.uid                    |                                                                                                                        | <code>TBD</code>                                   |
| clusterMetadata.name                   |                                                                                                                        | <code>TBD</code>                                   |
| clusterMetadata.clusterManagers        |                                                                                                                        | <code>[]</code>                                    |
| clusterMetadata.capi.provider          |                                                                                                                        | <code>""</code>                                    |
| clusterMetadata.capi.namespace         |                                                                                                                        | <code>""</code>                                    |
| hub.apiServer                          |                                                                                                                        | <code>""</code>                                    |
| hub.caData                             |                                                                                                                        | <code>""</code>                                    |
| hub.token                              |                                                                                                                        | <code>""</code>                                    |
| hub.kubeConfig                         |                                                                                                                        | <code>""</code>                                    |
| registry                               |                                                                                                                        | <code>quay.io/open-cluster-management</code>       |
| bundleVersion.registrationImageVersion |                                                                                                                        | <code>""</code>                                    |
| bundleVersion.placementImageVersion    |                                                                                                                        | <code>""</code>                                    |
| bundleVersion.workImageVersion         |                                                                                                                        | <code>""</code>                                    |
| bundleVersion.operatorImageVersion     |                                                                                                                        | <code>""</code>                                    |
| bundleVersion.clusteradmImageVersion   |                                                                                                                        | <code>latest</code>                                |
| managedKubeconfig                      |                                                                                                                        | <code>''</code>                                    |
| registrationFeatures                   |                                                                                                                        | <code>[]</code>                                    |
| workFeatures                           |                                                                                                                        | <code>[]</code>                                    |
| imagePullSecrets                       |                                                                                                                        | <code>[]</code>                                    |
| serviceAccount.create                  | Specifies whether a service account should be created                                                                  | <code>true</code>                                  |
| serviceAccount.annotations             | Annotations to add to the service account                                                                              | <code>{}</code>                                    |
| serviceAccount.name                    | The name of the service account to use. If not set and create is true, a name is generated using the fullname template | <code>""</code>                                    |
| podSecurityContext                     | podAnnotations: {} podLabels: {}                                                                                       | <code>{}</code>                                    |
| securityContext                        |                                                                                                                        | <code>{}</code>                                    |
| resources                              |                                                                                                                        | <code>{}</code>                                    |
| nodeSelector                           |                                                                                                                        | <code>{}</code>                                    |
| tolerations                            |                                                                                                                        | <code>[]</code>                                    |
| affinity                               |                                                                                                                        | <code>{}</code>                                    |
| kubectl.image                          |                                                                                                                        | <code>ghcr.io/appscode/kubectl-nonroot:1.25</code> |
| kubectl.pullPolicy                     |                                                                                                                        | <code>IfNotPresent</code>                          |
| secret.name                            |                                                                                                                        | <code>opscenter-features-values</code>             |
| secret.namespace                       |                                                                                                                        | <code>kubeops</code>                               |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-manager-spoke appscode/cluster-manager-spoke -n open-cluster-management --create-namespace --version=v2024.2.25 --set clusterMetadata.uid=TBD
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-manager-spoke appscode/cluster-manager-spoke -n open-cluster-management --create-namespace --version=v2024.2.25 --values values.yaml
```
