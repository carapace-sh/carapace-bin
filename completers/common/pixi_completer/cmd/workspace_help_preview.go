package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Commands to manage workspace preview flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_previewCmd).Standalone()
	workspace_helpCmd.AddCommand(workspace_help_previewCmd)
}
