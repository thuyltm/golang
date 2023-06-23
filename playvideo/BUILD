load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
package(default_visibility = ["//visibility:public"])

go_binary(
    name = "analyze",
    embed = [":analyze_lib"],
    visibility = ["//visibility:public"],
    cdeps = [
        "@system_libs//:libffmeg"],
    cgo = True
)

go_library(
    name = "analyze_lib",
    data = [
        "demo.mkv",
        "demo.mp4"
    ],
    srcs = ["main.go"],
    importpath = "golang-bazel-demo-app/playvideo",
    visibility = ["//visibility:public"],
    deps = [
        "@github_com_zergon321_reisen//:go_default_library"
    ]
)

