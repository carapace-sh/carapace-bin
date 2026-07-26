package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-undo",
	Short: "Remove the latest N commits",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-undo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("hard", "h", false, "Wipe commits, with confirmation")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("soft", "s", false, "Rollback, keep changes in staging")
}
