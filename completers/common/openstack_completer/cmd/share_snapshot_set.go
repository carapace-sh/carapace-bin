package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share snapshot properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_setCmd).Standalone()

	share_snapshot_setCmd.Flags().String("description", "", "Set a description to the snapshot.")
	share_snapshot_setCmd.Flags().String("name", "", "Set a name to the snapshot.")
	share_snapshot_setCmd.Flags().String("property", "", "Set a property to this snapshot (repeat option to set multiple properties)")
	share_snapshot_setCmd.Flags().String("status", "", "Assign a status to the snapshot (Admin only).")
	share_snapshotCmd.AddCommand(share_snapshot_setCmd)
}
