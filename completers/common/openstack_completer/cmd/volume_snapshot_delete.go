package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete volume snapshot(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_snapshot_deleteCmd).Standalone()

	volume_snapshot_deleteCmd.Flags().Bool("force", false, "Attempt forced removal of snapshot(s), regardless of state (defaults to False)")
	volume_snapshot_deleteCmd.Flags().Bool("remote", false, "Unmanage the snapshot, removing it from the Block Storage service management but not from the backend.")
	volume_snapshotCmd.AddCommand(volume_snapshot_deleteCmd)
}
