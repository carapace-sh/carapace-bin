package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var schedule_deleteCmd = &cobra.Command{
	Use:   "delete [flags]",
	Short: "Delete schedule with the specified ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schedule_deleteCmd).Standalone()

	scheduleCmd.AddCommand(schedule_deleteCmd)

	// TODO positional completion
}
