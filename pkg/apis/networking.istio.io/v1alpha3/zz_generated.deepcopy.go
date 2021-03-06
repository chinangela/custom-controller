// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha3

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Destination).DeepCopyInto(out.(*Destination))
			return nil
		}, InType: reflect.TypeOf(&Destination{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DestinationWeight).DeepCopyInto(out.(*DestinationWeight))
			return nil
		}, InType: reflect.TypeOf(&DestinationWeight{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HTTPRoute).DeepCopyInto(out.(*HTTPRoute))
			return nil
		}, InType: reflect.TypeOf(&HTTPRoute{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MatchRequest).DeepCopyInto(out.(*MatchRequest))
			return nil
		}, InType: reflect.TypeOf(&MatchRequest{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PortSelector).DeepCopyInto(out.(*PortSelector))
			return nil
		}, InType: reflect.TypeOf(&PortSelector{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PortSelector_Name).DeepCopyInto(out.(*PortSelector_Name))
			return nil
		}, InType: reflect.TypeOf(&PortSelector_Name{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VirtualService).DeepCopyInto(out.(*VirtualService))
			return nil
		}, InType: reflect.TypeOf(&VirtualService{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VirtualServiceSpec).DeepCopyInto(out.(*VirtualServiceSpec))
			return nil
		}, InType: reflect.TypeOf(&VirtualServiceSpec{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	out.Port = in.Port
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Destination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationWeight) DeepCopyInto(out *DestinationWeight) {
	*out = *in
	out.Destination = in.Destination
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationWeight.
func (in *DestinationWeight) DeepCopy() *DestinationWeight {
	if in == nil {
		return nil
	}
	out := new(DestinationWeight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DestinationWeight) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRoute) DeepCopyInto(out *HTTPRoute) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]MatchRequest, len(*in))
		copy(*out, *in)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = make([]DestinationWeight, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRoute.
func (in *HTTPRoute) DeepCopy() *HTTPRoute {
	if in == nil {
		return nil
	}
	out := new(HTTPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchRequest) DeepCopyInto(out *MatchRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchRequest.
func (in *MatchRequest) DeepCopy() *MatchRequest {
	if in == nil {
		return nil
	}
	out := new(MatchRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MatchRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	*out = *in
	out.Port = in.Port
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortSelector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSelector_Name) DeepCopyInto(out *PortSelector_Name) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector_Name.
func (in *PortSelector_Name) DeepCopy() *PortSelector_Name {
	if in == nil {
		return nil
	}
	out := new(PortSelector_Name)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortSelector_Name) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualService) DeepCopyInto(out *VirtualService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualService.
func (in *VirtualService) DeepCopy() *VirtualService {
	if in == nil {
		return nil
	}
	out := new(VirtualService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceSpec) DeepCopyInto(out *VirtualServiceSpec) {
	*out = *in
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]Gateway, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]Host, len(*in))
		copy(*out, *in)
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]HTTPRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceSpec.
func (in *VirtualServiceSpec) DeepCopy() *VirtualServiceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServiceSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
