package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "system_profiler",
	Short: "reports system hardware and software configuration",
	Long:  "https://keith.github.io/xcode-manpages/system_profiler.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("detailLevel", "detailLevel", false, "Set the detail level")
	rootCmd.Flags().StringP("file", "file", "", "Output to the specified file")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("list", "listDataTypes", false, "List available data types")
	rootCmd.Flags().StringP("timeout", "timeout", "", "Set the timeout in seconds")
	rootCmd.Flags().StringP("xml", "xml", "", "Output as XML plist")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"SPHardwareDataType", "Hardware information",
			"SPSoftwareDataType", "Software information",
			"SPNetworkDataType", "Network configuration",
			"SPDisplaysDataType", "Display information",
			"SPStorageDataType", "Storage information",
			"SPPCIDataType", "PCI devices",
			"SPUSBDataType", "USB devices",
			"SPThunderboltDataType", "Thunderbolt devices",
			"SPBluetoothDataType", "Bluetooth information",
			"SPPrintersDataType", "Printer information",
			"SPAudioDataType", "Audio devices",
			"SPPowerDataType", "Power information",
			"SPBatteryDataType", "Battery information",
			"SPSensorsDataType", "Sensor information",
			"SPNVMeDataType", "NVMe storage information",
			"SPPlatformDataType", "Platform information",
		),
	)
}
