package handlers

import (
	"net/http"

	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/raduul/kubeIT/pkg/deployment"
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

func CreateDeploymentHandler(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
	log.Println("Received request at /createDeployment")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var deploymentDetails deployment.CreateDeploymentDetails
	if err := json.NewDecoder(r.Body).Decode(&deploymentDetails); err != nil {
		log.Printf("Invalid JSON format: %v", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if deploymentDetails.DeploymentName == "" || deploymentDetails.Namespace == "" || deploymentDetails.Image == "" {
		log.Println("Missing required deployment details")
		http.Error(w, "Missing required deployment details", http.StatusBadRequest)
		return
	}

	createdDeployment, err := deployment.CreateDeployment(r.Context(), deploymentDetails, clientset)
	if err != nil {
		log.Printf("Failed to create deployment: %v", err)
		http.Error(w, "Failed to create deployment", http.StatusInternalServerError)
		return
	}

	log.Printf("Deployment %s created successfully", createdDeployment.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdDeployment)
}

func CreateServiceHandler(w http.ResponseWriter, r *http.Request, clientset kubernetes.Interface) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var serviceDetails deployment.CreateServiceDetails
	if err := json.NewDecoder(r.Body).Decode(&serviceDetails); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if serviceDetails.ServiceName == "" || serviceDetails.Namespace == "" || serviceDetails.Port == 0 {
		http.Error(w, "Missing required service details", http.StatusBadRequest)
		return
	}

	createdService, err := deployment.CreateService(r.Context(), serviceDetails, clientset)
	if err != nil {
		fmt.Printf("Failed to create service: %v\n", err)
		http.Error(w, "Failed to create service", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdService)
}
