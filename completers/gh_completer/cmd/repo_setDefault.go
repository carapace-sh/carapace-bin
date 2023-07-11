package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_setDefaultCmd = &cobra.Command{
	Use:     "set-default [<repository>]",
	Short:   "Configure default repository for this directory",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_setDefaultCmd).Standalone()

	repo_setDefaultCmd.Flags().BoolP("unset", "u", false, "unset the current default repository")
	repo_setDefaultCmd.Flags().BoolP("view", "v", false, "view the current default repository")
	repoCmd.AddCommand(repo_setDefaultCmd)

	carapace.Gen(repo_setDefaultCmd).PositionalCompletion(
		gh.ActionOwnerRepositories(gh.HostOpts{}),
	)
}
