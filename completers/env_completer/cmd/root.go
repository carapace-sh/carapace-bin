package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "env",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("block-signal", "", "block delivery of SIG signal(s) to COMMAND")
	rootCmd.Flags().StringP("chdir", "C", "", "change working directory to DIR")
	rootCmd.Flags().BoolP("debug", "v", false, "print verbose information for each processing step")
	rootCmd.Flags().String("default-signal", "", "reset handling of SIG signal(s) to the default")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-environment", "i", false, "start with an empty environment")
	rootCmd.Flags().String("ignore-signal", "", "set handling of SIG signals(s) to do nothing")
	rootCmd.Flags().Bool("list-signal-handling", false, "list non default signal handling to stderr")
	rootCmd.Flags().BoolP("null", "0", false, "end each output line with NUL, not newline")
	rootCmd.Flags().StringP("split-string", "S", "", "process and split S into separate arguments;")
	rootCmd.Flags().StringP("unset", "u", "", "remove variable from the environment")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"unset":          ActionEnvironmentVariables(),
		"chdir":          carapace.ActionDirectories(),
		"block-signal":   ActionSignals(),
		"default-signal": ActionSignals(),
		"ignore-signal":  ActionSignals(),
	})
}

func ActionEnvironmentVariables() carapace.Action {
	env := os.Environ()
	vars := make([]string, len(env))
	for index, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		vars[index] = pair[0]
	}
	return carapace.ActionValues(vars...)
}

func ActionSignals() carapace.Action {
	return carapace.ActionValues("ABRT", "CHLD", "HUP", "KILL", "PROF", "SEGV", "SYS", "TSTP", "URG", "VTALRM", "XFSZ", "ALRM", "CONT", "ILL", "PIPE", "PWR", "STKFLT", "TERM", "TTIN", "USR1", "WINCH", "BUS", "FPE", "INT", "POLL", "QUIT", "STOP", "TRAP", "TTOU", "USR2", "XCPU")
}
