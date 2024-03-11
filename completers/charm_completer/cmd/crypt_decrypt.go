package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var crypt_decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt stdin with your Charm account encryption key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crypt_decryptCmd).Standalone()

	cryptCmd.AddCommand(crypt_decryptCmd)
}
