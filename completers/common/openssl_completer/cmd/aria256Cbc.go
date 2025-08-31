package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var aria256CbcCmd = &cobra.Command{
	Use:                "aria-256-cbc",
	Short:              "Aria-256 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"aria-256-cfb", "aria-256-cfb1", "aria-256-cfb8", "aria-256-ctr", "aria-256-ecb", "aria-256-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(aria256CbcCmd).Standalone()

	rootCmd.AddCommand(aria256CbcCmd)

	carapace.Gen(aria256CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
