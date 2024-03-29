package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/bluetoothctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pairCmd = &cobra.Command{
	Use:   "pair <device>",
	Short: "Pair with device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pairCmd).Standalone()
	rootCmd.AddCommand(pairCmd)
	carapace.Gen(pairCmd).PositionalCompletion(
		action.ActionDevices(""),
	)
}
