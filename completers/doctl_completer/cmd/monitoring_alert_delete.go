package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var monitoring_alert_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an alert policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitoring_alert_deleteCmd).Standalone()
	monitoring_alert_deleteCmd.Flags().BoolP("force", "f", false, "Delete an alert policy without confirmation prompt")
	monitoring_alertCmd.AddCommand(monitoring_alert_deleteCmd)
}
