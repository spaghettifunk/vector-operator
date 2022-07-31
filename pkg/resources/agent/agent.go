package agent

import (
	"emperror.dev/errors"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	util "github.com/banzaicloud/operator-tools/pkg/utils"
	"github.com/go-logr/logr"
	"github.com/spaghettifunk/vector-operator/api/v1alpha1"
	"github.com/spaghettifunk/vector-operator/pkg/resources"
	"github.com/spaghettifunk/vector-operator/pkg/resources/aggregatordataprovider"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	defaultServiceAccountName      = "vector-agent"
	clusterRoleBindingName         = "vector-agent"
	clusterRoleName                = "vector-agent"
	agentSecretConfigName          = "vector-agent"
	agentDaemonSetName             = "vector-agent"
	agentPodSecurityPolicyName     = "vector-agent"
	agentServiceName               = "vector-agent"
	containerName                  = "vector-agent"
	defaultBufferVolumeMetricsPort = 9200
)

func generateVectorRefLabels(vectorRef string) map[string]string {
	return map[string]string{"app.kubernetes.io/managed-by": vectorRef}
}

func (r *Reconciler) getAgentLabels() map[string]string {
	return util.MergeLabels(r.Vector.Spec.AgentSpec.Labels, map[string]string{
		"app.kubernetes.io/name": "vector-agent"}, generateVectorRefLabels(r.Vector.ObjectMeta.GetName()))
}

func (r *Reconciler) getServiceAccount() string {
	if r.Vector.Spec.AgentSpec.Security.ServiceAccount != "" {
		return r.Vector.Spec.AgentSpec.Security.ServiceAccount
	}
	return r.Vector.QualifiedName(defaultServiceAccountName)
}

type DesiredObject struct {
	Object runtime.Object
	State  reconciler.DesiredState
}

// Reconciler holds info what resource to reconcile
type Reconciler struct {
	Vector *v1alpha1.Vector
	*reconciler.GenericResourceReconciler
	configs                map[string][]byte
	aggregatorDataProvider aggregatordataprovider.AggregatorDataProvider
}

// NewReconciler creates a new Agent reconciler
func New(client client.Client, logger logr.Logger, vector *v1alpha1.Vector, opts reconciler.ReconcilerOpts, aggregatorDataProvider aggregatordataprovider.AggregatorDataProvider) *Reconciler {
	return &Reconciler{
		Vector:                    vector,
		GenericResourceReconciler: reconciler.NewGenericReconciler(client, logger, opts),
		aggregatorDataProvider:    aggregatorDataProvider,
	}
}

// Reconcile reconciles the Agent resource
func (r *Reconciler) Reconcile() (*reconcile.Result, error) {
	for _, factory := range []resources.Resource{
		r.serviceAccount,
		r.clusterRole,
		r.clusterRoleBinding,
		r.clusterPodSecurityPolicy,
		r.pspClusterRole,
		r.pspClusterRoleBinding,
		// r.configSecret,
		r.daemonSet,
		r.serviceMetrics,
		r.monitorServiceMetrics,
		r.prometheusRules,
	} {
		o, state, err := factory()
		if err != nil {
			return nil, errors.WrapIf(err, "failed to create desired object")
		}
		if o == nil {
			return nil, errors.Errorf("Reconcile error! Resource %#v returns with nil object", factory)
		}
		result, err := r.ReconcileResource(o, state)
		if err != nil {
			return nil, errors.WrapWithDetails(err,
				"failed to reconcile resource", "resource", o.GetObjectKind().GroupVersionKind())
		}
		if result != nil {
			return result, nil
		}
	}

	return nil, nil
}

func RegisterWatches(builder *builder.Builder) *builder.Builder {
	return builder.
		Owns(&corev1.ConfigMap{}).
		Owns(&appsv1.DaemonSet{}).
		Owns(&rbacv1.ClusterRole{}).
		Owns(&rbacv1.ClusterRoleBinding{}).
		Owns(&corev1.ServiceAccount{})
}
