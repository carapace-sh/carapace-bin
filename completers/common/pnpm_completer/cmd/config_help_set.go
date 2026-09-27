package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the config key to the value provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_setCmd).Standalone()

	config_helpCmd.AddCommand(config_help_setCmd)
}
