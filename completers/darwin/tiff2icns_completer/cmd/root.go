package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiff2icns",
	Short: "converts TIFF to icns format",
	Long:  "https://keith.github.io/xcode-manpages/tiff2icns.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("noLarge", false, "Do not create a large 32x32 icon")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".tiff", ".tif"),
	)
}
