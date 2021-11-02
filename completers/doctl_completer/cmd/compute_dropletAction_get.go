package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletAction_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a specific Droplet action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletAction_getCmd).Standalone()
	compute_dropletAction_getCmd.Flags().Int("action-id", 0, "Action ID (required)")
	compute_dropletAction_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_dropletAction_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletActionCmd.AddCommand(compute_dropletAction_getCmd)

	carapace.Gen(compute_dropletAction_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
