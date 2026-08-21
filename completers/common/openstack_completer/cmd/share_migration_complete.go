package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_migration_completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Completes migration for a given share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_migration_completeCmd).Standalone()

	share_migrationCmd.AddCommand(share_migration_completeCmd)
}
