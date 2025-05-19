package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduleWorkCmd = &cobra.Command{
	Use:   "schedule:work [--run-output-file [RUN-OUTPUT-FILE]]",
	Short: "Start the schedule worker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleWorkCmd).Standalone()

	scheduleWorkCmd.Flags().String("run-output-file", "", "The file to direct <info>schedule:run</info> output to")
	rootCmd.AddCommand(scheduleWorkCmd)
}
