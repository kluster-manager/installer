#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

KLUSTER_MANAGER_FLUXCD_ADDON_TAG=${KLUSTER_MANAGER_FLUXCD_ADDON_TAG:-v0.0.2}
OPEN_CLUSTER_MANAGEMENT_IO_API_TAG=${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG:-v0.13.0}

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/operator/v1/0000_01_operator.open-cluster-management.io_clustermanagers.crd.yaml \
    --out=./charts/cluster-manager-hub/crds

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/operator/v1/0000_00_operator.open-cluster-management.io_klusterlets.crd.yaml \
    --out=./charts/cluster-manager-spoke/crds

crd-importer \
    --input=https://github.com/kluster-manager/fluxcd-addon/raw/${KLUSTER_MANAGER_FLUXCD_ADDON_TAG}/crds/fluxcd.open-cluster-management.io_fluxcdconfigs.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta1/0000_02_clusters.open-cluster-management.io_placements.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/cluster/v1beta2/0000_01_clusters.open-cluster-management.io_managedclustersetbindings.crd.yaml \
    --out=./charts/fluxcd-addon-manager/crds

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_01_addon.open-cluster-management.io_managedclusteraddons.crd.yaml \
    --out=./charts/managed-serviceaccount/crds

crd-importer \
    --input=https://github.com/open-cluster-management-io/api/raw/${OPEN_CLUSTER_MANAGEMENT_IO_API_TAG}/addon/v1alpha1/0000_00_addon.open-cluster-management.io_clustermanagementaddons.crd.yaml \
    --out=./charts/cluster-proxy/crds
