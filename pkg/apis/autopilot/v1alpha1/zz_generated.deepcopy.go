// +build !ignore_autogenerated

/*
Copyright 2019 Openstorage.org

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAction) DeepCopyInto(out *PolicyAction) {
	*out = *in
	in.ActionObject.DeepCopyInto(&out.ActionObject)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAction.
func (in *PolicyAction) DeepCopy() *PolicyAction {
	if in == nil {
		return nil
	}
	out := new(PolicyAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyObject) DeepCopyInto(out *PolicyObject) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyObject.
func (in *PolicyObject) DeepCopy() *PolicyObject {
	if in == nil {
		return nil
	}
	out := new(PolicyObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicy) DeepCopyInto(out *StoragePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicy.
func (in *StoragePolicy) DeepCopy() *StoragePolicy {
	if in == nil {
		return nil
	}
	out := new(StoragePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyList) DeepCopyInto(out *StoragePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyList.
func (in *StoragePolicyList) DeepCopy() *StoragePolicyList {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicySpec) DeepCopyInto(out *StoragePolicySpec) {
	*out = *in
	in.Object.DeepCopyInto(&out.Object)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*v1.LabelSelectorRequirement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.LabelSelectorRequirement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.Action.DeepCopyInto(&out.Action)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicySpec.
func (in *StoragePolicySpec) DeepCopy() *StoragePolicySpec {
	if in == nil {
		return nil
	}
	out := new(StoragePolicySpec)
	in.DeepCopyInto(out)
	return out
}