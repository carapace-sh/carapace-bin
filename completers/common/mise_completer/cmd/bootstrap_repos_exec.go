// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_repos_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command in each configured Git repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_repos_execCmd).Standalone()

	bootstrap_repos_execCmd.Flags().BoolP("continue-on-error", "c", false, "Continue running in other repositories after a command fails")
	bootstrap_repos_execCmd.Flags().BoolP("dry-run", "n", false, "Print the commands that would run without running them")
	bootstrap_repos_execCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_reposCmd.AddCommand(bootstrap_repos_execCmd)
}
