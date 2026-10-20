// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_repos_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Clone missing Git repositories and check out configured refs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_repos_applyCmd).Standalone()

	bootstrap_repos_applyCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_repos_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_repos_applyCmd.Flags().Bool("skip-dirty", false, "Skip repositories with local changes instead of failing")
	bootstrap_repos_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_reposCmd.AddCommand(bootstrap_repos_applyCmd)
}
