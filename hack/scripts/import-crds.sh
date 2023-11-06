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

KLUSTER_API_DOCKER_MACHINE_OPERATOR_TAG=${KLUSTER_API_DOCKER_MACHINE_OPERATOR_TAG:-master}

crd-importer \
    --input=https://github.com/kluster-api/docker-machine-operator/raw/${KLUSTER_API_DOCKER_MACHINE_OPERATOR_TAG}/crds/docker-machine.klusters.dev_drivers.yaml \
    --input=https://github.com/kluster-api/docker-machine-operator/raw/${KLUSTER_API_DOCKER_MACHINE_OPERATOR_TAG}/crds/docker-machine.klusters.dev_machines.yaml \
    --out=./charts/docker-machine-operator/crds
