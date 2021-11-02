package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var monitoring_alert_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all alert policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitoring_alert_listCmd).Standalone()
	monitoring_alert_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `UUID`, `Type`, `Description`, `Compare`, `Value`, `Window`, `Entities`, `Tags`, `Emails`, `Slack Channels`, `Enabled`")
	monitoring_alert_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	monitoring_alertCmd.AddCommand(monitoring_alert_listCmd)
}
