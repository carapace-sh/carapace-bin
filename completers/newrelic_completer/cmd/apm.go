package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/newrelic"
	"github.com/spf13/cobra"
)

var apmCmd = &cobra.Command{
	Use:   "apm",
	Short: "Interact with New Relic APM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apmCmd).Standalone()
	apmCmd.PersistentFlags().Int("applicationId", 0, "A New Relic APM application ID")
	rootCmd.AddCommand(apmCmd)

	carapace.Gen(apmCmd).FlagCompletion(carapace.ActionMap{
		"applicationId": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return newrelic.ActionApplicationIds(apmCmd.Flag("profile").Value.String())
		}),
	})
}
