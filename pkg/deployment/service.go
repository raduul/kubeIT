package deployment

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

type CreateServiceDetails struct {
	ServiceName string             `json:"serviceName"`
	Namespace   string             `json:"namespace"`
	Labels      map[string]string  `json:"labels"`
	Port        int32              `json:"port"`
	TargetPort  int32              `json:"targetPort"`
	Selector    map[string]string  `json:"selector"`
	Type        corev1.ServiceType `json:"type"` // ClusterIP, NodePort, LoadBalancer
}

func CreateService(ctx context.Context, serviceData CreateServiceDetails, clientset kubernetes.Interface) (*corev1.Service, error) {
	servicesClient := clientset.CoreV1().Services(serviceData.Namespace)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:   serviceData.ServiceName,
			Labels: serviceData.Labels,
		},
		Spec: corev1.ServiceSpec{
			Type: serviceData.Type,
			Ports: []corev1.ServicePort{
				{
					Port:       serviceData.Port,
					TargetPort: intstr.FromInt(int(serviceData.TargetPort)),
				},
			},
			Selector: serviceData.Selector,
		},
	}

	result, err := servicesClient.Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
	return result, nil
}

func DeleteService(ctx context.Context, name, namespace string, clientset kubernetes.Interface) error {
	servicesClient := clientset.CoreV1().Services(namespace)
	return servicesClient.Delete(ctx, name, metav1.DeleteOptions{})
}
