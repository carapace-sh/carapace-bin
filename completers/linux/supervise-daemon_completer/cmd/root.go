package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "supervise-daemon",
	Short: "supervise a daemon process",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("capabilities", "", "set the inheritable, ambient and bounding capabilities")
	rootCmd.Flags().StringP("chdir", "d", "", "change the PWD")
	rootCmd.Flags().StringP("chroot", "r", "", "chroot to this directory")
	rootCmd.Flags().StringP("env", "e", "", "set an environment string")
	rootCmd.Flags().StringP("group", "g", "", "change the process group")
	rootCmd.Flags().StringP("healthcheck-delay", "A", "", "set an initial health check delay")
	rootCmd.Flags().StringP("healthcheck-timer", "a", "", "set a health check timer")
	rootCmd.Flags().BoolP("help", "h", false, "display this help output")
	rootCmd.Flags().StringP("ionice", "I", "", "set an ionice class:data when starting")
	rootCmd.Flags().StringP("nicelevel", "N", "", "set a nicelevel when starting")
	rootCmd.Flags().Bool("no-new-privs", false, "set the No New Privs flag for the program")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().String("notify", "", "configures experimental notification behaviour")
	rootCmd.Flags().String("oom-score-adj", "", "set OOM score adjustment when starting")
	rootCmd.Flags().StringP("pidfile", "p", "", "match pid found in this file")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().BoolP("reexec", "3", false, "reexec (used internally)")
	rootCmd.Flags().StringP("respawn-delay", "D", "", "set a respawn delay")
	rootCmd.Flags().String("respawn-delay-cap", "", "set maximum respawn delay when respawn-delay-step is also active")
	rootCmd.Flags().String("respawn-delay-step", "", "increase the respawn delay by this amount for every retry")
	rootCmd.Flags().StringP("respawn-max", "m", "", "set maximum number of respawn attempts")
	rootCmd.Flags().StringP("respawn-period", "P", "", "set respawn time period")
	rootCmd.Flags().StringP("retry", "R", "", "retry schedule to use when stopping")
	rootCmd.Flags().String("secbits", "", "set the security-bits for the program")
	rootCmd.Flags().StringP("signal", "s", "", "send a signal to the daemon")
	rootCmd.Flags().BoolP("start", "S", false, "start daemon")
	rootCmd.Flags().StringP("stderr", "2", "", "redirect stderr to file")
	rootCmd.Flags().String("stderr-logger", "", "redirect stderr to process")
	rootCmd.Flags().String("stdin", "", "redirect stdin to file")
	rootCmd.Flags().StringP("stdout", "1", "", "redirect stdout to file")
	rootCmd.Flags().String("stdout-logger", "", "redirect stdout to process")
	rootCmd.Flags().BoolP("stop", "K", false, "stop daemon")
	rootCmd.Flags().BoolP("stop-group", "G", false, "stop the whole process group")
	rootCmd.Flags().StringP("umask", "k", "", "set the umask for the daemon")
	rootCmd.Flags().StringP("user", "u", "", "change the process user")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"capabilities":  carapace.ActionFiles(),
		"chdir":         carapace.ActionDirectories(),
		"chroot":        carapace.ActionDirectories(),
		"ionice":        carapace.ActionValues("0", "1", "2", "3"),
		"nicelevel":     carapace.ActionValues("-20", "0", "19"),
		"oom-score-adj": carapace.ActionValues("-1000", "0", "1000"),
		"pidfile":       carapace.ActionFiles(),
		"retry":         carapace.ActionValues("1", "5", "10", "15", "30", "60"),
		"signal":        carapace.ActionValues("HUP", "INT", "TERM", "KILL", "USR1", "USR2", "QUIT", "ALRM"),
		"stderr":        carapace.ActionFiles(),
		"stderr-logger": carapace.ActionFiles(),
		"stdin":         carapace.ActionFiles(),
		"stdout":        carapace.ActionFiles(),
		"stdout-logger": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
