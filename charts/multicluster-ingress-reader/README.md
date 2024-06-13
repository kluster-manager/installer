# Multicluster Ingress Reader

[Multicluster Ingress Reader](https://github.com/kluster-manager/installer) - Multicluster Ingress Reader

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/multicluster-ingress-reader --version=v2024.2.25
$ helm upgrade -i multicluster-ingress-reader appscode/multicluster-ingress-reader -n multicluster-controlplane --create-namespace --version=v2024.2.25
```

## Introduction

This chart deploys an Multicluster Ingress Reader on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `multicluster-ingress-reader`:

```bash
$ helm upgrade -i multicluster-ingress-reader appscode/multicluster-ingress-reader -n multicluster-controlplane --create-namespace --version=v2024.2.25
```

The command deploys an Multicluster Ingress Reader on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `multicluster-ingress-reader`:

```bash
$ helm uninstall multicluster-ingress-reader -n multicluster-controlplane
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `multicluster-ingress-reader` chart and their default values.

|         Parameter          |                                                      Description                                                       |                      Default                       |
|----------------------------|------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| nameOverride               |                                                                                                                        | <code>""</code>                                    |
| fullnameOverride           |                                                                                                                        | <code>""</code>                                    |
| imagePullSecrets           |                                                                                                                        | <code>[]</code>                                    |
| serviceAccount.create      | Specifies whether a service account should be created                                                                  | <code>true</code>                                  |
| serviceAccount.automount   | Automatically mount a ServiceAccount's API credentials?                                                                | <code>true</code>                                  |
| serviceAccount.annotations | Annotations to add to the service account                                                                              | <code>{}</code>                                    |
| serviceAccount.name        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template | <code>""</code>                                    |
| podSecurityContext         | podAnnotations: {} podLabels: {}                                                                                       | <code>{}</code>                                    |
| securityContext            |                                                                                                                        | <code>{}</code>                                    |
| resources                  |                                                                                                                        | <code>{}</code>                                    |
| nodeSelector               |                                                                                                                        | <code>{}</code>                                    |
| tolerations                |                                                                                                                        | <code>[]</code>                                    |
| affinity                   |                                                                                                                        | <code>{}</code>                                    |
| kubectl.image              |                                                                                                                        | <code>ghcr.io/appscode/kubectl-nonroot:1.25</code> |
| kubectl.pullPolicy         |                                                                                                                        | <code>IfNotPresent</code>                          |
| ingressServiceName         |                                                                                                                        | <code>ingress-nginx-mc-controller</code>           |
| secret.name                |                                                                                                                        | <code>multicluster-ingress-values</code>           |
| secret.namespace           |                                                                                                                        | <code>""</code>                                    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i multicluster-ingress-reader appscode/multicluster-ingress-reader -n multicluster-controlplane --create-namespace --version=v2024.2.25 --set kubectl.image=ghcr.io/appscode/kubectl-nonroot:1.25
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i multicluster-ingress-reader appscode/multicluster-ingress-reader -n multicluster-controlplane --create-namespace --version=v2024.2.25 --values values.yaml
```
