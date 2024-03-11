package fastfetch

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionModules completes modules
//
//	Board (Print mather board name and other info)
//	Break (Print a empty line)
func ActionModules() carapace.Action {
	return carapace.ActionExecCommand("fastfetch", "--list-modules", "autocompletion")(func(output []byte) carapace.Action {
		var texts []string
		for _, line := range strings.Split(strings.TrimRight(string(output), "\n"), "\n") {
			name, desc, _ := strings.Cut(line, ":")
			texts = append(texts, name, desc)
		}
		return carapace.ActionValuesDescribed(texts...)
	})
}
