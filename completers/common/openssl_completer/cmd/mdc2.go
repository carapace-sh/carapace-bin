package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var mdc2Cmd = &cobra.Command{
	Use:                "mdc2",
	Short:              "MDC2 Digest",
	GroupID:            "digest",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(mdc2Cmd).Standalone()

	rootCmd.AddCommand(mdc2Cmd)

	carapace.Gen(mdc2Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "dgst"),
	)
}
