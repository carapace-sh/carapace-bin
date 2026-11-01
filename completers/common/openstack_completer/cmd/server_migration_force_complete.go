package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_migration_force_completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Force an ongoing live migration to complete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migration_force_completeCmd).Standalone()

	server_migration_forceCmd.AddCommand(server_migration_force_completeCmd)

	carapace.Gen(server_migration_force_completeCmd).PositionalCompletion(
		action.ActionServers(server_migration_force_completeCmd),
		action.ActionServerMigrations(server_migration_force_completeCmd),
	)
}
