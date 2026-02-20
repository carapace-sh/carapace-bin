package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists global environments with their dependencies and exposed commands. Can also display all packages within a specific global environment when using the --environment flag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_listCmd).Standalone()

	global_helpCmd.AddCommand(global_help_listCmd)
}
