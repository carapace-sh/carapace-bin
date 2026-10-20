// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_user_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show whether your login shell matches `[bootstrap.user]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_user_statusCmd).Standalone()

	bootstrap_user_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_user_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_user_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if your login shell differs from the config")
	bootstrap_userCmd.AddCommand(bootstrap_user_statusCmd)
}
