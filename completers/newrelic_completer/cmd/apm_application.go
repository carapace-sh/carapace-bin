package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/newrelic"
	"github.com/spf13/cobra"
)

var apm_applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Interact with New Relic APM applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apm_applicationCmd).Standalone()
	apm_applicationCmd.PersistentFlags().StringP("guid", "g", "", "search for results matching the given APM application GUID")
	apmCmd.AddCommand(apm_applicationCmd)

	carapace.Gen(apm_applicationCmd).FlagCompletion(carapace.ActionMap{
		"guid": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return newrelic.ActionApplicationGuids(apm_applicationCmd.Flag("profile").Value.String())
		}),
	})
}
