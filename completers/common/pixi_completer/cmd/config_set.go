package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_setCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_setCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	configCmd.AddCommand(config_setCmd)
}
