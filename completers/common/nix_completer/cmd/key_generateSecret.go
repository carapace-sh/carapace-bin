package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var key_generateSecretCmd = &cobra.Command{
	Use:   "generate-secret",
	Short: "generate a secret key for signing store paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(key_generateSecretCmd).Standalone()

	key_generateSecretCmd.Flags().String("key-name", "", "Identifier of the key")
	keyCmd.AddCommand(key_generateSecretCmd)

	common.AddLoggingFlags(key_generateSecretCmd)
}
