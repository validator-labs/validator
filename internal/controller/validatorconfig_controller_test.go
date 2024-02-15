package controller

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kyaml "sigs.k8s.io/yaml"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/internal/test"
	"github.com/spectrocloud-labs/validator/pkg/helm"
	//+kubebuilder:scaffold:imports
)

const (
	networkPluginDeploymentName  = "validator-plugin-network-controller-manager"
	networkPluginDeploymentImage = "quay.io/spectrocloud-labs/validator-plugin-network"
	networkPluginVersionPre      = "v0.0.4"
	networkPluginVersionPost     = "v0.0.5"
)

var (
	vcNetwork         = filepath.Join("testdata", "vc-network.yaml")
	expectedImagePre  = fmt.Sprintf("%s:%s", networkPluginDeploymentImage, networkPluginVersionPre)
	expectedImagePost = fmt.Sprintf("%s:%s", networkPluginDeploymentImage, networkPluginVersionPost)
)

var _ = Describe("ValidatorConfig controller", Ordered, func() {

	BeforeEach(func() {
		// toggle true/false to enable/disable the ValidatorConfig controller specs
		if false {
			Skip("skipping")
		}
	})

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "validator-plugin-network-invalid-auth",
			Namespace: validatorNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("foo"),
		},
	}

	validSecret := secret.DeepCopy()
	validSecret.Name = "validator-plugin-network-chart-secret"
	validSecret.Data = map[string][]byte{
		"username": []byte("foo"),
		"password": []byte("bar"),
	}

	vc := &v1alpha1.ValidatorConfig{}
	vcKey := types.NamespacedName{Name: "validator-config-test", Namespace: validatorNamespace}
	networkPluginDeploymentKey := types.NamespacedName{Name: networkPluginDeploymentName, Namespace: validatorNamespace}
	networkPluginDeployment := &appsv1.Deployment{}

	It("Should deploy validator-plugin-network once a ValidatorConfig is created", func() {
		By("By creating a new ValidatorConfig")
		ctx := context.Background()

		vcBytes, err := os.ReadFile(vcNetwork)
		Expect(err).NotTo(HaveOccurred())

		err = kyaml.Unmarshal(vcBytes, vc)
		Expect(err).NotTo(HaveOccurred())

		Expect(k8sClient.Create(ctx, secret)).Should(Succeed())
		Expect(k8sClient.Create(ctx, validSecret)).Should(Succeed())
		Expect(k8sClient.Create(ctx, vc)).Should(Succeed())

		// Wait for the validator-plugin-network Deployment to be deployed
		Eventually(func() bool {
			if err := k8sClient.Get(ctx, networkPluginDeploymentKey, networkPluginDeployment); err != nil {
				return false
			}
			imageOk := networkPluginDeployment.Spec.Template.Spec.Containers[1].Image == expectedImagePre
			healthy := networkPluginDeployment.Status.ReadyReplicas == networkPluginDeployment.Status.Replicas
			return imageOk && healthy
		}, timeout, interval).Should(BeTrue(), "failed to deploy validator-plugin-network")
	})

	It("Should upgrade validator-plugin-network once the ValidatorConfig is updated", func() {
		By("Updating the ValidatorConfig")
		ctx := context.Background()

		Eventually(func() bool {
			if err := k8sClient.Get(ctx, vcKey, vc); err != nil {
				return false
			}
			vc.Spec.Plugins[0].Chart.Version = networkPluginVersionPost
			vc.Spec.Plugins[0].Values = strings.ReplaceAll(
				vc.Spec.Plugins[0].Values, networkPluginVersionPre, networkPluginVersionPost,
			)
			if err := k8sClient.Update(ctx, vc); err != nil {
				return false
			}
			return true
		}, timeout, interval).Should(BeTrue(), "failed to update validator-plugin-network")

		// Wait for the validator-plugin-network Deployment to be upgraded
		Eventually(func() bool {
			if err := k8sClient.Get(ctx, networkPluginDeploymentKey, networkPluginDeployment); err != nil {
				return false
			}
			imageOk := networkPluginDeployment.Spec.Template.Spec.Containers[1].Image == expectedImagePost
			healthy := networkPluginDeployment.Status.ReadyReplicas == networkPluginDeployment.Status.Replicas
			return imageOk && healthy
		}, timeout, interval).Should(BeTrue(), "failed to upgrade validator-plugin-network")
	})

	It("Should uninstall validator-plugin-network once the ValidatorConfig is deleted", func() {
		By("Deleting the ValidatorConfig")
		ctx := context.Background()

		Expect(k8sClient.Delete(ctx, vc)).Should(Succeed())

		// Wait for validator-plugin-network Deployment to be deleted
		Eventually(func() bool {
			err := k8sClient.Get(ctx, networkPluginDeploymentKey, networkPluginDeployment)
			return apierrs.IsNotFound(err)
		}, timeout, interval).Should(BeTrue(), "failed to uninstall validator-plugin-network")
	})

	It("Should fail to deploy validator-plugin-network-invalid-auth once a ValidatorConfig is created", func() {
		By("By creating a new ValidatorConfig")
		ctx := context.Background()

		vc := &v1alpha1.ValidatorConfig{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "validator-plugin-network-invalid-auth",
				Namespace: validatorNamespace,
			},
			Spec: v1alpha1.ValidatorConfigSpec{
				Plugins: []v1alpha1.HelmRelease{
					{
						Chart: v1alpha1.HelmChart{
							Name:           "foo",
							Repository:     "bar",
							AuthSecretName: "chart-secret",
						},
					},
				},
			},
		}
		vcKey := types.NamespacedName{Name: "validator-plugin-network-invalid-auth", Namespace: validatorNamespace}

		Expect(k8sClient.Create(ctx, vc)).Should(Succeed())

		// Wait for the validator-plugin-network-invalid-auth deployment to fail
		Eventually(func() bool {
			if err := k8sClient.Get(ctx, vcKey, vc); err != nil {
				return false
			}
			condition, ok := isConditionTrue(vc, "foo", v1alpha1.HelmChartDeployedCondition)
			return condition.Status == corev1.ConditionFalse && !ok
		}, timeout, interval).Should(BeTrue(), "failed to deploy validator-plugin-network")
	})
})

