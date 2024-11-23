package main

import (
	"net/http"
	"time"

	"fmt"
	"log"

	"github.com/raduul/kubeIT/pkg/client"
	"github.com/raduul/kubeIT/pkg/handlers"
	"github.com/raduul/kubeIT/pkg/job"
)

func main() {
	clientConfig, err := client.ClientSetup()
	if err != nil {
		panic(err)
	}

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := job.AutoRemoveSucceededJobs(clientConfig)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Called AutoRemoveSucceededJobs")
			}
		}
	}()

	http.HandleFunc("/createJob", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateJobHandler(w, r, clientConfig)
	})

	http.HandleFunc("/createDeployment", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateDeploymentHandler(w, r, clientConfig)
	})

	http.HandleFunc("/createService", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateServiceHandler(w, r, clientConfig)
	})

	fmt.Println("Starting server on :7080")
	log.Fatal(http.ListenAndServe(":7080", nil))
}
