package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_floatingIpAction_assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign a floating IP address to a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIpAction_assignCmd).Standalone()
	compute_floatingIpAction_assignCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_floatingIpAction_assignCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_floatingIpActionCmd.AddCommand(compute_floatingIpAction_assignCmd)

	carapace.Gen(compute_floatingIpAction_assignCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
