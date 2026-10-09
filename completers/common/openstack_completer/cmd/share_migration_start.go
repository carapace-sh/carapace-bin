package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_migration_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Migrates share to a new host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_migration_startCmd).Standalone()

	share_migration_startCmd.Flags().String("force-host-assisted-migration", "", "Enforces the use of the host-assisted migration approach, which bypasses driver optimizations.")
	share_migration_startCmd.Flags().String("new-share-network", "", "Specify the new share network for the share.")
	share_migration_startCmd.Flags().String("new-share-type", "", "Specify the new share type for the share.")
	share_migration_startCmd.Flags().String("nondisruptive", "", "Enforces migration to be nondisruptive.")
	share_migration_startCmd.Flags().String("preserve-metadata", "", "Enforces migration to preserve all file metadata when moving its contents.")
	share_migration_startCmd.Flags().String("preserve-snapshots", "", "Enforces migration of the share snapshots to the destination.")
	share_migration_startCmd.Flags().String("writable", "", "Enforces migration to keep the share writable while contents are being moved.")
	share_migration_startCmd.MarkFlagRequired("nondisruptive")
	share_migration_startCmd.MarkFlagRequired("preserve-metadata")
	share_migration_startCmd.MarkFlagRequired("preserve-snapshots")
	share_migration_startCmd.MarkFlagRequired("writable")
	share_migrationCmd.AddCommand(share_migration_startCmd)
}
