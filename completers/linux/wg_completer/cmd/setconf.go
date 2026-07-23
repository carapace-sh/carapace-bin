package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/wg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setconfCmd = &cobra.Command{
	Use:   "setconf",
	Short: "Applies a configuration file to a WireGuard interface",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setconfCmd).Standalone()

	rootCmd.AddCommand(setconfCmd)

	carapace.Gen(setconfCmd).PositionalCompletion(
		action.ActionInterfaces(),
		carapace.ActionFiles(),
	)
}
