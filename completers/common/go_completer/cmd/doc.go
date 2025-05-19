package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:                "doc",
	Short:              "show documentation for package or symbol",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(docCmd).Standalone()

	rootCmd.AddCommand(docCmd)

	carapace.Gen(docCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("go-tool-doc"),
	)
}
