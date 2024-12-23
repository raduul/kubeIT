package job_test

import (
	"context"
	"log"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v7"
	"github.com/raduul/kubeIT/pkg/job"
	"github.com/stretchr/testify/assert"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateJob(t *testing.T) {

	jobDetails := job.CreateJobDetails{
		JobName:       gofakeit.DigitN(10),
		ContainerName: gofakeit.DigitN(10),
		NameSpace:     gofakeit.DigitN(10),
		Image:         "busybox",
		Command:       []string{"echo", "hello there test user"},
	}

	clientset := fake.NewSimpleClientset()

	result, err := job.CreateJob(context.TODO(), jobDetails, clientset)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	createdJob, err := clientset.BatchV1().Jobs(jobDetails.NameSpace).Get(context.TODO(), jobDetails.JobName, metav1.GetOptions{})
	assert.NoError(t, err)
	assert.NotNil(t, createdJob)

	assert.Equal(t, jobDetails.JobName, createdJob.Name)
	assert.Equal(t, jobDetails.NameSpace, createdJob.Namespace)

	container := createdJob.Spec.Template.Spec.Containers[0]
	assert.Equal(t, jobDetails.ContainerName, container.Name)
	assert.Equal(t, jobDetails.Image, container.Image)
	assert.Equal(t, jobDetails.Command, container.Command)
	job.AutoRemoveSucceededJobs(clientset)
}

func TestDeleteJob_Success(t *testing.T) {

	jobName := "test-job"
	namespace := "default"
	deleteJobDetails := job.DeleteJobDetails{
		JobName:   jobName,
		NameSpace: namespace,
	}

	clientset := fake.NewSimpleClientset()

	_, err := clientset.BatchV1().Jobs(namespace).Create(context.TODO(), &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: namespace,
		},
	}, metav1.CreateOptions{})
	assert.NoError(t, err)

	result, err := job.DeleteJob(deleteJobDetails, clientset)

	assert.NoError(t, err)
	assert.True(t, result)

	_, err = clientset.BatchV1().Jobs(namespace).Get(context.TODO(), jobName, metav1.GetOptions{})
	assert.Error(t, err)
	job.AutoRemoveSucceededJobs(clientset)
}

func TestDeleteJob_ForwardsError(t *testing.T) {

	jobName := "test-job"
	namespace := "default"
	deleteJobDetails := job.DeleteJobDetails{
		JobName:   jobName,
		NameSpace: namespace,
	}

	clientset := fake.NewSimpleClientset()

	expectedErrorMessage := "error deleting job: jobs.batch \"test-job\" not found"

	_, err := job.DeleteJob(deleteJobDetails, clientset)

	assert.EqualError(t, err, expectedErrorMessage)

}

func TestAutoDeleteJob_JobIsDeletedThrowsNoError(t *testing.T) {
	jobDetails := job.CreateJobDetails{
		JobName:       gofakeit.DigitN(10),
		ContainerName: gofakeit.DigitN(10),
		NameSpace:     gofakeit.DigitN(10),
		Image:         "busybox",
		Command:       []string{"echo", "hello there test user"},
	}

	clientset := fake.NewSimpleClientset()

	nowCreatedJob, err := job.CreateJob(context.TODO(), jobDetails, clientset)
	log.Println("NOW CREATED---------", nowCreatedJob)
	assert.NoError(t, err)

	createdJob, err := clientset.BatchV1().Jobs(jobDetails.NameSpace).Get(context.TODO(), jobDetails.JobName, metav1.GetOptions{})
	assert.NoError(t, err)
	//fake.NewSimpleClientset() has limitations and doesn't change status of created job, so we need to do it manually
	createdJob.Status.Succeeded = 1

	_, err = clientset.BatchV1().Jobs(jobDetails.NameSpace).Update(context.TODO(), createdJob, metav1.UpdateOptions{})
	assert.NoError(t, err)

	err = job.AutoRemoveSucceededJobs(clientset)
	assert.NoError(t, err)

	_, err = clientset.BatchV1().Jobs(jobDetails.NameSpace).Get(context.TODO(), jobDetails.JobName, metav1.GetOptions{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
