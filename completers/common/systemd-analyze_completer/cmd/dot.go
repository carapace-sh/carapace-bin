package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var dotCmd = &cobra.Command{
	Use:   "dot",
	Short: "Output dependency graph in dot(1) format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotCmd).Standalone()

	rootCmd.AddCommand(dotCmd)

	carapace.Gen(dotCmd).PositionalAnyCompletion(
		action.ActionUnits(dotCmd).FilterArgs(),
	)
}
