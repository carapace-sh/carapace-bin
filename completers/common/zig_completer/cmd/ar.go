package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var arCmd = &cobra.Command{
	Use:                "ar",
	Short:              "Combine object files into static archive",
	GroupID:            "wrapper",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(arCmd).Standalone()

	rootCmd.AddCommand(arCmd)

	carapace.Gen(arCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("ar"),
	)
}
