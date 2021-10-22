package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "List all snapshots for a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_snapshotsCmd).Standalone()
	compute_droplet_snapshotsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Type`, `Distribution`, `Slug`, `Public`, `MinDisk`")
	compute_droplet_snapshotsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletCmd.AddCommand(compute_droplet_snapshotsCmd)

	carapace.Gen(compute_droplet_snapshotsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
