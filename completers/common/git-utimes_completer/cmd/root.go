package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-utimes",
	Short: "only touch files that are newer than their last commit date",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-utimes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("newer", false, "only touch files that are newer than their last commit date")
}
