package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var key_convertSecretToPublicCmd = &cobra.Command{
	Use:   "convert-secret-to-public",
	Short: "generate a public key for verifying store paths from a secret key read from standard input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(key_convertSecretToPublicCmd).Standalone()

	keyCmd.AddCommand(key_convertSecretToPublicCmd)
}
