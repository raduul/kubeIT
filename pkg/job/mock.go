package job

import (
	"github.com/stretchr/testify/mock"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes"
)

type MockJobService struct {
	mock.Mock
}

func (m *MockJobService) CreateJob(job CreateJobDetails, clientset kubernetes.Interface) (*batchv1.Job, error) {
	args := m.Called(job, clientset)
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

func (m *MockJobService) DeleteJob(jobName DeleteJobDetails, clientset kubernetes.Interface) (bool, error) {
	args := m.Called(jobName, clientset)
	return args.Bool(0), args.Error(1)
}
