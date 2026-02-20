package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_editCmd).Standalone()

	help_configCmd.AddCommand(help_config_editCmd)
}
