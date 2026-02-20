package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "Prepend a value to a list configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_prependCmd).Standalone()

	config_prependCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_prependCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_prependCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	configCmd.AddCommand(config_prependCmd)
}
