package agent

import (
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	util "github.com/banzaicloud/operator-tools/pkg/utils"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"

	"k8s.io/apimachinery/pkg/runtime"
)

func (r *Reconciler) clusterPodSecurityPolicy() (runtime.Object, reconciler.DesiredState, error) {
	if r.Vector.Spec.AgentSpec.Security.PodSecurityPolicyCreate {
		allowedHostPaths := []policyv1beta1.AllowedHostPath{{
			PathPrefix: r.Vector.Spec.AgentSpec.MountPath,
			ReadOnly:   true,
		}, {
			PathPrefix: "/var/log",
			ReadOnly:   true,
		}}

		for _, vMnt := range r.Vector.Spec.AgentSpec.ExtraVolumeMounts {
			allowedHostPaths = append(allowedHostPaths, policyv1beta1.AllowedHostPath{
				PathPrefix: vMnt.Source,
				ReadOnly:   *vMnt.ReadOnly,
			})
		}

		return &policyv1beta1.PodSecurityPolicy{
			ObjectMeta: r.AgentObjectMetaClusterScope(agentPodSecurityPolicyName),
			Spec: policyv1beta1.PodSecurityPolicySpec{
				Volumes: []policyv1beta1.FSType{
					"configMap",
					"emptyDir",
					"secret",
					"hostPath"},
				SELinux: policyv1beta1.SELinuxStrategyOptions{
					Rule: policyv1beta1.SELinuxStrategyRunAsAny,
				},
				RunAsUser: policyv1beta1.RunAsUserStrategyOptions{
					Rule: policyv1beta1.RunAsUserStrategyRunAsAny,
				},
				SupplementalGroups: policyv1beta1.SupplementalGroupsStrategyOptions{
					Rule: policyv1beta1.SupplementalGroupsStrategyRunAsAny,
				},
				FSGroup: policyv1beta1.FSGroupStrategyOptions{
					Rule: policyv1beta1.FSGroupStrategyRunAsAny,
				},
				ReadOnlyRootFilesystem:   true,
				AllowPrivilegeEscalation: util.BoolPointer(false),
				AllowedHostPaths:         allowedHostPaths,
			},
		}, reconciler.StatePresent, nil
	}
	return &policyv1beta1.PodSecurityPolicy{
		ObjectMeta: r.AgentObjectMeta(agentPodSecurityPolicyName),
		Spec:       policyv1beta1.PodSecurityPolicySpec{},
	}, reconciler.StateAbsent, nil
}

func (r *Reconciler) pspClusterRole() (runtime.Object, reconciler.DesiredState, error) {
	if *r.Vector.Spec.AgentSpec.Security.RoleBasedAccessControlCreate && r.Vector.Spec.AgentSpec.Security.PodSecurityPolicyCreate {
		return &rbacv1.ClusterRole{
			ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleName + "-psp"),
			Rules: []rbacv1.PolicyRule{
				{
					APIGroups:     []string{"policy"},
					Resources:     []string{"podsecuritypolicies"},
					ResourceNames: []string{r.Vector.QualifiedName(agentPodSecurityPolicyName)},
					Verbs:         []string{"use"},
				},
			},
		}, reconciler.StatePresent, nil
	}
	return &rbacv1.ClusterRole{
		ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleName + "-psp"),
		Rules:      []rbacv1.PolicyRule{}}, reconciler.StateAbsent, nil
}

func (r *Reconciler) pspClusterRoleBinding() (runtime.Object, reconciler.DesiredState, error) {
	if *r.Vector.Spec.AgentSpec.Security.RoleBasedAccessControlCreate && r.Vector.Spec.AgentSpec.Security.PodSecurityPolicyCreate {
		return &rbacv1.ClusterRoleBinding{
			ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleBindingName + "-psp"),
			RoleRef: rbacv1.RoleRef{
				Kind:     "ClusterRole",
				APIGroup: "rbac.authorization.k8s.io",
				Name:     r.Vector.QualifiedName(clusterRoleName + "-psp"),
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
		ObjectMeta: r.AgentObjectMetaClusterScope(clusterRoleBindingName + "-psp"),
		RoleRef:    rbacv1.RoleRef{}}, reconciler.StateAbsent, nil
}
