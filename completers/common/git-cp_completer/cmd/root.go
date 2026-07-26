package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-cp",
	Short: "Copy a file to another keeping its history",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-cp",
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
		carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
	)
}
