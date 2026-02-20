package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates environments in the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_updateCmd).Standalone()

	global_helpCmd.AddCommand(global_help_updateCmd)
}
