package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append a value to a list configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_appendCmd).Standalone()

	config_helpCmd.AddCommand(config_help_appendCmd)
}
