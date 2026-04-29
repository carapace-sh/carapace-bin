package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration value from your config.toml file or after defaults and overrides are applied",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	config_getCmd.Flags().BoolP("help", "h", false, "Print help")
	config_getCmd.Flags().BoolP("resolved", "r", false, "Print the value after defaults and overrides are applied")
	config_getCmd.Flags().BoolP("verbose", "v", false, "Print both the config file value and the resolved value")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalCompletion(
		atuin.ActionConfigKeys(),
	)
}
