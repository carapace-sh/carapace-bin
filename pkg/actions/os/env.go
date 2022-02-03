package os

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionEnvironmentVariables completes environment values
//   SHELL (/bin/elvish)
//   LANG (en_US.utf8)
func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		env := os.Environ()
		vals := make([]string, 0, len(env)*2)
		for _, e := range os.Environ() {
			vals = append(vals, strings.SplitN(e, "=", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
