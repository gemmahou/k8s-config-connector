//go:build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuxiliaryVersionConfig) DeepCopyInto(out *AuxiliaryVersionConfig) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuxiliaryVersionConfig.
func (in *AuxiliaryVersionConfig) DeepCopy() *AuxiliaryVersionConfig {
	if in == nil {
		return nil
	}
	out := new(AuxiliaryVersionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendMetastore) DeepCopyInto(out *BackendMetastore) {
	*out = *in
	if in.ServiceRef != nil {
		in, out := &in.ServiceRef, &out.ServiceRef
		*out = new(ServiceRef)
		**out = **in
	}
	if in.MetastoreType != nil {
		in, out := &in.MetastoreType, &out.MetastoreType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendMetastore.
func (in *BackendMetastore) DeepCopy() *BackendMetastore {
	if in == nil {
		return nil
	}
	out := new(BackendMetastore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfig) DeepCopyInto(out *EncryptionConfig) {
	*out = *in
	if in.KMSKeyRef != nil {
		in, out := &in.KMSKeyRef, &out.KMSKeyRef
		*out = new(v1beta1.KMSCryptoKeyRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfig.
func (in *EncryptionConfig) DeepCopy() *EncryptionConfig {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationIdentity) DeepCopyInto(out *FederationIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(FederationParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationIdentity.
func (in *FederationIdentity) DeepCopy() *FederationIdentity {
	if in == nil {
		return nil
	}
	out := new(FederationIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationParent) DeepCopyInto(out *FederationParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationParent.
func (in *FederationParent) DeepCopy() *FederationParent {
	if in == nil {
		return nil
	}
	out := new(FederationParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationRef) DeepCopyInto(out *FederationRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationRef.
func (in *FederationRef) DeepCopy() *FederationRef {
	if in == nil {
		return nil
	}
	out := new(FederationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HiveMetastoreConfig) DeepCopyInto(out *HiveMetastoreConfig) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KerberosConfig != nil {
		in, out := &in.KerberosConfig, &out.KerberosConfig
		*out = new(KerberosConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EndpointProtocol != nil {
		in, out := &in.EndpointProtocol, &out.EndpointProtocol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HiveMetastoreConfig.
func (in *HiveMetastoreConfig) DeepCopy() *HiveMetastoreConfig {
	if in == nil {
		return nil
	}
	out := new(HiveMetastoreConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Int32Value) DeepCopyInto(out *Int32Value) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Int32Value.
func (in *Int32Value) DeepCopy() *Int32Value {
	if in == nil {
		return nil
	}
	out := new(Int32Value)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KerberosConfig) DeepCopyInto(out *KerberosConfig) {
	*out = *in
	if in.Keytab != nil {
		in, out := &in.Keytab, &out.Keytab
		*out = new(Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(string)
		**out = **in
	}
	if in.Krb5ConfigGCSURI != nil {
		in, out := &in.Krb5ConfigGCSURI, &out.Krb5ConfigGCSURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KerberosConfig.
func (in *KerberosConfig) DeepCopy() *KerberosConfig {
	if in == nil {
		return nil
	}
	out := new(KerberosConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindow) DeepCopyInto(out *MaintenanceWindow) {
	*out = *in
	if in.HourOfDay != nil {
		in, out := &in.HourOfDay, &out.HourOfDay
		*out = new(Int32Value)
		(*in).DeepCopyInto(*out)
	}
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindow.
func (in *MaintenanceWindow) DeepCopy() *MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataExport) DeepCopyInto(out *MetadataExport) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataExport.
func (in *MetadataExport) DeepCopy() *MetadataExport {
	if in == nil {
		return nil
	}
	out := new(MetadataExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataExportObservedState) DeepCopyInto(out *MetadataExportObservedState) {
	*out = *in
	if in.DestinationGCSURI != nil {
		in, out := &in.DestinationGCSURI, &out.DestinationGCSURI
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.DatabaseDumpType != nil {
		in, out := &in.DatabaseDumpType, &out.DatabaseDumpType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataExportObservedState.
func (in *MetadataExportObservedState) DeepCopy() *MetadataExportObservedState {
	if in == nil {
		return nil
	}
	out := new(MetadataExportObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataManagementActivity) DeepCopyInto(out *MetadataManagementActivity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataManagementActivity.
func (in *MetadataManagementActivity) DeepCopy() *MetadataManagementActivity {
	if in == nil {
		return nil
	}
	out := new(MetadataManagementActivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataManagementActivityObservedState) DeepCopyInto(out *MetadataManagementActivityObservedState) {
	*out = *in
	if in.MetadataExports != nil {
		in, out := &in.MetadataExports, &out.MetadataExports
		*out = make([]MetadataExportObservedState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restores != nil {
		in, out := &in.Restores, &out.Restores
		*out = make([]RestoreObservedState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataManagementActivityObservedState.
func (in *MetadataManagementActivityObservedState) DeepCopy() *MetadataManagementActivityObservedState {
	if in == nil {
		return nil
	}
	out := new(MetadataManagementActivityObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreFederation) DeepCopyInto(out *MetastoreFederation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreFederation.
func (in *MetastoreFederation) DeepCopy() *MetastoreFederation {
	if in == nil {
		return nil
	}
	out := new(MetastoreFederation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetastoreFederation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreFederationList) DeepCopyInto(out *MetastoreFederationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetastoreFederation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreFederationList.
func (in *MetastoreFederationList) DeepCopy() *MetastoreFederationList {
	if in == nil {
		return nil
	}
	out := new(MetastoreFederationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetastoreFederationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreFederationObservedState) DeepCopyInto(out *MetastoreFederationObservedState) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.EndpointURI != nil {
		in, out := &in.EndpointURI, &out.EndpointURI
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateMessage != nil {
		in, out := &in.StateMessage, &out.StateMessage
		*out = new(string)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreFederationObservedState.
func (in *MetastoreFederationObservedState) DeepCopy() *MetastoreFederationObservedState {
	if in == nil {
		return nil
	}
	out := new(MetastoreFederationObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreFederationSpec) DeepCopyInto(out *MetastoreFederationSpec) {
	*out = *in
	in.Parent.DeepCopyInto(&out.Parent)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.BackendMetastores != nil {
		in, out := &in.BackendMetastores, &out.BackendMetastores
		*out = make(map[string]BackendMetastore, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreFederationSpec.
func (in *MetastoreFederationSpec) DeepCopy() *MetastoreFederationSpec {
	if in == nil {
		return nil
	}
	out := new(MetastoreFederationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreFederationStatus) DeepCopyInto(out *MetastoreFederationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(MetastoreFederationObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreFederationStatus.
func (in *MetastoreFederationStatus) DeepCopy() *MetastoreFederationStatus {
	if in == nil {
		return nil
	}
	out := new(MetastoreFederationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreService) DeepCopyInto(out *MetastoreService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreService.
func (in *MetastoreService) DeepCopy() *MetastoreService {
	if in == nil {
		return nil
	}
	out := new(MetastoreService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetastoreService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreServiceList) DeepCopyInto(out *MetastoreServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetastoreService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreServiceList.
func (in *MetastoreServiceList) DeepCopy() *MetastoreServiceList {
	if in == nil {
		return nil
	}
	out := new(MetastoreServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetastoreServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreServiceObservedState) DeepCopyInto(out *MetastoreServiceObservedState) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.EndpointURI != nil {
		in, out := &in.EndpointURI, &out.EndpointURI
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateMessage != nil {
		in, out := &in.StateMessage, &out.StateMessage
		*out = new(string)
		**out = **in
	}
	if in.ArtifactGCSURI != nil {
		in, out := &in.ArtifactGCSURI, &out.ArtifactGCSURI
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
	if in.MetadataManagementActivity != nil {
		in, out := &in.MetadataManagementActivity, &out.MetadataManagementActivity
		*out = new(MetadataManagementActivityObservedState)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = new(NetworkConfigObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreServiceObservedState.
func (in *MetastoreServiceObservedState) DeepCopy() *MetastoreServiceObservedState {
	if in == nil {
		return nil
	}
	out := new(MetastoreServiceObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreServiceParent) DeepCopyInto(out *MetastoreServiceParent) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreServiceParent.
func (in *MetastoreServiceParent) DeepCopy() *MetastoreServiceParent {
	if in == nil {
		return nil
	}
	out := new(MetastoreServiceParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreServiceSpec) DeepCopyInto(out *MetastoreServiceSpec) {
	*out = *in
	in.MetastoreServiceParent.DeepCopyInto(&out.MetastoreServiceParent)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.HiveMetastoreConfig != nil {
		in, out := &in.HiveMetastoreConfig, &out.HiveMetastoreConfig
		*out = new(HiveMetastoreConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(v1beta1.ComputeNetworkRef)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindow)
		(*in).DeepCopyInto(*out)
	}
	if in.ReleaseChannel != nil {
		in, out := &in.ReleaseChannel, &out.ReleaseChannel
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(EncryptionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = new(NetworkConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseType != nil {
		in, out := &in.DatabaseType, &out.DatabaseType
		*out = new(string)
		**out = **in
	}
	if in.TelemetryConfig != nil {
		in, out := &in.TelemetryConfig, &out.TelemetryConfig
		*out = new(TelemetryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalingConfig != nil {
		in, out := &in.ScalingConfig, &out.ScalingConfig
		*out = new(ScalingConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreServiceSpec.
func (in *MetastoreServiceSpec) DeepCopy() *MetastoreServiceSpec {
	if in == nil {
		return nil
	}
	out := new(MetastoreServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreServiceStatus) DeepCopyInto(out *MetastoreServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(MetastoreServiceObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreServiceStatus.
func (in *MetastoreServiceStatus) DeepCopy() *MetastoreServiceStatus {
	if in == nil {
		return nil
	}
	out := new(MetastoreServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]NetworkConfig_Consumer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigObservedState) DeepCopyInto(out *NetworkConfigObservedState) {
	*out = *in
	if in.Consumers != nil {
		in, out := &in.Consumers, &out.Consumers
		*out = make([]NetworkConfig_ConsumerObservedState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigObservedState.
func (in *NetworkConfigObservedState) DeepCopy() *NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig_Consumer) DeepCopyInto(out *NetworkConfig_Consumer) {
	*out = *in
	if in.SubnetworkRef != nil {
		in, out := &in.SubnetworkRef, &out.SubnetworkRef
		*out = new(v1beta1.ComputeSubnetworkRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig_Consumer.
func (in *NetworkConfig_Consumer) DeepCopy() *NetworkConfig_Consumer {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig_Consumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig_ConsumerObservedState) DeepCopyInto(out *NetworkConfig_ConsumerObservedState) {
	*out = *in
	if in.EndpointURI != nil {
		in, out := &in.EndpointURI, &out.EndpointURI
		*out = new(string)
		**out = **in
	}
	if in.EndpointLocation != nil {
		in, out := &in.EndpointLocation, &out.EndpointLocation
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig_ConsumerObservedState.
func (in *NetworkConfig_ConsumerObservedState) DeepCopy() *NetworkConfig_ConsumerObservedState {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig_ConsumerObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restore) DeepCopyInto(out *Restore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restore.
func (in *Restore) DeepCopy() *Restore {
	if in == nil {
		return nil
	}
	out := new(Restore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreObservedState) DeepCopyInto(out *RestoreObservedState) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreObservedState.
func (in *RestoreObservedState) DeepCopy() *RestoreObservedState {
	if in == nil {
		return nil
	}
	out := new(RestoreObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingConfig) DeepCopyInto(out *ScalingConfig) {
	*out = *in
	if in.InstanceSize != nil {
		in, out := &in.InstanceSize, &out.InstanceSize
		*out = new(string)
		**out = **in
	}
	if in.ScalingFactor != nil {
		in, out := &in.ScalingFactor, &out.ScalingFactor
		*out = new(float32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingConfig.
func (in *ScalingConfig) DeepCopy() *ScalingConfig {
	if in == nil {
		return nil
	}
	out := new(ScalingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(secretmanagerv1beta1.SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceIdentity) DeepCopyInto(out *ServiceIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(ServiceParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceIdentity.
func (in *ServiceIdentity) DeepCopy() *ServiceIdentity {
	if in == nil {
		return nil
	}
	out := new(ServiceIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceParent) DeepCopyInto(out *ServiceParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceParent.
func (in *ServiceParent) DeepCopy() *ServiceParent {
	if in == nil {
		return nil
	}
	out := new(ServiceParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRef) DeepCopyInto(out *ServiceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRef.
func (in *ServiceRef) DeepCopy() *ServiceRef {
	if in == nil {
		return nil
	}
	out := new(ServiceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetryConfig) DeepCopyInto(out *TelemetryConfig) {
	*out = *in
	if in.LogFormat != nil {
		in, out := &in.LogFormat, &out.LogFormat
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetryConfig.
func (in *TelemetryConfig) DeepCopy() *TelemetryConfig {
	if in == nil {
		return nil
	}
	out := new(TelemetryConfig)
	in.DeepCopyInto(out)
	return out
}
