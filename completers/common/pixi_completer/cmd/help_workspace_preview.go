package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Commands to manage workspace preview flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_previewCmd).Standalone()
	help_workspaceCmd.AddCommand(help_workspace_previewCmd)
}
