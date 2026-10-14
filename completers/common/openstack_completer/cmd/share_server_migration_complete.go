package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_migration_completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Completes migration for a given share server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_migration_completeCmd).Standalone()

	share_server_migrationCmd.AddCommand(share_server_migration_completeCmd)
}
