package main

import (
	"github.com/raduul/kubeIT/pkg/client"
	"github.com/raduul/kubeIT/pkg/job"
)

func main() {
	clientConfig, err := client.ClientSetup()
	if err != nil {
		panic(err)
	}

	// jobDetails := job.CreateJobDetails{
	// 	JobName:       "job1",
	// 	ContainerName: "container1",
	// 	NameSpace:     "default",
	// 	Image:         "busybox",
	// 	Command:       []string{"echo", "hello there test user"},
	// }
	//job.CreateJob(jobDetails, clientConfig)

	deleteJobDetails := job.DeleteJobDetails{
		JobName:   "job1",
		NameSpace: "default",
	}
	job.DeleteJob(deleteJobDetails, clientConfig)

}
