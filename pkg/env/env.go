package env

import (
	"os"
	"strings"
)

const (
	CARAPACE_BUILTINS = "CARAPACE_BUILTINS" // enabled shell builtin groups
	CARAPACE_EXCLUDES = "CARAPACE_EXCLUDES" // excluded completers
)

func Builtins() []string {
	if v, ok := os.LookupEnv(CARAPACE_BUILTINS); ok {
		return strings.Split(v, ",")
	}
	return []string{}
}

func Excludes() []string {
	// TODO: excludes should allow excluding variants/groups (`/zsh`, `@bsd`)
	if v, ok := os.LookupEnv(CARAPACE_EXCLUDES); ok {
		return strings.Split(v, ",")
	}
	return []string{}
}
