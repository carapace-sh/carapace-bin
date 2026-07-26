package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-count",
	Short: "Show commit count",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-count",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "Show verbose commit count details per author")
	rootCmd.Flags().Bool("help", false, "show help")
}
