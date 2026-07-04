package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xcode-select",
	Short: "switch the active Xcode developer directory",
	Long:  "https://keith.github.io/xcode-manpages/xcode-select.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("install", "i", false, "Install the Xcode command line tools")
	rootCmd.Flags().BoolP("print-path", "p", false, "Print the active Xcode developer directory")
	rootCmd.Flags().StringP("switch", "s", "", "Switch the active Xcode developer directory")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"switch": carapace.ActionDirectories(),
	})
}
