package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more share group snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshot_deleteCmd).Standalone()

	share_group_snapshot_deleteCmd.Flags().Bool("force", false, "Attempt to force delete the share group snapshot(s) (Default=False) (Admin only).")
	share_group_snapshot_deleteCmd.Flags().Bool("wait", false, "Wait for share group snapshot deletion")
	share_group_snapshotCmd.AddCommand(share_group_snapshot_deleteCmd)
}
