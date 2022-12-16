package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
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
	repo_listCmd.Flags().BoolP("mine", "m", true, "Only list projects you own")
	repo_listCmd.Flags().StringP("order", "o", "last_activity_at", "Return repositories ordered by id, created_at, or other fields")
	repo_listCmd.Flags().IntP("page", "p", 1, "Page number")
	repo_listCmd.Flags().IntP("per-page", "P", 30, "Number of items to list per page.")
	repo_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	repo_listCmd.Flags().StringP("sort", "s", "", "Return repositories sorted in asc or desc order")
	repo_listCmd.Flags().Bool("starred", false, "Only list starred projects")
	repoCmd.AddCommand(repo_listCmd)

	carapace.Gen(repo_listCmd).FlagCompletion(carapace.ActionMap{
		"order": carapace.ActionValues("id", "name", "path", "created_at", "updated_at", "last_activity_at", "repository_size", "storage_size", "packages_size", "wiki_size"),
		"repo":  action.ActionRepo(repo_listCmd),
		"sort":  carapace.ActionValues("asc", "desc"),
	})
}
