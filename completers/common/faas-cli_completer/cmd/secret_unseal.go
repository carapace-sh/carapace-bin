package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_unsealCmd = &cobra.Command{
	Use:   "unseal",
	Short: "Unseal and inspect a sealed secrets file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_unsealCmd).Standalone()

	secret_unsealCmd.Flags().String("in", "com.openfaas.secrets", "Path to the sealed secrets file")
	secret_unsealCmd.Flags().BoolP("json", "j", false, "Output secrets as JSON")
	secret_unsealCmd.Flags().String("key", "", "Unseal a single key (omit to print all)")
	secretCmd.AddCommand(secret_unsealCmd)

	carapace.Gen(secret_unsealCmd).FlagCompletion(carapace.ActionMap{
		"in": carapace.ActionFiles(),
	})
}
