package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sha3512Cmd = &cobra.Command{
	Use:                "sha3-512",
	Short:              "SHA-3 512 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sha3512Cmd).Standalone()

	rootCmd.AddCommand(sha3512Cmd)

	carapace.Gen(sha3512Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
