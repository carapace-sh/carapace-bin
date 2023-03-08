package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_unarchiveCmd = &cobra.Command{
	Use:     "unarchive [<repository>]",
	Short:   "Unarchive a repository",
	GroupID: "targeted",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_unarchiveCmd).Standalone()

	repo_unarchiveCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	repoCmd.AddCommand(repo_unarchiveCmd)

	carapace.Gen(repo_unarchiveCmd).PositionalCompletion(
		gh.ActionOwnerRepositories(gh.HostOpts{}), // TODO should only be archived repos
	)
}
