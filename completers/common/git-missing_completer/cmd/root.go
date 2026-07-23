package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-missing",
	Short: "Show commits on one branch or the other but not both",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-missing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionLocalBranches(),
		git.ActionLocalBranches(),
	)
}
