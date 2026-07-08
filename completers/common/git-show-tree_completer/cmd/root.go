package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-show-tree",
	Short: "Show decorated graph view of one-liner summarized commits from all branches",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-show-tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
}
