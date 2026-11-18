package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_typesCmd = &cobra.Command{
	Use:   "types",
	Short: "Print the Luau types of atuin.lua and .atuin.lua config files, for your editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_typesCmd).Standalone()

	help_configCmd.AddCommand(help_config_typesCmd)
}
