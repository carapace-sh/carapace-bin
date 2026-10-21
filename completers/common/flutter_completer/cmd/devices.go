package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List all connected devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicesCmd).Standalone()

	devicesCmd.Flags().String("device-connection", "", "Discover devices based on connection type.")
	devicesCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach. Longer timeouts may be necessary for networked devices.")
	devicesCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	devicesCmd.Flags().Bool("machine", false, "Outputs in a machine readable structured JSON format.")
	devicesCmd.Flags().StringP("timeout", "t", "", "(deprecated) This option has been replaced by \"--device-timeout\".")
	rootCmd.AddCommand(devicesCmd)
	carapace.Gen(devicesCmd).FlagCompletion(carapace.ActionMap{
		"device-connection": carapace.ActionValuesDescribed(
			"attached", "Only searches for devices connected by USB or built-in (such as simulators/emulators, MacOS/Windows, Chrome)",
			"wireless", "Only searches for devices connected wirelessly. Discovering wireless devices may take longer.",
			"both", "Searches for both attached and wireless devices.",
		),
	})

}
