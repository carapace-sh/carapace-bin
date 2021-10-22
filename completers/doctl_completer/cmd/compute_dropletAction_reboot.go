package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletAction_rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletAction_rebootCmd).Standalone()
	compute_dropletAction_rebootCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_dropletAction_rebootCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletAction_rebootCmd.Flags().Bool("wait", false, "Wait for action to complete")
	compute_dropletActionCmd.AddCommand(compute_dropletAction_rebootCmd)

	carapace.Gen(compute_dropletAction_rebootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
