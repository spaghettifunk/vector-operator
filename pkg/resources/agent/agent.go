package agent

import (
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	util "github.com/banzaicloud/operator-tools/pkg/utils"
	"github.com/go-logr/logr"
	"github.com/spaghettifunk/vector-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	defaultServiceAccountName      = "fluentbit"
	clusterRoleBindingName         = "fluentbit"
	clusterRoleName                = "fluentbit"
	fluentBitSecretConfigName      = "fluentbit"
	fluentbitDaemonSetName         = "fluentbit"
	fluentbitPodSecurityPolicyName = "fluentbit"
	fluentbitServiceName           = "fluentbit"
	containerName                  = "fluent-bit"
	defaultBufferVolumeMetricsPort = 9200
)

func generateLoggingRefLabels(loggingRef string) map[string]string {
	return map[string]string{"app.kubernetes.io/managed-by": loggingRef}
}

func (r *Reconciler) getAgentLabels() map[string]string {
	return util.MergeLabels(r.Vector.Spec.AgentSpec.Labels, map[string]string{
		"app.kubernetes.io/name": "fluentbit"}, generateLoggingRefLabels(r.Vector.ObjectMeta.GetName()))
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
	configs map[string][]byte
}

// NewReconciler creates a new Fluentbit reconciler
func New(client client.Client, logger logr.Logger, vector *v1alpha1.Vector, opts reconciler.ReconcilerOpts, fluentdDataProvider fluentddataprovider.FluentdDataProvider) *Reconciler {
	return &Reconciler{
		Vector:                    vector,
		GenericResourceReconciler: reconciler.NewGenericReconciler(client, logger, opts),
	}
}
