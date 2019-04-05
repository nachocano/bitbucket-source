// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitBucketConsumerValue) DeepCopyInto(out *BitBucketConsumerValue) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitBucketConsumerValue.
func (in *BitBucketConsumerValue) DeepCopy() *BitBucketConsumerValue {
	if in == nil {
		return nil
	}
	out := new(BitBucketConsumerValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitBucketSource) DeepCopyInto(out *BitBucketSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitBucketSource.
func (in *BitBucketSource) DeepCopy() *BitBucketSource {
	if in == nil {
		return nil
	}
	out := new(BitBucketSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BitBucketSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitBucketSourceList) DeepCopyInto(out *BitBucketSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BitBucketSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitBucketSourceList.
func (in *BitBucketSourceList) DeepCopy() *BitBucketSourceList {
	if in == nil {
		return nil
	}
	out := new(BitBucketSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BitBucketSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitBucketSourceSpec) DeepCopyInto(out *BitBucketSourceSpec) {
	*out = *in
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ConsumerKey.DeepCopyInto(&out.ConsumerKey)
	in.ConsumerSecret.DeepCopyInto(&out.ConsumerSecret)
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitBucketSourceSpec.
func (in *BitBucketSourceSpec) DeepCopy() *BitBucketSourceSpec {
	if in == nil {
		return nil
	}
	out := new(BitBucketSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BitBucketSourceStatus) DeepCopyInto(out *BitBucketSourceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitBucketSourceStatus.
func (in *BitBucketSourceStatus) DeepCopy() *BitBucketSourceStatus {
	if in == nil {
		return nil
	}
	out := new(BitBucketSourceStatus)
	in.DeepCopyInto(out)
	return out
}