package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/bluetoothctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var trustCmd = &cobra.Command{
	Use:   "trust [device]",
	Short: "Trust device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustCmd).Standalone()
	rootCmd.AddCommand(trustCmd)
	carapace.Gen(trustCmd).PositionalCompletion(
		action.ActionDevices(""),
	)
}
