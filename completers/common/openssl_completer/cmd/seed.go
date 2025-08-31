package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:                "seed",
	Short:              "SEED Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"seed-cbc", "seed-cfb", "seed-ecb", "seed-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(seedCmd).Standalone()

	rootCmd.AddCommand(seedCmd)

	carapace.Gen(seedCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
