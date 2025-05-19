package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-authors",
	Short: "Generate authors report",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-authors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "Display help message")
	rootCmd.Flags().BoolP("list", "l", false, "Show authors")
	rootCmd.Flags().Bool("no-email", false, "Don't show author's email")
}
