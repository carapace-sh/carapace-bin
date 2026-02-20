package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append a value to a list configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_appendCmd).Standalone()

	config_appendCmd.Flags().BoolP("global", "g", false, "Operation on global configuration")
	config_appendCmd.Flags().BoolP("local", "l", false, "Operation on project-local configuration")
	config_appendCmd.Flags().BoolP("system", "s", false, "Operation on system configuration")
	configCmd.AddCommand(config_appendCmd)
}
