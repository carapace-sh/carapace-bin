package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logger",
	Short: "make entries in the system log",
	Long:  "https://man.freebsd.org/cgi/man.cgi?logger",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("f", "f", "", "Read the contents of the specified file into syslog")
	rootCmd.Flags().BoolS("i", "i", false, "Log the process id")
	rootCmd.Flags().StringS("p", "p", "", "Enter the message with the specified priority")
	rootCmd.Flags().BoolS("s", "s", false, "Log to standard error as well")
	rootCmd.Flags().StringS("t", "t", "", "Mark every line with the specified tag")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
		"p": carapace.ActionValuesDescribed(
			"emerg", "Emergency",
			"alert", "Alert",
			"crit", "Critical",
			"err", "Error",
			"warning", "Warning",
			"notice", "Notice",
			"info", "Info",
			"debug", "Debug",
			"user.notice", "Default",
			"local0.info", "Local0",
			"local1.info", "Local1",
			"local2.info", "Local2",
			"local3.info", "Local3",
			"local4.info", "Local4",
			"local5.info", "Local5",
			"local6.info", "Local6",
			"local7.info", "Local7",
		),
	})
}
