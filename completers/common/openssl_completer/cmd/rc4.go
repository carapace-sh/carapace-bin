package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rc4Cmd = &cobra.Command{
	Use:                "rc4",
	Short:              "RC4 Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"rc4-40"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(rc4Cmd).Standalone()

	rootCmd.AddCommand(rc4Cmd)

	carapace.Gen(rc4Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
