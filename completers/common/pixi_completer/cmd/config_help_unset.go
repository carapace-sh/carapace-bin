package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_unsetCmd).Standalone()

	config_helpCmd.AddCommand(config_help_unsetCmd)
}
