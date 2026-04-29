package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var config_printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print all configuration values from your config.toml file in TOML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_printCmd).Standalone()

	config_printCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_printCmd)

	carapace.Gen(config_printCmd).PositionalCompletion(
		atuin.ActionConfigKeys(),
	)
}
