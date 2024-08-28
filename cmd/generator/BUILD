load("@rules_go//go:def.bzl", "go_binary", "go_library")
go_library(
    name = "bazelrc_generator_lib",
    srcs = ["main.go"],
    importpath = "github.com/WSUFan/bazelrc_generator/cmd/generator",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/bazel_options"
    ],
)

go_binary(
    name = "bazelrc_generator",
    embed = [":bazelrc_generator_lib"],
    pure = "on",
    visibility = ["//visibility:public"],
)