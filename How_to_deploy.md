### 1.  Push image to Artifact Registry
#### Use Docker Hub
```sh
docker login -u thuyltm2201
Password: Dockerhub@123
```

```sh
docker push YOUR-USER-NAME/getting-started
```

#### Or use Google Artifact Registry
```sh
# enable Google Artifact Registry
gcloud artifacts repositories create hello-repo --repository-format=docker --location=asia-southeast2   --description="Docker repository"
# authenticate to a repository
gcloud auth configure-docker asia-southeast2-docker.pkg.dev
# create 
docker build -t asia-southeast2-docker.pkg.dev/${PROJECT_ID}/${REPOSITORY_NAME}/hello-app:v1
```

```shell
dive bazel/playvideo/analyze:analyze-image
```

### 2. minikube, Skaffold, Helm Chart deploy to GKE

```shell
minikube update-check
sudo minikube delete             # remove your minikube cluster
sudo rm -rf ~/.minikube          # remove minikube
```

### 3. Setup Envoy, Ambassador Gateway, Linkerd Server Mesh,

### 4. Promethesus, Grafana, VictoriaMetric

### 5. CI/CD Jenkins, Spinaker, Argocd