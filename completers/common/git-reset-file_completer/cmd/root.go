package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-reset-file",
	Short: "Reset one file to HEAD or a certain commit",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-reset-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
