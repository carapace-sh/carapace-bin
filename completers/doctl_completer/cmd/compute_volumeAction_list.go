package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volumeAction_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of actions taken on a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volumeAction_listCmd).Standalone()
	compute_volumeAction_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_volumeAction_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_volumeActionCmd.AddCommand(compute_volumeAction_listCmd)

	carapace.Gen(compute_volumeAction_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
