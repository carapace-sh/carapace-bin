package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var cast5CbcCmd = &cobra.Command{
	Use:                "cast5-cbc",
	Short:              "CAST5 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"cast5-cfb", "cast5-ecb", "cast5-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(cast5CbcCmd).Standalone()

	rootCmd.AddCommand(cast5CbcCmd)

	carapace.Gen(cast5CbcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
