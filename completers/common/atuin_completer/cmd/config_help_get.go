package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration value from your config.toml file or after defaults and overrides are applied",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_getCmd).Standalone()

	config_helpCmd.AddCommand(config_help_getCmd)
}
