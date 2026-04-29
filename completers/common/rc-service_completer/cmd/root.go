package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-service",
	Short: "locate and run OpenRC services",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "run the service script with shell tracing")
	rootCmd.Flags().BoolP("exists", "e", false, "test whether the service exists")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("ifcrashed", "C", false, "run command only if service crashed")
	rootCmd.Flags().BoolP("ifexists", "i", false, "run command only if service exists")
	rootCmd.Flags().BoolP("ifinactive", "N", false, "run command only if service is inactive")
	rootCmd.Flags().BoolP("ifnotstarted", "n", false, "run command only if service is not started")
	rootCmd.Flags().BoolP("ifstarted", "s", false, "run command only if service is started")
	rootCmd.Flags().BoolP("list", "l", false, "list all available services")
	rootCmd.Flags().BoolP("resolve", "r", false, "resolve the service path")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).PositionalCompletion(
		actionServices(),
		actionServiceCommands(),
	)
}

func actionServices() carapace.Action {
	return carapace.ActionExecCommand("rc-service", "--list")(func(output []byte) carapace.Action {
		return carapace.ActionValues(strings.Fields(string(output))...)
	})
}

func actionServiceCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"start", "start the service",
		"stop", "stop the service",
		"restart", "restart the service",
		"status", "show service status",
		"zap", "reset crashed service state",
	)
}
