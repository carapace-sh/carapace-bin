package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var blake2s256Cmd = &cobra.Command{
	Use:                "blake2s256",
	Short:              "BLAKE2s-256 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(blake2s256Cmd).Standalone()

	rootCmd.AddCommand(blake2s256Cmd)

	carapace.Gen(blake2s256Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
