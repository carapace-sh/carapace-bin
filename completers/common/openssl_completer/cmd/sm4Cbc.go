package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sm4CbcCmd = &cobra.Command{
	Use:                "sm4-cbc",
	Short:              "SM4 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"sm4-cfb", "sm4-ctr", "sm4-ecb", "sm4-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sm4CbcCmd).Standalone()

	rootCmd.AddCommand(sm4CbcCmd)

	carapace.Gen(sm4CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
