package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sha3256Cmd = &cobra.Command{
	Use:                "sha3-256",
	Short:              "SHA-3 256 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sha3256Cmd).Standalone()

	rootCmd.AddCommand(sha3256Cmd)

	carapace.Gen(sha3256Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
