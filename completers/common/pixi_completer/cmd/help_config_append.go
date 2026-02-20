package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append a value to a list configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_appendCmd).Standalone()

	help_configCmd.AddCommand(help_config_appendCmd)
}
