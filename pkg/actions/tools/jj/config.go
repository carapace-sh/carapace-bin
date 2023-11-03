package jj

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionConfigs(includeDefaults bool) carapace.Action {
	if !includeDefaults {
		return actionConfigs(false)
	}
	return carapace.Batch(
		actionConfigs(true),
		actionConfigs(false).Style(style.Blue),
	).ToA()
}

func actionConfigs(includeDefaults bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"config", "list"}
		if includeDefaults {
			args = append(args, "--include-defaults")
		}
		return carapace.ActionExecCommand("jj", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, strings.SplitN(line, "=", 2)...)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
