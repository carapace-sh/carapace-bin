package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_preview_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add preview flags to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_preview_addCmd).Standalone()
	help_workspace_previewCmd.AddCommand(help_workspace_preview_addCmd)
}
