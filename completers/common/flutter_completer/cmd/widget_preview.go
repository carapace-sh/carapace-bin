package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var widget_previewCmd = &cobra.Command{
	Use:   "widget-preview",
	Short: "Manage the widget preview environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(widget_previewCmd).Standalone()

	widget_previewCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(widget_previewCmd)
}
