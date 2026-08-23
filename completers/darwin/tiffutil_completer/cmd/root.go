package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiffutil",
	Short: "manipulates tiff files",
	Long:  "https://keith.github.io/xcode-manpages/tiffutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("out", "out", "", "Output file")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("none", "lzw", "packbits", "cat", "catnosizecheck", "cathidpicheck", "extract", "info", "verboseinfo", "dump"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".tiff", ".tif"),
	)
}
