package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show log output for running Flutter apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().Bool("adb-log-filtering", false, "Filter adb logs so that logs not emitted by the current flutter application are not displayed.")
	logsCmd.Flags().BoolP("clear", "c", false, "Clear log history before reading from logs.")
	logsCmd.Flags().String("device-connection", "", "Discover devices based on connection type.")
	logsCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach. Longer timeouts may be necessary for networked devices.")
	logsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	logsCmd.Flags().Bool("no-adb-log-filtering", false, "Filter adb logs so that logs not emitted by the current flutter application are not displayed.")
	rootCmd.AddCommand(logsCmd)
	carapace.Gen(logsCmd).FlagCompletion(carapace.ActionMap{
		"device-connection": carapace.ActionValuesDescribed(
			"attached", "Only searches for devices connected by USB or built-in (such as simulators/emulators, MacOS/Windows, Chrome)",
			"wireless", "Only searches for devices connected wirelessly. Discovering wireless devices may take longer.",
			"both", "Searches for both attached and wireless devices.",
		),
	})

}
