package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Display serial output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitorCmd).Standalone()

	monitorCmd.Flags().StringP("command-file", "c", "", "Run IDF Monitor in command-stream mode from FILE")
	monitorCmd.Flags().Bool("disable-auto-color", false, "Disable auto coloring logs")
	monitorCmd.Flags().BoolP("encrypted", "E", false, "Enable encrypted flash targets")
	monitorCmd.Flags().Bool("force-color", false, "Always print ANSI for colors")
	monitorCmd.Flags().String("monitor-baud", "", "Baud rate for monitor")
	monitorCmd.Flags().Bool("no-reset", false, "Disable reset on monitor startup")
	monitorCmd.Flags().String("print-filter", "", "Filter monitor output")
	monitorCmd.Flags().String("timestamp-format", "", "Set formatting of timestamps compatible with strftime()")
	monitorCmd.Flags().Bool("timestamps", false, "Print a time stamp in the beginning of each line")

	rootCmd.AddCommand(monitorCmd)

	carapace.Gen(monitorCmd).FlagCompletion(carapace.ActionMap{
		"command-file": carapace.ActionFiles(),
		"monitor-baud": espidf.ActionBaudRates(),
	})
}
