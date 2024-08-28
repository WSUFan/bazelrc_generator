load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library (
    name = "bazel_options",
    srcs = [
        "bazel_options.go"
    ],
    importpath = "github.com/WSUFan/bazelrc_generator/pkg/bazel_options",
    visibility = ["//visibility:public"],
)

go_test(
    name = "bazel_options_test",
    srcs = [
        "bazel_options_test.go",
    ],
    deps = [
        ":bazel_options"
    ],
)