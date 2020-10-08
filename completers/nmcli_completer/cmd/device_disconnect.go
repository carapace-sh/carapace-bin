package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var device_disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_disconnectCmd).Standalone()

	deviceCmd.AddCommand(device_disconnectCmd)

	carapace.Gen(device_disconnectCmd).PositionalAnyCompletion(action.ActionDevices(""))
}
