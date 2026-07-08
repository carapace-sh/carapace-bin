package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-merge-repo",
	Short: "Merge two repository histories",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-merge-repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("squash", false, "Merge as a single commit without full history")
}
