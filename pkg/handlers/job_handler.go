package handlers

import (
	"net/http"

	"encoding/json"
	"fmt"
	"io"

	"github.com/raduul/kubeIT/pkg/job"
	"k8s.io/client-go/kubernetes"
)

func CreateJobHandler(w http.ResponseWriter, r *http.Request, clientConfig *kubernetes.Clientset) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var jobDetails job.CreateJobDetails
	if err := json.Unmarshal(body, &jobDetails); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	createdJob, err := job.CreateJob(r.Context(), jobDetails, clientConfig)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create job: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdJob)
}
