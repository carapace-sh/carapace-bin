package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshot_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share group snapshot properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshot_setCmd).Standalone()

	share_group_snapshot_setCmd.Flags().String("description", "", "Set a description to the snapshot.")
	share_group_snapshot_setCmd.Flags().String("name", "", "Set a name to the snapshot.")
	share_group_snapshot_setCmd.Flags().String("status", "", "Explicitly set the state of a share group snapshot(Admin only).")
	share_group_snapshotCmd.AddCommand(share_group_snapshot_setCmd)
}
