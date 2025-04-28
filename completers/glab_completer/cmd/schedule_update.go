package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schedule_updateCmd = &cobra.Command{
	Use:   "update <id> [flags]",
	Short: "Update a pipeline schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schedule_updateCmd).Standalone()

	schedule_updateCmd.Flags().Bool("active", false, "Whether or not the schedule is active.")
	schedule_updateCmd.Flags().StringSlice("create-variable", []string{}, "Pass new variables to schedule in format <key>:<value>.")
	schedule_updateCmd.Flags().String("cron", "", "Cron interval pattern.")
	schedule_updateCmd.Flags().String("cronTimeZone", "", "Cron timezone.")
	schedule_updateCmd.Flags().StringSlice("delete-variable", []string{}, "Pass variables you want to delete from schedule in format <key>.")
	schedule_updateCmd.Flags().String("description", "", "Description of the schedule.")
	schedule_updateCmd.Flags().String("ref", "", "Target branch or tag.")
	schedule_updateCmd.Flags().StringSlice("update-variable", []string{}, "Pass updated variables to schedule in format <key>:<value>.")
	scheduleCmd.AddCommand(schedule_updateCmd)

	// TODO add completions
}
