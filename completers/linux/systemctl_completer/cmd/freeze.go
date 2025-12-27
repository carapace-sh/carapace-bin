package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var freezeCmd = &cobra.Command{
	Use:     "freeze",
	Short:   "Freeze execution of unit processes",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freezeCmd).Standalone()

	rootCmd.AddCommand(freezeCmd)

	carapace.Gen(freezeCmd).PositionalAnyCompletion(
		action.ActionUnits(freezeCmd).FilterArgs(),
	)
}
