package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var integration_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show integration status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(integration_statusCmd).Standalone()

	integration_statusCmd.Flags().Bool("outdated-only", false, "")
	integrationCmd.AddCommand(integration_statusCmd)
}
