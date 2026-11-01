package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openrc-run",
	Short: "run an OpenRC service script",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "set xtrace when running the script")
	rootCmd.Flags().BoolP("dry-run", "Z", false, "show what would be done")
	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().BoolP("ifstarted", "s", false, "only run commands when started")
	rootCmd.Flags().BoolP("ifstopped", "S", false, "only run commands when stopped")
	rootCmd.Flags().StringP("lockfd", "l", "", "fd of the exclusive lock from rc")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().BoolP("nodeps", "D", false, "ignore dependencies")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).PositionalCompletion(
		openrc.ActionServices(),
		openrc.ActionCommands(),
	)
}
