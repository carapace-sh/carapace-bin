package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-brv",
	Short: "List branches sorted by their last commit date",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-brv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("reverse", "r", false, "reverse output order")
}
