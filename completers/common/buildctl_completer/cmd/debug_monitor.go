package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "display build events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_monitorCmd).Standalone()

	debug_monitorCmd.Flags().Bool("completed", false, "show completed builds")
	debug_monitorCmd.Flags().String("ref", "", "show events for a specific build")
	debugCmd.AddCommand(debug_monitorCmd)
}
