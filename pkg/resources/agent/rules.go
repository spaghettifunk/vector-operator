package agent

import (
	"fmt"

	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *Reconciler) prometheusRules() (runtime.Object, reconciler.DesiredState, error) {
	obj := &v1.PrometheusRule{
		ObjectMeta: r.AgentObjectMeta(agentServiceName),
	}
	state := reconciler.StateAbsent

	if r.Vector.Spec.AgentSpec.Metrics != nil && r.Vector.Spec.AgentSpec.Metrics.PrometheusRules {
		nsJobLabel := fmt.Sprintf(`job="%s", namespace="%s"`, obj.Name, obj.Namespace)
		state = reconciler.StatePresent
		obj.Spec.Groups = []v1.RuleGroup{{
			Name: "vector_agent",
			Rules: []v1.Rule{
				{
					Alert: "VectorAgentTooManyErrors",
					Expr:  intstr.FromString(fmt.Sprintf("rate(vector_agent_output_retries_failed_total{%s}[10m]) > 0", nsJobLabel)),
					For:   "10m",
					Labels: map[string]string{
						"service":  "vector-agent",
						"severity": "warning",
					},
					Annotations: map[string]string{
						"summary":     `Vector Agent too many errors.`,
						"description": `Vector Agent ({{ $labels.instance }}) is erroring.`,
					},
				},
			},
		},
		}
	}
	return obj, state, nil
}
