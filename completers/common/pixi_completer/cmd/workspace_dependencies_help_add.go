package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_dependencies_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add dependencies to the workspace dependencies table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_dependencies_help_addCmd).Standalone()
	workspace_dependencies_helpCmd.AddCommand(workspace_dependencies_help_addCmd)
}
