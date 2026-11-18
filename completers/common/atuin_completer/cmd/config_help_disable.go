package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a feature, leaving its dependencies enabled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_disableCmd).Standalone()

	config_helpCmd.AddCommand(config_help_disableCmd)
}
