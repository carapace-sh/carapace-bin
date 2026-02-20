package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_unsetCmd).Standalone()

	config_unsetCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_unsetCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_unsetCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	configCmd.AddCommand(config_unsetCmd)
}
