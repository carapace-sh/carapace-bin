package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dscacheutil",
	Short: "directory service cache utility",
	Long:  "https://keith.github.io/xcode-manpages/dscacheutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"-flushcache", "Flush the directory service cache",
			"-statistics", "Show cache statistics",
			"-q", "Query the cache",
			"-configuration", "Show the current configuration",
			"-cachedump", "Dump the cache contents",
		),
	)
}
