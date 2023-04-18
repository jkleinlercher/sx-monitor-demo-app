# sx-monitor-demo-app

## how to build the docker image and push to ghcr

    export CR_PAT=<your-github-personal-access-token>
    echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
    > Login Succeeded
    docker build -t ghcr.io/jkleinlercher/sx-monitor-demo-app:1.0 .
    docker push ghcr.io/jkleinlercher/sx-monitor-demo-app:1.0