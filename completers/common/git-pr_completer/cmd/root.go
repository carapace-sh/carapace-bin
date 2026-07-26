package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-pr",
	Short: "Check out a pull request from GitHub",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-pr",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("merge", "m", false, "Checkout the merge commit instead of the PR head")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("clean"),
		git.ActionRemotes(),
	)
}
