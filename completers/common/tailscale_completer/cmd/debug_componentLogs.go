package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_componentLogsCmd = &cobra.Command{
	Use:   "component-logs",
	Short: "Enable/disable debug logs for a component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_componentLogsCmd).Standalone()

	debug_componentLogsCmd.Flags().String("for", "", "duration to enable debug logs for")
	debugCmd.AddCommand(debug_componentLogsCmd)
}
