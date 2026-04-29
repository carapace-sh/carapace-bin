package actions

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionPackages(installedOnly bool) carapace.Action {
	args := []string{"list-available"}
	if installedOnly {
		args = []string{"list-installed"}
	}

	return carapace.ActionExecCommand("eopkg", args...)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		values := make([]string, 0, len(lines)*2)
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) == 0 {
				continue
			}
			name := fields[0]
			if strings.HasPrefix(name, "- ") || strings.HasPrefix(name, "Package") {
				continue
			}
			desc := ""
			if index := strings.Index(line, " - "); index != -1 {
				desc = strings.TrimSpace(line[index+3:])
			}
			values = append(values, name, desc)
		}
		return carapace.ActionValuesDescribed(values...)
	})
}
