package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiAccess_apiAccessDeleteKeysCmd = &cobra.Command{
	Use:   "apiAccessDeleteKeys",
	Short: "A mutation to delete keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiAccess_apiAccessDeleteKeysCmd).Standalone()
	apiAccess_apiAccessDeleteKeysCmd.Flags().String("keys", "", "A list of each key `id` that you want to delete. You can read more about managing keys on [this documentation page](https://docs.newrelic.com/docs/apis/nerdgraph/examples/use-nerdgraph-manage-license-keys-personal-api-keys).")
	apiAccessCmd.AddCommand(apiAccess_apiAccessDeleteKeysCmd)
}
