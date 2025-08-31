package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var camellia128CbcCmd = &cobra.Command{
	Use:                "camellia-128-cbc",
	Short:              "Camellia-128 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"camellia-128-ecb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(camellia128CbcCmd).Standalone()

	rootCmd.AddCommand(camellia128CbcCmd)

	carapace.Gen(camellia128CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
