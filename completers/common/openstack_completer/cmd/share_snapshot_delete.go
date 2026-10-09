package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more share snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_deleteCmd).Standalone()

	share_snapshot_deleteCmd.Flags().Bool("force", false, "Delete the snapshot(s) ignoring the current state.")
	share_snapshot_deleteCmd.Flags().Bool("wait", false, "Wait for share snapshot deletion")
	share_snapshotCmd.AddCommand(share_snapshot_deleteCmd)
}
