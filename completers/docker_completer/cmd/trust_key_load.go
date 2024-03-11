package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trust_key_loadCmd = &cobra.Command{
	Use:   "load [OPTIONS] KEYFILE",
	Short: "Load a private key file for signing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_key_loadCmd).Standalone()

	trust_key_loadCmd.Flags().String("name", "signer", "Name for the loaded key")
	trust_keyCmd.AddCommand(trust_key_loadCmd)

	carapace.Gen(trust_key_loadCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
