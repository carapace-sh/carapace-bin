package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Record and stream a terminal session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionCmd).Standalone()

	sessionCmd.Flags().BoolP("append", "a", false, "Append to an existing recording file")
	sessionCmd.Flags().String("audio-url", "", "Audio stream URL for synchronized playback")
	sessionCmd.Flags().String("capture-env", "", "Comma-separated list of environment variables to capture")
	sessionCmd.Flags().BoolP("capture-input", "I", false, "Enable input (keyboard) capture")
	sessionCmd.Flags().StringP("command", "c", "", "Command to start in the session")
	sessionCmd.Flags().String("description", "", "Description of the session (Markdown supported)")
	sessionCmd.Flags().Bool("headless", false, "Headless mode - don't use the terminal for I/O")
	sessionCmd.Flags().BoolP("help", "h", false, "Print help")
	sessionCmd.Flags().StringP("idle-time-limit", "i", "", "Limit idle time to a given number of seconds")
	sessionCmd.Flags().String("log-file", "", "Log file path")
	sessionCmd.Flags().StringP("output-file", "o", "", "Save the session to a file")
	sessionCmd.Flags().StringP("output-format", "f", "", "Output file format")
	sessionCmd.Flags().Bool("overwrite", false, "Overwrite the output file if it already exists")
	sessionCmd.Flags().Bool("return", false, "Return the session's exit status")
	sessionCmd.Flags().String("server-url", "", "asciinema server URL")
	sessionCmd.Flags().StringP("stream-local", "l", "", "Stream via the local HTTP server")
	sessionCmd.Flags().StringP("stream-remote", "r", "", "Stream via remote asciinema server")
	sessionCmd.Flags().StringP("title", "t", "", "Title of the session")
	sessionCmd.Flags().String("visibility", "", "Stream visibility level")
	sessionCmd.Flags().String("window-size", "", "Override session's terminal window size")
	rootCmd.AddCommand(sessionCmd)

	sessionCmd.Flag("stream-local").NoOptDefVal = "127.0.0.1:0"
	sessionCmd.Flag("stream-remote").NoOptDefVal = " "

	carapace.Gen(sessionCmd).FlagCompletion(carapace.ActionMap{
		"capture-env":   os.ActionEnvironmentVariables().UniqueList(","),
		"command":       bridge.ActionCarapaceBin().Split(),
		"log-file":      carapace.ActionFiles(),
		"output-file":   carapace.ActionFiles(),
		"output-format": carapace.ActionValues("asciicast-v2", "asciicast-v3", "raw", "txt"),
		"stream-local": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("0.0.0.0", "127.0.0.1")
			default:
				return net.ActionPorts()
			}
		}),
		"visibility": carapace.ActionValues("private", "public", "unlisted"),
	})
}
