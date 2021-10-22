package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletAction_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Take a Droplet snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletAction_snapshotCmd).Standalone()
	compute_dropletAction_snapshotCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_dropletAction_snapshotCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletAction_snapshotCmd.Flags().String("snapshot-name", "", "Snapshot name (required)")
	compute_dropletAction_snapshotCmd.Flags().Bool("wait", false, "Wait for action to complete")
	compute_dropletActionCmd.AddCommand(compute_dropletAction_snapshotCmd)

	carapace.Gen(compute_dropletAction_snapshotCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
