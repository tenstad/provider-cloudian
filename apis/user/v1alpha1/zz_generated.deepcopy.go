//go:build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKey) DeepCopyInto(out *AccessKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKey.
func (in *AccessKey) DeepCopy() *AccessKey {
	if in == nil {
		return nil
	}
	out := new(AccessKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeyList) DeepCopyInto(out *AccessKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeyList.
func (in *AccessKeyList) DeepCopy() *AccessKeyList {
	if in == nil {
		return nil
	}
	out := new(AccessKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeyObservation) DeepCopyInto(out *AccessKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeyObservation.
func (in *AccessKeyObservation) DeepCopy() *AccessKeyObservation {
	if in == nil {
		return nil
	}
	out := new(AccessKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeyParameters) DeepCopyInto(out *AccessKeyParameters) {
	*out = *in
	if in.UserIDRef != nil {
		in, out := &in.UserIDRef, &out.UserIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.UserIDSelector != nil {
		in, out := &in.UserIDSelector, &out.UserIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeyParameters.
func (in *AccessKeyParameters) DeepCopy() *AccessKeyParameters {
	if in == nil {
		return nil
	}
	out := new(AccessKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeySpec) DeepCopyInto(out *AccessKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeySpec.
func (in *AccessKeySpec) DeepCopy() *AccessKeySpec {
	if in == nil {
		return nil
	}
	out := new(AccessKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessKeyStatus) DeepCopyInto(out *AccessKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessKeyStatus.
func (in *AccessKeyStatus) DeepCopy() *AccessKeyStatus {
	if in == nil {
		return nil
	}
	out := new(AccessKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Group) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupList) DeepCopyInto(out *GroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Group, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupList.
func (in *GroupList) DeepCopy() *GroupList {
	if in == nil {
		return nil
	}
	out := new(GroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupObservation) DeepCopyInto(out *GroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupObservation.
func (in *GroupObservation) DeepCopy() *GroupObservation {
	if in == nil {
		return nil
	}
	out := new(GroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupParameters) DeepCopyInto(out *GroupParameters) {
	*out = *in
	if in.LDAPEnabled != nil {
		in, out := &in.LDAPEnabled, &out.LDAPEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LDAPGroup != nil {
		in, out := &in.LDAPGroup, &out.LDAPGroup
		*out = new(string)
		**out = **in
	}
	if in.LDAPMatchAttribute != nil {
		in, out := &in.LDAPMatchAttribute, &out.LDAPMatchAttribute
		*out = new(string)
		**out = **in
	}
	if in.LDAPSearch != nil {
		in, out := &in.LDAPSearch, &out.LDAPSearch
		*out = new(string)
		**out = **in
	}
	if in.LDAPSearchUserBase != nil {
		in, out := &in.LDAPSearchUserBase, &out.LDAPSearchUserBase
		*out = new(string)
		**out = **in
	}
	if in.LDAPServerURL != nil {
		in, out := &in.LDAPServerURL, &out.LDAPServerURL
		*out = new(string)
		**out = **in
	}
	if in.LDAPUserDNTemplate != nil {
		in, out := &in.LDAPUserDNTemplate, &out.LDAPUserDNTemplate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupParameters.
func (in *GroupParameters) DeepCopy() *GroupParameters {
	if in == nil {
		return nil
	}
	out := new(GroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupQualityOfServiceLimits) DeepCopyInto(out *GroupQualityOfServiceLimits) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupQualityOfServiceLimits.
func (in *GroupQualityOfServiceLimits) DeepCopy() *GroupQualityOfServiceLimits {
	if in == nil {
		return nil
	}
	out := new(GroupQualityOfServiceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupQualityOfServiceLimits) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupQualityOfServiceLimitsList) DeepCopyInto(out *GroupQualityOfServiceLimitsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupQualityOfServiceLimits, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupQualityOfServiceLimitsList.
func (in *GroupQualityOfServiceLimitsList) DeepCopy() *GroupQualityOfServiceLimitsList {
	if in == nil {
		return nil
	}
	out := new(GroupQualityOfServiceLimitsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupQualityOfServiceLimitsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupQualityOfServiceLimitsObservation) DeepCopyInto(out *GroupQualityOfServiceLimitsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupQualityOfServiceLimitsObservation.
func (in *GroupQualityOfServiceLimitsObservation) DeepCopy() *GroupQualityOfServiceLimitsObservation {
	if in == nil {
		return nil
	}
	out := new(GroupQualityOfServiceLimitsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupQualityOfServiceLimitsParameters) DeepCopyInto(out *GroupQualityOfServiceLimitsParameters) {
	*out = *in
	if in.GroupIDRef != nil {
		in, out := &in.GroupIDRef, &out.GroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GroupIDSelector != nil {
		in, out := &in.GroupIDSelector, &out.GroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(QualityOfServiceLimits)
		(*in).DeepCopyInto(*out)
	}
	if in.Hard != nil {
		in, out := &in.Hard, &out.Hard
		*out = new(QualityOfServiceLimits)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupQualityOfServiceLimitsParameters.
func (in *GroupQualityOfServiceLimitsParameters) DeepCopy() *GroupQualityOfServiceLimitsParameters {
	if in == nil {
		return nil
	}
	out := new(GroupQualityOfServiceLimitsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupQualityOfServiceLimitsSpec) DeepCopyInto(out *GroupQualityOfServiceLimitsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupQualityOfServiceLimitsSpec.
func (in *GroupQualityOfServiceLimitsSpec) DeepCopy() *GroupQualityOfServiceLimitsSpec {
	if in == nil {
		return nil
	}
	out := new(GroupQualityOfServiceLimitsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupQualityOfServiceLimitsStatus) DeepCopyInto(out *GroupQualityOfServiceLimitsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupQualityOfServiceLimitsStatus.
func (in *GroupQualityOfServiceLimitsStatus) DeepCopy() *GroupQualityOfServiceLimitsStatus {
	if in == nil {
		return nil
	}
	out := new(GroupQualityOfServiceLimitsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSpec) DeepCopyInto(out *GroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSpec.
func (in *GroupSpec) DeepCopy() *GroupSpec {
	if in == nil {
		return nil
	}
	out := new(GroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupStatus) DeepCopyInto(out *GroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatus.
func (in *GroupStatus) DeepCopy() *GroupStatus {
	if in == nil {
		return nil
	}
	out := new(GroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QualityOfServiceLimits) DeepCopyInto(out *QualityOfServiceLimits) {
	*out = *in
	if in.StorageQuotaBytes != nil {
		in, out := &in.StorageQuotaBytes, &out.StorageQuotaBytes
		*out = new(Quantity)
		**out = **in
	}
	if in.StorageQuotaCount != nil {
		in, out := &in.StorageQuotaCount, &out.StorageQuotaCount
		*out = new(uint32)
		**out = **in
	}
	if in.RequestsPerMin != nil {
		in, out := &in.RequestsPerMin, &out.RequestsPerMin
		*out = new(uint32)
		**out = **in
	}
	if in.InboundBytesPerMin != nil {
		in, out := &in.InboundBytesPerMin, &out.InboundBytesPerMin
		*out = new(Quantity)
		**out = **in
	}
	if in.OutboundBytesPerMin != nil {
		in, out := &in.OutboundBytesPerMin, &out.OutboundBytesPerMin
		*out = new(Quantity)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QualityOfServiceLimits.
func (in *QualityOfServiceLimits) DeepCopy() *QualityOfServiceLimits {
	if in == nil {
		return nil
	}
	out := new(QualityOfServiceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserObservation) DeepCopyInto(out *UserObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserObservation.
func (in *UserObservation) DeepCopy() *UserObservation {
	if in == nil {
		return nil
	}
	out := new(UserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserParameters) DeepCopyInto(out *UserParameters) {
	*out = *in
	if in.GroupIDRef != nil {
		in, out := &in.GroupIDRef, &out.GroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GroupIDSelector != nil {
		in, out := &in.GroupIDSelector, &out.GroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserParameters.
func (in *UserParameters) DeepCopy() *UserParameters {
	if in == nil {
		return nil
	}
	out := new(UserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserQualityOfServiceLimits) DeepCopyInto(out *UserQualityOfServiceLimits) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserQualityOfServiceLimits.
func (in *UserQualityOfServiceLimits) DeepCopy() *UserQualityOfServiceLimits {
	if in == nil {
		return nil
	}
	out := new(UserQualityOfServiceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserQualityOfServiceLimits) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserQualityOfServiceLimitsList) DeepCopyInto(out *UserQualityOfServiceLimitsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserQualityOfServiceLimits, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserQualityOfServiceLimitsList.
func (in *UserQualityOfServiceLimitsList) DeepCopy() *UserQualityOfServiceLimitsList {
	if in == nil {
		return nil
	}
	out := new(UserQualityOfServiceLimitsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserQualityOfServiceLimitsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserQualityOfServiceLimitsObservation) DeepCopyInto(out *UserQualityOfServiceLimitsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserQualityOfServiceLimitsObservation.
func (in *UserQualityOfServiceLimitsObservation) DeepCopy() *UserQualityOfServiceLimitsObservation {
	if in == nil {
		return nil
	}
	out := new(UserQualityOfServiceLimitsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserQualityOfServiceLimitsParameters) DeepCopyInto(out *UserQualityOfServiceLimitsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserQualityOfServiceLimitsParameters.
func (in *UserQualityOfServiceLimitsParameters) DeepCopy() *UserQualityOfServiceLimitsParameters {
	if in == nil {
		return nil
	}
	out := new(UserQualityOfServiceLimitsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserQualityOfServiceLimitsSpec) DeepCopyInto(out *UserQualityOfServiceLimitsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserQualityOfServiceLimitsSpec.
func (in *UserQualityOfServiceLimitsSpec) DeepCopy() *UserQualityOfServiceLimitsSpec {
	if in == nil {
		return nil
	}
	out := new(UserQualityOfServiceLimitsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserQualityOfServiceLimitsStatus) DeepCopyInto(out *UserQualityOfServiceLimitsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserQualityOfServiceLimitsStatus.
func (in *UserQualityOfServiceLimitsStatus) DeepCopy() *UserQualityOfServiceLimitsStatus {
	if in == nil {
		return nil
	}
	out := new(UserQualityOfServiceLimitsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}
