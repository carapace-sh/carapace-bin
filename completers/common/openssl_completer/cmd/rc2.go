package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rc2Cmd = &cobra.Command{
	Use:                "rc2",
	Short:              "RC2 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"rc2-40-cbc", "rc2-64-cbc", "rc2-cbc", "rc2-cfb", "rc2-ecb", "rc2-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(rc2Cmd).Standalone()

	rootCmd.AddCommand(rc2Cmd)

	carapace.Gen(rc2Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
