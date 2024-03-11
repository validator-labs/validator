package controller

import (
	"context"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	kyaml "sigs.k8s.io/yaml"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/pkg/constants"
	//+kubebuilder:scaffold:imports
)

const validationResultName = "validator-plugin-aws-service-quota"

var vrServiceQuota = filepath.Join("testdata", "vr-aws-service-quota.yaml")

var _ = Describe("ValidationResult controller", Ordered, func() {

	BeforeEach(func() {
		// toggle true/false to enable/disable the ValidationResult controller specs
		if false {
			Skip("skipping")
		}
	})

	vc := &v1alpha1.ValidatorConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.ValidatorConfig,
			Namespace: validatorNamespace,
		},
		Spec: v1alpha1.ValidatorConfigSpec{
			Sink: &v1alpha1.Sink{
				Type:       "slack",
				SecretName: "slack-secret",
			},
		},
	}

	vr := &v1alpha1.ValidationResult{
		ObjectMeta: metav1.ObjectMeta{
			Name:      validationResultName,
			Namespace: validatorNamespace,
		},
		Spec: v1alpha1.ValidationResultSpec{
			Plugin:          "AWS",
			ExpectedResults: 4,
		},
	}
	vrKey := types.NamespacedName{Name: validationResultName, Namespace: validatorNamespace}

	It("Should create a ValidationResult before the ValidatorConfig is created and fail initially", func() {
		By("By creating a ValidationResult before creating the ValidatorConfig")
		ctx := context.Background()

		vrPre := vr.DeepCopy()
		vrPre.Name = "vr-copy"
		Expect(k8sClient.Create(ctx, vrPre)).Should(Succeed())
	})

	It("Should hash the ValidationResult and update its Status once a ValidationResult is created", func() {
		By("By creating a new ValidationResult")
		ctx := context.Background()

		Expect(k8sClient.Create(ctx, vc)).Should(Succeed())
		Expect(k8sClient.Create(ctx, vr)).Should(Succeed())

		// Mark the ValidationResult as InProgress
		Eventually(func() bool {
			if err := k8sClient.Get(ctx, vrKey, vr); err != nil {
				return false
			}
			vr.Status.State = v1alpha1.ValidationInProgress
			if err := k8sClient.Status().Update(ctx, vr); err != nil {
				return false
			}
			return true
		}, timeout, interval).Should(BeTrue(), "failed to update ValidationResult Status")

		// Wait for the ValidationResult's Status to be updated
		Eventually(func() bool {
			if err := k8sClient.Get(ctx, vrKey, vr); err != nil {
				return false
			}
			stateOk := vr.Status.State == v1alpha1.ValidationInProgress
			sinkStateOk := len(vr.Status.Conditions) > 0 && vr.Status.Conditions[0].Type == v1alpha1.SinkEmission && vr.Status.Conditions[0].Reason == string(v1alpha1.SinkEmitNA)
			hashOk := vr.ObjectMeta.Annotations[ValidationResultHash] == vr.Hash()
			return stateOk && sinkStateOk && hashOk
		}, timeout, interval).Should(BeTrue(), "failed to update ValidationResult Status")
	})

	It("Should attempt to emit a message to Slack once the ValidationResult is updated", func() {
		By("By updating the ValidationResult")
		ctx := context.Background()

		vrBytes, err := os.ReadFile(vrServiceQuota)
		Expect(err).NotTo(HaveOccurred())

		statusUpdated := false
		vrUpdated := &v1alpha1.ValidationResult{}
		err = kyaml.Unmarshal(vrBytes, vrUpdated)
		Expect(err).NotTo(HaveOccurred())

		// Wait for the ValidationResult's Status to be updated
		Eventually(func() bool {
			if err := k8sClient.Get(ctx, vrKey, vr); err != nil {
				return false
			}

			if !statusUpdated {
				vr.Status.State = v1alpha1.ValidationSucceeded
				vr.Status.ValidationConditions = vrUpdated.Status.ValidationConditions
				if err := k8sClient.Status().Update(ctx, vr); err != nil {
					return false
				}
				statusUpdated = true
			}

			// expect the sink to fail, as we've not created the slack secret
			sinkStateOk := len(vr.Status.Conditions) > 0 && vr.Status.Conditions[0].Type == v1alpha1.SinkEmission && vr.Status.Conditions[0].Reason == string(v1alpha1.SinkEmitFailed)
			return sinkStateOk
		}, timeout, interval).Should(BeTrue(), "failed to update ValidationResult Status")
	})
})
