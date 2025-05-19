package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func ActionSecrets(cmd *cobra.Command) carapace.Action {
	return actionExecVercel(cmd, "secrets", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		if len(lines) > 3 {
			for _, line := range lines[3:] {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[0], strings.Join(fields[1:], " "))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
