package agent

import (
	util "github.com/banzaicloud/operator-tools/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AgentObjectMeta creates an objectMeta for resource Agent
func (r *Reconciler) AgentObjectMeta(name string) metav1.ObjectMeta {
	o := metav1.ObjectMeta{
		Name:      r.Vector.QualifiedName(name),
		Namespace: r.Vector.Spec.ControlNamespace,
		Labels:    r.getAgentLabels(),
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion: r.Vector.APIVersion,
				Kind:       r.Vector.Kind,
				Name:       r.Vector.Name,
				UID:        r.Vector.UID,
				Controller: util.BoolPointer(true),
			},
		},
	}
	return o
}

// AgentObjectMetaClusterScope creates an cluster scoped objectMeta for resource Agent
func (r *Reconciler) AgentObjectMetaClusterScope(name string) metav1.ObjectMeta {
	o := metav1.ObjectMeta{
		Name:   r.Vector.QualifiedName(name),
		Labels: r.getAgentLabels(),
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion: r.Vector.APIVersion,
				Kind:       r.Vector.Kind,
				Name:       r.Vector.Name,
				UID:        r.Vector.UID,
				Controller: util.BoolPointer(true),
			},
		},
	}
	return o
}
