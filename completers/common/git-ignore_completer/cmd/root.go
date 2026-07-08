package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-ignore",
	Short: "Add patterns to .gitignore, or display currently ignored patterns",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-ignore",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("global", "g", false, "Use global gitignore")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("local", "l", false, "Use local .gitignore (default)")
	rootCmd.Flags().BoolP("private", "p", false, "Use .git/info/exclude")
}
