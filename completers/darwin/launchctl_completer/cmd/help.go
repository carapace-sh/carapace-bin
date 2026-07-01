package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Prints the usage for a given subcommand",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()
	rootCmd.AddCommand(helpCmd)
	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.ActionValues(
			"bootstrap", "bootout", "enable", "disable", "kickstart", "kill", "blame",
			"print", "list", "start", "stop", "load", "unload", "remove", "setenv",
			"getenv", "unsetenv", "limit", "reboot", "version",
		),
	)
}
