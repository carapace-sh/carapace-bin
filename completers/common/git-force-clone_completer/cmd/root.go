package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-force-clone",
	Short: "Clone a repo, destroying local changes if the target directory exists",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-force-clone",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("branch", "", "Checkout this branch after clone")
	rootCmd.Flags().Bool("help", false, "show help")
}
