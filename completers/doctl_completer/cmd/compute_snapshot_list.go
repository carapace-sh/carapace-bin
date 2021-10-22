package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Droplet and volume snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_snapshot_listCmd).Standalone()
	compute_snapshot_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `CreatedAt`, `Regions`, `ResourceId`, `ResourceType`, `MinDiskSize`, `Size`, `Tags`")
	compute_snapshot_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_snapshot_listCmd.Flags().String("region", "", "Filter by regional availability")
	compute_snapshot_listCmd.Flags().String("resource", "", "Filter by resource type (`droplet` or `volume`)")
	compute_snapshotCmd.AddCommand(compute_snapshot_listCmd)

	carapace.Gen(compute_snapshot_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "CreatedAt", "Regions", "ResourceId", "ResourceType", "MinDiskSize", "Size", "Tags").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region":   action.ActionRegions(),
		"resource": carapace.ActionValues("droplet", "volume"),
	})
}
