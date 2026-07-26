package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show stored authentication entries and non-secret token metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_statusCmd).Standalone()

	auth_statusCmd.Flags().Bool("details", false, "Show endpoint URLs, client ID, and other IdP-introspection fields that are only useful for debugging")
	authCmd.AddCommand(auth_statusCmd)
}
