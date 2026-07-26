package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-reauthor",
	Short: "Rewrite history to change author's identity",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-reauthor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Rewrite ALL identities")
	rootCmd.Flags().StringP("correct-email", "e", "", "Correct email to set")
	rootCmd.Flags().StringP("correct-name", "n", "", "Correct name to set")
	rootCmd.Flags().BoolS("h", "h", false, "show help")
	rootCmd.Flags().StringP("old-email", "o", "", "Match identities with this old email")
	rootCmd.Flags().StringP("type", "t", "", "Type of identity to rewrite")
	rootCmd.Flags().BoolP("use-config", "c", false, "Use values from Git config")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("author", "committer", "both"),
	})
}
