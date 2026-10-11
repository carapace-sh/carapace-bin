package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_preview_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the enabled preview flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_preview_listCmd).Standalone()
	workspace_help_previewCmd.AddCommand(workspace_help_preview_listCmd)
}
