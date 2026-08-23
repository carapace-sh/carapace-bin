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

	rootCmd.Flags().Bool("json", false, "Output as JSON")
	rootCmd.Flags().Bool("listDataTypes", false, "List available data types")
	rootCmd.Flags().Bool("xml", false, "Output as XML plist")

	rootCmd.Flags().String("detailLevel", "", "Set the detail level: mini, basic, or full")
	rootCmd.Flags().String("timeout", "", "Set the timeout in seconds")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"detailLevel": carapace.ActionValues("mini", "basic", "full"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"SPAudioDataType", "Audio devices",
			"SPBatteryDataType", "Battery information",
			"SPBluetoothDataType", "Bluetooth information",
			"SPCameraDataType", "Camera information",
			"SPCardReaderDataType", "Card reader information",
			"SPDeveloperToolsDataType", "Developer tools information",
			"SPDiagnosticsDataType", "Diagnostics information",
			"SPDisabledApplicationsDataType", "Disabled applications",
			"SPDisplaysDataType", "Display information",
			"SPExtensionsDataType", "Extensions information",
			"SPFibreChannelDataType", "Fibre Channel information",
			"SPFireWireDataType", "FireWire devices",
			"SPFirewallDataType", "Firewall information",
			"SPHardwareDataType", "Hardware information",
			"SPHardwareRAIDDataType", "Hardware RAID information",
			"SPInstallHistoryDataType", "Install history",
			"SPLanguagesDataType", "Languages information",
			"SPLegacySoftwareDataType", "Legacy software information",
			"SPNetworkDataType", "Network configuration",
			"SPNVMeDataType", "NVMe storage information",
			"SPPCIDataType", "PCI devices",
			"SPPlatformDataType", "Platform information",
			"SPPowerDataType", "Power information",
			"SPPrintersDataType", "Printer information",
			"SPPrintersSoftwareDataType", "Printer software information",
			"SPSASDataType", "SAS information",
			"SPSerialATADataType", "Serial ATA information",
			"SPServicesDataType", "Services information",
			"SPSoftwareDataType", "Software information",
			"SPSoftwareRAIDDataType", "Software RAID information",
			"SPStorageDataType", "Storage information",
			"SPThunderboltDataType", "Thunderbolt devices",
			"SPUSBDataType", "USB devices",
		),
	)
}
