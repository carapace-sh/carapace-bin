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

	rootCmd.Flags().StringP("a", "a", "", "Add key value to query")
	rootCmd.Flags().Bool("buckets", false, "Dump cache buckets with -cachedump")
	rootCmd.Flags().Bool("cachedump", false, "Dump the cache contents")
	rootCmd.Flags().Bool("configuration", false, "Show the current configuration")
	rootCmd.Flags().StringP("entries", "e", "", "Limit cache dump to entries of category")
	rootCmd.Flags().Bool("flushcache", false, "Flush the directory service cache")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("q", "q", "", "Query the cache for category")
	rootCmd.Flags().Bool("statistics", false, "Show cache statistics")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"q": carapace.ActionValues("user", "group", "host", "mount", "user_byuid", "group_bygid", "user_byname", "group_byname"),
	})
}
