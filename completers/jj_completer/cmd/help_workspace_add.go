package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_addCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_addCmd)
}
