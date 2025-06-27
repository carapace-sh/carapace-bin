package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Show the underlying Git directory of a repository using the Git backend",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_rootCmd).Standalone()

	git_rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gitCmd.AddCommand(git_rootCmd)
}
