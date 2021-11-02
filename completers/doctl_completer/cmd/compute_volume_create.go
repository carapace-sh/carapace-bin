package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_volume_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a block storage volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volume_createCmd).Standalone()
	compute_volume_createCmd.Flags().String("desc", "", "Volume description")
	compute_volume_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Size`, `Region`, `Filesystem Type`, `Filesystem Label`, `DropletIDs`, `Tags`")
	compute_volume_createCmd.Flags().String("fs-label", "", "Volume filesystem label")
	compute_volume_createCmd.Flags().String("fs-type", "", "Volume filesystem type (ext4 or xfs)")
	compute_volume_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_volume_createCmd.Flags().String("region", "", "Volume region; should not be specified with a snapshot")
	compute_volume_createCmd.Flags().String("size", "4TiB", "Volume size (required)")
	compute_volume_createCmd.Flags().String("snapshot", "", "Volume snapshot; should not be specified with a region")
	compute_volume_createCmd.Flags().StringSlice("tag", []string{}, "Tags to apply to the volume; comma separate or repeat `--tag` to add multiple tags at once")
	compute_volumeCmd.AddCommand(compute_volume_createCmd)

	carapace.Gen(compute_volume_createCmd).FlagCompletion(carapace.ActionMap{"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		return carapace.ActionValues("ID", "Name", "Size", "Region", "Filesystem Type", "Filesystem Label", "DropletIDs", "Tags").Invoke(c).Filter(c.Parts).ToA()
	}),
		"fs-type": carapace.ActionValues("ext4", "xfs"),
		"region":  action.ActionRegions(),
	})
}
