#!/bin/bash

bazel build //hello:hello_oci_tarball
docker load -i $(bazel cquery --output=files //hello:hello_oci_tarball)
