package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_gardenCmd = &cobra.Command{
	Use:     "garden [<repository>]",
	Short:   "Explore a git repository as a garden",
	GroupID: "Targeted commands",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_gardenCmd).Standalone()

	repoCmd.AddCommand(repo_gardenCmd)
}
