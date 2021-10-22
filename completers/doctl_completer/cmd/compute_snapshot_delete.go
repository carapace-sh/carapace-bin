package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot of a Droplet or volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_snapshot_deleteCmd).Standalone()
	compute_snapshot_deleteCmd.Flags().BoolP("force", "f", false, "Delete the snapshot without confirmation")
	compute_snapshot_deleteCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `CreatedAt`, `Regions`, `ResourceId`, `ResourceType`, `MinDiskSize`, `Size`, `Tags`")
	compute_snapshot_deleteCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_snapshotCmd.AddCommand(compute_snapshot_deleteCmd)

	carapace.Gen(compute_snapshot_deleteCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "CreatedAt", "Regions", "ResourceId", "ResourceType", "MinDiskSize", "Size", "Tags").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
