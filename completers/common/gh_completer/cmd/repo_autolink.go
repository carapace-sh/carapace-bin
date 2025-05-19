package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_autolinkCmd = &cobra.Command{
	Use:     "autolink <command>",
	Short:   "Manage autolink references",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_autolinkCmd).Standalone()

	repo_autolinkCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	repoCmd.AddCommand(repo_autolinkCmd)

	carapace.Gen(repo_autolinkCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
