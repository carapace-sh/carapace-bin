package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unsetEnvironmentCmd = &cobra.Command{
	Use:   "unset-environment",
	Short: "Unset one or more environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsetEnvironmentCmd).Standalone()

	rootCmd.AddCommand(unsetEnvironmentCmd)

	carapace.Gen(unsetEnvironmentCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionEnvironmentVariables(unsetEnvironmentCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
