package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migration_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert server migration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migration_revertCmd).Standalone()

	server_migrationCmd.AddCommand(server_migration_revertCmd)
}
