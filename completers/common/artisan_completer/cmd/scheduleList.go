package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduleListCmd = &cobra.Command{
	Use:   "schedule:list [--timezone [TIMEZONE]] [--next]",
	Short: "List all scheduled tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleListCmd).Standalone()

	scheduleListCmd.Flags().Bool("next", false, "Sort the listed tasks by their next due date")
	scheduleListCmd.Flags().String("timezone", "", "The timezone that times should be displayed in")
	rootCmd.AddCommand(scheduleListCmd)
}
