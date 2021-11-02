package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_floatingIpAction_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve the status of a floating IP action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIpAction_getCmd).Standalone()
	compute_floatingIpAction_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_floatingIpAction_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_floatingIpActionCmd.AddCommand(compute_floatingIpAction_getCmd)

	carapace.Gen(compute_floatingIpAction_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
