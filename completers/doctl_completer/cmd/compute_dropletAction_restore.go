package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletAction_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a Droplet from a backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletAction_restoreCmd).Standalone()
	compute_dropletAction_restoreCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_dropletAction_restoreCmd.Flags().Int("image-id", 0, "Image ID (required)")
	compute_dropletAction_restoreCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletAction_restoreCmd.Flags().Bool("wait", false, "Wait for action to complete")
	compute_dropletActionCmd.AddCommand(compute_dropletAction_restoreCmd)

	carapace.Gen(compute_dropletAction_restoreCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
