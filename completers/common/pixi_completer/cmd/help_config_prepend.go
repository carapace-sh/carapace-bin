package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "Prepend a value to a list configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_prependCmd).Standalone()

	help_configCmd.AddCommand(help_config_prependCmd)
}
