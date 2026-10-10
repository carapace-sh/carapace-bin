package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
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

	repo_contributorsCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	repo_contributorsCmd.Flags().StringP("order", "o", "commits", "Return contributors ordered by name, email, or commits (orders by commit date) fields.")
	repo_contributorsCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	repo_contributorsCmd.Flags().StringP("page", "p", "1", "Page number.")
	repo_contributorsCmd.Flags().StringP("per-page", "P", "30", "Number of items to list per page.")
	repo_contributorsCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	repo_contributorsCmd.Flags().StringP("sort", "s", "", "Sort direction for --order field: asc or desc.")
	repoCmd.AddCommand(repo_contributorsCmd)

	carapace.Gen(repo_contributorsCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"order":  carapace.ActionValues("name", "email", "commits"),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(repo_contributorsCmd),
		"sort":   carapace.ActionValues("asc", "desc"),
	})
}
