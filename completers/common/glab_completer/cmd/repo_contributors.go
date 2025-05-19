package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_contributorsCmd = &cobra.Command{
	Use:     "contributors",
	Short:   "Get repository contributors list.",
	Aliases: []string{"users"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_contributorsCmd).Standalone()

	repo_contributorsCmd.Flags().StringP("order", "o", "", "Return contributors ordered by name, email, or commits (orders by commit date) fields.")
	repo_contributorsCmd.Flags().StringP("page", "p", "", "Page number.")
	repo_contributorsCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	repo_contributorsCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	repo_contributorsCmd.Flags().StringP("sort", "s", "", "Return contributors. Sort options: asc, desc.")
	repoCmd.AddCommand(repo_contributorsCmd)

	carapace.Gen(repo_contributorsCmd).FlagCompletion(carapace.ActionMap{
		"order": carapace.ActionValues("name", "email", "commits"),
		"repo":  action.ActionRepo(repo_contributorsCmd),
		"sort":  carapace.ActionValues("asc", "desc"),
	})
}
