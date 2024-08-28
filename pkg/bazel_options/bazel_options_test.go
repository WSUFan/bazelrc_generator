package bazel_options

import (
	"reflect"
	"testing"

	"github.com/WSUFan/bazelrc_generator/pkg/bazel_options"
)

func TestAddOptionWithValue(t *testing.T) {
	option := bazel_options.NewBazelOption("build")
	option.AddOption("option1", "value1")

	expected := []string{"build --option1=value1"}
	result := option.Generate()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAddOptionWithoutValue(t *testing.T) {
	option := bazel_options.NewBazelOption("test")
	option.AddOption("option2", "")

	expected := []string{"test --option2"}
	result := option.Generate()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMultipleOptions(t *testing.T) {
	option := bazel_options.NewBazelOption("query")
	option.AddOption("output", "package")
	option.AddOption("keep_going", "")

	expected := []string{"query --output=package", "query --keep_going"}
	result := option.Generate()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNoPrefixOption(t *testing.T) {
	option := bazel_options.NewBazelOption("")
	option.AddOption("option3", "value3")

	expected := []string{"option3=value3"}
	result := option.Generate()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
