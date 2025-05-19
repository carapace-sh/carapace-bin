package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_monitor_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List New Relic Synthetics monitors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_monitor_listCmd).Standalone()
	synthetics_monitor_listCmd.Flags().StringP("statusFilter", "s", "", "filter the results on the status field. Possible values ENABLED, DISABLED, MUTED. Comma separated.")
	synthetics_monitorCmd.AddCommand(synthetics_monitor_listCmd)

	carapace.Gen(synthetics_monitor_listCmd).FlagCompletion(carapace.ActionMap{
		"statusFilter": carapace.ActionValues("ENABLED", "DISABLED", "MUTED").UniqueList(","),
	})
}
