package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-update",
	Short: "add and remove OpenRC services from runlevels",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "process all runlevels")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("stack", "s", false, "stack runlevels instead of replacing")
	rootCmd.Flags().BoolP("update", "u", false, "force dependency tree update")
	rootCmd.Flags().BoolP("verbose", "v", false, "show extra output")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "add a service to a runlevel",
			"del", "remove a service from a runlevel",
			"delete", "remove a service from a runlevel",
			"show", "show services and runlevels",
		),
		actionServices(),
		actionRunlevels(),
	)
}

func actionServices() carapace.Action {
	return carapace.ActionExecCommand("rc-service", "--list")(func(output []byte) carapace.Action {
		return carapace.ActionValues(strings.Fields(string(output))...)
	})
}

func actionRunlevels() carapace.Action {
	return carapace.ActionValues("boot", "default", "nonetwork", "shutdown", "single", "sysinit")
}
