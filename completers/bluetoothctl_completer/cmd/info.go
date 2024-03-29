package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/bluetoothctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [device]",
	Short: "Device/Set information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	rootCmd.AddCommand(infoCmd)
	carapace.Gen(infoCmd).PositionalCompletion(
		action.ActionDevices(""),
	)
}
