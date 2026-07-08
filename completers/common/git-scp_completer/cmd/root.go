package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-scp",
	Short: "Copy files from the current working tree to a remote repo's working directory",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-scp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionRemotes(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
	)
}
