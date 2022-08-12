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

package controllers

import (
	"bytes"
	"context"
	"regexp"

	"emperror.dev/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	"github.com/banzaicloud/operator-tools/pkg/secret"
	"github.com/banzaicloud/operator-tools/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	vectorv1alpha1 "github.com/spaghettifunk/vector-operator/api/v1alpha1"
	"github.com/spaghettifunk/vector-operator/pkg/model"
	"github.com/spaghettifunk/vector-operator/pkg/resources"
	"github.com/spaghettifunk/vector-operator/pkg/resources/agent"
	"github.com/spaghettifunk/vector-operator/pkg/resources/aggregator"
	corev1 "k8s.io/api/core/v1"

	"github.com/go-logr/logr"
)

// VectorReconciler reconciles a Vector object
type VectorReconciler struct {
	client.Client
	Log logr.Logger
}

// NewVectorReconciler returns a new VectorReconciler instance
func NewVectorReconciler(client client.Client, log logr.Logger) *VectorReconciler {
	return &VectorReconciler{
		Client: client,
		Log:    log,
	}
}

//+kubebuilder:rbac:groups=dev.vector,resources=vectors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dev.vector,resources=vectors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dev.vector,resources=vectors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *VectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("vector", req.NamespacedName)

	var vector vectorv1alpha1.Vector
	if err := r.Client.Get(ctx, req.NamespacedName, &vector); err != nil {
		// If object is not found, return without error.
		// Created objects are automatically garbage collected.
		// For additional cleanup logic use finalizers.
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	if err := vector.SetDefaults(); err != nil {
		return reconcile.Result{}, err
	}

	reconcilerOpts := reconciler.ReconcilerOpts{
		RecreateErrorMessageCondition: reconciler.MatchImmutableErrorMessages,
	}

	vectorResources, err := model.NewVectorResourceRepository(r.Client).VectorResourcesFor(ctx, vector)
	if err != nil {
		return reconcile.Result{}, errors.WrapIfWithDetails(err, "failed to get vector resources", "vector", vector)
	}
	// metrics
	defer func() {
		gv := getResourceStateMetrics(log)
		gv.Reset()
		for _, ob := range vectorResources.Pipelines {
			updateResourceStateMetrics(&ob, utils.PointerToBool(ob.Status.Active), gv)
		}
	}()

	reconcilers := []resources.ComponentReconciler{
		// model.NewValidationReconciler(ctx, r.Client, vectorResources, &secretLoaderFactory{Client: r.Client}),
	}

	// if vector.Spec.AggregatorSpec != nil {
	// 	aggregatorConfig, err := r.clusterConfiguration(vectorResources)
	// 	if err != nil {
	// 		reconcilers = append(reconcilers, func() (*reconcile.Result, error) {
	// 			return &reconcile.Result{}, err
	// 		})
	// 	} else {
	// 		log.V(1).Info("agent configuration", "config", aggregatorConfig)

	// 		reconcilers = append(reconcilers, aggregator.New(r.Client, r.Log, &vector, &aggregatorConfig, reconcilerOpts).Reconcile)
	// 	}
	// }

	if vector.Spec.AgentSpec != nil {
		reconcilers = append(reconcilers, agent.New(r.Client, r.Log, &vector, reconcilerOpts, aggregator.NewDataProvider(r.Client)).Reconcile)
	}

	for _, rec := range reconcilers {
		result, err := rec()
		if err != nil {
			return reconcile.Result{}, err
		}
		if result != nil {
			// short circuit if requested explicitly
			return *result, err
		}
	}
	return ctrl.Result{}, nil
}

func updateResourceStateMetrics(obj client.Object, active bool, gv *prometheus.GaugeVec) {
	gv.With(prometheus.Labels{"name": obj.GetName(), "namespace": obj.GetNamespace(), "status": "active", "kind": obj.GetObjectKind().GroupVersionKind().Kind}).Set(boolToFloat64(active))
	gv.With(prometheus.Labels{"name": obj.GetName(), "namespace": obj.GetNamespace(), "status": "inactive", "kind": obj.GetObjectKind().GroupVersionKind().Kind}).Set(boolToFloat64(!active))
}

