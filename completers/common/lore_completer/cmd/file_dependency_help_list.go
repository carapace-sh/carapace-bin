package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dependency_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List dependencies or dependents for files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dependency_help_listCmd).Standalone()

	file_dependency_helpCmd.AddCommand(file_dependency_help_listCmd)
}
