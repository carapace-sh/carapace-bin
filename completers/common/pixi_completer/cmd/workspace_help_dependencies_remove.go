package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_dependencies_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove dependencies from the workspace dependencies table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_dependencies_removeCmd).Standalone()
	workspace_help_dependenciesCmd.AddCommand(workspace_help_dependencies_removeCmd)
}
