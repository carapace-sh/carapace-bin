package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_preview_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the enabled preview flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_preview_help_listCmd).Standalone()
	workspace_preview_helpCmd.AddCommand(workspace_preview_help_listCmd)
}
