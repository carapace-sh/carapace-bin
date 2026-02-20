package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists global environments with their dependencies and exposed commands. Can also display all packages within a specific global environment when using the --environment flag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_listCmd).Standalone()

	help_globalCmd.AddCommand(help_global_listCmd)
}
