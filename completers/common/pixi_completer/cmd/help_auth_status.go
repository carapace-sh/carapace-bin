package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auth_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show stored authentication entries and non-secret token metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auth_statusCmd).Standalone()

	help_authCmd.AddCommand(help_auth_statusCmd)
}
