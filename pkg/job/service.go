package job

import (
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes"
)

type JobService interface {
	CreateJob(jobData CreateJobDetails, clientset kubernetes.Interface) (*batchv1.Job, error)
	DeleteJob(jobName DeleteJobDetails, clientset kubernetes.Interface) (bool, error)
}
