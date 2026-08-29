package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start-stop-daemon",
	Short: "start and stop system daemon programs",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("background", "b", false, "force daemon to background")
	rootCmd.Flags().String("capabilities", "", "set the inheritable, ambient and bounding capabilities")
	rootCmd.Flags().StringP("chdir", "d", "", "change the PWD")
	rootCmd.Flags().StringP("chroot", "r", "", "chroot to this directory")
	rootCmd.Flags().StringP("chuid", "c", "", "deprecated, use --user")
	rootCmd.Flags().StringP("env", "e", "", "set an environment string")
	rootCmd.Flags().StringP("exec", "x", "", "binary to start/stop")
	rootCmd.Flags().StringP("group", "g", "", "change the process group")
	rootCmd.Flags().BoolP("help", "h", false, "display this help output")
	rootCmd.Flags().BoolP("interpreted", "i", false, "match process name by interpreter")
	rootCmd.Flags().StringP("ionice", "I", "", "set an ionice class:data when starting")
	rootCmd.Flags().BoolP("make-pidfile", "m", false, "create a pidfile")
	rootCmd.Flags().StringP("name", "n", "", "match process name")
	rootCmd.Flags().StringP("nicelevel", "N", "", "set a nicelevel when starting")
	rootCmd.Flags().Bool("no-new-privs", false, "set the No New Privs flag for the program")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().String("notify", "", "configures experimental notification behaviour")
	rootCmd.Flags().BoolP("oknodo", "o", false, "deprecated")
	rootCmd.Flags().String("oom-score-adj", "", "set OOM score adjustment when starting")
	rootCmd.Flags().StringP("pidfile", "p", "", "match pid found in this file")
	rootCmd.Flags().BoolP("progress", "P", false, "print dots each second while waiting")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().StringP("retry", "R", "", "retry schedule to use when stopping")
	rootCmd.Flags().String("scheduler", "", "set process scheduler")
	rootCmd.Flags().String("scheduler-priority", "", "set process scheduler priority")
	rootCmd.Flags().String("secbits", "", "set the security-bits for the program")
	rootCmd.Flags().StringP("signal", "s", "", "send a different signal")
	rootCmd.Flags().BoolP("start", "S", false, "start daemon")
	rootCmd.Flags().StringP("startas", "a", "", "deprecated, use --exec or --name")
	rootCmd.Flags().StringP("stderr", "2", "", "redirect stderr to file")
	rootCmd.Flags().String("stderr-logger", "", "redirect stderr to process")
	rootCmd.Flags().String("stdin", "", "redirect stdin to file")
	rootCmd.Flags().StringP("stdout", "1", "", "redirect stdout to file")
	rootCmd.Flags().String("stdout-logger", "", "redirect stdout to process")
	rootCmd.Flags().BoolP("stop", "K", false, "stop daemon")
	rootCmd.Flags().BoolP("stop-group", "G", false, "stop the whole process group")
	rootCmd.Flags().BoolP("test", "t", false, "test actions, don't do them")
	rootCmd.Flags().StringP("umask", "k", "", "set the umask for the daemon")
	rootCmd.Flags().StringP("user", "u", "", "change the process user")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")
	rootCmd.Flags().StringP("wait", "w", "", "milliseconds to wait for daemon start")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"capabilities":  carapace.ActionFiles(),
		"chdir":         carapace.ActionDirectories(),
		"chroot":        carapace.ActionDirectories(),
		"exec":          carapace.ActionFiles(),
		"ionice":        carapace.ActionValues("0", "1", "2", "3"),
		"nicelevel":     carapace.ActionValues("-20", "0", "19"),
		"oom-score-adj": carapace.ActionValues("-1000", "0", "1000"),
		"pidfile":       carapace.ActionFiles(),
		"retry":         carapace.ActionValues("1", "5", "10", "15", "30", "60"),
		"scheduler":     carapace.ActionValues("other", "batch", "idle", "fifo", "rr"),
		"signal": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return ps.ActionKillSignals("linux").
				Retain("ABRT", "ALRM", "BUS", "CHLD", "CONT", "FPE", "HUP", "ILL", "INT", "IO", "KILL", "PIPE", "PROF", "PWR", "QUIT", "SEGV", "STOP", "SYS", "TERM", "TRAP", "TSTP", "TTIN", "TTOU", "URG", "USR1", "USR2", "VTALRM", "WINCH", "XCPU", "XFSZ")
		}),
		"startas":       carapace.ActionFiles(),
		"stderr":        carapace.ActionFiles(),
		"stderr-logger": carapace.ActionFiles(),
		"stdin":         carapace.ActionFiles(),
		"stdout":        carapace.ActionFiles(),
		"stdout-logger": carapace.ActionFiles(),
		"wait":          carapace.ActionValues("1000", "5000", "10000"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
