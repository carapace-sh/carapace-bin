package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var castCmd = &cobra.Command{
	Use:                "cast",
	Short:              "CAST Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"cast-cbc"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(castCmd).Standalone()

	rootCmd.AddCommand(castCmd)

	carapace.Gen(castCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
