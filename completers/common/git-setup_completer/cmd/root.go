package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-setup",
	Short: "Set up a git repository, add all files, and make an initial commit",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-setup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().StringP("message", "m", "", "Initial commit message")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
