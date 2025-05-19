package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_viewCmd = &cobra.Command{
	Use:   "view [repository] [flags]",
	Short: "View a project or repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_viewCmd).Standalone()

	repo_viewCmd.Flags().StringP("branch", "b", "", "View a specific branch of the repository.")
	repo_viewCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	repo_viewCmd.Flags().BoolP("web", "w", false, "Open a project in the browser.")
	repoCmd.AddCommand(repo_viewCmd)

	carapace.Gen(repo_viewCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				action.FakeRepoFlag(repo_viewCmd, c.Args[0])
				return action.ActionBranches(repo_viewCmd)
			}
			return carapace.ActionValues()
		}),
		"output": carapace.ActionValues("text", "json"),
	})

	carapace.Gen(repo_viewCmd).PositionalCompletion(
		action.ActionRepo(repo_viewCmd),
	)
}
