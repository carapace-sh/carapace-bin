package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migrate_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "DEPRECATED: Use 'server migration revert' instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migrate_revertCmd).Standalone()

	server_migrateCmd.AddCommand(server_migrate_revertCmd)
}
