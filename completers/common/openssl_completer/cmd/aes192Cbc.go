package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var aes192CbcCmd = &cobra.Command{
	Use:                "aes-192-cbc",
	Short:              "AES-192 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"aes-192-ecb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(aes192CbcCmd).Standalone()

	rootCmd.AddCommand(aes192CbcCmd)

	carapace.Gen(aes192CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
