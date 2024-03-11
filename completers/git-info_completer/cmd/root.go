package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-info",
	Short: "Returns information on current repository",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("color", "c", false, "Use color for information titles")
	rootCmd.Flags().Bool("help", false, "Display help message")
	rootCmd.Flags().Bool("no-config", false, "Don't show list all variables set in config file, along with their values")
}
