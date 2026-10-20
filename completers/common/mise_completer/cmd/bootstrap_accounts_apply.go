// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_accounts_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create, update, or remove the configured Linux users and groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_accounts_applyCmd).Standalone()

	bootstrap_accounts_applyCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_accounts_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_accounts_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_accountsCmd.AddCommand(bootstrap_accounts_applyCmd)
}
