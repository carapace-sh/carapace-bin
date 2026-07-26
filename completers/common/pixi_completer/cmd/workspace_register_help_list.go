package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_register_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the registered workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_register_help_listCmd).Standalone()

	workspace_register_helpCmd.AddCommand(workspace_register_help_listCmd)
}
