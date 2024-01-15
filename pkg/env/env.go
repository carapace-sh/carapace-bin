package env

import (
	"os"
	"strings"
)

const (
	CARAPACE_EXCLUDES = "CARAPACE_EXCLUDES" // excluded internal completers
)

func Excludes() []string {
	if v, ok := os.LookupEnv(CARAPACE_EXCLUDES); ok {
		return strings.Split(v, ",")
	}
	return []string{}
}
