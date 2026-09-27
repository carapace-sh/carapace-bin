package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove the config key from the config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_deleteCmd).Standalone()

	help_configCmd.AddCommand(help_config_deleteCmd)
}
