package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var crypt_decryptLookupCmd = &cobra.Command{
	Use:   "decrypt-lookup",
	Short: "Decrypt arg deterministically",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crypt_decryptLookupCmd).Standalone()

	cryptCmd.AddCommand(crypt_decryptLookupCmd)
}
