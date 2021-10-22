package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volume_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve an existing block storage volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volume_getCmd).Standalone()
	compute_volume_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Size`, `Region`, `Filesystem Type`, `Filesystem Label`, `DropletIDs`, `Tags`")
	compute_volume_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_volumeCmd.AddCommand(compute_volume_getCmd)

	carapace.Gen(compute_volume_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Size", "Region", "Filesystem Type", "Filesystem Label", "DropletIDs", "Tags").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
