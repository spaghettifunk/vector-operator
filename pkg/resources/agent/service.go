package agent

import (
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	util "github.com/banzaicloud/operator-tools/pkg/utils"
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *Reconciler) serviceMetrics() (runtime.Object, reconciler.DesiredState, error) {
	if r.Vector.Spec.AgentSpec.Metrics != nil {
		return &corev1.Service{
			ObjectMeta: r.AgentObjectMeta(agentServiceName + "-monitor"),
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Protocol:   corev1.ProtocolTCP,
						Name:       "http-metrics",
						Port:       r.Vector.Spec.AgentSpec.Metrics.Port,
						TargetPort: intstr.IntOrString{IntVal: r.Vector.Spec.AgentSpec.Metrics.Port},
					},
				},
				Selector:  r.getAgentLabels(),
				Type:      corev1.ServiceTypeClusterIP,
				ClusterIP: "None",
			},
		}, reconciler.StatePresent, nil
	}
	return &corev1.Service{
		ObjectMeta: r.AgentObjectMeta(agentServiceName + "-monitor"),
		Spec:       corev1.ServiceSpec{}}, reconciler.StateAbsent, nil
}

func (r *Reconciler) monitorServiceMetrics() (runtime.Object, reconciler.DesiredState, error) {
	if r.Vector.Spec.AgentSpec.Metrics != nil && r.Vector.Spec.AgentSpec.Metrics.ServiceMonitor {
		objectMetadata := r.AgentObjectMeta(agentServiceName + "-metrics")
		if r.Vector.Spec.AgentSpec.Metrics.ServiceMonitorConfig.AdditionalLabels != nil {
			for k, v := range r.Vector.Spec.AgentSpec.Metrics.ServiceMonitorConfig.AdditionalLabels {
				objectMetadata.Labels[k] = v
			}
		}
		return &v1.ServiceMonitor{
			ObjectMeta: objectMetadata,
			Spec: v1.ServiceMonitorSpec{
				JobLabel:        "",
				TargetLabels:    nil,
				PodTargetLabels: nil,
				Endpoints: []v1.Endpoint{{
					Port:                 "http-metrics",
					Path:                 r.Vector.Spec.AgentSpec.Metrics.Path,
					HonorLabels:          r.Vector.Spec.AgentSpec.Metrics.ServiceMonitorConfig.HonorLabels,
					RelabelConfigs:       r.Vector.Spec.AgentSpec.Metrics.ServiceMonitorConfig.Relabelings,
					MetricRelabelConfigs: r.Vector.Spec.AgentSpec.Metrics.ServiceMonitorConfig.MetricsRelabelings,
					Scheme:               r.Vector.Spec.AgentSpec.Metrics.ServiceMonitorConfig.Scheme,
					TLSConfig:            r.Vector.Spec.AgentSpec.Metrics.ServiceMonitorConfig.TLSConfig,
				}},
				Selector: v12.LabelSelector{
					MatchLabels: util.MergeLabels(r.Vector.Spec.AgentSpec.Labels, r.getAgentLabels(), generateVectorRefLabels(r.Vector.ObjectMeta.GetName())),
				},
				NamespaceSelector: v1.NamespaceSelector{MatchNames: []string{r.Vector.Spec.ControlNamespace}},
				SampleLimit:       0,
			},
		}, reconciler.StatePresent, nil
	}
	return &v1.ServiceMonitor{
		ObjectMeta: r.AgentObjectMeta(agentServiceName + "-metrics"),
		Spec:       v1.ServiceMonitorSpec{},
	}, reconciler.StateAbsent, nil
}
