# multicluster-controlplane

[multicluster-controlplane](https://github.com/kluster-manager/multicluster-controlplane) - A Helm chart for multicluster-controlplane

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable
$ helm repo update
$ helm search repo appscode/multicluster-controlplane --version=v2025.4.30
$ helm upgrade -i multicluster-controlplane appscode/multicluster-controlplane -n multicluster-controlplane --create-namespace --version=v2025.4.30
```

## Introduction

This chart deploys a multicluster controlplane on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `multicluster-controlplane`:

```bash
$ helm upgrade -i multicluster-controlplane appscode/multicluster-controlplane -n multicluster-controlplane --create-namespace --version=v2025.4.30
```

The command deploys a multicluster controlplane on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `multicluster-controlplane`:

```bash
$ helm uninstall multicluster-controlplane -n multicluster-controlplane
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `multicluster-controlplane` chart and their default values.

|                     Parameter                     |                       Description                        |                                    Default                                    |
|---------------------------------------------------|----------------------------------------------------------|-------------------------------------------------------------------------------|
| image                                             |                                                          | <code>quay.io/open-cluster-management/multicluster-controlplane:latest</code> |
| imagePullPolicy                                   |                                                          | <code>IfNotPresent</code>                                                     |
| replicas                                          |                                                          | <code>1</code>                                                                |
| features                                          |                                                          | <code>"DefaultClusterSet=true,ManagedClusterAutoApproval=true"</code>         |
| autoApprovalBootstrapUsers                        |                                                          | <code>""</code>                                                               |
| enableSelfManagement                              | TODO: should add restriction while enable selfmanagement | <code>false</code>                                                            |
| selfManagementClusterName                         |                                                          | <code>""</code>                                                               |
| enableDelegatingAuthentication                    |                                                          | <code>false</code>                                                            |
| apiserver.externalHostname                        |                                                          | <code>""</code>                                                               |
| apiserver.externalPort                            |                                                          | <code>443</code>                                                              |
| apiserver.ca                                      |                                                          | <code>""</code>                                                               |
| apiserver.cakey                                   |                                                          | <code>""</code>                                                               |
| apiserver.generateCA                              |                                                          | <code>false</code>                                                            |
| etcd.mode                                         |                                                          | <code>"embed"</code>                                                          |
| etcd.snapshotCount                                |                                                          | <code>5000</code>                                                             |
| etcd.servers                                      |                                                          | <code>[]</code>                                                               |
| etcd.ca                                           |                                                          | <code>""</code>                                                               |
| etcd.cert                                         |                                                          | <code>""</code>                                                               |
| etcd.certkey                                      |                                                          | <code>""</code>                                                               |
| pvc.storageCapacity                               |                                                          | <code>1Gi</code>                                                              |
| pvc.storageClassName                              |                                                          | <code>""</code>                                                               |
| pvc.selector                                      |                                                          | <code>{}</code>                                                               |
| route.enabled                                     |                                                          | <code>false</code>                                                            |
| loadbalancer.enabled                              |                                                          | <code>false</code>                                                            |
| loadbalancer.ip                                   |                                                          | <code>""</code>                                                               |
| nodeport.enabled                                  |                                                          | <code>false</code>                                                            |
| nodeport.port                                     |                                                          | <code>30443</code>                                                            |
| containerSecurityContext.allowPrivilegeEscalation |                                                          | <code>false</code>                                                            |
| containerSecurityContext.privileged               |                                                          | <code>false</code>                                                            |
| containerSecurityContext.runAsNonRoot             |                                                          | <code>true</code>                                                             |
| containerSecurityContext.seccompProfile.type      |                                                          | <code>RuntimeDefault</code>                                                   |
| securityContext                                   |                                                          | <code></code>                                                                 |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i multicluster-controlplane appscode/multicluster-controlplane -n multicluster-controlplane --create-namespace --version=v2025.4.30 --set image=quay.io/open-cluster-management/multicluster-controlplane:latest
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i multicluster-controlplane appscode/multicluster-controlplane -n multicluster-controlplane --create-namespace --version=v2025.4.30 --values values.yaml
```
