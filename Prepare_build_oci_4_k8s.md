### 1. Install BuildKit
https://github.com/moby/buildkit/tree/master
#### 1.1 Install **_crun_** dependencies
https://github.com/containers/crun
```shell
# Ubuntu
$ sudo apt-get install -y make git gcc build-essential pkgconf libtool \
   libsystemd-dev libprotobuf-c-dev libcap-dev libseccomp-dev libyajl-dev \
   libgcrypt20-dev go-md2man autoconf python3 automake
```
Build From Source
```shell
./autogen.sh
./configure --enable-shared
make
sudo make install #install into default (/usr/local/bin)
crun --version
```
#### 1.2 Install **_containerd_** dependencies
https://github.com/containerd/containerd
Build from source
```shell
cd containerd
make
sudo make install #install into default (/usr/local/bin)
```
#### 1.3 Install **_buildkit_**
https://github.com/moby/buildkit/blob/master/.github/CONTRIBUTING.md
Build from source
```shell
make && sudo make install # installs buildkitd and buildctl to /usr/local/bin
```
OR move buildkitd and buildctl binary file to /usr/local/bin

### 3. Create Remote BuildKit over Unix Socket
https://docs.docker.com/build/drivers/remote/
You should use **_unix://$HOME/buildkitd.sock_** for accessing without sudo permission 
```shell
sudo buildkitd --group $(id -gn) --addr unix://$HOME/buildkitd.sock
ls -lh $HOME/buildkitd.sock
sudo buildctl debug info
docker buildx create --name remote-unix --driver remote unix://$HOME/buildkitd.sock
```
Check status `RUNNING` of `remote-unix`
```shell
docker buildx ls
```

### 4. Build oci
For example
```shell
docker buildx --builder=remote-unix build \
 --output "type=oci,dest=/tmp/thuyhome-oci-v1.tar" \
-f dockerfiles/base/Dockerfile .
```

### 5. Push and Pull
```shell
oras push docker.io/thuyltm2201/thuyhome:oc_v1 /tmp/thuyhome-oci-v1.tar
oras manifest fetch docker.io/thuyltm2201/thuyhome:oc_v1 | jq .
oras pull docker.io/thuyltm2201/thuyhome:oc_v1 -T
```

### 6. Test
```shell
docker load -i /tmp/thuyhome-oci-v1.tar
docker run <<IMAGE_ID>>
```