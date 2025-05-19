package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_searchCmd = &cobra.Command{
	Use:     "search [flags]",
	Short:   "Search for GitLab repositories and projects by name.",
	Aliases: []string{"find", "lookup"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_searchCmd).Standalone()

	repo_searchCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	repo_searchCmd.Flags().StringP("page", "p", "", "Page number.")
	repo_searchCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	repo_searchCmd.Flags().StringP("search", "s", "", "A string contained in the project name.")
	repoCmd.AddCommand(repo_searchCmd)

	carapace.Gen(repo_searchCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
