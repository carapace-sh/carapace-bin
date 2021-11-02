package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var monitoring_alert_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about an alert policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitoring_alert_getCmd).Standalone()
	monitoring_alert_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `UUID`, `Type`, `Description`, `Compare`, `Value`, `Window`, `Entities`, `Tags`, `Emails`, `Slack Channels`, `Enabled`")
	monitoring_alert_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	monitoring_alertCmd.AddCommand(monitoring_alert_getCmd)
}
