package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var camellia256CbcCmd = &cobra.Command{
	Use:                "camellia-256-cbc",
	Short:              "Camellia-256 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"camellia-256-ecb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(camellia256CbcCmd).Standalone()

	rootCmd.AddCommand(camellia256CbcCmd)

	carapace.Gen(camellia256CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
