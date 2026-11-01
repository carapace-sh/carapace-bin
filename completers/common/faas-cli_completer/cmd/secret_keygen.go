package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generate a keypair for sealing build secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_keygenCmd).Standalone()

	secret_keygenCmd.Flags().StringP("output", "o", "key", "Output path for the private key (public key gets .pub appended)")
	secretCmd.AddCommand(secret_keygenCmd)

	carapace.Gen(secret_keygenCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
