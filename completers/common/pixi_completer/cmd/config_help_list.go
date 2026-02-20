package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configuration values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_listCmd).Standalone()

	config_helpCmd.AddCommand(config_help_listCmd)
}
