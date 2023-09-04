package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repo_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get list of repositories.",
	Aliases: []string{"users"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_listCmd).Standalone()

	repo_listCmd.Flags().BoolP("all", "a", false, "List all projects on the instance")
	repo_listCmd.Flags().Bool("member", false, "Only list projects which you are a member")
	repo_listCmd.Flags().BoolP("mine", "m", false, "Only list projects you own (default if no filters are passed)")
	repo_listCmd.Flags().StringP("order", "o", "", "Return repositories ordered by id, created_at, or other fields")
	repo_listCmd.Flags().StringP("page", "p", "", "Page number")
	repo_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page")
	repo_listCmd.Flags().StringP("sort", "s", "", "Return repositories sorted in asc or desc order")
	repo_listCmd.Flags().Bool("starred", false, "Only list starred projects")
	repoCmd.AddCommand(repo_listCmd)

	carapace.Gen(repo_listCmd).FlagCompletion(carapace.ActionMap{
		"order": carapace.ActionValues("id", "name", "path", "created_at", "updated_at", "last_activity_at", "repository_size", "storage_size", "packages_size", "wiki_size"),
		"sort":  carapace.ActionValues("asc", "desc"),
	})
}
