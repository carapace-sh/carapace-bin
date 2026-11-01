package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove the config key from the config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_deleteCmd).Standalone()

	config_helpCmd.AddCommand(config_help_deleteCmd)
}
