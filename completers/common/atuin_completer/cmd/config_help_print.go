package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print all configuration values from your config.toml file in TOML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_printCmd).Standalone()

	config_helpCmd.AddCommand(config_help_printCmd)
}
