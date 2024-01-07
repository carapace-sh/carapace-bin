package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/helm"
	"github.com/spf13/cobra"
)

var repo_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove one or more chart repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_removeCmd).Standalone()
	repoCmd.AddCommand(repo_removeCmd)

	carapace.Gen(repo_removeCmd).PositionalAnyCompletion(
		helm.ActionRepositories().FilterArgs(),
	)
}
