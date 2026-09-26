package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the members of a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_listCmd).Standalone()

	workspace_listCmd.Flags().Bool("paths", false, "Show paths instead of names")
	workspace_listCmd.Flags().Bool("scripts", false, "List all standalone scripts with inline metadata in the workspace")
	workspaceCmd.AddCommand(workspace_listCmd)
}
