package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_snapshot_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a Droplet or volume snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_snapshot_getCmd).Standalone()
	compute_snapshot_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `CreatedAt`, `Regions`, `ResourceId`, `ResourceType`, `MinDiskSize`, `Size`, `Tags`")
	compute_snapshot_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_snapshotCmd.AddCommand(compute_snapshot_getCmd)

	carapace.Gen(compute_snapshot_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "CreatedAt", "Regions", "ResourceId", "ResourceType", "MinDiskSize", "Size", "Tags").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
