package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_updateStaleCmd = &cobra.Command{
	Use:   "update-stale",
	Short: "Update a workspace that has become stale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_updateStaleCmd).Standalone()

	workspace_updateStaleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspaceCmd.AddCommand(workspace_updateStaleCmd)
}
