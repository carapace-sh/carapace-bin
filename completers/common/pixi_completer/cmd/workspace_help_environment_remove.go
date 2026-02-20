package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_environment_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an environment from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_environment_removeCmd).Standalone()

	workspace_help_environmentCmd.AddCommand(workspace_help_environment_removeCmd)
}
