package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print all configuration values from your config.toml file in TOML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_printCmd).Standalone()

	help_configCmd.AddCommand(help_config_printCmd)
}
