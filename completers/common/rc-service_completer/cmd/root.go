package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-service [options] [-l | -e service | service command]",
	Short: "locate and run an OpenRC service with the given arguments",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "Run service scripts with shell tracing")
	rootCmd.Flags().BoolP("dry-run", "Z", false, "Display commands rather than executing them")
	rootCmd.Flags().BoolP("exists", "e", false, "Return success if the service exists")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information and exit")
	rootCmd.Flags().BoolP("ifcrashed", "c", false, "Run only if the service has crashed")
	rootCmd.Flags().BoolP("ifexists", "i", false, "Return success if the service does not exist")
	rootCmd.Flags().BoolP("ifinactive", "I", false, "Run only if the service is inactive")
	rootCmd.Flags().BoolP("ifnotstarted", "N", false, "Run only if the service is not started")
	rootCmd.Flags().BoolP("ifstarted", "s", false, "Run only if the service is started")
	rootCmd.Flags().BoolP("ifstopped", "S", false, "Run only if the service is stopped")
	rootCmd.Flags().BoolP("list", "l", false, "List all available services")
	rootCmd.Flags().BoolP("nocolor", "C", false, "Disable color output")
	rootCmd.Flags().BoolP("nodeps", "D", false, "Ignore dependencies when running the service")
	rootCmd.Flags().BoolP("quiet", "q", false, "Run quietly")
	rootCmd.Flags().BoolP("resolve", "r", false, "Resolve the full path of the service")
	rootCmd.Flags().BoolP("user", "U", false, "Operate on user services and runlevels")
	rootCmd.Flags().BoolP("verbose", "v", false, "Run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "Display software version")

	carapace.Gen(rootCmd).PositionalCompletion(
		openrc.ActionServices(),
		openrc.ActionServiceCommands(),
	)
}
