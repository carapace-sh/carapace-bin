package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_preview_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove preview flags from the workspace",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_preview_removeCmd).Standalone()

	workspace_preview_removeCmd.Flags().Bool("force", false, "Remove flags even when the manifest no longer loads without them")
	workspace_previewCmd.AddCommand(workspace_preview_removeCmd)

	carapace.Gen(workspace_preview_removeCmd).PositionalAnyCompletion(
		carapace.ActionValues("pixi-build"),
	)
}
