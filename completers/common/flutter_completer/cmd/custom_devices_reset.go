package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var custom_devices_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the config file to the default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(custom_devices_resetCmd).Standalone()

	custom_devices_resetCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	custom_devicesCmd.AddCommand(custom_devices_resetCmd)
}
