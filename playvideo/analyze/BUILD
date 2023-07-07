load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:image.bzl", "container_image")
package(default_visibility = ["//visibility:public"])

# bazel run //playvideo:analyze -- --file /home/thuy/Documents/Learn/golang/playvideo/src/demo.mp4
# docker run bazel/playvideo/analyze:analyze-image --file=/home/thuy/Documents/Learn/golang/playvideo/src/demo.mp4
go_binary(
    name = "analyze",
    embed = [":analyze_lib"],
    data = ["golang-bazel-demo-app/playvideo/src/demo.mkv"],
    visibility = ["//visibility:public"]
)

go_image(
    name = "analyze-image",
    embed = [":analyze_lib"],
    base = "@thuyhome_base//image",
    visibility = ["//visibility:public"]
)

# Error: docker: Error response from daemon: No command specified
container_image(
    name = "analyze-container-error",
    base = "@go_image_base//image",
    tars = [":analyze-image"],
)

container_image(
    name = "analyze-container",
    base = ":analyze-image"
)

go_library(
    name = "analyze_lib",
    srcs = ["main.go"],
    data = ["//playvideo/src:video_test"],
    importpath = "golang-bazel-demo-app/playvideo/analyze",
    visibility = ["//visibility:public"],
    deps = [
        "@github_com_zergon321_reisen//:go_default_library"
    ]
)

