package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduleInterruptCmd = &cobra.Command{
	Use:   "schedule:interrupt",
	Short: "Interrupt the current schedule run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleInterruptCmd).Standalone()

	rootCmd.AddCommand(scheduleInterruptCmd)
}
