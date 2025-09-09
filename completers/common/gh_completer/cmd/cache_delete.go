package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_deleteCmd = &cobra.Command{
	Use:   "delete [<cache-id> | <cache-key> | --all]",
	Short: "Delete GitHub Actions caches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_deleteCmd).Standalone()

	cache_deleteCmd.Flags().BoolP("all", "a", false, "Delete all caches")
	cache_deleteCmd.Flags().StringP("ref", "r", "", "Delete by cache key and ref, formatted as refs/heads/<branch name> or refs/pull/<number>/merge")
	cache_deleteCmd.Flags().Bool("succeed-on-no-caches", false, "Return exit code 0 if no caches found. Must be used in conjunction with `--all`")
	cacheCmd.AddCommand(cache_deleteCmd)

	carapace.Gen(cache_deleteCmd).FlagCompletion(carapace.ActionMap{
		"ref": action.ActionCacheRefs(cache_deleteCmd),
	})

	carapace.Gen(cache_deleteCmd).PositionalCompletion(
		action.ActionCaches(cache_deleteCmd),
	)
}
