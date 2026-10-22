package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate volume to a new host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_migrateCmd).Standalone()

	volume_migrateCmd.Flags().String("cluster", "", "Destination cluster to migrate the volume to (requires --os-volume-api-version 3.16 or higher)")
	volume_migrateCmd.Flags().Bool("force-host-copy", false, "Enable generic host-based force-migration, which bypasses driver optimizations")
	volume_migrateCmd.Flags().String("host", "", "Destination host (takes the form: host@backend-name#pool)")
	volume_migrateCmd.Flags().Bool("lock-volume", false, "If specified, the volume state will be locked and will not allow a migration to be aborted (possibly by another operation)")
	volumeCmd.AddCommand(volume_migrateCmd)
}
