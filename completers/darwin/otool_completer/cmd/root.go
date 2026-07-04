package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otool",
	Short: "Object file display tool",
	Long:  "https://keith.github.io/xcode-manpages/otool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "Display shared libraries used")
	rootCmd.Flags().BoolS("a", "a", false, "Display archive header")
	rootCmd.Flags().BoolS("h", "h", false, "Display Mach-O header")
	rootCmd.Flags().BoolS("l", "l", false, "Display load commands")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
