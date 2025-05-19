package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var reposCmd = &cobra.Command{
	Use:     "repos",
	Short:   "Show repository details",
	GroupID: "ENTITIES",
	Aliases: []string{"repo"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reposCmd).Standalone()

	reposCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	reposCmd.Flags().String("limit", "", "specify limit of items per page")
	reposCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	reposCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	reposCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	reposCmd.Flags().BoolP("starred", "s", false, "List your starred repos instead")
	reposCmd.Flags().StringP("type", "T", "", "Filter by type: fork, mirror, source")
	reposCmd.Flags().BoolP("watched", "w", false, "List your watched repos instead")
	rootCmd.AddCommand(reposCmd)

	// TODO completion
	carapace.Gen(reposCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionRepositoryFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"type":   carapace.ActionValues("fork", "mirror", "source"),
	})
}
