package job

import (
	"context"
	"errors"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type CreateJobDetails struct {
	JobName       string
	ContainerName string
	NameSpace     string
	labels        map[string]string
	Image         string
	Command       []string
}

type DeleteJobDetails struct {
	JobName   string
	NameSpace string
}

func CreateJob(jobData CreateJobDetails, clientset kubernetes.Interface) (*batchv1.Job, error) {
	jobsClient := clientset.BatchV1().Jobs(jobData.NameSpace)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobData.JobName,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyOnFailure,
					Containers: []corev1.Container{
						{
							Name:    jobData.ContainerName,
							Image:   jobData.Image,
							Command: jobData.Command,
						},
					},
				},
			},
		},
	}

	result, err := jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error creating job: %w", err)
	}
	if result == nil {
		return nil, errors.New("error creating job")
	}
	fmt.Printf("Created job %q.\n", result.GetObjectMeta().GetName())
	return result, nil
}
