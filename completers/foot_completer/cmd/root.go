package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "foot",
	Short: "A fast, lightweight and minimalistic Wayland terminal emulator",
	Long:  "https://codeberg.org/dnkl/foot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringP("app-id", "a", "", "window application ID (foot)")
	rootCmd.Flags().BoolP("check-config", "C", false, "verify configuration, exit with 0 if ok, otherwise exit with 1")
	rootCmd.Flags().StringP("config", "c", "", "load configuration from PATH ($XDG_CONFIG_HOME/foot/foot.ini)")
	rootCmd.Flags().BoolS("e", "e", false, "ignored (for compatibility with xterm -e)")
	rootCmd.Flags().StringP("font", "f", "", "comma separated list of fonts in fontconfig format (monospace)")
	rootCmd.Flags().BoolP("fullscreen", "F", false, "start in fullscreen mode")
	rootCmd.Flags().BoolP("hold", "H", false, "remain open after child process exits")
	rootCmd.Flags().StringP("log-colorize", "l", "", "enable/disable colorization of log output on stderr")
	rootCmd.Flags().StringP("log-level", "d", "", "log level (info)")
	rootCmd.Flags().Bool("log-no-syslog", false, "disable syslog logging (only applicable in server mode)")
	rootCmd.Flags().BoolP("login-shell", "L", false, "start shell as a login shell")
	rootCmd.Flags().BoolP("maximized", "m", false, "start in maximized mode")
	rootCmd.Flags().StringP("override", "o", "", "override configuration option")
	rootCmd.Flags().StringP("print-pid", "p", "", "print PID to file or FD (only applicable in server mode)")
	rootCmd.Flags().StringP("server", "s", "", "run as a server (use 'footclient' to start terminals).")
	rootCmd.Flags().StringP("term", "t", "", "value to set the environment variable TERM to (foot)")
	rootCmd.Flags().StringP("title", "T", "", "initial window title (foot)")
	rootCmd.Flags().BoolP("version", "v", false, "show the version number and quit")
	rootCmd.Flags().StringP("window-size-chars", "W", "", "initial width and height, in characters")
	rootCmd.Flags().StringP("window-size-pixels", "w", "", "initial width and height, in pixels")
	rootCmd.Flags().StringP("working-directory", "D", "", "directory to start in (CWD)")

	rootCmd.Flag("server").NoOptDefVal = " "
	rootCmd.Flag("log-colorize").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":            carapace.ActionFiles(),
		"font":              os.ActionFontFamilies(),
		"log-colorize":      carapace.ActionValues("never", "always", "auto").StyleF(style.ForKeyword),
		"log-level":         carapace.ActionValues("info", "warning", "error", "none").StyleF(style.ForLogLevel),
		"server":            carapace.ActionFiles(),
		"working-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
