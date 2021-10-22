package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_volume_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List block storage volumes by ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volume_listCmd).Standalone()
	compute_volume_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Size`, `Region`, `Filesystem Type`, `Filesystem Label`, `DropletIDs`, `Tags`")
	compute_volume_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_volume_listCmd.Flags().String("region", "", "Volume region")
	compute_volumeCmd.AddCommand(compute_volume_listCmd)

	carapace.Gen(compute_volume_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Size", "Region", "Filesystem Type", "Filesystem Label", "DropletIDs", "Tags").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region": action.ActionRegions(),
	})
}
