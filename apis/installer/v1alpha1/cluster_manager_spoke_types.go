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
	ResourceKindClusterManagerSpoke = "ClusterManagerSpoke"
	ResourceClusterManagerSpoke     = "clustermanagerspoke"
	ResourceClusterManagerSpokes    = "clustermanagerspokes"
)

// ClusterManagerSpoke defines the schama for ClusterManagerSpoke operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clustermanagerspokes,singular=clustermanagerspoke,categories={kubeops,appscode}
type ClusterManagerSpoke struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterManagerSpokeSpec `json:"spec,omitempty"`
}

// ClusterManagerSpokeSpec is the schema for Identity Server values file
type ClusterManagerSpokeSpec struct {
	// ClusterName: the name of the joined cluster on the hub
	ClusterName string `json:"clusterName"`

	// AgentNamespace: the namespace to deploy the agent
	AgentNamespace string `json:"agentNamespace"`

	// SpokeHub: Hub information
	// +optional
	Hub SpokeHub `json:"hub"`

	// Klusterlet is the klusterlet related configuration
	// +optional
	Klusterlet Klusterlet `json:"klusterlet"`

	// Registry is the image registry related configuration
	// +optional
	Registry string `json:"registry"`
	// bundle version
	// +optional
	BundleVersion SpokeBundleVersion `json:"bundleVersion"`
	// managed kubeconfig
	// +optional
	ManagedKubeconfig string `json:"managedKubeconfig"`

	// Features is the slice of feature for registration
	// +optional
	RegistrationFeatures []FeatureGate `json:"registrationFeatures"`

	// Features is the slice of feature for work
	// +optional
	WorkFeatures []FeatureGate `json:"workFeatures"`
}

// SpokeHub: The hub values for the template
type SpokeHub struct {
	// APIServer: The API Server external URL
	// +optional
	APIServer string `json:"apiServer"`
	// KubeConfig: The kubeconfig of the bootstrap secret to connect to the hub
	// +optional
	KubeConfig string `json:"kubeConfig"`
}

// Klusterlet is for templating klusterlet configuration
type Klusterlet struct {
	// APIServer: The API Server external URL
	// +optional
	APIServer string `json:"apiServer"`
	// +optional
	Mode string `json:"mode"`
	// +optional
	Name string `json:"name"`
	// +optional
	KlusterletNamespace string `json:"klusterletNamespace"`
}

type SpokeBundleVersion struct {
	// registration image version
	// +optional
	RegistrationImageVersion string `json:"registrationImageVersion"`
	// placement image version
	// +optional
	PlacementImageVersion string `json:"placementImageVersion"`
	// work image version
	// +optional
	WorkImageVersion string `json:"workImageVersion"`
	// operator image version
	// +optional
	OperatorImageVersion string `json:"operatorImageVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterManagerSpokeList is a list of ClusterManagerSpokes
type ClusterManagerSpokeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterManagerSpoke CRD objects
	Items []ClusterManagerSpoke `json:"items,omitempty"`
}
