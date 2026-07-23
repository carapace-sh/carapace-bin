package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-magic",
	Short: "Automate add/commit/push routines",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-magic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Add everything including untracked")
	rootCmd.Flags().BoolS("e", "e", false, "Edit generated commit message")
	rootCmd.Flags().BoolS("f", "f", false, "Force push with lease")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().StringS("m", "m", "", "Commit message")
	rootCmd.Flags().BoolS("p", "p", false, "Push after commit")
}
