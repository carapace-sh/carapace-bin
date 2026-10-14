package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var chargeonceCmd = &cobra.Command{
	Use:   "chargeonce",
	Short: "Charge battery to the stop charge threshold once",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chargeonceCmd).Standalone()

	rootCmd.AddCommand(chargeonceCmd)

	carapace.Gen(chargeonceCmd).PositionalCompletion(
		action.ActionBatteries(),
	)
}
