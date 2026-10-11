package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_preview_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove preview flags from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_preview_help_removeCmd).Standalone()
	workspace_preview_helpCmd.AddCommand(workspace_preview_help_removeCmd)
}
