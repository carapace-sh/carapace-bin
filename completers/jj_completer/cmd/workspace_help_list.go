package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_listCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_listCmd)
}
