package just

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionRecipes completes recipes
// default
// build (build project)
func ActionRecipes(justfile string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--list"}
		if justfile != "" {
			args = append(args, "--justfile", justfile)
		}

		return carapace.ActionExecCommand("just", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[1 : len(lines)-1] {
				if splitted := strings.SplitN(line, "#", 2); len(splitted) == 2 {
					vals = append(vals, strings.TrimSpace(splitted[0]), strings.TrimSpace(splitted[1]))
				} else {
					vals = append(vals, strings.TrimSpace(splitted[0]), "")
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
