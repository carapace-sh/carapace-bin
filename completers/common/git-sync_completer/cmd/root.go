package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-sync",
	Short: "Sync local branch with its remote branch",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-sync",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "No interaction")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("soft", "s", false, "Preserve untracked files")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionRemotes(),
		git.ActionLocalBranches(),
	)
}
