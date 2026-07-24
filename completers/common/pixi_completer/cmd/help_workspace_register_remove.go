package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_register_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a workspace from registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_register_removeCmd).Standalone()

	help_workspace_registerCmd.AddCommand(help_workspace_register_removeCmd)
}
