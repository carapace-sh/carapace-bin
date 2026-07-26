package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-merge-into",
	Short: "Merge source branch into destination, staying on current branch",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-merge-into",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("ff-only", false, "Refuse to merge unless fast-forward")
	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionLocalBranches(),
		git.ActionLocalBranches(),
	)
}
