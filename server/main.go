package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/raduul/kubeIT/pkg/client"
	handler "github.com/raduul/kubeIT/pkg/handlers"
)

func main() {
	clientConfig, err := client.ClientSetup()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/createJob", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateJobHandler(w, r, clientConfig)
	})

	fmt.Println("Starting server on :7080")
	if err := http.ListenAndServe(":7080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
