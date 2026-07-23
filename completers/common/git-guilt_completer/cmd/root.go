package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-guilt",
	Short: "Calculate the change in blame between two revisions",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-guilt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "Show debug output")
	rootCmd.Flags().BoolP("email", "e", false, "Display emails instead of names")
	rootCmd.Flags().BoolS("h", "h", false, "show help")
	rootCmd.Flags().BoolP("ignore-whitespace", "w", false, "Ignore whitespace changes")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
