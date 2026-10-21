package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var custom_devices_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the currently configured custom devices, both enabled and disabled, reachable or not",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(custom_devices_listCmd).Standalone()

	custom_devices_listCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	custom_devicesCmd.AddCommand(custom_devices_listCmd)
}
