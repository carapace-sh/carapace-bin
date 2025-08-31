package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var blake2b512Cmd = &cobra.Command{
	Use:                "blake2b512",
	Short:              "BLAKE2b-512 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(blake2b512Cmd).Standalone()

	rootCmd.AddCommand(blake2b512Cmd)

	carapace.Gen(blake2b512Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
