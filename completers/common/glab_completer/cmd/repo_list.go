package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get list of repositories.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_listCmd).Standalone()

	repo_listCmd.Flags().BoolP("all", "a", false, "List all projects on the instance.")
	repo_listCmd.Flags().Bool("archived", false, "Limit by archived status. Use 'false' to exclude archived repositories. Used with the '--group' flag.")
	repo_listCmd.Flags().StringP("group", "g", "", "Return repositories in only the given group.")
	repo_listCmd.Flags().BoolP("include-subgroups", "G", false, "Include projects in subgroups of this group. Default is false. Used with the '--group' flag.")
	repo_listCmd.Flags().Bool("member", false, "List only projects of which you are a member.")
	repo_listCmd.Flags().BoolP("mine", "m", false, "List only projects you own. Default if no filters are provided.")
	repo_listCmd.Flags().StringP("order", "o", "", "Return repositories ordered by id, created_at, or other fields.")
	repo_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	repo_listCmd.Flags().StringP("page", "p", "", "Page number.")
	repo_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	repo_listCmd.Flags().StringP("sort", "s", "", "Return repositories sorted in asc or desc order.")
	repo_listCmd.Flags().Bool("starred", false, "List only starred projects.")
	repo_listCmd.Flags().StringP("user", "u", "", "List user projects.")
	repoCmd.AddCommand(repo_listCmd)

	carapace.Gen(repo_listCmd).FlagCompletion(carapace.ActionMap{
		"order":  carapace.ActionValues("id", "name", "path", "created_at", "updated_at", "last_activity_at", "repository_size", "storage_size", "packages_size", "wiki_size"),
		"output": carapace.ActionValues("text", "json"),
		"sort":   carapace.ActionValues("asc", "desc"),
		"user":   action.ActionUsers(repo_listCmd),
	})
}
