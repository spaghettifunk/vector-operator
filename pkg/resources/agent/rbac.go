package agent

import (
	"emperror.dev/errors"
	"github.com/banzaicloud/operator-tools/pkg/merge"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func (r *Reconciler) clusterRole() (runtime.Object, reconciler.DesiredState, error) {
	if *r.Vector.Spec.AgentSpec.Security.RoleBasedAccessControlCreate {
		return &rbacv1.ClusterRole{
			ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleName),
			Rules: []rbacv1.PolicyRule{
				{
					APIGroups: []string{""},
					Resources: []string{"pods", "namespaces"},
					Verbs:     []string{"get", "list", "watch"},
				},
			},
		}, reconciler.StatePresent, nil
	}
	return &rbacv1.ClusterRole{
		ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleName),
		Rules:      []rbacv1.PolicyRule{}}, reconciler.StateAbsent, nil
}

func (r *Reconciler) clusterRoleBinding() (runtime.Object, reconciler.DesiredState, error) {
	if *r.Vector.Spec.AgentSpec.Security.RoleBasedAccessControlCreate {
		return &rbacv1.ClusterRoleBinding{
			ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleBindingName),
			RoleRef: rbacv1.RoleRef{
				Kind:     "ClusterRole",
				APIGroup: "rbac.authorization.k8s.io",
				Name:     r.Vector.QualifiedName(clusterRoleName),
			},
			Subjects: []rbacv1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      r.getServiceAccount(),
					Namespace: r.Vector.Spec.ControlNamespace,
				},
			},
		}, reconciler.StatePresent, nil
	}
	return &rbacv1.ClusterRoleBinding{
		ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleBindingName),
		RoleRef:    rbacv1.RoleRef{}}, reconciler.StateAbsent, nil
}

func (r *Reconciler) serviceAccount() (runtime.Object, reconciler.DesiredState, error) {
	if *r.Vector.Spec.AgentSpec.Security.RoleBasedAccessControlCreate && r.Vector.Spec.AgentSpec.Security.ServiceAccount == "" {
		desired := &corev1.ServiceAccount{
			ObjectMeta: r.AgentObjectMeta(defaultServiceAccountName),
		}
		err := merge.Merge(desired, r.Vector.Spec.AgentSpec.ServiceAccountOverrides)
		if err != nil {
			return desired, reconciler.StatePresent, errors.WrapIf(err, "unable to merge overrides to base object")
		}

		return desired, reconciler.StatePresent, nil
	} else {
		desired := &corev1.ServiceAccount{
			ObjectMeta: r.AgentObjectMeta(defaultServiceAccountName),
		}
		return desired, reconciler.StateAbsent, nil
	}
}
