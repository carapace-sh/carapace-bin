package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiAccess_apiAccessGetKeyCmd = &cobra.Command{
	Use:   "apiAccessGetKey",
	Short: "Fetch a single key by ID and type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiAccess_apiAccessGetKeyCmd).Standalone()
	apiAccess_apiAccessGetKeyCmd.Flags().String("id", "", "The `id` of the key. This can be used to identify a key without revealing the key itself (used to update and delete).")
	apiAccess_apiAccessGetKeyCmd.Flags().String("keyType", "", "The type of key.")
	apiAccessCmd.AddCommand(apiAccess_apiAccessGetKeyCmd)

	// TODO keytype
}
