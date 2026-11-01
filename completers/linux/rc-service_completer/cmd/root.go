package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-service",
	Short: "locate and run an OpenRC service",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "set xtrace when running the command")
	rootCmd.Flags().BoolP("dry-run", "Z", false, "dry run (show what would happen)")
	rootCmd.Flags().BoolP("exists", "e", false, "tests if the service exists or not")
	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().BoolP("ifcrashed", "c", false, "if the service is crashed run the command")
	rootCmd.Flags().BoolP("ifexists", "i", false, "if the service exists run the command")
	rootCmd.Flags().BoolP("ifinactive", "I", false, "if the service is inactive run the command")
	rootCmd.Flags().BoolP("ifnotstarted", "N", false, "if the service is not started run the command")
	rootCmd.Flags().BoolP("ifstarted", "s", false, "if the service is started run the command")
	rootCmd.Flags().BoolP("ifstopped", "S", false, "if the service is stopped run the command")
	rootCmd.Flags().BoolP("list", "l", false, "list all available services")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().BoolP("nodeps", "D", false, "ignore dependencies")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().StringP("resolve", "r", "", "resolve the service name to an init script")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"resolve": openrc.ActionServices(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		openrc.ActionServices(),
		openrc.ActionCommands(),
	)
}
