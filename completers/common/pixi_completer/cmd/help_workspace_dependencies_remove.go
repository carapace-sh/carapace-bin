package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_dependencies_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove dependencies from the workspace dependencies table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_dependencies_removeCmd).Standalone()
	help_workspace_dependenciesCmd.AddCommand(help_workspace_dependencies_removeCmd)
}
