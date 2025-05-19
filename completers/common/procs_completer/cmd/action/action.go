package action

import (
	"regexp"

	"github.com/carapace-sh/carapace"
)

func ActionColumns() carapace.Action {
	return carapace.ActionExecCommand("procs", "--list")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`(?m)^\s*(\w+)\s*:\s*(.+)$`)
		matches := re.FindAllSubmatch(output, -1)

		var vals []string
		for _, match := range matches {
			vals = append(vals, string(match[1]), string(match[2]))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
