package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/raduul/kubeIT/pkg/client"
	"github.com/raduul/kubeIT/pkg/handlers"
)

func main() {
	clientConfig, err := client.ClientSetup()
	if err != nil {
		panic(err)
	}

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
