package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var shake128Cmd = &cobra.Command{
	Use:                "shake128",
	Short:              "SHA-3 SHAKE128 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(shake128Cmd).Standalone()

	rootCmd.AddCommand(shake128Cmd)

	carapace.Gen(shake128Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
