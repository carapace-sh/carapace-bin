package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var screenshotCmd = &cobra.Command{
	Use:     "screenshot",
	Aliases: []string{"pic"},
	Short:   "Take a screenshot from a connected device",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(screenshotCmd).Standalone()

	screenshotCmd.Flags().String("device-connection", "", "Discover devices based on connection type.")
	screenshotCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach. Longer timeouts may be necessary for networked devices.")
	screenshotCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	screenshotCmd.Flags().StringP("out", "o", "", "Location to write the screenshot.")
	screenshotCmd.Flags().String("type", "", "The type of screenshot to retrieve.")
	screenshotCmd.Flags().String("vm-service-url", "", "The VM Service URL to which to connect. This is required when \"--type\" is \"skia\". To find the VM service URL, use \"flutter run\" and look for \"A Dart VM Service ... is available at\" in the output.")
	rootCmd.AddCommand(screenshotCmd)

	carapace.Gen(screenshotCmd).FlagCompletion(carapace.ActionMap{
		"device-connection": carapace.ActionValuesDescribed(
			"attached", "Only searches for devices connected by USB or built-in (such as simulators/emulators, MacOS/Windows, Chrome)",
			"wireless", "Only searches for devices connected wirelessly. Discovering wireless devices may take longer.",
			"both", "Searches for both attached and wireless devices.",
		),
		"out": carapace.ActionFiles(),
		"type": carapace.ActionValuesDescribed(
			"device", "Delegate to the device's native screenshot capabilities.",
			"rasterizer", "Render the Flutter app using the rasterizer.",
			"skia", "Render the Flutter app as a Skia picture.",
		),
	})
}
