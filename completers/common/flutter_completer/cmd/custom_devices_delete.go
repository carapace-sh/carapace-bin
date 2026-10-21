package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var custom_devices_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a device from the config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(custom_devices_deleteCmd).Standalone()

	custom_devices_deleteCmd.Flags().StringP("device-id", "d", "", "Target device id or name (prefixes allowed).")
	custom_devices_deleteCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	custom_devicesCmd.AddCommand(custom_devices_deleteCmd)

	carapace.Gen(custom_devices_deleteCmd).FlagCompletion(carapace.ActionMap{
		"device-id": action.ActionCustomDevices(),
	})
}
