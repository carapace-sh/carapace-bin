package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "defaults",
	Short: "access user preferences",
	Long:  "https://keith.github.io/xcode-manpages/defaults.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("current-host", "c", false, "Operate on the current host's defaults")
	rootCmd.Flags().BoolP("host", "h", false, "Operate on a specified host's defaults")
	rootCmd.Flags().Bool("version", false, "Print version information and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"read", "Prints the value for the given key",
			"read-type", "Prints the plist type for the given key",
			"write", "Writes the value for the given key",
			"rename", "Renames the given key",
			"delete", "Deletes the given key or domain",
			"import", "Reads the plist from stdin and sets it as the value",
			"export", "Exports the domain as XML plist to stdout",
			"domains", "Lists all domains",
			"find", "Searches for a key in all domains",
			"defaults", "Prints the defaults",
		),
	)
}
