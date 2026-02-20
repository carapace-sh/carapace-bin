package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_unsetCmd).Standalone()

	help_configCmd.AddCommand(help_config_unsetCmd)
}
