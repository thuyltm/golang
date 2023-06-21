load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
package(default_visibility = ["//visibility:public"])

go_binary(
    name = "playvideo",
    embed = [":playvideo_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "playvideo_lib",
    srcs = ["main.go"],
    importpath = "golang-bazel-demo-app/playvideo",
    visibility = ["//visibility:public"],
    deps = [
        "@github_com_zergon321_reisen//:go_default_library"
    ],
)
