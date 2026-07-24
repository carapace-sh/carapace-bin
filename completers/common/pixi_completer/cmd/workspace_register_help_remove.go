package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_register_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a workspace from registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_register_help_removeCmd).Standalone()

	workspace_register_helpCmd.AddCommand(workspace_register_help_removeCmd)
}