func TestConfigureHelmBasicAuth(t *testing.T) {
	cs := []struct {
		name       string
		reconciler ValidatorConfigReconciler
		nn         types.NamespacedName
		opts       *helm.Options
		expected   error
	}{
		{
			name: "Fail (get_secret)",
			nn:   types.NamespacedName{Name: "chart-secret", Namespace: validatorNamespace},
			reconciler: ValidatorConfigReconciler{
				Client: test.ClientMock{
					GetErrors: []error{errors.New("get failed")},
				},
			},
			opts:     &helm.Options{Chart: "foo", Repo: "bar"},
			expected: errors.New("failed to get auth secret chart-secret in namespace validator for chart foo in repo bar: get failed"),
		},
		{
			name: "Fail (get_username)",
			nn:   types.NamespacedName{Name: "chart-secret", Namespace: validatorNamespace},
			reconciler: ValidatorConfigReconciler{
				Client: test.ClientMock{},
			},
			opts:     &helm.Options{Chart: "foo", Repo: "bar"},
			expected: errors.New("auth secret for chart foo in repo bar missing required key: 'username'"),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		err := c.reconciler.configureHelmOpts(c.nn, c.opts)
		if err != nil && !reflect.DeepEqual(err.Error(), c.expected.Error()) {
			t.Errorf("expected (%v), got (%v)", c.expected, err)
		}
	}
}

func TestEmitFinalizeCleanup(t *testing.T) {
	cs := []struct {
		name         string
		reconciler   ValidatorConfigReconciler
		env          map[string]string
		expectedErrs []error
	}{
		{
			name:         "CLEANUP_GRPC_SERVER_ENABLED is empty",
			reconciler:   ValidatorConfigReconciler{},
			env:          map[string]string{},
			expectedErrs: nil,
		},
		{
			name:         "CLEANUP_GRPC_SERVER_ENABLED is disabled",
			reconciler:   ValidatorConfigReconciler{},
			env:          map[string]string{"CLEANUP_GRPC_SERVER_ENABLED": "false"},
			expectedErrs: nil,
		},
		{
			name:       "CLEANUP_GRPC_SERVER_HOST is empty",
			reconciler: ValidatorConfigReconciler{},
			env: map[string]string{
				"CLEANUP_GRPC_SERVER_ENABLED": "true",
				"CLEANUP_GRPC_SERVER_PORT":    "1234",
			},
			expectedErrs: []error{errors.New("CLEANUP_GRPC_SERVER_HOST is invalid")},
		},
		{
			name:       "CLEANUP_GRPC_SERVER_PORT is empty",
			reconciler: ValidatorConfigReconciler{},
			env: map[string]string{
				"CLEANUP_GRPC_SERVER_ENABLED": "true",
				"CLEANUP_GRPC_SERVER_HOST":    "localhost",
			},
			expectedErrs: []error{errors.New(`CLEANUP_GRPC_SERVER_PORT is invalid: strconv.Atoi: parsing "": invalid syntax`)},
		},
		{
			name:       "CLEANUP_GRPC_SERVER_PORT is invalid",
			reconciler: ValidatorConfigReconciler{},
			env: map[string]string{
				"CLEANUP_GRPC_SERVER_ENABLED": "true",
				"CLEANUP_GRPC_SERVER_HOST":    "localhost",
				"CLEANUP_GRPC_SERVER_PORT":    "abcd",
			},
			expectedErrs: []error{errors.New(`CLEANUP_GRPC_SERVER_PORT is invalid: strconv.Atoi: parsing "abcd": invalid syntax`)},
		},
		{
			name:       "Request fails",
			reconciler: ValidatorConfigReconciler{},
			env: map[string]string{
				"CLEANUP_GRPC_SERVER_ENABLED": "true",
				"CLEANUP_GRPC_SERVER_HOST":    "localhost",
				"CLEANUP_GRPC_SERVER_PORT":    "1234",
			},
			expectedErrs: []error{
				errors.New(`FinalizeCleanup request to http://localhost:1234 failed: unavailable: dial tcp [::1]:1234: connect: connection refused`),
				errors.New(`FinalizeCleanup request to http://localhost:1234 failed: unavailable: dial tcp 127.0.0.1:1234: connect: connection refused`),
			},
		},
	}
	for _, c := range cs {
		t.Log(c.name)

		os.Clearenv()
		for k, v := range c.env {
			os.Setenv(k, v)
		}

		err := c.reconciler.emitFinalizeCleanup()
		errOk := err == nil && c.expectedErrs == nil
		for _, expectedErr := range c.expectedErrs {
			if err != nil && reflect.DeepEqual(err.Error(), expectedErr.Error()) {
				errOk = true
				break
			}
		}
		if !errOk {
			t.Errorf("expected one of (%v), got (%v)", c.expectedErrs, err)
		}
	}
}
