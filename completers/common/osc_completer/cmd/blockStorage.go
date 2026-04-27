package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var blockStorageCmd = &cobra.Command{
	Use:   "block-storage",
	Short: "Block Storage (Volume) service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var blockStorageVolumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Volume commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var blockStorageSnapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Snapshot commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var blockStorageBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var blockStorageVolumeTypeCmd = &cobra.Command{
	Use:   "volume-type",
	Short: "Volume Type commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockStorageCmd).Standalone()
	carapace.Gen(blockStorageVolumeCmd).Standalone()
	carapace.Gen(blockStorageSnapshotCmd).Standalone()
	carapace.Gen(blockStorageBackupCmd).Standalone()
	carapace.Gen(blockStorageVolumeTypeCmd).Standalone()

	rootCmd.AddCommand(blockStorageCmd)
	blockStorageCmd.AddCommand(blockStorageVolumeCmd)
	blockStorageCmd.AddCommand(blockStorageSnapshotCmd)
	blockStorageCmd.AddCommand(blockStorageBackupCmd)
	blockStorageCmd.AddCommand(blockStorageVolumeTypeCmd)

	for _, cmd := range []*cobra.Command{blockStorageVolumeCmd, blockStorageSnapshotCmd, blockStorageBackupCmd, blockStorageVolumeTypeCmd} {
		listCmd := &cobra.Command{Use: "list", Short: "List resources", Run: func(cmd *cobra.Command, args []string) {}}
		showCmd := &cobra.Command{Use: "show", Short: "Show resource details", Run: func(cmd *cobra.Command, args []string) {}}
		createCmd := &cobra.Command{Use: "create", Short: "Create a resource", Run: func(cmd *cobra.Command, args []string) {}}
		deleteCmd := &cobra.Command{Use: "delete", Short: "Delete a resource", Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(listCmd).Standalone()
		carapace.Gen(showCmd).Standalone()
		carapace.Gen(createCmd).Standalone()
		carapace.Gen(deleteCmd).Standalone()
		cmd.AddCommand(listCmd, showCmd, createCmd, deleteCmd)
	}
}
