package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_config_getCmd).Standalone()

	repository_help_configCmd.AddCommand(repository_help_config_getCmd)
}
