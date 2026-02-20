package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes dependencies from an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_removeCmd).Standalone()

	global_helpCmd.AddCommand(global_help_removeCmd)
}
