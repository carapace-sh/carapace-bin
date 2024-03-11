package cmd

import (
	"runtime"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serialCmd = &cobra.Command{
	Use:   "serial [OPTIONS] <PORT>",
	Short: "Open a serial port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serialCmd).Standalone()

	serialCmd.Flags().String("baud", "", "Set the baud rate")
	serialCmd.Flags().String("class", "", "Override the default windowing system class")
	serialCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	serialCmd.Flags().String("position", "", "Override the position for the initial window launched by this process")
	rootCmd.AddCommand(serialCmd)

	carapace.Gen(serialCmd).FlagCompletion(carapace.ActionMap{
		"baud": net.ActionBaudRates(),
	})

	carapace.Gen(serialCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch runtime.GOOS {
			case "windows":
				return carapace.ActionValues() // TODO COM0...
			default:
				return carapace.ActionFiles() // TODO /dev/tty...
			}
		}),
	)
}
