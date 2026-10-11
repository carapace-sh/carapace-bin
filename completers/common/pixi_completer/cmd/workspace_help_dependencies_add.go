package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_dependencies_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add dependencies to the workspace dependencies table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_dependencies_addCmd).Standalone()
	workspace_help_dependenciesCmd.AddCommand(workspace_help_dependencies_addCmd)
}
