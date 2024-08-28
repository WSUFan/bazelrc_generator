package bazel_options

import (
	"fmt"
	"strings"
)

type BazelOption struct {
	prefix  string
	options []string
}

func NewBazelOption(prefix string) *BazelOption {
	return &BazelOption{
		prefix:  prefix,
		options: []string{},
	}
}

func (b *BazelOption) AddOption(flag string, value string) {
	if value != "" {
		b.options = append(b.options, fmt.Sprintf("%s=%s", flag, value))
	} else {
		b.options = append(b.options, flag)
	}
}

func (b *BazelOption) Generate() []string {
	var result []string
	for _, opt := range b.options {
		if b.prefix != "" {
			result = append(result, fmt.Sprintf("%s --%s", b.prefix, opt))
		} else {
			result = append(result, opt)
		}
	}
	return result
}

type BazelBuildOption struct {
	*BazelOption
}

func NewBazelBuildOption() BazelBuildOption {
	return BazelBuildOption{
		BazelOption: NewBazelOption("build"),
	}
}

type BazelTestOption struct {
	*BazelOption
}

func NewBazelTestOption() BazelTestOption {
	return BazelTestOption{
		BazelOption: NewBazelOption("test"),
	}
}

type BazelQueryOption struct {
	*BazelOption
}

func NewBazelQueryOption() BazelQueryOption {
	return BazelQueryOption{
		BazelOption: NewBazelOption("query"),
	}
}

type BazelRcConfig struct {
	BuildOptions *BazelBuildOption
	TestOptions  *BazelTestOption
	QueryOptions *BazelQueryOption
}

func NewBazelRcConfig(build *BazelBuildOption, test *BazelTestOption, query *BazelQueryOption) *BazelRcConfig {
	return &BazelRcConfig{
		BuildOptions: build,
		TestOptions:  test,
		QueryOptions: query,
	}
}

func (c *BazelRcConfig) GenerateBazelrc() string {
	var result []string

	if c.BuildOptions != nil {
		result = append(result, c.BuildOptions.Generate()...)
	}
	if c.TestOptions != nil {
		result = append(result, c.TestOptions.Generate()...)
	}
	if c.QueryOptions != nil {
		result = append(result, c.QueryOptions.Generate()...)
	}

	return strings.Join(result, "\n")
}
