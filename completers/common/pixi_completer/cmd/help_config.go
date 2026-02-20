package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_configCmd).Standalone()

	helpCmd.AddCommand(help_configCmd)
}
