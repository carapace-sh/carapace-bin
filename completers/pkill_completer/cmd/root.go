package cmd

import (
	ps "github.com/mitchellh/go-ps"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkill",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("count", "c", false, "count of matching processes")
	rootCmd.Flags().BoolP("echo", "e", false, "display what is killed")
	rootCmd.Flags().StringP("euid", "u", "", "match by effective IDs")
	rootCmd.Flags().BoolP("exact", "x", false, "match exactly with the command name")
	rootCmd.Flags().BoolP("full", "f", false, "use full process name to match")
	rootCmd.Flags().StringP("group", "G", "", "match real group IDs")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "match case insensitively")
	rootCmd.Flags().BoolP("logpidfile", "L", false, "fail if PID file is not locked")
	rootCmd.Flags().BoolP("newest", "n", false, "select most recently started")
	rootCmd.Flags().String("ns", "", "match the processes that belong to the same namespace as <pid>")
	rootCmd.Flags().String("nslist", "", "list which namespaces will be considered for the --ns option.")
	rootCmd.Flags().BoolP("oldest", "o", false, "select least recently started")
	rootCmd.Flags().StringP("parent", "P", "", "match only child processes of the given parent")
	rootCmd.Flags().StringP("pgroup", "g", "", "match listed process group IDs")
	rootCmd.Flags().StringP("pidfile", "F", "", "read PIDs from file")
	rootCmd.Flags().StringP("runstates", "r", "", "match runstates [D,S,Z,...]")
	rootCmd.Flags().StringP("session", "s", "", "match session IDs")
	//rootCmd.Flags().StringP("signal", "<sig>", "", "signal to send (either number or name)")
	rootCmd.Flags().String("signal", "", "signal to send (either number or name)")
	rootCmd.Flags().StringP("terminal", "t", "", "match by controlling terminal")
	rootCmd.Flags().StringP("uid", "U", "", "match by real IDs")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"euid":      carapace.ActionUsers(),
		"group":     carapace.ActionGroups(),
		"nslist":    carapace.ActionValues("ipc", "mnt", "net", "pid", "user", "uts"),
		"pidfile":   carapace.ActionFiles(""),
		"runstates": ActionProcessStates(),
		"signal":    carapace.ActionKillSignals(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ActionExecutables(),
	)
}

func ActionExecutables() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if processes, err := ps.Processes(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			executables := make([]string, 0)
			for _, process := range processes {
				executables = append(executables, process.Executable())
			}
			return carapace.ActionValues(executables...)
		}
	})
}

func ActionProcessStates() carapace.Action {
	return carapace.ActionValuesDescribed(
		"D", "uninterruptible sleep (usually IO)",
		"I", "Idle kernel thread",
		"R", "running or runnable (on run queue)",
		"S", "interruptible sleep (waiting for an event to complete)",
		"T", "stopped by job control signal",
		"W", "paging (not valid since the 2.6.xx kernel)",
		"X", "dead (should never be seen)",
		"Z", "defunct (zombie) process, terminated but not reaped by its parent",
		"t", "stopped by debugger during the tracing",
	)
}
