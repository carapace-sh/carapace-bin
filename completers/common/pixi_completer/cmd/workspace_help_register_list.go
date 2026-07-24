package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_register_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the registered workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_register_listCmd).Standalone()

	workspace_help_registerCmd.AddCommand(workspace_help_register_listCmd)
}
