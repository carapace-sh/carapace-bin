package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var recordCmd = &cobra.Command{
	Use:     "record",
	Aliases: []string{"rec"},
	Short:   "Record a terminal session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordCmd).Standalone()

	recordCmd.Flags().BoolP("append", "a", false, "Append to an existing recording file")
	recordCmd.Flags().String("capture-env", "", "Comma-separated list of environment variables to capture")
	recordCmd.Flags().BoolP("capture-input", "I", false, "Enable input (keyboard) capture")
	recordCmd.Flags().StringP("command", "c", "", "Command to start in the session")
	recordCmd.Flags().Bool("headless", false, "Headless mode - don't use the terminal for I/O")
	recordCmd.Flags().BoolP("help", "h", false, "Print help")
	recordCmd.Flags().StringP("idle-time-limit", "i", "", "Limit idle time to a given number of seconds")
	recordCmd.Flags().String("log-file", "", "Log file path")
	recordCmd.Flags().StringP("output-format", "f", "", "Output file format")
	recordCmd.Flags().Bool("overwrite", false, "Overwrite the output file if it already exists")
	recordCmd.Flags().Bool("return", false, "Return the session's exit status")
	recordCmd.Flags().StringP("title", "t", "", "Title of the recording")
	recordCmd.Flags().String("window-size", "", "Override session's terminal window size")
	rootCmd.AddCommand(recordCmd)

	carapace.Gen(recordCmd).FlagCompletion(carapace.ActionMap{
		"capture-env":   os.ActionEnvironmentVariables().UniqueList(","),
		"command":       bridge.ActionCarapaceBin().Split(),
		"log-file":      carapace.ActionFiles(),
		"output-format": carapace.ActionValues("asciicast-v2", "asciicast-v3", "raw", "txt"),
	})

	carapace.Gen(recordCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
