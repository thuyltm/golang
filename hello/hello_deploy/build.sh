#!/bin/bash

bazel run //hello:hello-image-publish
docker save docker.io/thuyltm2201/hello-server:latest > test.tar
docker load -i test.tar
