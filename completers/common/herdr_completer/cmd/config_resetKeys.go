package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_resetKeysCmd = &cobra.Command{
	Use:   "reset-keys",
	Short: "Reset custom keybindings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_resetKeysCmd).Standalone()

	configCmd.AddCommand(config_resetKeysCmd)
}
