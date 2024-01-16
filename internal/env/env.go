package env

import (
	"os"
	"strings"
)

const (
	CARAPACE_BRIDGES = "CARAPACE_BRIDGES" // order of implicit bridges
)

func Bridges() []string {
	v, ok := os.LookupEnv(CARAPACE_BRIDGES)
	if !ok {
		return []string{"fish", "zsh"}
	}
	return strings.Split(v, ",")
}
