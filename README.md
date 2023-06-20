### Run update golang external dependencies

Import repositories from go.mod and update macro

```
go get github.com/gorilla/websocket
go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```
