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
	ResourceKindHubManagerRobot = "HubManagerRobot"
	ResourceHubManagerRobot     = "hubmanagerrobot"
	ResourceHubManagerRobots    = "hubmanagerrobots"
)

// HubManagerRobot defines the schama for HubManagerRobot.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=hubmanagerrobots,singular=hubmanagerrobot,categories={kubeops,appscode}
type HubManagerRobot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HubManagerRobotSpec `json:"spec,omitempty"`
}

// HubManagerRobotSpec is the schema for Identity Server values file
type HubManagerRobotSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	// +optional
	KubeconfigSecretName string `json:"kubeconfigSecretName"`
	// +optional
	AddonManagerNamespace string      `json:"addonManagerNamespace"`
	Kubectl               DockerImage `json:"kubectl"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HubManagerRobotList is a list of HubManagerRobots
type HubManagerRobotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HubManagerRobot CRD objects
	Items []HubManagerRobot `json:"items,omitempty"`
}
