package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var cache_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List GitHub Actions caches",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_listCmd).Standalone()

	cache_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	cache_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	cache_listCmd.Flags().StringP("key", "k", "", "Filter by cache key prefix")
	cache_listCmd.Flags().StringP("limit", "L", "", "Maximum number of caches to fetch")
	cache_listCmd.Flags().StringP("order", "O", "", "Order of caches returned: {asc|desc}")
	cache_listCmd.Flags().StringP("ref", "r", "", "Filter by ref, formatted as refs/heads/<branch name> or refs/pull/<number>/merge")
	cache_listCmd.Flags().StringP("sort", "S", "", "Sort fetched caches: {created_at|last_accessed_at|size_in_bytes}")
	cache_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	cacheCmd.AddCommand(cache_listCmd)

	carapace.Gen(cache_listCmd).FlagCompletion(carapace.ActionMap{
		"json":  gh.ActionCacheFields().UniqueList(","),
		"order": carapace.ActionValues("asc", "desc").StyleF(style.ForKeyword),
		"ref":   action.ActionCacheRefs(cache_listCmd),
		"sort":  carapace.ActionValues("created_at", "last_accessed_at", "size_in_bytes"),
	})
}
