package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_listAlertsCmd = &cobra.Command{
	Use:   "list-alerts",
	Short: "List alerts on an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_listAlertsCmd).Standalone()
	apps_listAlertsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Spec.Rule`, `Trigger`, `ComponentName`, `Emails`, `SlackWebhooks`, `Spec.Disabled`")
	apps_listAlertsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	appsCmd.AddCommand(apps_listAlertsCmd)

	carapace.Gen(apps_listAlertsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Spec.Rule", "Trigger", "ComponentName", "Emails", "SlackWebhooks", "Spec.Disabled").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO positional
}
