package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "asr",
	Short: "Apple Software Restore",
	Long:  "https://keith.github.io/xcode-manpages/asr.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "Display help")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"restore", "Restore a disk image to a target",
			"server", "Serve a source over the network",
			"imagescan", "Scan an image for ASR",
			"help", "Display help",
			"version", "Display version",
		),
	)
}
