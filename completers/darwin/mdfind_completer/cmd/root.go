package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdfind",
	Short: "find files using Spotlight",
	Long:  "https://keith.github.io/xcode-manpages/mdfind.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("count", "count", false, "Show the count of matches")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("live", "live", false, "Continue searching as files are created/removed")
	rootCmd.Flags().StringP("name", "name", "", "Search for files with the specified name")
	rootCmd.Flags().StringP("onlyin", "onlyin", "", "Search only in the specified directory")
	rootCmd.Flags().BoolP("pl", "pl", false, "Output in plist format")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"onlyin": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
