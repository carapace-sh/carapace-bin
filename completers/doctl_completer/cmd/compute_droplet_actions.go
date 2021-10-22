package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "List Droplet actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_actionsCmd).Standalone()
	compute_droplet_actionsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_droplet_actionsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletCmd.AddCommand(compute_droplet_actionsCmd)

	carapace.Gen(compute_droplet_actionsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
