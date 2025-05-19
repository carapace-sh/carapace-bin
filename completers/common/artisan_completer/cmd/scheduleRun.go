package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduleRunCmd = &cobra.Command{
	Use:   "schedule:run",
	Short: "Run the scheduled commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleRunCmd).Standalone()

	rootCmd.AddCommand(scheduleRunCmd)
}
