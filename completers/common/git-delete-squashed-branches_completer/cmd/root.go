package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-delete-squashed-branches",
	Short: "Delete branches that have been squash-merged into a specified branch",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-delete-squashed-branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("proceed", "p", false, "Continue even if a branch can't be deleted")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionLocalBranches(),
	)
}
