package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sha512256Cmd = &cobra.Command{
	Use:                "sha512-256",
	Short:              "SHA-2 512/256 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sha512256Cmd).Standalone()

	rootCmd.AddCommand(sha512256Cmd)

	carapace.Gen(sha512256Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
