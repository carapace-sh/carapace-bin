package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var recCmd = &cobra.Command{
	Use:   "rec",
	Short: "Record terminal session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recCmd).Standalone()

	recCmd.Flags().BoolP("append", "a", false, "Append to an existing recording file")
	recCmd.Flags().StringP("command", "c", "", "Command to record")
	recCmd.Flags().String("env", "", "List of env vars to save")
	recCmd.Flags().String("filename", "", "Filename template, used when recording to a directory")
	recCmd.Flags().StringP("format", "f", "", "Recording file format")
	recCmd.Flags().Bool("headless", false, "Use headless mode - don't use TTY for input/output")
	recCmd.Flags().BoolP("help", "h", false, "Print help")
	recCmd.Flags().StringP("idle-time-limit", "i", "", "Limit idle time to a given number of seconds")
	recCmd.Flags().BoolP("input", "I", false, "Enable input recording")
	recCmd.Flags().Bool("overwrite", false, "Overwrite target file if it already exists")
	recCmd.Flags().BoolP("quiet", "q", false, "Quiet mode, i.e. suppress diagnostic messages")
	recCmd.Flags().StringP("title", "t", "", "Title of the recording")
	recCmd.Flags().String("tty-size", "", "Override terminal size for the recorded command")

	rootCmd.AddCommand(recCmd)

	carapace.Gen(recCmd).FlagCompletion(carapace.ActionMap{
		"command": carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionExecutables(),
		).ToA(),
		"env":      os.ActionEnvironmentVariables().UniqueList(","),
		"filename": carapace.ActionFiles(),
		"format":   carapace.ActionValues("asciicast", "raw", "txt"),
	})

	carapace.Gen(recCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
