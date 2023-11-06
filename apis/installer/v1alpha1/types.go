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

type FeatureGate struct {
	// Feature is the key of feature gate. e.g. featuregate/Foo.
	// +kubebuilder:validation:Required
	// +required
	Feature string `json:"feature"`

	// Mode is either Enable, Disable, "" where "" is Disable by default.
	// In Enable mode, a valid feature gate `featuregate/Foo` will be set to "--featuregate/Foo=true".
	// In Disable mode, a valid feature gate `featuregate/Foo` will be set to "--featuregate/Foo=false".
	// +kubebuilder:default:=Disable
	// +kubebuilder:validation:Enum:=Enable;Disable
	// +optional
	Mode FeatureGateModeType `json:"mode,omitempty"`
}

type FeatureGateModeType string

const (
	// FeatureGateModeTypeEnable is the feature gate type to enable a feature.
	FeatureGateModeTypeEnable FeatureGateModeType = "Enable"
	// FeatureGateModeTypeDisable is the feature gate type to disable a feature.
	FeatureGateModeTypeDisable FeatureGateModeType = "Disable"
)
