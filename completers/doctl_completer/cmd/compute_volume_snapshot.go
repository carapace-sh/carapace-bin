package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volume_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create a block storage volume snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volume_snapshotCmd).Standalone()
	compute_volume_snapshotCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Size`, `Region`, `Filesystem Type`, `Filesystem Label`, `DropletIDs`, `Tags`")
	compute_volume_snapshotCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_volume_snapshotCmd.Flags().String("snapshot-desc", "", "Snapshot description")
	compute_volume_snapshotCmd.Flags().String("snapshot-name", "", "Snapshot name (required)")
	compute_volume_snapshotCmd.Flags().StringSlice("tag", []string{}, "Tags to apply to the snapshot; comma separate or repeat `--tag` to add multiple tags at once")
	compute_volumeCmd.AddCommand(compute_volume_snapshotCmd)

	carapace.Gen(compute_volume_snapshotCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Size", "Region", "Filesystem Type", "Filesystem Label", "DropletIDs", "Tags").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
