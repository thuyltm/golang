### Run update golang external dependencies

Import repositories from go.mod and update macro

```
go get github.com/gorilla/websocket
go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

OR in WORKSPACE.bazel
```

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

gazelle_dependencies(go_repository_default_config = "//:WORKSPACE.bazel")

load("//:repositories.bzl", "go_repositories")

go_repository(
    name = "github_com_zergon321_reisen",
    importpath = "github.com/zergon321/reisen", # from go.sum data
    sum = "h1:rflBVtSoc4xXR3tXN+b2Hvk/Fb3qaJhgj+CJ/YWOYmI=", # from go.sum data
    version = "v0.1.8" # from go.sum data
)

# gazelle:repository_macro repositories.bzl%go_repositories
go_repositories()
```