package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var aes256CbcCmd = &cobra.Command{
	Use:                "aes-256-cbc",
	Short:              "AES-256 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"aes-256-ecb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(aes256CbcCmd).Standalone()

	rootCmd.AddCommand(aes256CbcCmd)

	carapace.Gen(aes256CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
