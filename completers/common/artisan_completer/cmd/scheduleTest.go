package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduleTestCmd = &cobra.Command{
	Use:   "schedule:test [--name [NAME]]",
	Short: "Run a scheduled command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleTestCmd).Standalone()

	scheduleTestCmd.Flags().String("name", "", "The name of the scheduled command to run")
	rootCmd.AddCommand(scheduleTestCmd)
}
