package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var aes128CbcCmd = &cobra.Command{
	Use:                "aes-128-cbc",
	Short:              "AES-128 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"aes-128-ecb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(aes128CbcCmd).Standalone()

	rootCmd.AddCommand(aes128CbcCmd)

	carapace.Gen(aes128CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
