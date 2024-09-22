package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-clear",
	Short: "Rigorously clean up a repository",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-clear",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "Force the clean, with no warning to the user")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
}
