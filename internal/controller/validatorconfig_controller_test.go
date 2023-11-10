package controller

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	kyaml "sigs.k8s.io/yaml"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
	//+kubebuilder:scaffold:imports
)

const (
	networkPluginDeploymentName  = "validator-plugin-network-controller-manager"
	networkPluginDeploymentImage = "quay.io/spectrocloud-labs/validator-plugin-network"
	networkPluginVersionPre      = "v0.0.4"
	networkPluginVersionPost     = "v0.0.5"
	validatorConfigName          = "validator-config"
)

var (
	vcNetwork         = filepath.Join("testdata", "vc-network.yaml")
	expectedImagePre  = fmt.Sprintf("%s:%s", networkPluginDeploymentImage, networkPluginVersionPre)
	expectedImagePost = fmt.Sprintf("%s:%s", networkPluginDeploymentImage, networkPluginVersionPost)
)

var _ = Describe("ValidatorConfig controller", Ordered, func() {
	BeforeEach(func() {
		// comment in or out to disable the ValidatorConfig controller specs
		if true {
			Skip("skipping")
		}
	})

	vc := &v1alpha1.ValidatorConfig{}
	vcKey := types.NamespacedName{Name: validatorConfigName, Namespace: validatorNamespace}
	networkPluginDeploymentKey := types.NamespacedName{Name: networkPluginDeploymentName, Namespace: validatorNamespace}
	networkPluginDeployment := &appsv1.Deployment{}

	It("Should deploy validator-plugin-network once a ValidatorConfig is created", func() {
		By("By creating a new ValidatorConfig")
		ctx := context.Background()

		vcBytes, err := os.ReadFile(vcNetwork)
		Expect(err).NotTo(HaveOccurred())

		err = kyaml.Unmarshal(vcBytes, vc)
		Expect(err).NotTo(HaveOccurred())

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

		Expect(k8sClient.Get(ctx, vcKey, vc)).Should(Succeed())

		vc.Spec.Plugins[0].Chart.Version = networkPluginVersionPost
		vc.Spec.Plugins[0].Values = strings.ReplaceAll(
			vc.Spec.Plugins[0].Values, networkPluginVersionPre, networkPluginVersionPost,
		)

		Expect(k8sClient.Update(ctx, vc)).Should(Succeed())

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
})
