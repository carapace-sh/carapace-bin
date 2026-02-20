package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_environment_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an environment from the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_environment_removeCmd).Standalone()

	help_workspace_environmentCmd.AddCommand(help_workspace_environment_removeCmd)
}
