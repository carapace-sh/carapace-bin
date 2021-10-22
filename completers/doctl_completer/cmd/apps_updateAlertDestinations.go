package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_updateAlertDestinationsCmd = &cobra.Command{
	Use:   "update-alert-destinations",
	Short: "Update alert destinations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_updateAlertDestinationsCmd).Standalone()
	apps_updateAlertDestinationsCmd.Flags().String("app-alert-destinations", "", "Path to an alert destinations file in JSON or YAML format.")
	apps_updateAlertDestinationsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Spec.Rule`, `Trigger`, `ComponentName`, `Emails`, `SlackWebhooks`, `Spec.Disabled`")
	apps_updateAlertDestinationsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	appsCmd.AddCommand(apps_updateAlertDestinationsCmd)

	carapace.Gen(apps_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Spec.Rule", "Trigger", "ComponentName", "Emails", "SlackWebhooks", "Spec.Disabled").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO positional
}
