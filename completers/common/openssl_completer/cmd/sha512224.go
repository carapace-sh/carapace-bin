package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sha512224Cmd = &cobra.Command{
	Use:                "sha512-224",
	Short:              "SHA-2 512/224 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sha512224Cmd).Standalone()

	rootCmd.AddCommand(sha512224Cmd)

	carapace.Gen(sha512224Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
