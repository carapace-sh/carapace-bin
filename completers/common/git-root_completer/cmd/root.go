package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-root",
	Short: "Show the path to the root directory of the git repo",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-root",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("relative", "r", false, "Show relative path from root to current directory")
}
