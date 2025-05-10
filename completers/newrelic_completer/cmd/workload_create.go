package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/newrelic"
	"github.com/spf13/cobra"
)

var workload_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a New Relic One workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workload_createCmd).Standalone()
	workload_createCmd.Flags().StringSliceP("entityGuid", "e", nil, "the list of entity Guids composing the workload")
	workload_createCmd.Flags().StringSliceP("entitySearchQuery", "q", nil, "a list of search queries, combined using an OR operator")
	workload_createCmd.Flags().StringP("name", "n", "", "the name of the workload")
	workload_createCmd.Flags().IntSliceP("scopeAccountIds", "s", []int{}, "accounts that will be used to get entities from")
	workloadCmd.AddCommand(workload_createCmd)

	// TODO guid completion
	carapace.Gen(workload_createCmd).FlagCompletion(carapace.ActionMap{
		"scopeAccountIds": newrelic.ActionAccountIds().UniqueList(","),
	})
}
