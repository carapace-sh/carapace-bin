package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_preview_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the enabled preview flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_preview_listCmd).Standalone()
	help_workspace_previewCmd.AddCommand(help_workspace_preview_listCmd)
}
