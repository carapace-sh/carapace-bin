package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:                "base64",
	Short:              "Base64 Encoding",
	GroupID:            "cipher",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(base64Cmd).Standalone()

	rootCmd.AddCommand(base64Cmd)

	carapace.Gen(base64Cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
