package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rmd160Cmd = &cobra.Command{
	Use:                "rmd160",
	Short:              "RMD-160 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(rmd160Cmd).Standalone()

	rootCmd.AddCommand(rmd160Cmd)

	carapace.Gen(rmd160Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
