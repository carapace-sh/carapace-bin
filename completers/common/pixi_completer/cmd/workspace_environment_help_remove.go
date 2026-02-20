package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_environment_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an environment from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_help_removeCmd).Standalone()

	workspace_environment_helpCmd.AddCommand(workspace_environment_help_removeCmd)
}
