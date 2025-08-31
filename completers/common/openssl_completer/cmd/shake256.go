package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var shake256Cmd = &cobra.Command{
	Use:                "shake256",
	Short:              "SHA-3 SHAKE256 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(shake256Cmd).Standalone()

	rootCmd.AddCommand(shake256Cmd)

	carapace.Gen(shake256Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
