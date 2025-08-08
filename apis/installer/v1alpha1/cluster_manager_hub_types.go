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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindClusterManagerHub = "ClusterManagerHub"
	ResourceClusterManagerHub     = "clustermanagerhub"
	ResourceClusterManagerHubs    = "clustermanagerhubs"
)

// ClusterManagerHub defines the schama for ClusterManagerHub operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clustermanagerhubs,singular=clustermanagerhub,categories={kubeops,appscode}
type ClusterManagerHub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterManagerHubSpec `json:"spec,omitempty"`
}

// ClusterManagerHubSpec is the schema for Identity Server values file
type ClusterManagerHubSpec struct {
	// The values related to the hub
	// +optional
	Hub HubInfo `json:"hub"`

	// +optional
	Aws AwsHubSpec `json:"aws"`

	// bundle version
	// +optional
	BundleVersion HubBundleVersion `json:"bundleVersion"`

	// if enable auto approve
	// +optional
	AutoApprove bool `json:"autoApprove"`

	// Features is the slice of feature for registration
	// +optional
	RegistrationFeatures []FeatureGate `json:"registrationFeatures"`

	// Features is the slice of feature for work
	// +optional
	WorkFeatures []FeatureGate `json:"workFeatures"`

	// Features is the slice of feature for addon manager
	// +optional
	AddonFeatures []FeatureGate `json:"addonFeatures"`
}

// HubInfo: The hub values for the template
type HubInfo struct {
	// If true the bootstrap token will be used instead of the service account token
	// +optional
	UseBootstrapToken bool `json:"useBootstrapToken"`
	// TokenID: A token id allowing the cluster to connect back to the hub
	// +optional
	TokenID string `json:"tokenID"`
	// TokenSecret: A token secret allowing the cluster to connect back to the hub
	// +optional
	TokenSecret string `json:"tokenSecret"`
	// Registry is the name of the image registry to pull.
	// +optional
	Registry string `json:"registry"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type AwsHubSpec struct {
	// +optional
	HubClusterArn string `json:"hubClusterArn"`
}

type HubBundleVersion struct {
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
	// addon manager image version
	// +optional
	AddonManagerImageVersion string `json:"addonManagerImageVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterManagerHubList is a list of ClusterManagerHubs
type ClusterManagerHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterManagerHub CRD objects
	Items []ClusterManagerHub `json:"items,omitempty"`
}
