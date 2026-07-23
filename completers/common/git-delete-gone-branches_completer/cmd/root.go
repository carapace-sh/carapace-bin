package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-delete-gone-branches",
	Short: "Delete local branches whose remote-tracking branch has been deleted",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-delete-gone-branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("dry-run", "n", false, "Preview only")
	rootCmd.Flags().BoolP("force", "f", false, "Skip confirmation")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("prune", "p", false, "Run git fetch --prune first")
}
