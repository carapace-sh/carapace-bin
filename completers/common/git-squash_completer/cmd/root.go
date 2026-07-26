package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-squash",
	Short: "Merge commits from a source branch into the current branch as a single commit",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-squash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("squash-msg", false, "Use concatenated squashed messages")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
