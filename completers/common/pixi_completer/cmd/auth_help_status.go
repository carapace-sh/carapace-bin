package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show stored authentication entries and non-secret token metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_help_statusCmd).Standalone()

	auth_helpCmd.AddCommand(auth_help_statusCmd)
}
