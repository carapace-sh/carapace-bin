package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var aria128CbcCmd = &cobra.Command{
	Use:                "aria-128-cbc",
	Short:              "Aria-128 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"aria-128-cfb", "aria-128-cfb1", "aria-128-cf8", "aria-128-ctr", "aria-128-ecb", "aria-128-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(aria128CbcCmd).Standalone()

	rootCmd.AddCommand(aria128CbcCmd)

	carapace.Gen(aria128CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
