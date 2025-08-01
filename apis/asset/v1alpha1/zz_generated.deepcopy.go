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
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetFeed) DeepCopyInto(out *AssetFeed) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetFeed.
func (in *AssetFeed) DeepCopy() *AssetFeed {
	if in == nil {
		return nil
	}
	out := new(AssetFeed)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssetFeed) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetFeedList) DeepCopyInto(out *AssetFeedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AssetFeed, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetFeedList.
func (in *AssetFeedList) DeepCopy() *AssetFeedList {
	if in == nil {
		return nil
	}
	out := new(AssetFeedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssetFeedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetFeedParent) DeepCopyInto(out *AssetFeedParent) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1beta1.OrganizationRef)
		**out = **in
	}
	if in.FolderRef != nil {
		in, out := &in.FolderRef, &out.FolderRef
		*out = new(v1beta1.FolderRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetFeedParent.
func (in *AssetFeedParent) DeepCopy() *AssetFeedParent {
	if in == nil {
		return nil
	}
	out := new(AssetFeedParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetFeedSpec) DeepCopyInto(out *AssetFeedSpec) {
	*out = *in
	in.Parent.DeepCopyInto(&out.Parent)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.AssetNames != nil {
		in, out := &in.AssetNames, &out.AssetNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AssetTypes != nil {
		in, out := &in.AssetTypes, &out.AssetTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.FeedOutputConfig != nil {
		in, out := &in.FeedOutputConfig, &out.FeedOutputConfig
		*out = new(FeedOutputConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(Expr)
		(*in).DeepCopyInto(*out)
	}
	if in.RelationshipTypes != nil {
		in, out := &in.RelationshipTypes, &out.RelationshipTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetFeedSpec.
func (in *AssetFeedSpec) DeepCopy() *AssetFeedSpec {
	if in == nil {
		return nil
	}
	out := new(AssetFeedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetFeedStatus) DeepCopyInto(out *AssetFeedStatus) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetFeedStatus.
func (in *AssetFeedStatus) DeepCopy() *AssetFeedStatus {
	if in == nil {
		return nil
	}
	out := new(AssetFeedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetSavedQuery) DeepCopyInto(out *AssetSavedQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetSavedQuery.
func (in *AssetSavedQuery) DeepCopy() *AssetSavedQuery {
	if in == nil {
		return nil
	}
	out := new(AssetSavedQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssetSavedQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetSavedQueryList) DeepCopyInto(out *AssetSavedQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AssetSavedQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetSavedQueryList.
func (in *AssetSavedQueryList) DeepCopy() *AssetSavedQueryList {
	if in == nil {
		return nil
	}
	out := new(AssetSavedQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssetSavedQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetSavedQueryObservedState) DeepCopyInto(out *AssetSavedQueryObservedState) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Creator != nil {
		in, out := &in.Creator, &out.Creator
		*out = new(string)
		**out = **in
	}
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = new(string)
		**out = **in
	}
	if in.LastUpdater != nil {
		in, out := &in.LastUpdater, &out.LastUpdater
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetSavedQueryObservedState.
func (in *AssetSavedQueryObservedState) DeepCopy() *AssetSavedQueryObservedState {
	if in == nil {
		return nil
	}
	out := new(AssetSavedQueryObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetSavedQuerySpec) DeepCopyInto(out *AssetSavedQuerySpec) {
	*out = *in
	in.Parent.DeepCopyInto(&out.Parent)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(SavedQuery_QueryContent)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetSavedQuerySpec.
func (in *AssetSavedQuerySpec) DeepCopy() *AssetSavedQuerySpec {
	if in == nil {
		return nil
	}
	out := new(AssetSavedQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetSavedQueryStatus) DeepCopyInto(out *AssetSavedQueryStatus) {
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
		*out = new(AssetSavedQueryObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetSavedQueryStatus.
func (in *AssetSavedQueryStatus) DeepCopy() *AssetSavedQueryStatus {
	if in == nil {
		return nil
	}
	out := new(AssetSavedQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Expr) DeepCopyInto(out *Expr) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Expr.
func (in *Expr) DeepCopy() *Expr {
	if in == nil {
		return nil
	}
	out := new(Expr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeedIdentity) DeepCopyInto(out *FeedIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(FeedParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeedIdentity.
func (in *FeedIdentity) DeepCopy() *FeedIdentity {
	if in == nil {
		return nil
	}
	out := new(FeedIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeedOutputConfig) DeepCopyInto(out *FeedOutputConfig) {
	*out = *in
	if in.PubsubDestination != nil {
		in, out := &in.PubsubDestination, &out.PubsubDestination
		*out = new(PubsubDestination)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeedOutputConfig.
func (in *FeedOutputConfig) DeepCopy() *FeedOutputConfig {
	if in == nil {
		return nil
	}
	out := new(FeedOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeedParent) DeepCopyInto(out *FeedParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeedParent.
func (in *FeedParent) DeepCopy() *FeedParent {
	if in == nil {
		return nil
	}
	out := new(FeedParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeedRef) DeepCopyInto(out *FeedRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeedRef.
func (in *FeedRef) DeepCopy() *FeedRef {
	if in == nil {
		return nil
	}
	out := new(FeedRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyAnalysisQuery) DeepCopyInto(out *IAMPolicyAnalysisQuery) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.ResourceSelector != nil {
		in, out := &in.ResourceSelector, &out.ResourceSelector
		*out = new(IAMPolicyAnalysisQuery_ResourceSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IdentitySelector != nil {
		in, out := &in.IdentitySelector, &out.IdentitySelector
		*out = new(IAMPolicyAnalysisQuery_IdentitySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessSelector != nil {
		in, out := &in.AccessSelector, &out.AccessSelector
		*out = new(IAMPolicyAnalysisQuery_AccessSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(IAMPolicyAnalysisQuery_Options)
		(*in).DeepCopyInto(*out)
	}
	if in.ConditionContext != nil {
		in, out := &in.ConditionContext, &out.ConditionContext
		*out = new(IAMPolicyAnalysisQuery_ConditionContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyAnalysisQuery.
func (in *IAMPolicyAnalysisQuery) DeepCopy() *IAMPolicyAnalysisQuery {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyAnalysisQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyAnalysisQuery_AccessSelector) DeepCopyInto(out *IAMPolicyAnalysisQuery_AccessSelector) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyAnalysisQuery_AccessSelector.
func (in *IAMPolicyAnalysisQuery_AccessSelector) DeepCopy() *IAMPolicyAnalysisQuery_AccessSelector {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyAnalysisQuery_AccessSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyAnalysisQuery_ConditionContext) DeepCopyInto(out *IAMPolicyAnalysisQuery_ConditionContext) {
	*out = *in
	if in.AccessTime != nil {
		in, out := &in.AccessTime, &out.AccessTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyAnalysisQuery_ConditionContext.
func (in *IAMPolicyAnalysisQuery_ConditionContext) DeepCopy() *IAMPolicyAnalysisQuery_ConditionContext {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyAnalysisQuery_ConditionContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyAnalysisQuery_IdentitySelector) DeepCopyInto(out *IAMPolicyAnalysisQuery_IdentitySelector) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyAnalysisQuery_IdentitySelector.
func (in *IAMPolicyAnalysisQuery_IdentitySelector) DeepCopy() *IAMPolicyAnalysisQuery_IdentitySelector {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyAnalysisQuery_IdentitySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyAnalysisQuery_Options) DeepCopyInto(out *IAMPolicyAnalysisQuery_Options) {
	*out = *in
	if in.ExpandGroups != nil {
		in, out := &in.ExpandGroups, &out.ExpandGroups
		*out = new(bool)
		**out = **in
	}
	if in.ExpandRoles != nil {
		in, out := &in.ExpandRoles, &out.ExpandRoles
		*out = new(bool)
		**out = **in
	}
	if in.ExpandResources != nil {
		in, out := &in.ExpandResources, &out.ExpandResources
		*out = new(bool)
		**out = **in
	}
	if in.OutputResourceEdges != nil {
		in, out := &in.OutputResourceEdges, &out.OutputResourceEdges
		*out = new(bool)
		**out = **in
	}
	if in.OutputGroupEdges != nil {
		in, out := &in.OutputGroupEdges, &out.OutputGroupEdges
		*out = new(bool)
		**out = **in
	}
	if in.AnalyzeServiceAccountImpersonation != nil {
		in, out := &in.AnalyzeServiceAccountImpersonation, &out.AnalyzeServiceAccountImpersonation
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyAnalysisQuery_Options.
func (in *IAMPolicyAnalysisQuery_Options) DeepCopy() *IAMPolicyAnalysisQuery_Options {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyAnalysisQuery_Options)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyAnalysisQuery_ResourceSelector) DeepCopyInto(out *IAMPolicyAnalysisQuery_ResourceSelector) {
	*out = *in
	if in.FullResourceName != nil {
		in, out := &in.FullResourceName, &out.FullResourceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyAnalysisQuery_ResourceSelector.
func (in *IAMPolicyAnalysisQuery_ResourceSelector) DeepCopy() *IAMPolicyAnalysisQuery_ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyAnalysisQuery_ResourceSelector)
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
	if in.FolderRef != nil {
		in, out := &in.FolderRef, &out.FolderRef
		*out = new(v1beta1.FolderRef)
		**out = **in
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1beta1.OrganizationRef)
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
func (in *PubsubDestination) DeepCopyInto(out *PubsubDestination) {
	*out = *in
	if in.TopicRef != nil {
		in, out := &in.TopicRef, &out.TopicRef
		*out = new(pubsubv1beta1.PubSubTopicRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubsubDestination.
func (in *PubsubDestination) DeepCopy() *PubsubDestination {
	if in == nil {
		return nil
	}
	out := new(PubsubDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SavedQueryIdentity) DeepCopyInto(out *SavedQueryIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(SavedQueryParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SavedQueryIdentity.
func (in *SavedQueryIdentity) DeepCopy() *SavedQueryIdentity {
	if in == nil {
		return nil
	}
	out := new(SavedQueryIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SavedQueryParent) DeepCopyInto(out *SavedQueryParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SavedQueryParent.
func (in *SavedQueryParent) DeepCopy() *SavedQueryParent {
	if in == nil {
		return nil
	}
	out := new(SavedQueryParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SavedQueryRef) DeepCopyInto(out *SavedQueryRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SavedQueryRef.
func (in *SavedQueryRef) DeepCopy() *SavedQueryRef {
	if in == nil {
		return nil
	}
	out := new(SavedQueryRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SavedQuery_QueryContent) DeepCopyInto(out *SavedQuery_QueryContent) {
	*out = *in
	if in.IAMPolicyAnalysisQuery != nil {
		in, out := &in.IAMPolicyAnalysisQuery, &out.IAMPolicyAnalysisQuery
		*out = new(IAMPolicyAnalysisQuery)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SavedQuery_QueryContent.
func (in *SavedQuery_QueryContent) DeepCopy() *SavedQuery_QueryContent {
	if in == nil {
		return nil
	}
	out := new(SavedQuery_QueryContent)
	in.DeepCopyInto(out)
	return out
}
