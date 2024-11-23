# kubeIT

kubeIT is a Kubernetes management tool designed to simplify the deployment, scaling, and management of containerized applications through REST API calls.

The kubeIT approach allows multiple microservices to internally spin up additional resources, for example if an additional jupyter kernel is needed to run python notebooks, this can be invoked directly through here. Moreover, same can be applied to jobs, if there is a need to run a data transformation, it can be executed as a docker container inside of a job, and using a storage class you can specify where the data should be ingested and transformed.

## Features

- Easy creation of deployments
- Easy creation of jobs

## Installation

To install kubeIT, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/kubeIT.git
    ```
2. Navigate to the project directory:
    ```sh
    cd kubeIT
    ```
3. Create an image:
    ```sh
    docker build -t kubeit:2.0.13 .
    ```
4. Deploy on kubernetes:
    ```sh
    kubectl apply -f kubernetes/
    ```
5. Deploy jobs/deployments on kubernetes:
    ```sh
    curl -X POST http://localhost:7080/createJob \
    -H "Content-Type: application/json" \
    -d '{
        "jobName": "2example-job-on",
        "containerName": "example-container",
        "namespace": "default",
        "image": "busybox",
        "command": ["echo", "Hello World"]
    }'
    ```
    ```sh
    curl -v -X POST http://localhost:7080/createDeployment \
    -H "Content-Type: application/json" \
    -d "{\"deploymentName\": \"example-deployment\", \"namespace\": \"default\", \"labels\": {\"app\": \"example\"}, \"image\": \"nginx\", \"replicas\": 1, \"containerPort\": 80}"

    ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, please contact me at rsrdan@proton.me.
