/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindSpokeClusterAddons = "SpokeClusterAddons"
	ResourceSpokeClusterAddons     = "spokeclusteraddons"
	ResourceSpokeClusterAddonss    = "spokeclusteraddonss"
)

// SpokeClusterAddons defines the schama for SpokeClusterAddons.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=spokeClusterAddonss,singular=spokeClusterAddons,categories={kubeops,appscode}
type SpokeClusterAddons struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpokeClusterAddonsSpec `json:"spec,omitempty"`
}

// SpokeClusterAddonsSpec is the schema for Identity Server values file
type SpokeClusterAddonsSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	// +optional
	KubeconfigSecretName string      `json:"kubeconfigSecretName"`
	Kubectl              DockerImage `json:"kubectl"`
	Addons               Addons      `json:"addons"`
}

type Addons struct {
	ClusterProxy AddonActivation `json:"cluster-proxy"`
	Kubeslice    AddonKubeslice  `json:"kubeslice"`
}

type AddonActivation struct {
	Enabled bool `json:"enabled"`
}

type AddonKubeslice struct {
	AddonActivation `json:",inline"`
	Values          AddonKubesliceValues `json:"values"`
}

type AddonKubesliceValues struct {
	ProjectNamespace string `json:"projectNamespace"`
	NetworkInterface string `json:"networkInterface"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpokeClusterAddonsList is a list of SpokeClusterAddonss
type SpokeClusterAddonsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpokeClusterAddons CRD objects
	Items []SpokeClusterAddons `json:"items,omitempty"`
}
