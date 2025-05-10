package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/newrelic"
	"github.com/spf13/cobra"
)

var workload_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a New Relic One workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workload_updateCmd).Standalone()
	workload_updateCmd.Flags().StringSliceP("entityGuid", "e", nil, "the list of entity Guids composing the workload")
	workload_updateCmd.Flags().StringSliceP("entitySearchQuery", "q", nil, "a list of search queries, combined using an OR operator")
	workload_updateCmd.Flags().StringP("guid", "g", "", "the GUID of the workload you want to update")
	workload_updateCmd.Flags().StringP("name", "n", "", "the name of the workload")
	workload_updateCmd.Flags().IntSliceP("scopeAccountIds", "s", nil, "accounts that will be used to get entities from")
	workloadCmd.AddCommand(workload_updateCmd)

	// TODO guid completion
	carapace.Gen(workload_updateCmd).FlagCompletion(carapace.ActionMap{
		"scopeAccountIds": newrelic.ActionAccountIds().UniqueList(","),
	})
}
