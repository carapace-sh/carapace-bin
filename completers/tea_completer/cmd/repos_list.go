package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var repos_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List repositories you have access to",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repos_listCmd).Standalone()

	repos_listCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	repos_listCmd.Flags().String("limit", "", "specify limit of items per page")
	repos_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	repos_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	repos_listCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	repos_listCmd.Flags().BoolP("starred", "s", false, "List your starred repos instead")
	repos_listCmd.Flags().StringP("type", "T", "", "Filter by type: fork, mirror, source")
	repos_listCmd.Flags().BoolP("watched", "w", false, "List your watched repos instead")
	reposCmd.AddCommand(repos_listCmd)

	// TODO completion
	carapace.Gen(repos_listCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionRepositoryFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"type":   carapace.ActionValues("fork", "mirror", "source"),
	})
}