func getResourceStateMetrics(logger logr.Logger) *prometheus.GaugeVec {
	gv := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "logging_resource_state"}, []string{"name", "namespace", "status", "kind"})
	err := metrics.Registry.Register(gv)
	if err != nil {
		if err, ok := err.(prometheus.AlreadyRegisteredError); ok {
			if gv, ok = err.ExistingCollector.(*prometheus.GaugeVec); !ok {
				logger.Error(err, "already registered metric name with different type ", "metric", gv)
			}
		} else {
			logger.Error(err, "couldn't register metrics vector for resource", "metric", gv)
		}
	}
	return gv
}

func boolToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func (r *VectorReconciler) clusterConfiguration(resources model.VectorResources) (string, error) {
	// vectorConfig, err := model.CreateSystem(resources, r.Log)
	// if err != nil {
	// 	return "", errors.WrapIfWithDetails(err, "failed to build model", "vector", resources.Vector)
	// }

	output := &bytes.Buffer{}
	// renderer := render.VectorRender{
	// 	Out:    output,
	// 	Indent: 2,
	// }
	// if err := renderer.Render(vectorConfig); err != nil {
	// 	return "", errors.WrapIfWithDetails(err, "failed to render vector config", "vector", resources.Vector)
	// }

	return output.String(), nil
}

type secretLoaderFactory struct {
	Client  client.Client
	Secrets secret.MountSecrets
}

func SetupVectorWithManager(mgr ctrl.Manager, logger logr.Logger) *ctrl.Builder {
	requestMapper := handler.EnqueueRequestsFromMapFunc(func(obj client.Object) []reconcile.Request {
		// get all the logging resources from the cache
		var vectorList vectorv1alpha1.VectorList
		if err := mgr.GetCache().List(context.TODO(), &vectorList); err != nil {
			logger.Error(err, "failed to list logging resources")
			return nil
		}

		switch o := obj.(type) {
		case *vectorv1alpha1.Pipeline:
			return reconcileRequestsForVectorRef(vectorList.Items, o.Spec.VectorRef)
		case *corev1.Secret:
			r := regexp.MustCompile("dev.vector/(.*)")
			var requestList []reconcile.Request
			for key := range o.Annotations {
				if result := r.FindStringSubmatch(key); len(result) > 1 {
					vectorRef := result[1]
					// When vectorRef is "default" we also trigger for the empty ("") vectorRef as well,
					// because the empty string cannot be used in the annotation, thus "default" refers to the empty case.
					if vectorRef == "default" {
						requestList = append(requestList, reconcileRequestsForVectorRef(vectorList.Items, "")...)
					}
					requestList = append(requestList, reconcileRequestsForVectorRef(vectorList.Items, vectorRef)...)
				}
			}
			return requestList
		}
		return nil
	})

	builder := ctrl.NewControllerManagedBy(mgr).
		For(&vectorv1alpha1.Vector{}).
		Owns(&corev1.Pod{}).
		Watches(&source.Kind{Type: &vectorv1alpha1.Pipeline{}}, requestMapper).
		Watches(&source.Kind{Type: &corev1.Secret{}}, requestMapper)

	agent.RegisterWatches(builder)
	aggregator.RegisterWatches(builder)

	return builder
}

func reconcileRequestsForVectorRef(loggings []vectorv1alpha1.Vector, vectorRef string) (reqs []reconcile.Request) {
	for _, l := range loggings {
		if l.Spec.VectorRef == vectorRef {
			reqs = append(reqs, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: l.Namespace, // this happens to be empty as long as Vector is cluster scoped
					Name:      l.Name,
				},
			})
		}
	}
	return
}
