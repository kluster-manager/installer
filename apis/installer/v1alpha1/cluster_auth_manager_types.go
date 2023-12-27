/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

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
	ResourceKindClusterAuthManager = "ClusterAuthManager"
	ResourceClusterAuthManager     = "kubeauthmanager"
	ResourceClusterAuthManagers    = "kubeauthmanagers"
)

// ClusterAuthManager defines the schama for ClusterAuthManager operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeauthmanagers,singular=kubeauthmanager,categories={kubeops,appscode}
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
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int       `json:"replicaCount"`
	RegistryFQDN     string    `json:"registryFQDN"`
	Image            Container `json:"image"`
	//+optional
	ImagePullSecrets []string           `json:"imagePullSecrets"`
	ImagePullPolicy  string             `json:"imagePullPolicy"`
	ServiceAccount   ServiceAccountSpec `json:"serviceAccount"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity   *core.Affinity `json:"affinity"`
	Monitoring Monitoring     `json:"monitoring"`
}

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Container struct {
	ImageRef `json:",inline"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

// +kubebuilder:validation:Enum=prometheus.io;prometheus.io/operator;prometheus.io/builtin
type MonitoringAgent string

type Monitoring struct {
	Agent          MonitoringAgent      `json:"agent"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
}

type ServiceMonitorLabels struct {
	// +optional
	Labels map[string]string `json:"labels"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterAuthManagerList is a list of ClusterAuthManagers
type ClusterAuthManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterAuthManager CRD objects
	Items []ClusterAuthManager `json:"items,omitempty"`
}
