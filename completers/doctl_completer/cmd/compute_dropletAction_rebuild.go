package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletAction_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletAction_rebuildCmd).Standalone()
	compute_dropletAction_rebuildCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_dropletAction_rebuildCmd.Flags().String("image", "", "Image ID or Slug (required)")
	compute_dropletAction_rebuildCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletAction_rebuildCmd.Flags().Bool("wait", false, "Wait for action to complete")
	compute_dropletActionCmd.AddCommand(compute_dropletAction_rebuildCmd)

	carapace.Gen(compute_dropletAction_rebuildCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
