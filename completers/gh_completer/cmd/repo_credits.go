package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_creditsCmd = &cobra.Command{
	Use:     "credits [<repository>]",
	Short:   "View credits for a repository",
	GroupID: "Targeted commands",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_creditsCmd).Standalone()

	repo_creditsCmd.Flags().BoolP("static", "s", false, "Print a static version of the credits")
	repoCmd.AddCommand(repo_creditsCmd)
}
