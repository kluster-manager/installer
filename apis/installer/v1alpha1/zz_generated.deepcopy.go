//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AceOcmAddonsPlatformSpec) DeepCopyInto(out *AceOcmAddonsPlatformSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AceOcmAddonsPlatformSpec.
func (in *AceOcmAddonsPlatformSpec) DeepCopy() *AceOcmAddonsPlatformSpec {
	if in == nil {
		return nil
	}
	out := new(AceOcmAddonsPlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiServerSpec) DeepCopyInto(out *ApiServerSpec) {
	*out = *in
	out.Healthcheck = in.Healthcheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiServerSpec.
func (in *ApiServerSpec) DeepCopy() *ApiServerSpec {
	if in == nil {
		return nil
	}
	out := new(ApiServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppImage) DeepCopyInto(out *AppImage) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppImage.
func (in *AppImage) DeepCopy() *AppImage {
	if in == nil {
		return nil
	}
	out := new(AppImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuth) DeepCopyInto(out *ClusterAuth) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuth.
func (in *ClusterAuth) DeepCopy() *ClusterAuth {
	if in == nil {
		return nil
	}
	out := new(ClusterAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthList) DeepCopyInto(out *ClusterAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthList.
func (in *ClusterAuthList) DeepCopy() *ClusterAuthList {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthManager) DeepCopyInto(out *ClusterAuthManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthManager.
func (in *ClusterAuthManager) DeepCopy() *ClusterAuthManager {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuthManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthManagerList) DeepCopyInto(out *ClusterAuthManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAuthManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthManagerList.
func (in *ClusterAuthManagerList) DeepCopy() *ClusterAuthManagerList {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuthManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthManagerSpec) DeepCopyInto(out *ClusterAuthManagerSpec) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	out.Placement = in.Placement
	out.Kubectl = in.Kubectl
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthManagerSpec.
func (in *ClusterAuthManagerSpec) DeepCopy() *ClusterAuthManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthSpec) DeepCopyInto(out *ClusterAuthSpec) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	out.ApiServer = in.ApiServer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthSpec.
func (in *ClusterAuthSpec) DeepCopy() *ClusterAuthSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerHub) DeepCopyInto(out *ClusterManagerHub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerHub.
func (in *ClusterManagerHub) DeepCopy() *ClusterManagerHub {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagerHub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerHubList) DeepCopyInto(out *ClusterManagerHubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterManagerHub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerHubList.
func (in *ClusterManagerHubList) DeepCopy() *ClusterManagerHubList {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerHubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagerHubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerHubSpec) DeepCopyInto(out *ClusterManagerHubSpec) {
	*out = *in
	in.Hub.DeepCopyInto(&out.Hub)
	out.BundleVersion = in.BundleVersion
	if in.RegistrationFeatures != nil {
		in, out := &in.RegistrationFeatures, &out.RegistrationFeatures
		*out = make([]FeatureGate, len(*in))
		copy(*out, *in)
	}
	if in.WorkFeatures != nil {
		in, out := &in.WorkFeatures, &out.WorkFeatures
		*out = make([]FeatureGate, len(*in))
		copy(*out, *in)
	}
	if in.AddonFeatures != nil {
		in, out := &in.AddonFeatures, &out.AddonFeatures
		*out = make([]FeatureGate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerHubSpec.
func (in *ClusterManagerHubSpec) DeepCopy() *ClusterManagerHubSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerHubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerSpoke) DeepCopyInto(out *ClusterManagerSpoke) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerSpoke.
func (in *ClusterManagerSpoke) DeepCopy() *ClusterManagerSpoke {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerSpoke)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagerSpoke) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerSpokeList) DeepCopyInto(out *ClusterManagerSpokeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterManagerSpoke, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerSpokeList.
func (in *ClusterManagerSpokeList) DeepCopy() *ClusterManagerSpokeList {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerSpokeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagerSpokeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagerSpokeSpec) DeepCopyInto(out *ClusterManagerSpokeSpec) {
	*out = *in
	out.ClusterMetadata = in.ClusterMetadata
	out.Hub = in.Hub
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	out.BundleVersion = in.BundleVersion
	if in.RegistrationFeatures != nil {
		in, out := &in.RegistrationFeatures, &out.RegistrationFeatures
		*out = make([]FeatureGate, len(*in))
		copy(*out, *in)
	}
	if in.WorkFeatures != nil {
		in, out := &in.WorkFeatures, &out.WorkFeatures
		*out = make([]FeatureGate, len(*in))
		copy(*out, *in)
	}
	in.Clusteradm.DeepCopyInto(&out.Clusteradm)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagerSpokeSpec.
func (in *ClusterManagerSpokeSpec) DeepCopy() *ClusterManagerSpokeSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterManagerSpokeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMetadata) DeepCopyInto(out *ClusterMetadata) {
	*out = *in
	out.Store = in.Store
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMetadata.
func (in *ClusterMetadata) DeepCopy() *ClusterMetadata {
	if in == nil {
		return nil
	}
	out := new(ClusterMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMetadataStore) DeepCopyInto(out *ClusterMetadataStore) {
	*out = *in
	out.ClusterClaim = in.ClusterClaim
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMetadataStore.
func (in *ClusterMetadataStore) DeepCopy() *ClusterMetadataStore {
	if in == nil {
		return nil
	}
	out := new(ClusterMetadataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfileManager) DeepCopyInto(out *ClusterProfileManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfileManager.
func (in *ClusterProfileManager) DeepCopy() *ClusterProfileManager {
	if in == nil {
		return nil
	}
	out := new(ClusterProfileManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterProfileManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfileManagerList) DeepCopyInto(out *ClusterProfileManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterProfileManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfileManagerList.
func (in *ClusterProfileManagerList) DeepCopy() *ClusterProfileManagerList {
	if in == nil {
		return nil
	}
	out := new(ClusterProfileManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterProfileManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfileManagerSpec) DeepCopyInto(out *ClusterProfileManagerSpec) {
	*out = *in
	in.Managed.DeepCopyInto(&out.Managed)
	out.Placement = in.Placement
	out.Kubectl = in.Kubectl
	in.BootstrapPresets.DeepCopyInto(&out.BootstrapPresets)
	out.Platform = in.Platform
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfileManagerSpec.
func (in *ClusterProfileManagerSpec) DeepCopy() *ClusterProfileManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterProfileManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusteradmSpec) DeepCopyInto(out *ClusteradmSpec) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusteradmSpec.
func (in *ClusteradmSpec) DeepCopy() *ClusteradmSpec {
	if in == nil {
		return nil
	}
	out := new(ClusteradmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	out.ImageRef = in.ImageRef
	in.Resources.DeepCopyInto(&out.Resources)
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerImage) DeepCopyInto(out *DockerImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerImage.
func (in *DockerImage) DeepCopy() *DockerImage {
	if in == nil {
		return nil
	}
	out := new(DockerImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGate) DeepCopyInto(out *FeatureGate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGate.
func (in *FeatureGate) DeepCopy() *FeatureGate {
	if in == nil {
		return nil
	}
	out := new(FeatureGate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcheckSpec) DeepCopyInto(out *HealthcheckSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcheckSpec.
func (in *HealthcheckSpec) DeepCopy() *HealthcheckSpec {
	if in == nil {
		return nil
	}
	out := new(HealthcheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubBundleVersion) DeepCopyInto(out *HubBundleVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubBundleVersion.
func (in *HubBundleVersion) DeepCopy() *HubBundleVersion {
	if in == nil {
		return nil
	}
	out := new(HubBundleVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubClusterRobot) DeepCopyInto(out *HubClusterRobot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubClusterRobot.
func (in *HubClusterRobot) DeepCopy() *HubClusterRobot {
	if in == nil {
		return nil
	}
	out := new(HubClusterRobot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubClusterRobot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubClusterRobotList) DeepCopyInto(out *HubClusterRobotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HubClusterRobot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubClusterRobotList.
func (in *HubClusterRobotList) DeepCopy() *HubClusterRobotList {
	if in == nil {
		return nil
	}
	out := new(HubClusterRobotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubClusterRobotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubClusterRobotSpec) DeepCopyInto(out *HubClusterRobotSpec) {
	*out = *in
	out.Kubectl = in.Kubectl
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubClusterRobotSpec.
func (in *HubClusterRobotSpec) DeepCopy() *HubClusterRobotSpec {
	if in == nil {
		return nil
	}
	out := new(HubClusterRobotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubInfo) DeepCopyInto(out *HubInfo) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubInfo.
func (in *HubInfo) DeepCopy() *HubInfo {
	if in == nil {
		return nil
	}
	out := new(HubInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRef) DeepCopyInto(out *ImageRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRef.
func (in *ImageRef) DeepCopy() *ImageRef {
	if in == nil {
		return nil
	}
	out := new(ImageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubectlSpec) DeepCopyInto(out *KubectlSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubectlSpec.
func (in *KubectlSpec) DeepCopy() *KubectlSpec {
	if in == nil {
		return nil
	}
	out := new(KubectlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	in.ServiceMonitor.DeepCopyInto(&out.ServiceMonitor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementSpec) DeepCopyInto(out *PlacementSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementSpec.
func (in *PlacementSpec) DeepCopy() *PlacementSpec {
	if in == nil {
		return nil
	}
	out := new(PlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSpec) DeepCopyInto(out *ServiceAccountSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSpec.
func (in *ServiceAccountSpec) DeepCopy() *ServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorLabels) DeepCopyInto(out *ServiceMonitorLabels) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorLabels.
func (in *ServiceMonitorLabels) DeepCopy() *ServiceMonitorLabels {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeBundleVersion) DeepCopyInto(out *SpokeBundleVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeBundleVersion.
func (in *SpokeBundleVersion) DeepCopy() *SpokeBundleVersion {
	if in == nil {
		return nil
	}
	out := new(SpokeBundleVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeHub) DeepCopyInto(out *SpokeHub) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeHub.
func (in *SpokeHub) DeepCopy() *SpokeHub {
	if in == nil {
		return nil
	}
	out := new(SpokeHub)
	in.DeepCopyInto(out)
	return out
}
