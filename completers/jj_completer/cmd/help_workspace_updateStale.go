package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_updateStaleCmd = &cobra.Command{
	Use:   "update-stale",
	Short: "Update a workspace that has become stale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_updateStaleCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_updateStaleCmd)
}
