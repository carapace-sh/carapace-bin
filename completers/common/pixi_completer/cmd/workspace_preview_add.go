package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_preview_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add preview flags to the workspace",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_preview_addCmd).Standalone()

	workspace_previewCmd.AddCommand(workspace_preview_addCmd)

	carapace.Gen(workspace_preview_addCmd).PositionalAnyCompletion(
		carapace.ActionValues("pixi-build"),
	)
}
