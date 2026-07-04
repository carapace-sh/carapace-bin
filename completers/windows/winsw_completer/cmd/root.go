package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "winsw",
	Short: "Windows Service Wrapper to run any executable as a Windows service",
	Long:  "https://github.com/winsw/winsw",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"install", "install the service",
			"uninstall", "uninstall the service",
			"start", "start the service",
			"stop", "stop the service",
			"restart", "restart the service",
			"status", "query the service status",
			"test", "run the executable in console mode",
		),
	)
}
