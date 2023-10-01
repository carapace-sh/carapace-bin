package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "carapace",
	Short:              "multi-shell multi-command argument completer",
	Long:               "https://github.com/rsteube/carapace-bin",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapace("carapace-new"), // TODO replace with carapace
	)
}
