package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func ActionPackageSearch(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Value) == 0 {
			return carapace.ActionMessage("package search needs at least one character") // it works but would overwhelm the completion menus
		}

		args := []string{"--quiet", "--cacheonly"}
		for _, name := range []string{"repo", "setopt", "installroot"} {
			if f := cmd.Root().Flag(name); f.Changed {
				args = append(args, "--"+f.Name, f.Value.String())
			}
		}
		args = append(args, "list", "--available", c.Value+"*")

		return carapace.ActionExecCommand("dnf", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines {
				if fields := strings.Fields(line); len(fields) == 3 {
					vals = append(vals, fields[0], fields[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
