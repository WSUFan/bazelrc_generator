package main

import (
	"bufio"
	"fmt"
	"github.com/WSUFan/bazelrc_generator/pkg/bazel_options"
	"os"
)

func main() {
	buildOptions := &bazel_options.BazelBuildOption{
		BazelOption: bazel_options.NewBazelOption("build"),
	}
	buildOptionsMap := map[string]string {
		"compilation_mode": "opt",
		"subcommand": "",
		"remote_cache": "grpc//:1234",
	}
	for flag, value := range(buildOptionsMap) {
		buildOptions.AddOption(flag, value)
	}

	testOptions := &bazel_options.BazelTestOption{
		BazelOption: bazel_options.NewBazelOption("test"),
	}
	testOptions.AddOption("test_output", "errors")
	testOptions.AddOption("cache_test_results", "")

	queryOptions := &bazel_options.BazelQueryOption{
		BazelOption: bazel_options.NewBazelOption("query"),
	}
	queryOptions.AddOption("output", "package")
	queryOptions.AddOption("keep_going", "")

	config := bazel_options.NewBazelRcConfig(buildOptions, testOptions, queryOptions)

	bazelrcContent := config.GenerateBazelrc()

	fmt.Println(bazelrcContent)

	f, err := os.Create("/tmp/bazelrc")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(bazelrcContent + "\n")
	if err != nil {
		panic(err)
	}

	w.Flush()
}
