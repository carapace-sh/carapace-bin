package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a feature, leaving its dependencies enabled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_disableCmd).Standalone()

	help_configCmd.AddCommand(help_config_disableCmd)
}
