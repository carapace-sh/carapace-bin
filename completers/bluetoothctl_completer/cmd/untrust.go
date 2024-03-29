package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/bluetoothctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var untrustCmd = &cobra.Command{
	Use:   "untrust <device>",
	Short: "Untrust device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untrustCmd).Standalone()
	rootCmd.AddCommand(untrustCmd)
	carapace.Gen(untrustCmd).PositionalCompletion(
		action.ActionDevices("Trusted"),
	)
}
