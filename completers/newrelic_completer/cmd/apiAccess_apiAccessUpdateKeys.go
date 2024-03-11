package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiAccess_apiAccessUpdateKeysCmd = &cobra.Command{
	Use:   "apiAccessUpdateKeys",
	Short: "Update keys. You can update keys for multiple accounts at once. You can read more about managing keys on [this documentation page](https://docs.newrelic.com/docs/apis/nerdgraph/examples/use-nerdgraph-manage-license-keys-personal-api-keys).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiAccess_apiAccessUpdateKeysCmd).Standalone()
	apiAccess_apiAccessUpdateKeysCmd.Flags().String("keys", "", "The configurations of each key you want to update.")
	apiAccessCmd.AddCommand(apiAccess_apiAccessUpdateKeysCmd)
}
