package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	os_actions "github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "idf_monitor.py",
	Short: "Serial output monitor for ESP-IDF",
	Long:  "https://github.com/espressif/esp-idf-monitor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("baud", "b", "", "Serial port baud rate")
	rootCmd.Flags().String("decode-coredumps", "", "Handling of core dumps found in serial output")
	rootCmd.Flags().String("decode-panic", "", "Handling of panic handler info found in serial output")
	rootCmd.Flags().BoolP("disable-address-decoding", "d", false, "Don't print decoded addresses from the ELF file")
	rootCmd.Flags().Bool("disable-auto-color", false, "Disable automatic color addition to monitor output")
	rootCmd.Flags().Bool("encrypted", false, "Use encrypted targets while running make")
	rootCmd.Flags().String("eol", "", "End of line to use when sending to the serial port")
	rootCmd.Flags().Bool("force-color", false, "Always colored monitor output")
	rootCmd.Flags().StringP("make", "m", "", "Command to run make")
	rootCmd.Flags().Bool("no-reset", false, "Do not reset the chip on monitor startup")
	rootCmd.Flags().String("open-port-attempts", "", "Number of attempts to wait for the port to appear")
	rootCmd.Flags().StringP("port", "p", "", "Serial port device")
	rootCmd.Flags().String("print_filter", "", "Filtering string")
	rootCmd.Flags().String("revision", "", "Revision of the target")
	rootCmd.Flags().String("rom-elf-file", "", "ELF file of target ROM for address decoding")
	rootCmd.Flags().BoolP("save-log", "s", false, "Save log of monitor")
	rootCmd.Flags().String("target", "", "Target name (used when stack dump decoding is enabled)")
	rootCmd.Flags().String("timestamp-format", "", "Set a strftime()-compatible timestamp format")
	rootCmd.Flags().Bool("timestamps", false, "Add timestamp for each line")
	rootCmd.Flags().String("toolchain-prefix", "", "Triplet prefix to add before cross-toolchain names")
	rootCmd.Flags().String("ws", "", "WebSocket URL for communicating with IDE tools")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"baud":             net.ActionBaudRates(),
		"decode-coredumps": carapace.ActionValues("disable", "info"),
		"decode-panic":     carapace.ActionValues("backtrace", "disable"),
		"eol":              carapace.ActionValues("CR", "CRLF", "LF"),
		"port":             os_actions.ActionSerialPorts(),
		"rom-elf-file":     carapace.ActionFiles(".elf"),
		"target":           espidf.ActionTargets(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".elf"),
	)
}
