package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func ActionActivationVariables(cmd *cobra.Command) carapace.Action {
	return actionWorkspaceEntries(cmd, []string{"workspace", "activation", "env", "list"}, false)
}

func ActionWorkspaceDependencies(cmd *cobra.Command) carapace.Action {
	return actionWorkspaceEntries(cmd, []string{"workspace", "dependencies", "list"}, true)
}

func actionWorkspaceEntries(cmd *cobra.Command, args []string, dependencies bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		commandArgs := append([]string{}, args...)
		if !dependencies && !cmd.Flag("feature").Changed && !cmd.Flag("environment").Changed {
			commandArgs = append(commandArgs, "--feature", "default")
		}
		for _, name := range []string{"manifest-path", "workspace", "environment", "feature", "target"} {
			if flag := cmd.Flag(name); flag != nil && flag.Changed {
				commandArgs = append(commandArgs, "--"+name, flag.Value.String())
			}
		}
		return carapace.ActionExecCommand("pixi", commandArgs...)(func(output []byte) carapace.Action {
			values := make([]string, 0)
			for line := range strings.SplitSeq(string(output), "\n") {
				line = strings.TrimSpace(line)
				if dependencies {
					var ok bool
					if line, ok = strings.CutPrefix(line, "- "); !ok {
						continue
					}
				}
				if name, value, ok := strings.Cut(line, " = "); ok {
					values = append(values, name, value)
				}
			}
			return carapace.ActionValuesDescribed(values...)
		})
	})
}
