package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mdfind"
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

	rootCmd.Flags().BoolP("0", "0", false, "Use NUL as a path separator")
	rootCmd.Flags().StringP("attr", "attr", "", "Fetches the value of the specified attribute")
	rootCmd.Flags().BoolP("count", "count", false, "Show the count of matches")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("live", "live", false, "Continue searching as files are created/removed")
	rootCmd.Flags().StringP("name", "name", "", "Search for files with the specified name")
	rootCmd.Flags().StringP("onlyin", "onlyin", "", "Search only in the specified directory")
	rootCmd.Flags().BoolP("pl", "pl", false, "Output in plist format")
	rootCmd.Flags().BoolP("reprint", "reprint", false, "Reprint results on live update")
	rootCmd.Flags().StringP("s", "s", "", "Show contents of smart folder <name>")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"attr":   carapace.ActionValues(),
		"onlyin": carapace.ActionDirectories(),
		"s":      mdfind.ActionSavedSearches(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
