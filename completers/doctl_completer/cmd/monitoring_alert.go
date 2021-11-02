package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var monitoring_alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Display commands for managing alert policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitoring_alertCmd).Standalone()
	monitoringCmd.AddCommand(monitoring_alertCmd)
}
