package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_migration_cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancels migration of a given share when copying",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_migration_cancelCmd).Standalone()

	share_migrationCmd.AddCommand(share_migration_cancelCmd)
}
