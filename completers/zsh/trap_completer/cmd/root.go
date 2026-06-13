package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trap",
	Short: "Trap signals and execute commands",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-trap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(
			"HUP", "INT", "QUIT", "ILL", "TRAP", "IOT", "BUS", "FPE", "KILL",
			"USR1", "SEGV", "USR2", "PIPE", "ALRM", "TERM", "STKFLT", "CHLD",
			"CONT", "STOP", "TSTP", "TTIN", "TTOU", "URG", "XCPU", "XFSZ",
			"VTALRM", "PROF", "WINCH", "POLL", "PWR", "SYS", "ZERR", "DEBUG",
			"EXIT", "0",
		).Usage("signal"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues().Usage("command"),
	)
}
