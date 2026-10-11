package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_preview_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add preview flags to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_preview_help_addCmd).Standalone()
	workspace_preview_helpCmd.AddCommand(workspace_preview_help_addCmd)
}
