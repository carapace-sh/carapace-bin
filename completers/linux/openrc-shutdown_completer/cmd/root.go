package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openrc-shutdown",
	Short: "shut down or reboot the system",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("cancel", "c", false, "cancel a pending shutdown")
	rootCmd.Flags().BoolP("dry-run", "D", false, "print actions instead of executing them")
	rootCmd.Flags().BoolP("halt", "H", false, "halt the system")
	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().BoolP("kexec", "K", false, "reboot the system using kexec")
	rootCmd.Flags().BoolP("no-write", "d", false, "do not write wtmp record")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().BoolP("poweroff", "p", false, "power off the system")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().BoolP("reboot", "r", false, "reboot the system")
	rootCmd.Flags().BoolP("reexec", "R", false, "re-execute init (use after upgrading)")
	rootCmd.Flags().BoolP("single", "s", false, "single user mode")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")
	rootCmd.Flags().BoolP("write-only", "w", false, "write wtmp boot record and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("now"),
	)
}
