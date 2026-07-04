package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "launchctl",
	Short: "control service instances",
	Long:  "https://keith.github.io/xcode-manpages/launchctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("domain", "D", "", "Apply the command to the specified domain")
	rootCmd.Flags().BoolP("help", "h", false, "Display a brief usage message and exit")
	rootCmd.Flags().BoolP("list", "l", false, "List information about services")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress informational output")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase the verbosity of informational output")
	rootCmd.Flags().BoolP("version", "V", false, "Print the version of launchctl and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"bootstrap", "Bootstraps a service into the specified domain",
			"bootout", "Unbootstraps a service from the specified domain",
			"enable", "Enables a service in the specified domain",
			"disable", "Disables a service in the specified domain",
			"kickstart", "Forces a service to start in the specified domain",
			"kill", "Sends a signal to the service",
			"list", "Lists information about services",
			"load", "Load a service for execution",
			"unload", "Unload a service",
			"start", "Start a service",
			"stop", "Stop a service",
			"print", "Prints information about the service",
			"print-disabled", "Prints whether a service is disabled",
			"print-system", "Prints the system domain",
			"print-user", "Prints the user domain",
			"plist", "Prints the property list for a service",
			"setenv", "Sets an environment variable in the specified domain",
			"unsetenv", "Unsets an environment variable in the specified domain",
			"limit", "Sets or gets resource limits",
			"hostinfo", "Prints host information",
			"userinfo", "Prints user information",
			"service", "Look up a service name",
		),
	)
}
