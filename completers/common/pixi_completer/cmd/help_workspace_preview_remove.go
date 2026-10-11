package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_preview_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove preview flags from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_preview_removeCmd).Standalone()
	help_workspace_previewCmd.AddCommand(help_workspace_preview_removeCmd)
}
