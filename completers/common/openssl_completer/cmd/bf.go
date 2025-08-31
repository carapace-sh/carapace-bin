package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var bfCmd = &cobra.Command{
	Use:                "bf",
	Short:              "Blowfish Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"bf-cbc", "bf-cfb", "bf-ecb", "bf-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(bfCmd).Standalone()

	rootCmd.AddCommand(bfCmd)

	carapace.Gen(bfCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
