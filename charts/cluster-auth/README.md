# Cluster Auth Agent

[Cluster Auth Agent](https://github.com/kluster-manager/cluster-auth) - Cluster Auth Agent

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/cluster-auth-agent --version=v2025.5.16
$ helm upgrade -i cluster-auth appscode/cluster-auth-agent -n open-cluster-management-cluster-auth --create-namespace --version=v2025.5.16
```

## Introduction

This chart deploys an Cluster Auth Agent on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-auth`:

```bash
$ helm upgrade -i cluster-auth appscode/cluster-auth-agent -n open-cluster-management-cluster-auth --create-namespace --version=v2025.5.16
```

The command deploys an Cluster Auth Agent on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-auth`:

```bash
$ helm uninstall cluster-auth -n open-cluster-management-cluster-auth
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-auth-agent` chart and their default values.

|            Parameter             |                                                                                                                Description                                                                                                                |                                                                                            Default                                                                                             |
|----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                     | Overrides name template                                                                                                                                                                                                                   | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                 | Overrides fullname template                                                                                                                                                                                                               | <code>""</code>                                                                                                                                                                                |
| replicaCount                     | Number of cluster-auth operator replicas to create (only 1 is supported)                                                                                                                                                                  | <code>1</code>                                                                                                                                                                                 |
| registryFQDN                     | Docker registry fqdn used to pull docker images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                            | <code>ghcr.io</code>                                                                                                                                                                           |
| image.registry                   | Docker registry used to pull operator image                                                                                                                                                                                               | <code>kluster-manager</code>                                                                                                                                                                   |
| image.repository                 | Name of operator container image                                                                                                                                                                                                          | <code>cluster-auth</code>                                                                                                                                                                      |
| image.tag                        | Operator container image tag                                                                                                                                                                                                              | <code>""</code>                                                                                                                                                                                |
| image.resources                  | Compute Resources required by the operator container                                                                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| image.securityContext            | Security options this container should run with                                                                                                                                                                                           | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/cluster-auth \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                                                                                                                                                                                |
| imagePullPolicy                  | Container image pull policy                                                                                                                                                                                                               | <code>IfNotPresent</code>                                                                                                                                                                      |
| criticalAddon                    | If true, installs cluster-auth operator as critical addon                                                                                                                                                                                 | <code>false</code>                                                                                                                                                                             |
| logLevel                         | Log level for operator                                                                                                                                                                                                                    | <code>3</code>                                                                                                                                                                                 |
| annotations                      | Annotations applied to operator deployment                                                                                                                                                                                                | <code>{}</code>                                                                                                                                                                                |
| podAnnotations                   | Annotations passed to operator pod(s).                                                                                                                                                                                                    | <code>{}</code>                                                                                                                                                                                |
| nodeSelector                     | Node labels for pod assignment                                                                                                                                                                                                            | <code>{"kubernetes.io/os":"linux"}</code>                                                                                                                                                      |
| tolerations                      | Tolerations for pod assignment                                                                                                                                                                                                            | <code>[]</code>                                                                                                                                                                                |
| affinity                         | Affinity rules for pod assignment                                                                                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| podSecurityContext               | Security options the operator pod should run with.                                                                                                                                                                                        | <code>{"fsGroup":65535}</code>                                                                                                                                                                 |
| serviceAccount.create            | Specifies whether a service account should be created                                                                                                                                                                                     | <code>true</code>                                                                                                                                                                              |
| serviceAccount.annotations       | Annotations to add to the service account                                                                                                                                                                                                 | <code>{}</code>                                                                                                                                                                                |
| serviceAccount.name              | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                    | <code></code>                                                                                                                                                                                  |
| monitoring.agent                 | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                      | <code>prometheus.io/operator</code>                                                                                                                                                            |
| monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                       | <code>{}</code>                                                                                                                                                                                |
| apiServer.healthcheck.enabled    |                                                                                                                                                                                                                                           | <code>false</code>                                                                                                                                                                             |
| hubKubeconfigSecretName          | Name of OCM Hub Kubeconfig secret                                                                                                                                                                                                         | <code>""</code>                                                                                                                                                                                |
| clusterName                      | We need to pass the cluster name because the OCM-MC host cluster doesn't have Klusterlet object.                                                                                                                                          | <code>""</code>                                                                                                                                                                                |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-auth appscode/cluster-auth-agent -n open-cluster-management-cluster-auth --create-namespace --version=v2025.5.16 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-auth appscode/cluster-auth-agent -n open-cluster-management-cluster-auth --create-namespace --version=v2025.5.16 --values values.yaml
```
