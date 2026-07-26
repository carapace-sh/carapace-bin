package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-delete-merged-branches",
	Short: "Delete branches that are listed in git branch --merged",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-delete-merged-branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
}
