// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for gateway.gloo.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for GatewayParameters

func (in *GatewayParameters) DeepCopyInto(out *GatewayParameters) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *GatewayParameters) DeepCopy() *GatewayParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayParameters)
	in.DeepCopyInto(out)
	return out
}

func (in *GatewayParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *GatewayParametersList) DeepCopyInto(out *GatewayParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatewayParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *GatewayParametersList) DeepCopy() *GatewayParametersList {
	if in == nil {
		return nil
	}
	out := new(GatewayParametersList)
	in.DeepCopyInto(out)
	return out
}

func (in *GatewayParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
