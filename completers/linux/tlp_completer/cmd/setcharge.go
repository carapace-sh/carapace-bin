package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setchargeCmd = &cobra.Command{
	Use:   "setcharge",
	Short: "Change battery charge thresholds temporarily",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setchargeCmd).Standalone()

	rootCmd.AddCommand(setchargeCmd)

	carapace.Gen(setchargeCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
		action.ActionBatteries(),
	)

	carapace.Gen(setchargeCmd).DashAnyCompletion(action.ActionConfigParams())
}
