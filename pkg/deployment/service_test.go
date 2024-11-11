package deployment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateService(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	serviceDetails := CreateServiceDetails{
		ServiceName: "test-service",
		Namespace:   "default",
		Labels:      map[string]string{"app": "test"},
		Port:        80,
		TargetPort:  8080,
		Selector:    map[string]string{"app": "test"},
		Type:        corev1.ServiceTypeClusterIP,
	}

	result, err := CreateService(context.TODO(), serviceDetails, clientset)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, serviceDetails.ServiceName, result.Name)
	assert.Equal(t, serviceDetails.Namespace, result.Namespace)
	assert.Equal(t, serviceDetails.Labels, result.Labels)
	assert.Equal(t, serviceDetails.Selector, result.Spec.Selector)
}
