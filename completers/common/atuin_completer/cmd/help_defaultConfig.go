package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_defaultConfigCmd = &cobra.Command{
	Use:   "default-config",
	Short: "Print the default atuin configuration (config.toml)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_defaultConfigCmd).Standalone()

	helpCmd.AddCommand(help_defaultConfigCmd)
}
