package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the config value for the provided key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_getCmd).Standalone()

	help_configCmd.AddCommand(help_config_getCmd)
}
