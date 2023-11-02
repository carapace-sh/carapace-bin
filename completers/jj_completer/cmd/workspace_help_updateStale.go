package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_updateStaleCmd = &cobra.Command{
	Use:   "update-stale",
	Short: "Update a workspace that has become stale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_updateStaleCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_updateStaleCmd)
}
