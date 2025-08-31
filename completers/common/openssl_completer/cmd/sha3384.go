package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sha3384Cmd = &cobra.Command{
	Use:                "sha3-384",
	Short:              "SHA-3 384 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(sha3384Cmd).Standalone()

	rootCmd.AddCommand(sha3384Cmd)

	carapace.Gen(sha3384Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
