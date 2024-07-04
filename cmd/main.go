/*
Copyright 2023.

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

// Package main initializes the ValidationConfig and ValidationResult controllers.
package main

import (
	"flag"
	"os"
	"time"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	validationv1alpha1 "github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/internal/controller"
	"github.com/validator-labs/validator/internal/kube"
	"github.com/validator-labs/validator/internal/sinks"
	"github.com/validator-labs/validator/pkg/helm"
	"github.com/validator-labs/validator/pkg/helm/release"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(validationv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false, "Enable leader election for controller manager. "+
		"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	ns := os.Getenv("NAMESPACE")

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "0d82a3fe.spectrocloud.labs",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,

		// NewCache: func(config *rest.Config, opts cache.Options) (cache.Cache, error) {
		// 	opts.DefaultNamespaces = map[string]cache.Config{
		// 		ns: {},
		// 	}
		// 	return cache.New(config, opts)
		// },
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	rawConfig, err := kube.ConvertRestConfigToRawConfig(mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "unable to get config")
		os.Exit(1)
	}

	timeout, ok := os.LookupEnv("SINK_WEBHOOK_TIMEOUT_SECONDS")
	if !ok {
		timeout = "30s"
	}
	sinkTimeout, err := time.ParseDuration(timeout)
	if err != nil {
		setupLog.Error(err, "failed to parse webhook timeout")
		os.Exit(1)
	}
	sinkClient := sinks.NewClient(sinkTimeout)

	if err = (&controller.ValidationResultReconciler{
		Client:     mgr.GetClient(),
		Log:        ctrl.Log.WithName("controllers").WithName("ValidationResult"),
		Namespace:  ns,
		Scheme:     mgr.GetScheme(),
		SinkClient: sinkClient,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ValidationResult")
		os.Exit(1)
	}

	if err = (&controller.ValidatorConfigReconciler{
		Client:            mgr.GetClient(),
		HelmClient:        helm.NewHelmClient(rawConfig),
		HelmReleaseClient: release.NewHelmReleaseClient(mgr.GetClient()),
		Log:               ctrl.Log.WithName("controllers").WithName("ValidatorConfig"),
		Scheme:            mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ValidatorConfig")
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
