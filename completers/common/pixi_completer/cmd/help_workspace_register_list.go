package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_register_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the registered workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_register_listCmd).Standalone()

	help_workspace_registerCmd.AddCommand(help_workspace_register_listCmd)
}
