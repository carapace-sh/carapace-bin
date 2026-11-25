package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove cached repositories, revisions or files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_rmCmd).Standalone()

	cache_rmCmd.Flags().String("cache-dir", "", "Cache directory to scan")
	cache_rmCmd.Flags().Bool("dry-run", false, "Preview deletions without removing anything")
	cache_rmCmd.Flags().String("format", "", "Output format")
	cache_rmCmd.Flags().BoolP("help", "h", false, "Show help")
	cache_rmCmd.Flags().Bool("json", false, "JSON output")
	cache_rmCmd.Flags().Bool("no-dry-run", false, "Delete for real (default)")
	cache_rmCmd.Flags().Bool("no-truncate", false, "Do not truncate values in human tables")
	cache_rmCmd.Flags().BoolP("quiet", "q", false, "Quiet output")
	cache_rmCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	cacheCmd.AddCommand(cache_rmCmd)

	carapace.Gen(cache_rmCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"format":    carapace.ActionValues("agent", "auto", "human", "json", "quiet"),
	})

	carapace.Gen(cache_rmCmd).PositionalAnyCompletion(
		action.ActionCachedRepos(cache_rmCmd).FilterArgs(),
	)
}
