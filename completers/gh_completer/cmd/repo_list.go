package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories owned by user or organization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_listCmd).Standalone()
	repo_listCmd.Flags().Bool("archived", false, "Show only archived repositories")
	repo_listCmd.Flags().Bool("fork", false, "Show only forks")
	repo_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_listCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	repo_listCmd.Flags().StringP("language", "l", "", "Filter by primary coding language")
	repo_listCmd.Flags().IntP("limit", "L", 30, "Maximum number of repositories to list")
	repo_listCmd.Flags().Bool("no-archived", false, "Omit archived repositories")
	repo_listCmd.Flags().Bool("private", false, "Show only private repositories")
	repo_listCmd.Flags().Bool("public", false, "Show only public repositories")
	repo_listCmd.Flags().Bool("source", false, "Show only non-forks")
	repo_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	repo_listCmd.Flags().String("topic", "", "Filter by topic")
	repoCmd.AddCommand(repo_listCmd)

	carapace.Gen(repo_listCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionRepositoryFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"language": action.ActionLanguages(),
		"topic": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionTopics(repo_listCmd, c.Args[0])
			}
			return action.ActionTopics(repo_listCmd, "")
		}),
	})

	carapace.Gen(repo_listCmd).PositionalCompletion(
		action.ActionUsers(repo_listCmd, action.UserOpts{Users: true, Organizations: true}),
	)
}
