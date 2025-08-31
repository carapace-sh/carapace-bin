package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var desCmd = &cobra.Command{
	Use:                "des",
	Short:              "DES Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"des-cbc", "des-cfb", "des-ecb", "des-ede", "des-ede-cbc", "des-ede-cfb", "des-ede-ofb", "des-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(desCmd).Standalone()

	rootCmd.AddCommand(desCmd)

	carapace.Gen(desCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
