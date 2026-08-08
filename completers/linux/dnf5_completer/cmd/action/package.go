package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func actionPackages(cmd *cobra.Command, listFlag string, requireInput bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if requireInput && len(c.Value) == 0 {
			return carapace.ActionMessage("package search needs at least one character")
		}

		args := []string{"--quiet", "--cacheonly"}
		for _, name := range []string{"repo", "setopt", "installroot"} {
			if f := cmd.Root().Flag(name); f.Changed {
				args = append(args, "--"+f.Name, f.Value.String())
			}
		}
		args = append(args, "repoquery", "--queryformat", "%{name}.%{arch}\t%{summary}\n", listFlag, c.Value+"*")

		return carapace.ActionExecCommand("dnf5", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines {
				if fields := strings.SplitN(line, "\t", 2); len(fields) == 2 {
					vals = append(vals, fields[0], fields[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionPackageSearch(cmd *cobra.Command) carapace.Action {
	return actionPackages(cmd, "--available", true)
}

func ActionInstalledPackages(cmd *cobra.Command) carapace.Action {
	return actionPackages(cmd, "--installed", false)
}
