package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var dischargeCmd = &cobra.Command{
	Use:   "discharge",
	Short: "Force a complete or partial discharge of the battery while on AC power",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dischargeCmd).Standalone()

	rootCmd.AddCommand(dischargeCmd)

	carapace.Gen(dischargeCmd).PositionalCompletion(
		action.ActionBatteries(),
		carapace.ActionValues(),
	)

	carapace.Gen(dischargeCmd).DashAnyCompletion(action.ActionConfigParams())
}
