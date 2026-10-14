package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fullchargeCmd = &cobra.Command{
	Use:   "fullcharge",
	Short: "Charge battery to full capacity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fullchargeCmd).Standalone()

	rootCmd.AddCommand(fullchargeCmd)

	carapace.Gen(fullchargeCmd).PositionalCompletion(
		action.ActionBatteries(),
	)
}
