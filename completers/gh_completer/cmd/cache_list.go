package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var cache_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Github Actions caches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_listCmd).Standalone()

	cache_listCmd.Flags().StringP("limit", "L", "", "Maximum number of caches to fetch")
	cache_listCmd.Flags().StringP("order", "O", "", "Order of caches returned: {asc|desc}")
	cache_listCmd.Flags().StringP("sort", "S", "", "Sort fetched caches: {created_at|last_accessed_at|size_in_bytes}")
	cacheCmd.AddCommand(cache_listCmd)

	carapace.Gen(cache_listCmd).FlagCompletion(carapace.ActionMap{
		"order": carapace.ActionValues("asc", "desc").StyleF(style.ForKeyword),
		"sort":  carapace.ActionValues("created_at", "last_accessed_at", "size_in_bytes"),
	})
}
