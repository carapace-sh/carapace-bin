package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream a terminal session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(streamCmd).Standalone()

	streamCmd.Flags().String("audio-url", "", "Audio stream URL for synchronized playback")
	streamCmd.Flags().String("capture-env", "", "Comma-separated list of environment variables to capture")
	streamCmd.Flags().BoolP("capture-input", "I", false, "Enable input (keyboard) capture")
	streamCmd.Flags().StringP("command", "c", "", "Command to start in the session")
	streamCmd.Flags().String("description", "", "Description of the session (Markdown supported)")
	streamCmd.Flags().Bool("headless", false, "Headless mode - don't use the terminal for I/O")
	streamCmd.Flags().BoolP("help", "h", false, "Print help")
	streamCmd.Flags().StringP("local", "l", "", "Stream via the local HTTP server")
	streamCmd.Flags().String("log-file", "", "Log file path")
	streamCmd.Flags().StringP("remote", "r", "", "Stream via remote asciinema server")
	streamCmd.Flags().Bool("return", false, "Return the session's exit status")
	streamCmd.Flags().String("server-url", "", "asciinema server URL")
	streamCmd.Flags().StringP("title", "t", "", "Title of the session")
	streamCmd.Flags().String("visibility", "", "Visibility level")
	streamCmd.Flags().String("window-size", "", "Override session's terminal window size")
	rootCmd.AddCommand(streamCmd)

	streamCmd.Flag("local").NoOptDefVal = "127.0.0.1:0"
	streamCmd.Flag("remote").NoOptDefVal = " "

	carapace.Gen(streamCmd).FlagCompletion(carapace.ActionMap{
		"capture-env": os.ActionEnvironmentVariables().UniqueList(","),
		"command":     bridge.ActionCarapaceBin().Split(),
		"local": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("0.0.0.0", "127.0.0.1")
			default:
				return net.ActionPorts()
			}
		}),
		"log-file":   carapace.ActionFiles(),
		"visibility": carapace.ActionValues("private", "public", "unlisted"),
	})
}
