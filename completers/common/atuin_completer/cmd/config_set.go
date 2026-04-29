package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value in your config.toml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_setCmd.Flags().StringP("type", "t", "", "Store value as an explicit type")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("auto", "string", "boolean", "integer", "float"),
	})

	carapace.Gen(config_setCmd).PositionalCompletion(
		atuin.ActionConfigKeys(),
	)
}
