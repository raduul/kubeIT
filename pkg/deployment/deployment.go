package deployment

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type CreateDeploymentDetails struct {
	DeploymentName string            `json:"deploymentName"`
	Namespace      string            `json:"namespace"`
	Labels         map[string]string `json:"labels"`
	Image          string            `json:"image"`
	Replicas       int32             `json:"replicas"`
	ContainerPort  int32             `json:"containerPort"`
}

func CreateDeployment(ctx context.Context, deploymentData CreateDeploymentDetails, clientset kubernetes.Interface) (*appsv1.Deployment, error) {
	deploymentsClient := clientset.AppsV1().Deployments(deploymentData.Namespace)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   deploymentData.DeploymentName,
			Labels: deploymentData.Labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deploymentData.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: deploymentData.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: deploymentData.Labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentData.DeploymentName,
							Image: deploymentData.Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: deploymentData.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentsClient.Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment: %w", err)
	}

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	return result, nil
}

func DeleteDeployment(ctx context.Context, name, namespace string, clientset kubernetes.Interface) error {
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	return deploymentsClient.Delete(ctx, name, metav1.DeleteOptions{})
}
