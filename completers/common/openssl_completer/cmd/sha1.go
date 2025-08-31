package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sha1Cmd = &cobra.Command{
	Use:                "sha1",
	Short:              "SHA-1 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sha1Cmd).Standalone()

	rootCmd.AddCommand(sha1Cmd)

	carapace.Gen(sha1Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
