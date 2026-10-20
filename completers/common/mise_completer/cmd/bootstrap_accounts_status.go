// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_accounts_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the state of the configured Linux users and groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_accounts_statusCmd).Standalone()

	bootstrap_accounts_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_accounts_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_accounts_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if any account differs from the config (all are still listed)")
	bootstrap_accountsCmd.AddCommand(bootstrap_accounts_statusCmd)
}
