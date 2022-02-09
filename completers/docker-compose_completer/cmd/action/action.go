package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func actionExecCompose(cmd *cobra.Command, arg ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			a := []string{}
			if files, err := cmd.Root().Flags().GetStringArray("file"); err == nil {
				for _, file := range files {
					a = append(a, "--file", file)
				}
			} else {
				return carapace.ActionMessage(err.Error())
			}
			a = append(a, arg...)
			return carapace.ActionExecCommand("docker-compose", a...)(f)
		})
	}
}
