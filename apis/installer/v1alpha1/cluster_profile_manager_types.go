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
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindClusterProfileManager = "ClusterProfileManager"
	ResourceClusterProfileManager     = "clusterprofilemanager"
	ResourceClusterProfileManagers    = "clusterprofilemanagers"
)

// ClusterProfileManager defines the schama for ClusterProfileManager operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusterprofilemanagers,singular=clusterprofilemanager,categories={kubeops,appscode}
type ClusterProfileManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterProfileManagerSpec `json:"spec,omitempty"`
}

// ClusterProfileManagerSpec is the schema for Identity Server values file
type ClusterProfileManagerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string        `json:"fullnameOverride"`
	RegistryFQDN     string        `json:"registryFQDN"`
	Managed          AppImage      `json:"manager"`
	Placement        PlacementSpec `json:"placement"`

	// opscenter-feature value

	shared.BootstrapPresets `json:",inline,omitempty"`
	Platform                AceOcmAddonsPlatformSpec `json:"platform"`
}

type AceOcmAddonsPlatformSpec struct {
	BaseURL string `json:"baseURL"`
	// +optional
	Token string `json:"token"`
	// +optional
	CaBundle string `json:"caBundle"`
}

type AppImage struct {
	Image string `json:"image"`
	// +optional
	Tag             string                `json:"tag"`
	ImagePullPolicy string                `json:"imagePullPolicy"`
	SecurityContext *core.SecurityContext `json:"securityContext"`
	// List of sources to populate environment variables in the container.
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys
	// will be reported as an event when the container is starting. When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	// +listType=atomic
	EnvFrom []core.EnvFromSource `json:"envFrom"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Env []core.EnvVar `json:"env"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterProfileManagerList is a list of ClusterProfileManagers
type ClusterProfileManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterProfileManager CRD objects
	Items []ClusterProfileManager `json:"items,omitempty"`
}
