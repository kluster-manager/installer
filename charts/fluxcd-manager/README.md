# FluxCD Manager

[FluxCD Manager](https://github.com/kluster-manager/fluxcd-addon) - Installs FluxCD in managed clusters

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/fluxcd-manager --version=v2024.2.25
$ helm upgrade -i fluxcd-manager appscode/fluxcd-manager -n open-cluster-management-addon --create-namespace --version=v2024.2.25
```

## Introduction

This chart deploys a FluxCD Manager on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `fluxcd-manager`:

```bash
$ helm upgrade -i fluxcd-manager appscode/fluxcd-manager -n open-cluster-management-addon --create-namespace --version=v2024.2.25
```

The command deploys a FluxCD Manager on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `fluxcd-manager`:

```bash
$ helm uninstall fluxcd-manager -n open-cluster-management-addon
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `fluxcd-manager` chart and their default values.

|                   Parameter                   | Description |                      Default                       |
|-----------------------------------------------|-------------|----------------------------------------------------|
| nameOverride                                  |             | <code>""</code>                                    |
| fullnameOverride                              |             | <code>""</code>                                    |
| image                                         |             | <code>ghcr.io/kluster-manager/fluxcd-addon</code>  |
| tag                                           |             | <code></code>                                      |
| kubeconfigSecretName                          |             | <code>""</code>                                    |
| addonManagerNamespace                         |             | <code>open-cluster-management-fluxcd</code>        |
| placement.create                              |             | <code>true</code>                                  |
| placement.name                                |             | <code>global</code>                                |
| kubectl.image                                 |             | <code>ghcr.io/appscode/kubectl-nonroot:1.25</code> |
| fluxcdConfig.installCRDs                      |             | <code>true</code>                                  |
| fluxcdConfig.helmController.create            |             | <code>true</code>                                  |
| fluxcdConfig.sourceController.create          |             | <code>true</code>                                  |
| fluxcdConfig.imageReflectionController.create |             | <code>false</code>                                 |
| fluxcdConfig.imageAutomationController.create |             | <code>false</code>                                 |
| fluxcdConfig.kustomizeController.create       |             | <code>false</code>                                 |
| fluxcdConfig.notificationController.create    |             | <code>false</code>                                 |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i fluxcd-manager appscode/fluxcd-manager -n open-cluster-management-addon --create-namespace --version=v2024.2.25 --set image=ghcr.io/kluster-manager/fluxcd-addon
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i fluxcd-manager appscode/fluxcd-manager -n open-cluster-management-addon --create-namespace --version=v2024.2.25 --values values.yaml
```
