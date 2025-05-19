package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listUnitsCmd = &cobra.Command{
	Use:     "list-units",
	Short:   "List units currently in memory",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listUnitsCmd).Standalone()

	rootCmd.AddCommand(listUnitsCmd)

	carapace.Gen(listUnitsCmd).PositionalAnyCompletion(
		action.ActionUnits(listUnitsCmd).FilterArgs(),
	)
}
