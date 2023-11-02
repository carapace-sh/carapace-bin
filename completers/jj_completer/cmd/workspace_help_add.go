package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_addCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_addCmd)
}
