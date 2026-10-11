package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_preview_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove preview flags from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_preview_removeCmd).Standalone()
	workspace_help_previewCmd.AddCommand(workspace_help_preview_removeCmd)
}
