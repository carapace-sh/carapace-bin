package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "syslog",
	Short: "Apple System Log utility",
	Long:  "https://keith.github.io/xcode-manpages/syslog.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("B", false, "Binary output")
	rootCmd.Flags().Bool("C", false, "Fetch and print console messages")
	rootCmd.Flags().Bool("config", false, "Configuration options")
	rootCmd.Flags().StringP("dir", "d", "", "Read log messages from a directory")
	rootCmd.Flags().StringP("encoding", "E", "", "Encoding format")
	rootCmd.Flags().BoolP("extract", "x", false, "Copy log messages to a file")
	rootCmd.Flags().StringP("file", "f", "", "Read log messages from a file")
	rootCmd.Flags().StringP("format", "F", "", "Output format string")
	rootCmd.Flags().Bool("help", false, "Print usage message")
	rootCmd.Flags().StringP("host", "r", "", "Send to remote host")
	rootCmd.Flags().BoolP("keyval", "k", false, "Use key-value pairs for structured message")
	rootCmd.Flags().StringP("level", "l", "", "Set log level (priority) of the message")
	rootCmd.Flags().String("module", "", "Module name")
	rootCmd.Flags().StringP("process", "c", "", "Get or set log mask for a process")
	rootCmd.Flags().BoolP("send", "s", false, "Send log messages")
	rootCmd.Flags().StringP("timeformat", "T", "", "Time format string")
	rootCmd.Flags().StringP("wait", "w", "", "Wait for new log messages")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir":  carapace.ActionDirectories(),
		"file": carapace.ActionFiles(),
		"level": carapace.ActionValuesDescribed(
			"0", "Emergency",
			"1", "Alert",
			"2", "Critical",
			"3", "Error",
			"4", "Warning",
			"5", "Notice",
			"6", "Info",
			"7", "Debug",
		),
	})
}
