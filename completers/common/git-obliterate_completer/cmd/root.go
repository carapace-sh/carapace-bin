package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-obliterate",
	Short: "Completely remove a file from the repository, including past commits and tags",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-obliterate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
	)
}
