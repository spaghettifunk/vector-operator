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

package v1alpha1

import (
	"github.com/banzaicloud/operator-tools/pkg/typeoverride"
	"github.com/banzaicloud/operator-tools/pkg/volume"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/spaghettifunk/vector-operator/pkg/sdk/vector/sources"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = new(PipelineSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DaemonSetAnnotations != nil {
		in, out := &in.DaemonSetAnnotations, &out.DaemonSetAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Image.DeepCopyInto(&out.Image)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(Security)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]*VolumeMount, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VolumeMount)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountOverrides != nil {
		in, out := &in.ServiceAccountOverrides, &out.ServiceAccountOverrides
		*out = new(typeoverride.ServiceAccount)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec.
func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatorScaling) DeepCopyInto(out *AggregatorScaling) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatorScaling.
func (in *AggregatorScaling) DeepCopy() *AggregatorScaling {
	if in == nil {
		return nil
	}
	out := new(AggregatorScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatorSpec) DeepCopyInto(out *AggregatorSpec) {
	*out = *in
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = new(PipelineSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StatefulSetAnnotations != nil {
		in, out := &in.StatefulSetAnnotations, &out.StatefulSetAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigCheckAnnotations != nil {
		in, out := &in.ConfigCheckAnnotations, &out.ConfigCheckAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Image.DeepCopyInto(&out.Image)
	in.BufferStorageVolume.DeepCopyInto(&out.BufferStorageVolume)
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]ExtraVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.ConfigCheckResources.DeepCopyInto(&out.ConfigCheckResources)
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	out.ReadinessDefaultCheck = in.ReadinessDefaultCheck
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(Security)
		(*in).DeepCopyInto(*out)
	}
	if in.Scaling != nil {
		in, out := &in.Scaling, &out.Scaling
		*out = new(AggregatorScaling)
		**out = **in
	}
	if in.ServiceAccountOverrides != nil {
		in, out := &in.ServiceAccountOverrides, &out.ServiceAccountOverrides
		*out = new(typeoverride.ServiceAccount)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatorSpec.
func (in *AggregatorSpec) DeepCopy() *AggregatorSpec {
	if in == nil {
		return nil
	}
	out := new(AggregatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnrichmentTables) DeepCopyInto(out *EnrichmentTables) {
	*out = *in
	in.File.DeepCopyInto(&out.File)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnrichmentTables.
func (in *EnrichmentTables) DeepCopy() *EnrichmentTables {
	if in == nil {
		return nil
	}
	out := new(EnrichmentTables)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraVolume) DeepCopyInto(out *ExtraVolume) {
	*out = *in
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(volume.KubernetesVolume)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraVolume.
func (in *ExtraVolume) DeepCopy() *ExtraVolume {
	if in == nil {
		return nil
	}
	out := new(ExtraVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalOptionsSpec) DeepCopyInto(out *GlobalOptionsSpec) {
	*out = *in
	if in.EnrichmentTables != nil {
		in, out := &in.EnrichmentTables, &out.EnrichmentTables
		*out = new(EnrichmentTables)
		(*in).DeepCopyInto(*out)
	}
	if in.Healthcheck != nil {
		in, out := &in.Healthcheck, &out.Healthcheck
		*out = new(Healthcheck)
		**out = **in
	}
	if in.LogSchema != nil {
		in, out := &in.LogSchema, &out.LogSchema
		*out = new(LogSchema)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(Proxy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalOptionsSpec.
func (in *GlobalOptionsSpec) DeepCopy() *GlobalOptionsSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalOptionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Healthcheck) DeepCopyInto(out *Healthcheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Healthcheck.
func (in *Healthcheck) DeepCopy() *Healthcheck {
	if in == nil {
		return nil
	}
	out := new(Healthcheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSchema) DeepCopyInto(out *LogSchema) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSchema.
func (in *LogSchema) DeepCopy() *LogSchema {
	if in == nil {
		return nil
	}
	out := new(LogSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metrics) DeepCopyInto(out *Metrics) {
	*out = *in
	in.ServiceMonitorConfig.DeepCopyInto(&out.ServiceMonitorConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metrics.
func (in *Metrics) DeepCopy() *Metrics {
	if in == nil {
		return nil
	}
	out := new(Metrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]*Source, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Source)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Transforms != nil {
		in, out := &in.Transforms, &out.Transforms
		*out = make([]*Transform, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Transform)
				**out = **in
			}
		}
	}
	if in.Sinks != nil {
		in, out := &in.Sinks, &out.Sinks
		*out = make([]*Sink, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Sink)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.Problems != nil {
		in, out := &in.Problems, &out.Problems
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
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
func (in *ReadinessDefaultCheck) DeepCopyInto(out *ReadinessDefaultCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessDefaultCheck.
func (in *ReadinessDefaultCheck) DeepCopy() *ReadinessDefaultCheck {
	if in == nil {
		return nil
	}
	out := new(ReadinessDefaultCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Security) DeepCopyInto(out *Security) {
	*out = *in
	if in.RoleBasedAccessControlCreate != nil {
		in, out := &in.RoleBasedAccessControlCreate, &out.RoleBasedAccessControlCreate
		*out = new(bool)
		**out = **in
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Security.
func (in *Security) DeepCopy() *Security {
	if in == nil {
		return nil
	}
	out := new(Security)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorConfig) DeepCopyInto(out *ServiceMonitorConfig) {
	*out = *in
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Relabelings != nil {
		in, out := &in.Relabelings, &out.Relabelings
		*out = make([]*monitoringv1.RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(monitoringv1.RelabelConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MetricsRelabelings != nil {
		in, out := &in.MetricsRelabelings, &out.MetricsRelabelings
		*out = make([]*monitoringv1.RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(monitoringv1.RelabelConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(monitoringv1.TLSConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorConfig.
func (in *ServiceMonitorConfig) DeepCopy() *ServiceMonitorConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sink) DeepCopyInto(out *Sink) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sink.
func (in *Sink) DeepCopy() *Sink {
	if in == nil {
		return nil
	}
	out := new(Sink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	if in.ApacheMetricsSpec != nil {
		in, out := &in.ApacheMetricsSpec, &out.ApacheMetricsSpec
		*out = new(sources.ApacheMetricsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DemoLogsSpec != nil {
		in, out := &in.DemoLogsSpec, &out.DemoLogsSpec
		*out = new(sources.DemoLogsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VectorSpec != nil {
		in, out := &in.VectorSpec, &out.VectorSpec
		*out = new(sources.VectorSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transform) DeepCopyInto(out *Transform) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transform.
func (in *Transform) DeepCopy() *Transform {
	if in == nil {
		return nil
	}
	out := new(Transform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vector) DeepCopyInto(out *Vector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vector.
func (in *Vector) DeepCopy() *Vector {
	if in == nil {
		return nil
	}
	out := new(Vector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorList) DeepCopyInto(out *VectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorList.
func (in *VectorList) DeepCopy() *VectorList {
	if in == nil {
		return nil
	}
	out := new(VectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorSpec) DeepCopyInto(out *VectorSpec) {
	*out = *in
	if in.AgentSpec != nil {
		in, out := &in.AgentSpec, &out.AgentSpec
		*out = new(AgentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AggregatorSpec != nil {
		in, out := &in.AggregatorSpec, &out.AggregatorSpec
		*out = new(AggregatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.GlobalOptionsSpec != nil {
		in, out := &in.GlobalOptionsSpec, &out.GlobalOptionsSpec
		*out = new(GlobalOptionsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WatchNamespaces != nil {
		in, out := &in.WatchNamespaces, &out.WatchNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VectorStatus) DeepCopyInto(out *VectorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VectorStatus.
func (in *VectorStatus) DeepCopy() *VectorStatus {
	if in == nil {
		return nil
	}
	out := new(VectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeMount) DeepCopyInto(out *VolumeMount) {
	*out = *in
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeMount.
func (in *VolumeMount) DeepCopy() *VolumeMount {
	if in == nil {
		return nil
	}
	out := new(VolumeMount)
	in.DeepCopyInto(out)
	return out
}
