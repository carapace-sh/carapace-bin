package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_action_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a  list of all recent actions taken on your resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_action_listCmd).Standalone()
	compute_action_listCmd.Flags().String("action-type", "", "Action type")
	compute_action_listCmd.Flags().String("after", "", "Action completed after in RFC3339 format")
	compute_action_listCmd.Flags().String("before", "", "Action completed before in RFC3339 format")
	compute_action_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_action_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_action_listCmd.Flags().String("region", "", "Action region")
	compute_action_listCmd.Flags().String("resource-type", "", "Action resource type")
	compute_action_listCmd.Flags().String("status", "", "Action status")
	compute_actionCmd.AddCommand(compute_action_listCmd)

	// TODO flags completion
	carapace.Gen(compute_action_listCmd).FlagCompletion(carapace.ActionMap{
		"action-type": action.ActionActionTypes(),
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region": action.ActionRegions(),
		"status": action.ActionActionStatus(),
	})
}
