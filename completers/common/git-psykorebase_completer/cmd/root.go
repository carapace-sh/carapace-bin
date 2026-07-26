package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-psykorebase",
	Short: "Rebase a branch on top of another using a merge commit and one conflict handling",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-psykorebase",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("continue", false, "Continue after conflict resolution")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("no-ff", false, "Force commit message even without conflicts")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionLocalBranches(),
		git.ActionLocalBranches(),
	)
}
