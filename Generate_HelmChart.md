### 1.  imagePullSecrets from docker hub
```shell
k create secret generic regcred --from-file=.dockerconfigjson=/home/thuy/.docker/config.json --type=kubernetes.io/dockerconfigjson
```
Generate an auto-scale yaml file according to kubernetes version 
```shell
kubectl autoscale deployment hello-server-deployment --cpu-percent=50 --min=1 --max=10 -n hello
```
### 2. Deploy by helm
```shell
helm install hello-server hello/hello_deploy --values hello/hello_deploy/values.yaml -n hello
```

```shell
helm upgrade hello-server hello/hello_deploy --values hello/hello_deploy/values.yaml -n hello
```

### 3. Setup Ingress Controller
Like Nginx, Ambassador

#### 3.1 Enable the Ingress Controler Minikube
https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/
```
minikube addons enable ingress
curl --resolve "hello-server-service.hello:80:$( minikube ip )" -i http://hello-server-service.hello/hello-server
```
#### 3.2 Setup fake host `minikube.myhome` which point to $(minikube ip)
```sh
sudo vim /etc/hosts
#192.168.49.2    minikube.myhome
ping minikube.myhome
```

#### 3.4. Add flag `--enable_bzlmod` to ~/.bazelrc
```sh
common --enable_bzlmod
```