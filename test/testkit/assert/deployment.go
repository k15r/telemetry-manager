package assert

import (
	"context"
	"fmt"

	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kyma-project/telemetry-manager/test/testkit/periodic"
)

func DeploymentReady(ctx context.Context, k8sClient client.Client, name types.NamespacedName) {
	Eventually(func(g Gomega) {
		ready, err := isDeploymentReady(ctx, k8sClient, name)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(ready).To(BeTrueBecause("Deployment not ready"))
	}, periodic.EventuallyTimeout, periodic.DefaultInterval).Should(Succeed())
}

func isDeploymentReady(ctx context.Context, k8sClient client.Client, name types.NamespacedName) (bool, error) {
	var deployment appsv1.Deployment

	err := k8sClient.Get(ctx, name, &deployment)
	if err != nil {
		return false, fmt.Errorf("failed to get deployment: %w", err)
	}

	listOptions := client.ListOptions{
		LabelSelector: labels.SelectorFromSet(deployment.Spec.Selector.MatchLabels),
		Namespace:     name.Namespace,
	}

	return arePodsReady(ctx, k8sClient, listOptions)
}

func DeploymentHasPriorityClass(ctx context.Context, k8sClient client.Client, name types.NamespacedName, expectedPriorityClassName string) {
	Eventually(func(g Gomega) {
		var deployment appsv1.Deployment
		g.Expect(k8sClient.Get(ctx, name, &deployment)).To(Succeed())

		g.Expect(deployment.Spec.Template.Spec.PriorityClassName).To(Equal(expectedPriorityClassName))
	}, periodic.EventuallyTimeout, periodic.DefaultInterval).Should(Succeed())
}
