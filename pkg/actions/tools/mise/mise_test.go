package mise

import (
	"reflect"
	"testing"
)

func TestParseDescribed(t *testing.T) {
	output := []byte("build:Build it\ntest\\:unit:\nlint\n\nfmt:Format\\: all files\r\n")
	expected := []string{
		"build", "Build it",
		"test:unit", "",
		"lint", "",
		"fmt", "Format: all files",
	}
	if actual := parseDescribed(output); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestParseLines(t *testing.T) {
	output := []byte("bun\n  deno \n\nnode\n")
	expected := []string{"bun", "deno", "node"}
	if actual := parseLines(output); !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
