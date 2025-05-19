package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiAccess_apiAccessCreateKeysCmd = &cobra.Command{
	Use:   "apiAccessCreateKeys",
	Short: "Create keys. You can create keys for multiple accounts at once. You can read more about managing keys on [this documentation page](https://docs.newrelic.com/docs/apis/nerdgraph/examples/use-nerdgraph-manage-license-keys-personal-api-keys).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiAccess_apiAccessCreateKeysCmd).Standalone()
	apiAccess_apiAccessCreateKeysCmd.Flags().String("keys", "", "A list of the configurations for each key you want to create.")
	apiAccessCmd.AddCommand(apiAccess_apiAccessCreateKeysCmd)
}
