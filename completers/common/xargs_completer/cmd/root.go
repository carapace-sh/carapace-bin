package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xargs",
	Short: "build and execute command lines from standard input",
	Long:  "https://linux.die.net/man/1/xargs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("E", "E", "", "set logical EOF string")
	rootCmd.Flags().StringS("I", "I", "", "same as --replace=R")
	rootCmd.Flags().StringP("arg-file", "a", "", "read arguments from FILE")
	rootCmd.Flags().StringP("delimiter", "d", "", "items in input stream are separated by CHARACTER")
	rootCmd.Flags().StringP("eof", "e", "", "equivalent to -E END if END is specified")
	rootCmd.Flags().BoolP("exit", "x", false, "exit if the size (see -s) is exceeded")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("interactive", "p", false, "prompt before running commands")
	rootCmd.Flags().StringS("l", "l", "", "similar to -L but defaults to at most one non- blank input line")
	rootCmd.Flags().StringP("max-args", "n", "", "use at most MAX-ARGS arguments per command line")
	rootCmd.Flags().StringP("max-chars", "s", "", "limit length of command line to MAX-CHARS")
	rootCmd.Flags().StringP("max-lines", "L", "", "use at most MAX-LINES non-blank input lines per command line")
	rootCmd.Flags().StringP("max-procs", "P", "", "run at most MAX-PROCS processes at a time")
	rootCmd.Flags().BoolP("no-run-if-empty", "r", false, "if there are no arguments, then do not run COMMAND")
	rootCmd.Flags().BoolP("null", "0", false, "items are separated by a null")
	rootCmd.Flags().BoolP("open-tty", "o", false, "Reopen stdin as /dev/tty in the child process")
	rootCmd.Flags().String("process-slot-var", "", "set environment variable VAR in child processes")
	rootCmd.Flags().StringP("replace", "i", "", "replace R in INITIAL-ARGS with names")
	rootCmd.Flags().Bool("show-limits", false, "show limits on command-line length")
	rootCmd.Flags().BoolP("verbose", "t", false, "print commands before executing them")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	rootCmd.Flag("eof").NoOptDefVal = " "
	rootCmd.Flag("replace").NoOptDefVal = "{}"

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"arg-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
