package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List cached repositories or revisions",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_listCmd).Standalone()

	cache_listCmd.Flags().String("cache-dir", "", "Cache directory to scan")
	cache_listCmd.Flags().StringArrayP("filter", "f", nil, "Filter entries (e.g. size>1GB, type=model, accessed>7d)")
	cache_listCmd.Flags().String("format", "", "Output format")
	cache_listCmd.Flags().BoolP("help", "h", false, "Show help")
	cache_listCmd.Flags().Bool("json", false, "JSON output")
	cache_listCmd.Flags().Int("limit", 0, "Limit the number of results")
	cache_listCmd.Flags().Bool("no-revisions", false, "Aggregate by repository (default)")
	cache_listCmd.Flags().Bool("no-show-warnings", false, "Hide cache warnings (default)")
	cache_listCmd.Flags().Bool("no-truncate", false, "Do not truncate values in human tables")
	cache_listCmd.Flags().BoolP("quiet", "q", false, "Quiet output")
	cache_listCmd.Flags().Bool("revisions", false, "Include revisions instead of aggregated repositories")
	cache_listCmd.Flags().Bool("show-warnings", false, "Show warnings about cache inconsistencies")
	cache_listCmd.Flags().String("sort", "", "Sort entries by key")
	cacheCmd.AddCommand(cache_listCmd)

	carapace.Gen(cache_listCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"format":    carapace.ActionValues("agent", "auto", "human", "json", "quiet"),
		"sort": carapace.ActionValues(
			"accessed", "accessed:asc", "accessed:desc",
			"modified", "modified:asc", "modified:desc",
			"name", "name:asc", "name:desc",
			"size", "size:asc", "size:desc",
		).Tag("sort keys"),
	})
}
