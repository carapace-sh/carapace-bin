package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schedule_runCmd = &cobra.Command{
	Use:   "run <id>",
	Short: "Trigger a pipeline schedule to run immediately.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schedule_runCmd).Standalone()

	scheduleCmd.AddCommand(schedule_runCmd)

	// TODO positional completion
}
