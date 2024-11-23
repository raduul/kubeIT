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
	JobName       string            `json:"jobName"`
	ContainerName string            `json:"containerName"`
	NameSpace     string            `json:"namespace"`
	Labels        map[string]string `json:"labels"`
	Image         string            `json:"image"`
	Command       []string          `json:"command"`
}

type DeleteJobDetails struct {
	JobName   string
	NameSpace string
}

func CreateJob(ctx context.Context, jobData CreateJobDetails, clientset kubernetes.Interface) (*batchv1.Job, error) {
	jobsClient := clientset.BatchV1().Jobs(jobData.NameSpace)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:   jobData.JobName,
			Labels: jobData.Labels,
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

	response, err := jobsClient.Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error creating job: %w", err)
	}
	if response == nil {
		return nil, errors.New("error creating job")
	}
	fmt.Printf("Created job %q.\n", response.Name)
	return response, nil
}

func DeleteJob(jobName DeleteJobDetails, clientset kubernetes.Interface) (bool, error) {
	jobsClient := clientset.BatchV1().Jobs(jobName.NameSpace)

	err := jobsClient.Delete(context.TODO(), jobName.JobName, metav1.DeleteOptions{})
	if err != nil {
		return false, fmt.Errorf("error deleting job: %w", err)
	}
	fmt.Printf("Deleted job %q.\n", jobName.JobName)
	return true, nil
}

func AutoRemoveSucceededJobs(clientset kubernetes.Interface) error {
	list, err := clientset.BatchV1().Jobs(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error listing jobs: %w", err)
	}

	for _, job := range list.Items {
		if job.Status.Succeeded > 0 {
			deletePolicy := metav1.DeletePropagationForeground
			err := clientset.BatchV1().Jobs(job.Namespace).Delete(context.TODO(),
				job.Name,
				metav1.DeleteOptions{
					PropagationPolicy: &deletePolicy,
				})

			if err != nil {
				return fmt.Errorf("error deleting job: %s in namespace %s: %w", job.Name, job.Namespace, err)
			}
			fmt.Printf("Deleted following job: %q in namespace: %q.\n", job.Name, job.Namespace)
		}
	}
	return nil
}
