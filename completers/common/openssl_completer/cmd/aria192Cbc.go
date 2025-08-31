package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var aria192CbcCmd = &cobra.Command{
	Use:                "aria-192-cbc",
	Short:              "Aria-192 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"aria-192-cfb", "aria-192-cfb1", "aria-192-cfb8", "aria-192-ctr", "aria-192-ecb", "aria-192-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(aria192CbcCmd).Standalone()

	rootCmd.AddCommand(aria192CbcCmd)

	carapace.Gen(aria192CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
