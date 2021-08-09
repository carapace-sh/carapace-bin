package action

import (
	"github.com/rsteube/carapace"
	"strings"
)

func ActionLanguages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("bat", "--list-languages")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			values := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				splitted := strings.Split(line, ":")
				if len(splitted) == 1 {
					values = append(values, splitted[0], "")
				} else if len(splitted) > 1 {
					values = append(values, splitted[0], splitted[1])
				}
			}
			return carapace.ActionValuesDescribed(values...)
		})
	})
}

func ActionThemes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("bat", "--list-themes")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
