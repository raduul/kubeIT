package deployment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateDeployment(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	deploymentDetails := CreateDeploymentDetails{
		DeploymentName: "test-deployment",
		Namespace:      "default",
		Labels:         map[string]string{"app": "test"},
		Image:          "nginx",
		Replicas:       2,
		ContainerPort:  80,
	}

	result, err := CreateDeployment(context.TODO(), deploymentDetails, clientset)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, deploymentDetails.DeploymentName, result.Name)
	assert.Equal(t, deploymentDetails.Namespace, result.Namespace)
	assert.Equal(t, deploymentDetails.Labels, result.Labels)
	assert.Equal(t, *result.Spec.Replicas, deploymentDetails.Replicas)
}

func TestCreateDeployment_ExceptionHandling_ErrorForwarded_DeploymentName(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	deploymentDetails := CreateDeploymentDetails{
		DeploymentName: "",
		Namespace:      "default",
		Labels:         map[string]string{"app": "test"},
		Image:          "nginx",
		Replicas:       1,
		ContainerPort:  80,
	}

	expectedErrorMessage := "error DeploymentName is empty"

	_, err := CreateDeployment(context.TODO(), deploymentDetails, clientset)

	assert.EqualError(t, err, expectedErrorMessage)
}

func TestCreateDeployment_ExceptionHandling_ErrorForwarded_Namespace(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	deploymentDetails := CreateDeploymentDetails{
		DeploymentName: "test-deployment",
		Namespace:      "",
		Labels:         map[string]string{"app": "test"},
		Image:          "nginx",
		Replicas:       1,
		ContainerPort:  80,
	}

	expectedErrorMessage := "error Namespace is empty"

	_, err := CreateDeployment(context.TODO(), deploymentDetails, clientset)

	assert.EqualError(t, err, expectedErrorMessage)
}

func TestCreateDeployment_ExceptionHandling_ErrorForwarded_Image(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	deploymentDetails := CreateDeploymentDetails{
		DeploymentName: "test-deployment",
		Namespace:      "default",
		Labels:         map[string]string{"app": "test"},
		Image:          "",
		Replicas:       1,
		ContainerPort:  80,
	}

	expectedErrorMessage := "error Image is empty"

	_, err := CreateDeployment(context.TODO(), deploymentDetails, clientset)

	assert.EqualError(t, err, expectedErrorMessage)
}
