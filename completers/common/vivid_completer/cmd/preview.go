package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vivid"
	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview a given theme",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previewCmd).Standalone()

	previewCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(previewCmd)

	carapace.Gen(previewCmd).PositionalCompletion(
		vivid.ActionThemes(),
	)
}
