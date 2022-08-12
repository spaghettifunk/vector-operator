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

package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"

	prometheusOperator "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/spf13/cast"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	ctrl "sigs.k8s.io/controller-runtime"

	vectorv1alpha1 "github.com/spaghettifunk/vector-operator/api/v1alpha1"
	"github.com/spaghettifunk/vector-operator/controllers"
	"github.com/spaghettifunk/vector-operator/pkg/k8sutil"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(vectorv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
	_ = prometheusOperator.AddToScheme(scheme)
	_ = apiextensions.AddToScheme(scheme)
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var verboseLogging bool
	var loggingOutputFormat string
	var enableprofile bool
	var namespace string
	var vectorRef string
	var klogLevel int

	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.BoolVar(&verboseLogging, "verbose", false, "Enable verbose logging")
	flag.StringVar(&loggingOutputFormat, "output-format", "", "Logging output format (json, console)")
	flag.IntVar(&klogLevel, "klogLevel", 0, "Global log level for klog (0-9)")
	flag.BoolVar(&enableprofile, "pprof", false, "Enable pprof")
	flag.StringVar(&namespace, "watch-namespace", "", "Namespace to filter the list of watched objects")
	flag.StringVar(&vectorRef, "watch-vector-name", "", "Vector resource name to optionally filter the list of watched objects based on which vector they belong to by checking the app.kubernetes.io/managed-by label")
	flag.Parse()

	zapLogger := zap.New(func(o *zap.Options) {
		o.Development = verboseLogging

		switch loggingOutputFormat {
		case "json":
			encoder := zap.JSONEncoder()
			encoder(o)
		case "console":
			encoder := zap.ConsoleEncoder()
			encoder(o)
		case "":
			break
		default:
			fmt.Printf("invalid encoder value \"%s\"", loggingOutputFormat)
			os.Exit(1)
		}
	})

	ctrl.SetLogger(zapLogger)

	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	err := klogFlags.Set("v", cast.ToString(klogLevel))
	if err != nil {
		fmt.Printf("%s - failed to set log level for klog, moving on.\n", err)
	}
	klog.SetLogger(zapLogger)

	mgrOptions := ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "vector-operator." + vectorv1alpha1.GroupVersion.Group,
		MapperProvider:     k8sutil.NewCached,
		Port:               9443,
	}

	customMgrOptions, err := setupCustomCache(&mgrOptions, namespace, vectorRef)
	if err != nil {
		setupLog.Error(err, "unable to set up custom cache settings")
		os.Exit(1)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), *customMgrOptions)

	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if enableprofile {
		setupLog.Info("enabling pprof")
		err = mgr.AddMetricsExtraHandler("/debug/pprof/", http.HandlerFunc(pprof.Index))
		if err != nil {
			setupLog.Error(err, "unable to attach pprof to webserver")
			os.Exit(1)
		}
	}

	vectorReconciler := controllers.NewVectorReconciler(mgr.GetClient(), ctrl.Log.WithName("controllers").WithName("Logging"))

	if err := controllers.SetupVectorWithManager(mgr, ctrl.Log.WithName("manager")).Complete(vectorReconciler); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Logging")
		os.Exit(1)
	}

	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func setupCustomCache(mgrOptions *ctrl.Options, namespace string, vectorRef string) (*ctrl.Options, error) {
	if namespace == "" && vectorRef == "" {
		return mgrOptions, nil
	}

	var namespaceSelector fields.Selector
	var labelSelector labels.Selector
	if namespace != "" {
		namespaceSelector = fields.Set{"metadata.namespace": namespace}.AsSelector()
	}
	if vectorRef != "" {
		labelSelector = labels.Set{"app.kubernetes.io/managed-by": vectorRef}.AsSelector()
	}

	selectorsByObject := cache.SelectorsByObject{
		&corev1.Pod{}: {
			Field: namespaceSelector,
			Label: labelSelector,
		},
		&appsv1.DaemonSet{}: {
			Field: namespaceSelector,
			Label: labelSelector,
		},
		&appsv1.StatefulSet{}: {
			Field: namespaceSelector,
			Label: labelSelector,
		},
		&appsv1.Deployment{}: {
			Field: namespaceSelector,
			Label: labelSelector,
		},
		&corev1.PersistentVolumeClaim{}: {
			Field: namespaceSelector,
			Label: labelSelector,
		},
	}

	mgrOptions.NewCache = cache.BuilderWithOptions(cache.Options{SelectorsByObject: selectorsByObject})

	return mgrOptions, nil
}
