package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unsetEnvironmentCmd = &cobra.Command{
	Use:     "unset-environment",
	Short:   "Unset one or more environment variables",
	GroupID: "environment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsetEnvironmentCmd).Standalone()

	rootCmd.AddCommand(unsetEnvironmentCmd)

	carapace.Gen(unsetEnvironmentCmd).PositionalAnyCompletion(
		action.ActionEnvironmentVariables(unsetEnvironmentCmd).FilterArgs(),
	)
}
