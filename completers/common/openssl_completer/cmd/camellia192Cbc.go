package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var camellia192CbcCmd = &cobra.Command{
	Use:                "camellia-192-cbc",
	Short:              "Camellia-192 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"camellia-192-ecb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(camellia192CbcCmd).Standalone()

	rootCmd.AddCommand(camellia192CbcCmd)

	carapace.Gen(camellia192CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
