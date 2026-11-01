package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var widget_preview_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans up widget preview state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(widget_preview_cleanCmd).Standalone()

	widget_preview_cleanCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	widget_previewCmd.AddCommand(widget_preview_cleanCmd)
}
