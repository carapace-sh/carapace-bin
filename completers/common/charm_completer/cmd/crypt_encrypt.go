package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var crypt_encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt stdin with your Charm account encryption key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crypt_encryptCmd).Standalone()

	cryptCmd.AddCommand(crypt_encryptCmd)
}
