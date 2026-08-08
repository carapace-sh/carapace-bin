package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var schedule_createCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a new pipeline schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schedule_createCmd).Standalone()

	schedule_createCmd.Flags().Bool("active", false, "Whether or not the schedule is active.")
	schedule_createCmd.Flags().String("cron", "", "Cron interval pattern.")
	schedule_createCmd.Flags().String("cronTimeZone", "", "Cron timezone.")
	schedule_createCmd.Flags().String("description", "", "Description of the schedule.")
	schedule_createCmd.Flags().String("ref", "", "Target branch or tag.")
	schedule_createCmd.Flags().StringSlice("variable", nil, "Pass variables to schedule in the format <key>:<value>. Repeat flag for multiple variables.")
	schedule_createCmd.MarkFlagRequired("cron")
	schedule_createCmd.MarkFlagRequired("description")
	schedule_createCmd.MarkFlagRequired("ref")
	scheduleCmd.AddCommand(schedule_createCmd)

	carapace.Gen(schedule_createCmd).FlagCompletion(carapace.ActionMap{
		"ref": carapace.Batch(
			action.ActionBranches(schedule_createCmd),
			action.ActionTags(schedule_createCmd),
		).ToA(),
	})
}
