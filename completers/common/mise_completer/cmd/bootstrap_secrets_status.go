// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_secrets_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show whether declared bootstrap secret inputs are available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_secrets_statusCmd).Standalone()

	bootstrap_secrets_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_secrets_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_secrets_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if a declared secret input is unset, empty, or not valid Unicode")
	bootstrap_secretsCmd.AddCommand(bootstrap_secrets_statusCmd)
}
