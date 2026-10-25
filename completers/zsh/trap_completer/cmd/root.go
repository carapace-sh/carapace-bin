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

	// TODO fix positional completion
	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(
			"0", "ALRM", "BUS", "CHLD", "CONT", "DEBUG", "EXIT", "FPE", "HUP",
			"ILL", "INT", "IOT", "KILL", "PIPE", "POLL", "PROF", "PWR",
			"QUIT", "SEGV", "STKFLT", "STOP", "SYS", "TERM", "TRAP", "TSTP",
			"TTIN", "TTOU", "URG", "USR1", "USR2", "VTALRM", "WINCH", "XCPU",
			"XFSZ", "ZERR",
		).Usage("signal"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues().Usage("command"),
	)
}
