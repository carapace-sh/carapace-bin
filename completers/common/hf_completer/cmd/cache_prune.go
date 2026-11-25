package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove detached revisions and incomplete downloads from the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_pruneCmd).Standalone()

	cache_pruneCmd.Flags().String("cache-dir", "", "Cache directory to scan")
	cache_pruneCmd.Flags().Bool("dry-run", false, "Preview deletions without removing anything")
	cache_pruneCmd.Flags().String("format", "", "Output format")
	cache_pruneCmd.Flags().BoolP("help", "h", false, "Show help")
	cache_pruneCmd.Flags().Bool("json", false, "JSON output")
	cache_pruneCmd.Flags().Bool("no-dry-run", false, "Delete for real (default)")
	cache_pruneCmd.Flags().Bool("no-truncate", false, "Do not truncate values in human tables")
	cache_pruneCmd.Flags().BoolP("quiet", "q", false, "Quiet output")
	cache_pruneCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	cacheCmd.AddCommand(cache_pruneCmd)

	carapace.Gen(cache_pruneCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"format":    carapace.ActionValues("agent", "auto", "human", "json", "quiet"),
	})
}
