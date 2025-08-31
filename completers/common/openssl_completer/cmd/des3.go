package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var des3Cmd = &cobra.Command{
	Use:                "des3",
	Short:              "Triple-DES Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"desx", "des-ede3", "des-ede3-cbc", "des-ede3-cfb", "des-ede3-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(des3Cmd).Standalone()

	rootCmd.AddCommand(des3Cmd)

	carapace.Gen(des3Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
