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

package controller

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/internal/kube"
	"github.com/validator-labs/validator/internal/sinks"
	"github.com/validator-labs/validator/pkg/helm"
	"github.com/validator-labs/validator/pkg/util"
	//+kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

const (
	k8sVersion         = "1.27.1"
	validatorNamespace = "validator"

	timeout  = time.Second * 30
	interval = time.Millisecond * 250
)

var (
	cfg       *rest.Config
	k8sClient client.Client
	testEnv   *envtest.Environment
	ctx       context.Context
	cancel    context.CancelFunc

	validatorNs    *corev1.Namespace
	validatorNsKey = types.NamespacedName{Name: validatorNamespace, Namespace: validatorNamespace}
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	ctx, cancel = context.WithCancel(context.TODO())

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: true,

		// The BinaryAssetsDirectory is only required if you want to run the tests directly
		// without call the makefile target test. If not informed it will look for the
		// default path defined in controller-runtime which is /usr/local/kubebuilder/.
		// Note that you must have the required binaries setup under the bin directory to perform
		// the tests directly. When we run make test it will be setup and used automatically.
		BinaryAssetsDirectory: filepath.Join(
			"..", "..", "bin", "k8s", fmt.Sprintf("%s-%s-%s", k8sVersion, runtime.GOOS, runtime.GOARCH),
		),
		UseExistingCluster: util.Ptr(false),
	}

	if os.Getenv("KUBECONFIG") != "" {
		testEnv.UseExistingCluster = util.Ptr(true)
	}

	// monkey-patch binary paths
	helm.CommandPath = filepath.Join("..", "..", "bin", "helm")

	var err error
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = v1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	//+kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	validatorNs = &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: validatorNamespace}}
	err = k8sClient.Create(ctx, validatorNs)
	Expect(err).NotTo(HaveOccurred(), "failed to create validator namespace")

	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred(), "failed to init manager")

	rawConfig, err := kube.ConvertRestConfigToRawConfig(k8sManager.GetConfig())
	Expect(err).ToNot(HaveOccurred(), "failed to generate api.Config from rest.Config")

	// start the ValidationResult controller
	err = (&ValidationResultReconciler{
		Client:     k8sManager.GetClient(),
		Log:        ctrl.Log.WithName("controllers").WithName("ValidationResult"),
		Namespace:  validatorNamespace,
		Scheme:     k8sManager.GetScheme(),
		SinkClient: sinks.NewClient(10 * time.Second),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred(), "failed to start ValidationResult controller")

	// start the ValidatorConfig controller
	err = (&ValidatorConfigReconciler{
		Client:            k8sManager.GetClient(),
		HelmClient:        helm.NewHelmClient(rawConfig),
		HelmSecretsClient: helm.NewSecretsClient(k8sManager.GetClient()),
		Log:               ctrl.Log.WithName("controllers").WithName("ValidatorConfig"),
		Scheme:            k8sManager.GetScheme(),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred(), "failed to start ValidatorConfig controller")

	go func() {
		defer GinkgoRecover()
		err = k8sManager.Start(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to run manager")
		gexec.KillAndWait(4 * time.Second)
	}()
})

var _ = AfterSuite(func() {
	By("tearing down the validator namespace")
	err := k8sClient.Delete(ctx, validatorNs)
	Expect(err).ToNot(HaveOccurred(), "failed to tear down the validator namespace")
	Eventually(func() bool {
		// Skip namespace teardown if running in EnvTest
		// https://book.kubebuilder.io/reference/envtest.html#namespace-usage-limitation
		if os.Getenv("KUBECONFIG") == "" {
			return true
		}
		err := k8sClient.Get(ctx, validatorNsKey, validatorNs)
		return apierrs.IsNotFound(err)
	}, 1*time.Minute, interval).Should(BeTrue(), "failed to tear down the validator namespace")

	By("tearing down the test environment")
	cancel()
	err = (func() (err error) {
		// Need to sleep if the first stop fails due to a bug:
		// https://github.com/kubernetes-sigs/controller-runtime/issues/1571
		sleepTime := 1 * time.Millisecond
		for i := 0; i < 12; i++ { // Exponentially sleep up to ~4s
			if err = testEnv.Stop(); err == nil {
				return
			}
			sleepTime *= 2
			time.Sleep(sleepTime)
		}
		return
	})()
	Expect(err).NotTo(HaveOccurred(), "failed to tear down the test environment")
})
