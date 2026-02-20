package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds dependencies to an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_addCmd).Standalone()

	global_helpCmd.AddCommand(global_help_addCmd)
}
