package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_deployKeyCmd = &cobra.Command{
	Use:     "deploy-key <command>",
	Short:   "Manage deploy keys in a repository",
	GroupID: "targeted",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deployKeyCmd).Standalone()

	repo_deployKeyCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	repoCmd.AddCommand(repo_deployKeyCmd)

	carapace.Gen(repo_deployKeyCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepoOverride(repo_deployKeyCmd),
	})
}
