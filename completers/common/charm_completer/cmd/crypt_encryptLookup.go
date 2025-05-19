package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var crypt_encryptLookupCmd = &cobra.Command{
	Use:   "encrypt-lookup",
	Short: "Encrypt arg deterministically",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crypt_encryptLookupCmd).Standalone()

	cryptCmd.AddCommand(crypt_encryptLookupCmd)
}
