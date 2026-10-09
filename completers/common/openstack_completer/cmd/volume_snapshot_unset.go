package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_snapshot_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset volume snapshot properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_snapshot_unsetCmd).Standalone()

	volume_snapshot_unsetCmd.Flags().String("property", "", "Property to remove from snapshot (repeat option to remove multiple properties)")
	volume_snapshotCmd.AddCommand(volume_snapshot_unsetCmd)
}
