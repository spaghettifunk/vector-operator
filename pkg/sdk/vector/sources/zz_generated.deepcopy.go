//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package sources

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApacheMetricsSpec) DeepCopyInto(out *ApacheMetricsSpec) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(Proxy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApacheMetricsSpec.
func (in *ApacheMetricsSpec) DeepCopy() *ApacheMetricsSpec {
	if in == nil {
		return nil
	}
	out := new(ApacheMetricsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DemoLogsSpec) DeepCopyInto(out *DemoLogsSpec) {
	*out = *in
	if in.Decoding != nil {
		in, out := &in.Decoding, &out.Decoding
		*out = new(Decoding)
		**out = **in
	}
	if in.Framing != nil {
		in, out := &in.Framing, &out.Framing
		*out = new(Framing)
		(*in).DeepCopyInto(*out)
	}
	if in.Lines != nil {
		in, out := &in.Lines, &out.Lines
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DemoLogsSpec.
func (in *DemoLogsSpec) DeepCopy() *DemoLogsSpec {
	if in == nil {
		return nil
	}
	out := new(DemoLogsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Framing) DeepCopyInto(out *Framing) {
	*out = *in
	if in.CharacterDelimited != nil {
		in, out := &in.CharacterDelimited, &out.CharacterDelimited
		*out = new(CharacterDelimited)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Framing.
func (in *Framing) DeepCopy() *Framing {
	if in == nil {
		return nil
	}
	out := new(Framing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Proxy) DeepCopyInto(out *Proxy) {
	*out = *in
	if in.NoProxy != nil {
		in, out := &in.NoProxy, &out.NoProxy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Proxy.
func (in *Proxy) DeepCopy() *Proxy {
	if in == nil {
		return nil
	}
	out := new(Proxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorSpec) DeepCopyInto(out *VectorSpec) {
	*out = *in
	if in.KeepAlive != nil {
		in, out := &in.KeepAlive, &out.KeepAlive
		*out = new(KeepAlive)
		**out = **in
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(TLSConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorSpec.
func (in *VectorSpec) DeepCopy() *VectorSpec {
	if in == nil {
		return nil
	}
	out := new(VectorSpec)
	in.DeepCopyInto(out)
	return out
}
