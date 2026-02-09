package env

import (
	"os"
	"strings"
)

const (
	CARAPACE_EXCLUDES = "CARAPACE_EXCLUDES" // excluded completers
)

func Excludes() []string {
	// TODO: excludes should allow excluding variants/groups (`/zsh`, `@bsd`)
	if v, ok := os.LookupEnv(CARAPACE_EXCLUDES); ok {
		return strings.Split(v, ",")
	}
	return []string{}
}
