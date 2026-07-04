package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nssm",
	Short: "Non-Sucking Service Manager for Windows",
	Long:  "https://nssm.cc/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"install", "install a service",
			"remove", "remove a service",
			"start", "start a service",
			"stop", "stop a service",
			"restart", "restart a service",
			"status", "query service status",
			"get", "get a service parameter",
			"set", "set a service parameter",
			"edit", "edit service parameters",
		),
	)
}
