package os

import (
	"github.com/rsteube/carapace"
	"os"
	"strings"
)

func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		env := os.Environ()
		vars := make([]string, len(env))
		for index, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			vars[index] = pair[0]
		}
		return carapace.ActionValues(vars...)
	})
}
