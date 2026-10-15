package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var recalibrateCmd = &cobra.Command{
	Use:   "recalibrate",
	Short: "Perform a battery recalibration while on AC power",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recalibrateCmd).Standalone()

	rootCmd.AddCommand(recalibrateCmd)

	carapace.Gen(recalibrateCmd).PositionalCompletion(
		action.ActionBatteries(),
	)

	carapace.Gen(recalibrateCmd).DashAnyCompletion(action.ActionConfigParams())
}
