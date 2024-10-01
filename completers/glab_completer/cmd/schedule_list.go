package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schedule_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "Get the list of schedules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schedule_listCmd).Standalone()

	schedule_listCmd.Flags().StringP("page", "p", "", "Page number.")
	schedule_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	scheduleCmd.AddCommand(schedule_listCmd)
}
