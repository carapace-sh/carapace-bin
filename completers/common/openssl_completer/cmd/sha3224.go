package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sha3224Cmd = &cobra.Command{
	Use:                "sha3-224",
	Short:              "SHA-3 224 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sha3224Cmd).Standalone()

	rootCmd.AddCommand(sha3224Cmd)

	carapace.Gen(sha3224Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
