package job

import (
	"errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateJob(t *testing.T) {
	jobDetails := CreateJobDetails{
		JobName:       gofakeit.DigitN(10),
		ContainerName: gofakeit.DigitN(10),
		NameSpace:     gofakeit.DigitN(10),
		Image:         "busybox",
		Command:       []string{"echo", "hello there test user"},
	}
	clientset := fake.NewSimpleClientset()
	mockJobService := new(MockJobService)

	expectedResult := &batchv1.Job{}

	mockJobService.On("CreateJob", mock.Anything, mock.Anything).Return(expectedResult, nil)
	result, err := mockJobService.CreateJob(jobDetails, clientset)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	mockJobService.AssertCalled(t, "CreateJob", jobDetails, clientset)

	//t.Logf("Mocked Created Job Name: %s", result.Name)
}

func TestCreateJobWithCorrectArgsReturnsNotNil(t *testing.T) {
	jobDetails := CreateJobDetails{
		JobName:       gofakeit.DigitN(1),
		ContainerName: gofakeit.Sentence(1),
		NameSpace:     gofakeit.DigitN(10),
		Image:         gofakeit.Sentence(1),
		Command:       []string{"echo", "hello there test user"},
	}
	clientset := fake.NewSimpleClientset()
	mockJobService := new(MockJobService)

	expectedResult := &batchv1.Job{}

	mockJobService.On("CreateJob", jobDetails, clientset).Return(expectedResult, nil)
	result, err := mockJobService.CreateJob(jobDetails, clientset)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockJobService.AssertCalled(t, "CreateJob", jobDetails, clientset)

	//t.Logf("Mock Created with the following fake Job Name: %s", result.Name)
}

func TestCreateJobWithCorrectArgsReturnsCorrectName(t *testing.T) {
	jobDetails := CreateJobDetails{
		JobName:       gofakeit.DigitN(10),
		ContainerName: gofakeit.DigitN(10),
		NameSpace:     gofakeit.DigitN(10),
		Image:         "busybox",
		Command:       []string{"echo", "hello there test user"},
	}
	clientset := fake.NewSimpleClientset()
	mockJobService := new(MockJobService)

	expectedResult := &batchv1.Job{}

	mockJobService.On("CreateJob", jobDetails, clientset).Return(expectedResult, nil)
	result, err := mockJobService.CreateJob(jobDetails, clientset)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResult.Name, result.Name)
	mockJobService.AssertCalled(t, "CreateJob", jobDetails, clientset)

	//t.Logf("Mock Created with the following fake Job Name: %s", result.Name)
}

func TestCreateJobWithCorrectArgsReturnsCorrectNamespace(t *testing.T) {
	jobDetails := CreateJobDetails{
		JobName:       gofakeit.DigitN(10),
		ContainerName: gofakeit.DigitN(10),
		NameSpace:     gofakeit.DigitN(10),
		Image:         "busybox",
		Command:       []string{"echo", "hello there test user"},
	}
	clientset := fake.NewSimpleClientset()
	mockJobService := new(MockJobService)

	expectedResult := &batchv1.Job{}

	mockJobService.On("CreateJob", jobDetails, clientset).Return(expectedResult, nil)
	result, err := mockJobService.CreateJob(jobDetails, clientset)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResult.Namespace, result.Namespace)
	mockJobService.AssertCalled(t, "CreateJob", jobDetails, clientset)

	//t.Logf("Mock Created with the following fake Job Name: %s", result.Name)
}

func TestCreateJobWithCorrectArgsForwardsError(t *testing.T) {
	jobDetails := CreateJobDetails{
		JobName:       gofakeit.DigitN(10),
		ContainerName: gofakeit.DigitN(10),
		NameSpace:     gofakeit.DigitN(10),
		Image:         "busybox",
		Command:       []string{"echo", "hello there test user"},
	}
	clientset := fake.NewSimpleClientset()
	mockJobService := new(MockJobService)

	expectedErrorMessage := gofakeit.Sentence(5)

	mockJobService.On("CreateJob", jobDetails, clientset).Return((*batchv1.Job)(nil), errors.New(expectedErrorMessage))
	_, err := mockJobService.CreateJob(jobDetails, clientset)

	// Assertions
	assert.EqualError(t, err, expectedErrorMessage)
	mockJobService.AssertCalled(t, "CreateJob", jobDetails, clientset)

	t.Logf("ExpectedErrorMessage: %s, ReturnedErrorMessage: %s", expectedErrorMessage, err.Error())
}
