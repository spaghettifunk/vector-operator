package aggregator

import (
	"context"

	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	"github.com/go-logr/logr"
	"github.com/spaghettifunk/vector-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	ComponentAggregator = "vector-aggregator"
	StatefulSetName     = "vector-aggregator"
)

// Reconciler holds info what resource to reconcile
type Reconciler struct {
	Vector *v1alpha1.Vector
	*reconciler.GenericResourceReconciler
	config *string
}

type Desire struct {
	DesiredObject runtime.Object
	DesiredState  reconciler.DesiredState
	// BeforeUpdateHook has the ability to change the desired object
	// or even to change the desired state in case the object should be recreated
	BeforeUpdateHook func(runtime.Object) (reconciler.DesiredState, error)
}

func New(client client.Client, log logr.Logger,
	vector *v1alpha1.Vector, config *string, opts reconciler.ReconcilerOpts) *Reconciler {
	return &Reconciler{
		Vector:                    vector,
		GenericResourceReconciler: reconciler.NewGenericReconciler(client, log, opts),
		config:                    config,
	}
}

// Reconcile reconciles the aggregator resource
func (r *Reconciler) Reconcile() (*reconcile.Result, error) {
	_ = context.Background()

	// instantiate the statefulset here...

	return &reconcile.Result{}, nil
}

func RegisterWatches(builder *builder.Builder) *builder.Builder {
	return builder.
		Owns(&corev1.ConfigMap{}).
		Owns(&appsv1.StatefulSet{}).
		Owns(&rbacv1.ClusterRole{}).
		Owns(&rbacv1.ClusterRoleBinding{}).
		Owns(&corev1.ServiceAccount{})
}
