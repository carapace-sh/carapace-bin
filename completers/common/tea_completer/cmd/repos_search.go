package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var repos_searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Find any repo on an Gitea instance",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repos_searchCmd).Standalone()

	repos_searchCmd.Flags().String("archived", "", "Filter archived repos (true|false)")
	repos_searchCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	repos_searchCmd.Flags().String("limit", "", "specify limit of items per page")
	repos_searchCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	repos_searchCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	repos_searchCmd.Flags().StringP("owner", "O", "", "Filter by owner")
	repos_searchCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	repos_searchCmd.Flags().String("private", "", "Filter private repos (true|false)")
	repos_searchCmd.Flags().BoolP("topic", "t", false, "Search for term in repo topics instead of name")
	repos_searchCmd.Flags().StringP("type", "T", "", "Filter by type: fork, mirror, source")
	reposCmd.AddCommand(repos_searchCmd)

	// TODO completion
	carapace.Gen(repos_searchCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionRepositoryFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"type":   carapace.ActionValues("fork", "mirror", "source"),
	})
}
