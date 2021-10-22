package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_action_waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Block thread until an action completes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_action_waitCmd).Standalone()
	compute_action_waitCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_action_waitCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_action_waitCmd.Flags().Int("poll-timeout", 5, "Re-poll time in seconds")
	compute_actionCmd.AddCommand(compute_action_waitCmd)

	carapace.Gen(compute_action_waitCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
