package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Validate repetitive calendar time events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(calendarCmd).Standalone()

	rootCmd.AddCommand(calendarCmd)
}
