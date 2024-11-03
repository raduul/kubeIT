package main

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Job struct {
	JobName       string
	ContainerName string
	NameSpace     string
	labels        map[string]string
	Image         string
	Command       []string
}

func clientSetup() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", "/Users/radesrdanovic/.kube/config")
		if err != nil {
			panic(err)
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset, err
}

func createJob(jobData Job, clientset *kubernetes.Clientset) {
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
		fmt.Printf("Error creating job: %v\n", err)
		panic(err)
	}
	fmt.Printf("Created job %q.\n", result.GetObjectMeta().GetName())
}

func main() {
	clientConfig, err := clientSetup()
	if err != nil {
		fmt.Println("Failed to setup client")
		panic(err)
	}

	//TODO: Call function which creates the job
	jobDetails := Job{
		JobName:       "job1",
		ContainerName: "container1",
		NameSpace:     "default",
		Image:         "busybox",
		Command:       []string{"echo", "hello there test user"},
	}
	createJob(jobDetails, clientConfig)
}
