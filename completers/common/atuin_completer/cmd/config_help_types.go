package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_typesCmd = &cobra.Command{
	Use:   "types",
	Short: "Print the Luau types of atuin.lua and .atuin.lua config files, for your editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_typesCmd).Standalone()

	config_helpCmd.AddCommand(config_help_typesCmd)
}
