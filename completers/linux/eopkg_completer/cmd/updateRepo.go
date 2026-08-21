package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var updateRepoCmd = &cobra.Command{
	Use:     "update-repo <reponame?>",
	Aliases: []string{"ur"},
	Short:   "update repository indexes by fetching them from their origin if a change has occurred",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateRepoCmd).Standalone()

	updateRepoCmd.Flags().BoolP("force", "f", false, "forcibly update the repository indexes, even if the checksum file hasn't changed")

	rootCmd.AddCommand(updateRepoCmd)

	carapace.Gen(updateRepoCmd).PositionalAnyCompletion(
		eopkg.ActionRepositories().FilterArgs(),
	)
}
