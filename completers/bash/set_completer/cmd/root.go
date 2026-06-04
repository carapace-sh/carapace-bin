package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "set",
	Short: "Set or unset values of shell options and positional parameters",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("B", "B", false, "the shell will perform brace expansion")
	rootCmd.Flags().BoolS("C", "C", false, "disallow existing regular files to be overwritten by redirection of output")
	rootCmd.Flags().BoolS("E", "E", false, "the ERR trap is inherited by shell functions")
	rootCmd.Flags().BoolS("H", "H", false, "enable ! style history substitution")
	rootCmd.Flags().BoolS("P", "P", false, "do not resolve symbolic links when executing commands")
	rootCmd.Flags().BoolS("T", "T", false, "the DEBUG and RETURN traps are inherited by shell functions")
	rootCmd.Flags().BoolS("a", "a", false, "mark variables which are modified or created for export")
	rootCmd.Flags().BoolS("b", "b", false, "cause the status of terminated background jobs to be reported immediately")
	rootCmd.Flags().BoolS("e", "e", false, "exit immediately if a command exits with a non-zero status")
	rootCmd.Flags().BoolS("f", "f", false, "disable file name generation (globbing)")
	rootCmd.Flags().BoolS("h", "h", false, "remember the location of commands as they are looked up")
	rootCmd.Flags().BoolS("k", "k", false, "all arguments in the form of assignment statements are placed in the environment")
	rootCmd.Flags().BoolS("m", "m", false, "enable job control")
	rootCmd.Flags().BoolS("n", "n", false, "read commands but do not execute them")
	rootCmd.Flags().StringS("o", "o", "", "set the option corresponding to option-name")
	rootCmd.Flags().BoolS("p", "p", false, "turned on whenever the real and effective user ids do not match")
	rootCmd.Flags().BoolS("t", "t", false, "exit after reading and executing one command")
	rootCmd.Flags().BoolS("u", "u", false, "treat unset variables as an error when substituting")
	rootCmd.Flags().BoolS("v", "v", false, "print shell input lines as they are read")
	rootCmd.Flags().BoolS("x", "x", false, "print commands and their arguments as they are executed")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionValuesDescribed(
			"allexport", "same as -a",
			"braceexpand", "same as -B",
			"emacs", "use an emacs-style line editing interface",
			"errexit", "same as -e",
			"errtrace", "same as -E",
			"functrace", "same as -T",
			"hashall", "same as -h",
			"histexpand", "same as -H",
			"history", "enable command history",
			"ignoreeof", "shell does not exit on EOF",
			"keyword", "same as -k",
			"monitor", "same as -m",
			"noclobber", "same as -C",
			"noexec", "same as -n",
			"noglob", "same as -f",
			"nolog", "do not save function definitions in history",
			"notify", "same as -b",
			"nounset", "same as -u",
			"onecmd", "same as -t",
			"physical", "same as -P",
			"pipefail", "return value of pipeline is the rightmost command to exit with non-zero status",
			"posix", "change the behavior of bash to match the POSIX standard",
			"privileged", "same as -p",
			"verbose", "same as -v",
			"vi", "use a vi-style line editing interface",
			"xtrace", "same as -x",
		),
	})
}
