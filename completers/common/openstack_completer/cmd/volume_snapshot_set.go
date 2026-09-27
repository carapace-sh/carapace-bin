package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_snapshot_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set volume snapshot properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_snapshot_setCmd).Standalone()

	volume_snapshot_setCmd.Flags().String("description", "", "New snapshot description")
	volume_snapshot_setCmd.Flags().String("name", "", "New snapshot name")
	volume_snapshot_setCmd.Flags().Bool("no-property", false, "Remove all properties from <snapshot> (specify both --no-property and --property to remove the current properties before setting new properties.)")
	volume_snapshot_setCmd.Flags().String("property", "", "Property to add/change for this snapshot (repeat option to set multiple properties)")
	volume_snapshot_setCmd.Flags().String("state", "", "New snapshot state.")
	volume_snapshotCmd.AddCommand(volume_snapshot_setCmd)
}
