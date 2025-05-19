package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream a terminal session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(streamCmd).Standalone()

	streamCmd.Flags().StringP("command", "c", "", "Command to stream")
	streamCmd.Flags().Bool("headless", false, "Use headless mode - don't use TTY for input/output")
	streamCmd.Flags().BoolP("input", "I", false, "Enable input capture")
	streamCmd.Flags().String("log-file", "", "Log file path")
	streamCmd.Flags().StringP("relay", "r", "", "Relay the stream via an asciinema server")
	streamCmd.Flags().StringP("serve", "s", "", "Serve the stream with the built-in HTTP server")
	streamCmd.Flags().String("tty-size", "", "Override terminal size for the session")
	rootCmd.AddCommand(streamCmd)

	carapace.Gen(streamCmd).FlagCompletion(carapace.ActionMap{
		"command": carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionExecutables(),
		).ToA(),
		"log-file": carapace.ActionFiles(),
		"serve": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return carapace.ActionValues()
			}
			return net.ActionPorts()
		}),
	})
}
