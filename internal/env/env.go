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
		return []string{"zsh", "fish"}
	}
	return strings.Split(v, ",")
}
