package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-stamp",
	Short: "Stamp the last commit message with issue references or review links",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-stamp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("replace", "r", false, "Replace all previous stamps with same identifier")
}
