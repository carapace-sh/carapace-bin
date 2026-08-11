package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_config_getCmd).Standalone()

	help_repository_configCmd.AddCommand(help_repository_config_getCmd)
}
