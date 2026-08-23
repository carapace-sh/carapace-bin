package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iconutil",
	Short: "convert between .iconset and .icns files",
	Long:  "https://keith.github.io/xcode-manpages/iconutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("convert", "c", "", "Convert to the specified format")
	rootCmd.Flags().StringP("output", "o", "", "Override the default output file name")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"convert": carapace.ActionValues("icns", "iconset"),
		"output":  carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
