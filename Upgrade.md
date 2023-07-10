### 1. Upgrade docker server

https://askubuntu.com/questions/1072759/how-can-i-update-my-docker-version-in-linux-environment
```shell
sudo apt-get update

dpkg -l | grep docker (Check installed packages in the system)

sudo apt-cache madison docker-ce (Lists all the available docker versions)

sudo apt-get install docker-ce=<VERSION>
sudo apt-get install docker-ce-cli=<VERSION>

sudo service docker restart

docker --version
docker version
```


### 3. Upgrade kubectl client

```shell
curl -LO https://storage.googleapis.com/kubernetes-release/release/<specific-kubectl-version>/bin/darwin/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl $(which kubectl)
kubectl version --client --output=yaml
```