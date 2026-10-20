package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_repos_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Clone missing Git repositories and fast-forward the rest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_repos_updateCmd).Standalone()

	bootstrap_repos_updateCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_repos_updateCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_repos_updateCmd.Flags().Bool("skip-dirty", false, "Skip repositories with local changes instead of failing")
	bootstrap_repos_updateCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_reposCmd.AddCommand(bootstrap_repos_updateCmd)
}
