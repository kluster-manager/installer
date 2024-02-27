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
	ResourceKindClusterAuthManager = "ClusterAuthManager"
	ResourceClusterAuthManager     = "clusterauthmanager"
	ResourceClusterAuthManagers    = "clusterauthmanagers"
)

// ClusterAuthManager defines the schama for ClusterAuthManager operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusterauthmanagers,singular=clusterauthmanager,categories={kubeops,appscode}
type ClusterAuthManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterAuthManagerSpec `json:"spec,omitempty"`
}

// ClusterAuthManagerSpec is the schema for Identity Server values file
type ClusterAuthManagerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	RegistryFQDN     string `json:"registryFQDN"`
	Image            string `json:"image"`
	// +optional
	Tag             string `json:"tag"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	// +optional
	KubeconfigSecretName string      `json:"kubeconfigSecretName"`
	Kubectl              DockerImage `json:"kubectl"`
}

type DockerImage struct {
	Image string `json:"image"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterAuthManagerList is a list of ClusterAuthManagers
type ClusterAuthManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterAuthManager CRD objects
	Items []ClusterAuthManager `json:"items,omitempty"`
}
