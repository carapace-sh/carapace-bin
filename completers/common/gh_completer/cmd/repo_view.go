package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_viewCmd = &cobra.Command{
	Use:     "view [<repository>]",
	Short:   "View a repository",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_viewCmd).Standalone()

	repo_viewCmd.Flags().StringP("branch", "b", "", "View a specific branch of the repository")
	repo_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	repo_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	repo_viewCmd.Flags().BoolP("web", "w", false, "Open a repository in the browser")
	repoCmd.AddCommand(repo_viewCmd)

	carapace.Gen(repo_viewCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				repo_viewCmd.Flags().String("repo", c.Args[0], "fake repo flag for ActionBranches")
				return action.ActionBranches(repo_viewCmd)
			} else {
				return carapace.ActionValues()
			}
		}),
		"json": action.ActionRepositoryFields().UniqueList(","),
	})

	carapace.Gen(repo_viewCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_viewCmd),
	)
}
