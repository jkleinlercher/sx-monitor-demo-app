# sx-monitor-demo-app

This is a simple web application written in Go. It includes several routes that demonstrate different aspects of web development, such as generating high CPU load, using lots of memory, and generating high bandwidth responses.

# Routes

The following routes are available in this application:

- /site/200: Returns a simple 200 OK response.
- /site/503: Returns a 503 Service Unavailable response.
- /site/delay?sleep=10: Returns a response after a configurable delay. The delay can be set using the response query parameter, which should be a duration in seconds.
- /site/high-cpu: Generates high CPU load by adding up a large number of integers in a loop.
- /site/high-memory: Allocates a large amount of memory by creating a slice with 10 million elements.
- /site/high-bandwidth: Generates a large amount of bandwidth by sending a 100 MB binary file as the response.

# Docker

This application can also be run as a Docker container. To build the Docker image, run the following command:

    docker build -t sx-monitor-demo-app:latest .

This will build a Docker image with the application and its dependencies. You can then run the Docker container using the following command:

    docker run -p 8080:8080 sx-monitor-demo-app:latest
    
This will start the container and expose the application on port 8080.

## how to build the docker image and push to ghcr

    export CR_PAT=<your-github-personal-access-token>
    echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
    > Login Succeeded
    docker build -t ghcr.io/jkleinlercher/sx-monitor-demo-app:1.0 .
    docker push ghcr.io/jkleinlercher/sx-monitor-demo-app:1.0
    
# Kubernetes

This application can also be deployed to a Kubernetes cluster using the included Kubernetes manifests. To deploy the application, run the following commands:

    kubectl apply -f deployment.yaml
    kubectl apply -f service.yaml
    kubectl apply -f ingress.yaml

This will create a Kubernetes deployment, service, and ingress for the application. You can then access the application by opening a web browser and navigating to the URL specified in the ingress.

# locust load test

## install locust

    sudo apt install python3-pip
    pip3 install locust

## run locust

    ~/.local/bin/locust -f locust/test_default.py -H http://sx-monitor-demo-app-suxess-it-dev.apps.cluster1.mcp.pitagora.at -r 1 -u 10 --headless

This will run the test with 10 users with a ramp-up of 1 user per second.
There are several other test files in the locust folder.