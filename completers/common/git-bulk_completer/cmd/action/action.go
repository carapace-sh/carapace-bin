package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionBulkWorkspaces() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("git", "config", "--global", "--get-regexp", "^bulkworkspaces\\.")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}
			vals := make([]string, 0)
			for line := range strings.SplitSeq(string(output), "\n") {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				parts := strings.SplitN(line, " ", 2)
				if len(parts) == 2 && strings.HasPrefix(parts[0], "bulkworkspaces.") {
					name := strings.TrimPrefix(parts[0], "bulkworkspaces.")
					vals = append(vals, name, parts[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...).Tag("workspaces")
		})
	})
}
