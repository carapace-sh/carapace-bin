package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate server to different host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_migrateCmd).Standalone()

	server_migrateCmd.Flags().Bool("block-migration", false, "Perform a block live migration (auto-configured from --os-compute-api-version 2.25)")
	server_migrateCmd.Flags().Bool("disk-overcommit", false, "Allow disk over-commit on the destination host (supported with --os-compute-api-version 2.24 or below)")
	server_migrateCmd.Flags().String("host", "", "Migrate the server to the specified host.")
	server_migrateCmd.Flags().Bool("live-migration", false, "Live migrate the server; use the ``--host`` option to specify a target host for the migration which will be validated by the scheduler")
	server_migrateCmd.Flags().Bool("no-disk-overcommit", false, "Do not over-commit disk on the destination host (default) (supported with --os-compute-api-version 2.24 or below)")
	server_migrateCmd.Flags().Bool("shared-migration", false, "Perform a shared live migration (default before --os-compute-api-version 2.25, auto after)")
	server_migrateCmd.Flags().Bool("wait", false, "Wait for migrate to complete")
	serverCmd.AddCommand(server_migrateCmd)
}
